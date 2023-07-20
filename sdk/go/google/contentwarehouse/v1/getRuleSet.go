// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a ruleset. Returns NOT_FOUND if the ruleset does not exist.
func LookupRuleSet(ctx *pulumi.Context, args *LookupRuleSetArgs, opts ...pulumi.InvokeOption) (*LookupRuleSetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRuleSetResult
	err := ctx.Invoke("google-native:contentwarehouse/v1:getRuleSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRuleSetArgs struct {
	Location  string  `pulumi:"location"`
	Project   *string `pulumi:"project"`
	RuleSetId string  `pulumi:"ruleSetId"`
}

type LookupRuleSetResult struct {
	// Short description of the rule-set.
	Description string `pulumi:"description"`
	// The resource name of the rule set. Managed internally. Format: projects/{project_number}/locations/{location}/ruleSet/{rule_set_id}. The name is ignored when creating a rule set.
	Name string `pulumi:"name"`
	// List of rules given by the customer.
	Rules []GoogleCloudContentwarehouseV1RuleResponse `pulumi:"rules"`
	// Source of the rules i.e., customer name.
	Source string `pulumi:"source"`
}

func LookupRuleSetOutput(ctx *pulumi.Context, args LookupRuleSetOutputArgs, opts ...pulumi.InvokeOption) LookupRuleSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRuleSetResult, error) {
			args := v.(LookupRuleSetArgs)
			r, err := LookupRuleSet(ctx, &args, opts...)
			var s LookupRuleSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRuleSetResultOutput)
}

type LookupRuleSetOutputArgs struct {
	Location  pulumi.StringInput    `pulumi:"location"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
	RuleSetId pulumi.StringInput    `pulumi:"ruleSetId"`
}

func (LookupRuleSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleSetArgs)(nil)).Elem()
}

type LookupRuleSetResultOutput struct{ *pulumi.OutputState }

func (LookupRuleSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleSetResult)(nil)).Elem()
}

func (o LookupRuleSetResultOutput) ToLookupRuleSetResultOutput() LookupRuleSetResultOutput {
	return o
}

func (o LookupRuleSetResultOutput) ToLookupRuleSetResultOutputWithContext(ctx context.Context) LookupRuleSetResultOutput {
	return o
}

// Short description of the rule-set.
func (o LookupRuleSetResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleSetResult) string { return v.Description }).(pulumi.StringOutput)
}

// The resource name of the rule set. Managed internally. Format: projects/{project_number}/locations/{location}/ruleSet/{rule_set_id}. The name is ignored when creating a rule set.
func (o LookupRuleSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleSetResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of rules given by the customer.
func (o LookupRuleSetResultOutput) Rules() GoogleCloudContentwarehouseV1RuleResponseArrayOutput {
	return o.ApplyT(func(v LookupRuleSetResult) []GoogleCloudContentwarehouseV1RuleResponse { return v.Rules }).(GoogleCloudContentwarehouseV1RuleResponseArrayOutput)
}

// Source of the rules i.e., customer name.
func (o LookupRuleSetResultOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleSetResult) string { return v.Source }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRuleSetResultOutput{})
}
