import { ServiceClient } from "@fraym/proto/dist/index.freym.auth.management";
import { PaginatedResponse } from "./paginatedResponse";

export interface Role {
    id: string;
    allowedScopes: RoleScope[];
}

export interface RoleScope {
    clientId: string;
    scopeName: string;
}

export const getAllRoles = async (
    tenantId: string,
    limit: number,
    page: number,
    serviceClient: ServiceClient
): Promise<PaginatedResponse<Role>> => {
    return new Promise<PaginatedResponse<Role>>((resolve, reject) => {
        serviceClient.getRoles(
            {
                tenantId,
                limit: limit.toString(),
                page: page.toString(),
            },
            (error, response) => {
                if (error) {
                    reject(error.message);
                    return;
                }

                resolve({
                    limit: parseInt(response.limit, 10),
                    page: parseInt(response.page, 10),
                    total: parseInt(response.total, 10),
                    data: response.roles,
                });
            }
        );
    });
};
