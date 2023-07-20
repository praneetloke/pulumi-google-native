// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Replay. Each `Replay` is available for at least 7 days.
func LookupFolderReplay(ctx *pulumi.Context, args *LookupFolderReplayArgs, opts ...pulumi.InvokeOption) (*LookupFolderReplayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFolderReplayResult
	err := ctx.Invoke("google-native:policysimulator/v1:getFolderReplay", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFolderReplayArgs struct {
	FolderId string `pulumi:"folderId"`
	Location string `pulumi:"location"`
	ReplayId string `pulumi:"replayId"`
}

type LookupFolderReplayResult struct {
	// The configuration used for the `Replay`.
	Config GoogleCloudPolicysimulatorV1ReplayConfigResponse `pulumi:"config"`
	// The resource name of the `Replay`, which has the following format: `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}`, where `{resource-id}` is the ID of the project, folder, or organization that owns the Replay. Example: `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36`
	Name string `pulumi:"name"`
	// Summary statistics about the replayed log entries.
	ResultsSummary GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponse `pulumi:"resultsSummary"`
	// The current state of the `Replay`.
	State string `pulumi:"state"`
}

func LookupFolderReplayOutput(ctx *pulumi.Context, args LookupFolderReplayOutputArgs, opts ...pulumi.InvokeOption) LookupFolderReplayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFolderReplayResult, error) {
			args := v.(LookupFolderReplayArgs)
			r, err := LookupFolderReplay(ctx, &args, opts...)
			var s LookupFolderReplayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFolderReplayResultOutput)
}

type LookupFolderReplayOutputArgs struct {
	FolderId pulumi.StringInput `pulumi:"folderId"`
	Location pulumi.StringInput `pulumi:"location"`
	ReplayId pulumi.StringInput `pulumi:"replayId"`
}

func (LookupFolderReplayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderReplayArgs)(nil)).Elem()
}

type LookupFolderReplayResultOutput struct{ *pulumi.OutputState }

func (LookupFolderReplayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderReplayResult)(nil)).Elem()
}

func (o LookupFolderReplayResultOutput) ToLookupFolderReplayResultOutput() LookupFolderReplayResultOutput {
	return o
}

func (o LookupFolderReplayResultOutput) ToLookupFolderReplayResultOutputWithContext(ctx context.Context) LookupFolderReplayResultOutput {
	return o
}

// The configuration used for the `Replay`.
func (o LookupFolderReplayResultOutput) Config() GoogleCloudPolicysimulatorV1ReplayConfigResponseOutput {
	return o.ApplyT(func(v LookupFolderReplayResult) GoogleCloudPolicysimulatorV1ReplayConfigResponse { return v.Config }).(GoogleCloudPolicysimulatorV1ReplayConfigResponseOutput)
}

// The resource name of the `Replay`, which has the following format: `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}`, where `{resource-id}` is the ID of the project, folder, or organization that owns the Replay. Example: `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36`
func (o LookupFolderReplayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderReplayResult) string { return v.Name }).(pulumi.StringOutput)
}

// Summary statistics about the replayed log entries.
func (o LookupFolderReplayResultOutput) ResultsSummary() GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput {
	return o.ApplyT(func(v LookupFolderReplayResult) GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponse {
		return v.ResultsSummary
	}).(GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput)
}

// The current state of the `Replay`.
func (o LookupFolderReplayResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderReplayResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFolderReplayResultOutput{})
}
