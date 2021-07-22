// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an entity type in the specified agent. Note: You should always train an agent prior to sending it queries. See the [training documentation](https://cloud.google.com/dialogflow/es/docs/training).
 * Auto-naming is currently not supported for this resource.
 */
export class EntityType extends pulumi.CustomResource {
    /**
     * Get an existing EntityType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EntityType {
        return new EntityType(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v2beta1:EntityType';

    /**
     * Returns true if the given object is an instance of EntityType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EntityType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EntityType.__pulumiType;
    }

    /**
     * Optional. Indicates whether the entity type can be automatically expanded.
     */
    public readonly autoExpansionMode!: pulumi.Output<string>;
    /**
     * The name of the entity type.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Optional. Enables fuzzy entity extraction during classification.
     */
    public readonly enableFuzzyExtraction!: pulumi.Output<boolean>;
    /**
     * Optional. The collection of entity entries associated with the entity type.
     */
    public readonly entities!: pulumi.Output<outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1EntityTypeEntityResponse[]>;
    /**
     * Indicates the kind of entity type.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType and EntityTypes.BatchUpdateEntityTypes methods. Supported formats: - `projects//agent/entityTypes/` - `projects//locations//agent/entityTypes/`
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a EntityType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EntityTypeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.kind === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kind'");
            }
            inputs["autoExpansionMode"] = args ? args.autoExpansionMode : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["enableFuzzyExtraction"] = args ? args.enableFuzzyExtraction : undefined;
            inputs["entities"] = args ? args.entities : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["languageCode"] = args ? args.languageCode : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
        } else {
            inputs["autoExpansionMode"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["enableFuzzyExtraction"] = undefined /*out*/;
            inputs["entities"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EntityType.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a EntityType resource.
 */
export interface EntityTypeArgs {
    /**
     * Optional. Indicates whether the entity type can be automatically expanded.
     */
    autoExpansionMode?: pulumi.Input<enums.dialogflow.v2beta1.EntityTypeAutoExpansionMode>;
    /**
     * The name of the entity type.
     */
    displayName: pulumi.Input<string>;
    /**
     * Optional. Enables fuzzy entity extraction during classification.
     */
    enableFuzzyExtraction?: pulumi.Input<boolean>;
    /**
     * Optional. The collection of entity entries associated with the entity type.
     */
    entities?: pulumi.Input<pulumi.Input<inputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1EntityTypeEntityArgs>[]>;
    /**
     * Indicates the kind of entity type.
     */
    kind: pulumi.Input<enums.dialogflow.v2beta1.EntityTypeKind>;
    languageCode?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType and EntityTypes.BatchUpdateEntityTypes methods. Supported formats: - `projects//agent/entityTypes/` - `projects//locations//agent/entityTypes/`
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
