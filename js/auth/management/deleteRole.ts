import { ServiceClient } from "@fraym/proto/dist/index.freym.auth.management";
import { Metadata, fillMetadataWithDefaults } from "./metadata";

export const deleteExistingRole = async (
    tenantId: string,
    id: string,
    metadata: Partial<Metadata> | null,
    serviceClient: ServiceClient
): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        serviceClient.deleteRole(
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
