// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the access control policy for a resource. Returns an empty policy if the resource exists and does not have a policy set.
func LookupReusableConfigIamPolicy(ctx *pulumi.Context, args *LookupReusableConfigIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupReusableConfigIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupReusableConfigIamPolicyResult
	err := ctx.Invoke("google-native:privateca/v1beta1:getReusableConfigIamPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReusableConfigIamPolicyArgs struct {
	Location                      string  `pulumi:"location"`
	OptionsRequestedPolicyVersion *int    `pulumi:"optionsRequestedPolicyVersion"`
	Project                       *string `pulumi:"project"`
	ReusableConfigId              string  `pulumi:"reusableConfigId"`
}

type LookupReusableConfigIamPolicyResult struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfigResponse `pulumi:"auditConfigs"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings []BindingResponse `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag string `pulumi:"etag"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version int `pulumi:"version"`
}

func LookupReusableConfigIamPolicyOutput(ctx *pulumi.Context, args LookupReusableConfigIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupReusableConfigIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReusableConfigIamPolicyResult, error) {
			args := v.(LookupReusableConfigIamPolicyArgs)
			r, err := LookupReusableConfigIamPolicy(ctx, &args, opts...)
			var s LookupReusableConfigIamPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReusableConfigIamPolicyResultOutput)
}

type LookupReusableConfigIamPolicyOutputArgs struct {
	Location                      pulumi.StringInput    `pulumi:"location"`
	OptionsRequestedPolicyVersion pulumi.IntPtrInput    `pulumi:"optionsRequestedPolicyVersion"`
	Project                       pulumi.StringPtrInput `pulumi:"project"`
	ReusableConfigId              pulumi.StringInput    `pulumi:"reusableConfigId"`
}

func (LookupReusableConfigIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReusableConfigIamPolicyArgs)(nil)).Elem()
}

type LookupReusableConfigIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupReusableConfigIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReusableConfigIamPolicyResult)(nil)).Elem()
}

func (o LookupReusableConfigIamPolicyResultOutput) ToLookupReusableConfigIamPolicyResultOutput() LookupReusableConfigIamPolicyResultOutput {
	return o
}

func (o LookupReusableConfigIamPolicyResultOutput) ToLookupReusableConfigIamPolicyResultOutputWithContext(ctx context.Context) LookupReusableConfigIamPolicyResultOutput {
	return o
}

// Specifies cloud audit logging configuration for this policy.
func (o LookupReusableConfigIamPolicyResultOutput) AuditConfigs() AuditConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupReusableConfigIamPolicyResult) []AuditConfigResponse { return v.AuditConfigs }).(AuditConfigResponseArrayOutput)
}

// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
func (o LookupReusableConfigIamPolicyResultOutput) Bindings() BindingResponseArrayOutput {
	return o.ApplyT(func(v LookupReusableConfigIamPolicyResult) []BindingResponse { return v.Bindings }).(BindingResponseArrayOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
func (o LookupReusableConfigIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReusableConfigIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
func (o LookupReusableConfigIamPolicyResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v LookupReusableConfigIamPolicyResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReusableConfigIamPolicyResultOutput{})
}
