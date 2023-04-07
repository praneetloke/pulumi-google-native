// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a conversation profile in the specified project. ConversationProfile.CreateTime and ConversationProfile.UpdateTime aren't populated in the response. You can retrieve them via GetConversationProfile API.
type ConversationProfile struct {
	pulumi.CustomResourceState

	// Configuration for an automated agent to use with this profile.
	AutomatedAgentConfig GoogleCloudDialogflowV2beta1AutomatedAgentConfigResponseOutput `pulumi:"automatedAgentConfig"`
	// Create time of the conversation profile.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Human readable name for this profile. Max length 1024 bytes.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Configuration for agent assistance to use with this profile.
	HumanAgentAssistantConfig GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigResponseOutput `pulumi:"humanAgentAssistantConfig"`
	// Configuration for connecting to a live agent. Currently, this feature is not general available, please contact Google to get access.
	HumanAgentHandoffConfig GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigResponseOutput `pulumi:"humanAgentHandoffConfig"`
	// Language code for the conversation profile. If not specified, the language is en-US. Language at ConversationProfile should be set for all non en-us languages. This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag. Example: "en-US".
	LanguageCode pulumi.StringOutput `pulumi:"languageCode"`
	Location     pulumi.StringOutput `pulumi:"location"`
	// Configuration for logging conversation lifecycle events.
	LoggingConfig GoogleCloudDialogflowV2beta1LoggingConfigResponseOutput `pulumi:"loggingConfig"`
	// The unique identifier of this conversation profile. Format: `projects//locations//conversationProfiles/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration for publishing new message events. Event will be sent in format of ConversationEvent
	NewMessageEventNotificationConfig GoogleCloudDialogflowV2beta1NotificationConfigResponseOutput `pulumi:"newMessageEventNotificationConfig"`
	// Configuration for publishing conversation lifecycle events.
	NotificationConfig GoogleCloudDialogflowV2beta1NotificationConfigResponseOutput `pulumi:"notificationConfig"`
	Project            pulumi.StringOutput                                          `pulumi:"project"`
	// Name of the CX SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
	SecuritySettings pulumi.StringOutput `pulumi:"securitySettings"`
	// Settings for speech transcription.
	SttConfig GoogleCloudDialogflowV2beta1SpeechToTextConfigResponseOutput `pulumi:"sttConfig"`
	// The time zone of this conversational profile from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris. Defaults to America/New_York.
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
	// Configuration for Text-to-Speech synthesization. Used by Phone Gateway to specify synthesization options. If agent defines synthesization options as well, agent settings overrides the option here.
	TtsConfig GoogleCloudDialogflowV2beta1SynthesizeSpeechConfigResponseOutput `pulumi:"ttsConfig"`
	// Update time of the conversation profile.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewConversationProfile registers a new resource with the given unique name, arguments, and options.
func NewConversationProfile(ctx *pulumi.Context,
	name string, args *ConversationProfileArgs, opts ...pulumi.ResourceOption) (*ConversationProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource ConversationProfile
	err := ctx.RegisterResource("google-native:dialogflow/v2beta1:ConversationProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConversationProfile gets an existing ConversationProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConversationProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConversationProfileState, opts ...pulumi.ResourceOption) (*ConversationProfile, error) {
	var resource ConversationProfile
	err := ctx.ReadResource("google-native:dialogflow/v2beta1:ConversationProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConversationProfile resources.
type conversationProfileState struct {
}

type ConversationProfileState struct {
}

func (ConversationProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*conversationProfileState)(nil)).Elem()
}

type conversationProfileArgs struct {
	// Configuration for an automated agent to use with this profile.
	AutomatedAgentConfig *GoogleCloudDialogflowV2beta1AutomatedAgentConfig `pulumi:"automatedAgentConfig"`
	// Human readable name for this profile. Max length 1024 bytes.
	DisplayName string `pulumi:"displayName"`
	// Configuration for agent assistance to use with this profile.
	HumanAgentAssistantConfig *GoogleCloudDialogflowV2beta1HumanAgentAssistantConfig `pulumi:"humanAgentAssistantConfig"`
	// Configuration for connecting to a live agent. Currently, this feature is not general available, please contact Google to get access.
	HumanAgentHandoffConfig *GoogleCloudDialogflowV2beta1HumanAgentHandoffConfig `pulumi:"humanAgentHandoffConfig"`
	// Language code for the conversation profile. If not specified, the language is en-US. Language at ConversationProfile should be set for all non en-us languages. This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag. Example: "en-US".
	LanguageCode *string `pulumi:"languageCode"`
	Location     *string `pulumi:"location"`
	// Configuration for logging conversation lifecycle events.
	LoggingConfig *GoogleCloudDialogflowV2beta1LoggingConfig `pulumi:"loggingConfig"`
	// The unique identifier of this conversation profile. Format: `projects//locations//conversationProfiles/`.
	Name *string `pulumi:"name"`
	// Configuration for publishing new message events. Event will be sent in format of ConversationEvent
	NewMessageEventNotificationConfig *GoogleCloudDialogflowV2beta1NotificationConfig `pulumi:"newMessageEventNotificationConfig"`
	// Configuration for publishing conversation lifecycle events.
	NotificationConfig *GoogleCloudDialogflowV2beta1NotificationConfig `pulumi:"notificationConfig"`
	Project            *string                                         `pulumi:"project"`
	// Name of the CX SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
	SecuritySettings *string `pulumi:"securitySettings"`
	// Settings for speech transcription.
	SttConfig *GoogleCloudDialogflowV2beta1SpeechToTextConfig `pulumi:"sttConfig"`
	// The time zone of this conversational profile from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris. Defaults to America/New_York.
	TimeZone *string `pulumi:"timeZone"`
	// Configuration for Text-to-Speech synthesization. Used by Phone Gateway to specify synthesization options. If agent defines synthesization options as well, agent settings overrides the option here.
	TtsConfig *GoogleCloudDialogflowV2beta1SynthesizeSpeechConfig `pulumi:"ttsConfig"`
}

// The set of arguments for constructing a ConversationProfile resource.
type ConversationProfileArgs struct {
	// Configuration for an automated agent to use with this profile.
	AutomatedAgentConfig GoogleCloudDialogflowV2beta1AutomatedAgentConfigPtrInput
	// Human readable name for this profile. Max length 1024 bytes.
	DisplayName pulumi.StringInput
	// Configuration for agent assistance to use with this profile.
	HumanAgentAssistantConfig GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigPtrInput
	// Configuration for connecting to a live agent. Currently, this feature is not general available, please contact Google to get access.
	HumanAgentHandoffConfig GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigPtrInput
	// Language code for the conversation profile. If not specified, the language is en-US. Language at ConversationProfile should be set for all non en-us languages. This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag. Example: "en-US".
	LanguageCode pulumi.StringPtrInput
	Location     pulumi.StringPtrInput
	// Configuration for logging conversation lifecycle events.
	LoggingConfig GoogleCloudDialogflowV2beta1LoggingConfigPtrInput
	// The unique identifier of this conversation profile. Format: `projects//locations//conversationProfiles/`.
	Name pulumi.StringPtrInput
	// Configuration for publishing new message events. Event will be sent in format of ConversationEvent
	NewMessageEventNotificationConfig GoogleCloudDialogflowV2beta1NotificationConfigPtrInput
	// Configuration for publishing conversation lifecycle events.
	NotificationConfig GoogleCloudDialogflowV2beta1NotificationConfigPtrInput
	Project            pulumi.StringPtrInput
	// Name of the CX SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
	SecuritySettings pulumi.StringPtrInput
	// Settings for speech transcription.
	SttConfig GoogleCloudDialogflowV2beta1SpeechToTextConfigPtrInput
	// The time zone of this conversational profile from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris. Defaults to America/New_York.
	TimeZone pulumi.StringPtrInput
	// Configuration for Text-to-Speech synthesization. Used by Phone Gateway to specify synthesization options. If agent defines synthesization options as well, agent settings overrides the option here.
	TtsConfig GoogleCloudDialogflowV2beta1SynthesizeSpeechConfigPtrInput
}

func (ConversationProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*conversationProfileArgs)(nil)).Elem()
}

