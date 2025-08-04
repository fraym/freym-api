import { ServiceClient } from "@fraym/proto/dist/index.freym.streams.management";
import { ErroneousEvent, getErroneousEvent } from "./event";
import { retry } from "./util";

export const listErroneousEvents = async (
    tenantId: string,
    topic: string,
    eventTypes: string[],
    limit: number,
    serviceClient: ServiceClient
) => {
    return retry(
        () =>
            new Promise<ErroneousEvent[]>((resolve, reject) => {
                serviceClient.listErroneousEvents(
                    {
                        tenantId,
                        topic,
                        eventTypes,
                        limit: limit.toString(),
                    },
                    async (error, data) => {
                        if (error) {
                            reject(error);
                            return;
                        }

                        const erroneousEvents: ErroneousEvent[] = [];

                        for (const event of data.events) {
                            const erroneousEvent = getErroneousEvent(event);

                            if (erroneousEvent) {
                                erroneousEvents.push(erroneousEvent);
                            }
                        }

                        resolve(erroneousEvents);
                    }
                );
            })
    );
};

export const resendErroneousEvent = async (
    tenantId: string,
    topic: string,
    consumerGroup: string,
    eventId: string,
    serviceClient: ServiceClient
) => {
    return retry(
        () =>
            new Promise<void>((resolve, reject) => {
                serviceClient.resendErroneousEvent(
                    {
                        tenantId,
                        topic,
                        consumerGroup,
                        eventId,
                    },
                    async error => {
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
