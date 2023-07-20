// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Fetches the representation of an existing Response Policy Rule.
func GetResponsePolicyRule(ctx *pulumi.Context, args *GetResponsePolicyRuleArgs, opts ...pulumi.InvokeOption) (*GetResponsePolicyRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetResponsePolicyRuleResult
	err := ctx.Invoke("google-native:dns/v1:getResponsePolicyRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetResponsePolicyRuleArgs struct {
	ClientOperationId  *string `pulumi:"clientOperationId"`
	Project            *string `pulumi:"project"`
	ResponsePolicy     string  `pulumi:"responsePolicy"`
	ResponsePolicyRule string  `pulumi:"responsePolicyRule"`
}

type GetResponsePolicyRuleResult struct {
	// Answer this query with a behavior rather than DNS data.
	Behavior string `pulumi:"behavior"`
	// The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
	DnsName string `pulumi:"dnsName"`
	Kind    string `pulumi:"kind"`
	// Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name; in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
	LocalData ResponsePolicyRuleLocalDataResponse `pulumi:"localData"`
	// An identifier for this rule. Must be unique with the ResponsePolicy.
	RuleName string `pulumi:"ruleName"`
}

func GetResponsePolicyRuleOutput(ctx *pulumi.Context, args GetResponsePolicyRuleOutputArgs, opts ...pulumi.InvokeOption) GetResponsePolicyRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetResponsePolicyRuleResult, error) {
			args := v.(GetResponsePolicyRuleArgs)
			r, err := GetResponsePolicyRule(ctx, &args, opts...)
			var s GetResponsePolicyRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetResponsePolicyRuleResultOutput)
}

type GetResponsePolicyRuleOutputArgs struct {
	ClientOperationId  pulumi.StringPtrInput `pulumi:"clientOperationId"`
	Project            pulumi.StringPtrInput `pulumi:"project"`
	ResponsePolicy     pulumi.StringInput    `pulumi:"responsePolicy"`
	ResponsePolicyRule pulumi.StringInput    `pulumi:"responsePolicyRule"`
}

func (GetResponsePolicyRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResponsePolicyRuleArgs)(nil)).Elem()
}

type GetResponsePolicyRuleResultOutput struct{ *pulumi.OutputState }

func (GetResponsePolicyRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResponsePolicyRuleResult)(nil)).Elem()
}

func (o GetResponsePolicyRuleResultOutput) ToGetResponsePolicyRuleResultOutput() GetResponsePolicyRuleResultOutput {
	return o
}

func (o GetResponsePolicyRuleResultOutput) ToGetResponsePolicyRuleResultOutputWithContext(ctx context.Context) GetResponsePolicyRuleResultOutput {
	return o
}

// Answer this query with a behavior rather than DNS data.
func (o GetResponsePolicyRuleResultOutput) Behavior() pulumi.StringOutput {
	return o.ApplyT(func(v GetResponsePolicyRuleResult) string { return v.Behavior }).(pulumi.StringOutput)
}

// The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
func (o GetResponsePolicyRuleResultOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v GetResponsePolicyRuleResult) string { return v.DnsName }).(pulumi.StringOutput)
}

func (o GetResponsePolicyRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v GetResponsePolicyRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name; in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
func (o GetResponsePolicyRuleResultOutput) LocalData() ResponsePolicyRuleLocalDataResponseOutput {
	return o.ApplyT(func(v GetResponsePolicyRuleResult) ResponsePolicyRuleLocalDataResponse { return v.LocalData }).(ResponsePolicyRuleLocalDataResponseOutput)
}

// An identifier for this rule. Must be unique with the ResponsePolicy.
func (o GetResponsePolicyRuleResultOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v GetResponsePolicyRuleResult) string { return v.RuleName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetResponsePolicyRuleResultOutput{})
}
