import { ServiceClient } from "@fraym/proto/dist/index.freym.auth.management";
import { EventMetadata, fillMetadataWithDefaults } from "./eventMetadata";

export interface UpsertRoleScope {
    scopeName: string;
    clientId?: string;
}

export const createOrUpdateRole = async (
    tenantId: string,
    id: string,
    allowedScopes: UpsertRoleScope[],
    eventMetadata: Partial<EventMetadata> | null,
    serviceClient: ServiceClient
): Promise<string> => {
    return new Promise<string>((resolve, reject) => {
        serviceClient.upsertRole(
            {
                tenantId,
                id,
                allowedScopes: allowedScopes.map(scope => {
                    return {
                        scopeName: scope.scopeName,
                        clientId: scope.clientId ?? "",
                    };
                }),
                eventMetadata: fillMetadataWithDefaults(eventMetadata),
            },
            (error, response) => {
                if (error) {
                    reject(error.message);
                    return;
                }

                resolve(response.id);
            }
        );
    });
};
