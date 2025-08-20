import { cacheExchange, createClient, fetchExchange } from "@urql/core";
import { requestPolicyExchange } from "@urql/exchange-request-policy";
import { retryExchange } from "@urql/exchange-retry";
import { getAuthExchange, getSubscriptionExchange } from "./exchange";
import { ClientOptions, getClientOptions } from "./options";

export type JWT = string | null;

export const getOptionalClient = (url: string, jwt: JWT, options?: ClientOptions) => {
    if (!jwt) {
        return null;
    }

    return getClient(url, jwt, options);
};

export const getClient = (url: string, jwt: NonNullable<JWT>, options?: ClientOptions) => {
    const { cacheTtl, enableCache } = getClientOptions(options);

    return createClient({
        url,
        fetchOptions: {
            headers: {
                "Content-Type": "application/json",
            },
        },
        exchanges: [
            getSubscriptionExchange(url, jwt),
            requestPolicyExchange({
                ttl: cacheTtl,
            }),
            enableCache ? cacheExchange : null,
            getAuthExchange(jwt),
            retryExchange({
                initialDelayMs: 1000,
                maxDelayMs: 5000,
                randomDelay: true,
                maxNumberAttempts: 3,
            }),
            fetchExchange,
        ].filter(exchange => exchange !== null),
        fetch: async (input, init) => {
            const response = await fetch(input, init);

            if (options?.onOutdated && response.headers.get("x-outdated") === "true") {
                options.onOutdated();
            }

            if (options?.onStatusCode) {
                const onStatusCode = options.onStatusCode[response.status];

                if (onStatusCode) {
                    onStatusCode();
                }
            }

            return response;
        },
    });
};
