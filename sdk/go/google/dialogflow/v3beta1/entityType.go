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

// Creates an entity type in the specified agent.
type EntityType struct {
	pulumi.CustomResourceState

	AgentId pulumi.StringOutput `pulumi:"agentId"`
	// Indicates whether the entity type can be automatically expanded.
	AutoExpansionMode pulumi.StringOutput `pulumi:"autoExpansionMode"`
	// The human-readable name of the entity type, unique within the agent.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction pulumi.BoolOutput `pulumi:"enableFuzzyExtraction"`
	// The collection of entity entries associated with the entity type.
	Entities GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponseArrayOutput `pulumi:"entities"`
	// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry `giant`(an adjective), you might consider adding `giants`(a noun) as an exclusion. If the kind of entity type is `KIND_MAP`, then the phrases specified by entities and excluded phrases should be mutually exclusive.
	ExcludedPhrases GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseResponseArrayOutput `pulumi:"excludedPhrases"`
	// Indicates the kind of entity type.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The language of the following fields in `entity_type`: * `EntityType.entities.value` * `EntityType.entities.synonyms` * `EntityType.excluded_phrases.value` If not specified, the agent's default language is used. [Many languages](https://cloud.google.com/dialogflow/cx/docs/reference/language) are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrOutput `pulumi:"languageCode"`
	Location     pulumi.StringOutput    `pulumi:"location"`
	// The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType. Format: `projects//locations//agents//entityTypes/`.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name during logging.
	Redact pulumi.BoolOutput `pulumi:"redact"`
}

// NewEntityType registers a new resource with the given unique name, arguments, and options.
func NewEntityType(ctx *pulumi.Context,
	name string, args *EntityTypeArgs, opts ...pulumi.ResourceOption) (*EntityType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentId == nil {
		return nil, errors.New("invalid value for required argument 'AgentId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"agentId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EntityType
	err := ctx.RegisterResource("google-native:dialogflow/v3beta1:EntityType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntityType gets an existing EntityType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntityType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntityTypeState, opts ...pulumi.ResourceOption) (*EntityType, error) {
	var resource EntityType
	err := ctx.ReadResource("google-native:dialogflow/v3beta1:EntityType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EntityType resources.
type entityTypeState struct {
}

type EntityTypeState struct {
}

func (EntityTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*entityTypeState)(nil)).Elem()
}

type entityTypeArgs struct {
	AgentId string `pulumi:"agentId"`
	// Indicates whether the entity type can be automatically expanded.
	AutoExpansionMode *EntityTypeAutoExpansionMode `pulumi:"autoExpansionMode"`
	// The human-readable name of the entity type, unique within the agent.
	DisplayName string `pulumi:"displayName"`
	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction *bool `pulumi:"enableFuzzyExtraction"`
	// The collection of entity entries associated with the entity type.
	Entities []GoogleCloudDialogflowCxV3beta1EntityTypeEntity `pulumi:"entities"`
	// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry `giant`(an adjective), you might consider adding `giants`(a noun) as an exclusion. If the kind of entity type is `KIND_MAP`, then the phrases specified by entities and excluded phrases should be mutually exclusive.
	ExcludedPhrases []GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhrase `pulumi:"excludedPhrases"`
	// Indicates the kind of entity type.
	Kind EntityTypeKind `pulumi:"kind"`
	// The language of the following fields in `entity_type`: * `EntityType.entities.value` * `EntityType.entities.synonyms` * `EntityType.excluded_phrases.value` If not specified, the agent's default language is used. [Many languages](https://cloud.google.com/dialogflow/cx/docs/reference/language) are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode *string `pulumi:"languageCode"`
	Location     *string `pulumi:"location"`
	// The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType. Format: `projects//locations//agents//entityTypes/`.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name during logging.
	Redact *bool `pulumi:"redact"`
}

// The set of arguments for constructing a EntityType resource.
type EntityTypeArgs struct {
	AgentId pulumi.StringInput
	// Indicates whether the entity type can be automatically expanded.
	AutoExpansionMode EntityTypeAutoExpansionModePtrInput
	// The human-readable name of the entity type, unique within the agent.
	DisplayName pulumi.StringInput
	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction pulumi.BoolPtrInput
	// The collection of entity entries associated with the entity type.
	Entities GoogleCloudDialogflowCxV3beta1EntityTypeEntityArrayInput
	// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry `giant`(an adjective), you might consider adding `giants`(a noun) as an exclusion. If the kind of entity type is `KIND_MAP`, then the phrases specified by entities and excluded phrases should be mutually exclusive.
	ExcludedPhrases GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseArrayInput
	// Indicates the kind of entity type.
	Kind EntityTypeKindInput
	// The language of the following fields in `entity_type`: * `EntityType.entities.value` * `EntityType.entities.synonyms` * `EntityType.excluded_phrases.value` If not specified, the agent's default language is used. [Many languages](https://cloud.google.com/dialogflow/cx/docs/reference/language) are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrInput
	Location     pulumi.StringPtrInput
	// The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType. Format: `projects//locations//agents//entityTypes/`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name during logging.
	Redact pulumi.BoolPtrInput
}

func (EntityTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entityTypeArgs)(nil)).Elem()
}

type EntityTypeInput interface {
	pulumi.Input

	ToEntityTypeOutput() EntityTypeOutput
	ToEntityTypeOutputWithContext(ctx context.Context) EntityTypeOutput
}

func (*EntityType) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityType)(nil)).Elem()
}

