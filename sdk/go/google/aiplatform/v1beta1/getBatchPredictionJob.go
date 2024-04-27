// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a BatchPredictionJob
func LookupBatchPredictionJob(ctx *pulumi.Context, args *LookupBatchPredictionJobArgs, opts ...pulumi.InvokeOption) (*LookupBatchPredictionJobResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBatchPredictionJobResult
	err := ctx.Invoke("google-native:aiplatform/v1beta1:getBatchPredictionJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBatchPredictionJobArgs struct {
	BatchPredictionJobId string  `pulumi:"batchPredictionJobId"`
	Location             string  `pulumi:"location"`
	Project              *string `pulumi:"project"`
}

type LookupBatchPredictionJobResult struct {
	// Statistics on completed and failed prediction instances.
	CompletionStats GoogleCloudAiplatformV1beta1CompletionStatsResponse `pulumi:"completionStats"`
	// Time when the BatchPredictionJob was created.
	CreateTime string `pulumi:"createTime"`
	// The config of resources used by the Model during the batch prediction. If the Model supports DEDICATED_RESOURCES this config may be provided (and the job will use these resources), if the Model doesn't support AUTOMATIC_RESOURCES, this config must be provided.
	DedicatedResources GoogleCloudAiplatformV1beta1BatchDedicatedResourcesResponse `pulumi:"dedicatedResources"`
	// For custom-trained Models and AutoML Tabular Models, the container of the DeployedModel instances will send `stderr` and `stdout` streams to Cloud Logging by default. Please note that the logs incur cost, which are subject to [Cloud Logging pricing](https://cloud.google.com/logging/pricing). User can disable container logging by setting this flag to true.
	DisableContainerLogging bool `pulumi:"disableContainerLogging"`
	// The user-defined name of this BatchPredictionJob.
	DisplayName string `pulumi:"displayName"`
	// Customer-managed encryption key options for a BatchPredictionJob. If this is set, then all resources created by the BatchPredictionJob will be encrypted with the provided encryption key.
	EncryptionSpec GoogleCloudAiplatformV1beta1EncryptionSpecResponse `pulumi:"encryptionSpec"`
	// Time when the BatchPredictionJob entered any of the following states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
	EndTime string `pulumi:"endTime"`
	// Only populated when the job's state is JOB_STATE_FAILED or JOB_STATE_CANCELLED.
	Error GoogleRpcStatusResponse `pulumi:"error"`
	// Explanation configuration for this BatchPredictionJob. Can be specified only if generate_explanation is set to `true`. This value overrides the value of Model.explanation_spec. All fields of explanation_spec are optional in the request. If a field of the explanation_spec object is not populated, the corresponding field of the Model.explanation_spec object is inherited.
	ExplanationSpec GoogleCloudAiplatformV1beta1ExplanationSpecResponse `pulumi:"explanationSpec"`
	// Generate explanation with the batch prediction results. When set to `true`, the batch prediction output changes based on the `predictions_format` field of the BatchPredictionJob.output_config object: * `bigquery`: output includes a column named `explanation`. The value is a struct that conforms to the Explanation object. * `jsonl`: The JSON objects on each line include an additional entry keyed `explanation`. The value of the entry is a JSON object that conforms to the Explanation object. * `csv`: Generating explanations for CSV format is not supported. If this field is set to true, either the Model.explanation_spec or explanation_spec must be populated.
	GenerateExplanation bool `pulumi:"generateExplanation"`
	// Input configuration of the instances on which predictions are performed. The schema of any single instance may be specified via the Model's PredictSchemata's instance_schema_uri.
	InputConfig GoogleCloudAiplatformV1beta1BatchPredictionJobInputConfigResponse `pulumi:"inputConfig"`
	// Configuration for how to convert batch prediction input instances to the prediction instances that are sent to the Model.
	InstanceConfig GoogleCloudAiplatformV1beta1BatchPredictionJobInstanceConfigResponse `pulumi:"instanceConfig"`
	// The labels with user-defined metadata to organize BatchPredictionJobs. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. Parameters configuring the batch behavior. Currently only applicable when dedicated_resources are used (in other cases Vertex AI does the tuning itself).
	ManualBatchTuningParameters GoogleCloudAiplatformV1beta1ManualBatchTuningParametersResponse `pulumi:"manualBatchTuningParameters"`
	// The name of the Model resource that produces the predictions via this job, must share the same ancestor Location. Starting this job has no impact on any existing deployments of the Model and their resources. Exactly one of model and unmanaged_container_model must be set. The model resource name may contain version id or version alias to specify the version. Example: `projects/{project}/locations/{location}/models/{model}@2` or `projects/{project}/locations/{location}/models/{model}@golden` if no version is specified, the default version will be deployed. The model resource could also be a publisher model. Example: `publishers/{publisher}/models/{model}` or `projects/{project}/locations/{location}/publishers/{publisher}/models/{model}`
	Model string `pulumi:"model"`
	// Model monitoring config will be used for analysis model behaviors, based on the input and output to the batch prediction job, as well as the provided training dataset.
	ModelMonitoringConfig GoogleCloudAiplatformV1beta1ModelMonitoringConfigResponse `pulumi:"modelMonitoringConfig"`
	// Get batch prediction job monitoring statistics.
	ModelMonitoringStatsAnomalies []GoogleCloudAiplatformV1beta1ModelMonitoringStatsAnomaliesResponse `pulumi:"modelMonitoringStatsAnomalies"`
	// The running status of the model monitoring pipeline.
	ModelMonitoringStatus GoogleRpcStatusResponse `pulumi:"modelMonitoringStatus"`
	// The parameters that govern the predictions. The schema of the parameters may be specified via the Model's PredictSchemata's parameters_schema_uri.
	ModelParameters interface{} `pulumi:"modelParameters"`
	// The version ID of the Model that produces the predictions via this job.
	ModelVersionId string `pulumi:"modelVersionId"`
	// Resource name of the BatchPredictionJob.
	Name string `pulumi:"name"`
	// The Configuration specifying where output predictions should be written. The schema of any single prediction may be specified as a concatenation of Model's PredictSchemata's instance_schema_uri and prediction_schema_uri.
	OutputConfig GoogleCloudAiplatformV1beta1BatchPredictionJobOutputConfigResponse `pulumi:"outputConfig"`
	// Information further describing the output of this job.
	OutputInfo GoogleCloudAiplatformV1beta1BatchPredictionJobOutputInfoResponse `pulumi:"outputInfo"`
	// Partial failures encountered. For example, single files that can't be read. This field never exceeds 20 entries. Status details fields contain standard Google Cloud error details.
	PartialFailures []GoogleRpcStatusResponse `pulumi:"partialFailures"`
	// Information about resources that had been consumed by this job. Provided in real time at best effort basis, as well as a final value once the job completes. Note: This field currently may be not populated for batch predictions that use AutoML Models.
	ResourcesConsumed GoogleCloudAiplatformV1beta1ResourcesConsumedResponse `pulumi:"resourcesConsumed"`
	// The service account that the DeployedModel's container runs as. If not specified, a system generated one will be used, which has minimal permissions and the custom container, if used, may not have enough permission to access other Google Cloud resources. Users deploying the Model must have the `iam.serviceAccounts.actAs` permission on this service account.
	ServiceAccount string `pulumi:"serviceAccount"`
	// Time when the BatchPredictionJob for the first time entered the `JOB_STATE_RUNNING` state.
	StartTime string `pulumi:"startTime"`
	// The detailed state of the job.
	State string `pulumi:"state"`
	// Contains model information necessary to perform batch prediction without requiring uploading to model registry. Exactly one of model and unmanaged_container_model must be set.
	UnmanagedContainerModel GoogleCloudAiplatformV1beta1UnmanagedContainerModelResponse `pulumi:"unmanagedContainerModel"`
	// Time when the BatchPredictionJob was most recently updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupBatchPredictionJobOutput(ctx *pulumi.Context, args LookupBatchPredictionJobOutputArgs, opts ...pulumi.InvokeOption) LookupBatchPredictionJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBatchPredictionJobResult, error) {
			args := v.(LookupBatchPredictionJobArgs)
			r, err := LookupBatchPredictionJob(ctx, &args, opts...)
			var s LookupBatchPredictionJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBatchPredictionJobResultOutput)
}

