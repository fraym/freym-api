import { networkInterfaces } from "os";

export const getOwnIpAddress = () => {
    for (const iface of Object.values(networkInterfaces())) {
        if (iface === undefined) {
            continue;
        }

        for (const alias of iface) {
            if (alias.family === "IPv4" && alias.address !== "127.0.0.1" && !alias.internal) {
                return alias.address;
            }
        }
    }

    throw new Error("No local IP address found!");
};
