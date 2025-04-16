import { ServiceClient } from "@fraym/proto/dist/index.freym.auth.management";
import { Metadata, fillMetadataWithDefaults } from "./metadata";

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
    metadata: Partial<Metadata> | null,
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
                metadata: fillMetadataWithDefaults(metadata),
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
