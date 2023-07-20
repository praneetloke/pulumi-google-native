// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the access control policy for a contentitem resource. A NOT_FOUND error is returned if the resource does not exist. An empty policy is returned if the resource exists but does not have a policy set on it.Caller must have Google IAM dataplex.content.getIamPolicy permission on the resource.
func LookupLakeContentIamPolicy(ctx *pulumi.Context, args *LookupLakeContentIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupLakeContentIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLakeContentIamPolicyResult
	err := ctx.Invoke("google-native:dataplex/v1:getLakeContentIamPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLakeContentIamPolicyArgs struct {
	ContentId                     string  `pulumi:"contentId"`
	LakeId                        string  `pulumi:"lakeId"`
	Location                      string  `pulumi:"location"`
	OptionsRequestedPolicyVersion *int    `pulumi:"optionsRequestedPolicyVersion"`
	Project                       *string `pulumi:"project"`
}

type LookupLakeContentIamPolicyResult struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []GoogleIamV1AuditConfigResponse `pulumi:"auditConfigs"`
	// Associates a list of members, or principals, with a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one principal.The bindings in a Policy can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the bindings grant 50 different roles to user:alice@example.com, and not to any other principal, then you can add another 1,450 principals to the bindings in the Policy.
	Bindings []GoogleIamV1BindingResponse `pulumi:"bindings"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
	Etag string `pulumi:"etag"`
	// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
	Version int `pulumi:"version"`
}

func LookupLakeContentIamPolicyOutput(ctx *pulumi.Context, args LookupLakeContentIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupLakeContentIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLakeContentIamPolicyResult, error) {
			args := v.(LookupLakeContentIamPolicyArgs)
			r, err := LookupLakeContentIamPolicy(ctx, &args, opts...)
			var s LookupLakeContentIamPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLakeContentIamPolicyResultOutput)
}

type LookupLakeContentIamPolicyOutputArgs struct {
	ContentId                     pulumi.StringInput    `pulumi:"contentId"`
	LakeId                        pulumi.StringInput    `pulumi:"lakeId"`
	Location                      pulumi.StringInput    `pulumi:"location"`
	OptionsRequestedPolicyVersion pulumi.IntPtrInput    `pulumi:"optionsRequestedPolicyVersion"`
	Project                       pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupLakeContentIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLakeContentIamPolicyArgs)(nil)).Elem()
}

type LookupLakeContentIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupLakeContentIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLakeContentIamPolicyResult)(nil)).Elem()
}

func (o LookupLakeContentIamPolicyResultOutput) ToLookupLakeContentIamPolicyResultOutput() LookupLakeContentIamPolicyResultOutput {
	return o
}

func (o LookupLakeContentIamPolicyResultOutput) ToLookupLakeContentIamPolicyResultOutputWithContext(ctx context.Context) LookupLakeContentIamPolicyResultOutput {
	return o
}

// Specifies cloud audit logging configuration for this policy.
func (o LookupLakeContentIamPolicyResultOutput) AuditConfigs() GoogleIamV1AuditConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupLakeContentIamPolicyResult) []GoogleIamV1AuditConfigResponse { return v.AuditConfigs }).(GoogleIamV1AuditConfigResponseArrayOutput)
}

// Associates a list of members, or principals, with a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one principal.The bindings in a Policy can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the bindings grant 50 different roles to user:alice@example.com, and not to any other principal, then you can add another 1,450 principals to the bindings in the Policy.
func (o LookupLakeContentIamPolicyResultOutput) Bindings() GoogleIamV1BindingResponseArrayOutput {
	return o.ApplyT(func(v LookupLakeContentIamPolicyResult) []GoogleIamV1BindingResponse { return v.Bindings }).(GoogleIamV1BindingResponseArrayOutput)
}

// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
func (o LookupLakeContentIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLakeContentIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
func (o LookupLakeContentIamPolicyResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLakeContentIamPolicyResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLakeContentIamPolicyResultOutput{})
}
