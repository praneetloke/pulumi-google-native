// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new Gateway in a given project and location.
 */
export class Gateway extends pulumi.CustomResource {
    /**
     * Get an existing Gateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Gateway {
        return new Gateway(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:networkservices/v1:Gateway';

    /**
     * Returns true if the given object is an instance of Gateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Gateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Gateway.__pulumiType;
    }

    /**
     * Optional. Zero or one IPv4-address on which the Gateway will receive the traffic. When no address is provided, an IP from the subnetwork is allocated This field only applies to gateways of type 'SECURE_WEB_GATEWAY'. Gateways of type 'OPEN_MESH' listen on 0.0.0.0.
     */
    public readonly addresses!: pulumi.Output<string[]>;
    /**
     * Optional. A fully-qualified Certificates URL reference. The proxy presents a Certificate (selected based on SNI) when establishing a TLS connection. This feature only applies to gateways of type 'SECURE_WEB_GATEWAY'.
     */
    public readonly certificateUrls!: pulumi.Output<string[]>;
    /**
     * The timestamp when the resource was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. A free-text description of the resource. Max length 1024 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Required. Short name of the Gateway resource to be created.
     */
    public readonly gatewayId!: pulumi.Output<string>;
    /**
     * Optional. A fully-qualified GatewaySecurityPolicy URL reference. Defines how a server should apply security policy to inbound (VM to Proxy) initiated connections. For example: `projects/*&#47;locations/*&#47;gatewaySecurityPolicies/swg-policy`. This policy is specific to gateways of type 'SECURE_WEB_GATEWAY'.
     */
    public readonly gatewaySecurityPolicy!: pulumi.Output<string>;
    /**
     * Optional. Set of label tags associated with the Gateway resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Name of the Gateway resource. It matches pattern `projects/*&#47;locations/*&#47;gateways/`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Optional. The relative resource name identifying the VPC network that is using this configuration. For example: `projects/*&#47;global/networks/network-1`. Currently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY'.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * One or more port numbers (1-65535), on which the Gateway will receive traffic. The proxy binds to the specified ports. Gateways of type 'SECURE_WEB_GATEWAY' are limited to 1 port. Gateways of type 'OPEN_MESH' listen on 0.0.0.0 and support multiple ports.
     */
    public readonly ports!: pulumi.Output<number[]>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Optional. Scope determines how configuration across multiple Gateway instances are merged. The configuration for multiple Gateway instances with the same scope will be merged as presented as a single coniguration to the proxy/load balancer. Max length 64 characters. Scope should start with a letter and can only have letters, numbers, hyphens.
     */
    public readonly scope!: pulumi.Output<string>;
    /**
     * Server-defined URL of this resource
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Optional. A fully-qualified ServerTLSPolicy URL reference. Specifies how TLS traffic is terminated. If empty, TLS termination is disabled.
     */
    public readonly serverTlsPolicy!: pulumi.Output<string>;
    /**
     * Optional. The relative resource name identifying the subnetwork in which this SWG is allocated. For example: `projects/*&#47;regions/us-central1/subnetworks/network-1` Currently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY".
     */
    public readonly subnetwork!: pulumi.Output<string>;
    /**
     * Immutable. The type of the customer managed gateway. This field is required. If unspecified, an error is returned.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The timestamp when the resource was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Gateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewayArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.gatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayId'");
            }
            if ((!args || args.ports === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ports'");
            }
            resourceInputs["addresses"] = args ? args.addresses : undefined;
            resourceInputs["certificateUrls"] = args ? args.certificateUrls : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["gatewayId"] = args ? args.gatewayId : undefined;
            resourceInputs["gatewaySecurityPolicy"] = args ? args.gatewaySecurityPolicy : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["ports"] = args ? args.ports : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["serverTlsPolicy"] = args ? args.serverTlsPolicy : undefined;
            resourceInputs["subnetwork"] = args ? args.subnetwork : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["addresses"] = undefined /*out*/;
            resourceInputs["certificateUrls"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["gatewayId"] = undefined /*out*/;
            resourceInputs["gatewaySecurityPolicy"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["network"] = undefined /*out*/;
            resourceInputs["ports"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["scope"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["serverTlsPolicy"] = undefined /*out*/;
            resourceInputs["subnetwork"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["gatewayId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Gateway.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Gateway resource.
 */
export interface GatewayArgs {
    /**
     * Optional. Zero or one IPv4-address on which the Gateway will receive the traffic. When no address is provided, an IP from the subnetwork is allocated This field only applies to gateways of type 'SECURE_WEB_GATEWAY'. Gateways of type 'OPEN_MESH' listen on 0.0.0.0.
     */
    addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional. A fully-qualified Certificates URL reference. The proxy presents a Certificate (selected based on SNI) when establishing a TLS connection. This feature only applies to gateways of type 'SECURE_WEB_GATEWAY'.
     */
    certificateUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional. A free-text description of the resource. Max length 1024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Required. Short name of the Gateway resource to be created.
     */
    gatewayId: pulumi.Input<string>;
    /**
     * Optional. A fully-qualified GatewaySecurityPolicy URL reference. Defines how a server should apply security policy to inbound (VM to Proxy) initiated connections. For example: `projects/*&#47;locations/*&#47;gatewaySecurityPolicies/swg-policy`. This policy is specific to gateways of type 'SECURE_WEB_GATEWAY'.
     */
    gatewaySecurityPolicy?: pulumi.Input<string>;
    /**
     * Optional. Set of label tags associated with the Gateway resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Name of the Gateway resource. It matches pattern `projects/*&#47;locations/*&#47;gateways/`.
     */
    name?: pulumi.Input<string>;
    /**
     * Optional. The relative resource name identifying the VPC network that is using this configuration. For example: `projects/*&#47;global/networks/network-1`. Currently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY'.
     */
    network?: pulumi.Input<string>;
    /**
     * One or more port numbers (1-65535), on which the Gateway will receive traffic. The proxy binds to the specified ports. Gateways of type 'SECURE_WEB_GATEWAY' are limited to 1 port. Gateways of type 'OPEN_MESH' listen on 0.0.0.0 and support multiple ports.
     */
    ports: pulumi.Input<pulumi.Input<number>[]>;
    project?: pulumi.Input<string>;
    /**
     * Optional. Scope determines how configuration across multiple Gateway instances are merged. The configuration for multiple Gateway instances with the same scope will be merged as presented as a single coniguration to the proxy/load balancer. Max length 64 characters. Scope should start with a letter and can only have letters, numbers, hyphens.
     */
    scope?: pulumi.Input<string>;
    /**
     * Optional. A fully-qualified ServerTLSPolicy URL reference. Specifies how TLS traffic is terminated. If empty, TLS termination is disabled.
     */
    serverTlsPolicy?: pulumi.Input<string>;
    /**
     * Optional. The relative resource name identifying the subnetwork in which this SWG is allocated. For example: `projects/*&#47;regions/us-central1/subnetworks/network-1` Currently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY".
     */
    subnetwork?: pulumi.Input<string>;
    /**
     * Immutable. The type of the customer managed gateway. This field is required. If unspecified, an error is returned.
     */
    type?: pulumi.Input<enums.networkservices.v1.GatewayType>;
}