type ConversationProfileInput interface {
	pulumi.Input

	ToConversationProfileOutput() ConversationProfileOutput
	ToConversationProfileOutputWithContext(ctx context.Context) ConversationProfileOutput
}

func (*ConversationProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**ConversationProfile)(nil)).Elem()
}

func (i *ConversationProfile) ToConversationProfileOutput() ConversationProfileOutput {
	return i.ToConversationProfileOutputWithContext(context.Background())
}

func (i *ConversationProfile) ToConversationProfileOutputWithContext(ctx context.Context) ConversationProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConversationProfileOutput)
}

type ConversationProfileOutput struct{ *pulumi.OutputState }

func (ConversationProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConversationProfile)(nil)).Elem()
}

func (o ConversationProfileOutput) ToConversationProfileOutput() ConversationProfileOutput {
	return o
}

func (o ConversationProfileOutput) ToConversationProfileOutputWithContext(ctx context.Context) ConversationProfileOutput {
	return o
}

// Configuration for an automated agent to use with this profile.
func (o ConversationProfileOutput) AutomatedAgentConfig() GoogleCloudDialogflowV2beta1AutomatedAgentConfigResponseOutput {
	return o.ApplyT(func(v *ConversationProfile) GoogleCloudDialogflowV2beta1AutomatedAgentConfigResponseOutput {
		return v.AutomatedAgentConfig
	}).(GoogleCloudDialogflowV2beta1AutomatedAgentConfigResponseOutput)
}

