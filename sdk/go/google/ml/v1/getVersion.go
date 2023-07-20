// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a model version. Models can have multiple versions. You can call projects.models.versions.list to get the same information that this method returns for all of the versions of a model.
func LookupVersion(ctx *pulumi.Context, args *LookupVersionArgs, opts ...pulumi.InvokeOption) (*LookupVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVersionResult
	err := ctx.Invoke("google-native:ml/v1:getVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVersionArgs struct {
	ModelId   string  `pulumi:"modelId"`
	Project   *string `pulumi:"project"`
	VersionId string  `pulumi:"versionId"`
}

type LookupVersionResult struct {
	// Optional. Accelerator config for using GPUs for online prediction (beta). Only specify this field if you have specified a Compute Engine (N1) machine type in the `machineType` field. Learn more about [using GPUs for online prediction](/ml-engine/docs/machine-types-online-prediction#gpus).
	AcceleratorConfig GoogleCloudMlV1__AcceleratorConfigResponse `pulumi:"acceleratorConfig"`
	// Automatically scale the number of nodes used to serve the model in response to increases and decreases in traffic. Care should be taken to ramp up traffic according to the model's ability to scale or you will start seeing increases in latency and 429 response codes.
	AutoScaling GoogleCloudMlV1__AutoScalingResponse `pulumi:"autoScaling"`
	// Optional. Specifies a custom container to use for serving predictions. If you specify this field, then `machineType` is required. If you specify this field, then `deploymentUri` is optional. If you specify this field, then you must not specify `runtimeVersion`, `packageUris`, `framework`, `pythonVersion`, or `predictionClass`.
	Container GoogleCloudMlV1__ContainerSpecResponse `pulumi:"container"`
	// The time the version was created.
	CreateTime string `pulumi:"createTime"`
	// The Cloud Storage URI of a directory containing trained model artifacts to be used to create the model version. See the [guide to deploying models](/ai-platform/prediction/docs/deploying-models) for more information. The total number of files under this directory must not exceed 1000. During projects.models.versions.create, AI Platform Prediction copies all files from the specified directory to a location managed by the service. From then on, AI Platform Prediction uses these copies of the model artifacts to serve predictions, not the original files in Cloud Storage, so this location is useful only as a historical record. If you specify container, then this field is optional. Otherwise, it is required. Learn [how to use this field with a custom container](/ai-platform/prediction/docs/custom-container-requirements#artifacts).
	DeploymentUri string `pulumi:"deploymentUri"`
	// Optional. The description specified for the version when it was created.
	Description string `pulumi:"description"`
	// The details of a failure or a cancellation.
	ErrorMessage string `pulumi:"errorMessage"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a model from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform model updates in order to avoid race conditions: An `etag` is returned in the response to `GetVersion`, and systems are expected to put that etag in the request to `UpdateVersion` to ensure that their change will be applied to the model as intended.
	Etag string `pulumi:"etag"`
	// Optional. Configures explainability features on the model's version. Some explanation features require additional metadata to be loaded as part of the model payload.
	ExplanationConfig GoogleCloudMlV1__ExplanationConfigResponse `pulumi:"explanationConfig"`
	// Optional. The machine learning framework AI Platform uses to train this version of the model. Valid values are `TENSORFLOW`, `SCIKIT_LEARN`, `XGBOOST`. If you do not specify a framework, AI Platform will analyze files in the deployment_uri to determine a framework. If you choose `SCIKIT_LEARN` or `XGBOOST`, you must also set the runtime version of the model to 1.4 or greater. Do **not** specify a framework if you're deploying a [custom prediction routine](/ai-platform/prediction/docs/custom-prediction-routines) or if you're using a [custom container](/ai-platform/prediction/docs/use-custom-container).
	Framework string `pulumi:"framework"`
	// If true, this version will be used to handle prediction requests that do not specify a version. You can change the default version by calling projects.methods.versions.setDefault.
	IsDefault bool `pulumi:"isDefault"`
	// Optional. One or more labels that you can add, to organize your model versions. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels. Note that this field is not updatable for mls1* models.
	Labels map[string]string `pulumi:"labels"`
	// The [AI Platform (Unified) `Model`](https://cloud.google.com/ai-platform-unified/docs/reference/rest/v1beta1/projects.locations.models) ID for the last [model migration](https://cloud.google.com/ai-platform-unified/docs/start/migrating-to-ai-platform-unified).
	LastMigrationModelId string `pulumi:"lastMigrationModelId"`
	// The last time this version was successfully [migrated to AI Platform (Unified)](https://cloud.google.com/ai-platform-unified/docs/start/migrating-to-ai-platform-unified).
	LastMigrationTime string `pulumi:"lastMigrationTime"`
	// The time the version was last used for prediction.
	LastUseTime string `pulumi:"lastUseTime"`
	// Optional. The type of machine on which to serve the model. Currently only applies to online prediction service. To learn about valid values for this field, read [Choosing a machine type for online prediction](/ai-platform/prediction/docs/machine-types-online-prediction). If this field is not specified and you are using a [regional endpoint](/ai-platform/prediction/docs/regional-endpoints), then the machine type defaults to `n1-standard-2`. If this field is not specified and you are using the global endpoint (`ml.googleapis.com`), then the machine type defaults to `mls1-c1-m2`.
	MachineType string `pulumi:"machineType"`
	// Manually select the number of nodes to use for serving the model. You should generally use `auto_scaling` with an appropriate `min_nodes` instead, but this option is available if you want more predictable billing. Beware that latency and error rates will increase if the traffic exceeds that capability of the system to serve it based on the selected number of nodes.
	ManualScaling GoogleCloudMlV1__ManualScalingResponse `pulumi:"manualScaling"`
	// The name specified for the version when it was created. The version name must be unique within the model it is created in.
	Name string `pulumi:"name"`
	// Optional. Cloud Storage paths (`gs://…`) of packages for [custom prediction routines](/ml-engine/docs/tensorflow/custom-prediction-routines) or [scikit-learn pipelines with custom code](/ml-engine/docs/scikit/exporting-for-prediction#custom-pipeline-code). For a custom prediction routine, one of these packages must contain your Predictor class (see [`predictionClass`](#Version.FIELDS.prediction_class)). Additionally, include any dependencies used by your Predictor or scikit-learn pipeline uses that are not already included in your selected [runtime version](/ml-engine/docs/tensorflow/runtime-version-list). If you specify this field, you must also set [`runtimeVersion`](#Version.FIELDS.runtime_version) to 1.4 or greater.
	PackageUris []string `pulumi:"packageUris"`
	// Optional. The fully qualified name (module_name.class_name) of a class that implements the Predictor interface described in this reference field. The module containing this class should be included in a package provided to the [`packageUris` field](#Version.FIELDS.package_uris). Specify this field if and only if you are deploying a [custom prediction routine (beta)](/ml-engine/docs/tensorflow/custom-prediction-routines). If you specify this field, you must set [`runtimeVersion`](#Version.FIELDS.runtime_version) to 1.4 or greater and you must set `machineType` to a [legacy (MLS1) machine type](/ml-engine/docs/machine-types-online-prediction). The following code sample provides the Predictor interface: class Predictor(object): """Interface for constructing custom predictors.""" def predict(self, instances, **kwargs): """Performs custom prediction. Instances are the decoded values from the request. They have already been deserialized from JSON. Args: instances: A list of prediction input instances. **kwargs: A dictionary of keyword args provided as additional fields on the predict request body. Returns: A list of outputs containing the prediction results. This list must be JSON serializable. """ raise NotImplementedError() @classmethod def from_path(cls, model_dir): """Creates an instance of Predictor using the given path. Loading of the predictor should be done in this method. Args: model_dir: The local directory that contains the exported model file along with any additional files uploaded when creating the version resource. Returns: An instance implementing this Predictor class. """ raise NotImplementedError() Learn more about [the Predictor interface and custom prediction routines](/ml-engine/docs/tensorflow/custom-prediction-routines).
	PredictionClass string `pulumi:"predictionClass"`
	// The version of Python used in prediction. The following Python versions are available: * Python '3.7' is available when `runtime_version` is set to '1.15' or later. * Python '3.5' is available when `runtime_version` is set to a version from '1.4' to '1.14'. * Python '2.7' is available when `runtime_version` is set to '1.15' or earlier. Read more about the Python versions available for [each runtime version](/ml-engine/docs/runtime-version-list).
	PythonVersion string `pulumi:"pythonVersion"`
	// Optional. *Only* specify this field in a projects.models.versions.patch request. Specifying it in a projects.models.versions.create request has no effect. Configures the request-response pair logging on predictions from this Version.
	RequestLoggingConfig GoogleCloudMlV1__RequestLoggingConfigResponse `pulumi:"requestLoggingConfig"`
	// Optional. Specifies paths on a custom container's HTTP server where AI Platform Prediction sends certain requests. If you specify this field, then you must also specify the `container` field. If you specify the `container` field and do not specify this field, it defaults to the following: ``` json { "predict": "/v1/models/MODEL/versions/VERSION:predict", "health": "/v1/models/MODEL/versions/VERSION" }  ``` See RouteMap for more details about these default values.
	Routes GoogleCloudMlV1__RouteMapResponse `pulumi:"routes"`
	// The AI Platform runtime version to use for this deployment. For more information, see the [runtime version list](/ml-engine/docs/runtime-version-list) and [how to manage runtime versions](/ml-engine/docs/versioning).
	RuntimeVersion string `pulumi:"runtimeVersion"`
	// Optional. Specifies the service account for resource access control. If you specify this field, then you must also specify either the `containerSpec` or the `predictionClass` field. Learn more about [using a custom service account](/ai-platform/prediction/docs/custom-service-account).
	ServiceAccount string `pulumi:"serviceAccount"`
	// The state of a version.
	State string `pulumi:"state"`
}

func LookupVersionOutput(ctx *pulumi.Context, args LookupVersionOutputArgs, opts ...pulumi.InvokeOption) LookupVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVersionResult, error) {
			args := v.(LookupVersionArgs)
			r, err := LookupVersion(ctx, &args, opts...)
			var s LookupVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVersionResultOutput)
}

