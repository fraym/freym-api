import { Config } from "../cmd/config";
import { Deployment, DeploymentResponse, DeploymentTarget } from "../schema/data";

export const createDeployment = async (
    deployment: Deployment,
    config: Config
): Promise<DeploymentResponse> => {
    const response = await fetch(`${config.serverAddress}/api/deployment`, {
        method: "POST",
        headers: {
            Authorization: `Bearer ${config.apiToken}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(deployment),
    });

    if (!response.ok) {
        throw new Error(await response.text());
    }

    return await response.json();
};

export type DeploymentStatus = Record<string, number>;

export const getDeploymentStatus = async (
    deploymentId: number,
    config: Config
): Promise<DeploymentStatus> => {
    const response = await fetch(`${config.serverAddress}/api/deployment/${deploymentId}`, {
        method: "GET",
        headers: {
            Authorization: `Bearer ${config.apiToken}`,
            "Content-Type": "application/json",
        },
    });

    if (!response.ok) {
        throw new Error(await response.text());
    }

    return await response.json();
};

export const activateDeployment = async (deploymentId: number, config: Config): Promise<void> => {
    const response = await fetch(
        `${config.serverAddress}/api/deployment/${deploymentId}/activate`,
        {
            method: "POST",
            headers: {
                Authorization: `Bearer ${config.apiToken}`,
                "Content-Type": "application/json",
            },
        }
    );

    if (!response.ok) {
        throw new Error(await response.text());
    }
};

export const confirmDeployment = async (deploymentId: number, config: Config): Promise<void> => {
    const response = await fetch(`${config.serverAddress}/api/deployment/${deploymentId}/confirm`, {
        method: "POST",
        headers: {
            Authorization: `Bearer ${config.apiToken}`,
            "Content-Type": "application/json",
        },
    });

    if (!response.ok) {
        throw new Error(await response.text());
    }
};

export const rollbackDeployment = async (deploymentId: number, config: Config): Promise<void> => {
    const response = await fetch(
        `${config.serverAddress}/api/deployment/${deploymentId}/rollback`,
        {
            method: "POST",
            headers: {
                Authorization: `Bearer ${config.apiToken}`,
                "Content-Type": "application/json",
            },
        }
    );

    if (!response.ok) {
        throw new Error(await response.text());
    }
};

export const rollbackDeploymentByNamespace = async (
    config: Config,
    target: DeploymentTarget
): Promise<void> => {
    const response = await fetch(`${config.serverAddress}/api/deployment/rollback/namespace`, {
        method: "POST",
        headers: {
            Authorization: `Bearer ${config.apiToken}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            namespace: config.namespace,
            target,
        }),
    });

    if (!response.ok) {
        throw new Error(await response.text());
    }
};

export const getCurrentDeployment = async (config: Config): Promise<DeploymentTarget> => {
    const response = await fetch(
        `${config.serverAddress}/api/target/production/${config.namespace}`,
        {
            method: "GET",
            headers: {
                Authorization: `Bearer ${config.apiToken}`,
                "Content-Type": "application/json",
            },
        }
    );

    if (!response.ok) {
        throw new Error(await response.text());
    }

    const data = await response.json();

    return data.target;
};

export const setCurrentDeployment = async (
    config: Config,
    target: DeploymentTarget
): Promise<void> => {
    const response = await fetch(`${config.serverAddress}/api/target/production`, {
        method: "POST",
        headers: {
            Authorization: `Bearer ${config.apiToken}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            namespace: config.namespace,
            target,
        }),
    });

    if (!response.ok) {
        throw new Error(await response.text());
    }
};
