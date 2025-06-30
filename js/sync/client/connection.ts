import { createResolvablePromise, racePromises } from "./promise";

export interface Connection {
    waitForConnect: () => Promise<void>;
    disconnect: () => Promise<void>;
    connect: () => Promise<void>;
    stop: () => void;
}

type OnConnect = () => void;

export const newConnection = async (): Promise<Connection> => {
    const [waitForStop, onStop] = await createResolvablePromise<void>();
    let connected = false;
    let onConnectCallbacks: OnConnect[] = [];

    return {
        connect: async () => {
            connected = true;

            for (const callback of onConnectCallbacks) {
                callback();
            }

            onConnectCallbacks = [];
        },
        disconnect: async () => {
            connected = false;
        },
        waitForConnect: async () => {
            const [waitForConnect, onConnect] = await createResolvablePromise<void>();

            if (connected) {
                onConnect();
                return;
            }

            onConnectCallbacks.push(onConnect);

            const response = await racePromises({
                connect: waitForConnect(),
                stop: waitForStop(),
            });

            if (response === "stop") {
                onConnect();
                throw new Error("Stopping locking service...");
            }
        },
        stop: onStop,
    };
};
