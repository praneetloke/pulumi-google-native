// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified ServiceAttachment resource in the given scope.
func LookupServiceAttachment(ctx *pulumi.Context, args *LookupServiceAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupServiceAttachmentResult, error) {
	var rv LookupServiceAttachmentResult
	err := ctx.Invoke("google-native:compute/beta:getServiceAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceAttachmentArgs struct {
	Project           *string `pulumi:"project"`
	Region            string  `pulumi:"region"`
	ServiceAttachment string  `pulumi:"serviceAttachment"`
}

type LookupServiceAttachmentResult struct {
	// An array of connections for all the consumers connected to this service attachment.
	ConnectedEndpoints []ServiceAttachmentConnectedEndpointResponse `pulumi:"connectedEndpoints"`
	// The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
	ConnectionPreference string `pulumi:"connectionPreference"`
	// Projects that are allowed to connect to this service attachment.
	ConsumerAcceptLists []ServiceAttachmentConsumerProjectLimitResponse `pulumi:"consumerAcceptLists"`
	// An array of forwarding rules for all the consumers connected to this service attachment.
	ConsumerForwardingRules []ServiceAttachmentConsumerForwardingRuleResponse `pulumi:"consumerForwardingRules"`
	// Projects that are not allowed to connect to this service attachment. The project can be specified using its id or number.
	ConsumerRejectLists []string `pulumi:"consumerRejectLists"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
	EnableProxyProtocol bool `pulumi:"enableProxyProtocol"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a ServiceAttachment. An up-to-date fingerprint must be provided in order to patch/update the ServiceAttachment; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the ServiceAttachment.
	Fingerprint string `pulumi:"fingerprint"`
	// Type of the resource. Always compute#serviceAttachment for service attachments.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// An array of URLs where each entry is the URL of a subnet provided by the service producer to use for NAT in this service attachment.
	NatSubnets []string `pulumi:"natSubnets"`
	// The URL of a forwarding rule with loadBalancingScheme INTERNAL* that is serving the endpoint identified by this service attachment.
	ProducerForwardingRule string `pulumi:"producerForwardingRule"`
	// An 128-bit global unique ID of the PSC service attachment.
	PscServiceAttachmentId Uint128Response `pulumi:"pscServiceAttachmentId"`
	// URL of the region where the service attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// The URL of a service serving the endpoint identified by this service attachment.
	TargetService string `pulumi:"targetService"`
}