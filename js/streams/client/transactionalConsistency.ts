import { ServiceClient } from "@fraym/proto/dist/index.freym.streams.management";
import { retry } from "./util";

export const waitForTransactionalConsistency = async (
    tenantId: string,
    topic: string,
    correlationId: string,
    consumerGroups: string[],
    parallelTopicProcessing: boolean,
    serviceClient: ServiceClient
) => {
    return retry(
        () =>
            new Promise<void>((resolve, reject) => {
                serviceClient.waitForTransactionalConsistency(
                    {
                        tenantId,
                        topic,
                        correlationId,
                        consumerGroups,
                        parallelTopicProcessing,
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
