import { useCallback, useState } from "react";
import { getErrorFromResult } from "@/error";
import { UseMutationOptions } from "@/types";
import { AnyVariables, CombinedError, DocumentInput } from "@urql/core";
import { useClient } from "./useClient";
import { useUpdatingRef } from "./useUpdatingRef";

export const useMutation = <
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    Data = any,
    Variables extends AnyVariables = AnyVariables,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    Element extends { id: string } = any,
>(
    mutation: DocumentInput<Data, Variables>,
    variables: Variables,
    getElementsFromMutation: (data: Data) => Element[],
    { jwt, url, requestPolicy, onGraphqlError, clientOptions }: UseMutationOptions
) => {
    const client = useClient(url, jwt, clientOptions);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<CombinedError | null>(null);
    const getFromMutation = useUpdatingRef(getElementsFromMutation);
    const onError = useUpdatingRef(onGraphqlError);

    const execute = useCallback(async () => {
        if (!client) {
            return;
        }

        const result = await client.mutation(mutation, variables, {
            requestPolicy: "network-only",
        });
        const error = getErrorFromResult(result, result.operation);

        if (error) {
            setError(onError.current ? onError.current(error, result.operation) : error);
            setLoading(false);
            return;
        }

        setError(null);
        setLoading(false);
        return getFromMutation.current(result.data! ?? []);

        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [mutation, requestPolicy, variables, client]);

    return {
        execute,
        loading,
        error,
    };
};
