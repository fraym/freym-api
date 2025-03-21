import { Lock as PbLock, ServiceClient } from "@fraym/proto/dist/index.freym.sync.management";
import { E_CANCELED, Mutex } from "async-mutex";
import { ServiceError } from "@grpc/grpc-js";
import { ClientConfig } from "./config";
import { Connection } from "./connection";
import { getOwnIpAddress } from "./ip";
import { createResolvablePromise, racePromises, wait } from "./promise";
import { retry } from "./retry";

export interface Lease {
    stop: () => void;
    runWithLeaseId: (callback: (leaseId: string) => Promise<void>) => Promise<void>;
    waitForStop: () => Promise<void>;
    renew: () => Promise<void>;
    track: (tenant: string, resource: string[], read: boolean) => Promise<void>;
    untrack: (tenant: string, resource: string[], read: boolean) => Promise<void>;
}

interface Lock {
    tenant: string;
    resource: string[];
}

type CreateLeaseResponse =
    | {
          error: null;
          leaseId: string;
      }
    | {
          error: ServiceError;
      };

export const createLease = async (
    connection: Connection,
    config: ClientConfig,
    serviceClient: ServiceClient
): Promise<Lease> => {
    const mutex = new Mutex();
    const ttl = 20;
    let locks: Lock[] = [];
    let rLocks: Lock[] = [];
    let leaseId = "";

    const [waitForStopExecution, stopExecution] = await createResolvablePromise<void>();
    const [waitForStop, stopped] = await createResolvablePromise<void>();

    const renew = async () => {
        const [waitForCreateLease, onCreateLease] =
            await createResolvablePromise<CreateLeaseResponse>();

        try {
            const release = await mutex.acquire();
            serviceClient.createLease(
                {
                    ip: getOwnIpAddress(),
                    app: config.appPrefix,
                    ttl,
                    leaseId,
                    alreadyRegisteredLocks: getPbLocks(leaseId, locks),
                    alreadyRegisteredRlocks: getPbLocks(leaseId, rLocks),
                },
                (error, response) => {
                    if (error) {
                        onCreateLease({ error });
                        return;
                    }

                    onCreateLease({ error: null, leaseId: response.leaseId });
                }
            );

            const response = await waitForCreateLease();

            if (response.error) {
                release();
                throw response.error;
            }

            leaseId = response.leaseId;
            keepLeaseAlive(); // do not await this!
            release();
            await connection.connect();
        } catch (e) {
            if (e === E_CANCELED) {
                return;
            }
            throw e;
        }
    };

    const keepLeaseAlive = async () => {
        while (true) {
            const response = await racePromises({
                stop: waitForStopExecution(),
                keepalive: wait((ttl / 2) * 1000),
            });

            if (response === "stop") {
                await runDropLease();
                return;
            }

            if (!(await runKeepalive())) {
                return;
            }
        }
    };

    // returns true if processing can be continued
    const runKeepalive = async () => {
        const [waitForKeepalive, onKeepalive] =
            await createResolvablePromise<ServiceError | null>();

        serviceClient.keepLease({ leaseId, ttl }, onKeepalive);

        const error = await waitForKeepalive();

        if (error) {
            await connection.disconnect();

            try {
                await retry(renew, 100, 50);
                return false;
            } catch {
                throw new Error("Unable to renew lease");
            }
        }

        return true;
    };

    const runDropLease = async () => {
        const [waitForDropped, onDrop] = await createResolvablePromise<ServiceError | null>();

        serviceClient.dropLease({ leaseId }, onDrop);

        const error = await waitForDropped();

        if (error) {
            throw error;
        }

        stopped();
    };

    return {
        stop: () => {
            mutex.cancel();
            stopExecution();
        },
        runWithLeaseId: async (callback: (leaseId: string) => Promise<void>) => {
            await mutex.runExclusive(async () => {
                await callback(leaseId);
            });
        },
        renew,
        waitForStop,
        track: async (tenant: string, resource: string[], read: boolean) => {
            try {
                await mutex.runExclusive(() => {
                    if (read) {
                        rLocks.push({ tenant, resource });
                        return;
                    }

                    locks.push({ tenant, resource });
                });
            } catch (e) {
                if (e === E_CANCELED) {
                    return;
                }
                throw e;
            }
        },
        untrack: async (tenant: string, resource: string[], read: boolean) => {
            try {
                await mutex.runExclusive(() => {
                    if (read) {
                        let found = false;
                        const newRLocks: Lock[] = [];

                        rLocks.forEach(lock => {
                            if (
                                !found &&
                                lock.tenant === tenant &&
                                arraysEqual(lock.resource, resource)
                            ) {
                                found = true;
                                return;
                            }

                            newRLocks.push(lock);
                        });

                        rLocks = newRLocks;
                        return;
                    }

                    locks = locks.filter(lock => {
                        return lock.tenant !== tenant || !arraysEqual(lock.resource, resource);
                    });
                });
            } catch (e) {
                if (e === E_CANCELED) {
                    return;
                }
                throw e;
            }
        },
    };
};

const arraysEqual = <T>(a: T[], b: T[]): boolean => {
    if (a.length !== b.length) {
        return false;
    }

    for (let i = 0; i < a.length; i++) {
        if (a[i] !== b[i]) {
            return false;
        }
    }

    return true;
};

const getPbLocks = (leaseId: string, locks: Lock[]): PbLock[] => {
    return locks.map(lock => {
        return {
            leaseId,
            tenantId: lock.tenant,
            resource: lock.resource,
        };
    });
};
