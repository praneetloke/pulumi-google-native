// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the specified intent.
func LookupIntent(ctx *pulumi.Context, args *LookupIntentArgs, opts ...pulumi.InvokeOption) (*LookupIntentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIntentResult
	err := ctx.Invoke("google-native:dialogflow/v3beta1:getIntent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntentArgs struct {
	AgentId      string  `pulumi:"agentId"`
	IntentId     string  `pulumi:"intentId"`
	LanguageCode *string `pulumi:"languageCode"`
	Location     string  `pulumi:"location"`
	Project      *string `pulumi:"project"`
}

type LookupIntentResult struct {
	// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
	Description string `pulumi:"description"`
	// The human-readable name of the intent, unique within the agent.
	DisplayName string `pulumi:"displayName"`
	// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
	IsFallback bool `pulumi:"isFallback"`
	// The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys-contextual" means the intent is a contextual intent.
	Labels map[string]string `pulumi:"labels"`
	// The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
	Name string `pulumi:"name"`
	// The collection of parameters associated with the intent.
	Parameters []GoogleCloudDialogflowCxV3beta1IntentParameterResponse `pulumi:"parameters"`
	// The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority int `pulumi:"priority"`
	// The collection of training phrases the agent is trained on to identify the intent.
	TrainingPhrases []GoogleCloudDialogflowCxV3beta1IntentTrainingPhraseResponse `pulumi:"trainingPhrases"`
}

func LookupIntentOutput(ctx *pulumi.Context, args LookupIntentOutputArgs, opts ...pulumi.InvokeOption) LookupIntentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntentResult, error) {
			args := v.(LookupIntentArgs)
			r, err := LookupIntent(ctx, &args, opts...)
			var s LookupIntentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntentResultOutput)
}

type LookupIntentOutputArgs struct {
	AgentId      pulumi.StringInput    `pulumi:"agentId"`
	IntentId     pulumi.StringInput    `pulumi:"intentId"`
	LanguageCode pulumi.StringPtrInput `pulumi:"languageCode"`
	Location     pulumi.StringInput    `pulumi:"location"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupIntentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntentArgs)(nil)).Elem()
}

type LookupIntentResultOutput struct{ *pulumi.OutputState }

func (LookupIntentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntentResult)(nil)).Elem()
}

func (o LookupIntentResultOutput) ToLookupIntentResultOutput() LookupIntentResultOutput {
	return o
}

func (o LookupIntentResultOutput) ToLookupIntentResultOutputWithContext(ctx context.Context) LookupIntentResultOutput {
	return o
}

// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
func (o LookupIntentResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntentResult) string { return v.Description }).(pulumi.StringOutput)
}

// The human-readable name of the intent, unique within the agent.
func (o LookupIntentResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntentResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
func (o LookupIntentResultOutput) IsFallback() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIntentResult) bool { return v.IsFallback }).(pulumi.BoolOutput)
}

// The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys-contextual" means the intent is a contextual intent.
func (o LookupIntentResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntentResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
func (o LookupIntentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The collection of parameters associated with the intent.
func (o LookupIntentResultOutput) Parameters() GoogleCloudDialogflowCxV3beta1IntentParameterResponseArrayOutput {
	return o.ApplyT(func(v LookupIntentResult) []GoogleCloudDialogflowCxV3beta1IntentParameterResponse {
		return v.Parameters
	}).(GoogleCloudDialogflowCxV3beta1IntentParameterResponseArrayOutput)
}

// The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
func (o LookupIntentResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIntentResult) int { return v.Priority }).(pulumi.IntOutput)
}

// The collection of training phrases the agent is trained on to identify the intent.
func (o LookupIntentResultOutput) TrainingPhrases() GoogleCloudDialogflowCxV3beta1IntentTrainingPhraseResponseArrayOutput {
	return o.ApplyT(func(v LookupIntentResult) []GoogleCloudDialogflowCxV3beta1IntentTrainingPhraseResponse {
		return v.TrainingPhrases
	}).(GoogleCloudDialogflowCxV3beta1IntentTrainingPhraseResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntentResultOutput{})
}