func (i *EntityType) ToEntityTypeOutput() EntityTypeOutput {
	return i.ToEntityTypeOutputWithContext(context.Background())
}

func (i *EntityType) ToEntityTypeOutputWithContext(ctx context.Context) EntityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityTypeOutput)
}

type EntityTypeOutput struct{ *pulumi.OutputState }

func (EntityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityType)(nil)).Elem()
}

func (o EntityTypeOutput) ToEntityTypeOutput() EntityTypeOutput {
	return o
}

func (o EntityTypeOutput) ToEntityTypeOutputWithContext(ctx context.Context) EntityTypeOutput {
	return o
}

func (o EntityTypeOutput) AgentId() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityType) pulumi.StringOutput { return v.AgentId }).(pulumi.StringOutput)
}

// Indicates whether the entity type can be automatically expanded.
func (o EntityTypeOutput) AutoExpansionMode() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityType) pulumi.StringOutput { return v.AutoExpansionMode }).(pulumi.StringOutput)
}

// The human-readable name of the entity type, unique within the agent.
func (o EntityTypeOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityType) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Enables fuzzy entity extraction during classification.
func (o EntityTypeOutput) EnableFuzzyExtraction() pulumi.BoolOutput {
	return o.ApplyT(func(v *EntityType) pulumi.BoolOutput { return v.EnableFuzzyExtraction }).(pulumi.BoolOutput)
}

// The collection of entity entries associated with the entity type.
func (o EntityTypeOutput) Entities() GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponseArrayOutput {
	return o.ApplyT(func(v *EntityType) GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponseArrayOutput {
		return v.Entities
	}).(GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponseArrayOutput)
}

// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry `giant`(an adjective), you might consider adding `giants`(a noun) as an exclusion. If the kind of entity type is `KIND_MAP`, then the phrases specified by entities and excluded phrases should be mutually exclusive.
func (o EntityTypeOutput) ExcludedPhrases() GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseResponseArrayOutput {
	return o.ApplyT(func(v *EntityType) GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseResponseArrayOutput {
		return v.ExcludedPhrases
	}).(GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseResponseArrayOutput)
}

// Indicates the kind of entity type.
func (o EntityTypeOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityType) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The language of the following fields in `entity_type`: * `EntityType.entities.value` * `EntityType.entities.synonyms` * `EntityType.excluded_phrases.value` If not specified, the agent's default language is used. [Many languages](https://cloud.google.com/dialogflow/cx/docs/reference/language) are supported. Note: languages must be enabled in the agent before they can be used.
func (o EntityTypeOutput) LanguageCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EntityType) pulumi.StringPtrOutput { return v.LanguageCode }).(pulumi.StringPtrOutput)
}

func (o EntityTypeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityType) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType. Format: `projects//locations//agents//entityTypes/`.
func (o EntityTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EntityTypeOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityType) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name during logging.
func (o EntityTypeOutput) Redact() pulumi.BoolOutput {
	return o.ApplyT(func(v *EntityType) pulumi.BoolOutput { return v.Redact }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EntityTypeInput)(nil)).Elem(), &EntityType{})
	pulumi.RegisterOutputType(EntityTypeOutput{})
}