type LookupBatchPredictionJobOutputArgs struct {
	BatchPredictionJobId pulumi.StringInput    `pulumi:"batchPredictionJobId"`
	Location             pulumi.StringInput    `pulumi:"location"`
	Project              pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupBatchPredictionJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBatchPredictionJobArgs)(nil)).Elem()
}

type LookupBatchPredictionJobResultOutput struct{ *pulumi.OutputState }

func (LookupBatchPredictionJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBatchPredictionJobResult)(nil)).Elem()
}

func (o LookupBatchPredictionJobResultOutput) ToLookupBatchPredictionJobResultOutput() LookupBatchPredictionJobResultOutput {
	return o
}

func (o LookupBatchPredictionJobResultOutput) ToLookupBatchPredictionJobResultOutputWithContext(ctx context.Context) LookupBatchPredictionJobResultOutput {
	return o
}

// Statistics on completed and failed prediction instances.
func (o LookupBatchPredictionJobResultOutput) CompletionStats() GoogleCloudAiplatformV1beta1CompletionStatsResponseOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) GoogleCloudAiplatformV1beta1CompletionStatsResponse {
		return v.CompletionStats
	}).(GoogleCloudAiplatformV1beta1CompletionStatsResponseOutput)
}

// Time when the BatchPredictionJob was created.
func (o LookupBatchPredictionJobResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The config of resources used by the Model during the batch prediction. If the Model supports DEDICATED_RESOURCES this config may be provided (and the job will use these resources), if the Model doesn't support AUTOMATIC_RESOURCES, this config must be provided.
func (o LookupBatchPredictionJobResultOutput) DedicatedResources() GoogleCloudAiplatformV1beta1BatchDedicatedResourcesResponseOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) GoogleCloudAiplatformV1beta1BatchDedicatedResourcesResponse {
		return v.DedicatedResources
	}).(GoogleCloudAiplatformV1beta1BatchDedicatedResourcesResponseOutput)
}

