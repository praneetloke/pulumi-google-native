// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Policy.
type Policy struct {
	pulumi.CustomResourceState

	// Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
	AlternativeNameServerConfig PolicyAlternativeNameServerConfigResponseOutput `pulumi:"alternativeNameServerConfig"`
	// For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
	ClientOperationId pulumi.StringPtrOutput `pulumi:"clientOperationId"`
	// A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the policy's function.
	Description pulumi.StringOutput `pulumi:"description"`
	// Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy.
	EnableInboundForwarding pulumi.BoolOutput `pulumi:"enableInboundForwarding"`
	// Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
	EnableLogging pulumi.BoolOutput   `pulumi:"enableLogging"`
	Kind          pulumi.StringOutput `pulumi:"kind"`
	// User-assigned name for this policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of network names specifying networks to which this policy is applied.
	Networks PolicyNetworkResponseArrayOutput `pulumi:"networks"`
	Project  pulumi.StringOutput              `pulumi:"project"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		args = &PolicyArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Policy
	err := ctx.RegisterResource("google-native:dns/v1beta2:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("google-native:dns/v1beta2:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
}

type PolicyState struct {
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
	AlternativeNameServerConfig *PolicyAlternativeNameServerConfig `pulumi:"alternativeNameServerConfig"`
	// For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
	ClientOperationId *string `pulumi:"clientOperationId"`
	// A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the policy's function.
	Description *string `pulumi:"description"`
	// Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy.
	EnableInboundForwarding *bool `pulumi:"enableInboundForwarding"`
	// Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
	EnableLogging *bool   `pulumi:"enableLogging"`
	Kind          *string `pulumi:"kind"`
	// User-assigned name for this policy.
	Name *string `pulumi:"name"`
	// List of network names specifying networks to which this policy is applied.
	Networks []PolicyNetwork `pulumi:"networks"`
	Project  *string         `pulumi:"project"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
	AlternativeNameServerConfig PolicyAlternativeNameServerConfigPtrInput
	// For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
	ClientOperationId pulumi.StringPtrInput
	// A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the policy's function.
	Description pulumi.StringPtrInput
	// Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy.
	EnableInboundForwarding pulumi.BoolPtrInput
	// Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
	EnableLogging pulumi.BoolPtrInput
	Kind          pulumi.StringPtrInput
	// User-assigned name for this policy.
	Name pulumi.StringPtrInput
	// List of network names specifying networks to which this policy is applied.
	Networks PolicyNetworkArrayInput
	Project  pulumi.StringPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

// Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
func (o PolicyOutput) AlternativeNameServerConfig() PolicyAlternativeNameServerConfigResponseOutput {
	return o.ApplyT(func(v *Policy) PolicyAlternativeNameServerConfigResponseOutput { return v.AlternativeNameServerConfig }).(PolicyAlternativeNameServerConfigResponseOutput)
}

// For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
func (o PolicyOutput) ClientOperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.ClientOperationId }).(pulumi.StringPtrOutput)
}

// A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the policy's function.
func (o PolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy.
func (o PolicyOutput) EnableInboundForwarding() pulumi.BoolOutput {
	return o.ApplyT(func(v *Policy) pulumi.BoolOutput { return v.EnableInboundForwarding }).(pulumi.BoolOutput)
}

// Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
func (o PolicyOutput) EnableLogging() pulumi.BoolOutput {
	return o.ApplyT(func(v *Policy) pulumi.BoolOutput { return v.EnableLogging }).(pulumi.BoolOutput)
}

func (o PolicyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// User-assigned name for this policy.
func (o PolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of network names specifying networks to which this policy is applied.
func (o PolicyOutput) Networks() PolicyNetworkResponseArrayOutput {
	return o.ApplyT(func(v *Policy) PolicyNetworkResponseArrayOutput { return v.Networks }).(PolicyNetworkResponseArrayOutput)
}

func (o PolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyInput)(nil)).Elem(), &Policy{})
	pulumi.RegisterOutputType(PolicyOutput{})
}
