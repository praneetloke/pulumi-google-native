// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new policy in the specified project using the data included in the request.
 * Auto-naming is currently not supported for this resource.
 */
export class NetworkFirewallPolicy extends pulumi.CustomResource {
    /**
     * Get an existing NetworkFirewallPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NetworkFirewallPolicy {
        return new NetworkFirewallPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:NetworkFirewallPolicy';

    /**
     * Returns true if the given object is an instance of NetworkFirewallPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkFirewallPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkFirewallPolicy.__pulumiType;
    }

    /**
     * A list of associations that belong to this firewall policy.
     */
    public readonly associations!: pulumi.Output<outputs.compute.alpha.FirewallPolicyAssociationResponse[]>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the firewall policy.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * [Output only] Type of the resource. Always compute#firewallPolicyfor firewall policies
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The parent of the firewall policy.
     */
    public /*out*/ readonly parent!: pulumi.Output<string>;
    /**
     * URL of the region where the regional firewall policy resides. This field is not applicable to global firewall policies. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public /*out*/ readonly region!: pulumi.Output<string>;
    /**
     * Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
     */
    public /*out*/ readonly ruleTupleCount!: pulumi.Output<number>;
    /**
     * A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
     */
    public readonly rules!: pulumi.Output<outputs.compute.alpha.FirewallPolicyRuleResponse[]>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly shortName!: pulumi.Output<string>;
    /**
     * The scope of networks allowed to be associated with the firewall policy. This field can be either GLOBAL_VPC_NETWORK or REGIONAL_VPC_NETWORK. A firewall policy with the VPC scope set to GLOBAL_VPC_NETWORK is allowed to be attached only to global networks. When the VPC scope is set to REGIONAL_VPC_NETWORK the firewall policy is allowed to be attached only to regional networks in the same scope as the firewall policy. Note: if not specified then GLOBAL_VPC_NETWORK will be used.
     */
    public readonly vpcNetworkScope!: pulumi.Output<string>;

    /**
     * Create a NetworkFirewallPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NetworkFirewallPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["associations"] = args ? args.associations : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["shortName"] = args ? args.shortName : undefined;
            resourceInputs["vpcNetworkScope"] = args ? args.vpcNetworkScope : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parent"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["ruleTupleCount"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
        } else {
            resourceInputs["associations"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parent"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["ruleTupleCount"] = undefined /*out*/;
            resourceInputs["rules"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["shortName"] = undefined /*out*/;
            resourceInputs["vpcNetworkScope"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkFirewallPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a NetworkFirewallPolicy resource.
 */
export interface NetworkFirewallPolicyArgs {
    /**
     * A list of associations that belong to this firewall policy.
     */
    associations?: pulumi.Input<pulumi.Input<inputs.compute.alpha.FirewallPolicyAssociationArgs>[]>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.compute.alpha.FirewallPolicyRuleArgs>[]>;
    /**
     * User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    shortName?: pulumi.Input<string>;
    /**
     * The scope of networks allowed to be associated with the firewall policy. This field can be either GLOBAL_VPC_NETWORK or REGIONAL_VPC_NETWORK. A firewall policy with the VPC scope set to GLOBAL_VPC_NETWORK is allowed to be attached only to global networks. When the VPC scope is set to REGIONAL_VPC_NETWORK the firewall policy is allowed to be attached only to regional networks in the same scope as the firewall policy. Note: if not specified then GLOBAL_VPC_NETWORK will be used.
     */
    vpcNetworkScope?: pulumi.Input<enums.compute.alpha.NetworkFirewallPolicyVpcNetworkScope>;
}
