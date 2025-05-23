import { CombinedError, Operation } from "@urql/core";

export type GraphqlError = CombinedError;
export type GraphqlOperation = Operation;

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const getErrorFromResult = <Data = any>(
    result: {
        data?: Data;
        error?: GraphqlError;
    },
    operation: Operation
): GraphqlError | null => {
    return (
        result.error ??
        (result.data
            ? null
            : new CombinedError({
                  graphQLErrors: [],
                  networkError: new Error(`No data returned from ${operation.kind}`),
              }))
    );
};
