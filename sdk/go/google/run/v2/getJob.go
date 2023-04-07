// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a Job.
func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("google-native:run/v2:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	JobId    string  `pulumi:"jobId"`
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
}

type LookupJobResult struct {
	// KRM-style annotations for the resource. Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects. Cloud Run API v2 does not support annotations with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system annotations in v1 now have a corresponding field in v2 Job. This field follows Kubernetes annotations' namespacing, limits, and rules. More info: https://kubernetes.io/docs/user-guide/annotations
	Annotations map[string]string `pulumi:"annotations"`
	// Settings for the Binary Authorization feature.
	BinaryAuthorization GoogleCloudRunV2BinaryAuthorizationResponse `pulumi:"binaryAuthorization"`
	// Arbitrary identifier for the API client.
	Client string `pulumi:"client"`
	// Arbitrary version identifier for the API client.
	ClientVersion string `pulumi:"clientVersion"`
	// The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Job does not reach its desired state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
	Conditions []GoogleCloudRunV2ConditionResponse `pulumi:"conditions"`
	// The creation time.
	CreateTime string `pulumi:"createTime"`
	// Email address of the authenticated creator.
	Creator string `pulumi:"creator"`
	// The deletion time.
	DeleteTime string `pulumi:"deleteTime"`
	// A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.
	Etag string `pulumi:"etag"`
	// Number of executions created for this job.
	ExecutionCount int `pulumi:"executionCount"`
	// For a deleted resource, the time after which it will be permamently deleted.
	ExpireTime string `pulumi:"expireTime"`
	// A number that monotonically increases every time the user modifies the desired state.
	Generation string `pulumi:"generation"`
	// KRM-style labels for the resource. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run API v2 does not support labels with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system labels in v1 now have a corresponding field in v2 Job.
	Labels map[string]string `pulumi:"labels"`
	// Email address of the last authenticated modifier.
	LastModifier string `pulumi:"lastModifier"`
	// Name of the last created execution.
	LatestCreatedExecution GoogleCloudRunV2ExecutionReferenceResponse `pulumi:"latestCreatedExecution"`
	// The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/terms/launch-stages). Cloud Run supports `ALPHA`, `BETA`, and `GA`. If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features. For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output.
	LaunchStage string `pulumi:"launchStage"`
	// The fully qualified name of this Job. Format: projects/{project}/locations/{location}/jobs/{job}
	Name string `pulumi:"name"`
	// The generation of this Job. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
	ObservedGeneration string `pulumi:"observedGeneration"`
	// Returns true if the Job is currently being acted upon by the system to bring it into the desired state. When a new Job is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Job to the desired state. This process is called reconciliation. While reconciliation is in process, `observed_generation` and `latest_succeeded_execution`, will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the state matches the Job, or there was an error, and reconciliation failed. This state can be found in `terminal_condition.state`. If reconciliation succeeded, the following fields will match: `observed_generation` and `generation`, `latest_succeeded_execution` and `latest_created_execution`. If reconciliation failed, `observed_generation` and `latest_succeeded_execution` will have the state of the last succeeded execution or empty for newly created Job. Additional information on the failure can be found in `terminal_condition` and `conditions`.
	Reconciling bool `pulumi:"reconciling"`
	// Reserved for future use.
	SatisfiesPzs bool `pulumi:"satisfiesPzs"`
	// The template used to create executions for this Job.
	Template GoogleCloudRunV2ExecutionTemplateResponse `pulumi:"template"`
	// The Condition of this Job, containing its readiness status, and detailed error information in case it did not reach the desired state.
	TerminalCondition GoogleCloudRunV2ConditionResponse `pulumi:"terminalCondition"`
	// Server assigned unique identifier for the Execution. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
	Uid string `pulumi:"uid"`
	// The last-modified time.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupJobOutput(ctx *pulumi.Context, args LookupJobOutputArgs, opts ...pulumi.InvokeOption) LookupJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobResult, error) {
			args := v.(LookupJobArgs)
			r, err := LookupJob(ctx, &args, opts...)
			var s LookupJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobResultOutput)
}

type LookupJobOutputArgs struct {
	JobId    pulumi.StringInput    `pulumi:"jobId"`
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobArgs)(nil)).Elem()
}

type LookupJobResultOutput struct{ *pulumi.OutputState }

func (LookupJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobResult)(nil)).Elem()
}

func (o LookupJobResultOutput) ToLookupJobResultOutput() LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) ToLookupJobResultOutputWithContext(ctx context.Context) LookupJobResultOutput {
	return o
}

// KRM-style annotations for the resource. Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects. Cloud Run API v2 does not support annotations with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system annotations in v1 now have a corresponding field in v2 Job. This field follows Kubernetes annotations' namespacing, limits, and rules. More info: https://kubernetes.io/docs/user-guide/annotations
func (o LookupJobResultOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupJobResult) map[string]string { return v.Annotations }).(pulumi.StringMapOutput)
}

