// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a NodeGroup resource in the specified project using the data included in the request.
 */
export class NodeGroup extends pulumi.CustomResource {
    /**
     * Get an existing NodeGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NodeGroup {
        return new NodeGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:NodeGroup';

    /**
     * Returns true if the given object is an instance of NodeGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NodeGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NodeGroup.__pulumiType;
    }

    /**
     * Specifies how autoscaling should behave.
     */
    public readonly autoscalingPolicy!: pulumi.Output<outputs.compute.alpha.NodeGroupAutoscalingPolicyResponse>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The type of the resource. Always compute#nodeGroup for node group.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * An opaque location hint used to place the Node close to other resources. This field is for use by internal tools that use the public API. The location hint here on the NodeGroup overrides any location_hint present in the NodeTemplate.
     */
    public readonly locationHint!: pulumi.Output<string>;
    /**
     * Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT. For more information, see Maintenance policies.
     */
    public readonly maintenancePolicy!: pulumi.Output<string>;
    public readonly maintenanceWindow!: pulumi.Output<outputs.compute.alpha.NodeGroupMaintenanceWindowResponse>;
    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * URL of the node template to create the node group from.
     */
    public readonly nodeTemplate!: pulumi.Output<string>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * Share-settings for the node group
     */
    public readonly shareSettings!: pulumi.Output<outputs.compute.alpha.ShareSettingsResponse>;
    /**
     * The total number of nodes in the node group.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    public readonly status!: pulumi.Output<string>;
    /**
     * The name of the zone where the node group resides, such as us-central1-a.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a NodeGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodeGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.initialNodeCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'initialNodeCount'");
            }
            inputs["autoscalingPolicy"] = args ? args.autoscalingPolicy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["initialNodeCount"] = args ? args.initialNodeCount : undefined;
            inputs["locationHint"] = args ? args.locationHint : undefined;
            inputs["maintenancePolicy"] = args ? args.maintenancePolicy : undefined;
            inputs["maintenanceWindow"] = args ? args.maintenanceWindow : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeTemplate"] = args ? args.nodeTemplate : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["shareSettings"] = args ? args.shareSettings : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
            inputs["size"] = undefined /*out*/;
        } else {
            inputs["autoscalingPolicy"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["locationHint"] = undefined /*out*/;
            inputs["maintenancePolicy"] = undefined /*out*/;
            inputs["maintenanceWindow"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nodeTemplate"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
            inputs["shareSettings"] = undefined /*out*/;
            inputs["size"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["zone"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NodeGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NodeGroup resource.
 */
export interface NodeGroupArgs {
    /**
     * Specifies how autoscaling should behave.
     */
    autoscalingPolicy?: pulumi.Input<inputs.compute.alpha.NodeGroupAutoscalingPolicyArgs>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    initialNodeCount: pulumi.Input<string>;
    /**
     * An opaque location hint used to place the Node close to other resources. This field is for use by internal tools that use the public API. The location hint here on the NodeGroup overrides any location_hint present in the NodeTemplate.
     */
    locationHint?: pulumi.Input<string>;
    /**
     * Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT. For more information, see Maintenance policies.
     */
    maintenancePolicy?: pulumi.Input<enums.compute.alpha.NodeGroupMaintenancePolicy>;
    maintenanceWindow?: pulumi.Input<inputs.compute.alpha.NodeGroupMaintenanceWindowArgs>;
    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * URL of the node template to create the node group from.
     */
    nodeTemplate?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * Share-settings for the node group
     */
    shareSettings?: pulumi.Input<inputs.compute.alpha.ShareSettingsArgs>;
    status?: pulumi.Input<enums.compute.alpha.NodeGroupStatus>;
    zone?: pulumi.Input<string>;
}
