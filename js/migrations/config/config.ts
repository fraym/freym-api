import { config } from "dotenv";

export interface ClientConfig {
    // serverAddress: address of the projection service
    serverAddress: string;
    // apiToken: auth token for the api
    apiToken: string;
    // namespace: namespace for the api
    namespace: string;
}

export const getEnvConfig = (): ClientConfig => {
    config();

    const secure =
        process.env.MIGRATIONS_SECURE === "1" ||
        process.env.MIGRATIONS_SECURE?.toLowerCase() === "true";

    const httpProtocoll = secure ? "https" : "http";

    return {
        serverAddress: process.env.MIGRATIONS_SERVER_ADDRESS
            ? `${httpProtocoll}${process.env.MIGRATIONS_SERVER_ADDRESS}`
            : "",
        apiToken: process.env.MIGRATIONS_API_TOKEN ?? "",
        namespace: process.env.MIGRATIONS_NAMESPACE ?? "",
    };
};

export const useConfigDefaults = (config?: ClientConfig): Required<ClientConfig> => {
    if (!config) {
        config = getEnvConfig();
    }

    return {
        serverAddress: config.serverAddress,
        apiToken: config.apiToken,
        namespace: config.namespace,
    };
};