// Settings for the Binary Authorization feature.
func (o LookupJobResultOutput) BinaryAuthorization() GoogleCloudRunV2BinaryAuthorizationResponseOutput {
	return o.ApplyT(func(v LookupJobResult) GoogleCloudRunV2BinaryAuthorizationResponse { return v.BinaryAuthorization }).(GoogleCloudRunV2BinaryAuthorizationResponseOutput)
}

// Arbitrary identifier for the API client.
func (o LookupJobResultOutput) Client() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Client }).(pulumi.StringOutput)
}

// Arbitrary version identifier for the API client.
func (o LookupJobResultOutput) ClientVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.ClientVersion }).(pulumi.StringOutput)
}

// The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Job does not reach its desired state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
func (o LookupJobResultOutput) Conditions() GoogleCloudRunV2ConditionResponseArrayOutput {
	return o.ApplyT(func(v LookupJobResult) []GoogleCloudRunV2ConditionResponse { return v.Conditions }).(GoogleCloudRunV2ConditionResponseArrayOutput)
}

// The creation time.
func (o LookupJobResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Email address of the authenticated creator.
func (o LookupJobResultOutput) Creator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Creator }).(pulumi.StringOutput)
}

// The deletion time.
func (o LookupJobResultOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.DeleteTime }).(pulumi.StringOutput)
}

// A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.
func (o LookupJobResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Number of executions created for this job.
func (o LookupJobResultOutput) ExecutionCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupJobResult) int { return v.ExecutionCount }).(pulumi.IntOutput)
}

// For a deleted resource, the time after which it will be permamently deleted.
func (o LookupJobResultOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.ExpireTime }).(pulumi.StringOutput)
}

// A number that monotonically increases every time the user modifies the desired state.
func (o LookupJobResultOutput) Generation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Generation }).(pulumi.StringOutput)
}

// KRM-style labels for the resource. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run API v2 does not support labels with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system labels in v1 now have a corresponding field in v2 Job.
func (o LookupJobResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupJobResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Email address of the last authenticated modifier.
func (o LookupJobResultOutput) LastModifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.LastModifier }).(pulumi.StringOutput)
}

// Name of the last created execution.
func (o LookupJobResultOutput) LatestCreatedExecution() GoogleCloudRunV2ExecutionReferenceResponseOutput {
	return o.ApplyT(func(v LookupJobResult) GoogleCloudRunV2ExecutionReferenceResponse { return v.LatestCreatedExecution }).(GoogleCloudRunV2ExecutionReferenceResponseOutput)
}

// The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/terms/launch-stages). Cloud Run supports `ALPHA`, `BETA`, and `GA`. If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features. For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output.
func (o LookupJobResultOutput) LaunchStage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.LaunchStage }).(pulumi.StringOutput)
}

// The fully qualified name of this Job. Format: projects/{project}/locations/{location}/jobs/{job}
func (o LookupJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Name }).(pulumi.StringOutput)
}

// The generation of this Job. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
func (o LookupJobResultOutput) ObservedGeneration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.ObservedGeneration }).(pulumi.StringOutput)
}

// Returns true if the Job is currently being acted upon by the system to bring it into the desired state. When a new Job is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Job to the desired state. This process is called reconciliation. While reconciliation is in process, `observed_generation` and `latest_succeeded_execution`, will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the state matches the Job, or there was an error, and reconciliation failed. This state can be found in `terminal_condition.state`. If reconciliation succeeded, the following fields will match: `observed_generation` and `generation`, `latest_succeeded_execution` and `latest_created_execution`. If reconciliation failed, `observed_generation` and `latest_succeeded_execution` will have the state of the last succeeded execution or empty for newly created Job. Additional information on the failure can be found in `terminal_condition` and `conditions`.
func (o LookupJobResultOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupJobResult) bool { return v.Reconciling }).(pulumi.BoolOutput)
}

// Reserved for future use.
func (o LookupJobResultOutput) SatisfiesPzs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupJobResult) bool { return v.SatisfiesPzs }).(pulumi.BoolOutput)
}

// The template used to create executions for this Job.
func (o LookupJobResultOutput) Template() GoogleCloudRunV2ExecutionTemplateResponseOutput {
	return o.ApplyT(func(v LookupJobResult) GoogleCloudRunV2ExecutionTemplateResponse { return v.Template }).(GoogleCloudRunV2ExecutionTemplateResponseOutput)
}

// The Condition of this Job, containing its readiness status, and detailed error information in case it did not reach the desired state.
func (o LookupJobResultOutput) TerminalCondition() GoogleCloudRunV2ConditionResponseOutput {
	return o.ApplyT(func(v LookupJobResult) GoogleCloudRunV2ConditionResponse { return v.TerminalCondition }).(GoogleCloudRunV2ConditionResponseOutput)
}

// Server assigned unique identifier for the Execution. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
func (o LookupJobResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Uid }).(pulumi.StringOutput)
}

// The last-modified time.
func (o LookupJobResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobResultOutput{})
}
