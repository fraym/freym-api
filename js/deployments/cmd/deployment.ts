import {
    DeploymentStatus,
    activateDeployment,
    confirmDeployment,
    createDeployment,
    getDeploymentStatus,
    rollbackDeployment,
    rollbackDeploymentByNamespace,
} from "@/api/deployment";
import { getMigrationFromSchema } from "@/schema";
import { DeploymentTarget } from "@/schema/data";
import { loadSchema } from "@graphql-tools/load";
import { useConfig } from "./config";
import { replaceEnvPlaceholdersGraphQLFileLoader } from "./loader";

export const runPrintDeploymentStatus = async (id: number) => {
    console.log("get deployment status ...");
    const config = await useConfig();

    const status = await getDeploymentStatus(id, config);
    console.log(status);
};

export const runCreateDeployment = async (target: DeploymentTarget) => {
    console.log("creating deployment ...");
    const config = await useConfig();

    const schema = await loadSchema(config.schemaGlob, {
        loaders: [replaceEnvPlaceholdersGraphQLFileLoader],
    });

    // @todo: make options configurable
    const deployment = await getMigrationFromSchema(schema, config.namespace, {
        dangerouslyRemoveGdprFields: false,
        skipServices: [],
        force: false,
        target,
    });

    const response = await createDeployment({ ...deployment }, config);
    console.log(`created deployment ${response.deploymentId}`);
};

export const runActivateDeployment = async (id: number) => {
    console.log("activating deployment ...");
    const config = await useConfig();

    await activateDeployment(id, config);
    console.log("activated deployment");
};

export const runConfirmDeployment = async (id: number) => {
    console.log("confirming deployment ...");
    const config = await useConfig();

    await confirmDeployment(id, config);
    console.log("confirmed deployment");
};

export const runRollbackDeployment = async (id: number) => {
    console.log("rolling back deployment ...");
    const config = await useConfig();

    await rollbackDeployment(id, config);
    console.log("rolled back deployment");
};

export const runRollbackNamespaceDeployment = async () => {
    console.log("rolling back deployment ...");
    const config = await useConfig();

    await rollbackDeploymentByNamespace(config);
    console.log("rolled back deployment");
};

export const runWait = async (id: number) => {
    console.log("waiting for deployment to be done...");
    const config = await useConfig();

    const isReady = false;
    let percentage = 0;

    while (!isReady) {
        const status = await getDeploymentStatus(id, config);

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

    console.log("deployment is done");
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
