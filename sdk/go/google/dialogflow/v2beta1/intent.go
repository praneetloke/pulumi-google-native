// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an intent in the specified agent. Note: You should always train an agent prior to sending it queries. See the [training documentation](https://cloud.google.com/dialogflow/es/docs/training).
// Auto-naming is currently not supported for this resource.
type Intent struct {
	pulumi.CustomResourceState

	// Optional. The name of the action associated with the intent. Note: The action name must not contain whitespaces.
	Action pulumi.StringOutput `pulumi:"action"`
	// Optional. The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED (i.e. default platform).
	DefaultResponsePlatforms pulumi.StringArrayOutput `pulumi:"defaultResponsePlatforms"`
	// The name of this intent.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Optional. Indicates that this intent ends an interaction. Some integrations (e.g., Actions on Google or Dialogflow phone gateway) use this information to close interaction with an end user. Default is false.
	EndInteraction pulumi.BoolOutput `pulumi:"endInteraction"`
	// Optional. The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of the contexts must be present in the active user session for an event to trigger this intent. Event names are limited to 150 characters.
	Events pulumi.StringArrayOutput `pulumi:"events"`
	// Information about all followup intents that have this intent as a direct or indirect parent. We populate this field only in the output.
	FollowupIntentInfo GoogleCloudDialogflowV2beta1IntentFollowupIntentInfoResponseArrayOutput `pulumi:"followupIntentInfo"`
	// Optional. The list of context names required for this intent to be triggered. Formats: - `projects//agent/sessions/-/contexts/` - `projects//locations//agent/sessions/-/contexts/`
	InputContextNames pulumi.StringArrayOutput `pulumi:"inputContextNames"`
	// Optional. Indicates whether this is a fallback intent.
	IsFallback pulumi.BoolOutput `pulumi:"isFallback"`
	// Optional. Indicates that a live agent should be brought in to handle the interaction with the user. In most cases, when you set this flag to true, you would also want to set end_interaction to true as well. Default is false.
	LiveAgentHandoff pulumi.BoolOutput `pulumi:"liveAgentHandoff"`
	// Optional. The collection of rich messages corresponding to the `Response` field in the Dialogflow console.
	Messages GoogleCloudDialogflowV2beta1IntentMessageResponseArrayOutput `pulumi:"messages"`
	// Optional. Indicates whether Machine Learning is disabled for the intent. Note: If `ml_disabled` setting is set to true, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off.
	MlDisabled pulumi.BoolOutput `pulumi:"mlDisabled"`
	// Optional. The unique identifier of this intent. Required for Intents.UpdateIntent and Intents.BatchUpdateIntents methods. Supported formats: - `projects//agent/intents/` - `projects//locations//agent/intents/`
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. The collection of contexts that are activated when the intent is matched. Context messages in this collection should not set the parameters field. Setting the `lifespan_count` to 0 will reset the context when the intent is matched. Format: `projects//agent/sessions/-/contexts/`.
	OutputContexts GoogleCloudDialogflowV2beta1ContextResponseArrayOutput `pulumi:"outputContexts"`
	// Optional. The collection of parameters associated with the intent.
	Parameters GoogleCloudDialogflowV2beta1IntentParameterResponseArrayOutput `pulumi:"parameters"`
	// Optional. The unique identifier of the parent intent in the chain of followup intents. You can set this field when creating an intent, for example with CreateIntent or BatchUpdateIntents, in order to make this intent a followup intent. It identifies the parent followup intent. Format: `projects//agent/intents/`.
	ParentFollowupIntentName pulumi.StringOutput `pulumi:"parentFollowupIntentName"`
	// Optional. The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Optional. Indicates whether to delete all contexts in the current session when this intent is matched.
	ResetContexts pulumi.BoolOutput `pulumi:"resetContexts"`
	// The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup intents chain for this intent. Format: `projects//agent/intents/`.
	RootFollowupIntentName pulumi.StringOutput `pulumi:"rootFollowupIntentName"`
	// Optional. The collection of examples that the agent is trained on.
	TrainingPhrases GoogleCloudDialogflowV2beta1IntentTrainingPhraseResponseArrayOutput `pulumi:"trainingPhrases"`
	// Optional. Indicates whether webhooks are enabled for the intent.
	WebhookState pulumi.StringOutput `pulumi:"webhookState"`
}

