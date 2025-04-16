import { ServiceClient } from "@fraym/proto/dist/index.freym.auth.management";
import { Metadata, fillMetadataWithDefaults } from "./metadata";

export const updateExistingUser = async (
    tenantId: string,
    id: string,
    login: string,
    email: string,
    displayName: string,
    password: string,
    assignedRoleIds: string[],
    active: boolean,
    blockedUntil: Date,
    metadata: Partial<Metadata> | null,
    serviceClient: ServiceClient
): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        serviceClient.updateUser(
            {
                tenantId,
                id,
                login,
                email,
                displayName,
                password,
                active,
                assignedRoleIds,
                blockedUntil: blockedUntil.getTime().toString(),
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
