import { ServiceClient } from "@fraym/proto/dist/index.freym.streams.management";
import { SubscriptionEvent, getSubscriptionEvent } from "./event";
import { retry } from "./util";

export const getLastHandledEvent = async (
    tenantId: string,
    topic: string,
    group: string,
    parallelTopicProcessing: boolean,
    serviceClient: ServiceClient
) => {
    return retry(
        () =>
            new Promise<SubscriptionEvent | null>((resolve, reject) => {
                serviceClient.getLastHandledEvent(
                    {
                        tenantId,
                        topic,
                        group,
                        parallelTopicProcessing,
                    },
                    (error, response) => {
                        if (error?.details.includes("unable to find last handled event")) {
                            resolve(null);
                            return;
                        }

                        if (error) {
                            reject(error);
                            return;
                        }
                        const event = getSubscriptionEvent(response);

                        if (event) {
                            resolve(event);
                            return;
                        }

                        reject("unable to resolve last handled event from event data");
                    }
                );
            })
    );
};