type LookupVersionOutputArgs struct {
	ModelId   pulumi.StringInput    `pulumi:"modelId"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
	VersionId pulumi.StringInput    `pulumi:"versionId"`
}

func (LookupVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVersionArgs)(nil)).Elem()
}

type LookupVersionResultOutput struct{ *pulumi.OutputState }

func (LookupVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVersionResult)(nil)).Elem()
}

func (o LookupVersionResultOutput) ToLookupVersionResultOutput() LookupVersionResultOutput {
	return o
}

func (o LookupVersionResultOutput) ToLookupVersionResultOutputWithContext(ctx context.Context) LookupVersionResultOutput {
	return o
}

// Optional. Accelerator config for using GPUs for online prediction (beta). Only specify this field if you have specified a Compute Engine (N1) machine type in the `machineType` field. Learn more about [using GPUs for online prediction](/ml-engine/docs/machine-types-online-prediction#gpus).
func (o LookupVersionResultOutput) AcceleratorConfig() GoogleCloudMlV1__AcceleratorConfigResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) GoogleCloudMlV1__AcceleratorConfigResponse { return v.AcceleratorConfig }).(GoogleCloudMlV1__AcceleratorConfigResponseOutput)
}

// Automatically scale the number of nodes used to serve the model in response to increases and decreases in traffic. Care should be taken to ramp up traffic according to the model's ability to scale or you will start seeing increases in latency and 429 response codes.
func (o LookupVersionResultOutput) AutoScaling() GoogleCloudMlV1__AutoScalingResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) GoogleCloudMlV1__AutoScalingResponse { return v.AutoScaling }).(GoogleCloudMlV1__AutoScalingResponseOutput)
}

// Optional. Specifies a custom container to use for serving predictions. If you specify this field, then `machineType` is required. If you specify this field, then `deploymentUri` is optional. If you specify this field, then you must not specify `runtimeVersion`, `packageUris`, `framework`, `pythonVersion`, or `predictionClass`.
func (o LookupVersionResultOutput) Container() GoogleCloudMlV1__ContainerSpecResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) GoogleCloudMlV1__ContainerSpecResponse { return v.Container }).(GoogleCloudMlV1__ContainerSpecResponseOutput)
}

// The time the version was created.
func (o LookupVersionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The Cloud Storage URI of a directory containing trained model artifacts to be used to create the model version. See the [guide to deploying models](/ai-platform/prediction/docs/deploying-models) for more information. The total number of files under this directory must not exceed 1000. During projects.models.versions.create, AI Platform Prediction copies all files from the specified directory to a location managed by the service. From then on, AI Platform Prediction uses these copies of the model artifacts to serve predictions, not the original files in Cloud Storage, so this location is useful only as a historical record. If you specify container, then this field is optional. Otherwise, it is required. Learn [how to use this field with a custom container](/ai-platform/prediction/docs/custom-container-requirements#artifacts).
func (o LookupVersionResultOutput) DeploymentUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.DeploymentUri }).(pulumi.StringOutput)
}

// Optional. The description specified for the version when it was created.
func (o LookupVersionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.Description }).(pulumi.StringOutput)
}

// The details of a failure or a cancellation.
func (o LookupVersionResultOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a model from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform model updates in order to avoid race conditions: An `etag` is returned in the response to `GetVersion`, and systems are expected to put that etag in the request to `UpdateVersion` to ensure that their change will be applied to the model as intended.
func (o LookupVersionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Optional. Configures explainability features on the model's version. Some explanation features require additional metadata to be loaded as part of the model payload.
func (o LookupVersionResultOutput) ExplanationConfig() GoogleCloudMlV1__ExplanationConfigResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) GoogleCloudMlV1__ExplanationConfigResponse { return v.ExplanationConfig }).(GoogleCloudMlV1__ExplanationConfigResponseOutput)
}

// Optional. The machine learning framework AI Platform uses to train this version of the model. Valid values are `TENSORFLOW`, `SCIKIT_LEARN`, `XGBOOST`. If you do not specify a framework, AI Platform will analyze files in the deployment_uri to determine a framework. If you choose `SCIKIT_LEARN` or `XGBOOST`, you must also set the runtime version of the model to 1.4 or greater. Do **not** specify a framework if you're deploying a [custom prediction routine](/ai-platform/prediction/docs/custom-prediction-routines) or if you're using a [custom container](/ai-platform/prediction/docs/use-custom-container).
func (o LookupVersionResultOutput) Framework() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.Framework }).(pulumi.StringOutput)
}

// If true, this version will be used to handle prediction requests that do not specify a version. You can change the default version by calling projects.methods.versions.setDefault.
func (o LookupVersionResultOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVersionResult) bool { return v.IsDefault }).(pulumi.BoolOutput)
}

// Optional. One or more labels that you can add, to organize your model versions. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels. Note that this field is not updatable for mls1* models.
func (o LookupVersionResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVersionResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The [AI Platform (Unified) `Model`](https://cloud.google.com/ai-platform-unified/docs/reference/rest/v1beta1/projects.locations.models) ID for the last [model migration](https://cloud.google.com/ai-platform-unified/docs/start/migrating-to-ai-platform-unified).
func (o LookupVersionResultOutput) LastMigrationModelId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.LastMigrationModelId }).(pulumi.StringOutput)
}

// The last time this version was successfully [migrated to AI Platform (Unified)](https://cloud.google.com/ai-platform-unified/docs/start/migrating-to-ai-platform-unified).
func (o LookupVersionResultOutput) LastMigrationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.LastMigrationTime }).(pulumi.StringOutput)
}

// The time the version was last used for prediction.
func (o LookupVersionResultOutput) LastUseTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.LastUseTime }).(pulumi.StringOutput)
}

// Optional. The type of machine on which to serve the model. Currently only applies to online prediction service. To learn about valid values for this field, read [Choosing a machine type for online prediction](/ai-platform/prediction/docs/machine-types-online-prediction). If this field is not specified and you are using a [regional endpoint](/ai-platform/prediction/docs/regional-endpoints), then the machine type defaults to `n1-standard-2`. If this field is not specified and you are using the global endpoint (`ml.googleapis.com`), then the machine type defaults to `mls1-c1-m2`.
func (o LookupVersionResultOutput) MachineType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.MachineType }).(pulumi.StringOutput)
}

// Manually select the number of nodes to use for serving the model. You should generally use `auto_scaling` with an appropriate `min_nodes` instead, but this option is available if you want more predictable billing. Beware that latency and error rates will increase if the traffic exceeds that capability of the system to serve it based on the selected number of nodes.
func (o LookupVersionResultOutput) ManualScaling() GoogleCloudMlV1__ManualScalingResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) GoogleCloudMlV1__ManualScalingResponse { return v.ManualScaling }).(GoogleCloudMlV1__ManualScalingResponseOutput)
}

// The name specified for the version when it was created. The version name must be unique within the model it is created in.
func (o LookupVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. Cloud Storage paths (`gs://…`) of packages for [custom prediction routines](/ml-engine/docs/tensorflow/custom-prediction-routines) or [scikit-learn pipelines with custom code](/ml-engine/docs/scikit/exporting-for-prediction#custom-pipeline-code). For a custom prediction routine, one of these packages must contain your Predictor class (see [`predictionClass`](#Version.FIELDS.prediction_class)). Additionally, include any dependencies used by your Predictor or scikit-learn pipeline uses that are not already included in your selected [runtime version](/ml-engine/docs/tensorflow/runtime-version-list). If you specify this field, you must also set [`runtimeVersion`](#Version.FIELDS.runtime_version) to 1.4 or greater.
func (o LookupVersionResultOutput) PackageUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVersionResult) []string { return v.PackageUris }).(pulumi.StringArrayOutput)
}

