// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a ServiceAttachment in the specified project in the given scope using the parameters that are included in the request.
 */
export class ServiceAttachment extends pulumi.CustomResource {
    /**
     * Get an existing ServiceAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServiceAttachment {
        return new ServiceAttachment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:ServiceAttachment';

    /**
     * Returns true if the given object is an instance of ServiceAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceAttachment.__pulumiType;
    }

    /**
     * An array of connections for all the consumers connected to this service attachment.
     */
    public /*out*/ readonly connectedEndpoints!: pulumi.Output<outputs.compute.alpha.ServiceAttachmentConnectedEndpointResponse[]>;
    /**
     * The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
     */
    public readonly connectionPreference!: pulumi.Output<string>;
    /**
     * Projects that are allowed to connect to this service attachment.
     */
    public readonly consumerAcceptLists!: pulumi.Output<outputs.compute.alpha.ServiceAttachmentConsumerProjectLimitResponse[]>;
    /**
     * Projects that are not allowed to connect to this service attachment. The project can be specified using its id or number.
     */
    public readonly consumerRejectLists!: pulumi.Output<string[]>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * If specified, the domain name will be used during the integration between the PSC connected endpoints and the Cloud DNS. For example, this is a valid domain name: "p.mycompany.com.". Current max number of domain names supported is 1.
     */
    public readonly domainNames!: pulumi.Output<string[]>;
    /**
     * If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
     */
    public readonly enableProxyProtocol!: pulumi.Output<boolean>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a ServiceAttachment. An up-to-date fingerprint must be provided in order to patch/update the ServiceAttachment; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the ServiceAttachment.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * Type of the resource. Always compute#serviceAttachment for service attachments.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An array of URLs where each entry is the URL of a subnet provided by the service producer to use for NAT in this service attachment.
     */
    public readonly natSubnets!: pulumi.Output<string[]>;
    /**
     * The URL of a forwarding rule with loadBalancingScheme INTERNAL* that is serving the endpoint identified by this service attachment.
     */
    public readonly producerForwardingRule!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * An 128-bit global unique ID of the PSC service attachment.
     */
    public /*out*/ readonly pscServiceAttachmentId!: pulumi.Output<outputs.compute.alpha.Uint128Response>;
    /**
     * This flag determines whether a consumer accept/reject list change can reconcile the statuses of existing ACCEPTED or REJECTED PSC endpoints. - If false, connection policy update will only affect existing PENDING PSC endpoints. Existing ACCEPTED/REJECTED endpoints will remain untouched regardless how the connection policy is modified . - If true, update will affect both PENDING and ACCEPTED/REJECTED PSC endpoints. For example, an ACCEPTED PSC endpoint will be moved to REJECTED if its project is added to the reject list. For newly created service attachment, this boolean defaults to true.
     */
    public readonly reconcileConnections!: pulumi.Output<boolean>;
    public readonly region!: pulumi.Output<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The URL of a service serving the endpoint identified by this service attachment.
     */
    public readonly targetService!: pulumi.Output<string>;

    /**
     * Create a ServiceAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceAttachmentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["connectionPreference"] = args ? args.connectionPreference : undefined;
            resourceInputs["consumerAcceptLists"] = args ? args.consumerAcceptLists : undefined;
            resourceInputs["consumerRejectLists"] = args ? args.consumerRejectLists : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["domainNames"] = args ? args.domainNames : undefined;
            resourceInputs["enableProxyProtocol"] = args ? args.enableProxyProtocol : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["natSubnets"] = args ? args.natSubnets : undefined;
            resourceInputs["producerForwardingRule"] = args ? args.producerForwardingRule : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["reconcileConnections"] = args ? args.reconcileConnections : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["targetService"] = args ? args.targetService : undefined;
            resourceInputs["connectedEndpoints"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["pscServiceAttachmentId"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        } else {
            resourceInputs["connectedEndpoints"] = undefined /*out*/;
            resourceInputs["connectionPreference"] = undefined /*out*/;
            resourceInputs["consumerAcceptLists"] = undefined /*out*/;
            resourceInputs["consumerRejectLists"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["domainNames"] = undefined /*out*/;
            resourceInputs["enableProxyProtocol"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["natSubnets"] = undefined /*out*/;
            resourceInputs["producerForwardingRule"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["pscServiceAttachmentId"] = undefined /*out*/;
            resourceInputs["reconcileConnections"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["targetService"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["project", "region"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ServiceAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServiceAttachment resource.
 */
export interface ServiceAttachmentArgs {
    /**
     * The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
     */
    connectionPreference?: pulumi.Input<enums.compute.alpha.ServiceAttachmentConnectionPreference>;
    /**
     * Projects that are allowed to connect to this service attachment.
     */
    consumerAcceptLists?: pulumi.Input<pulumi.Input<inputs.compute.alpha.ServiceAttachmentConsumerProjectLimitArgs>[]>;
    /**
     * Projects that are not allowed to connect to this service attachment. The project can be specified using its id or number.
     */
    consumerRejectLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * If specified, the domain name will be used during the integration between the PSC connected endpoints and the Cloud DNS. For example, this is a valid domain name: "p.mycompany.com.". Current max number of domain names supported is 1.
     */
    domainNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
     */
    enableProxyProtocol?: pulumi.Input<boolean>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * An array of URLs where each entry is the URL of a subnet provided by the service producer to use for NAT in this service attachment.
     */
    natSubnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The URL of a forwarding rule with loadBalancingScheme INTERNAL* that is serving the endpoint identified by this service attachment.
     */
    producerForwardingRule?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * This flag determines whether a consumer accept/reject list change can reconcile the statuses of existing ACCEPTED or REJECTED PSC endpoints. - If false, connection policy update will only affect existing PENDING PSC endpoints. Existing ACCEPTED/REJECTED endpoints will remain untouched regardless how the connection policy is modified . - If true, update will affect both PENDING and ACCEPTED/REJECTED PSC endpoints. For example, an ACCEPTED PSC endpoint will be moved to REJECTED if its project is added to the reject list. For newly created service attachment, this boolean defaults to true.
     */
    reconcileConnections?: pulumi.Input<boolean>;
    region: pulumi.Input<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * The URL of a service serving the endpoint identified by this service attachment.
     */
    targetService?: pulumi.Input<string>;
}
