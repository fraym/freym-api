import { SubscriptionEvent } from "client/event";
import { Status } from "@grpc/grpc-js/build/src/constants";

export const sleep = (delay: number) =>
    new Promise<void>(resolve => {
        setTimeout(() => resolve(), delay);
    });

export const retry = async <T>(
    fn: () => Promise<T>,
    pause: number = 100,
    retries: number = 50
): Promise<T> => {
    try {
        return await fn();
        /* eslint-disable-next-line @typescript-eslint/no-explicit-any */
    } catch (err: any) {
        if (retries === 0 || (err && err.code && err.code === Status.UNKNOWN)) {
            throw err;
        }

        await sleep(pause);

        return await retry(fn, pause, retries - 1);
    }
};

export type StopLoadingMoreFunc = (lastEvent: SubscriptionEvent | null) => boolean;
