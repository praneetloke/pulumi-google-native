// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and starts a Replay using the given ReplayConfig.
type Replay struct {
	pulumi.CustomResourceState

	// Required. The configuration used for the `Replay`.
	Config GoogleCloudPolicysimulatorV1beta1ReplayConfigResponseOutput `pulumi:"config"`
	// The resource name of the `Replay`, which has the following format: `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}`, where `{resource-id}` is the ID of the project, folder, or organization that owns the Replay. Example: `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36`
	Name pulumi.StringOutput `pulumi:"name"`
	// Summary statistics about the replayed log entries.
	ResultsSummary GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponseOutput `pulumi:"resultsSummary"`
	// The current state of the `Replay`.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewReplay registers a new resource with the given unique name, arguments, and options.
func NewReplay(ctx *pulumi.Context,
	name string, args *ReplayArgs, opts ...pulumi.ResourceOption) (*Replay, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.ReplayId == nil {
		return nil, errors.New("invalid value for required argument 'ReplayId'")
	}
	var resource Replay
	err := ctx.RegisterResource("google-native:policysimulator/v1beta1:Replay", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplay gets an existing Replay resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplay(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplayState, opts ...pulumi.ResourceOption) (*Replay, error) {
	var resource Replay
	err := ctx.ReadResource("google-native:policysimulator/v1beta1:Replay", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Replay resources.
type replayState struct {
	// Required. The configuration used for the `Replay`.
	Config *GoogleCloudPolicysimulatorV1beta1ReplayConfigResponse `pulumi:"config"`
	// The resource name of the `Replay`, which has the following format: `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}`, where `{resource-id}` is the ID of the project, folder, or organization that owns the Replay. Example: `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36`
	Name *string `pulumi:"name"`
	// Summary statistics about the replayed log entries.
	ResultsSummary *GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponse `pulumi:"resultsSummary"`
	// The current state of the `Replay`.
	State *string `pulumi:"state"`
}

type ReplayState struct {
	// Required. The configuration used for the `Replay`.
	Config GoogleCloudPolicysimulatorV1beta1ReplayConfigResponsePtrInput
	// The resource name of the `Replay`, which has the following format: `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}`, where `{resource-id}` is the ID of the project, folder, or organization that owns the Replay. Example: `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36`
	Name pulumi.StringPtrInput
	// Summary statistics about the replayed log entries.
	ResultsSummary GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponsePtrInput
	// The current state of the `Replay`.
	State pulumi.StringPtrInput
}

func (ReplayState) ElementType() reflect.Type {
	return reflect.TypeOf((*replayState)(nil)).Elem()
}

type replayArgs struct {
	// Required. The configuration used for the `Replay`.
	Config   *GoogleCloudPolicysimulatorV1beta1ReplayConfig `pulumi:"config"`
	Location string                                         `pulumi:"location"`
	Project  string                                         `pulumi:"project"`
	ReplayId string                                         `pulumi:"replayId"`
}

// The set of arguments for constructing a Replay resource.
type ReplayArgs struct {
	// Required. The configuration used for the `Replay`.
	Config   GoogleCloudPolicysimulatorV1beta1ReplayConfigPtrInput
	Location pulumi.StringInput
	Project  pulumi.StringInput
	ReplayId pulumi.StringInput
}

func (ReplayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replayArgs)(nil)).Elem()
}

type ReplayInput interface {
	pulumi.Input

	ToReplayOutput() ReplayOutput
	ToReplayOutputWithContext(ctx context.Context) ReplayOutput
}

func (*Replay) ElementType() reflect.Type {
	return reflect.TypeOf((*Replay)(nil))
}

func (i *Replay) ToReplayOutput() ReplayOutput {
	return i.ToReplayOutputWithContext(context.Background())
}

func (i *Replay) ToReplayOutputWithContext(ctx context.Context) ReplayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplayOutput)
}

type ReplayOutput struct {
	*pulumi.OutputState
}

func (ReplayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Replay)(nil))
}

func (o ReplayOutput) ToReplayOutput() ReplayOutput {
	return o
}

func (o ReplayOutput) ToReplayOutputWithContext(ctx context.Context) ReplayOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReplayOutput{})
}
