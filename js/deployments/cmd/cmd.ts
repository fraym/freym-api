#! /usr/bin/env node
import { DeploymentTarget } from "../schema/data";
import {
    runActivateDeployment,
    runConfirmDeployment,
    runCreateDeployment,
    runPrintDeploymentStatus,
    runRollbackDeployment,
    runRollbackNamespaceDeployment,
    runWait,
} from "./deployment";

const COMMAND_STATUS = "status";
const COMMAND_CREATE = "create";
const COMMAND_ACTIVATE = "activate";
const COMMAND_CONFIRM = "confirm";
const COMMAND_ROLLBACK = "rollback";
const COMMAND_ROLLBACK_NAMESPACE = "rollback-namespace";
const COMMAND_WAIT = "wait";

const arg = process.argv[2] ?? COMMAND_CREATE;

const getIdFromArgs = (): number => {
    const idString = process.argv[3];
    const id = idString ? parseInt(idString, 10) : undefined;

    if (!id || id <= 0) {
        throw new Error("id is required and has to be a number > 0");
    }

    return id;
};

const getTargetFromArgs = (): DeploymentTarget => {
    const targetString = process.argv[3];

    if (targetString !== "blue" && targetString !== "green") {
        throw new Error("deployment target is required and has be one of 'blue' or 'green'");
    }

    return targetString;
};

const isForced = (): boolean => {
    for (const arg of process.argv) {
        if (arg === "force") {
            return true;
        }
    }
    return false;
};

const skipServices = (): string[] => {
    for (const arg of process.argv) {
        if (arg.startsWith("skip=")) {
            const skipArg = arg.split("=")[1];

            if (!skipArg) {
                throw new Error("invalid skip argument");
            }

            return skipArg.split(",");
        }
    }

    return [];
};

switch (arg) {
    case COMMAND_STATUS:
        runPrintDeploymentStatus(getIdFromArgs());
        break;
    case COMMAND_CREATE:
        runCreateDeployment(getTargetFromArgs(), isForced(), skipServices());
        break;
    case COMMAND_ACTIVATE:
        runActivateDeployment(getIdFromArgs());
        break;
    case COMMAND_CONFIRM:
        runConfirmDeployment(getIdFromArgs());
        break;
    case COMMAND_ROLLBACK:
        runRollbackDeployment(getIdFromArgs());
        break;
    case COMMAND_ROLLBACK_NAMESPACE:
        runRollbackNamespaceDeployment(getTargetFromArgs());
        break;
    case COMMAND_WAIT:
        runWait(getIdFromArgs());
        break;
    default:
        throw new Error(`unknown command: ${arg}`);
}
