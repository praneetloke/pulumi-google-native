// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified Interconnect. Get a list of available Interconnects by making a list() request.
func LookupInterconnect(ctx *pulumi.Context, args *LookupInterconnectArgs, opts ...pulumi.InvokeOption) (*LookupInterconnectResult, error) {
	var rv LookupInterconnectResult
	err := ctx.Invoke("google-native:compute/beta:getInterconnect", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInterconnectArgs struct {
	Interconnect string  `pulumi:"interconnect"`
	Project      *string `pulumi:"project"`
}

type LookupInterconnectResult struct {
	// Administrative status of the interconnect. When this is set to true, the Interconnect is functional and can carry traffic. When set to false, no packets can be carried over the interconnect and no BGP routes are exchanged over it. By default, the status is set to true.
	AdminEnabled bool `pulumi:"adminEnabled"`
	// A list of CircuitInfo objects, that describe the individual circuits in this LAG.
	CircuitInfos []InterconnectCircuitInfoResponse `pulumi:"circuitInfos"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// Customer name, to put in the Letter of Authorization as the party authorized to request a crossconnect.
	CustomerName string `pulumi:"customerName"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// A list of outages expected for this Interconnect.
	ExpectedOutages []InterconnectOutageNotificationResponse `pulumi:"expectedOutages"`
	// IP address configured on the Google side of the Interconnect link. This can be used only for ping tests.
	GoogleIpAddress string `pulumi:"googleIpAddress"`
	// Google reference ID to be used when raising support tickets with Google or otherwise to debug backend connectivity issues.
	GoogleReferenceId string `pulumi:"googleReferenceId"`
	// A list of the URLs of all InterconnectAttachments configured to use this Interconnect.
	InterconnectAttachments []string `pulumi:"interconnectAttachments"`
	// Type of interconnect, which can take one of the following values: - PARTNER: A partner-managed interconnection shared between customers though a partner. - DEDICATED: A dedicated physical interconnection with the customer. Note that a value IT_PRIVATE has been deprecated in favor of DEDICATED.
	InterconnectType string `pulumi:"interconnectType"`
	// Type of the resource. Always compute#interconnect for interconnects.
	Kind string `pulumi:"kind"`
	// A fingerprint for the labels being applied to this Interconnect, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an Interconnect.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Type of link requested, which can take one of the following values: - LINK_TYPE_ETHERNET_10G_LR: A 10G Ethernet with LR optics - LINK_TYPE_ETHERNET_100G_LR: A 100G Ethernet with LR optics. Note that this field indicates the speed of each of the links in the bundle, not the speed of the entire bundle.
	LinkType string `pulumi:"linkType"`
	// URL of the InterconnectLocation object that represents where this connection is to be provisioned.
	Location string `pulumi:"location"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// Email address to contact the customer NOC for operations and maintenance notifications regarding this Interconnect. If specified, this will be used for notifications in addition to all other forms described, such as Cloud Monitoring logs alerting and Cloud Notifications. This field is required for users who sign up for Cloud Interconnect using workforce identity federation.
	NocContactEmail string `pulumi:"nocContactEmail"`
	// The current status of this Interconnect's functionality, which can take one of the following values: - OS_ACTIVE: A valid Interconnect, which is turned up and is ready to use. Attachments may be provisioned on this Interconnect. - OS_UNPROVISIONED: An Interconnect that has not completed turnup. No attachments may be provisioned on this Interconnect. - OS_UNDER_MAINTENANCE: An Interconnect that is undergoing internal maintenance. No attachments may be provisioned or updated on this Interconnect.
	OperationalStatus string `pulumi:"operationalStatus"`
	// IP address configured on the customer side of the Interconnect link. The customer should configure this IP address during turnup when prompted by Google NOC. This can be used only for ping tests.
	PeerIpAddress string `pulumi:"peerIpAddress"`
	// Number of links actually provisioned in this interconnect.
	ProvisionedLinkCount int `pulumi:"provisionedLinkCount"`
	// Target number of physical links in the link bundle, as requested by the customer.
	RequestedLinkCount int `pulumi:"requestedLinkCount"`
	// Reserved for future use.
	SatisfiesPzs bool `pulumi:"satisfiesPzs"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// The current state of Interconnect functionality, which can take one of the following values: - ACTIVE: The Interconnect is valid, turned up and ready to use. Attachments may be provisioned on this Interconnect. - UNPROVISIONED: The Interconnect has not completed turnup. No attachments may be provisioned on this Interconnect. - UNDER_MAINTENANCE: The Interconnect is undergoing internal maintenance. No attachments may be provisioned or updated on this Interconnect.
	State string `pulumi:"state"`
}

func LookupInterconnectOutput(ctx *pulumi.Context, args LookupInterconnectOutputArgs, opts ...pulumi.InvokeOption) LookupInterconnectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInterconnectResult, error) {
			args := v.(LookupInterconnectArgs)
			r, err := LookupInterconnect(ctx, &args, opts...)
			var s LookupInterconnectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInterconnectResultOutput)
}

