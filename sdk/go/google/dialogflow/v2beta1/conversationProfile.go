// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
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
	// Language code for the conversation profile. If not specified, the language is en-US. Language at ConversationProfile should be set for all non en-us languages.
	LanguageCode pulumi.StringOutput `pulumi:"languageCode"`
	// Configuration for logging conversation lifecycle events.
	LoggingConfig GoogleCloudDialogflowV2beta1LoggingConfigResponseOutput `pulumi:"loggingConfig"`
	// The unique identifier of this conversation profile. Format: `projects//locations//conversationProfiles/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration for publishing new message events. Event will be sent in format of ConversationEvent
	NewMessageEventNotificationConfig GoogleCloudDialogflowV2beta1NotificationConfigResponseOutput `pulumi:"newMessageEventNotificationConfig"`
	// Configuration for publishing conversation lifecycle events.
	NotificationConfig GoogleCloudDialogflowV2beta1NotificationConfigResponseOutput `pulumi:"notificationConfig"`
	// Settings for speech transcription.
	SttConfig GoogleCloudDialogflowV2beta1SpeechToTextConfigResponseOutput `pulumi:"sttConfig"`
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
	// Language code for the conversation profile. If not specified, the language is en-US. Language at ConversationProfile should be set for all non en-us languages.
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
	// Settings for speech transcription.
	SttConfig *GoogleCloudDialogflowV2beta1SpeechToTextConfig `pulumi:"sttConfig"`
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
	// Language code for the conversation profile. If not specified, the language is en-US. Language at ConversationProfile should be set for all non en-us languages.
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
	// Settings for speech transcription.
	SttConfig GoogleCloudDialogflowV2beta1SpeechToTextConfigPtrInput
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
	return reflect.TypeOf((*ConversationProfile)(nil))
}

func (i *ConversationProfile) ToConversationProfileOutput() ConversationProfileOutput {
	return i.ToConversationProfileOutputWithContext(context.Background())
}

func (i *ConversationProfile) ToConversationProfileOutputWithContext(ctx context.Context) ConversationProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConversationProfileOutput)
}

type ConversationProfileOutput struct {
	*pulumi.OutputState
}

func (ConversationProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConversationProfile)(nil))
}

func (o ConversationProfileOutput) ToConversationProfileOutput() ConversationProfileOutput {
	return o
}

func (o ConversationProfileOutput) ToConversationProfileOutputWithContext(ctx context.Context) ConversationProfileOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConversationProfileOutput{})
}
