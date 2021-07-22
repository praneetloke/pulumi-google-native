// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a flow in the specified agent. Note: You should always train a flow prior to sending it queries. See the [training documentation](https://cloud.google.com/dialogflow/cx/docs/concept/training).
type Flow struct {
	pulumi.CustomResourceState

	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringOutput `pulumi:"description"`
	// The human-readable name of the flow.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
	EventHandlers GoogleCloudDialogflowCxV3beta1EventHandlerResponseArrayOutput `pulumi:"eventHandlers"`
	// The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// NLU related settings of the flow.
	NluSettings GoogleCloudDialogflowCxV3beta1NluSettingsResponseOutput `pulumi:"nluSettings"`
	// A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
	TransitionRouteGroups pulumi.StringArrayOutput `pulumi:"transitionRouteGroups"`
	// A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
	TransitionRoutes GoogleCloudDialogflowCxV3beta1TransitionRouteResponseArrayOutput `pulumi:"transitionRoutes"`
}

// NewFlow registers a new resource with the given unique name, arguments, and options.
func NewFlow(ctx *pulumi.Context,
	name string, args *FlowArgs, opts ...pulumi.ResourceOption) (*Flow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentId == nil {
		return nil, errors.New("invalid value for required argument 'AgentId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource Flow
	err := ctx.RegisterResource("google-native:dialogflow/v3beta1:Flow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlow gets an existing Flow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowState, opts ...pulumi.ResourceOption) (*Flow, error) {
	var resource Flow
	err := ctx.ReadResource("google-native:dialogflow/v3beta1:Flow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Flow resources.
type flowState struct {
}

type FlowState struct {
}

func (FlowState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowState)(nil)).Elem()
}

type flowArgs struct {
	AgentId string `pulumi:"agentId"`
	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `pulumi:"description"`
	// The human-readable name of the flow.
	DisplayName string `pulumi:"displayName"`
	// A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
	EventHandlers []GoogleCloudDialogflowCxV3beta1EventHandler `pulumi:"eventHandlers"`
	LanguageCode  *string                                      `pulumi:"languageCode"`
	Location      *string                                      `pulumi:"location"`
	// The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
	Name *string `pulumi:"name"`
	// NLU related settings of the flow.
	NluSettings *GoogleCloudDialogflowCxV3beta1NluSettings `pulumi:"nluSettings"`
	Project     *string                                    `pulumi:"project"`
	// A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
	TransitionRouteGroups []string `pulumi:"transitionRouteGroups"`
	// A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
	TransitionRoutes []GoogleCloudDialogflowCxV3beta1TransitionRoute `pulumi:"transitionRoutes"`
}

// The set of arguments for constructing a Flow resource.
type FlowArgs struct {
	AgentId pulumi.StringInput
	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringPtrInput
	// The human-readable name of the flow.
	DisplayName pulumi.StringInput
	// A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
	EventHandlers GoogleCloudDialogflowCxV3beta1EventHandlerArrayInput
	LanguageCode  pulumi.StringPtrInput
	Location      pulumi.StringPtrInput
	// The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
	Name pulumi.StringPtrInput
	// NLU related settings of the flow.
	NluSettings GoogleCloudDialogflowCxV3beta1NluSettingsPtrInput
	Project     pulumi.StringPtrInput
	// A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
	TransitionRouteGroups pulumi.StringArrayInput
	// A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
	TransitionRoutes GoogleCloudDialogflowCxV3beta1TransitionRouteArrayInput
}

func (FlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowArgs)(nil)).Elem()
}

type FlowInput interface {
	pulumi.Input

	ToFlowOutput() FlowOutput
	ToFlowOutputWithContext(ctx context.Context) FlowOutput
}

func (*Flow) ElementType() reflect.Type {
	return reflect.TypeOf((*Flow)(nil))
}

func (i *Flow) ToFlowOutput() FlowOutput {
	return i.ToFlowOutputWithContext(context.Background())
}

func (i *Flow) ToFlowOutputWithContext(ctx context.Context) FlowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowOutput)
}

type FlowOutput struct {
	*pulumi.OutputState
}

func (FlowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Flow)(nil))
}

func (o FlowOutput) ToFlowOutput() FlowOutput {
	return o
}

func (o FlowOutput) ToFlowOutputWithContext(ctx context.Context) FlowOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FlowOutput{})
}
