import { ServiceClient } from "@fraym/proto/dist/index.freym.auth.management";
import { ClientConfig, useConfigDefaults } from "@/config/config";
import { credentials } from "@grpc/grpc-js";
import { CreateUserResponse, createNewUser } from "./createUser";
import { deleteExistingRole } from "./deleteRole";
import { deleteExistingUser } from "./deleteUser";
import { Role, getAllRoles } from "./getRoles";
import { User, getAllUsers } from "./getUsers";
import { Metadata } from "./metadata";
import { PaginatedResponse } from "./paginatedResponse";
import { updateExistingUser } from "./updateUser";
import { UpsertRoleScope, createOrUpdateRole } from "./upsertRole";

export interface ManagementClient {
    upsertRole: (
        tenantId: string,
        allowedScopes: UpsertRoleScope[],
        id?: string,
        eventMetadata?: Partial<Metadata>
    ) => Promise<string>;
    deleteRole: (tenantId: string, id: string, eventMetadata?: Partial<Metadata>) => Promise<void>;
    getRoles: (tenantId: string, limit?: number, page?: number) => Promise<PaginatedResponse<Role>>;
    createUser: (
        tenantId: string,
        email: string,
        assignedRoleIds: string[],
        login?: string,
        displayName?: string,
        password?: string,
        active?: boolean,
        blockedUntil?: Date,
        eventMetadata?: Partial<Metadata>
    ) => Promise<CreateUserResponse>;
    updateUser: (
        tenantId: string,
        id: string,
        email: string,
        assignedRoleIds: string[],
        login?: string,
        displayName?: string,
        password?: string,
        active?: boolean,
        blockedUntil?: Date,
        eventMetadata?: Partial<Metadata>
    ) => Promise<void>;
    deleteUser: (tenantId: string, id: string, eventMetadata?: Partial<Metadata>) => Promise<void>;
    getUsers: (tenantId: string, limit?: number, page?: number) => Promise<PaginatedResponse<User>>;
    close: () => Promise<void>;
}

export const newManagementClient = async (config?: ClientConfig): Promise<ManagementClient> => {
    const currentConfig = useConfigDefaults(config);
    const serviceClient = new ServiceClient(
        currentConfig.serverAddress,
        credentials.createInsecure(),
        {
            "grpc.keepalive_time_ms": currentConfig.keepaliveInterval,
            "grpc.keepalive_timeout_ms": currentConfig.keepaliveTimeout,
            "grpc.keepalive_permit_without_calls": 1,
        }
    );

    const upsertRole = async (
        tenantId: string,
        allowedScopes: UpsertRoleScope[],
        id: string = "",
        eventMetadata: Partial<Metadata> | null = null
    ) => {
        return await createOrUpdateRole(tenantId, id, allowedScopes, eventMetadata, serviceClient);
    };

    const deleteRole = async (
        tenantId: string,
        id: string,
        eventMetadata: Partial<Metadata> | null = null
    ) => {
        return await deleteExistingRole(tenantId, id, eventMetadata, serviceClient);
    };

    const getRoles = async (tenantId: string, limit: number = 0, page: number = 1) => {
        return await getAllRoles(tenantId, limit, page, serviceClient);
    };

    const createUser = async (
        tenantId: string,
        email: string,
        assignedRoleIds: string[],
        login: string = "",
        displayName: string = "",
        password: string = "",
        active: boolean = false,
        blockedUntil: Date = new Date(0),
        eventMetadata: Partial<Metadata> | null = null
    ) => {
        return await createNewUser(
            tenantId,
            login,
            email,
            displayName,
            password,
            assignedRoleIds,
            active,
            blockedUntil,
            eventMetadata,
            serviceClient
        );
    };

    const updateUser = async (
        tenantId: string,
        id: string,
        email: string,
        assignedRoleIds: string[],
        login: string = "",
        displayName: string = "",
        password: string = "",
        active: boolean = false,
        blockedUntil: Date = new Date(0),
        eventMetadata: Partial<Metadata> | null = null
    ) => {
        return await updateExistingUser(
            tenantId,
            id,
            login,
            email,
            displayName,
            password,
            assignedRoleIds,
            active,
            blockedUntil,
            eventMetadata,
            serviceClient
        );
    };

    const deleteUser = async (
        tenantId: string,
        id: string,
        eventMetadata: Partial<Metadata> | null = null
    ) => {
        return await deleteExistingUser(tenantId, id, eventMetadata, serviceClient);
    };

    const getUsers = async (tenantId: string, limit: number = 0, page: number = 1) => {
        return await getAllUsers(tenantId, limit, page, serviceClient);
    };

    const close = async () => {
        serviceClient.close();
    };

    return {
        upsertRole,
        deleteRole,
        getRoles,
        createUser,
        updateUser,
        deleteUser,
        getUsers,
        close,
    };
};
