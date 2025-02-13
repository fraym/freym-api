import { config } from "dotenv";
import { hideBin } from "yargs/helpers";
import yargs from "yargs/yargs";

export interface CmdConfig {
    schemaGlob: string;
    namespace: string;
    serverAddress: string;
    apiToken: string;
}

export const useConfig = async (): Promise<CmdConfig> => {
    config();

    const argv = await yargs(hideBin(process.argv))
        .config({
            schemaGlob: "./src/**/*.graphql",
            serverAddress: "http://127.0.0.1",
            apiToken: "",
            namespace: "",
        })
        .pkgConf("projections").argv;

    let schemaGlob: string = argv.schemaGlob as string;
    let serverAddress: string = argv.serverAddress as string;
    let apiToken: string = argv.apiToken as string;
    let namespace: string = argv.namespace as string;

    const secure =
        process.env.MIGRATIONS_SECURE === "1" ||
        process.env.MIGRATIONS_SECURE?.toLowerCase() === "true";

    const httpProtocoll = secure ? "https" : "http";

    if (process.env.MIGRATIONS_SCHEMA_GLOB) {
        schemaGlob = process.env.MIGRATIONS_SCHEMA_GLOB;
    }

    if (process.env.MIGRATIONS_SERVER_ADDRESS) {
        serverAddress = `${httpProtocoll}${process.env.MIGRATIONS_SERVER_ADDRESS}`;
    }

    if (process.env.MIGRATIONS_API_TOKEN) {
        apiToken = process.env.MIGRATIONS_API_TOKEN;
    }

    if (process.env.MIGRATIONS_NAMESPACE) {
        namespace = process.env.MIGRATIONS_NAMESPACE;
    }

    if (namespace === "Fraym") {
        throw new Error("Cannot use Fraym as namespace as it is reserved for fraym apps");
    }

    return {
        apiToken,
        namespace,
        schemaGlob,
        serverAddress,
    };
};
