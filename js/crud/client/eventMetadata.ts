import {
    EventMetadata as CrudEventMetadata,
    DeploymentTarget,
} from "@fraym/proto/dist/index.freym.crud.delivery";

export interface EventMetadata {
    correlationId: string;
    causationId: string;
    target: DeploymentTarget;
}

export const fillMetadataWithDefaults = (
    metadata: Partial<EventMetadata> | null
): CrudEventMetadata => {
    if (!metadata) {
        return {
            causationId: "",
            correlationId: "",
            target: "DEPLOYMENT_TARGET_BLUE",
        };
    }

    return {
        causationId: metadata.causationId || "",
        correlationId: metadata.correlationId || "",
        target: metadata.target || "DEPLOYMENT_TARGET_BLUE",
    };
};
