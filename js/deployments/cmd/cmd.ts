#! /usr/bin/env node
import {
    runConfirmDeployment,
    runCreateDeployment,
    runPrintDeploymentStatus,
    runRollbackDeployment,
    runRollbackNamespaceDeployment,
    runWait,
} from "./deployment";

const COMMAND_STATUS = "status";
const COMMAND_CREATE = "create";
const COMMAND_CONFIRM = "confirm";
const COMMAND_ROLLBACK = "rollback";
const COMMAND_ROLLBACK_NAMESPACE = "rollback-namespace";
const COMMAND_WAIT = "wait";

const arg = process.argv[2] ?? COMMAND_CREATE;

const getIdFromArgs = (): number => {
    const idString = process.argv[3];
    const id = idString ? parseInt(idString, 10) : undefined;

    if (!id || id > 0) {
        throw new Error("id is required and has to by a number > 0");
    }

    return id;
};

switch (arg) {
    case COMMAND_STATUS:
        runPrintDeploymentStatus(getIdFromArgs());
        break;
    case COMMAND_CREATE:
        runCreateDeployment();
        break;
    case COMMAND_CONFIRM:
        runConfirmDeployment(getIdFromArgs());
        break;
    case COMMAND_ROLLBACK:
        runRollbackDeployment(getIdFromArgs());
        break;
    case COMMAND_ROLLBACK_NAMESPACE:
        runRollbackNamespaceDeployment();
        break;
    case COMMAND_WAIT:
        runWait(getIdFromArgs());
        break;
    default:
        throw new Error(`unknown command: ${arg}`);
}
