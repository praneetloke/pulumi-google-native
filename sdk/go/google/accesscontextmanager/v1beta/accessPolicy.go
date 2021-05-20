// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create an `AccessPolicy`. Fails if this organization already has a `AccessPolicy`. The longrunning Operation will have a successful status once the `AccessPolicy` has propagated to long-lasting storage. Syntactic and basic semantic errors will be returned in `metadata` as a BadRequest proto.
type AccessPolicy struct {
	pulumi.CustomResourceState

	// Resource name of the `AccessPolicy`. Format: `accessPolicies/{policy_id}`
	Name pulumi.StringOutput `pulumi:"name"`
	// Required. The parent of this `AccessPolicy` in the Cloud Resource Hierarchy. Currently immutable once created. Format: `organizations/{organization_id}`
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Required. Human readable title. Does not affect behavior.
	Title pulumi.StringOutput `pulumi:"title"`
}

// NewAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewAccessPolicy(ctx *pulumi.Context,
	name string, args *AccessPolicyArgs, opts ...pulumi.ResourceOption) (*AccessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'AccessPolicyId'")
	}
	var resource AccessPolicy
	err := ctx.RegisterResource("google-native:accesscontextmanager/v1beta:AccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPolicy gets an existing AccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPolicyState, opts ...pulumi.ResourceOption) (*AccessPolicy, error) {
	var resource AccessPolicy
	err := ctx.ReadResource("google-native:accesscontextmanager/v1beta:AccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPolicy resources.
type accessPolicyState struct {
	// Resource name of the `AccessPolicy`. Format: `accessPolicies/{policy_id}`
	Name *string `pulumi:"name"`
	// Required. The parent of this `AccessPolicy` in the Cloud Resource Hierarchy. Currently immutable once created. Format: `organizations/{organization_id}`
	Parent *string `pulumi:"parent"`
	// Required. Human readable title. Does not affect behavior.
	Title *string `pulumi:"title"`
}

type AccessPolicyState struct {
	// Resource name of the `AccessPolicy`. Format: `accessPolicies/{policy_id}`
	Name pulumi.StringPtrInput
	// Required. The parent of this `AccessPolicy` in the Cloud Resource Hierarchy. Currently immutable once created. Format: `organizations/{organization_id}`
	Parent pulumi.StringPtrInput
	// Required. Human readable title. Does not affect behavior.
	Title pulumi.StringPtrInput
}

func (AccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyState)(nil)).Elem()
}

type accessPolicyArgs struct {
	AccessPolicyId string `pulumi:"accessPolicyId"`
	// Resource name of the `AccessPolicy`. Format: `accessPolicies/{policy_id}`
	Name *string `pulumi:"name"`
	// Required. The parent of this `AccessPolicy` in the Cloud Resource Hierarchy. Currently immutable once created. Format: `organizations/{organization_id}`
	Parent *string `pulumi:"parent"`
	// Required. Human readable title. Does not affect behavior.
	Title *string `pulumi:"title"`
}

// The set of arguments for constructing a AccessPolicy resource.
type AccessPolicyArgs struct {
	AccessPolicyId pulumi.StringInput
	// Resource name of the `AccessPolicy`. Format: `accessPolicies/{policy_id}`
	Name pulumi.StringPtrInput
	// Required. The parent of this `AccessPolicy` in the Cloud Resource Hierarchy. Currently immutable once created. Format: `organizations/{organization_id}`
	Parent pulumi.StringPtrInput
	// Required. Human readable title. Does not affect behavior.
	Title pulumi.StringPtrInput
}

func (AccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyArgs)(nil)).Elem()
}

type AccessPolicyInput interface {
	pulumi.Input

	ToAccessPolicyOutput() AccessPolicyOutput
	ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput
}

func (*AccessPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicy)(nil))
}

func (i *AccessPolicy) ToAccessPolicyOutput() AccessPolicyOutput {
	return i.ToAccessPolicyOutputWithContext(context.Background())
}

func (i *AccessPolicy) ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyOutput)
}

type AccessPolicyOutput struct {
	*pulumi.OutputState
}

func (AccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicy)(nil))
}

func (o AccessPolicyOutput) ToAccessPolicyOutput() AccessPolicyOutput {
	return o
}

func (o AccessPolicyOutput) ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccessPolicyOutput{})
}
