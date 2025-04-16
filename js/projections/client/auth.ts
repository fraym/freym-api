import { AuthData as PbAuthData } from "@fraym/proto/dist/index.freym.projections.delivery";

export interface AuthData {
    tenantId: string;
    scopes: string[];
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    data: Record<string, any>;
    userId?: string;
}

export const getProtobufAuthData = (auth: AuthData): PbAuthData => {
    const data: Record<string, string> = {};

    for (const key in auth.data) {
        data[key] = JSON.stringify(auth.data[key]);
    }

    return {
        tenantId: auth.tenantId,
        userId: auth.userId ?? "",
        scopes: auth.scopes,
        data,
    };
};
