export interface Deployment {
    namespace: string;
    permissions: string[];
    projectionTypes: ObjectType[];
    crudTypes: ObjectType[];
    nestedTypes: ObjectType[];
    enumTypes: EnumType[];
    views: View[];
    baseViews: View[];
    options: DeploymentOptions;
}

export interface DeploymentResponse {
    deploymentId: number;
}

export type DeploymentTarget = "blue" | "green";

export interface DeploymentOptions {
    target: DeploymentTarget;
    force: boolean;
    skipServices: string[];
}

export interface View {
    name: string;
    sql: string;
    directives: TypeDirective[];
    fields: TypeField[];
}

export interface ObjectType {
    name: string;
    directives: TypeDirective[];
    fields: TypeField[];
}

export interface EnumType {
    name: string;
    values: string[];
}

export interface TypeField {
    name: string;
    type: string[];
    directives: TypeDirective[];
}

export interface TypeDirective {
    name: string;
    arguments: TypeArgument[];
}

export interface TypeArgument {
    name: string;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    value: any;
}
