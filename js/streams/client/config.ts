import { config } from "dotenv";

export interface ClientConfig {
    // serverAddress: address of the streams service
    serverAddress: string;
    // groupId: your services group identifier
    groupId: string;
    // keepaliveInterval: grpc connection keepalive ping interval in milliseconds
    keepaliveInterval?: number;
    // keepaliveTimeout: grpc connection keepalive ping timeout in milliseconds
    keepaliveTimeout?: number;
}

export const useConfigDefaults = (config: ClientConfig): Required<ClientConfig> => {
    return {
        serverAddress: config.serverAddress,
        groupId: config.groupId,
        keepaliveInterval: config.keepaliveInterval ?? 40 * 1000,
        keepaliveTimeout: config.keepaliveTimeout ?? 3 * 1000,
    };
};

export const getEnvDeliveryConfig = (): ClientConfig => {
    config();

    let keepaliveInterval: number | undefined;
    let keepaliveTimeout: number | undefined;

    const keepaliveIntervalString = process.env.STREAMS_CLIENT_KEEPALIVE_INTERVAL;
    const keepaliveTimeoutString = process.env.STREAMS_CLIENT_KEEPALIVE_TIMEOUT;

    if (keepaliveIntervalString) {
        keepaliveInterval = parseInt(keepaliveIntervalString, 10);
    }

    if (keepaliveTimeoutString) {
        keepaliveTimeout = parseInt(keepaliveTimeoutString, 10);
    }

    return {
        serverAddress: process.env.STREAMS_CLIENT_ADDRESS ?? "",
        groupId: process.env.STREAMS_CLIENT_GROUP_ID ?? "",
        keepaliveInterval,
        keepaliveTimeout,
    };
};
