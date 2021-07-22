// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets an access control policy for a resource. Replaces any existing policy. Supported resources are: - Tag templates - Entry groups Note: This method sets policies only within Data Catalog and can't be used to manage policies in BigQuery, Pub/Sub, Dataproc Metastore, and any external Google Cloud Platform resources synced with the Data Catalog. To call this method, you must have the following Google IAM permissions: - `datacatalog.tagTemplates.setIamPolicy` to set policies on tag templates. - `datacatalog.entryGroups.setIamPolicy` to set policies on entry groups.
type EntryGroupIamPolicy struct {
	pulumi.CustomResourceState

	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewEntryGroupIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewEntryGroupIamPolicy(ctx *pulumi.Context,
	name string, args *EntryGroupIamPolicyArgs, opts ...pulumi.ResourceOption) (*EntryGroupIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EntryGroupId == nil {
		return nil, errors.New("invalid value for required argument 'EntryGroupId'")
	}
	var resource EntryGroupIamPolicy
	err := ctx.RegisterResource("google-native:datacatalog/v1:EntryGroupIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntryGroupIamPolicy gets an existing EntryGroupIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntryGroupIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntryGroupIamPolicyState, opts ...pulumi.ResourceOption) (*EntryGroupIamPolicy, error) {
	var resource EntryGroupIamPolicy
	err := ctx.ReadResource("google-native:datacatalog/v1:EntryGroupIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EntryGroupIamPolicy resources.
type entryGroupIamPolicyState struct {
}

type EntryGroupIamPolicyState struct {
}

func (EntryGroupIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*entryGroupIamPolicyState)(nil)).Elem()
}

type entryGroupIamPolicyArgs struct {
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings     []Binding `pulumi:"bindings"`
	EntryGroupId string    `pulumi:"entryGroupId"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     *string `pulumi:"etag"`
	Location *string `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a EntryGroupIamPolicy resource.
type EntryGroupIamPolicyArgs struct {
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings     BindingArrayInput
	EntryGroupId pulumi.StringInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     pulumi.StringPtrInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (EntryGroupIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entryGroupIamPolicyArgs)(nil)).Elem()
}

type EntryGroupIamPolicyInput interface {
	pulumi.Input

	ToEntryGroupIamPolicyOutput() EntryGroupIamPolicyOutput
	ToEntryGroupIamPolicyOutputWithContext(ctx context.Context) EntryGroupIamPolicyOutput
}

func (*EntryGroupIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*EntryGroupIamPolicy)(nil))
}

func (i *EntryGroupIamPolicy) ToEntryGroupIamPolicyOutput() EntryGroupIamPolicyOutput {
	return i.ToEntryGroupIamPolicyOutputWithContext(context.Background())
}

func (i *EntryGroupIamPolicy) ToEntryGroupIamPolicyOutputWithContext(ctx context.Context) EntryGroupIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryGroupIamPolicyOutput)
}

type EntryGroupIamPolicyOutput struct {
	*pulumi.OutputState
}

func (EntryGroupIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntryGroupIamPolicy)(nil))
}

func (o EntryGroupIamPolicyOutput) ToEntryGroupIamPolicyOutput() EntryGroupIamPolicyOutput {
	return o
}

func (o EntryGroupIamPolicyOutput) ToEntryGroupIamPolicyOutputWithContext(ctx context.Context) EntryGroupIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EntryGroupIamPolicyOutput{})
}
