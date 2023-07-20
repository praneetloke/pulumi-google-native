// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new TlsInspectionPolicy in a given project and location.
type TlsInspectionPolicy struct {
	pulumi.CustomResourceState

	// A CA pool resource used to issue interception certificates. The CA pool string has a relative resource path following the form "projects/{project}/locations/{location}/caPools/{ca_pool}".
	CaPool pulumi.StringOutput `pulumi:"caPool"`
	// The timestamp when the resource was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Free-text description of the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	Location    pulumi.StringOutput `pulumi:"location"`
	// Name of the resource. Name is of the form projects/{project}/locations/{location}/tlsInspectionPolicies/{tls_inspection_policy} tls_inspection_policy should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Required. Short name of the TlsInspectionPolicy resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "tls_inspection_policy1".
	TlsInspectionPolicyId pulumi.StringOutput `pulumi:"tlsInspectionPolicyId"`
	// The timestamp when the resource was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewTlsInspectionPolicy registers a new resource with the given unique name, arguments, and options.
func NewTlsInspectionPolicy(ctx *pulumi.Context,
	name string, args *TlsInspectionPolicyArgs, opts ...pulumi.ResourceOption) (*TlsInspectionPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaPool == nil {
		return nil, errors.New("invalid value for required argument 'CaPool'")
	}
	if args.TlsInspectionPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'TlsInspectionPolicyId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"tlsInspectionPolicyId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TlsInspectionPolicy
	err := ctx.RegisterResource("google-native:networksecurity/v1beta1:TlsInspectionPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTlsInspectionPolicy gets an existing TlsInspectionPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTlsInspectionPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TlsInspectionPolicyState, opts ...pulumi.ResourceOption) (*TlsInspectionPolicy, error) {
	var resource TlsInspectionPolicy
	err := ctx.ReadResource("google-native:networksecurity/v1beta1:TlsInspectionPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TlsInspectionPolicy resources.
type tlsInspectionPolicyState struct {
}

type TlsInspectionPolicyState struct {
}

func (TlsInspectionPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*tlsInspectionPolicyState)(nil)).Elem()
}

type tlsInspectionPolicyArgs struct {
	// A CA pool resource used to issue interception certificates. The CA pool string has a relative resource path following the form "projects/{project}/locations/{location}/caPools/{ca_pool}".
	CaPool string `pulumi:"caPool"`
	// Optional. Free-text description of the resource.
	Description *string `pulumi:"description"`
	Location    *string `pulumi:"location"`
	// Name of the resource. Name is of the form projects/{project}/locations/{location}/tlsInspectionPolicies/{tls_inspection_policy} tls_inspection_policy should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Required. Short name of the TlsInspectionPolicy resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "tls_inspection_policy1".
	TlsInspectionPolicyId string `pulumi:"tlsInspectionPolicyId"`
}

// The set of arguments for constructing a TlsInspectionPolicy resource.
type TlsInspectionPolicyArgs struct {
	// A CA pool resource used to issue interception certificates. The CA pool string has a relative resource path following the form "projects/{project}/locations/{location}/caPools/{ca_pool}".
	CaPool pulumi.StringInput
	// Optional. Free-text description of the resource.
	Description pulumi.StringPtrInput
	Location    pulumi.StringPtrInput
	// Name of the resource. Name is of the form projects/{project}/locations/{location}/tlsInspectionPolicies/{tls_inspection_policy} tls_inspection_policy should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Required. Short name of the TlsInspectionPolicy resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "tls_inspection_policy1".
	TlsInspectionPolicyId pulumi.StringInput
}

func (TlsInspectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tlsInspectionPolicyArgs)(nil)).Elem()
}

type TlsInspectionPolicyInput interface {
	pulumi.Input

	ToTlsInspectionPolicyOutput() TlsInspectionPolicyOutput
	ToTlsInspectionPolicyOutputWithContext(ctx context.Context) TlsInspectionPolicyOutput
}

func (*TlsInspectionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsInspectionPolicy)(nil)).Elem()
}

func (i *TlsInspectionPolicy) ToTlsInspectionPolicyOutput() TlsInspectionPolicyOutput {
	return i.ToTlsInspectionPolicyOutputWithContext(context.Background())
}

func (i *TlsInspectionPolicy) ToTlsInspectionPolicyOutputWithContext(ctx context.Context) TlsInspectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsInspectionPolicyOutput)
}

type TlsInspectionPolicyOutput struct{ *pulumi.OutputState }

func (TlsInspectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsInspectionPolicy)(nil)).Elem()
}

func (o TlsInspectionPolicyOutput) ToTlsInspectionPolicyOutput() TlsInspectionPolicyOutput {
	return o
}

func (o TlsInspectionPolicyOutput) ToTlsInspectionPolicyOutputWithContext(ctx context.Context) TlsInspectionPolicyOutput {
	return o
}

// A CA pool resource used to issue interception certificates. The CA pool string has a relative resource path following the form "projects/{project}/locations/{location}/caPools/{ca_pool}".
func (o TlsInspectionPolicyOutput) CaPool() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.CaPool }).(pulumi.StringOutput)
}

// The timestamp when the resource was created.
func (o TlsInspectionPolicyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Free-text description of the resource.
func (o TlsInspectionPolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o TlsInspectionPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the resource. Name is of the form projects/{project}/locations/{location}/tlsInspectionPolicies/{tls_inspection_policy} tls_inspection_policy should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
func (o TlsInspectionPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TlsInspectionPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Required. Short name of the TlsInspectionPolicy resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "tls_inspection_policy1".
func (o TlsInspectionPolicyOutput) TlsInspectionPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.TlsInspectionPolicyId }).(pulumi.StringOutput)
}

// The timestamp when the resource was updated.
func (o TlsInspectionPolicyOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TlsInspectionPolicyInput)(nil)).Elem(), &TlsInspectionPolicy{})
	pulumi.RegisterOutputType(TlsInspectionPolicyOutput{})
}
