import { promises as fsPromises } from "fs";
import { GraphQLObjectType, GraphQLSchema } from "graphql";
import { Deployment, ObjectType, View } from "./data";
import { getObjectDirectives, hasDirective } from "./directive";
import { getEnums, getPermissions } from "./enum";
import { getObjectFields } from "./field";
import { ensureValidName } from "./util";

export const getMigrationFromSchema = async (
    schema: GraphQLSchema,
    namespace: string,
    dangerouslyRemoveGdprFields: boolean
): Promise<Deployment> => {
    return {
        ...(await getTypes(schema, namespace)),
        namespace,
        enumTypes: getEnums(schema, namespace),
        permissions: getPermissions(schema),
        options: {
            dangerouslyRemoveGdprFields, // @todo: make configurable
            skipServices: [], // @todo: make configurable
            force: false, // @todo: make configurable
        },
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
}> => {
    const usedNames: string[] = [];
    const crudTypes: ObjectType[] = [];
    const projectionTypes: ObjectType[] = [];
    const nestedTypes: ObjectType[] = [];
    const views: View[] = [];

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
        } else if (hasDirective(t, "view")) {
            const directives = getObjectDirectives(t);
            const viewDirective = directives.find(d => d.name === "view");
            const sqlFileArg = viewDirective?.arguments.find(a => a.name === "sqlFile");
            const sqlFileName = sqlFileArg?.value;

            if (!sqlFileName || typeof sqlFileName !== "string") {
                throw new Error(`view directive on type "${name}" requires a sqlFile argument`);
            }

            const sql = await fsPromises.readFile(sqlFileName, {
                encoding: "utf8",
            });

            views.push({
                name,
                sql,
                directives,
                fields: getObjectFields(t, namespace),
            });
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
    };
};
