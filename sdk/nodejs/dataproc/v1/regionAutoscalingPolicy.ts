// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates new autoscaling policy.
 */
export class RegionAutoscalingPolicy extends pulumi.CustomResource {
    /**
     * Get an existing RegionAutoscalingPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegionAutoscalingPolicy {
        return new RegionAutoscalingPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dataproc/v1:RegionAutoscalingPolicy';

    /**
     * Returns true if the given object is an instance of RegionAutoscalingPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionAutoscalingPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionAutoscalingPolicy.__pulumiType;
    }

    public readonly basicAlgorithm!: pulumi.Output<outputs.dataproc.v1.BasicAutoscalingAlgorithmResponse>;
    /**
     * The "resource name" of the autoscaling policy, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.autoscalingPolicies, the resource name of the policy has the following format: projects/{project_id}/regions/{region}/autoscalingPolicies/{policy_id} For projects.locations.autoscalingPolicies, the resource name of the policy has the following format: projects/{project_id}/locations/{location}/autoscalingPolicies/{policy_id}
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Optional. Describes how the autoscaler will operate for secondary workers.
     */
    public readonly secondaryWorkerConfig!: pulumi.Output<outputs.dataproc.v1.InstanceGroupAutoscalingPolicyConfigResponse>;
    /**
     * Required. Describes how the autoscaler will operate for primary workers.
     */
    public readonly workerConfig!: pulumi.Output<outputs.dataproc.v1.InstanceGroupAutoscalingPolicyConfigResponse>;

    /**
     * Create a RegionAutoscalingPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionAutoscalingPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.autoscalingPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoscalingPolicyId'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.regionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'regionId'");
            }
            inputs["autoscalingPolicyId"] = args ? args.autoscalingPolicyId : undefined;
            inputs["basicAlgorithm"] = args ? args.basicAlgorithm : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["regionId"] = args ? args.regionId : undefined;
            inputs["secondaryWorkerConfig"] = args ? args.secondaryWorkerConfig : undefined;
            inputs["workerConfig"] = args ? args.workerConfig : undefined;
            inputs["name"] = undefined /*out*/;
        } else {
            inputs["basicAlgorithm"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["secondaryWorkerConfig"] = undefined /*out*/;
            inputs["workerConfig"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RegionAutoscalingPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegionAutoscalingPolicy resource.
 */
export interface RegionAutoscalingPolicyArgs {
    readonly autoscalingPolicyId: pulumi.Input<string>;
    readonly basicAlgorithm?: pulumi.Input<inputs.dataproc.v1.BasicAutoscalingAlgorithmArgs>;
    /**
     * Required. The policy id.The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
     */
    readonly id?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    readonly regionId: pulumi.Input<string>;
    /**
     * Optional. Describes how the autoscaler will operate for secondary workers.
     */
    readonly secondaryWorkerConfig?: pulumi.Input<inputs.dataproc.v1.InstanceGroupAutoscalingPolicyConfigArgs>;
    /**
     * Required. Describes how the autoscaler will operate for primary workers.
     */
    readonly workerConfig?: pulumi.Input<inputs.dataproc.v1.InstanceGroupAutoscalingPolicyConfigArgs>;
}
