export type ClientOptions = {
    enableCache?: boolean; // default is false
    cacheTtl?: number; // default is 5 minutes = 300_000
};

export const getClientOptions = (options?: ClientOptions) => {
    return {
        enableCache: options?.enableCache ?? false,
        cacheTtl: options?.cacheTtl ?? 300_000,
    };
};