// For custom-trained Models and AutoML Tabular Models, the container of the DeployedModel instances will send `stderr` and `stdout` streams to Cloud Logging by default. Please note that the logs incur cost, which are subject to [Cloud Logging pricing](https://cloud.google.com/logging/pricing). User can disable container logging by setting this flag to true.
func (o LookupBatchPredictionJobResultOutput) DisableContainerLogging() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) bool { return v.DisableContainerLogging }).(pulumi.BoolOutput)
}

// The user-defined name of this BatchPredictionJob.
func (o LookupBatchPredictionJobResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Customer-managed encryption key options for a BatchPredictionJob. If this is set, then all resources created by the BatchPredictionJob will be encrypted with the provided encryption key.
func (o LookupBatchPredictionJobResultOutput) EncryptionSpec() GoogleCloudAiplatformV1beta1EncryptionSpecResponseOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) GoogleCloudAiplatformV1beta1EncryptionSpecResponse {
		return v.EncryptionSpec
	}).(GoogleCloudAiplatformV1beta1EncryptionSpecResponseOutput)
}

// Time when the BatchPredictionJob entered any of the following states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
func (o LookupBatchPredictionJobResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// Only populated when the job's state is JOB_STATE_FAILED or JOB_STATE_CANCELLED.
func (o LookupBatchPredictionJobResultOutput) Error() GoogleRpcStatusResponseOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) GoogleRpcStatusResponse { return v.Error }).(GoogleRpcStatusResponseOutput)
}

// Explanation configuration for this BatchPredictionJob. Can be specified only if generate_explanation is set to `true`. This value overrides the value of Model.explanation_spec. All fields of explanation_spec are optional in the request. If a field of the explanation_spec object is not populated, the corresponding field of the Model.explanation_spec object is inherited.
func (o LookupBatchPredictionJobResultOutput) ExplanationSpec() GoogleCloudAiplatformV1beta1ExplanationSpecResponseOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) GoogleCloudAiplatformV1beta1ExplanationSpecResponse {
		return v.ExplanationSpec
	}).(GoogleCloudAiplatformV1beta1ExplanationSpecResponseOutput)
}

