import { ServiceClient } from "@fraym/proto/dist/index.freym.sync.management";
import { Lease } from "./lease";
import { retry } from "./retry";

export const unlock = async (
    lease: Lease,
    tenantId: string,
    resource: string[],
    serviceClient: ServiceClient
) => {
    await retry(() =>
        lease.runWithLeaseId(async leaseId => {
            return new Promise<void>((resolve, reject) => {
                serviceClient.unlock(
                    {
                        leaseId,
                        tenantId,
                        resource,
                    },
                    error => {
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

    await lease.untrack(tenantId, resource, false);
};
