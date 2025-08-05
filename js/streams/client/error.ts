export interface RetryableError {
    message: string;
    shouldRetry: boolean;
}

export class StreamHandlerError extends Error implements RetryableError {
    public shouldRetry: boolean;

    constructor(message: string, shouldRetry: boolean = false) {
        super(message);
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
