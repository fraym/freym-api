import { config } from "dotenv";
import { hideBin } from "yargs/helpers";
import yargs from "yargs/yargs";

export interface Config {
    schemaGlob: string;
    namespace: string;
    serverAddress: string;
    apiToken: string;
}

export const useConfig = async (): Promise<Config> => {
    config({ quiet: true });

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

    if (process.env.DEPLOYMENTS_SCHEMA_GLOB) {
        schemaGlob = process.env.DEPLOYMENTS_SCHEMA_GLOB;
    }

    if (process.env.DEPLOYMENTS_SERVER_ADDRESS) {
        serverAddress = process.env.DEPLOYMENTS_SERVER_ADDRESS ?? "";
    }

    if (process.env.DEPLOYMENTS_API_TOKEN) {
        apiToken = process.env.DEPLOYMENTS_API_TOKEN;
    }

    if (process.env.DEPLOYMENTS_NAMESPACE) {
        namespace = process.env.DEPLOYMENTS_NAMESPACE;
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
