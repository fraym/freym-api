import { ServiceClient } from "@fraym/proto/dist/index.freym.auth.management";
import { EventMetadata, fillMetadataWithDefaults } from "./eventMetadata";

export interface CreateUserResponse {
    id: string;
    setInitialPasswordToken: string;
}

export const createNewUser = async (
    tenantId: string,
    login: string,
    email: string,
    displayName: string,
    password: string,
    assignedRoleIds: string[],
    active: boolean,
    blockedUntil: Date,
    eventMetadata: Partial<EventMetadata> | null,
    serviceClient: ServiceClient
): Promise<CreateUserResponse> => {
    return new Promise<CreateUserResponse>((resolve, reject) => {
        serviceClient.createUser(
            {
                tenantId,
                login,
                email,
                displayName,
                password,
                active,
                assignedRoleIds,
                blockedUntil: blockedUntil.getTime().toString(),
                eventMetadata: fillMetadataWithDefaults(eventMetadata),
            },
            (error, response) => {
                if (error) {
                    reject(error.message);
                    return;
                }

                resolve({
                    id: response.id,
                    setInitialPasswordToken: response.setInitialPasswordToken,
                });
            }
        );
    });
};
