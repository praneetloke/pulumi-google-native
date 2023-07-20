// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new network firewall policy in the specified project and region.
type RegionNetworkFirewallPolicy struct {
	pulumi.CustomResourceState

	// A list of associations that belong to this firewall policy.
	Associations FirewallPolicyAssociationResponseArrayOutput `pulumi:"associations"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Deprecated, please use short name instead. User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This field is not applicable to network firewall policies. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	//
	// Deprecated: Deprecated, please use short name instead. User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This field is not applicable to network firewall policies. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the firewall policy.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// [Output only] Type of the resource. Always compute#firewallPolicyfor firewall policies
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. For Organization Firewall Policies it's a [Output Only] numeric ID allocated by Google Cloud which uniquely identifies the Organization Firewall Policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent of the firewall policy. This field is not applicable to network firewall policies.
	Parent  pulumi.StringOutput `pulumi:"parent"`
	Project pulumi.StringOutput `pulumi:"project"`
	Region  pulumi.StringOutput `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount pulumi.IntOutput `pulumi:"ruleTupleCount"`
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
	Rules FirewallPolicyRuleResponseArrayOutput `pulumi:"rules"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This field is not applicable to network firewall policies. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	ShortName pulumi.StringOutput `pulumi:"shortName"`
}

// NewRegionNetworkFirewallPolicy registers a new resource with the given unique name, arguments, and options.
func NewRegionNetworkFirewallPolicy(ctx *pulumi.Context,
	name string, args *RegionNetworkFirewallPolicyArgs, opts ...pulumi.ResourceOption) (*RegionNetworkFirewallPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
		"region",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RegionNetworkFirewallPolicy
	err := ctx.RegisterResource("google-native:compute/v1:RegionNetworkFirewallPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionNetworkFirewallPolicy gets an existing RegionNetworkFirewallPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionNetworkFirewallPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionNetworkFirewallPolicyState, opts ...pulumi.ResourceOption) (*RegionNetworkFirewallPolicy, error) {
	var resource RegionNetworkFirewallPolicy
	err := ctx.ReadResource("google-native:compute/v1:RegionNetworkFirewallPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionNetworkFirewallPolicy resources.
type regionNetworkFirewallPolicyState struct {
}

type RegionNetworkFirewallPolicyState struct {
}

func (RegionNetworkFirewallPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionNetworkFirewallPolicyState)(nil)).Elem()
}

type regionNetworkFirewallPolicyArgs struct {
	// A list of associations that belong to this firewall policy.
	Associations []FirewallPolicyAssociation `pulumi:"associations"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Deprecated, please use short name instead. User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This field is not applicable to network firewall policies. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	//
	// Deprecated: Deprecated, please use short name instead. User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This field is not applicable to network firewall policies. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName *string `pulumi:"displayName"`
	// Name of the resource. For Organization Firewall Policies it's a [Output Only] numeric ID allocated by Google Cloud which uniquely identifies the Organization Firewall Policy.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	Region  string  `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
	Rules []FirewallPolicyRule `pulumi:"rules"`
	// User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This field is not applicable to network firewall policies. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	ShortName *string `pulumi:"shortName"`
}

// The set of arguments for constructing a RegionNetworkFirewallPolicy resource.
type RegionNetworkFirewallPolicyArgs struct {
	// A list of associations that belong to this firewall policy.
	Associations FirewallPolicyAssociationArrayInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Deprecated, please use short name instead. User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This field is not applicable to network firewall policies. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	//
	// Deprecated: Deprecated, please use short name instead. User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This field is not applicable to network firewall policies. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName pulumi.StringPtrInput
	// Name of the resource. For Organization Firewall Policies it's a [Output Only] numeric ID allocated by Google Cloud which uniquely identifies the Organization Firewall Policy.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	Region  pulumi.StringInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
	Rules FirewallPolicyRuleArrayInput
	// User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This field is not applicable to network firewall policies. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	ShortName pulumi.StringPtrInput
}

