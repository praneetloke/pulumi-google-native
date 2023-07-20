// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the specified TransitionRouteGroup.
func LookupTransitionRouteGroup(ctx *pulumi.Context, args *LookupTransitionRouteGroupArgs, opts ...pulumi.InvokeOption) (*LookupTransitionRouteGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTransitionRouteGroupResult
	err := ctx.Invoke("google-native:dialogflow/v3beta1:getTransitionRouteGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTransitionRouteGroupArgs struct {
	AgentId                string  `pulumi:"agentId"`
	FlowId                 string  `pulumi:"flowId"`
	LanguageCode           *string `pulumi:"languageCode"`
	Location               string  `pulumi:"location"`
	Project                *string `pulumi:"project"`
	TransitionRouteGroupId string  `pulumi:"transitionRouteGroupId"`
}

type LookupTransitionRouteGroupResult struct {
	// The human-readable name of the transition route group, unique within the flow. The display name can be no longer than 30 characters.
	DisplayName string `pulumi:"displayName"`
	// The unique identifier of the transition route group. TransitionRouteGroups.CreateTransitionRouteGroup populates the name automatically. Format: `projects//locations//agents//flows//transitionRouteGroups/`.
	Name string `pulumi:"name"`
	// Transition routes associated with the TransitionRouteGroup.
	TransitionRoutes []GoogleCloudDialogflowCxV3beta1TransitionRouteResponse `pulumi:"transitionRoutes"`
}

func LookupTransitionRouteGroupOutput(ctx *pulumi.Context, args LookupTransitionRouteGroupOutputArgs, opts ...pulumi.InvokeOption) LookupTransitionRouteGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTransitionRouteGroupResult, error) {
			args := v.(LookupTransitionRouteGroupArgs)
			r, err := LookupTransitionRouteGroup(ctx, &args, opts...)
			var s LookupTransitionRouteGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTransitionRouteGroupResultOutput)
}

type LookupTransitionRouteGroupOutputArgs struct {
	AgentId                pulumi.StringInput    `pulumi:"agentId"`
	FlowId                 pulumi.StringInput    `pulumi:"flowId"`
	LanguageCode           pulumi.StringPtrInput `pulumi:"languageCode"`
	Location               pulumi.StringInput    `pulumi:"location"`
	Project                pulumi.StringPtrInput `pulumi:"project"`
	TransitionRouteGroupId pulumi.StringInput    `pulumi:"transitionRouteGroupId"`
}

func (LookupTransitionRouteGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransitionRouteGroupArgs)(nil)).Elem()
}

type LookupTransitionRouteGroupResultOutput struct{ *pulumi.OutputState }

func (LookupTransitionRouteGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransitionRouteGroupResult)(nil)).Elem()
}

func (o LookupTransitionRouteGroupResultOutput) ToLookupTransitionRouteGroupResultOutput() LookupTransitionRouteGroupResultOutput {
	return o
}

func (o LookupTransitionRouteGroupResultOutput) ToLookupTransitionRouteGroupResultOutputWithContext(ctx context.Context) LookupTransitionRouteGroupResultOutput {
	return o
}

// The human-readable name of the transition route group, unique within the flow. The display name can be no longer than 30 characters.
func (o LookupTransitionRouteGroupResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransitionRouteGroupResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The unique identifier of the transition route group. TransitionRouteGroups.CreateTransitionRouteGroup populates the name automatically. Format: `projects//locations//agents//flows//transitionRouteGroups/`.
func (o LookupTransitionRouteGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransitionRouteGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Transition routes associated with the TransitionRouteGroup.
func (o LookupTransitionRouteGroupResultOutput) TransitionRoutes() GoogleCloudDialogflowCxV3beta1TransitionRouteResponseArrayOutput {
	return o.ApplyT(func(v LookupTransitionRouteGroupResult) []GoogleCloudDialogflowCxV3beta1TransitionRouteResponse {
		return v.TransitionRoutes
	}).(GoogleCloudDialogflowCxV3beta1TransitionRouteResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTransitionRouteGroupResultOutput{})
}
