import { AuthData as pbAuthData } from "@fraym/proto/dist/index.freym.crud.delivery";

export interface AuthData {
    tenantId: string;
    scopes: string[];
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    data: Record<string, any>;
}

export const getProtobufAuthData = (auth: AuthData): pbAuthData => {
    const data: Record<string, string> = {};

    for (const key in auth.data) {
        data[key] = JSON.stringify(auth.data[key]);
    }

    return {
        tenantId: auth.tenantId,
        scopes: auth.scopes,
        data,
    };
};
