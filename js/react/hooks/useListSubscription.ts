import { useCallback, useEffect, useState } from "react";
import { getErrorFromResult } from "@/error";
import { UseSubscriptionOptions, defaultOnSubscriptionUpdate } from "@/types";
import { AnyVariables, CombinedError, DocumentInput } from "@urql/core";
import { useClient } from "./useClient";
import { useUpdatingRef } from "./useUpdatingRef";

export const useListSubscription = <
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
    getElementsFromQuery: (data: Data) => Element[],
    getElementFromSubscription: (data: SubscriptionData) => { id: string; data: Element } | null,
    {
        jwt,
        url,
        requestPolicy,
        paused,
        onGraphqlError,
        clientOptions,
        onSubscriptionUpdate,
    }: UseSubscriptionOptions<Element[]>
) => {
    const client = useClient(url, jwt, clientOptions);
    const [loading, setLoading] = useState<boolean>(true);
    const [data, setData] = useState<Element[] | null>(null);
    const [error, setError] = useState<CombinedError | null>(null);
    const getFromQuery = useUpdatingRef(getElementsFromQuery);
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
            setData(onUpdate.current(getFromQuery.current(result.data! ?? [])));
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

                setError(null);
                setData(prev => {
                    const element = getFromSubscription.current(result.data! ?? []);

                    if (!element) {
                        return prev;
                    }

                    if (!prev) {
                        return onUpdate.current([element.data]);
                    }

                    const index = prev.findIndex(item => item.id === element.id);

                    if (index === -1) {
                        return onUpdate.current([...prev, element.data]);
                    }

                    const newData = [...prev];
                    newData[index] = element.data;

                    return onUpdate.current(newData);
                });
            });

        return () => {
            unsubscribe();
        };
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [subscription, subscriptionVariables, paused, client]);

    return {
        data,
        loading,
        error,
    };
};
