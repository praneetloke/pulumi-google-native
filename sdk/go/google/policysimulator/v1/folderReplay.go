// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and starts a Replay using the given ReplayConfig.
type FolderReplay struct {
	pulumi.CustomResourceState

	// Required. The configuration used for the `Replay`.
	Config GoogleCloudPolicysimulatorV1ReplayConfigResponseOutput `pulumi:"config"`
	// The resource name of the `Replay`, which has the following format: `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}`, where `{resource-id}` is the ID of the project, folder, or organization that owns the Replay. Example: `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36`
	Name pulumi.StringOutput `pulumi:"name"`
	// Summary statistics about the replayed log entries.
	ResultsSummary GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput `pulumi:"resultsSummary"`
	// The current state of the `Replay`.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewFolderReplay registers a new resource with the given unique name, arguments, and options.
func NewFolderReplay(ctx *pulumi.Context,
	name string, args *FolderReplayArgs, opts ...pulumi.ResourceOption) (*FolderReplay, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FolderId == nil {
		return nil, errors.New("invalid value for required argument 'FolderId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ReplayId == nil {
		return nil, errors.New("invalid value for required argument 'ReplayId'")
	}
	var resource FolderReplay
	err := ctx.RegisterResource("google-native:policysimulator/v1:FolderReplay", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolderReplay gets an existing FolderReplay resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolderReplay(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderReplayState, opts ...pulumi.ResourceOption) (*FolderReplay, error) {
	var resource FolderReplay
	err := ctx.ReadResource("google-native:policysimulator/v1:FolderReplay", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FolderReplay resources.
type folderReplayState struct {
	// Required. The configuration used for the `Replay`.
	Config *GoogleCloudPolicysimulatorV1ReplayConfigResponse `pulumi:"config"`
	// The resource name of the `Replay`, which has the following format: `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}`, where `{resource-id}` is the ID of the project, folder, or organization that owns the Replay. Example: `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36`
	Name *string `pulumi:"name"`
	// Summary statistics about the replayed log entries.
	ResultsSummary *GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponse `pulumi:"resultsSummary"`
	// The current state of the `Replay`.
	State *string `pulumi:"state"`
}

type FolderReplayState struct {
	// Required. The configuration used for the `Replay`.
	Config GoogleCloudPolicysimulatorV1ReplayConfigResponsePtrInput
	// The resource name of the `Replay`, which has the following format: `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}`, where `{resource-id}` is the ID of the project, folder, or organization that owns the Replay. Example: `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36`
	Name pulumi.StringPtrInput
	// Summary statistics about the replayed log entries.
	ResultsSummary GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponsePtrInput
	// The current state of the `Replay`.
	State pulumi.StringPtrInput
}

func (FolderReplayState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderReplayState)(nil)).Elem()
}

type folderReplayArgs struct {
	// Required. The configuration used for the `Replay`.
	Config   *GoogleCloudPolicysimulatorV1ReplayConfig `pulumi:"config"`
	FolderId string                                    `pulumi:"folderId"`
	Location string                                    `pulumi:"location"`
	ReplayId string                                    `pulumi:"replayId"`
}

// The set of arguments for constructing a FolderReplay resource.
type FolderReplayArgs struct {
	// Required. The configuration used for the `Replay`.
	Config   GoogleCloudPolicysimulatorV1ReplayConfigPtrInput
	FolderId pulumi.StringInput
	Location pulumi.StringInput
	ReplayId pulumi.StringInput
}

func (FolderReplayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderReplayArgs)(nil)).Elem()
}

type FolderReplayInput interface {
	pulumi.Input

	ToFolderReplayOutput() FolderReplayOutput
	ToFolderReplayOutputWithContext(ctx context.Context) FolderReplayOutput
}

func (*FolderReplay) ElementType() reflect.Type {
	return reflect.TypeOf((*FolderReplay)(nil))
}

func (i *FolderReplay) ToFolderReplayOutput() FolderReplayOutput {
	return i.ToFolderReplayOutputWithContext(context.Background())
}

func (i *FolderReplay) ToFolderReplayOutputWithContext(ctx context.Context) FolderReplayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderReplayOutput)
}

type FolderReplayOutput struct {
	*pulumi.OutputState
}

func (FolderReplayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FolderReplay)(nil))
}

func (o FolderReplayOutput) ToFolderReplayOutput() FolderReplayOutput {
	return o
}

func (o FolderReplayOutput) ToFolderReplayOutputWithContext(ctx context.Context) FolderReplayOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FolderReplayOutput{})
}
