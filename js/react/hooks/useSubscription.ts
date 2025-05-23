import { useCallback, useEffect, useState } from "react";
import { getErrorFromResult } from "@fraym/graphql/error";
import { UseSubscriptionOptions, defaultOnSubscriptionUpdate } from "@fraym/graphql/types";
import { AnyVariables, CombinedError, DocumentInput } from "@urql/core";
import { useClient } from "./useClient";
import { useUpdatingRef } from "./useUpdatingRef";

export const useSubscription = <
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    Data = any,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    SubscriptionData = any,
    Variables extends AnyVariables = AnyVariables,
    SubscriptionVariables extends AnyVariables = AnyVariables,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    Element extends { id: string } = any,
>(
    query: DocumentInput<Data, Variables>,
    subscription: DocumentInput<SubscriptionData, SubscriptionVariables>,
    variables: Variables,
    subscriptionVariables: SubscriptionVariables,
    getElementFromQuery: (data: Data) => Element,
    getElementFromSubscription: (data: SubscriptionData, prev: Element | null) => Element | null,
    {
        jwt,
        url,
        requestPolicy,
        paused,
        onGraphqlError,
        clientOptions,
        onSubscriptionUpdate,
    }: UseSubscriptionOptions<Element | null>
) => {
    const client = useClient(url, jwt, clientOptions);
    const [loading, setLoading] = useState<boolean>(true);
    const [data, setData] = useState<Element | null>(null);
    const [error, setError] = useState<CombinedError | null>(null);
    const getFromQuery = useUpdatingRef(getElementFromQuery);
    const getFromSubscription = useUpdatingRef(getElementFromSubscription);
    const onError = useUpdatingRef(onGraphqlError);
    const onUpdate = useUpdatingRef(onSubscriptionUpdate ?? defaultOnSubscriptionUpdate);

    const loadData = useCallback(async () => {
        if (!client) {
            return;
        }

        const result = await client.query(query, variables, {
            requestPolicy,
        });
        const error = getErrorFromResult(result, result.operation);

        if (error) {
            setError(onError.current ? onError.current(error, result.operation) : error);
        } else {
            setData(onUpdate.current(getFromQuery.current(result.data!)));
            setError(null);
        }

        setLoading(false);

        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [query, requestPolicy, variables, client]);

    useEffect(() => {
        if (paused) {
            return;
        }

        loadData();
    }, [loadData, paused]);

    useEffect(() => {
        if (!client || paused) {
            return;
        }

        const { unsubscribe } = client
            .subscription(subscription, subscriptionVariables, {
                requestPolicy: "network-only",
            })
            .subscribe(result => {
                const error = getErrorFromResult(result, result.operation);

                if (error) {
                    setError(onError.current ? onError.current(error, result.operation) : error);
                    return;
                }

                setData(prev => onUpdate.current(getFromSubscription.current(result.data!, prev)));
                setError(null);
            });

        return unsubscribe;
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [subscription, subscriptionVariables, paused, client]);

    return {
        data,
        loading,
        error,
    };
};
