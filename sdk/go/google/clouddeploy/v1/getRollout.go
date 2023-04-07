// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single Rollout.
func LookupRollout(ctx *pulumi.Context, args *LookupRolloutArgs, opts ...pulumi.InvokeOption) (*LookupRolloutResult, error) {
	var rv LookupRolloutResult
	err := ctx.Invoke("google-native:clouddeploy/v1:getRollout", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRolloutArgs struct {
	DeliveryPipelineId string  `pulumi:"deliveryPipelineId"`
	Location           string  `pulumi:"location"`
	Project            *string `pulumi:"project"`
	ReleaseId          string  `pulumi:"releaseId"`
	RolloutId          string  `pulumi:"rolloutId"`
}

type LookupRolloutResult struct {
	// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	Annotations map[string]string `pulumi:"annotations"`
	// Approval state of the `Rollout`.
	ApprovalState string `pulumi:"approvalState"`
	// Time at which the `Rollout` was approved.
	ApproveTime string `pulumi:"approveTime"`
	// Name of the `ControllerRollout`. Format is projects/{project}/ locations/{location}/deliveryPipelines/{deliveryPipeline}/ releases/{release}/rollouts/a-z{0,62}.
	ControllerRollout string `pulumi:"controllerRollout"`
	// Time at which the `Rollout` was created.
	CreateTime string `pulumi:"createTime"`
	// Time at which the `Rollout` finished deploying.
	DeployEndTime string `pulumi:"deployEndTime"`
	// The reason this rollout failed. This will always be unspecified while the rollout is in progress.
	DeployFailureCause string `pulumi:"deployFailureCause"`
	// Time at which the `Rollout` started deploying.
	DeployStartTime string `pulumi:"deployStartTime"`
	// The resource name of the Cloud Build `Build` object that is used to deploy the Rollout. Format is `projects/{project}/locations/{location}/builds/{build}`.
	DeployingBuild string `pulumi:"deployingBuild"`
	// Description of the `Rollout` for user purposes. Max length is 255 characters.
	Description string `pulumi:"description"`
	// Time at which the `Rollout` was enqueued.
	EnqueueTime string `pulumi:"enqueueTime"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag string `pulumi:"etag"`
	// Additional information about the rollout failure, if available.
	FailureReason string `pulumi:"failureReason"`
	// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	Labels map[string]string `pulumi:"labels"`
	// Metadata contains information about the rollout.
	Metadata MetadataResponse `pulumi:"metadata"`
	// Optional. Name of the `Rollout`. Format is projects/{project}/ locations/{location}/deliveryPipelines/{deliveryPipeline}/ releases/{release}/rollouts/a-z{0,62}.
	Name string `pulumi:"name"`
	// The phases that represent the workflows of this `Rollout`.
	Phases []PhaseResponse `pulumi:"phases"`
	// Current state of the `Rollout`.
	State string `pulumi:"state"`
	// The ID of Target to which this `Rollout` is deploying.
	TargetId string `pulumi:"targetId"`
	// Unique identifier of the `Rollout`.
	Uid string `pulumi:"uid"`
}

func LookupRolloutOutput(ctx *pulumi.Context, args LookupRolloutOutputArgs, opts ...pulumi.InvokeOption) LookupRolloutResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRolloutResult, error) {
			args := v.(LookupRolloutArgs)
			r, err := LookupRollout(ctx, &args, opts...)
			var s LookupRolloutResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRolloutResultOutput)
}

type LookupRolloutOutputArgs struct {
	DeliveryPipelineId pulumi.StringInput    `pulumi:"deliveryPipelineId"`
	Location           pulumi.StringInput    `pulumi:"location"`
	Project            pulumi.StringPtrInput `pulumi:"project"`
	ReleaseId          pulumi.StringInput    `pulumi:"releaseId"`
	RolloutId          pulumi.StringInput    `pulumi:"rolloutId"`
}

func (LookupRolloutOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRolloutArgs)(nil)).Elem()
}

type LookupRolloutResultOutput struct{ *pulumi.OutputState }

func (LookupRolloutResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRolloutResult)(nil)).Elem()
}

func (o LookupRolloutResultOutput) ToLookupRolloutResultOutput() LookupRolloutResultOutput {
	return o
}

func (o LookupRolloutResultOutput) ToLookupRolloutResultOutputWithContext(ctx context.Context) LookupRolloutResultOutput {
	return o
}

// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
func (o LookupRolloutResultOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRolloutResult) map[string]string { return v.Annotations }).(pulumi.StringMapOutput)
}

// Approval state of the `Rollout`.
func (o LookupRolloutResultOutput) ApprovalState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.ApprovalState }).(pulumi.StringOutput)
}

// Time at which the `Rollout` was approved.
func (o LookupRolloutResultOutput) ApproveTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.ApproveTime }).(pulumi.StringOutput)
}

// Name of the `ControllerRollout`. Format is projects/{project}/ locations/{location}/deliveryPipelines/{deliveryPipeline}/ releases/{release}/rollouts/a-z{0,62}.
func (o LookupRolloutResultOutput) ControllerRollout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.ControllerRollout }).(pulumi.StringOutput)
}

// Time at which the `Rollout` was created.
func (o LookupRolloutResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Time at which the `Rollout` finished deploying.
func (o LookupRolloutResultOutput) DeployEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.DeployEndTime }).(pulumi.StringOutput)
}

// The reason this rollout failed. This will always be unspecified while the rollout is in progress.
func (o LookupRolloutResultOutput) DeployFailureCause() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.DeployFailureCause }).(pulumi.StringOutput)
}

// Time at which the `Rollout` started deploying.
func (o LookupRolloutResultOutput) DeployStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.DeployStartTime }).(pulumi.StringOutput)
}

// The resource name of the Cloud Build `Build` object that is used to deploy the Rollout. Format is `projects/{project}/locations/{location}/builds/{build}`.
func (o LookupRolloutResultOutput) DeployingBuild() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.DeployingBuild }).(pulumi.StringOutput)
}

// Description of the `Rollout` for user purposes. Max length is 255 characters.
func (o LookupRolloutResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.Description }).(pulumi.StringOutput)
}

// Time at which the `Rollout` was enqueued.
func (o LookupRolloutResultOutput) EnqueueTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.EnqueueTime }).(pulumi.StringOutput)
}

// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
func (o LookupRolloutResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Additional information about the rollout failure, if available.
func (o LookupRolloutResultOutput) FailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.FailureReason }).(pulumi.StringOutput)
}

// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
func (o LookupRolloutResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRolloutResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Metadata contains information about the rollout.
func (o LookupRolloutResultOutput) Metadata() MetadataResponseOutput {
	return o.ApplyT(func(v LookupRolloutResult) MetadataResponse { return v.Metadata }).(MetadataResponseOutput)
}

// Optional. Name of the `Rollout`. Format is projects/{project}/ locations/{location}/deliveryPipelines/{deliveryPipeline}/ releases/{release}/rollouts/a-z{0,62}.
func (o LookupRolloutResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.Name }).(pulumi.StringOutput)
}

// The phases that represent the workflows of this `Rollout`.
func (o LookupRolloutResultOutput) Phases() PhaseResponseArrayOutput {
	return o.ApplyT(func(v LookupRolloutResult) []PhaseResponse { return v.Phases }).(PhaseResponseArrayOutput)
}

// Current state of the `Rollout`.
func (o LookupRolloutResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.State }).(pulumi.StringOutput)
}

// The ID of Target to which this `Rollout` is deploying.
func (o LookupRolloutResultOutput) TargetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.TargetId }).(pulumi.StringOutput)
}

// Unique identifier of the `Rollout`.
func (o LookupRolloutResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.Uid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRolloutResultOutput{})
}
