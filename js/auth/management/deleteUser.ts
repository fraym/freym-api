import { ServiceClient } from "@fraym/proto/dist/index.freym.auth.management";
import { EventMetadata, fillMetadataWithDefaults } from "./eventMetadata";

export const deleteExistingUser = async (
    tenantId: string,
    id: string,
    eventMetadata: Partial<EventMetadata> | null,
    serviceClient: ServiceClient
): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        serviceClient.deleteUser(
            {
                tenantId,
                id,
                eventMetadata: fillMetadataWithDefaults(eventMetadata),
            },
            error => {
                if (error) {
                    reject(error.message);
                    return;
                }

                resolve();
            }
        );
    });
};
