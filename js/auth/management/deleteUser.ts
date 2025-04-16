import { ServiceClient } from "@fraym/proto/dist/index.freym.auth.management";
import { Metadata, fillMetadataWithDefaults } from "./metadata";

export const deleteExistingUser = async (
    tenantId: string,
    id: string,
    metadata: Partial<Metadata> | null,
    serviceClient: ServiceClient
): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        serviceClient.deleteUser(
            {
                tenantId,
                id,
                metadata: fillMetadataWithDefaults(metadata),
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
