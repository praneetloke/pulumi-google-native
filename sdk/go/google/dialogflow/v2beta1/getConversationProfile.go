// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the specified conversation profile.
func LookupConversationProfile(ctx *pulumi.Context, args *LookupConversationProfileArgs, opts ...pulumi.InvokeOption) (*LookupConversationProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupConversationProfileResult
	err := ctx.Invoke("google-native:dialogflow/v2beta1:getConversationProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConversationProfileArgs struct {
	ConversationProfileId string  `pulumi:"conversationProfileId"`
	Location              string  `pulumi:"location"`
	Project               *string `pulumi:"project"`
}

type LookupConversationProfileResult struct {
	// Configuration for an automated agent to use with this profile.
	AutomatedAgentConfig GoogleCloudDialogflowV2beta1AutomatedAgentConfigResponse `pulumi:"automatedAgentConfig"`
	// Create time of the conversation profile.
	CreateTime string `pulumi:"createTime"`
	// Human readable name for this profile. Max length 1024 bytes.
	DisplayName string `pulumi:"displayName"`
	// Configuration for agent assistance to use with this profile.
	HumanAgentAssistantConfig GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigResponse `pulumi:"humanAgentAssistantConfig"`
	// Configuration for connecting to a live agent. Currently, this feature is not general available, please contact Google to get access.
	HumanAgentHandoffConfig GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigResponse `pulumi:"humanAgentHandoffConfig"`
	// Language code for the conversation profile. If not specified, the language is en-US. Language at ConversationProfile should be set for all non en-us languages. This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag. Example: "en-US".
	LanguageCode string `pulumi:"languageCode"`
	// Configuration for logging conversation lifecycle events.
	LoggingConfig GoogleCloudDialogflowV2beta1LoggingConfigResponse `pulumi:"loggingConfig"`
	// The unique identifier of this conversation profile. Format: `projects//locations//conversationProfiles/`.
	Name string `pulumi:"name"`
	// Configuration for publishing new message events. Event will be sent in format of ConversationEvent
	NewMessageEventNotificationConfig GoogleCloudDialogflowV2beta1NotificationConfigResponse `pulumi:"newMessageEventNotificationConfig"`
	// Configuration for publishing conversation lifecycle events.
	NotificationConfig GoogleCloudDialogflowV2beta1NotificationConfigResponse `pulumi:"notificationConfig"`
	// Name of the CX SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
	SecuritySettings string `pulumi:"securitySettings"`
	// Settings for speech transcription.
	SttConfig GoogleCloudDialogflowV2beta1SpeechToTextConfigResponse `pulumi:"sttConfig"`
	// The time zone of this conversational profile from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris. Defaults to America/New_York.
	TimeZone string `pulumi:"timeZone"`
	// Configuration for Text-to-Speech synthesization. Used by Phone Gateway to specify synthesization options. If agent defines synthesization options as well, agent settings overrides the option here.
	TtsConfig GoogleCloudDialogflowV2beta1SynthesizeSpeechConfigResponse `pulumi:"ttsConfig"`
	// Update time of the conversation profile.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupConversationProfileOutput(ctx *pulumi.Context, args LookupConversationProfileOutputArgs, opts ...pulumi.InvokeOption) LookupConversationProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConversationProfileResult, error) {
			args := v.(LookupConversationProfileArgs)
			r, err := LookupConversationProfile(ctx, &args, opts...)
			var s LookupConversationProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConversationProfileResultOutput)
}

type LookupConversationProfileOutputArgs struct {
	ConversationProfileId pulumi.StringInput    `pulumi:"conversationProfileId"`
	Location              pulumi.StringInput    `pulumi:"location"`
	Project               pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupConversationProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConversationProfileArgs)(nil)).Elem()
}

type LookupConversationProfileResultOutput struct{ *pulumi.OutputState }

func (LookupConversationProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConversationProfileResult)(nil)).Elem()
}

func (o LookupConversationProfileResultOutput) ToLookupConversationProfileResultOutput() LookupConversationProfileResultOutput {
	return o
}

func (o LookupConversationProfileResultOutput) ToLookupConversationProfileResultOutputWithContext(ctx context.Context) LookupConversationProfileResultOutput {
	return o
}

// Configuration for an automated agent to use with this profile.
func (o LookupConversationProfileResultOutput) AutomatedAgentConfig() GoogleCloudDialogflowV2beta1AutomatedAgentConfigResponseOutput {
	return o.ApplyT(func(v LookupConversationProfileResult) GoogleCloudDialogflowV2beta1AutomatedAgentConfigResponse {
		return v.AutomatedAgentConfig
	}).(GoogleCloudDialogflowV2beta1AutomatedAgentConfigResponseOutput)
}