func (RegionNetworkFirewallPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionNetworkFirewallPolicyArgs)(nil)).Elem()
}

type RegionNetworkFirewallPolicyInput interface {
	pulumi.Input

	ToRegionNetworkFirewallPolicyOutput() RegionNetworkFirewallPolicyOutput
	ToRegionNetworkFirewallPolicyOutputWithContext(ctx context.Context) RegionNetworkFirewallPolicyOutput
}

func (*RegionNetworkFirewallPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionNetworkFirewallPolicy)(nil)).Elem()
}

func (i *RegionNetworkFirewallPolicy) ToRegionNetworkFirewallPolicyOutput() RegionNetworkFirewallPolicyOutput {
	return i.ToRegionNetworkFirewallPolicyOutputWithContext(context.Background())
}

func (i *RegionNetworkFirewallPolicy) ToRegionNetworkFirewallPolicyOutputWithContext(ctx context.Context) RegionNetworkFirewallPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionNetworkFirewallPolicyOutput)
}

type RegionNetworkFirewallPolicyOutput struct{ *pulumi.OutputState }

func (RegionNetworkFirewallPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionNetworkFirewallPolicy)(nil)).Elem()
}

func (o RegionNetworkFirewallPolicyOutput) ToRegionNetworkFirewallPolicyOutput() RegionNetworkFirewallPolicyOutput {
	return o
}

func (o RegionNetworkFirewallPolicyOutput) ToRegionNetworkFirewallPolicyOutputWithContext(ctx context.Context) RegionNetworkFirewallPolicyOutput {
	return o
}

// A list of associations that belong to this firewall policy.
func (o RegionNetworkFirewallPolicyOutput) Associations() FirewallPolicyAssociationResponseArrayOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicy) FirewallPolicyAssociationResponseArrayOutput {
		return v.Associations
	}).(FirewallPolicyAssociationResponseArrayOutput)
}

// Creation timestamp in RFC3339 text format.
func (o RegionNetworkFirewallPolicyOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicy) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o RegionNetworkFirewallPolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Deprecated, please use short name instead. User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This field is not applicable to network firewall policies. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
//
// Deprecated: Deprecated, please use short name instead. User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This field is not applicable to network firewall policies. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o RegionNetworkFirewallPolicyOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicy) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the firewall policy.
func (o RegionNetworkFirewallPolicyOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicy) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// [Output only] Type of the resource. Always compute#firewallPolicyfor firewall policies
func (o RegionNetworkFirewallPolicyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicy) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. For Organization Firewall Policies it's a [Output Only] numeric ID allocated by Google Cloud which uniquely identifies the Organization Firewall Policy.
func (o RegionNetworkFirewallPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The parent of the firewall policy. This field is not applicable to network firewall policies.
func (o RegionNetworkFirewallPolicyOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicy) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

func (o RegionNetworkFirewallPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o RegionNetworkFirewallPolicyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicy) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o RegionNetworkFirewallPolicyOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicy) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
func (o RegionNetworkFirewallPolicyOutput) RuleTupleCount() pulumi.IntOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicy) pulumi.IntOutput { return v.RuleTupleCount }).(pulumi.IntOutput)
}

// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
func (o RegionNetworkFirewallPolicyOutput) Rules() FirewallPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicy) FirewallPolicyRuleResponseArrayOutput { return v.Rules }).(FirewallPolicyRuleResponseArrayOutput)
}

// Server-defined URL for the resource.
func (o RegionNetworkFirewallPolicyOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicy) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o RegionNetworkFirewallPolicyOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicy) pulumi.StringOutput { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This field is not applicable to network firewall policies. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o RegionNetworkFirewallPolicyOutput) ShortName() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicy) pulumi.StringOutput { return v.ShortName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionNetworkFirewallPolicyInput)(nil)).Elem(), &RegionNetworkFirewallPolicy{})
	pulumi.RegisterOutputType(RegionNetworkFirewallPolicyOutput{})
}
