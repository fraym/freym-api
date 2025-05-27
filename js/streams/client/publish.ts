import {
    EventPayload,
    PublishEvent as ProtobufPublishEvent,
    ServiceClient,
} from "@fraym/proto/dist/index.freym.streams.management";
import { PublishEvent, isGdprEventData } from "./event";
import { retry } from "./util";

export const sendPublish = async (
    topic: string,
    events: PublishEvent[],
    deploymentId: number,
    serviceClient: ServiceClient
) => {
    return retry(
        () =>
            new Promise<void>((resolve, reject) => {
                serviceClient.publish(
                    {
                        events: events.map(event =>
                            getProtobufPublishEventFromPublishedEvent(event, deploymentId)
                        ),
                        topic,
                    },
                    error => {
                        if (error) {
                            reject(error);
                            return;
                        }

                        resolve();
                    }
                );
            })
    );
};

export const getProtobufPublishEventFromPublishedEvent = (
    event: PublishEvent,
    deploymentId: number
): ProtobufPublishEvent => {
    const payload: Record<string, EventPayload> = {};

    for (const key in event.payload) {
        const currentData = event.payload[key];

        payload[key] = isGdprEventData(currentData)
            ? {
                  value: JSON.stringify(currentData.value),
                  gdpr: {
                      default: JSON.stringify(currentData.gdprDefault),
                      id: currentData.id ?? "",
                      isInvalidated: false,
                  },
              }
            : {
                  value: JSON.stringify(currentData),
                  gdpr: undefined,
              };
    }

    return {
        id: event.id,
        metadata: {
            causationId: event.causationId ?? "",
            correlationId: event.correlationId ?? "",
            orderSerial: "0",
            deploymentId: deploymentId.toString(),
            userId: event.userId ?? "",
        },
        options: {
            broadcast: event.broadcast ?? false,
        },
        reason: event.reason ?? "",
        stream: event.stream ?? "",
        tenantId: event.tenantId,
        type: event.type ?? "",
        payload,
    };
};
