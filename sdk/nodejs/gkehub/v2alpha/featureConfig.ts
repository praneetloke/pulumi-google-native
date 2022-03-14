// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates FeatureConfig under a given parent.
 * Auto-naming is currently not supported for this resource.
 */
export class FeatureConfig extends pulumi.CustomResource {
    /**
     * Get an existing FeatureConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FeatureConfig {
        return new FeatureConfig(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:gkehub/v2alpha:FeatureConfig';

    /**
     * Returns true if the given object is an instance of FeatureConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FeatureConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FeatureConfig.__pulumiType;
    }

    /**
     * When the FeatureConfig resource was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * When the FeatureConfig resource was deleted.
     */
    public /*out*/ readonly deleteTime!: pulumi.Output<string>;
    /**
     * GCP labels for this FeatureConfig.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource name of this FeatureConfig, in the format: `projects/{project}/locations/global/FeatureConfigs/{feature_type}/{feature_config}`
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Input only. Immutable. User input of feature spec. Note that this field is immutable. Must create a new FeatureConfig if a new feature spec is needed.
     */
    public readonly spec!: pulumi.Output<outputs.gkehub.v2alpha.FeatureSpecResponse>;
    /**
     * Lifecycle information of the FeatureConfig.
     */
    public readonly state!: pulumi.Output<outputs.gkehub.v2alpha.FeatureConfigStateResponse>;
    /**
     * Google-generated UUID for this resource. This is unique across all FeatureConfig resources. If a Membership resource is deleted and another resource with the same name is created, it gets a different unique_id.
     */
    public /*out*/ readonly uniqueId!: pulumi.Output<string>;
    /**
     * When the FeatureConfig resource was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a FeatureConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FeatureConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["featureConfigId"] = args ? args.featureConfigId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["spec"] = args ? args.spec : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["uniqueId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteTime"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["spec"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uniqueId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FeatureConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a FeatureConfig resource.
 */
export interface FeatureConfigArgs {
    /**
     * The ID of the feature config to create.
     */
    featureConfigId?: pulumi.Input<string>;
    /**
     * GCP labels for this FeatureConfig.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Idempotent request UUID.
     */
    requestId?: pulumi.Input<string>;
    /**
     * Input only. Immutable. User input of feature spec. Note that this field is immutable. Must create a new FeatureConfig if a new feature spec is needed.
     */
    spec?: pulumi.Input<inputs.gkehub.v2alpha.FeatureSpecArgs>;
    /**
     * Lifecycle information of the FeatureConfig.
     */
    state?: pulumi.Input<inputs.gkehub.v2alpha.FeatureConfigStateArgs>;
}
