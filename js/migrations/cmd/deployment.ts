import {
    DeploymentStatus,
    confirmDeployment,
    createDeployment,
    getDeploymentStatus,
    rollbackDeployment,
} from "@/api/deployment";
import { getMigrationFromSchema } from "@/schema";
import { loadSchema } from "@graphql-tools/load";
import { useConfig } from "./config";
import { replaceEnvPlaceholdersGraphQLFileLoader } from "./loader";

export const runPrintDeploymentStatus = async (id: number) => {
    console.log("get deployment status ...");
    const { serverAddress, apiToken, namespace } = await useConfig();

    const status = await getDeploymentStatus(id, {
        apiToken,
        serverAddress,
        namespace,
    });
    console.log(status);
};

export const runCreateDeployment = async () => {
    console.log("creating deployment ...");
    const { schemaGlob, namespace, serverAddress, apiToken } = await useConfig();

    const schema = await loadSchema(schemaGlob, {
        loaders: [replaceEnvPlaceholdersGraphQLFileLoader],
    });
    const deployment = await getMigrationFromSchema(schema, namespace, false);

    const response = await createDeployment(
        { ...deployment },
        { apiToken, serverAddress, namespace }
    );
    console.log(`created deployment ${response.target}/${response.deploymentId}`);
};

export const runConfirmDeployment = async (id: number) => {
    console.log("confirming deployment ...");
    const { serverAddress, apiToken, namespace } = await useConfig();

    await confirmDeployment(id, {
        apiToken,
        serverAddress,

        namespace,
    });
    console.log("done confirming deployment");
};

export const runRollbackDeployment = async (id: number) => {
    console.log("rolling back deployment ...");
    const { serverAddress, apiToken, namespace } = await useConfig();

    await rollbackDeployment(id, {
        apiToken,
        serverAddress,

        namespace,
    });
    console.log("done rolling back deployment");
};

export const runWait = async (id: number) => {
    console.log("waiting for migration to be ready to finish ...");
    const { serverAddress, apiToken, namespace } = await useConfig();

    const isReady = false;
    let percentage = 0;

    while (!isReady) {
        const status = await getDeploymentStatus(id, {
            apiToken,
            serverAddress,
            namespace,
        });

        const newPercentage = getPercentage(status);

        if (newPercentage === 100) {
            console.log("deployment is ready");
            break;
        }

        if (percentage !== newPercentage) {
            console.log(`current progress: ${newPercentage}%`);
        }

        percentage = newPercentage;

        await sleep(1000);
    }

    console.log("migration is now ready to finish");
};

const sleep = async (time: number): Promise<void> => {
    return new Promise(resolve => {
        setTimeout(() => {
            resolve();
        }, time);
    });
};

const getPercentage = (status: DeploymentStatus): number => {
    let percentageSum = 0;
    let services = 0;

    for (const key in status) {
        if (key === "crud" || key === "auth") {
            // these are directly finished
            if (status[key] === 100) {
                continue;
            }
        }

        if (status[key] !== undefined) {
            percentageSum += status[key];
            services++;
        }
    }

    return Math.floor(percentageSum / services);
};
