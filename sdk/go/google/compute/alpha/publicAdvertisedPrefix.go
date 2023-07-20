// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a PublicAdvertisedPrefix in the specified project using the parameters that are included in the request.
type PublicAdvertisedPrefix struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// The IPv4 address to be used for reverse DNS verification.
	DnsVerificationIp pulumi.StringOutput `pulumi:"dnsVerificationIp"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicAdvertisedPrefix. An up-to-date fingerprint must be provided in order to update the PublicAdvertisedPrefix, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a PublicAdvertisedPrefix.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The IPv4 address range, in CIDR format, represented by this public advertised prefix.
	IpCidrRange pulumi.StringOutput `pulumi:"ipCidrRange"`
	// Type of the resource. Always compute#publicAdvertisedPrefix for public advertised prefixes.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies how child public delegated prefix will be scoped. It could be one of following values: - `REGIONAL`: The public delegated prefix is regional only. The provisioning will take a few minutes. - `GLOBAL`: The public delegated prefix is global only. The provisioning will take ~4 weeks. - `GLOBAL_AND_REGIONAL` [output only]: The public delegated prefixes is BYOIP V1 legacy prefix. This is output only value and no longer supported in BYOIP V2.
	PdpScope pulumi.StringOutput `pulumi:"pdpScope"`
	Project  pulumi.StringOutput `pulumi:"project"`
	// The list of public delegated prefixes that exist for this public advertised prefix.
	PublicDelegatedPrefixs PublicAdvertisedPrefixPublicDelegatedPrefixResponseArrayOutput `pulumi:"publicDelegatedPrefixs"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL with id for the resource.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// The shared secret to be used for reverse DNS verification.
	SharedSecret pulumi.StringOutput `pulumi:"sharedSecret"`
	// The status of the public advertised prefix. Possible values include: - `INITIAL`: RPKI validation is complete. - `PTR_CONFIGURED`: User has configured the PTR. - `VALIDATED`: Reverse DNS lookup is successful. - `REVERSE_DNS_LOOKUP_FAILED`: Reverse DNS lookup failed. - `PREFIX_CONFIGURATION_IN_PROGRESS`: The prefix is being configured. - `PREFIX_CONFIGURATION_COMPLETE`: The prefix is fully configured. - `PREFIX_REMOVAL_IN_PROGRESS`: The prefix is being removed.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewPublicAdvertisedPrefix registers a new resource with the given unique name, arguments, and options.
