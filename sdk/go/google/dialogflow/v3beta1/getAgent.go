// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the specified agent.
func LookupAgent(ctx *pulumi.Context, args *LookupAgentArgs, opts ...pulumi.InvokeOption) (*LookupAgentResult, error) {
	var rv LookupAgentResult
	err := ctx.Invoke("google-native:dialogflow/v3beta1:getAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAgentArgs struct {
	AgentId  string  `pulumi:"agentId"`
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
}

type LookupAgentResult struct {
	// Hierarchical advanced settings for this agent. The settings exposed at the lower level overrides the settings exposed at the higher level.
	AdvancedSettings GoogleCloudDialogflowCxV3beta1AdvancedSettingsResponse `pulumi:"advancedSettings"`
	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted [Web Demo](https://cloud.google.com/dialogflow/docs/integrations/web-demo) integration.
	AvatarUri string `pulumi:"avatarUri"`
	// Immutable. The default language of the agent as a language tag. See [Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language codes. This field cannot be set by the Agents.UpdateAgent method.
	DefaultLanguageCode string `pulumi:"defaultLanguageCode"`
	// The description of the agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description string `pulumi:"description"`
	// The human-readable name of the agent, unique within the location.
	DisplayName string `pulumi:"displayName"`
	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection bool `pulumi:"enableSpellCorrection"`
	// Indicates if stackdriver logging is enabled for the agent. Please use agent.advanced_settings instead.
	EnableStackdriverLogging bool `pulumi:"enableStackdriverLogging"`
	// Indiciates whether the agent is locked for changes. If the agent is locked, modifications to the agent will be rejected except for RestoreAgent.
	Locked bool `pulumi:"locked"`
	// The unique identifier of the agent. Required for the Agents.UpdateAgent method. Agents.CreateAgent populates the name automatically. Format: `projects//locations//agents/`.
	Name string `pulumi:"name"`
	// Name of the SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
	SecuritySettings string `pulumi:"securitySettings"`
	// Speech recognition related settings.
	SpeechToTextSettings GoogleCloudDialogflowCxV3beta1SpeechToTextSettingsResponse `pulumi:"speechToTextSettings"`
	// Immutable. Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: `projects//locations//agents//flows/`.
	StartFlow string `pulumi:"startFlow"`
	// The list of all languages supported by the agent (except for the `default_language_code`).
	SupportedLanguageCodes []string `pulumi:"supportedLanguageCodes"`
	// The time zone of the agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris.
	TimeZone string `pulumi:"timeZone"`
}

func LookupAgentOutput(ctx *pulumi.Context, args LookupAgentOutputArgs, opts ...pulumi.InvokeOption) LookupAgentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAgentResult, error) {
			args := v.(LookupAgentArgs)
			r, err := LookupAgent(ctx, &args, opts...)
			return *r, err
		}).(LookupAgentResultOutput)
}

type LookupAgentOutputArgs struct {
	AgentId  pulumi.StringInput    `pulumi:"agentId"`
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupAgentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentArgs)(nil)).Elem()
}

type LookupAgentResultOutput struct{ *pulumi.OutputState }

func (LookupAgentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentResult)(nil)).Elem()
}

func (o LookupAgentResultOutput) ToLookupAgentResultOutput() LookupAgentResultOutput {
	return o
}

func (o LookupAgentResultOutput) ToLookupAgentResultOutputWithContext(ctx context.Context) LookupAgentResultOutput {
	return o
}

// Hierarchical advanced settings for this agent. The settings exposed at the lower level overrides the settings exposed at the higher level.
func (o LookupAgentResultOutput) AdvancedSettings() GoogleCloudDialogflowCxV3beta1AdvancedSettingsResponseOutput {
	return o.ApplyT(func(v LookupAgentResult) GoogleCloudDialogflowCxV3beta1AdvancedSettingsResponse {
		return v.AdvancedSettings
	}).(GoogleCloudDialogflowCxV3beta1AdvancedSettingsResponseOutput)
}

// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted [Web Demo](https://cloud.google.com/dialogflow/docs/integrations/web-demo) integration.
func (o LookupAgentResultOutput) AvatarUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.AvatarUri }).(pulumi.StringOutput)
}

// Immutable. The default language of the agent as a language tag. See [Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language codes. This field cannot be set by the Agents.UpdateAgent method.
func (o LookupAgentResultOutput) DefaultLanguageCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.DefaultLanguageCode }).(pulumi.StringOutput)
}

// The description of the agent. The maximum length is 500 characters. If exceeded, the request is rejected.
func (o LookupAgentResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.Description }).(pulumi.StringOutput)
}

// The human-readable name of the agent, unique within the location.
func (o LookupAgentResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Indicates if automatic spell correction is enabled in detect intent requests.
func (o LookupAgentResultOutput) EnableSpellCorrection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAgentResult) bool { return v.EnableSpellCorrection }).(pulumi.BoolOutput)
}

// Indicates if stackdriver logging is enabled for the agent. Please use agent.advanced_settings instead.
func (o LookupAgentResultOutput) EnableStackdriverLogging() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAgentResult) bool { return v.EnableStackdriverLogging }).(pulumi.BoolOutput)
}

// Indiciates whether the agent is locked for changes. If the agent is locked, modifications to the agent will be rejected except for RestoreAgent.
func (o LookupAgentResultOutput) Locked() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAgentResult) bool { return v.Locked }).(pulumi.BoolOutput)
}

// The unique identifier of the agent. Required for the Agents.UpdateAgent method. Agents.CreateAgent populates the name automatically. Format: `projects//locations//agents/`.
func (o LookupAgentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Name of the SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
func (o LookupAgentResultOutput) SecuritySettings() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.SecuritySettings }).(pulumi.StringOutput)
}

// Speech recognition related settings.
func (o LookupAgentResultOutput) SpeechToTextSettings() GoogleCloudDialogflowCxV3beta1SpeechToTextSettingsResponseOutput {
	return o.ApplyT(func(v LookupAgentResult) GoogleCloudDialogflowCxV3beta1SpeechToTextSettingsResponse {
		return v.SpeechToTextSettings
	}).(GoogleCloudDialogflowCxV3beta1SpeechToTextSettingsResponseOutput)
}

// Immutable. Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: `projects//locations//agents//flows/`.
func (o LookupAgentResultOutput) StartFlow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.StartFlow }).(pulumi.StringOutput)
}

// The list of all languages supported by the agent (except for the `default_language_code`).
func (o LookupAgentResultOutput) SupportedLanguageCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAgentResult) []string { return v.SupportedLanguageCodes }).(pulumi.StringArrayOutput)
}

// The time zone of the agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris.
func (o LookupAgentResultOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.TimeZone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAgentResultOutput{})
}
