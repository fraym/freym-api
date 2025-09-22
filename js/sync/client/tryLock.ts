import { ServiceClient } from "@fraym/proto/dist/index.freym.sync.management";
import { Lease } from "./lease";
import { retry } from "./retry";

export const tryLock = async (
    lease: Lease,
    tenantId: string,
    resource: string[],
    serviceClient: ServiceClient
) => {
    let locked = false;

    await retry(() =>
        lease.runWithLeaseId(async leaseId => {
            return new Promise<void>((resolve, reject) => {
                serviceClient.tryLock(
                    {
                        leaseId,
                        tenantId,
                        resource,
                    },
                    (error, response) => {
                        if (error) {
                            reject(error);
                            return;
                        }
                        locked = response.locked;
                        resolve();
                    }
                );
            });
        })
    );

    await lease.track(tenantId, resource, false);

    return locked;
};