func NewPublicAdvertisedPrefix(ctx *pulumi.Context,
	name string, args *PublicAdvertisedPrefixArgs, opts ...pulumi.ResourceOption) (*PublicAdvertisedPrefix, error) {
	if args == nil {
		args = &PublicAdvertisedPrefixArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PublicAdvertisedPrefix
	err := ctx.RegisterResource("google-native:compute/alpha:PublicAdvertisedPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicAdvertisedPrefix gets an existing PublicAdvertisedPrefix resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicAdvertisedPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicAdvertisedPrefixState, opts ...pulumi.ResourceOption) (*PublicAdvertisedPrefix, error) {
	var resource PublicAdvertisedPrefix
	err := ctx.ReadResource("google-native:compute/alpha:PublicAdvertisedPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicAdvertisedPrefix resources.
type publicAdvertisedPrefixState struct {
}

type PublicAdvertisedPrefixState struct {
}

func (PublicAdvertisedPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicAdvertisedPrefixState)(nil)).Elem()
}

type publicAdvertisedPrefixArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// The IPv4 address to be used for reverse DNS verification.
	DnsVerificationIp *string `pulumi:"dnsVerificationIp"`
	// The IPv4 address range, in CIDR format, represented by this public advertised prefix.
	IpCidrRange *string `pulumi:"ipCidrRange"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Specifies how child public delegated prefix will be scoped. It could be one of following values: - `REGIONAL`: The public delegated prefix is regional only. The provisioning will take a few minutes. - `GLOBAL`: The public delegated prefix is global only. The provisioning will take ~4 weeks. - `GLOBAL_AND_REGIONAL` [output only]: The public delegated prefixes is BYOIP V1 legacy prefix. This is output only value and no longer supported in BYOIP V2.
	PdpScope *PublicAdvertisedPrefixPdpScope `pulumi:"pdpScope"`
	Project  *string                         `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// The status of the public advertised prefix. Possible values include: - `INITIAL`: RPKI validation is complete. - `PTR_CONFIGURED`: User has configured the PTR. - `VALIDATED`: Reverse DNS lookup is successful. - `REVERSE_DNS_LOOKUP_FAILED`: Reverse DNS lookup failed. - `PREFIX_CONFIGURATION_IN_PROGRESS`: The prefix is being configured. - `PREFIX_CONFIGURATION_COMPLETE`: The prefix is fully configured. - `PREFIX_REMOVAL_IN_PROGRESS`: The prefix is being removed.
	Status *PublicAdvertisedPrefixStatus `pulumi:"status"`
}

// The set of arguments for constructing a PublicAdvertisedPrefix resource.
type PublicAdvertisedPrefixArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// The IPv4 address to be used for reverse DNS verification.
	DnsVerificationIp pulumi.StringPtrInput
	// The IPv4 address range, in CIDR format, represented by this public advertised prefix.
	IpCidrRange pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Specifies how child public delegated prefix will be scoped. It could be one of following values: - `REGIONAL`: The public delegated prefix is regional only. The provisioning will take a few minutes. - `GLOBAL`: The public delegated prefix is global only. The provisioning will take ~4 weeks. - `GLOBAL_AND_REGIONAL` [output only]: The public delegated prefixes is BYOIP V1 legacy prefix. This is output only value and no longer supported in BYOIP V2.
	PdpScope PublicAdvertisedPrefixPdpScopePtrInput
	Project  pulumi.StringPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// The status of the public advertised prefix. Possible values include: - `INITIAL`: RPKI validation is complete. - `PTR_CONFIGURED`: User has configured the PTR. - `VALIDATED`: Reverse DNS lookup is successful. - `REVERSE_DNS_LOOKUP_FAILED`: Reverse DNS lookup failed. - `PREFIX_CONFIGURATION_IN_PROGRESS`: The prefix is being configured. - `PREFIX_CONFIGURATION_COMPLETE`: The prefix is fully configured. - `PREFIX_REMOVAL_IN_PROGRESS`: The prefix is being removed.
	Status PublicAdvertisedPrefixStatusPtrInput
}

func (PublicAdvertisedPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicAdvertisedPrefixArgs)(nil)).Elem()
}

type PublicAdvertisedPrefixInput interface {
	pulumi.Input

	ToPublicAdvertisedPrefixOutput() PublicAdvertisedPrefixOutput
	ToPublicAdvertisedPrefixOutputWithContext(ctx context.Context) PublicAdvertisedPrefixOutput
}

func (*PublicAdvertisedPrefix) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicAdvertisedPrefix)(nil)).Elem()
}

func (i *PublicAdvertisedPrefix) ToPublicAdvertisedPrefixOutput() PublicAdvertisedPrefixOutput {
	return i.ToPublicAdvertisedPrefixOutputWithContext(context.Background())
}

func (i *PublicAdvertisedPrefix) ToPublicAdvertisedPrefixOutputWithContext(ctx context.Context) PublicAdvertisedPrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicAdvertisedPrefixOutput)
}

type PublicAdvertisedPrefixOutput struct{ *pulumi.OutputState }

func (PublicAdvertisedPrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicAdvertisedPrefix)(nil)).Elem()
}

func (o PublicAdvertisedPrefixOutput) ToPublicAdvertisedPrefixOutput() PublicAdvertisedPrefixOutput {
	return o
}

