import { subscriptionExchange } from "@urql/core";
import { authExchange } from "@urql/exchange-auth";
import { getWsClient } from "./wsClient";

export const getAuthExchange = (jwt: string) => {
    return authExchange(async utils => {
        return {
            addAuthToOperation: operation => {
                return utils.appendHeaders(operation, {
                    Authorization: `Bearer ${jwt}`,
                });
            },
            didAuthError(error) {
                return error.graphQLErrors.some(e => e.extensions?.code === "FORBIDDEN");
            },
            refreshAuth: async () => {
                throw new Error("JWT expired");
            },
        };
    });
};

export const getSubscriptionExchange = (url: string, jwt: string) => {
    const wsClient = getWsClient(url, jwt);

    return wsClient
        ? subscriptionExchange({
              forwardSubscription(request) {
                  const input = { ...request, query: request.query || "" };

                  return {
                      subscribe(sink) {
                          const unsubscribe = wsClient.subscribe(input, {
                              next: value => {
                                  sink.next({ data: value });
                              },
                              error: console.error,
                              complete: console.log,
                          });
                          return {
                              unsubscribe,
                          };
                      },
                  };
              },
          })
        : null;
};
