import { Mutex } from "async-mutex";
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
    const mutex = new Mutex();
    let connected = false;
    let onConnectCallbacks: OnConnect[] = [];

    return {
        connect: async () => {
            await mutex.runExclusive(async () => {
                connected = true;

                for (const callback of onConnectCallbacks) {
                    callback();
                }

                onConnectCallbacks = [];
            });
        },
        disconnect: async () => {
            await mutex.runExclusive(async () => {
                connected = false;
            });
        },
        waitForConnect: async () => {
            const [waitForConnect, onConnect] = await createResolvablePromise<void>();

            await mutex.runExclusive(async () => {
                if (connected) {
                    onConnect();
                    return;
                }

                onConnectCallbacks.push(onConnect);
            });

            const response = await racePromises({
                connect: waitForConnect(),
                stop: waitForStop(),
            });

            if (response === "stop") {
                throw new Error("Stopping locking service...");
            }
        },
        stop: onStop,
    };
};
