// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an TransitionRouteGroup in the specified flow. Note: You should always train a flow prior to sending it queries. See the [training documentation](https://cloud.google.com/dialogflow/cx/docs/concept/training).
type TransitionRouteGroup struct {
	pulumi.CustomResourceState

	// The human-readable name of the transition route group, unique within the Agent. The display name can be no longer than 30 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The unique identifier of the transition route group. TransitionRouteGroups.CreateTransitionRouteGroup populates the name automatically. Format: `projects//locations//agents//flows//transitionRouteGroups/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Transition routes associated with the TransitionRouteGroup.
	TransitionRoutes GoogleCloudDialogflowCxV3TransitionRouteResponseArrayOutput `pulumi:"transitionRoutes"`
}

// NewTransitionRouteGroup registers a new resource with the given unique name, arguments, and options.
func NewTransitionRouteGroup(ctx *pulumi.Context,
	name string, args *TransitionRouteGroupArgs, opts ...pulumi.ResourceOption) (*TransitionRouteGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentId == nil {
		return nil, errors.New("invalid value for required argument 'AgentId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.FlowId == nil {
		return nil, errors.New("invalid value for required argument 'FlowId'")
	}
	var resource TransitionRouteGroup
	err := ctx.RegisterResource("google-native:dialogflow/v3:TransitionRouteGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitionRouteGroup gets an existing TransitionRouteGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitionRouteGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitionRouteGroupState, opts ...pulumi.ResourceOption) (*TransitionRouteGroup, error) {
	var resource TransitionRouteGroup
	err := ctx.ReadResource("google-native:dialogflow/v3:TransitionRouteGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitionRouteGroup resources.
type transitionRouteGroupState struct {
}

type TransitionRouteGroupState struct {
}

func (TransitionRouteGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitionRouteGroupState)(nil)).Elem()
}

type transitionRouteGroupArgs struct {
	AgentId string `pulumi:"agentId"`
	// The human-readable name of the transition route group, unique within the Agent. The display name can be no longer than 30 characters.
	DisplayName  string  `pulumi:"displayName"`
	FlowId       string  `pulumi:"flowId"`
	LanguageCode *string `pulumi:"languageCode"`
	Location     *string `pulumi:"location"`
	// The unique identifier of the transition route group. TransitionRouteGroups.CreateTransitionRouteGroup populates the name automatically. Format: `projects//locations//agents//flows//transitionRouteGroups/`.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Transition routes associated with the TransitionRouteGroup.
	TransitionRoutes []GoogleCloudDialogflowCxV3TransitionRoute `pulumi:"transitionRoutes"`
}

// The set of arguments for constructing a TransitionRouteGroup resource.
type TransitionRouteGroupArgs struct {
	AgentId pulumi.StringInput
	// The human-readable name of the transition route group, unique within the Agent. The display name can be no longer than 30 characters.
	DisplayName  pulumi.StringInput
	FlowId       pulumi.StringInput
	LanguageCode pulumi.StringPtrInput
	Location     pulumi.StringPtrInput
	// The unique identifier of the transition route group. TransitionRouteGroups.CreateTransitionRouteGroup populates the name automatically. Format: `projects//locations//agents//flows//transitionRouteGroups/`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Transition routes associated with the TransitionRouteGroup.
	TransitionRoutes GoogleCloudDialogflowCxV3TransitionRouteArrayInput
}

func (TransitionRouteGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitionRouteGroupArgs)(nil)).Elem()
}

type TransitionRouteGroupInput interface {
	pulumi.Input

	ToTransitionRouteGroupOutput() TransitionRouteGroupOutput
	ToTransitionRouteGroupOutputWithContext(ctx context.Context) TransitionRouteGroupOutput
}

func (*TransitionRouteGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*TransitionRouteGroup)(nil))
}

func (i *TransitionRouteGroup) ToTransitionRouteGroupOutput() TransitionRouteGroupOutput {
	return i.ToTransitionRouteGroupOutputWithContext(context.Background())
}

func (i *TransitionRouteGroup) ToTransitionRouteGroupOutputWithContext(ctx context.Context) TransitionRouteGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitionRouteGroupOutput)
}

type TransitionRouteGroupOutput struct {
	*pulumi.OutputState
}

func (TransitionRouteGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransitionRouteGroup)(nil))
}

func (o TransitionRouteGroupOutput) ToTransitionRouteGroupOutput() TransitionRouteGroupOutput {
	return o
}

func (o TransitionRouteGroupOutput) ToTransitionRouteGroupOutputWithContext(ctx context.Context) TransitionRouteGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TransitionRouteGroupOutput{})
}
