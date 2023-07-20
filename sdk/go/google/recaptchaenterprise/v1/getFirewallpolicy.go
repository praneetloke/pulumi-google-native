// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified firewall policy.
func LookupFirewallpolicy(ctx *pulumi.Context, args *LookupFirewallpolicyArgs, opts ...pulumi.InvokeOption) (*LookupFirewallpolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFirewallpolicyResult
	err := ctx.Invoke("google-native:recaptchaenterprise/v1:getFirewallpolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallpolicyArgs struct {
	FirewallpolicyId string  `pulumi:"firewallpolicyId"`
	Project          *string `pulumi:"project"`
}

type LookupFirewallpolicyResult struct {
	// The actions that the caller should take regarding user access. There should be at most one terminal action. A terminal action is any action that forces a response, such as AllowAction, BlockAction or SubstituteAction. Zero or more non-terminal actions such as SetHeader might be specified. A single policy can contain up to 16 actions.
	Actions []GoogleCloudRecaptchaenterpriseV1FirewallActionResponse `pulumi:"actions"`
	// A CEL (Common Expression Language) conditional expression that specifies if this policy applies to an incoming user request. If this condition evaluates to true and the requested path matched the path pattern, the associated actions should be executed by the caller. The condition string is checked for CEL syntax correctness on creation. For more information, see the [CEL spec](https://github.com/google/cel-spec) and its [language definition](https://github.com/google/cel-spec/blob/master/doc/langdef.md). A condition has a max length of 500 characters.
	Condition string `pulumi:"condition"`
	// A description of what this policy aims to achieve, for convenience purposes. The description can at most include 256 UTF-8 characters.
	Description string `pulumi:"description"`
	// The resource name for the FirewallPolicy in the format "projects/{project}/firewallpolicies/{firewallpolicy}".
	Name string `pulumi:"name"`
	// The path for which this policy applies, specified as a glob pattern. For more information on glob, see the [manual page](https://man7.org/linux/man-pages/man7/glob.7.html). A path has a max length of 200 characters.
	Path string `pulumi:"path"`
}

func LookupFirewallpolicyOutput(ctx *pulumi.Context, args LookupFirewallpolicyOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallpolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallpolicyResult, error) {
			args := v.(LookupFirewallpolicyArgs)
			r, err := LookupFirewallpolicy(ctx, &args, opts...)
			var s LookupFirewallpolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallpolicyResultOutput)
}

type LookupFirewallpolicyOutputArgs struct {
	FirewallpolicyId pulumi.StringInput    `pulumi:"firewallpolicyId"`
	Project          pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupFirewallpolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallpolicyArgs)(nil)).Elem()
}

type LookupFirewallpolicyResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallpolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallpolicyResult)(nil)).Elem()
}

func (o LookupFirewallpolicyResultOutput) ToLookupFirewallpolicyResultOutput() LookupFirewallpolicyResultOutput {
	return o
}

func (o LookupFirewallpolicyResultOutput) ToLookupFirewallpolicyResultOutputWithContext(ctx context.Context) LookupFirewallpolicyResultOutput {
	return o
}

// The actions that the caller should take regarding user access. There should be at most one terminal action. A terminal action is any action that forces a response, such as AllowAction, BlockAction or SubstituteAction. Zero or more non-terminal actions such as SetHeader might be specified. A single policy can contain up to 16 actions.
func (o LookupFirewallpolicyResultOutput) Actions() GoogleCloudRecaptchaenterpriseV1FirewallActionResponseArrayOutput {
	return o.ApplyT(func(v LookupFirewallpolicyResult) []GoogleCloudRecaptchaenterpriseV1FirewallActionResponse {
		return v.Actions
	}).(GoogleCloudRecaptchaenterpriseV1FirewallActionResponseArrayOutput)
}

// A CEL (Common Expression Language) conditional expression that specifies if this policy applies to an incoming user request. If this condition evaluates to true and the requested path matched the path pattern, the associated actions should be executed by the caller. The condition string is checked for CEL syntax correctness on creation. For more information, see the [CEL spec](https://github.com/google/cel-spec) and its [language definition](https://github.com/google/cel-spec/blob/master/doc/langdef.md). A condition has a max length of 500 characters.
func (o LookupFirewallpolicyResultOutput) Condition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallpolicyResult) string { return v.Condition }).(pulumi.StringOutput)
}

// A description of what this policy aims to achieve, for convenience purposes. The description can at most include 256 UTF-8 characters.
func (o LookupFirewallpolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallpolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// The resource name for the FirewallPolicy in the format "projects/{project}/firewallpolicies/{firewallpolicy}".
func (o LookupFirewallpolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallpolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The path for which this policy applies, specified as a glob pattern. For more information on glob, see the [manual page](https://man7.org/linux/man-pages/man7/glob.7.html). A path has a max length of 200 characters.
func (o LookupFirewallpolicyResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallpolicyResult) string { return v.Path }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallpolicyResultOutput{})
}