// Optional. The fully qualified name (module_name.class_name) of a class that implements the Predictor interface described in this reference field. The module containing this class should be included in a package provided to the [`packageUris` field](#Version.FIELDS.package_uris). Specify this field if and only if you are deploying a [custom prediction routine (beta)](/ml-engine/docs/tensorflow/custom-prediction-routines). If you specify this field, you must set [`runtimeVersion`](#Version.FIELDS.runtime_version) to 1.4 or greater and you must set `machineType` to a [legacy (MLS1) machine type](/ml-engine/docs/machine-types-online-prediction). The following code sample provides the Predictor interface: class Predictor(object): """Interface for constructing custom predictors.""" def predict(self, instances, **kwargs): """Performs custom prediction. Instances are the decoded values from the request. They have already been deserialized from JSON. Args: instances: A list of prediction input instances. **kwargs: A dictionary of keyword args provided as additional fields on the predict request body. Returns: A list of outputs containing the prediction results. This list must be JSON serializable. """ raise NotImplementedError() @classmethod def from_path(cls, model_dir): """Creates an instance of Predictor using the given path. Loading of the predictor should be done in this method. Args: model_dir: The local directory that contains the exported model file along with any additional files uploaded when creating the version resource. Returns: An instance implementing this Predictor class. """ raise NotImplementedError() Learn more about [the Predictor interface and custom prediction routines](/ml-engine/docs/tensorflow/custom-prediction-routines).
func (o LookupVersionResultOutput) PredictionClass() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.PredictionClass }).(pulumi.StringOutput)
}

