export interface RequestGate {
    enter: () => Promise<void>;
    leave: () => void;
}

interface WaitingRequest {
    resolve: () => void;
}

export const newRequestGate = (maxRequests: number): RequestGate => {
    let activeRequests = 0;
    const waitingQueue: WaitingRequest[] = [];

    return {
        enter: async () => {
            // If we can enter immediately, do so
            if (activeRequests < maxRequests) {
                activeRequests++;
                return;
            }

            // Otherwise, create a promise and add it to the waiting queue
            return new Promise<void>(resolve => {
                waitingQueue.push({ resolve });
            });
        },
        leave: () => {
            activeRequests--;

            // If there are waiting requests, allow the next one to proceed
            if (waitingQueue.length > 0) {
                const waiting = waitingQueue.shift()!;
                activeRequests++;
                waiting.resolve();
            }
        },
    };
};