// Create time of the conversation profile.
func (o LookupConversationProfileResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationProfileResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Human readable name for this profile. Max length 1024 bytes.
func (o LookupConversationProfileResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationProfileResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Configuration for agent assistance to use with this profile.
func (o LookupConversationProfileResultOutput) HumanAgentAssistantConfig() GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigResponseOutput {
	return o.ApplyT(func(v LookupConversationProfileResult) GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigResponse {
		return v.HumanAgentAssistantConfig
	}).(GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigResponseOutput)
}

// Configuration for connecting to a live agent. Currently, this feature is not general available, please contact Google to get access.
func (o LookupConversationProfileResultOutput) HumanAgentHandoffConfig() GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigResponseOutput {
	return o.ApplyT(func(v LookupConversationProfileResult) GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigResponse {
		return v.HumanAgentHandoffConfig
	}).(GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigResponseOutput)
}

// Language code for the conversation profile. If not specified, the language is en-US. Language at ConversationProfile should be set for all non en-us languages. This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag. Example: "en-US".
func (o LookupConversationProfileResultOutput) LanguageCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationProfileResult) string { return v.LanguageCode }).(pulumi.StringOutput)
}

// Configuration for logging conversation lifecycle events.
func (o LookupConversationProfileResultOutput) LoggingConfig() GoogleCloudDialogflowV2beta1LoggingConfigResponseOutput {
	return o.ApplyT(func(v LookupConversationProfileResult) GoogleCloudDialogflowV2beta1LoggingConfigResponse {
		return v.LoggingConfig
	}).(GoogleCloudDialogflowV2beta1LoggingConfigResponseOutput)
}

// The unique identifier of this conversation profile. Format: `projects//locations//conversationProfiles/`.
func (o LookupConversationProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// Configuration for publishing new message events. Event will be sent in format of ConversationEvent
func (o LookupConversationProfileResultOutput) NewMessageEventNotificationConfig() GoogleCloudDialogflowV2beta1NotificationConfigResponseOutput {
	return o.ApplyT(func(v LookupConversationProfileResult) GoogleCloudDialogflowV2beta1NotificationConfigResponse {
		return v.NewMessageEventNotificationConfig
	}).(GoogleCloudDialogflowV2beta1NotificationConfigResponseOutput)
}

// Configuration for publishing conversation lifecycle events.
func (o LookupConversationProfileResultOutput) NotificationConfig() GoogleCloudDialogflowV2beta1NotificationConfigResponseOutput {
	return o.ApplyT(func(v LookupConversationProfileResult) GoogleCloudDialogflowV2beta1NotificationConfigResponse {
		return v.NotificationConfig
	}).(GoogleCloudDialogflowV2beta1NotificationConfigResponseOutput)
}

// Name of the CX SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
func (o LookupConversationProfileResultOutput) SecuritySettings() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationProfileResult) string { return v.SecuritySettings }).(pulumi.StringOutput)
}

// Settings for speech transcription.
func (o LookupConversationProfileResultOutput) SttConfig() GoogleCloudDialogflowV2beta1SpeechToTextConfigResponseOutput {
	return o.ApplyT(func(v LookupConversationProfileResult) GoogleCloudDialogflowV2beta1SpeechToTextConfigResponse {
		return v.SttConfig
	}).(GoogleCloudDialogflowV2beta1SpeechToTextConfigResponseOutput)
}

// The time zone of this conversational profile from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris. Defaults to America/New_York.
func (o LookupConversationProfileResultOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationProfileResult) string { return v.TimeZone }).(pulumi.StringOutput)
}

// Configuration for Text-to-Speech synthesization. Used by Phone Gateway to specify synthesization options. If agent defines synthesization options as well, agent settings overrides the option here.
func (o LookupConversationProfileResultOutput) TtsConfig() GoogleCloudDialogflowV2beta1SynthesizeSpeechConfigResponseOutput {
	return o.ApplyT(func(v LookupConversationProfileResult) GoogleCloudDialogflowV2beta1SynthesizeSpeechConfigResponse {
		return v.TtsConfig
	}).(GoogleCloudDialogflowV2beta1SynthesizeSpeechConfigResponseOutput)
}

// Update time of the conversation profile.
func (o LookupConversationProfileResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationProfileResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConversationProfileResultOutput{})
}