type LookupInterconnectOutputArgs struct {
	Interconnect pulumi.StringInput    `pulumi:"interconnect"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupInterconnectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInterconnectArgs)(nil)).Elem()
}

type LookupInterconnectResultOutput struct{ *pulumi.OutputState }

func (LookupInterconnectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInterconnectResult)(nil)).Elem()
}

func (o LookupInterconnectResultOutput) ToLookupInterconnectResultOutput() LookupInterconnectResultOutput {
	return o
}

func (o LookupInterconnectResultOutput) ToLookupInterconnectResultOutputWithContext(ctx context.Context) LookupInterconnectResultOutput {
	return o
}

// Administrative status of the interconnect. When this is set to true, the Interconnect is functional and can carry traffic. When set to false, no packets can be carried over the interconnect and no BGP routes are exchanged over it. By default, the status is set to true.
func (o LookupInterconnectResultOutput) AdminEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInterconnectResult) bool { return v.AdminEnabled }).(pulumi.BoolOutput)
}

// A list of CircuitInfo objects, that describe the individual circuits in this LAG.
func (o LookupInterconnectResultOutput) CircuitInfos() InterconnectCircuitInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupInterconnectResult) []InterconnectCircuitInfoResponse { return v.CircuitInfos }).(InterconnectCircuitInfoResponseArrayOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupInterconnectResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// Customer name, to put in the Letter of Authorization as the party authorized to request a crossconnect.
func (o LookupInterconnectResultOutput) CustomerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectResult) string { return v.CustomerName }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupInterconnectResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectResult) string { return v.Description }).(pulumi.StringOutput)
}

// A list of outages expected for this Interconnect.
func (o LookupInterconnectResultOutput) ExpectedOutages() InterconnectOutageNotificationResponseArrayOutput {
	return o.ApplyT(func(v LookupInterconnectResult) []InterconnectOutageNotificationResponse { return v.ExpectedOutages }).(InterconnectOutageNotificationResponseArrayOutput)
}

// IP address configured on the Google side of the Interconnect link. This can be used only for ping tests.
func (o LookupInterconnectResultOutput) GoogleIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectResult) string { return v.GoogleIpAddress }).(pulumi.StringOutput)
}

// Google reference ID to be used when raising support tickets with Google or otherwise to debug backend connectivity issues.
func (o LookupInterconnectResultOutput) GoogleReferenceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectResult) string { return v.GoogleReferenceId }).(pulumi.StringOutput)
}

// A list of the URLs of all InterconnectAttachments configured to use this Interconnect.
func (o LookupInterconnectResultOutput) InterconnectAttachments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInterconnectResult) []string { return v.InterconnectAttachments }).(pulumi.StringArrayOutput)
}

// Type of interconnect, which can take one of the following values: - PARTNER: A partner-managed interconnection shared between customers though a partner. - DEDICATED: A dedicated physical interconnection with the customer. Note that a value IT_PRIVATE has been deprecated in favor of DEDICATED.
func (o LookupInterconnectResultOutput) InterconnectType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectResult) string { return v.InterconnectType }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#interconnect for interconnects.
func (o LookupInterconnectResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectResult) string { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this Interconnect, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an Interconnect.
func (o LookupInterconnectResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
func (o LookupInterconnectResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInterconnectResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Type of link requested, which can take one of the following values: - LINK_TYPE_ETHERNET_10G_LR: A 10G Ethernet with LR optics - LINK_TYPE_ETHERNET_100G_LR: A 100G Ethernet with LR optics. Note that this field indicates the speed of each of the links in the bundle, not the speed of the entire bundle.
func (o LookupInterconnectResultOutput) LinkType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectResult) string { return v.LinkType }).(pulumi.StringOutput)
}

// URL of the InterconnectLocation object that represents where this connection is to be provisioned.
func (o LookupInterconnectResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectResult) string { return v.Location }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupInterconnectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectResult) string { return v.Name }).(pulumi.StringOutput)
}

// Email address to contact the customer NOC for operations and maintenance notifications regarding this Interconnect. If specified, this will be used for notifications in addition to all other forms described, such as Cloud Monitoring logs alerting and Cloud Notifications. This field is required for users who sign up for Cloud Interconnect using workforce identity federation.
func (o LookupInterconnectResultOutput) NocContactEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectResult) string { return v.NocContactEmail }).(pulumi.StringOutput)
}

// The current status of this Interconnect's functionality, which can take one of the following values: - OS_ACTIVE: A valid Interconnect, which is turned up and is ready to use. Attachments may be provisioned on this Interconnect. - OS_UNPROVISIONED: An Interconnect that has not completed turnup. No attachments may be provisioned on this Interconnect. - OS_UNDER_MAINTENANCE: An Interconnect that is undergoing internal maintenance. No attachments may be provisioned or updated on this Interconnect.
func (o LookupInterconnectResultOutput) OperationalStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectResult) string { return v.OperationalStatus }).(pulumi.StringOutput)
}

// IP address configured on the customer side of the Interconnect link. The customer should configure this IP address during turnup when prompted by Google NOC. This can be used only for ping tests.
func (o LookupInterconnectResultOutput) PeerIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectResult) string { return v.PeerIpAddress }).(pulumi.StringOutput)
}

// Number of links actually provisioned in this interconnect.
func (o LookupInterconnectResultOutput) ProvisionedLinkCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInterconnectResult) int { return v.ProvisionedLinkCount }).(pulumi.IntOutput)
}

// Target number of physical links in the link bundle, as requested by the customer.
func (o LookupInterconnectResultOutput) RequestedLinkCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInterconnectResult) int { return v.RequestedLinkCount }).(pulumi.IntOutput)
}

// Reserved for future use.
func (o LookupInterconnectResultOutput) SatisfiesPzs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInterconnectResult) bool { return v.SatisfiesPzs }).(pulumi.BoolOutput)
}

// Server-defined URL for the resource.
func (o LookupInterconnectResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// The current state of Interconnect functionality, which can take one of the following values: - ACTIVE: The Interconnect is valid, turned up and ready to use. Attachments may be provisioned on this Interconnect. - UNPROVISIONED: The Interconnect has not completed turnup. No attachments may be provisioned on this Interconnect. - UNDER_MAINTENANCE: The Interconnect is undergoing internal maintenance. No attachments may be provisioned or updated on this Interconnect.
func (o LookupInterconnectResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInterconnectResultOutput{})
}
