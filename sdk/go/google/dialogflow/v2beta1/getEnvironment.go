// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the specified agent environment.
func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentResult
	err := ctx.Invoke("google-native:dialogflow/v2beta1:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
	EnvironmentId string  `pulumi:"environmentId"`
	Location      string  `pulumi:"location"`
	Project       *string `pulumi:"project"`
}

type LookupEnvironmentResult struct {
	// Optional. The agent version loaded into this environment. Supported formats: - `projects//agent/versions/` - `projects//locations//agent/versions/`
	AgentVersion string `pulumi:"agentVersion"`
	// Optional. The developer-provided description for this environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description string `pulumi:"description"`
	// Optional. The fulfillment settings to use for this environment.
	Fulfillment GoogleCloudDialogflowV2beta1FulfillmentResponse `pulumi:"fulfillment"`
	// The unique identifier of this agent environment. Supported formats: - `projects//agent/environments/` - `projects//locations//agent/environments/`
	Name string `pulumi:"name"`
	// The state of this environment. This field is read-only, i.e., it cannot be set by create and update methods.
	State string `pulumi:"state"`
	// Optional. Text to speech settings for this environment.
	TextToSpeechSettings GoogleCloudDialogflowV2beta1TextToSpeechSettingsResponse `pulumi:"textToSpeechSettings"`
	// The last update time of this environment. This field is read-only, i.e., it cannot be set by create and update methods.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupEnvironmentOutput(ctx *pulumi.Context, args LookupEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentResult, error) {
			args := v.(LookupEnvironmentArgs)
			r, err := LookupEnvironment(ctx, &args, opts...)
			var s LookupEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentResultOutput)
}

type LookupEnvironmentOutputArgs struct {
	EnvironmentId pulumi.StringInput    `pulumi:"environmentId"`
	Location      pulumi.StringInput    `pulumi:"location"`
	Project       pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentArgs)(nil)).Elem()
}

type LookupEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentResult)(nil)).Elem()
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutput() LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutputWithContext(ctx context.Context) LookupEnvironmentResultOutput {
	return o
}

// Optional. The agent version loaded into this environment. Supported formats: - `projects//agent/versions/` - `projects//locations//agent/versions/`
func (o LookupEnvironmentResultOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.AgentVersion }).(pulumi.StringOutput)
}

// Optional. The developer-provided description for this environment. The maximum length is 500 characters. If exceeded, the request is rejected.
func (o LookupEnvironmentResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Description }).(pulumi.StringOutput)
}

// Optional. The fulfillment settings to use for this environment.
func (o LookupEnvironmentResultOutput) Fulfillment() GoogleCloudDialogflowV2beta1FulfillmentResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) GoogleCloudDialogflowV2beta1FulfillmentResponse { return v.Fulfillment }).(GoogleCloudDialogflowV2beta1FulfillmentResponseOutput)
}

// The unique identifier of this agent environment. Supported formats: - `projects//agent/environments/` - `projects//locations//agent/environments/`
func (o LookupEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The state of this environment. This field is read-only, i.e., it cannot be set by create and update methods.
func (o LookupEnvironmentResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.State }).(pulumi.StringOutput)
}

// Optional. Text to speech settings for this environment.
func (o LookupEnvironmentResultOutput) TextToSpeechSettings() GoogleCloudDialogflowV2beta1TextToSpeechSettingsResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) GoogleCloudDialogflowV2beta1TextToSpeechSettingsResponse {
		return v.TextToSpeechSettings
	}).(GoogleCloudDialogflowV2beta1TextToSpeechSettingsResponseOutput)
}

// The last update time of this environment. This field is read-only, i.e., it cannot be set by create and update methods.
func (o LookupEnvironmentResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentResultOutput{})
}
