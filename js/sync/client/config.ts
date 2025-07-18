import { config } from "dotenv";

export interface ClientConfig {
    // serverAddress: address of the streams service
    serverAddress: string;
    // appPrefix: your services app prefix
    appPrefix: string;
    // keepaliveInterval: grpc connection keepalive ping interval in milliseconds
    keepaliveInterval?: number;
    // keepaliveTimeout: grpc connection keepalive ping timeout in milliseconds
    keepaliveTimeout?: number;
}

export const useConfigDefaults = (config: ClientConfig): Required<ClientConfig> => {
    return {
        serverAddress: config.serverAddress,
        appPrefix: config.appPrefix,
        keepaliveInterval: config.keepaliveInterval ?? 40 * 1000,
        keepaliveTimeout: config.keepaliveTimeout ?? 3 * 1000,
    };
};

export const getEnvConfig = (): ClientConfig => {
    config({ quiet: true });

    let keepaliveInterval: number | undefined;
    let keepaliveTimeout: number | undefined;

    const keepaliveIntervalString = process.env.SYNC_CLIENT_KEEPALIVE_INTERVAL;
    const keepaliveTimeoutString = process.env.SYNC_CLIENT_KEEPALIVE_TIMEOUT;

    if (keepaliveIntervalString) {
        keepaliveInterval = parseInt(keepaliveIntervalString, 10);
    }

    if (keepaliveTimeoutString) {
        keepaliveTimeout = parseInt(keepaliveTimeoutString, 10);
    }

    return {
        serverAddress: process.env.SYNC_CLIENT_ADDRESS ?? "",
        appPrefix: process.env.SYNC_CLIENT_APP_PREFIX ?? "",
        keepaliveInterval,
        keepaliveTimeout,
    };
};
