export type ClientOptions = {
    enableCache?: boolean; // default is false
    cacheTtl?: number; // default is 5 minutes = 300_000
    onOutdated?: () => void;
    onStatusCode?: Record<number, () => void>;
};

export const getClientOptions = (options?: ClientOptions) => {
    return {
        enableCache: options?.enableCache ?? false,
        cacheTtl: options?.cacheTtl ?? 300_000,
    };
};
