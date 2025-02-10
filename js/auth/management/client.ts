import { ServiceClient } from "@fraym/proto/dist/index.freym.auth.management";
import { ClientConfig, useConfigDefaults } from "@/config/config";
import { credentials } from "@grpc/grpc-js";
import { CreateUserResponse, createNewUser } from "./createUser";
import { deleteExistingRole } from "./deleteRole";
import { deleteExistingUser } from "./deleteUser";
import { EventMetadata } from "./eventMetadata";
import { Role, getAllRoles } from "./getRoles";
import { User, getAllUsers } from "./getUsers";
import { updateExistingUser } from "./updateUser";
import { UpsertRoleScope, createOrUpdateRole } from "./upsertRole";

export interface ManagementClient {
    upsertRole: (
        tenantId: string,
        allowedScopes: UpsertRoleScope[],
        id?: string,
        eventMetadata?: Partial<EventMetadata>
    ) => Promise<string>;
    deleteRole: (
        tenantId: string,
        id: string,
        eventMetadata?: Partial<EventMetadata>
    ) => Promise<void>;
    getRoles: (tenantId: string) => Promise<Role[]>;
    createUser: (
        tenantId: string,
        email: string,
        assignedRoleIds: string[],
        login?: string,
        displayName?: string,
        password?: string,
        active?: boolean,
        blockedUntil?: Date,
        eventMetadata?: Partial<EventMetadata>
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
        eventMetadata?: Partial<EventMetadata>
    ) => Promise<void>;
    deleteUser: (
        tenantId: string,
        id: string,
        eventMetadata?: Partial<EventMetadata>
    ) => Promise<void>;
    getUsers: (tenantId: string) => Promise<User[]>;
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
        eventMetadata: Partial<EventMetadata> | null = null
    ) => {
        return await createOrUpdateRole(tenantId, id, allowedScopes, eventMetadata, serviceClient);
    };

    const deleteRole = async (
        tenantId: string,
        id: string,
        eventMetadata: Partial<EventMetadata> | null = null
    ) => {
        return await deleteExistingRole(tenantId, id, eventMetadata, serviceClient);
    };

    const getRoles = async (tenantId: string) => {
        return await getAllRoles(tenantId, serviceClient);
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
        eventMetadata: Partial<EventMetadata> | null = null
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
        eventMetadata: Partial<EventMetadata> | null = null
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
        eventMetadata: Partial<EventMetadata> | null = null
    ) => {
        return await deleteExistingUser(tenantId, id, eventMetadata, serviceClient);
    };

    const getUsers = async (tenantId: string) => {
        return await getAllUsers(tenantId, serviceClient);
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
