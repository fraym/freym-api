import { ServiceClient } from "@fraym/proto/dist/index.freym.sync.management";
import { credentials } from "@grpc/grpc-js";
import { ClientConfig, useConfigDefaults } from "./config";
import { newConnection } from "./connection";
import { createLease } from "./lease";
import { lock } from "./lock";
import { rLock } from "./rlock";
import { rUnlock } from "./runlock";
import { unlock } from "./unlock";

export interface Client {
    lock: (tenantId: string, ...resource: string[]) => Promise<void>;
    unlock: (tenantId: string, ...resource: string[]) => Promise<void>;
    rLock: (tenantId: string, ...resource: string[]) => Promise<void>;
    rUnlock: (tenantId: string, ...resource: string[]) => Promise<void>;
    stop: () => Promise<void>;
}

export const newClient = async (config: ClientConfig): Promise<Client> => {
    config = useConfigDefaults(config);
    const serviceClient = new ServiceClient(config.serverAddress, credentials.createInsecure(), {
        "grpc.keepalive_time_ms": config.keepaliveInterval,
        "grpc.keepalive_timeout_ms": config.keepaliveTimeout,
        "grpc.keepalive_permit_without_calls": 1,
        "grpc.max_receive_message_length": 2147483647,
    });

    const connection = await newConnection();
    const lease = await createLease(connection, config, serviceClient);

    await lease.renew();

    return {
        lock: async (tenantId: string, ...resource: string[]) => {
            await connection.waitForConnect();
            await lock(lease, tenantId, resource, serviceClient);
        },
        unlock: async (tenantId: string, ...resource: string[]) => {
            await connection.waitForConnect();
            await unlock(lease, tenantId, resource, serviceClient);
        },
        rLock: async (tenantId: string, ...resource: string[]) => {
            await connection.waitForConnect();
            await rLock(lease, tenantId, resource, serviceClient);
        },
        rUnlock: async (tenantId: string, ...resource: string[]) => {
            await connection.waitForConnect();
            await rUnlock(lease, tenantId, resource, serviceClient);
        },
        stop: async () => {
            connection.stop();
            lease.stop();
            await lease.waitForStop();
        },
    };
};
