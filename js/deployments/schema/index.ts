import { promises as fsPromises } from "fs";
import { GraphQLObjectType, GraphQLSchema } from "graphql";
import path from "path";
import { Deployment, DeploymentOptions, ObjectType, View } from "./data";
import { getObjectDirectives, hasDirective } from "./directive";
import { getEnums, getPermissions } from "./enum";
import { getObjectFields } from "./field";
import { ensureValidName } from "./util";

export const getDeploymentFromSchema = async (
    schema: GraphQLSchema,
    namespace: string,
    options: DeploymentOptions
): Promise<Deployment> => {
    return {
        ...(await getTypes(schema, namespace)),
        namespace,
        enumTypes: getEnums(schema, namespace),
        permissions: getPermissions(schema),
        options,
    };
};

const getTypes = async (
    schema: GraphQLSchema,
    namespace: string
): Promise<{
    crudTypes: ObjectType[];
    projectionTypes: ObjectType[];
    nestedTypes: ObjectType[];
    views: View[];
    baseViews: View[];
}> => {
    const usedNames: string[] = [];
    const crudTypes: ObjectType[] = [];
    const projectionTypes: ObjectType[] = [];
    const nestedTypes: ObjectType[] = [];
    const views: View[] = [];
    const baseViews: View[] = [];

    for (const t of schema.toConfig().types) {
        if (!(t instanceof GraphQLObjectType) || t.toString().startsWith("__")) {
            continue;
        }

        const name = `${namespace}${t.toString()}`;
        ensureValidName(name);

        if (usedNames.includes(name)) {
            throw new Error(
                `duplicate definition for type "${name}" detected, try renaming one of them as they have to be uniquely named`
            );
        }

        usedNames.push(name);

        if (hasDirective(t, "crudType")) {
            crudTypes.push({
                name,
                directives: getObjectDirectives(t),
                fields: getObjectFields(t, namespace),
            });
        } else if (hasDirective(t, "upsertOn") || hasDirective(t, "dangerouslyUpsertOn")) {
            projectionTypes.push({
                name,
                directives: getObjectDirectives(t),
                fields: getObjectFields(t, namespace),
            });
        } else if (hasDirective(t, "view") || hasDirective(t, "baseView")) {
            if (!t.astNode?.loc) {
                throw new Error(`cannot resolve path of file that contains @view "${name}"`);
            }

            const absolutePath = t.astNode.loc.source.name;
            const relativePath = absolutePath
                .split("/")
                .slice(0, -1)
                .join("/")
                .replace(process.cwd(), ".");
            const directives = getObjectDirectives(t);
            const viewDirective = directives.find(d => d.name === "view" || d.name === "baseView");
            const sqlFileArg = viewDirective?.arguments.find(a => a.name === "sqlFile");
            const sqlFileName = sqlFileArg?.value;

            if (!sqlFileName || typeof sqlFileName !== "string") {
                throw new Error(`view directive on type "${name}" requires a sqlFile argument`);
            }

            const sqlFilePath = sqlFileName.startsWith(".")
                ? "./" + path.join(relativePath, sqlFileName)
                : sqlFileName.startsWith("/")
                  ? sqlFileName.slice(1)
                  : sqlFileName;

            const sql = await fsPromises.readFile(sqlFilePath, {
                encoding: "utf8",
            });

            if (hasDirective(t, "baseView")) {
                baseViews.push({
                    name,
                    sql,
                    directives,
                    fields: getObjectFields(t, namespace),
                });
            } else {
                views.push({
                    name,
                    sql,
                    directives,
                    fields: getObjectFields(t, namespace),
                });
            }
        } else {
            nestedTypes.push({
                name,
                directives: getObjectDirectives(t),
                fields: getObjectFields(t, namespace),
            });
        }
    }

    return {
        crudTypes,
        projectionTypes,
        nestedTypes,
        views,
        baseViews,
    };
};
