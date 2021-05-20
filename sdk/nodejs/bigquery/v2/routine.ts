// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new routine in the dataset.
 */
export class Routine extends pulumi.CustomResource {
    /**
     * Get an existing Routine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Routine {
        return new Routine(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:bigquery/v2:Routine';

    /**
     * Returns true if the given object is an instance of Routine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Routine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Routine.__pulumiType;
    }

    /**
     * Optional.
     */
    public readonly arguments!: pulumi.Output<outputs.bigquery.v2.ArgumentResponse[]>;
    /**
     * The time when this routine was created, in milliseconds since the epoch.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * Required. The body of the routine. For functions, this is the expression in the AS clause. If language=SQL, it is the substring inside (but excluding) the parentheses. For example, for the function created with the following statement: `CREATE FUNCTION JoinLines(x string, y string) as (concat(x, "\n", y))` The definition_body is `concat(x, "\n", y)` (\n is not replaced with linebreak). If language=JAVASCRIPT, it is the evaluated string in the AS clause. For example, for the function created with the following statement: `CREATE FUNCTION f() RETURNS STRING LANGUAGE js AS 'return "\n";\n'` The definition_body is `return "\n";\n` Note that both \n are replaced with linebreaks.
     */
    public readonly definitionBody!: pulumi.Output<string>;
    /**
     * Optional. [Experimental] The description of the routine if defined.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. [Experimental] The determinism level of the JavaScript UDF if defined.
     */
    public readonly determinismLevel!: pulumi.Output<string>;
    /**
     * A hash of this resource.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Optional. If language = "JAVASCRIPT", this field stores the path of the imported JAVASCRIPT libraries.
     */
    public readonly importedLibraries!: pulumi.Output<string[]>;
    /**
     * Optional. Defaults to "SQL".
     */
    public readonly language!: pulumi.Output<string>;
    /**
     * The time when this routine was last modified, in milliseconds since the epoch.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;
    /**
     * Optional. Set only if Routine is a "TABLE_VALUED_FUNCTION".
     */
    public readonly returnTableType!: pulumi.Output<outputs.bigquery.v2.StandardSqlTableTypeResponse>;
    /**
     * Optional if language = "SQL"; required otherwise. If absent, the return type is inferred from definition_body at query time in each query that references this routine. If present, then the evaluated result will be cast to the specified returned type at query time. For example, for the functions created with the following statements: * `CREATE FUNCTION Add(x FLOAT64, y FLOAT64) RETURNS FLOAT64 AS (x + y);` * `CREATE FUNCTION Increment(x FLOAT64) AS (Add(x, 1));` * `CREATE FUNCTION Decrement(x FLOAT64) RETURNS FLOAT64 AS (Add(x, -1));` The return_type is `{type_kind: "FLOAT64"}` for `Add` and `Decrement`, and is absent for `Increment` (inferred as FLOAT64 at query time). Suppose the function `Add` is replaced by `CREATE OR REPLACE FUNCTION Add(x INT64, y INT64) AS (x + y);` Then the inferred return type of `Increment` is automatically changed to INT64 at query time, while the return type of `Decrement` remains FLOAT64.
     */
    public readonly returnType!: pulumi.Output<outputs.bigquery.v2.StandardSqlDataTypeResponse>;
    /**
     * Required. Reference describing the ID of this routine.
     */
    public readonly routineReference!: pulumi.Output<outputs.bigquery.v2.RoutineReferenceResponse>;
    /**
     * Required. The type of routine.
     */
    public readonly routineType!: pulumi.Output<string>;