// Create time of the conversation profile.
func (o ConversationProfileOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ConversationProfile) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Human readable name for this profile. Max length 1024 bytes.
func (o ConversationProfileOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConversationProfile) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Configuration for agent assistance to use with this profile.
func (o ConversationProfileOutput) HumanAgentAssistantConfig() GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigResponseOutput {
	return o.ApplyT(func(v *ConversationProfile) GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigResponseOutput {
		return v.HumanAgentAssistantConfig
	}).(GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigResponseOutput)
}

// Configuration for connecting to a live agent. Currently, this feature is not general available, please contact Google to get access.
func (o ConversationProfileOutput) HumanAgentHandoffConfig() GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigResponseOutput {
	return o.ApplyT(func(v *ConversationProfile) GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigResponseOutput {
		return v.HumanAgentHandoffConfig
	}).(GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigResponseOutput)
}

// Language code for the conversation profile. If not specified, the language is en-US. Language at ConversationProfile should be set for all non en-us languages. This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag. Example: "en-US".
func (o ConversationProfileOutput) LanguageCode() pulumi.StringOutput {
	return o.ApplyT(func(v *ConversationProfile) pulumi.StringOutput { return v.LanguageCode }).(pulumi.StringOutput)
}

func (o ConversationProfileOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ConversationProfile) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Configuration for logging conversation lifecycle events.
func (o ConversationProfileOutput) LoggingConfig() GoogleCloudDialogflowV2beta1LoggingConfigResponseOutput {
	return o.ApplyT(func(v *ConversationProfile) GoogleCloudDialogflowV2beta1LoggingConfigResponseOutput {
		return v.LoggingConfig
	}).(GoogleCloudDialogflowV2beta1LoggingConfigResponseOutput)
}

// The unique identifier of this conversation profile. Format: `projects//locations//conversationProfiles/`.
func (o ConversationProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConversationProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Configuration for publishing new message events. Event will be sent in format of ConversationEvent
func (o ConversationProfileOutput) NewMessageEventNotificationConfig() GoogleCloudDialogflowV2beta1NotificationConfigResponseOutput {
	return o.ApplyT(func(v *ConversationProfile) GoogleCloudDialogflowV2beta1NotificationConfigResponseOutput {
		return v.NewMessageEventNotificationConfig
	}).(GoogleCloudDialogflowV2beta1NotificationConfigResponseOutput)
}

// Configuration for publishing conversation lifecycle events.
func (o ConversationProfileOutput) NotificationConfig() GoogleCloudDialogflowV2beta1NotificationConfigResponseOutput {
	return o.ApplyT(func(v *ConversationProfile) GoogleCloudDialogflowV2beta1NotificationConfigResponseOutput {
		return v.NotificationConfig
	}).(GoogleCloudDialogflowV2beta1NotificationConfigResponseOutput)
}

func (o ConversationProfileOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ConversationProfile) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Name of the CX SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
func (o ConversationProfileOutput) SecuritySettings() pulumi.StringOutput {
	return o.ApplyT(func(v *ConversationProfile) pulumi.StringOutput { return v.SecuritySettings }).(pulumi.StringOutput)
}

// Settings for speech transcription.
func (o ConversationProfileOutput) SttConfig() GoogleCloudDialogflowV2beta1SpeechToTextConfigResponseOutput {
	return o.ApplyT(func(v *ConversationProfile) GoogleCloudDialogflowV2beta1SpeechToTextConfigResponseOutput {
		return v.SttConfig
	}).(GoogleCloudDialogflowV2beta1SpeechToTextConfigResponseOutput)
}

// The time zone of this conversational profile from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris. Defaults to America/New_York.
func (o ConversationProfileOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *ConversationProfile) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

// Configuration for Text-to-Speech synthesization. Used by Phone Gateway to specify synthesization options. If agent defines synthesization options as well, agent settings overrides the option here.
func (o ConversationProfileOutput) TtsConfig() GoogleCloudDialogflowV2beta1SynthesizeSpeechConfigResponseOutput {
	return o.ApplyT(func(v *ConversationProfile) GoogleCloudDialogflowV2beta1SynthesizeSpeechConfigResponseOutput {
		return v.TtsConfig
	}).(GoogleCloudDialogflowV2beta1SynthesizeSpeechConfigResponseOutput)
}

// Update time of the conversation profile.
func (o ConversationProfileOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ConversationProfile) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConversationProfileInput)(nil)).Elem(), &ConversationProfile{})
	pulumi.RegisterOutputType(ConversationProfileOutput{})
}