// NewIntent registers a new resource with the given unique name, arguments, and options.
func NewIntent(ctx *pulumi.Context,
	name string, args *IntentArgs, opts ...pulumi.ResourceOption) (*Intent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource Intent
	err := ctx.RegisterResource("google-native:dialogflow/v2beta1:Intent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntent gets an existing Intent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntentState, opts ...pulumi.ResourceOption) (*Intent, error) {
	var resource Intent
	err := ctx.ReadResource("google-native:dialogflow/v2beta1:Intent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Intent resources.
type intentState struct {
}

type IntentState struct {
}

func (IntentState) ElementType() reflect.Type {
	return reflect.TypeOf((*intentState)(nil)).Elem()
}

type intentArgs struct {
	// Optional. The name of the action associated with the intent. Note: The action name must not contain whitespaces.
	Action *string `pulumi:"action"`
	// Optional. The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED (i.e. default platform).
	DefaultResponsePlatforms []IntentDefaultResponsePlatformsItem `pulumi:"defaultResponsePlatforms"`
	// The name of this intent.
	DisplayName string `pulumi:"displayName"`
	// Optional. Indicates that this intent ends an interaction. Some integrations (e.g., Actions on Google or Dialogflow phone gateway) use this information to close interaction with an end user. Default is false.
	EndInteraction *bool `pulumi:"endInteraction"`
	// Optional. The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of the contexts must be present in the active user session for an event to trigger this intent. Event names are limited to 150 characters.
	Events []string `pulumi:"events"`
	// Optional. The list of context names required for this intent to be triggered. Formats: - `projects//agent/sessions/-/contexts/` - `projects//locations//agent/sessions/-/contexts/`
	InputContextNames []string `pulumi:"inputContextNames"`
	IntentView        *string  `pulumi:"intentView"`
	// Optional. Indicates whether this is a fallback intent.
	IsFallback   *bool   `pulumi:"isFallback"`
	LanguageCode *string `pulumi:"languageCode"`
	// Optional. Indicates that a live agent should be brought in to handle the interaction with the user. In most cases, when you set this flag to true, you would also want to set end_interaction to true as well. Default is false.
	LiveAgentHandoff *bool   `pulumi:"liveAgentHandoff"`
	Location         *string `pulumi:"location"`
	// Optional. The collection of rich messages corresponding to the `Response` field in the Dialogflow console.
	Messages []GoogleCloudDialogflowV2beta1IntentMessage `pulumi:"messages"`
	// Optional. Indicates whether Machine Learning is disabled for the intent. Note: If `ml_disabled` setting is set to true, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off.
	MlDisabled *bool `pulumi:"mlDisabled"`
	// Optional. The unique identifier of this intent. Required for Intents.UpdateIntent and Intents.BatchUpdateIntents methods. Supported formats: - `projects//agent/intents/` - `projects//locations//agent/intents/`
	Name *string `pulumi:"name"`
	// Optional. The collection of contexts that are activated when the intent is matched. Context messages in this collection should not set the parameters field. Setting the `lifespan_count` to 0 will reset the context when the intent is matched. Format: `projects//agent/sessions/-/contexts/`.
	OutputContexts []GoogleCloudDialogflowV2beta1Context `pulumi:"outputContexts"`
	// Optional. The collection of parameters associated with the intent.
	Parameters []GoogleCloudDialogflowV2beta1IntentParameter `pulumi:"parameters"`
	// Optional. The unique identifier of the parent intent in the chain of followup intents. You can set this field when creating an intent, for example with CreateIntent or BatchUpdateIntents, in order to make this intent a followup intent. It identifies the parent followup intent. Format: `projects//agent/intents/`.
	ParentFollowupIntentName *string `pulumi:"parentFollowupIntentName"`
	// Optional. The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority *int    `pulumi:"priority"`
	Project  *string `pulumi:"project"`
	// Optional. Indicates whether to delete all contexts in the current session when this intent is matched.
	ResetContexts *bool `pulumi:"resetContexts"`
	// Optional. The collection of examples that the agent is trained on.
	TrainingPhrases []GoogleCloudDialogflowV2beta1IntentTrainingPhrase `pulumi:"trainingPhrases"`
	// Optional. Indicates whether webhooks are enabled for the intent.
	WebhookState *IntentWebhookState `pulumi:"webhookState"`
}

// The set of arguments for constructing a Intent resource.
type IntentArgs struct {
	// Optional. The name of the action associated with the intent. Note: The action name must not contain whitespaces.
	Action pulumi.StringPtrInput
	// Optional. The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED (i.e. default platform).
	DefaultResponsePlatforms IntentDefaultResponsePlatformsItemArrayInput
	// The name of this intent.
	DisplayName pulumi.StringInput
	// Optional. Indicates that this intent ends an interaction. Some integrations (e.g., Actions on Google or Dialogflow phone gateway) use this information to close interaction with an end user. Default is false.
	EndInteraction pulumi.BoolPtrInput
	// Optional. The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of the contexts must be present in the active user session for an event to trigger this intent. Event names are limited to 150 characters.
	Events pulumi.StringArrayInput
	// Optional. The list of context names required for this intent to be triggered. Formats: - `projects//agent/sessions/-/contexts/` - `projects//locations//agent/sessions/-/contexts/`
	InputContextNames pulumi.StringArrayInput
	IntentView        pulumi.StringPtrInput
	// Optional. Indicates whether this is a fallback intent.
	IsFallback   pulumi.BoolPtrInput
	LanguageCode pulumi.StringPtrInput
	// Optional. Indicates that a live agent should be brought in to handle the interaction with the user. In most cases, when you set this flag to true, you would also want to set end_interaction to true as well. Default is false.
	LiveAgentHandoff pulumi.BoolPtrInput
	Location         pulumi.StringPtrInput
	// Optional. The collection of rich messages corresponding to the `Response` field in the Dialogflow console.
	Messages GoogleCloudDialogflowV2beta1IntentMessageArrayInput
	// Optional. Indicates whether Machine Learning is disabled for the intent. Note: If `ml_disabled` setting is set to true, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off.
	MlDisabled pulumi.BoolPtrInput
	// Optional. The unique identifier of this intent. Required for Intents.UpdateIntent and Intents.BatchUpdateIntents methods. Supported formats: - `projects//agent/intents/` - `projects//locations//agent/intents/`
	Name pulumi.StringPtrInput
	// Optional. The collection of contexts that are activated when the intent is matched. Context messages in this collection should not set the parameters field. Setting the `lifespan_count` to 0 will reset the context when the intent is matched. Format: `projects//agent/sessions/-/contexts/`.
	OutputContexts GoogleCloudDialogflowV2beta1ContextArrayInput
	// Optional. The collection of parameters associated with the intent.
	Parameters GoogleCloudDialogflowV2beta1IntentParameterArrayInput
	// Optional. The unique identifier of the parent intent in the chain of followup intents. You can set this field when creating an intent, for example with CreateIntent or BatchUpdateIntents, in order to make this intent a followup intent. It identifies the parent followup intent. Format: `projects//agent/intents/`.
	ParentFollowupIntentName pulumi.StringPtrInput
	// Optional. The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority pulumi.IntPtrInput
	Project  pulumi.StringPtrInput
	// Optional. Indicates whether to delete all contexts in the current session when this intent is matched.
	ResetContexts pulumi.BoolPtrInput
	// Optional. The collection of examples that the agent is trained on.
	TrainingPhrases GoogleCloudDialogflowV2beta1IntentTrainingPhraseArrayInput
	// Optional. Indicates whether webhooks are enabled for the intent.
	WebhookState IntentWebhookStatePtrInput
}

func (IntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*intentArgs)(nil)).Elem()
}

type IntentInput interface {
	pulumi.Input

	ToIntentOutput() IntentOutput
	ToIntentOutputWithContext(ctx context.Context) IntentOutput
}

func (*Intent) ElementType() reflect.Type {
	return reflect.TypeOf((*Intent)(nil))
}

func (i *Intent) ToIntentOutput() IntentOutput {
	return i.ToIntentOutputWithContext(context.Background())
}

func (i *Intent) ToIntentOutputWithContext(ctx context.Context) IntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntentOutput)
}

type IntentOutput struct {
	*pulumi.OutputState
}

func (IntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Intent)(nil))
}

func (o IntentOutput) ToIntentOutput() IntentOutput {
	return o
}

func (o IntentOutput) ToIntentOutputWithContext(ctx context.Context) IntentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntentOutput{})
}
