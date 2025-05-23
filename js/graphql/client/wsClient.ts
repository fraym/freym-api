import { Client, createClient } from "graphql-ws";

const wsClients: Record<string, Record<string, Client>> = {};

export const getWsClient = (url: string, jwt: string) => {
    let wsClient: Client | null = null;

    if (wsClients[url] && wsClients[url][jwt]) {
        wsClient = wsClients[url][jwt];
    } else {
        if (typeof window !== "undefined") {
            if (!wsClients[url]) {
                wsClients[url] = {};
            }

            wsClients[url][jwt] = createClient({
                url: getWsUrl(url),
                lazy: true,
                retryAttempts: Infinity,
                shouldRetry: () => true,
                connectionParams: () => {
                    return {
                        Authorization: `Bearer ${jwt}`,
                    };
                },
            });

            wsClient = wsClients[url][jwt];
        }
    }

    return wsClient;
};

const getWsUrl = (url: string) => {
    if (url.startsWith("https://")) {
        return url.replace("https://", "wss://");
    } else {
        return url.replace("http://", "ws://");
    }
};