// Generate explanation with the batch prediction results. When set to `true`, the batch prediction output changes based on the `predictions_format` field of the BatchPredictionJob.output_config object: * `bigquery`: output includes a column named `explanation`. The value is a struct that conforms to the Explanation object. * `jsonl`: The JSON objects on each line include an additional entry keyed `explanation`. The value of the entry is a JSON object that conforms to the Explanation object. * `csv`: Generating explanations for CSV format is not supported. If this field is set to true, either the Model.explanation_spec or explanation_spec must be populated.
func (o LookupBatchPredictionJobResultOutput) GenerateExplanation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) bool { return v.GenerateExplanation }).(pulumi.BoolOutput)
}

// Input configuration of the instances on which predictions are performed. The schema of any single instance may be specified via the Model's PredictSchemata's instance_schema_uri.
func (o LookupBatchPredictionJobResultOutput) InputConfig() GoogleCloudAiplatformV1beta1BatchPredictionJobInputConfigResponseOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) GoogleCloudAiplatformV1beta1BatchPredictionJobInputConfigResponse {
		return v.InputConfig
	}).(GoogleCloudAiplatformV1beta1BatchPredictionJobInputConfigResponseOutput)
}

// Configuration for how to convert batch prediction input instances to the prediction instances that are sent to the Model.
func (o LookupBatchPredictionJobResultOutput) InstanceConfig() GoogleCloudAiplatformV1beta1BatchPredictionJobInstanceConfigResponseOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) GoogleCloudAiplatformV1beta1BatchPredictionJobInstanceConfigResponse {
		return v.InstanceConfig
	}).(GoogleCloudAiplatformV1beta1BatchPredictionJobInstanceConfigResponseOutput)
}

// The labels with user-defined metadata to organize BatchPredictionJobs. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
func (o LookupBatchPredictionJobResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Immutable. Parameters configuring the batch behavior. Currently only applicable when dedicated_resources are used (in other cases Vertex AI does the tuning itself).
func (o LookupBatchPredictionJobResultOutput) ManualBatchTuningParameters() GoogleCloudAiplatformV1beta1ManualBatchTuningParametersResponseOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) GoogleCloudAiplatformV1beta1ManualBatchTuningParametersResponse {
		return v.ManualBatchTuningParameters
	}).(GoogleCloudAiplatformV1beta1ManualBatchTuningParametersResponseOutput)
}

// The name of the Model resource that produces the predictions via this job, must share the same ancestor Location. Starting this job has no impact on any existing deployments of the Model and their resources. Exactly one of model and unmanaged_container_model must be set. The model resource name may contain version id or version alias to specify the version. Example: `projects/{project}/locations/{location}/models/{model}@2` or `projects/{project}/locations/{location}/models/{model}@golden` if no version is specified, the default version will be deployed. The model resource could also be a publisher model. Example: `publishers/{publisher}/models/{model}` or `projects/{project}/locations/{location}/publishers/{publisher}/models/{model}`
func (o LookupBatchPredictionJobResultOutput) Model() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) string { return v.Model }).(pulumi.StringOutput)
}

// Model monitoring config will be used for analysis model behaviors, based on the input and output to the batch prediction job, as well as the provided training dataset.
func (o LookupBatchPredictionJobResultOutput) ModelMonitoringConfig() GoogleCloudAiplatformV1beta1ModelMonitoringConfigResponseOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) GoogleCloudAiplatformV1beta1ModelMonitoringConfigResponse {
		return v.ModelMonitoringConfig
	}).(GoogleCloudAiplatformV1beta1ModelMonitoringConfigResponseOutput)
}

// Get batch prediction job monitoring statistics.
func (o LookupBatchPredictionJobResultOutput) ModelMonitoringStatsAnomalies() GoogleCloudAiplatformV1beta1ModelMonitoringStatsAnomaliesResponseArrayOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) []GoogleCloudAiplatformV1beta1ModelMonitoringStatsAnomaliesResponse {
		return v.ModelMonitoringStatsAnomalies
	}).(GoogleCloudAiplatformV1beta1ModelMonitoringStatsAnomaliesResponseArrayOutput)
}

