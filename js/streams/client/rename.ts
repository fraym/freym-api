import { ServiceClient } from "@fraym/proto/dist/index.freym.streams.management";
import { retry } from "./util";

export const renameEventType = async (
    topic: string,
    oldEventType: string,
    newEventType: string,
    serviceClient: ServiceClient
) => {
    return retry(
        () =>
            new Promise<void>((resolve, reject) => {
                serviceClient.renameEventType(
                    {
                        topic,
                        newType: newEventType,
                        oldType: oldEventType,
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
