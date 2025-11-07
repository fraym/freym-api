import { ServiceClient } from "@fraym/proto/dist/index.freym.sync.management";
import { credentials } from "@grpc/grpc-js";
import { ClientConfig, useConfigDefaults } from "./config";
import { newConnection } from "./connection";
import { createLease } from "./lease";
import { lock } from "./lock";
import { newRequestGate } from "./requestGate";
import { rLock } from "./rlock";
import { rUnlock } from "./runlock";
import { tryLock } from "./tryLock";
import { tryRLock } from "./tryRLock";
import { unlock } from "./unlock";

export interface Client {
    lock: (tenantId: string, ...resource: string[]) => Promise<void>;
    tryLock: (tenantId: string, ...resource: string[]) => Promise<boolean>;
    unlock: (tenantId: string, ...resource: string[]) => Promise<void>;
    rLock: (tenantId: string, ...resource: string[]) => Promise<void>;
    tryRLock: (tenantId: string, ...resource: string[]) => Promise<boolean>;
    rUnlock: (tenantId: string, ...resource: string[]) => Promise<void>;
    close: () => Promise<void>;
}

export const newClient = async (config: ClientConfig): Promise<Client> => {
    config = useConfigDefaults(config);
    const serviceClient = new ServiceClient(config.serverAddress, credentials.createInsecure(), {
        "grpc.keepalive_time_ms": config.keepaliveInterval,
        "grpc.keepalive_timeout_ms": config.keepaliveTimeout,
        "grpc.keepalive_permit_without_calls": 1,
        "grpc.max_receive_message_length": 2147483647,
    });

    const lockRequestGate = newRequestGate(900);
    const unlockRequestGate = newRequestGate(80);

    const connection = await newConnection();
    const lease = await createLease(connection, config, serviceClient);

    await lease.renew();

    return {
        lock: async (tenantId: string, ...resource: string[]) => {
            await connection.waitForConnect();
            await lock(lease, tenantId, resource, serviceClient, lockRequestGate);
        },
        tryLock: async (tenantId: string, ...resource: string[]) => {
            await connection.waitForConnect();
            return await tryLock(lease, tenantId, resource, serviceClient, lockRequestGate);
        },
        unlock: async (tenantId: string, ...resource: string[]) => {
            await connection.waitForConnect();
            await unlock(lease, tenantId, resource, serviceClient, unlockRequestGate);
        },
        rLock: async (tenantId: string, ...resource: string[]) => {
            await connection.waitForConnect();
            await rLock(lease, tenantId, resource, serviceClient, lockRequestGate);
        },
        tryRLock: async (tenantId: string, ...resource: string[]) => {
            await connection.waitForConnect();
            return await tryRLock(lease, tenantId, resource, serviceClient, lockRequestGate);
        },
        rUnlock: async (tenantId: string, ...resource: string[]) => {
            await connection.waitForConnect();
            await rUnlock(lease, tenantId, resource, serviceClient, unlockRequestGate);
        },
        close: async () => {
            connection.stop();
            lease.stop();
            await lease.waitForStop();
        },
    };
};
