import { ServiceClient } from "@fraym/proto/dist/index.freym.auth.management";
import { PaginatedResponse } from "./paginatedResponse";

export interface User {
    id: string;
    login: string;
    email: string;
    displayName: string;
    assignedRoleIds: string[];
    active: boolean;
    failedAttempts: number;
    lastAttempt: number;
    blockedUntil: number;
}

export const getAllUsers = async (
    tenantId: string,
    limit: number,
    page: number,
    serviceClient: ServiceClient
): Promise<PaginatedResponse<User>> => {
    return new Promise<PaginatedResponse<User>>((resolve, reject) => {
        serviceClient.getUsers(
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
                    data: response.users.map(user => {
                        return {
                            active: user.active,
                            assignedRoleIds: user.assignedRoleIds,
                            blockedUntil: parseInt(user.blockedUntil),
                            displayName: user.displayName,
                            email: user.email,
                            failedAttempts: parseInt(user.failedAttempts),
                            id: user.id,
                            lastAttempt: parseInt(user.lastAttempt),
                            login: user.login,
                        };
                    }),
                });
            }
        );
    });
};
