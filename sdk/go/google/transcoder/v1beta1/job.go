// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a job in the specified region.
type Job struct {
	pulumi.CustomResourceState

	// The configuration for this job.
	Config JobConfigResponseOutput `pulumi:"config"`
	// The time the job was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The time the transcoding finished.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// List of failure details. This property may contain additional information about the failure when `failure_reason` is present. *Note*: This feature is not yet available.
	FailureDetails FailureDetailResponseArrayOutput `pulumi:"failureDetails"`
	// A description of the reason for the failure. This property is always present when `state` is `FAILED`.
	FailureReason pulumi.StringOutput `pulumi:"failureReason"`
	// Input only. Specify the `input_uri` to populate empty `uri` fields in each element of `Job.config.inputs` or `JobTemplate.config.inputs` when using template. URI of the media. Input files must be at least 5 seconds in duration and stored in Cloud Storage (for example, `gs://bucket/inputs/file.mp4`).
	InputUri pulumi.StringOutput `pulumi:"inputUri"`
	// The resource name of the job. Format: `projects/{project}/locations/{location}/jobs/{job}`
	Name pulumi.StringOutput `pulumi:"name"`
	// The origin URI. *Note*: This feature is not yet available.
	OriginUri OriginUriResponseOutput `pulumi:"originUri"`
	// Input only. Specify the `output_uri` to populate an empty `Job.config.output.uri` or `JobTemplate.config.output.uri` when using template. URI for the output file(s). For example, `gs://my-bucket/outputs/`.
	OutputUri pulumi.StringOutput `pulumi:"outputUri"`
	// Specify the priority of the job. Enter a value between 0 and 100, where 0 is the lowest priority and 100 is the highest priority. The default is 0.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Estimated fractional progress, from `0` to `1` for each step. *Note*: This feature is not yet available.
	Progress ProgressResponseOutput `pulumi:"progress"`
	// The time the transcoding started.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// The current state of the job.
	State pulumi.StringOutput `pulumi:"state"`
	// Input only. Specify the `template_id` to use for populating `Job.config`. The default is `preset/web-hd`. Preset Transcoder templates: - `preset/{preset_id}` - User defined JobTemplate: `{job_template_id}`
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
	// Job time to live value in days, which will be effective after job completion. Job should be deleted automatically after the given TTL. Enter a value between 1 and 90. The default is 30.
	TtlAfterCompletionDays pulumi.IntOutput `pulumi:"ttlAfterCompletionDays"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobId == nil {
		return nil, errors.New("invalid value for required argument 'JobId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Job
	err := ctx.RegisterResource("google-native:transcoder/v1beta1:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("google-native:transcoder/v1beta1:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// The configuration for this job.
	Config *JobConfigResponse `pulumi:"config"`
	// The time the job was created.
	CreateTime *string `pulumi:"createTime"`
	// The time the transcoding finished.
	EndTime *string `pulumi:"endTime"`
	// List of failure details. This property may contain additional information about the failure when `failure_reason` is present. *Note*: This feature is not yet available.
	FailureDetails []FailureDetailResponse `pulumi:"failureDetails"`
	// A description of the reason for the failure. This property is always present when `state` is `FAILED`.
	FailureReason *string `pulumi:"failureReason"`
	// Input only. Specify the `input_uri` to populate empty `uri` fields in each element of `Job.config.inputs` or `JobTemplate.config.inputs` when using template. URI of the media. Input files must be at least 5 seconds in duration and stored in Cloud Storage (for example, `gs://bucket/inputs/file.mp4`).
	InputUri *string `pulumi:"inputUri"`
	// The resource name of the job. Format: `projects/{project}/locations/{location}/jobs/{job}`
	Name *string `pulumi:"name"`
	// The origin URI. *Note*: This feature is not yet available.
	OriginUri *OriginUriResponse `pulumi:"originUri"`
	// Input only. Specify the `output_uri` to populate an empty `Job.config.output.uri` or `JobTemplate.config.output.uri` when using template. URI for the output file(s). For example, `gs://my-bucket/outputs/`.
	OutputUri *string `pulumi:"outputUri"`
	// Specify the priority of the job. Enter a value between 0 and 100, where 0 is the lowest priority and 100 is the highest priority. The default is 0.
	Priority *int `pulumi:"priority"`
	// Estimated fractional progress, from `0` to `1` for each step. *Note*: This feature is not yet available.
	Progress *ProgressResponse `pulumi:"progress"`
	// The time the transcoding started.
	StartTime *string `pulumi:"startTime"`
	// The current state of the job.
	State *string `pulumi:"state"`
	// Input only. Specify the `template_id` to use for populating `Job.config`. The default is `preset/web-hd`. Preset Transcoder templates: - `preset/{preset_id}` - User defined JobTemplate: `{job_template_id}`
	TemplateId *string `pulumi:"templateId"`
	// Job time to live value in days, which will be effective after job completion. Job should be deleted automatically after the given TTL. Enter a value between 1 and 90. The default is 30.
	TtlAfterCompletionDays *int `pulumi:"ttlAfterCompletionDays"`
}