    /**
     * Create a Routine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoutineArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.datasetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasetId'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.routineId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routineId'");
            }
            inputs["arguments"] = args ? args.arguments : undefined;
            inputs["datasetId"] = args ? args.datasetId : undefined;
            inputs["definitionBody"] = args ? args.definitionBody : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["determinismLevel"] = args ? args.determinismLevel : undefined;
            inputs["importedLibraries"] = args ? args.importedLibraries : undefined;
            inputs["language"] = args ? args.language : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["returnTableType"] = args ? args.returnTableType : undefined;
            inputs["returnType"] = args ? args.returnType : undefined;
            inputs["routineId"] = args ? args.routineId : undefined;
            inputs["routineReference"] = args ? args.routineReference : undefined;
            inputs["routineType"] = args ? args.routineType : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
        } else {
            inputs["arguments"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["definitionBody"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["determinismLevel"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["importedLibraries"] = undefined /*out*/;
            inputs["language"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["returnTableType"] = undefined /*out*/;
            inputs["returnType"] = undefined /*out*/;
            inputs["routineReference"] = undefined /*out*/;
            inputs["routineType"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Routine.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Routine resource.
 */
export interface RoutineArgs {
    /**
     * Optional.
     */
    readonly arguments?: pulumi.Input<pulumi.Input<inputs.bigquery.v2.ArgumentArgs>[]>;
    readonly datasetId: pulumi.Input<string>;
    /**
     * Required. The body of the routine. For functions, this is the expression in the AS clause. If language=SQL, it is the substring inside (but excluding) the parentheses. For example, for the function created with the following statement: `CREATE FUNCTION JoinLines(x string, y string) as (concat(x, "\n", y))` The definition_body is `concat(x, "\n", y)` (\n is not replaced with linebreak). If language=JAVASCRIPT, it is the evaluated string in the AS clause. For example, for the function created with the following statement: `CREATE FUNCTION f() RETURNS STRING LANGUAGE js AS 'return "\n";\n'` The definition_body is `return "\n";\n` Note that both \n are replaced with linebreaks.
     */
    readonly definitionBody?: pulumi.Input<string>;
    /**
     * Optional. [Experimental] The description of the routine if defined.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Optional. [Experimental] The determinism level of the JavaScript UDF if defined.
     */
    readonly determinismLevel?: pulumi.Input<string>;
    /**
     * Optional. If language = "JAVASCRIPT", this field stores the path of the imported JAVASCRIPT libraries.
     */
    readonly importedLibraries?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional. Defaults to "SQL".
     */
    readonly language?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * Optional. Set only if Routine is a "TABLE_VALUED_FUNCTION".
     */
    readonly returnTableType?: pulumi.Input<inputs.bigquery.v2.StandardSqlTableTypeArgs>;
    /**
     * Optional if language = "SQL"; required otherwise. If absent, the return type is inferred from definition_body at query time in each query that references this routine. If present, then the evaluated result will be cast to the specified returned type at query time. For example, for the functions created with the following statements: * `CREATE FUNCTION Add(x FLOAT64, y FLOAT64) RETURNS FLOAT64 AS (x + y);` * `CREATE FUNCTION Increment(x FLOAT64) AS (Add(x, 1));` * `CREATE FUNCTION Decrement(x FLOAT64) RETURNS FLOAT64 AS (Add(x, -1));` The return_type is `{type_kind: "FLOAT64"}` for `Add` and `Decrement`, and is absent for `Increment` (inferred as FLOAT64 at query time). Suppose the function `Add` is replaced by `CREATE OR REPLACE FUNCTION Add(x INT64, y INT64) AS (x + y);` Then the inferred return type of `Increment` is automatically changed to INT64 at query time, while the return type of `Decrement` remains FLOAT64.
     */
    readonly returnType?: pulumi.Input<inputs.bigquery.v2.StandardSqlDataTypeArgs>;
    readonly routineId: pulumi.Input<string>;
    /**
     * Required. Reference describing the ID of this routine.
     */
    readonly routineReference?: pulumi.Input<inputs.bigquery.v2.RoutineReferenceArgs>;
    /**
     * Required. The type of routine.
     */
    readonly routineType?: pulumi.Input<string>;
}
