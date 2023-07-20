// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified interconnect attachment.
func LookupInterconnectAttachment(ctx *pulumi.Context, args *LookupInterconnectAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupInterconnectAttachmentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInterconnectAttachmentResult
	err := ctx.Invoke("google-native:compute/v1:getInterconnectAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInterconnectAttachmentArgs struct {
	InterconnectAttachment string  `pulumi:"interconnectAttachment"`
	Project                *string `pulumi:"project"`
	Region                 string  `pulumi:"region"`
}

type LookupInterconnectAttachmentResult struct {
	// Determines whether this Attachment will carry packets. Not present for PARTNER_PROVIDER.
	AdminEnabled bool `pulumi:"adminEnabled"`
	// Provisioned bandwidth capacity for the interconnect attachment. For attachments of type DEDICATED, the user can set the bandwidth. For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth. Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED, and can take one of the following values: - BPS_50M: 50 Mbit/s - BPS_100M: 100 Mbit/s - BPS_200M: 200 Mbit/s - BPS_300M: 300 Mbit/s - BPS_400M: 400 Mbit/s - BPS_500M: 500 Mbit/s - BPS_1G: 1 Gbit/s - BPS_2G: 2 Gbit/s - BPS_5G: 5 Gbit/s - BPS_10G: 10 Gbit/s - BPS_20G: 20 Gbit/s - BPS_50G: 50 Gbit/s
	Bandwidth string `pulumi:"bandwidth"`
	// This field is not available.
	CandidateIpv6Subnets []string `pulumi:"candidateIpv6Subnets"`
	// Up to 16 candidate prefixes that can be used to restrict the allocation of cloudRouterIpAddress and customerRouterIpAddress for this attachment. All prefixes must be within link-local address space (169.254.0.0/16) and must be /29 or shorter (/28, /27, etc). Google will attempt to select an unused /29 from the supplied candidate prefix(es). The request will fail if all possible /29s are in use on Google's edge. If not supplied, Google will randomly select an unused /29 from all of link-local space.
	CandidateSubnets []string `pulumi:"candidateSubnets"`
	// IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
	CloudRouterIpAddress string `pulumi:"cloudRouterIpAddress"`
	// IPv6 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
	CloudRouterIpv6Address string `pulumi:"cloudRouterIpv6Address"`
	// This field is not available.
	CloudRouterIpv6InterfaceId string `pulumi:"cloudRouterIpv6InterfaceId"`
	// Constraints for this attachment, if any. The attachment does not work if these constraints are not met.
	ConfigurationConstraints InterconnectAttachmentConfigurationConstraintsResponse `pulumi:"configurationConstraints"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
	CustomerRouterIpAddress string `pulumi:"customerRouterIpAddress"`
	// IPv6 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
	CustomerRouterIpv6Address string `pulumi:"customerRouterIpv6Address"`
	// This field is not available.
	CustomerRouterIpv6InterfaceId string `pulumi:"customerRouterIpv6InterfaceId"`
	// Dataplane version for this InterconnectAttachment. This field is only present for Dataplane version 2 and higher. Absence of this field in the API output indicates that the Dataplane is version 1.
	DataplaneVersion int `pulumi:"dataplaneVersion"`
	// An optional description of this resource.
	Description string `pulumi:"description"`
	// Desired availability domain for the attachment. Only available for type PARTNER, at creation time, and can take one of the following values: - AVAILABILITY_DOMAIN_ANY - AVAILABILITY_DOMAIN_1 - AVAILABILITY_DOMAIN_2 For improved reliability, customers should configure a pair of attachments, one per availability domain. The selected availability domain will be provided to the Partner via the pairing key, so that the provisioned circuit will lie in the specified domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
	EdgeAvailabilityDomain string `pulumi:"edgeAvailabilityDomain"`
	// Indicates the user-supplied encryption option of this VLAN attachment (interconnectAttachment). Can only be specified at attachment creation for PARTNER or DEDICATED attachments. Possible values are: - NONE - This is the default value, which means that the VLAN attachment carries unencrypted traffic. VMs are able to send traffic to, or receive traffic from, such a VLAN attachment. - IPSEC - The VLAN attachment carries only encrypted traffic that is encrypted by an IPsec device, such as an HA VPN gateway or third-party IPsec VPN. VMs cannot directly send traffic to, or receive traffic from, such a VLAN attachment. To use *HA VPN over Cloud Interconnect*, the VLAN attachment must be created with this option.
	Encryption string `pulumi:"encryption"`
	// Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity issues. [Deprecated] This field is not used.
	//
	// Deprecated: [Output Only] Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity issues. [Deprecated] This field is not used.
	GoogleReferenceId string `pulumi:"googleReferenceId"`
	// URL of the underlying Interconnect object that this attachment's traffic will traverse through.
	Interconnect string `pulumi:"interconnect"`
	// A list of URLs of addresses that have been reserved for the VLAN attachment. Used only for the VLAN attachment that has the encryption option as IPSEC. The addresses must be regional internal IP address ranges. When creating an HA VPN gateway over the VLAN attachment, if the attachment is configured to use a regional internal IP address, then the VPN gateway's IP address is allocated from the IP address range specified here. For example, if the HA VPN gateway's interface 0 is paired to this VLAN attachment, then a regional internal IP address for the VPN gateway interface 0 will be allocated from the IP address specified for this VLAN attachment. If this field is not specified when creating the VLAN attachment, then later on when creating an HA VPN gateway on this VLAN attachment, the HA VPN gateway's IP address is allocated from the regional external IP address pool.
	IpsecInternalAddresses []string `pulumi:"ipsecInternalAddresses"`
	// Type of the resource. Always compute#interconnectAttachment for interconnect attachments.
	Kind string `pulumi:"kind"`
	// A fingerprint for the labels being applied to this InterconnectAttachment, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an InterconnectAttachment.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Maximum Transmission Unit (MTU), in bytes, of packets passing through this interconnect attachment. Only 1440 and 1500 are allowed. If not specified, the value will default to 1440.
	Mtu int `pulumi:"mtu"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// The current status of whether or not this interconnect attachment is functional, which can take one of the following values: - OS_ACTIVE: The attachment has been turned up and is ready to use. - OS_UNPROVISIONED: The attachment is not ready to use yet, because turnup is not complete.
	OperationalStatus string `pulumi:"operationalStatus"`
	// [Output only for type PARTNER. Input only for PARTNER_PROVIDER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
	PairingKey string `pulumi:"pairingKey"`
	// Optional BGP ASN for the router supplied by a Layer 3 Partner if they configured BGP on behalf of the customer. Output only for PARTNER type, input only for PARTNER_PROVIDER, not available for DEDICATED.
	PartnerAsn string `pulumi:"partnerAsn"`
	// Informational metadata about Partner attachments from Partners to display to customers. Output only for for PARTNER type, mutable for PARTNER_PROVIDER, not available for DEDICATED.
	PartnerMetadata InterconnectAttachmentPartnerMetadataResponse `pulumi:"partnerMetadata"`
	// Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached to is of type DEDICATED.
	PrivateInterconnectInfo InterconnectAttachmentPrivateInfoResponse `pulumi:"privateInterconnectInfo"`
	// URL of the region where the regional interconnect attachment resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region string `pulumi:"region"`
	// If the attachment is on a Cross-Cloud Interconnect connection, this field contains the interconnect's remote location service provider. Example values: "Amazon Web Services" "Microsoft Azure". The field is set only for attachments on Cross-Cloud Interconnect connections. Its value is copied from the InterconnectRemoteLocation remoteService field.
	RemoteService string `pulumi:"remoteService"`
	// URL of the Cloud Router to be used for dynamic routing. This router must be in the same region as this InterconnectAttachment. The InterconnectAttachment will automatically connect the Interconnect to the network & region within which the Cloud Router is configured.
	Router string `pulumi:"router"`
	// Reserved for future use.
	SatisfiesPzs bool `pulumi:"satisfiesPzs"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// The stack type for this interconnect attachment to identify whether the IPv6 feature is enabled or not. If not specified, IPV4_ONLY will be used. This field can be both set at interconnect attachments creation and update interconnect attachment operations.
	StackType string `pulumi:"stackType"`
	// The current state of this attachment's functionality. Enum values ACTIVE and UNPROVISIONED are shared by DEDICATED/PRIVATE, PARTNER, and PARTNER_PROVIDER interconnect attachments, while enum values PENDING_PARTNER, PARTNER_REQUEST_RECEIVED, and PENDING_CUSTOMER are used for only PARTNER and PARTNER_PROVIDER interconnect attachments. This state can take one of the following values: - ACTIVE: The attachment has been turned up and is ready to use. - UNPROVISIONED: The attachment is not ready to use yet, because turnup is not complete. - PENDING_PARTNER: A newly-created PARTNER attachment that has not yet been configured on the Partner side. - PARTNER_REQUEST_RECEIVED: A PARTNER attachment is in the process of provisioning after a PARTNER_PROVIDER attachment was created that references it. - PENDING_CUSTOMER: A PARTNER or PARTNER_PROVIDER attachment that is waiting for a customer to activate it. - DEFUNCT: The attachment was deleted externally and is no longer functional. This could be because the associated Interconnect was removed, or because the other side of a Partner attachment was deleted.
	State string `pulumi:"state"`
	// Length of the IPv4 subnet mask. Allowed values: - 29 (default) - 30 The default value is 29, except for Cross-Cloud Interconnect connections that use an InterconnectRemoteLocation with a constraints.subnetLengthRange.min equal to 30. For example, connections that use an Azure remote location fall into this category. In these cases, the default value is 30, and requesting 29 returns an error. Where both 29 and 30 are allowed, 29 is preferred, because it gives Google Cloud Support more debugging visibility.
	SubnetLength int `pulumi:"subnetLength"`
	// The type of interconnect attachment this is, which can take one of the following values: - DEDICATED: an attachment to a Dedicated Interconnect. - PARTNER: an attachment to a Partner Interconnect, created by the customer. - PARTNER_PROVIDER: an attachment to a Partner Interconnect, created by the partner.
	Type string `pulumi:"type"`
	// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4093. Only specified at creation time.
	VlanTag8021q int `pulumi:"vlanTag8021q"`
}

func LookupInterconnectAttachmentOutput(ctx *pulumi.Context, args LookupInterconnectAttachmentOutputArgs, opts ...pulumi.InvokeOption) LookupInterconnectAttachmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInterconnectAttachmentResult, error) {
			args := v.(LookupInterconnectAttachmentArgs)
			r, err := LookupInterconnectAttachment(ctx, &args, opts...)
			var s LookupInterconnectAttachmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInterconnectAttachmentResultOutput)
}

type LookupInterconnectAttachmentOutputArgs struct {
	InterconnectAttachment pulumi.StringInput    `pulumi:"interconnectAttachment"`
	Project                pulumi.StringPtrInput `pulumi:"project"`
	Region                 pulumi.StringInput    `pulumi:"region"`
}

func (LookupInterconnectAttachmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInterconnectAttachmentArgs)(nil)).Elem()
}

type LookupInterconnectAttachmentResultOutput struct{ *pulumi.OutputState }

func (LookupInterconnectAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInterconnectAttachmentResult)(nil)).Elem()
}

func (o LookupInterconnectAttachmentResultOutput) ToLookupInterconnectAttachmentResultOutput() LookupInterconnectAttachmentResultOutput {
	return o
}

func (o LookupInterconnectAttachmentResultOutput) ToLookupInterconnectAttachmentResultOutputWithContext(ctx context.Context) LookupInterconnectAttachmentResultOutput {
	return o
}

// Determines whether this Attachment will carry packets. Not present for PARTNER_PROVIDER.
func (o LookupInterconnectAttachmentResultOutput) AdminEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) bool { return v.AdminEnabled }).(pulumi.BoolOutput)
}

// Provisioned bandwidth capacity for the interconnect attachment. For attachments of type DEDICATED, the user can set the bandwidth. For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth. Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED, and can take one of the following values: - BPS_50M: 50 Mbit/s - BPS_100M: 100 Mbit/s - BPS_200M: 200 Mbit/s - BPS_300M: 300 Mbit/s - BPS_400M: 400 Mbit/s - BPS_500M: 500 Mbit/s - BPS_1G: 1 Gbit/s - BPS_2G: 2 Gbit/s - BPS_5G: 5 Gbit/s - BPS_10G: 10 Gbit/s - BPS_20G: 20 Gbit/s - BPS_50G: 50 Gbit/s
func (o LookupInterconnectAttachmentResultOutput) Bandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.Bandwidth }).(pulumi.StringOutput)
}

// This field is not available.
func (o LookupInterconnectAttachmentResultOutput) CandidateIpv6Subnets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) []string { return v.CandidateIpv6Subnets }).(pulumi.StringArrayOutput)
}

// Up to 16 candidate prefixes that can be used to restrict the allocation of cloudRouterIpAddress and customerRouterIpAddress for this attachment. All prefixes must be within link-local address space (169.254.0.0/16) and must be /29 or shorter (/28, /27, etc). Google will attempt to select an unused /29 from the supplied candidate prefix(es). The request will fail if all possible /29s are in use on Google's edge. If not supplied, Google will randomly select an unused /29 from all of link-local space.
func (o LookupInterconnectAttachmentResultOutput) CandidateSubnets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) []string { return v.CandidateSubnets }).(pulumi.StringArrayOutput)
}

// IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
func (o LookupInterconnectAttachmentResultOutput) CloudRouterIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.CloudRouterIpAddress }).(pulumi.StringOutput)
}

// IPv6 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
func (o LookupInterconnectAttachmentResultOutput) CloudRouterIpv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.CloudRouterIpv6Address }).(pulumi.StringOutput)
}

// This field is not available.
func (o LookupInterconnectAttachmentResultOutput) CloudRouterIpv6InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.CloudRouterIpv6InterfaceId }).(pulumi.StringOutput)
}

// Constraints for this attachment, if any. The attachment does not work if these constraints are not met.
func (o LookupInterconnectAttachmentResultOutput) ConfigurationConstraints() InterconnectAttachmentConfigurationConstraintsResponseOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) InterconnectAttachmentConfigurationConstraintsResponse {
		return v.ConfigurationConstraints
	}).(InterconnectAttachmentConfigurationConstraintsResponseOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupInterconnectAttachmentResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
func (o LookupInterconnectAttachmentResultOutput) CustomerRouterIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.CustomerRouterIpAddress }).(pulumi.StringOutput)
}

// IPv6 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
func (o LookupInterconnectAttachmentResultOutput) CustomerRouterIpv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.CustomerRouterIpv6Address }).(pulumi.StringOutput)
}

// This field is not available.
func (o LookupInterconnectAttachmentResultOutput) CustomerRouterIpv6InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.CustomerRouterIpv6InterfaceId }).(pulumi.StringOutput)
}

// Dataplane version for this InterconnectAttachment. This field is only present for Dataplane version 2 and higher. Absence of this field in the API output indicates that the Dataplane is version 1.
func (o LookupInterconnectAttachmentResultOutput) DataplaneVersion() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) int { return v.DataplaneVersion }).(pulumi.IntOutput)
}

// An optional description of this resource.
func (o LookupInterconnectAttachmentResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.Description }).(pulumi.StringOutput)
}

// Desired availability domain for the attachment. Only available for type PARTNER, at creation time, and can take one of the following values: - AVAILABILITY_DOMAIN_ANY - AVAILABILITY_DOMAIN_1 - AVAILABILITY_DOMAIN_2 For improved reliability, customers should configure a pair of attachments, one per availability domain. The selected availability domain will be provided to the Partner via the pairing key, so that the provisioned circuit will lie in the specified domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
func (o LookupInterconnectAttachmentResultOutput) EdgeAvailabilityDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.EdgeAvailabilityDomain }).(pulumi.StringOutput)
}

// Indicates the user-supplied encryption option of this VLAN attachment (interconnectAttachment). Can only be specified at attachment creation for PARTNER or DEDICATED attachments. Possible values are: - NONE - This is the default value, which means that the VLAN attachment carries unencrypted traffic. VMs are able to send traffic to, or receive traffic from, such a VLAN attachment. - IPSEC - The VLAN attachment carries only encrypted traffic that is encrypted by an IPsec device, such as an HA VPN gateway or third-party IPsec VPN. VMs cannot directly send traffic to, or receive traffic from, such a VLAN attachment. To use *HA VPN over Cloud Interconnect*, the VLAN attachment must be created with this option.
func (o LookupInterconnectAttachmentResultOutput) Encryption() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.Encryption }).(pulumi.StringOutput)
}

// Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity issues. [Deprecated] This field is not used.
//
// Deprecated: [Output Only] Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity issues. [Deprecated] This field is not used.
func (o LookupInterconnectAttachmentResultOutput) GoogleReferenceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.GoogleReferenceId }).(pulumi.StringOutput)
}

// URL of the underlying Interconnect object that this attachment's traffic will traverse through.
func (o LookupInterconnectAttachmentResultOutput) Interconnect() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.Interconnect }).(pulumi.StringOutput)
}

// A list of URLs of addresses that have been reserved for the VLAN attachment. Used only for the VLAN attachment that has the encryption option as IPSEC. The addresses must be regional internal IP address ranges. When creating an HA VPN gateway over the VLAN attachment, if the attachment is configured to use a regional internal IP address, then the VPN gateway's IP address is allocated from the IP address range specified here. For example, if the HA VPN gateway's interface 0 is paired to this VLAN attachment, then a regional internal IP address for the VPN gateway interface 0 will be allocated from the IP address specified for this VLAN attachment. If this field is not specified when creating the VLAN attachment, then later on when creating an HA VPN gateway on this VLAN attachment, the HA VPN gateway's IP address is allocated from the regional external IP address pool.
func (o LookupInterconnectAttachmentResultOutput) IpsecInternalAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) []string { return v.IpsecInternalAddresses }).(pulumi.StringArrayOutput)
}

// Type of the resource. Always compute#interconnectAttachment for interconnect attachments.
func (o LookupInterconnectAttachmentResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this InterconnectAttachment, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an InterconnectAttachment.
func (o LookupInterconnectAttachmentResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
func (o LookupInterconnectAttachmentResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Maximum Transmission Unit (MTU), in bytes, of packets passing through this interconnect attachment. Only 1440 and 1500 are allowed. If not specified, the value will default to 1440.
func (o LookupInterconnectAttachmentResultOutput) Mtu() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) int { return v.Mtu }).(pulumi.IntOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupInterconnectAttachmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current status of whether or not this interconnect attachment is functional, which can take one of the following values: - OS_ACTIVE: The attachment has been turned up and is ready to use. - OS_UNPROVISIONED: The attachment is not ready to use yet, because turnup is not complete.
func (o LookupInterconnectAttachmentResultOutput) OperationalStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.OperationalStatus }).(pulumi.StringOutput)
}

// [Output only for type PARTNER. Input only for PARTNER_PROVIDER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
func (o LookupInterconnectAttachmentResultOutput) PairingKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.PairingKey }).(pulumi.StringOutput)
}

// Optional BGP ASN for the router supplied by a Layer 3 Partner if they configured BGP on behalf of the customer. Output only for PARTNER type, input only for PARTNER_PROVIDER, not available for DEDICATED.
func (o LookupInterconnectAttachmentResultOutput) PartnerAsn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.PartnerAsn }).(pulumi.StringOutput)
}

// Informational metadata about Partner attachments from Partners to display to customers. Output only for for PARTNER type, mutable for PARTNER_PROVIDER, not available for DEDICATED.
func (o LookupInterconnectAttachmentResultOutput) PartnerMetadata() InterconnectAttachmentPartnerMetadataResponseOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) InterconnectAttachmentPartnerMetadataResponse {
		return v.PartnerMetadata
	}).(InterconnectAttachmentPartnerMetadataResponseOutput)
}

// Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached to is of type DEDICATED.
func (o LookupInterconnectAttachmentResultOutput) PrivateInterconnectInfo() InterconnectAttachmentPrivateInfoResponseOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) InterconnectAttachmentPrivateInfoResponse {
		return v.PrivateInterconnectInfo
	}).(InterconnectAttachmentPrivateInfoResponseOutput)
}

// URL of the region where the regional interconnect attachment resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupInterconnectAttachmentResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.Region }).(pulumi.StringOutput)
}

// If the attachment is on a Cross-Cloud Interconnect connection, this field contains the interconnect's remote location service provider. Example values: "Amazon Web Services" "Microsoft Azure". The field is set only for attachments on Cross-Cloud Interconnect connections. Its value is copied from the InterconnectRemoteLocation remoteService field.
func (o LookupInterconnectAttachmentResultOutput) RemoteService() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.RemoteService }).(pulumi.StringOutput)
}

// URL of the Cloud Router to be used for dynamic routing. This router must be in the same region as this InterconnectAttachment. The InterconnectAttachment will automatically connect the Interconnect to the network & region within which the Cloud Router is configured.
func (o LookupInterconnectAttachmentResultOutput) Router() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.Router }).(pulumi.StringOutput)
}

// Reserved for future use.
func (o LookupInterconnectAttachmentResultOutput) SatisfiesPzs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) bool { return v.SatisfiesPzs }).(pulumi.BoolOutput)
}

// Server-defined URL for the resource.
func (o LookupInterconnectAttachmentResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// The stack type for this interconnect attachment to identify whether the IPv6 feature is enabled or not. If not specified, IPV4_ONLY will be used. This field can be both set at interconnect attachments creation and update interconnect attachment operations.
func (o LookupInterconnectAttachmentResultOutput) StackType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.StackType }).(pulumi.StringOutput)
}

// The current state of this attachment's functionality. Enum values ACTIVE and UNPROVISIONED are shared by DEDICATED/PRIVATE, PARTNER, and PARTNER_PROVIDER interconnect attachments, while enum values PENDING_PARTNER, PARTNER_REQUEST_RECEIVED, and PENDING_CUSTOMER are used for only PARTNER and PARTNER_PROVIDER interconnect attachments. This state can take one of the following values: - ACTIVE: The attachment has been turned up and is ready to use. - UNPROVISIONED: The attachment is not ready to use yet, because turnup is not complete. - PENDING_PARTNER: A newly-created PARTNER attachment that has not yet been configured on the Partner side. - PARTNER_REQUEST_RECEIVED: A PARTNER attachment is in the process of provisioning after a PARTNER_PROVIDER attachment was created that references it. - PENDING_CUSTOMER: A PARTNER or PARTNER_PROVIDER attachment that is waiting for a customer to activate it. - DEFUNCT: The attachment was deleted externally and is no longer functional. This could be because the associated Interconnect was removed, or because the other side of a Partner attachment was deleted.
func (o LookupInterconnectAttachmentResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.State }).(pulumi.StringOutput)
}

// Length of the IPv4 subnet mask. Allowed values: - 29 (default) - 30 The default value is 29, except for Cross-Cloud Interconnect connections that use an InterconnectRemoteLocation with a constraints.subnetLengthRange.min equal to 30. For example, connections that use an Azure remote location fall into this category. In these cases, the default value is 30, and requesting 29 returns an error. Where both 29 and 30 are allowed, 29 is preferred, because it gives Google Cloud Support more debugging visibility.
func (o LookupInterconnectAttachmentResultOutput) SubnetLength() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) int { return v.SubnetLength }).(pulumi.IntOutput)
}

// The type of interconnect attachment this is, which can take one of the following values: - DEDICATED: an attachment to a Dedicated Interconnect. - PARTNER: an attachment to a Partner Interconnect, created by the customer. - PARTNER_PROVIDER: an attachment to a Partner Interconnect, created by the partner.
func (o LookupInterconnectAttachmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) string { return v.Type }).(pulumi.StringOutput)
}

// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4093. Only specified at creation time.
func (o LookupInterconnectAttachmentResultOutput) VlanTag8021q() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInterconnectAttachmentResult) int { return v.VlanTag8021q }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInterconnectAttachmentResultOutput{})
}
