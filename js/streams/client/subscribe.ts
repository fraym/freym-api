import {
    ServiceClient,
    SubscribeRequest,
    SubscribeResponse,
} from "@fraym/proto/dist/index.freym.streams.management";
import { v4 as uuid } from "uuid";
import { ClientDuplexStream } from "@grpc/grpc-js";
import { Status } from "@grpc/grpc-js/build/src/constants";
import { ClientConfig } from "./config";
import { isRetryableError } from "./error";
import { HandlerFunc, getSubscriptionEvent } from "./event";

export interface Subscription {
    useHandler: (type: string, handler: HandlerFunc) => void;
    useHandlerForAllTypes: (handler: HandlerFunc) => void;
    start: () => void;
    stop: () => void;
}

export type Stream = ClientDuplexStream<SubscribeRequest, SubscribeResponse>;

export const newSubscription = (
    topics: string[],
    ignoreUnhandledEvents: boolean,
    parallelTopicProcessing: boolean,
    config: Required<ClientConfig>,
    serviceClient: ServiceClient
): Subscription => {
    let stream: Stream | null = null;
    let closed = false;
    const typeHandlerMap: Record<string, HandlerFunc> = {};
    let globalHandler: HandlerFunc | null = null;

    const rebuildConnection = (currentStream: Stream, retries: number) => {
        currentStream.cancel();
        currentStream.removeAllListeners();

        setTimeout(() => {
            stream = null;
            reconnect(retries);
        }, 100);
    };

    const reconnect = async (retries: number) => {
        const newStream = serviceClient.subscribe();
        newStream.on("end", () => {
            if (closed) {
                newStream.cancel();
                return;
            }

            rebuildConnection(newStream, retries - 1);
        });
        /* eslint-disable-next-line @typescript-eslint/no-explicit-any */
        newStream.on("error", (err: any) => {
            if (closed) {
                return;
            }

            if (retries === 0 || (err && err.code && err.code === Status.UNKNOWN)) {
                closed = true;
                throw err;
            }

            rebuildConnection(newStream, retries - 1);
        });

        const dataFn = async (data: SubscribeResponse) => {
            if (
                !data.payload ||
                data.payload?.$case === "panic" ||
                data.payload?.$case === "subscribed"
            ) {
                newStream.cancel();
                return;
            }

            const event = getSubscriptionEvent(data.payload.event);
            if (!event) {
                return;
            }

            if (!event.type) {
                return;
            }

            try {
                const typedHandler = typeHandlerMap[event.type];

                if (typedHandler) {
                    await typedHandler(event);
                } else if (globalHandler) {
                    await globalHandler(event);
                } else {
                    if (ignoreUnhandledEvents) {
                        newStream.write(
                            newHandledRequest(event.tenantId, event.topic, event.stream)
                        );
                        return;
                    }

                    newStream.write(
                        newHandledRequest(
                            event.tenantId,
                            event.topic,
                            event.stream,
                            "no handlers for this event, maybe you forgot to register an event handler",
                            true
                        )
                    );
                    return;
                }

                newStream.write(newHandledRequest(event.tenantId, event.topic));
            } catch (err) {
                const shouldRetry = isRetryableError(err) && err.shouldRetry;
                const errorMessage = err instanceof Error ? err.message : String(err);

                newStream.write(
                    newHandledRequest(
                        event.tenantId,
                        event.topic,
                        event.stream,
                        errorMessage,
                        shouldRetry
                    )
                );

                throw err;
            }
        };
        stream = newStream;

        await initStream(topics, config, parallelTopicProcessing, newStream);

        retries = 50;
        newStream.on("data", dataFn);
    };

    return {
        useHandler: (type: string, handler: HandlerFunc) => {
            typeHandlerMap[type] = handler;
        },
        useHandlerForAllTypes: (handler: HandlerFunc) => {
            globalHandler = handler;
        },
        start: () => {
            reconnect(50);
        },
        stop: () => {
            if (stream) {
                stream.cancel();
                stream = null;
            }

            closed = true;
        },
    };
};

export const initStream = async (
    topics: string[],
    config: Required<ClientConfig>,
    parallelTopicProcessing: boolean,
    stream: Stream
): Promise<Stream> => {
    return new Promise<Stream>((resolve, reject) => {
        stream.write({
            payload: {
                $case: "subscribe",
                subscribe: {
                    metadata: {
                        group: config.groupId,
                        subscriberId: uuid(),
                        deploymentId: config.deploymentId.toString(),
                        parallelTopicProcessing,
                    },
                    topics,
                },
            },
        });

        stream.once("data", (data: SubscribeResponse) => {
            if (data.payload?.$case !== "subscribed") {
                reject("connection to streams service was not initialized correctly");
                return;
            }

            if (data.payload.subscribed.error) {
                reject(`unable to subscribe to streams service: ${data.payload.subscribed.error}`);
                return;
            }

            resolve(stream);
        });
    });
};

const newHandledRequest = (
    tenantId: string,
    topic: string,
    stream?: string,
    error?: string,
    retry?: boolean
): SubscribeRequest => {
    return {
        payload: {
            $case: "handled",
            handled: {
                tenantId,
                topic,
                stream: stream ?? "",
                error: error ?? "",
                retry: retry ?? false,
            },
        },
    };
};