func (o PublicAdvertisedPrefixOutput) ToPublicAdvertisedPrefixOutputWithContext(ctx context.Context) PublicAdvertisedPrefixOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o PublicAdvertisedPrefixOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAdvertisedPrefix) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o PublicAdvertisedPrefixOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAdvertisedPrefix) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The IPv4 address to be used for reverse DNS verification.
func (o PublicAdvertisedPrefixOutput) DnsVerificationIp() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAdvertisedPrefix) pulumi.StringOutput { return v.DnsVerificationIp }).(pulumi.StringOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicAdvertisedPrefix. An up-to-date fingerprint must be provided in order to update the PublicAdvertisedPrefix, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a PublicAdvertisedPrefix.
func (o PublicAdvertisedPrefixOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAdvertisedPrefix) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// The IPv4 address range, in CIDR format, represented by this public advertised prefix.
func (o PublicAdvertisedPrefixOutput) IpCidrRange() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAdvertisedPrefix) pulumi.StringOutput { return v.IpCidrRange }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#publicAdvertisedPrefix for public advertised prefixes.
func (o PublicAdvertisedPrefixOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAdvertisedPrefix) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o PublicAdvertisedPrefixOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAdvertisedPrefix) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies how child public delegated prefix will be scoped. It could be one of following values: - `REGIONAL`: The public delegated prefix is regional only. The provisioning will take a few minutes. - `GLOBAL`: The public delegated prefix is global only. The provisioning will take ~4 weeks. - `GLOBAL_AND_REGIONAL` [output only]: The public delegated prefixes is BYOIP V1 legacy prefix. This is output only value and no longer supported in BYOIP V2.
func (o PublicAdvertisedPrefixOutput) PdpScope() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAdvertisedPrefix) pulumi.StringOutput { return v.PdpScope }).(pulumi.StringOutput)
}

func (o PublicAdvertisedPrefixOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAdvertisedPrefix) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The list of public delegated prefixes that exist for this public advertised prefix.
func (o PublicAdvertisedPrefixOutput) PublicDelegatedPrefixs() PublicAdvertisedPrefixPublicDelegatedPrefixResponseArrayOutput {
	return o.ApplyT(func(v *PublicAdvertisedPrefix) PublicAdvertisedPrefixPublicDelegatedPrefixResponseArrayOutput {
		return v.PublicDelegatedPrefixs
	}).(PublicAdvertisedPrefixPublicDelegatedPrefixResponseArrayOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o PublicAdvertisedPrefixOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicAdvertisedPrefix) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Server-defined URL for the resource.
func (o PublicAdvertisedPrefixOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAdvertisedPrefix) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL with id for the resource.
func (o PublicAdvertisedPrefixOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAdvertisedPrefix) pulumi.StringOutput { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// The shared secret to be used for reverse DNS verification.
func (o PublicAdvertisedPrefixOutput) SharedSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAdvertisedPrefix) pulumi.StringOutput { return v.SharedSecret }).(pulumi.StringOutput)
}

// The status of the public advertised prefix. Possible values include: - `INITIAL`: RPKI validation is complete. - `PTR_CONFIGURED`: User has configured the PTR. - `VALIDATED`: Reverse DNS lookup is successful. - `REVERSE_DNS_LOOKUP_FAILED`: Reverse DNS lookup failed. - `PREFIX_CONFIGURATION_IN_PROGRESS`: The prefix is being configured. - `PREFIX_CONFIGURATION_COMPLETE`: The prefix is fully configured. - `PREFIX_REMOVAL_IN_PROGRESS`: The prefix is being removed.
func (o PublicAdvertisedPrefixOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAdvertisedPrefix) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicAdvertisedPrefixInput)(nil)).Elem(), &PublicAdvertisedPrefix{})
	pulumi.RegisterOutputType(PublicAdvertisedPrefixOutput{})
}
