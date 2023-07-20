// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the IAM Access Control policy currently in effect for the given Cloud Run service. This result does not include any inherited policies.
func LookupServiceIamPolicy(ctx *pulumi.Context, args *LookupServiceIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupServiceIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceIamPolicyResult
	err := ctx.Invoke("google-native:run/v1:getServiceIamPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceIamPolicyArgs struct {
	Location                      string  `pulumi:"location"`
	OptionsRequestedPolicyVersion *int    `pulumi:"optionsRequestedPolicyVersion"`
	Project                       *string `pulumi:"project"`
	ServiceId                     string  `pulumi:"serviceId"`
}

type LookupServiceIamPolicyResult struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfigResponse `pulumi:"auditConfigs"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings []BindingResponse `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag string `pulumi:"etag"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version int `pulumi:"version"`
}

func LookupServiceIamPolicyOutput(ctx *pulumi.Context, args LookupServiceIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupServiceIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceIamPolicyResult, error) {
			args := v.(LookupServiceIamPolicyArgs)
			r, err := LookupServiceIamPolicy(ctx, &args, opts...)
			var s LookupServiceIamPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceIamPolicyResultOutput)
}

type LookupServiceIamPolicyOutputArgs struct {
	Location                      pulumi.StringInput    `pulumi:"location"`
	OptionsRequestedPolicyVersion pulumi.IntPtrInput    `pulumi:"optionsRequestedPolicyVersion"`
	Project                       pulumi.StringPtrInput `pulumi:"project"`
	ServiceId                     pulumi.StringInput    `pulumi:"serviceId"`
}

func (LookupServiceIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceIamPolicyArgs)(nil)).Elem()
}

type LookupServiceIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupServiceIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceIamPolicyResult)(nil)).Elem()
}

func (o LookupServiceIamPolicyResultOutput) ToLookupServiceIamPolicyResultOutput() LookupServiceIamPolicyResultOutput {
	return o
}

func (o LookupServiceIamPolicyResultOutput) ToLookupServiceIamPolicyResultOutputWithContext(ctx context.Context) LookupServiceIamPolicyResultOutput {
	return o
}

// Specifies cloud audit logging configuration for this policy.
func (o LookupServiceIamPolicyResultOutput) AuditConfigs() AuditConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupServiceIamPolicyResult) []AuditConfigResponse { return v.AuditConfigs }).(AuditConfigResponseArrayOutput)
}

// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
func (o LookupServiceIamPolicyResultOutput) Bindings() BindingResponseArrayOutput {
	return o.ApplyT(func(v LookupServiceIamPolicyResult) []BindingResponse { return v.Bindings }).(BindingResponseArrayOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
func (o LookupServiceIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
func (o LookupServiceIamPolicyResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceIamPolicyResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceIamPolicyResultOutput{})
}