// The version of Python used in prediction. The following Python versions are available: * Python '3.7' is available when `runtime_version` is set to '1.15' or later. * Python '3.5' is available when `runtime_version` is set to a version from '1.4' to '1.14'. * Python '2.7' is available when `runtime_version` is set to '1.15' or earlier. Read more about the Python versions available for [each runtime version](/ml-engine/docs/runtime-version-list).
func (o LookupVersionResultOutput) PythonVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.PythonVersion }).(pulumi.StringOutput)
}

// Optional. *Only* specify this field in a projects.models.versions.patch request. Specifying it in a projects.models.versions.create request has no effect. Configures the request-response pair logging on predictions from this Version.
func (o LookupVersionResultOutput) RequestLoggingConfig() GoogleCloudMlV1__RequestLoggingConfigResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) GoogleCloudMlV1__RequestLoggingConfigResponse {
		return v.RequestLoggingConfig
	}).(GoogleCloudMlV1__RequestLoggingConfigResponseOutput)
}

// Optional. Specifies paths on a custom container's HTTP server where AI Platform Prediction sends certain requests. If you specify this field, then you must also specify the `container` field. If you specify the `container` field and do not specify this field, it defaults to the following: ``` json { "predict": "/v1/models/MODEL/versions/VERSION:predict", "health": "/v1/models/MODEL/versions/VERSION" }  ``` See RouteMap for more details about these default values.
func (o LookupVersionResultOutput) Routes() GoogleCloudMlV1__RouteMapResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) GoogleCloudMlV1__RouteMapResponse { return v.Routes }).(GoogleCloudMlV1__RouteMapResponseOutput)
}

// The AI Platform runtime version to use for this deployment. For more information, see the [runtime version list](/ml-engine/docs/runtime-version-list) and [how to manage runtime versions](/ml-engine/docs/versioning).
func (o LookupVersionResultOutput) RuntimeVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.RuntimeVersion }).(pulumi.StringOutput)
}

// Optional. Specifies the service account for resource access control. If you specify this field, then you must also specify either the `containerSpec` or the `predictionClass` field. Learn more about [using a custom service account](/ai-platform/prediction/docs/custom-service-account).
func (o LookupVersionResultOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.ServiceAccount }).(pulumi.StringOutput)
}

// The state of a version.
func (o LookupVersionResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVersionResultOutput{})
}
