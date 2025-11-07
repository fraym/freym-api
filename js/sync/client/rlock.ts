import { ServiceClient } from "@fraym/proto/dist/index.freym.sync.management";
import { Lease } from "./lease";
import { RequestGate } from "./requestGate";
import { retry } from "./retry";

export const rLock = async (
    lease: Lease,
    tenantId: string,
    resource: string[],
    serviceClient: ServiceClient,
    requestGate: RequestGate
) => {
    await retry(() =>
        lease.runWithLeaseId(async leaseId => {
            await requestGate.enter();

            return new Promise<void>((resolve, reject) => {
                serviceClient.rLock(
                    {
                        leaseId,
                        tenantId,
                        resource,
                    },
                    error => {
                        requestGate.leave();

                        if (error) {
                            reject(error);
                            return;
                        }
                        resolve();
                    }
                );
            });
        })
    );

    await lease.track(tenantId, resource, true);
};