// The running status of the model monitoring pipeline.
func (o LookupBatchPredictionJobResultOutput) ModelMonitoringStatus() GoogleRpcStatusResponseOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) GoogleRpcStatusResponse { return v.ModelMonitoringStatus }).(GoogleRpcStatusResponseOutput)
}

// The parameters that govern the predictions. The schema of the parameters may be specified via the Model's PredictSchemata's parameters_schema_uri.
func (o LookupBatchPredictionJobResultOutput) ModelParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) interface{} { return v.ModelParameters }).(pulumi.AnyOutput)
}

// The version ID of the Model that produces the predictions via this job.
func (o LookupBatchPredictionJobResultOutput) ModelVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) string { return v.ModelVersionId }).(pulumi.StringOutput)
}

// Resource name of the BatchPredictionJob.
func (o LookupBatchPredictionJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Configuration specifying where output predictions should be written. The schema of any single prediction may be specified as a concatenation of Model's PredictSchemata's instance_schema_uri and prediction_schema_uri.
func (o LookupBatchPredictionJobResultOutput) OutputConfig() GoogleCloudAiplatformV1beta1BatchPredictionJobOutputConfigResponseOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) GoogleCloudAiplatformV1beta1BatchPredictionJobOutputConfigResponse {
		return v.OutputConfig
	}).(GoogleCloudAiplatformV1beta1BatchPredictionJobOutputConfigResponseOutput)
}

// Information further describing the output of this job.
func (o LookupBatchPredictionJobResultOutput) OutputInfo() GoogleCloudAiplatformV1beta1BatchPredictionJobOutputInfoResponseOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) GoogleCloudAiplatformV1beta1BatchPredictionJobOutputInfoResponse {
		return v.OutputInfo
	}).(GoogleCloudAiplatformV1beta1BatchPredictionJobOutputInfoResponseOutput)
}

// Partial failures encountered. For example, single files that can't be read. This field never exceeds 20 entries. Status details fields contain standard Google Cloud error details.
func (o LookupBatchPredictionJobResultOutput) PartialFailures() GoogleRpcStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) []GoogleRpcStatusResponse { return v.PartialFailures }).(GoogleRpcStatusResponseArrayOutput)
}

// Information about resources that had been consumed by this job. Provided in real time at best effort basis, as well as a final value once the job completes. Note: This field currently may be not populated for batch predictions that use AutoML Models.
func (o LookupBatchPredictionJobResultOutput) ResourcesConsumed() GoogleCloudAiplatformV1beta1ResourcesConsumedResponseOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) GoogleCloudAiplatformV1beta1ResourcesConsumedResponse {
		return v.ResourcesConsumed
	}).(GoogleCloudAiplatformV1beta1ResourcesConsumedResponseOutput)
}

// The service account that the DeployedModel's container runs as. If not specified, a system generated one will be used, which has minimal permissions and the custom container, if used, may not have enough permission to access other Google Cloud resources. Users deploying the Model must have the `iam.serviceAccounts.actAs` permission on this service account.
func (o LookupBatchPredictionJobResultOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) string { return v.ServiceAccount }).(pulumi.StringOutput)
}

// Time when the BatchPredictionJob for the first time entered the `JOB_STATE_RUNNING` state.
func (o LookupBatchPredictionJobResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// The detailed state of the job.
func (o LookupBatchPredictionJobResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) string { return v.State }).(pulumi.StringOutput)
}

// Contains model information necessary to perform batch prediction without requiring uploading to model registry. Exactly one of model and unmanaged_container_model must be set.
func (o LookupBatchPredictionJobResultOutput) UnmanagedContainerModel() GoogleCloudAiplatformV1beta1UnmanagedContainerModelResponseOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) GoogleCloudAiplatformV1beta1UnmanagedContainerModelResponse {
		return v.UnmanagedContainerModel
	}).(GoogleCloudAiplatformV1beta1UnmanagedContainerModelResponseOutput)
}

// Time when the BatchPredictionJob was most recently updated.
func (o LookupBatchPredictionJobResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchPredictionJobResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBatchPredictionJobResultOutput{})
}