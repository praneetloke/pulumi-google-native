// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a session entity type.
type SessionEntityType struct {
	pulumi.CustomResourceState

	AgentId pulumi.StringOutput `pulumi:"agentId"`
	// The collection of entities to override or supplement the custom entity type.
	Entities GoogleCloudDialogflowCxV3EntityTypeEntityResponseArrayOutput `pulumi:"entities"`
	// Indicates whether the additional data should override or supplement the custom entity type definition.
	EntityOverrideMode pulumi.StringOutput `pulumi:"entityOverrideMode"`
	EnvironmentId      pulumi.StringOutput `pulumi:"environmentId"`
	Location           pulumi.StringOutput `pulumi:"location"`
	// The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
	Name      pulumi.StringOutput `pulumi:"name"`
	Project   pulumi.StringOutput `pulumi:"project"`
	SessionId pulumi.StringOutput `pulumi:"sessionId"`
}

// NewSessionEntityType registers a new resource with the given unique name, arguments, and options.
func NewSessionEntityType(ctx *pulumi.Context,
	name string, args *SessionEntityTypeArgs, opts ...pulumi.ResourceOption) (*SessionEntityType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentId == nil {
		return nil, errors.New("invalid value for required argument 'AgentId'")
	}
	if args.Entities == nil {
		return nil, errors.New("invalid value for required argument 'Entities'")
	}
	if args.EntityOverrideMode == nil {
		return nil, errors.New("invalid value for required argument 'EntityOverrideMode'")
	}
	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.SessionId == nil {
		return nil, errors.New("invalid value for required argument 'SessionId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"agentId",
		"environmentId",
		"location",
		"project",
		"sessionId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SessionEntityType
	err := ctx.RegisterResource("google-native:dialogflow/v3:SessionEntityType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSessionEntityType gets an existing SessionEntityType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSessionEntityType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SessionEntityTypeState, opts ...pulumi.ResourceOption) (*SessionEntityType, error) {
	var resource SessionEntityType
	err := ctx.ReadResource("google-native:dialogflow/v3:SessionEntityType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SessionEntityType resources.
type sessionEntityTypeState struct {
}

type SessionEntityTypeState struct {
}

func (SessionEntityTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*sessionEntityTypeState)(nil)).Elem()
}

type sessionEntityTypeArgs struct {
	AgentId string `pulumi:"agentId"`
	// The collection of entities to override or supplement the custom entity type.
	Entities []GoogleCloudDialogflowCxV3EntityTypeEntity `pulumi:"entities"`
	// Indicates whether the additional data should override or supplement the custom entity type definition.
	EntityOverrideMode SessionEntityTypeEntityOverrideMode `pulumi:"entityOverrideMode"`
	EnvironmentId      string                              `pulumi:"environmentId"`
	Location           *string                             `pulumi:"location"`
	// The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
	Name      *string `pulumi:"name"`
	Project   *string `pulumi:"project"`
	SessionId string  `pulumi:"sessionId"`
}

// The set of arguments for constructing a SessionEntityType resource.
type SessionEntityTypeArgs struct {
	AgentId pulumi.StringInput
	// The collection of entities to override or supplement the custom entity type.
	Entities GoogleCloudDialogflowCxV3EntityTypeEntityArrayInput
	// Indicates whether the additional data should override or supplement the custom entity type definition.
	EntityOverrideMode SessionEntityTypeEntityOverrideModeInput
	EnvironmentId      pulumi.StringInput
	Location           pulumi.StringPtrInput
	// The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
	Name      pulumi.StringPtrInput
	Project   pulumi.StringPtrInput
	SessionId pulumi.StringInput
}

func (SessionEntityTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sessionEntityTypeArgs)(nil)).Elem()
}

type SessionEntityTypeInput interface {
	pulumi.Input

	ToSessionEntityTypeOutput() SessionEntityTypeOutput
	ToSessionEntityTypeOutputWithContext(ctx context.Context) SessionEntityTypeOutput
}

func (*SessionEntityType) ElementType() reflect.Type {
	return reflect.TypeOf((**SessionEntityType)(nil)).Elem()
}

func (i *SessionEntityType) ToSessionEntityTypeOutput() SessionEntityTypeOutput {
	return i.ToSessionEntityTypeOutputWithContext(context.Background())
}

func (i *SessionEntityType) ToSessionEntityTypeOutputWithContext(ctx context.Context) SessionEntityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SessionEntityTypeOutput)
}

type SessionEntityTypeOutput struct{ *pulumi.OutputState }

func (SessionEntityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SessionEntityType)(nil)).Elem()
}

func (o SessionEntityTypeOutput) ToSessionEntityTypeOutput() SessionEntityTypeOutput {
	return o
}

func (o SessionEntityTypeOutput) ToSessionEntityTypeOutputWithContext(ctx context.Context) SessionEntityTypeOutput {
	return o
}

func (o SessionEntityTypeOutput) AgentId() pulumi.StringOutput {
	return o.ApplyT(func(v *SessionEntityType) pulumi.StringOutput { return v.AgentId }).(pulumi.StringOutput)
}

// The collection of entities to override or supplement the custom entity type.
func (o SessionEntityTypeOutput) Entities() GoogleCloudDialogflowCxV3EntityTypeEntityResponseArrayOutput {
	return o.ApplyT(func(v *SessionEntityType) GoogleCloudDialogflowCxV3EntityTypeEntityResponseArrayOutput {
		return v.Entities
	}).(GoogleCloudDialogflowCxV3EntityTypeEntityResponseArrayOutput)
}

// Indicates whether the additional data should override or supplement the custom entity type definition.
func (o SessionEntityTypeOutput) EntityOverrideMode() pulumi.StringOutput {
	return o.ApplyT(func(v *SessionEntityType) pulumi.StringOutput { return v.EntityOverrideMode }).(pulumi.StringOutput)
}

func (o SessionEntityTypeOutput) EnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *SessionEntityType) pulumi.StringOutput { return v.EnvironmentId }).(pulumi.StringOutput)
}

func (o SessionEntityTypeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SessionEntityType) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
func (o SessionEntityTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SessionEntityType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SessionEntityTypeOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SessionEntityType) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o SessionEntityTypeOutput) SessionId() pulumi.StringOutput {
	return o.ApplyT(func(v *SessionEntityType) pulumi.StringOutput { return v.SessionId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SessionEntityTypeInput)(nil)).Elem(), &SessionEntityType{})
	pulumi.RegisterOutputType(SessionEntityTypeOutput{})
}
