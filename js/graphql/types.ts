import { CombinedError, Operation, RequestPolicy } from "@urql/core";
import { JWT } from "./client/client";
import { ClientOptions } from "./client/options";

export type OnGraphqlErrorFunction = (
    error: CombinedError,
    operation: Operation
) => CombinedError | null;

interface BaseOptions {
    url: string;
    jwt: JWT;
    requestPolicy?: RequestPolicy;
    clientOptions?: ClientOptions;
    onGraphqlError?: OnGraphqlErrorFunction; // return null in case the error should be ignored
}

export interface UseQueryOptions extends BaseOptions {
    paused?: boolean; // pauses the query
}

export interface UseSubscriptionOptions<T> extends UseQueryOptions {
    onSubscriptionUpdate?: (updateData: T) => T; // called when the subscription updates, can be used for sorting in list subscriptions
}

export type UseMutationOptions = BaseOptions;

export const defaultOnSubscriptionUpdate = <T>(v: T) => v;
