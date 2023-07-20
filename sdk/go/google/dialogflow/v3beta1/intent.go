// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an intent in the specified agent. Note: You should always train a flow prior to sending it queries. See the [training documentation](https://cloud.google.com/dialogflow/cx/docs/concept/training).
type Intent struct {
	pulumi.CustomResourceState

	AgentId pulumi.StringOutput `pulumi:"agentId"`
	// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// The human-readable name of the intent, unique within the agent.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
	IsFallback pulumi.BoolOutput `pulumi:"isFallback"`
	// The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys-contextual" means the intent is a contextual intent.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The language of the following fields in `intent`: * `Intent.training_phrases.parts.text` If not specified, the agent's default language is used. [Many languages](https://cloud.google.com/dialogflow/cx/docs/reference/language) are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrOutput `pulumi:"languageCode"`
	Location     pulumi.StringOutput    `pulumi:"location"`
	// The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The collection of parameters associated with the intent.
	Parameters GoogleCloudDialogflowCxV3beta1IntentParameterResponseArrayOutput `pulumi:"parameters"`
	// The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority pulumi.IntOutput    `pulumi:"priority"`
	Project  pulumi.StringOutput `pulumi:"project"`
	// The collection of training phrases the agent is trained on to identify the intent.
	TrainingPhrases GoogleCloudDialogflowCxV3beta1IntentTrainingPhraseResponseArrayOutput `pulumi:"trainingPhrases"`
}

// NewIntent registers a new resource with the given unique name, arguments, and options.
func NewIntent(ctx *pulumi.Context,
	name string, args *IntentArgs, opts ...pulumi.ResourceOption) (*Intent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentId == nil {
		return nil, errors.New("invalid value for required argument 'AgentId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"agentId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Intent
	err := ctx.RegisterResource("google-native:dialogflow/v3beta1:Intent", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-native:dialogflow/v3beta1:Intent", name, id, state, &resource, opts...)
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
	AgentId string `pulumi:"agentId"`
	// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
	Description *string `pulumi:"description"`
	// The human-readable name of the intent, unique within the agent.
	DisplayName string `pulumi:"displayName"`
	// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
	IsFallback *bool `pulumi:"isFallback"`
	// The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys-contextual" means the intent is a contextual intent.
	Labels map[string]string `pulumi:"labels"`
	// The language of the following fields in `intent`: * `Intent.training_phrases.parts.text` If not specified, the agent's default language is used. [Many languages](https://cloud.google.com/dialogflow/cx/docs/reference/language) are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode *string `pulumi:"languageCode"`
	Location     *string `pulumi:"location"`
	// The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
	Name *string `pulumi:"name"`
	// The collection of parameters associated with the intent.
	Parameters []GoogleCloudDialogflowCxV3beta1IntentParameter `pulumi:"parameters"`
	// The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority *int    `pulumi:"priority"`
	Project  *string `pulumi:"project"`
	// The collection of training phrases the agent is trained on to identify the intent.
	TrainingPhrases []GoogleCloudDialogflowCxV3beta1IntentTrainingPhrase `pulumi:"trainingPhrases"`
}

// The set of arguments for constructing a Intent resource.
type IntentArgs struct {
	AgentId pulumi.StringInput
	// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
	Description pulumi.StringPtrInput
	// The human-readable name of the intent, unique within the agent.
	DisplayName pulumi.StringInput
	// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
	IsFallback pulumi.BoolPtrInput
	// The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys-contextual" means the intent is a contextual intent.
	Labels pulumi.StringMapInput
	// The language of the following fields in `intent`: * `Intent.training_phrases.parts.text` If not specified, the agent's default language is used. [Many languages](https://cloud.google.com/dialogflow/cx/docs/reference/language) are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrInput
	Location     pulumi.StringPtrInput
	// The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
	Name pulumi.StringPtrInput
	// The collection of parameters associated with the intent.
	Parameters GoogleCloudDialogflowCxV3beta1IntentParameterArrayInput
	// The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority pulumi.IntPtrInput
	Project  pulumi.StringPtrInput
	// The collection of training phrases the agent is trained on to identify the intent.
	TrainingPhrases GoogleCloudDialogflowCxV3beta1IntentTrainingPhraseArrayInput
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
	return reflect.TypeOf((**Intent)(nil)).Elem()
}

func (i *Intent) ToIntentOutput() IntentOutput {
	return i.ToIntentOutputWithContext(context.Background())
}

func (i *Intent) ToIntentOutputWithContext(ctx context.Context) IntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntentOutput)
}

type IntentOutput struct{ *pulumi.OutputState }

func (IntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Intent)(nil)).Elem()
}

func (o IntentOutput) ToIntentOutput() IntentOutput {
	return o
}

func (o IntentOutput) ToIntentOutputWithContext(ctx context.Context) IntentOutput {
	return o
}

func (o IntentOutput) AgentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringOutput { return v.AgentId }).(pulumi.StringOutput)
}

// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
func (o IntentOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The human-readable name of the intent, unique within the agent.
func (o IntentOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
func (o IntentOutput) IsFallback() pulumi.BoolOutput {
	return o.ApplyT(func(v *Intent) pulumi.BoolOutput { return v.IsFallback }).(pulumi.BoolOutput)
}

// The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys-contextual" means the intent is a contextual intent.
func (o IntentOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The language of the following fields in `intent`: * `Intent.training_phrases.parts.text` If not specified, the agent's default language is used. [Many languages](https://cloud.google.com/dialogflow/cx/docs/reference/language) are supported. Note: languages must be enabled in the agent before they can be used.
func (o IntentOutput) LanguageCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringPtrOutput { return v.LanguageCode }).(pulumi.StringPtrOutput)
}

func (o IntentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
func (o IntentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The collection of parameters associated with the intent.
func (o IntentOutput) Parameters() GoogleCloudDialogflowCxV3beta1IntentParameterResponseArrayOutput {
	return o.ApplyT(func(v *Intent) GoogleCloudDialogflowCxV3beta1IntentParameterResponseArrayOutput { return v.Parameters }).(GoogleCloudDialogflowCxV3beta1IntentParameterResponseArrayOutput)
}

// The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
func (o IntentOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *Intent) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

func (o IntentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The collection of training phrases the agent is trained on to identify the intent.
func (o IntentOutput) TrainingPhrases() GoogleCloudDialogflowCxV3beta1IntentTrainingPhraseResponseArrayOutput {
	return o.ApplyT(func(v *Intent) GoogleCloudDialogflowCxV3beta1IntentTrainingPhraseResponseArrayOutput {
		return v.TrainingPhrases
	}).(GoogleCloudDialogflowCxV3beta1IntentTrainingPhraseResponseArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntentInput)(nil)).Elem(), &Intent{})
	pulumi.RegisterOutputType(IntentOutput{})
}
