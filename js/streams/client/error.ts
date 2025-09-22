export interface RetryableError {
    message: string;
    shouldRetry: boolean;
}

export class StreamHandlerError extends Error implements RetryableError {
    public shouldRetry: boolean;

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    constructor(error: any, shouldRetry: boolean = false) {
        super(error instanceof Error ? error.message : String(error));
        this.name = "StreamHandlerError";
        this.shouldRetry = shouldRetry;
    }
}

export const isRetryableError = (error: unknown): error is RetryableError => {
    return (
        error instanceof StreamHandlerError ||
        (typeof error === "object" &&
            error !== null &&
            "shouldRetry" in error &&
            "message" in error)
    );
};
