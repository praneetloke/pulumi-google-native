// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an entity type in the specified agent.
type AgentEntityType struct {
	pulumi.CustomResourceState

	// Optional. Indicates whether the entity type can be automatically expanded.
	AutoExpansionMode pulumi.StringOutput `pulumi:"autoExpansionMode"`
	// Required. The name of the entity type.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Optional. Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction pulumi.BoolOutput `pulumi:"enableFuzzyExtraction"`
	// Optional. The collection of entity entries associated with the entity type.
	Entities GoogleCloudDialogflowV2EntityTypeEntityResponseArrayOutput `pulumi:"entities"`
	// Required. Indicates the kind of entity type.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType and EntityTypes.BatchUpdateEntityTypes methods. Format: `projects//agent/entityTypes/`.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewAgentEntityType registers a new resource with the given unique name, arguments, and options.
func NewAgentEntityType(ctx *pulumi.Context,
	name string, args *AgentEntityTypeArgs, opts ...pulumi.ResourceOption) (*AgentEntityType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EntityTypeId == nil {
		return nil, errors.New("invalid value for required argument 'EntityTypeId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource AgentEntityType
	err := ctx.RegisterResource("google-native:dialogflow/v2:AgentEntityType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgentEntityType gets an existing AgentEntityType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgentEntityType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentEntityTypeState, opts ...pulumi.ResourceOption) (*AgentEntityType, error) {
	var resource AgentEntityType
	err := ctx.ReadResource("google-native:dialogflow/v2:AgentEntityType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AgentEntityType resources.
type agentEntityTypeState struct {
	// Optional. Indicates whether the entity type can be automatically expanded.
	AutoExpansionMode *string `pulumi:"autoExpansionMode"`
	// Required. The name of the entity type.
	DisplayName *string `pulumi:"displayName"`
	// Optional. Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction *bool `pulumi:"enableFuzzyExtraction"`
	// Optional. The collection of entity entries associated with the entity type.
	Entities []GoogleCloudDialogflowV2EntityTypeEntityResponse `pulumi:"entities"`
	// Required. Indicates the kind of entity type.
	Kind *string `pulumi:"kind"`
	// The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType and EntityTypes.BatchUpdateEntityTypes methods. Format: `projects//agent/entityTypes/`.
	Name *string `pulumi:"name"`
}

type AgentEntityTypeState struct {
	// Optional. Indicates whether the entity type can be automatically expanded.
	AutoExpansionMode pulumi.StringPtrInput
	// Required. The name of the entity type.
	DisplayName pulumi.StringPtrInput
	// Optional. Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction pulumi.BoolPtrInput
	// Optional. The collection of entity entries associated with the entity type.
	Entities GoogleCloudDialogflowV2EntityTypeEntityResponseArrayInput
	// Required. Indicates the kind of entity type.
	Kind pulumi.StringPtrInput
	// The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType and EntityTypes.BatchUpdateEntityTypes methods. Format: `projects//agent/entityTypes/`.
	Name pulumi.StringPtrInput
}

func (AgentEntityTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentEntityTypeState)(nil)).Elem()
}

type agentEntityTypeArgs struct {
	// Optional. Indicates whether the entity type can be automatically expanded.
	AutoExpansionMode *string `pulumi:"autoExpansionMode"`
	// Required. The name of the entity type.
	DisplayName *string `pulumi:"displayName"`
	// Optional. Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction *bool `pulumi:"enableFuzzyExtraction"`
	// Optional. The collection of entity entries associated with the entity type.
	Entities     []GoogleCloudDialogflowV2EntityTypeEntity `pulumi:"entities"`
	EntityTypeId string                                    `pulumi:"entityTypeId"`
	// Required. Indicates the kind of entity type.
	Kind         *string `pulumi:"kind"`
	LanguageCode *string `pulumi:"languageCode"`
	Location     string  `pulumi:"location"`
	// The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType and EntityTypes.BatchUpdateEntityTypes methods. Format: `projects//agent/entityTypes/`.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
}

// The set of arguments for constructing a AgentEntityType resource.
type AgentEntityTypeArgs struct {
	// Optional. Indicates whether the entity type can be automatically expanded.
	AutoExpansionMode pulumi.StringPtrInput
	// Required. The name of the entity type.
	DisplayName pulumi.StringPtrInput
	// Optional. Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction pulumi.BoolPtrInput
	// Optional. The collection of entity entries associated with the entity type.
	Entities     GoogleCloudDialogflowV2EntityTypeEntityArrayInput
	EntityTypeId pulumi.StringInput
	// Required. Indicates the kind of entity type.
	Kind         pulumi.StringPtrInput
	LanguageCode pulumi.StringPtrInput
	Location     pulumi.StringInput
	// The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType and EntityTypes.BatchUpdateEntityTypes methods. Format: `projects//agent/entityTypes/`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
}

func (AgentEntityTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentEntityTypeArgs)(nil)).Elem()
}

type AgentEntityTypeInput interface {
	pulumi.Input

	ToAgentEntityTypeOutput() AgentEntityTypeOutput
	ToAgentEntityTypeOutputWithContext(ctx context.Context) AgentEntityTypeOutput
}

func (*AgentEntityType) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentEntityType)(nil))
}

func (i *AgentEntityType) ToAgentEntityTypeOutput() AgentEntityTypeOutput {
	return i.ToAgentEntityTypeOutputWithContext(context.Background())
}

func (i *AgentEntityType) ToAgentEntityTypeOutputWithContext(ctx context.Context) AgentEntityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentEntityTypeOutput)
}

type AgentEntityTypeOutput struct {
	*pulumi.OutputState
}

func (AgentEntityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentEntityType)(nil))
}

func (o AgentEntityTypeOutput) ToAgentEntityTypeOutput() AgentEntityTypeOutput {
	return o
}

func (o AgentEntityTypeOutput) ToAgentEntityTypeOutputWithContext(ctx context.Context) AgentEntityTypeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AgentEntityTypeOutput{})
}
