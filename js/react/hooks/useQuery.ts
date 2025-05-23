import { useCallback, useEffect, useState } from "react";
import { getErrorFromResult } from "@fraym/graphql/error";
import { UseQueryOptions } from "@fraym/graphql/types";
import { AnyVariables, CombinedError, DocumentInput } from "@urql/core";
import { useClient } from "./useClient";
import { useUpdatingRef } from "./useUpdatingRef";

export const useQuery = <
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    Data = any,
    Variables extends AnyVariables = AnyVariables,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    Element extends { id: string } = any,
>(
    query: DocumentInput<Data, Variables>,
    variables: Variables,
    getElementsFromQuery: (data: Data) => Element[],
    { jwt, url, requestPolicy, paused, onGraphqlError, clientOptions }: UseQueryOptions
) => {
    const client = useClient(url, jwt, clientOptions);
    const [loading, setLoading] = useState<boolean>(true);
    const [data, setData] = useState<Element[] | null>(null);
    const [error, setError] = useState<CombinedError | null>(null);
    const getFromQuery = useUpdatingRef(getElementsFromQuery);
    const onError = useUpdatingRef(onGraphqlError);

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
            setData(getFromQuery.current(result.data! ?? []));
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

    return {
        data,
        loading,
        error,
    };
};
