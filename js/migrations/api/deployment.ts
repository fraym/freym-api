import { ClientConfig } from "@/config/config";
import { Deployment, DeploymentResponse } from "@/schema/data";

export const createDeployment = async (
    deployment: Deployment,
    config: ClientConfig
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
    config: ClientConfig
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

export const confirmDeployment = async (
    deploymentId: number,
    config: ClientConfig
): Promise<void> => {
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

export const rollbackDeployment = async (
    deploymentId: number,
    config: ClientConfig
): Promise<void> => {
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
