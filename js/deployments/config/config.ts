import { config } from "dotenv";

export interface ClientConfig {
    // serverAddress: address of the service
    serverAddress: string;
    // apiToken: auth token for the api
    apiToken: string;
    // namespace: namespace for the api
    namespace: string;
}

export const getEnvConfig = (): ClientConfig => {
    config();

    return {
        serverAddress: process.env.DEPLOYMENTS_SERVER_ADDRESS ?? "",
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
