import { ServiceClient } from "@fraym/proto/dist/index.freym.auth.management";
import { EventMetadata, fillMetadataWithDefaults } from "./eventMetadata";

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
    eventMetadata: Partial<EventMetadata> | null,
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
