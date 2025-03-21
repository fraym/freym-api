type ResolvablePromiseResolveFunction<T> = (value: T) => void;

type ResolvablePromise<T> = [() => Promise<T>, ResolvablePromiseResolveFunction<T>];

export const createResolvablePromise = async <T>(
    onResolve?: (resolvePromise: (value: T) => void) => Promise<void>
) => {
    return new Promise<ResolvablePromise<T>>(resolve => {
        const promise = new Promise<T>(resolvePromise => {
            resolve([
                () => promise,
                async (value: T) => {
                    if (onResolve) {
                        await onResolve(resolvePromise);
                    } else {
                        resolvePromise(value);
                    }
                },
            ]);
        });
    });
};

export const wait = async (milliSeconds: number) => {
    return new Promise<void>(resolve => {
        setTimeout(() => {
            resolve();
        }, milliSeconds);
    });
};

type WaitResponse<T> = {
    response: T;
};

export const racePromises = async <T extends string | number>(
    promises: Record<T, Promise<void>>
): Promise<T> => {
    const waitFunctions: Promise<WaitResponse<T>>[] = [];

    for (const key in promises) {
        const waitFn = async () => {
            await promises[key];
            return {
                response: key,
            } as WaitResponse<T>;
        };

        waitFunctions.push(waitFn());
    }

    const { response } = await Promise.race(waitFunctions);
    return response;
};
