// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a network in the specified project and region using the data included in the request.
 */
export class RegionNetwork extends pulumi.CustomResource {
    /**
     * Get an existing RegionNetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegionNetwork {
        return new RegionNetwork(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:RegionNetwork';

    /**
     * Returns true if the given object is an instance of RegionNetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionNetwork {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionNetwork.__pulumiType;
    }

    /**
     * Must be set to create a VPC network. If not set, a legacy network is created. When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode. An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges. For custom mode VPC networks, you can add subnets using the subnetworks insert method.
     */
    public readonly autoCreateSubnetworks!: pulumi.Output<boolean>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this field when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix fd20::/20. .
     */
    public readonly enableUlaInternalIpv6!: pulumi.Output<boolean>;
    /**
     * URL of the firewall policy the network is associated with.
     */
    public /*out*/ readonly firewallPolicy!: pulumi.Output<string>;
    /**
     * The gateway address for default routing out of the network, selected by GCP.
     */
    public /*out*/ readonly gatewayIPv4!: pulumi.Output<string>;
    /**
     * When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field. .
     */
    public readonly internalIpv6Range!: pulumi.Output<string>;
    /**
     * Type of the resource. Always compute#network for networks.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Maximum Transmission Unit in bytes. The minimum value for this field is 1460 and the maximum value is 1500 bytes. If unspecified, defaults to 1460.
     */
    public readonly mtu!: pulumi.Output<number>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The network firewall policy enforcement order. Can be either AFTER_CLASSIC_FIREWALL or BEFORE_CLASSIC_FIREWALL. Defaults to AFTER_CLASSIC_FIREWALL if the field is not specified.
     */
    public readonly networkFirewallPolicyEnforcementOrder!: pulumi.Output<string>;
    /**
     * A list of network peerings for the resource.
     */
    public /*out*/ readonly peerings!: pulumi.Output<outputs.compute.alpha.NetworkPeeringResponse[]>;
    /**
     * URL of the region where the regional network resides. This field is not applicable to global network. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
     */
    public readonly routingConfig!: pulumi.Output<outputs.compute.alpha.NetworkRoutingConfigResponse>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * Server-defined fully-qualified URLs for all subnetworks in this VPC network.
     */
    public /*out*/ readonly subnetworks!: pulumi.Output<string[]>;

    /**
     * Create a RegionNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionNetworkArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["autoCreateSubnetworks"] = args ? args.autoCreateSubnetworks : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enableUlaInternalIpv6"] = args ? args.enableUlaInternalIpv6 : undefined;
            resourceInputs["internalIpv6Range"] = args ? args.internalIpv6Range : undefined;
            resourceInputs["mtu"] = args ? args.mtu : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkFirewallPolicyEnforcementOrder"] = args ? args.networkFirewallPolicyEnforcementOrder : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["routingConfig"] = args ? args.routingConfig : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["firewallPolicy"] = undefined /*out*/;
            resourceInputs["gatewayIPv4"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["peerings"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["subnetworks"] = undefined /*out*/;
        } else {
            resourceInputs["autoCreateSubnetworks"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["enableUlaInternalIpv6"] = undefined /*out*/;
            resourceInputs["firewallPolicy"] = undefined /*out*/;
            resourceInputs["gatewayIPv4"] = undefined /*out*/;
            resourceInputs["internalIpv6Range"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["mtu"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["networkFirewallPolicyEnforcementOrder"] = undefined /*out*/;
            resourceInputs["peerings"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["routingConfig"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["subnetworks"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RegionNetwork.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegionNetwork resource.
 */
export interface RegionNetworkArgs {
    /**
     * Must be set to create a VPC network. If not set, a legacy network is created. When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode. An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges. For custom mode VPC networks, you can add subnets using the subnetworks insert method.
     */
    autoCreateSubnetworks?: pulumi.Input<boolean>;
    /**
     * An optional description of this resource. Provide this field when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix fd20::/20. .
     */
    enableUlaInternalIpv6?: pulumi.Input<boolean>;
    /**
     * When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field. .
     */
    internalIpv6Range?: pulumi.Input<string>;
    /**
     * Maximum Transmission Unit in bytes. The minimum value for this field is 1460 and the maximum value is 1500 bytes. If unspecified, defaults to 1460.
     */
    mtu?: pulumi.Input<number>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
     */
    name?: pulumi.Input<string>;
    /**
     * The network firewall policy enforcement order. Can be either AFTER_CLASSIC_FIREWALL or BEFORE_CLASSIC_FIREWALL. Defaults to AFTER_CLASSIC_FIREWALL if the field is not specified.
     */
    networkFirewallPolicyEnforcementOrder?: pulumi.Input<enums.compute.alpha.RegionNetworkNetworkFirewallPolicyEnforcementOrder>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
     */
    routingConfig?: pulumi.Input<inputs.compute.alpha.NetworkRoutingConfigArgs>;
}
