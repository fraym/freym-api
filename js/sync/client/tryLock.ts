import { ServiceClient } from "@fraym/proto/dist/index.freym.sync.management";
import { Lease } from "./lease";
import { RequestGate } from "./requestGate";
import { retry } from "./retry";

export const tryLock = async (
    lease: Lease,
    tenantId: string,
    resource: string[],
    serviceClient: ServiceClient,
    requestGate: RequestGate
) => {
    let locked = false;

    await retry(() =>
        lease.runWithLeaseId(async leaseId => {
            await requestGate.enter();

            return new Promise<void>((resolve, reject) => {
                serviceClient.tryLock(
                    {
                        leaseId,
                        tenantId,
                        resource,
                    },
                    (error, response) => {
                        requestGate.leave();

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