type JobState struct {
	// The configuration for this job.
	Config JobConfigResponsePtrInput
	// The time the job was created.
	CreateTime pulumi.StringPtrInput
	// The time the transcoding finished.
	EndTime pulumi.StringPtrInput
	// List of failure details. This property may contain additional information about the failure when `failure_reason` is present. *Note*: This feature is not yet available.
	FailureDetails FailureDetailResponseArrayInput
	// A description of the reason for the failure. This property is always present when `state` is `FAILED`.
	FailureReason pulumi.StringPtrInput
	// Input only. Specify the `input_uri` to populate empty `uri` fields in each element of `Job.config.inputs` or `JobTemplate.config.inputs` when using template. URI of the media. Input files must be at least 5 seconds in duration and stored in Cloud Storage (for example, `gs://bucket/inputs/file.mp4`).
	InputUri pulumi.StringPtrInput
	// The resource name of the job. Format: `projects/{project}/locations/{location}/jobs/{job}`
	Name pulumi.StringPtrInput
	// The origin URI. *Note*: This feature is not yet available.
	OriginUri OriginUriResponsePtrInput
	// Input only. Specify the `output_uri` to populate an empty `Job.config.output.uri` or `JobTemplate.config.output.uri` when using template. URI for the output file(s). For example, `gs://my-bucket/outputs/`.
	OutputUri pulumi.StringPtrInput
	// Specify the priority of the job. Enter a value between 0 and 100, where 0 is the lowest priority and 100 is the highest priority. The default is 0.
	Priority pulumi.IntPtrInput
	// Estimated fractional progress, from `0` to `1` for each step. *Note*: This feature is not yet available.
	Progress ProgressResponsePtrInput
	// The time the transcoding started.
	StartTime pulumi.StringPtrInput
	// The current state of the job.
	State pulumi.StringPtrInput
	// Input only. Specify the `template_id` to use for populating `Job.config`. The default is `preset/web-hd`. Preset Transcoder templates: - `preset/{preset_id}` - User defined JobTemplate: `{job_template_id}`
	TemplateId pulumi.StringPtrInput
	// Job time to live value in days, which will be effective after job completion. Job should be deleted automatically after the given TTL. Enter a value between 1 and 90. The default is 30.
	TtlAfterCompletionDays pulumi.IntPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// The configuration for this job.
	Config *JobConfig `pulumi:"config"`
	// Input only. Specify the `input_uri` to populate empty `uri` fields in each element of `Job.config.inputs` or `JobTemplate.config.inputs` when using template. URI of the media. Input files must be at least 5 seconds in duration and stored in Cloud Storage (for example, `gs://bucket/inputs/file.mp4`).
	InputUri *string `pulumi:"inputUri"`
	JobId    string  `pulumi:"jobId"`
	Location string  `pulumi:"location"`
	// The resource name of the job. Format: `projects/{project}/locations/{location}/jobs/{job}`
	Name *string `pulumi:"name"`
	// Input only. Specify the `output_uri` to populate an empty `Job.config.output.uri` or `JobTemplate.config.output.uri` when using template. URI for the output file(s). For example, `gs://my-bucket/outputs/`.
	OutputUri *string `pulumi:"outputUri"`
	// Specify the priority of the job. Enter a value between 0 and 100, where 0 is the lowest priority and 100 is the highest priority. The default is 0.
	Priority *int   `pulumi:"priority"`
	Project  string `pulumi:"project"`
	// Input only. Specify the `template_id` to use for populating `Job.config`. The default is `preset/web-hd`. Preset Transcoder templates: - `preset/{preset_id}` - User defined JobTemplate: `{job_template_id}`
	TemplateId *string `pulumi:"templateId"`
	// Job time to live value in days, which will be effective after job completion. Job should be deleted automatically after the given TTL. Enter a value between 1 and 90. The default is 30.
	TtlAfterCompletionDays *int `pulumi:"ttlAfterCompletionDays"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// The configuration for this job.
	Config JobConfigPtrInput
	// Input only. Specify the `input_uri` to populate empty `uri` fields in each element of `Job.config.inputs` or `JobTemplate.config.inputs` when using template. URI of the media. Input files must be at least 5 seconds in duration and stored in Cloud Storage (for example, `gs://bucket/inputs/file.mp4`).
	InputUri pulumi.StringPtrInput
	JobId    pulumi.StringInput
	Location pulumi.StringInput
	// The resource name of the job. Format: `projects/{project}/locations/{location}/jobs/{job}`
	Name pulumi.StringPtrInput
	// Input only. Specify the `output_uri` to populate an empty `Job.config.output.uri` or `JobTemplate.config.output.uri` when using template. URI for the output file(s). For example, `gs://my-bucket/outputs/`.
	OutputUri pulumi.StringPtrInput
	// Specify the priority of the job. Enter a value between 0 and 100, where 0 is the lowest priority and 100 is the highest priority. The default is 0.
	Priority pulumi.IntPtrInput
	Project  pulumi.StringInput
	// Input only. Specify the `template_id` to use for populating `Job.config`. The default is `preset/web-hd`. Preset Transcoder templates: - `preset/{preset_id}` - User defined JobTemplate: `{job_template_id}`
	TemplateId pulumi.StringPtrInput
	// Job time to live value in days, which will be effective after job completion. Job should be deleted automatically after the given TTL. Enter a value between 1 and 90. The default is 30.
	TtlAfterCompletionDays pulumi.IntPtrInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((*Job)(nil))
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

type JobOutput struct {
	*pulumi.OutputState
}

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Job)(nil))
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(JobOutput{})
}
