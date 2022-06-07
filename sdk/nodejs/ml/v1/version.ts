// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new version of a model from a trained TensorFlow model. If the version created in the cloud by this call is the first deployed version of the specified model, it will be made the default version of the model. When you add a version to a model that already has one or more versions, the default version does not automatically change. If you want a new version to be the default, you must call projects.models.versions.setDefault.
 */
export class Version extends pulumi.CustomResource {
    /**
     * Get an existing Version resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Version {
        return new Version(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:ml/v1:Version';

    /**
     * Returns true if the given object is an instance of Version.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Version {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Version.__pulumiType;
    }

    /**
     * Optional. Accelerator config for using GPUs for online prediction (beta). Only specify this field if you have specified a Compute Engine (N1) machine type in the `machineType` field. Learn more about [using GPUs for online prediction](/ml-engine/docs/machine-types-online-prediction#gpus).
     */
    public readonly acceleratorConfig!: pulumi.Output<outputs.ml.v1.GoogleCloudMlV1__AcceleratorConfigResponse>;
    /**
     * Automatically scale the number of nodes used to serve the model in response to increases and decreases in traffic. Care should be taken to ramp up traffic according to the model's ability to scale or you will start seeing increases in latency and 429 response codes.
     */
    public readonly autoScaling!: pulumi.Output<outputs.ml.v1.GoogleCloudMlV1__AutoScalingResponse>;
    /**
     * Optional. Specifies a custom container to use for serving predictions. If you specify this field, then `machineType` is required. If you specify this field, then `deploymentUri` is optional. If you specify this field, then you must not specify `runtimeVersion`, `packageUris`, `framework`, `pythonVersion`, or `predictionClass`.
     */
    public readonly container!: pulumi.Output<outputs.ml.v1.GoogleCloudMlV1__ContainerSpecResponse>;
    /**
     * The time the version was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The Cloud Storage URI of a directory containing trained model artifacts to be used to create the model version. See the [guide to deploying models](/ai-platform/prediction/docs/deploying-models) for more information. The total number of files under this directory must not exceed 1000. During projects.models.versions.create, AI Platform Prediction copies all files from the specified directory to a location managed by the service. From then on, AI Platform Prediction uses these copies of the model artifacts to serve predictions, not the original files in Cloud Storage, so this location is useful only as a historical record. If you specify container, then this field is optional. Otherwise, it is required. Learn [how to use this field with a custom container](/ai-platform/prediction/docs/custom-container-requirements#artifacts).
     */
    public readonly deploymentUri!: pulumi.Output<string>;
    /**
     * Optional. The description specified for the version when it was created.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The details of a failure or a cancellation.
     */
    public /*out*/ readonly errorMessage!: pulumi.Output<string>;
    /**
     * `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a model from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform model updates in order to avoid race conditions: An `etag` is returned in the response to `GetVersion`, and systems are expected to put that etag in the request to `UpdateVersion` to ensure that their change will be applied to the model as intended.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * Optional. Configures explainability features on the model's version. Some explanation features require additional metadata to be loaded as part of the model payload.
     */
    public readonly explanationConfig!: pulumi.Output<outputs.ml.v1.GoogleCloudMlV1__ExplanationConfigResponse>;
    /**
     * Optional. The machine learning framework AI Platform uses to train this version of the model. Valid values are `TENSORFLOW`, `SCIKIT_LEARN`, `XGBOOST`. If you do not specify a framework, AI Platform will analyze files in the deployment_uri to determine a framework. If you choose `SCIKIT_LEARN` or `XGBOOST`, you must also set the runtime version of the model to 1.4 or greater. Do **not** specify a framework if you're deploying a [custom prediction routine](/ai-platform/prediction/docs/custom-prediction-routines) or if you're using a [custom container](/ai-platform/prediction/docs/use-custom-container).
     */
    public readonly framework!: pulumi.Output<string>;
    /**
     * If true, this version will be used to handle prediction requests that do not specify a version. You can change the default version by calling projects.methods.versions.setDefault.
     */
    public /*out*/ readonly isDefault!: pulumi.Output<boolean>;
    /**
     * Optional. One or more labels that you can add, to organize your model versions. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels. Note that this field is not updatable for mls1* models.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The [AI Platform (Unified) `Model`](https://cloud.google.com/ai-platform-unified/docs/reference/rest/v1beta1/projects.locations.models) ID for the last [model migration](https://cloud.google.com/ai-platform-unified/docs/start/migrating-to-ai-platform-unified).
     */
    public /*out*/ readonly lastMigrationModelId!: pulumi.Output<string>;
    /**
     * The last time this version was successfully [migrated to AI Platform (Unified)](https://cloud.google.com/ai-platform-unified/docs/start/migrating-to-ai-platform-unified).
     */
    public /*out*/ readonly lastMigrationTime!: pulumi.Output<string>;
    /**
     * The time the version was last used for prediction.
     */
    public /*out*/ readonly lastUseTime!: pulumi.Output<string>;
    /**
     * Optional. The type of machine on which to serve the model. Currently only applies to online prediction service. To learn about valid values for this field, read [Choosing a machine type for online prediction](/ai-platform/prediction/docs/machine-types-online-prediction). If this field is not specified and you are using a [regional endpoint](/ai-platform/prediction/docs/regional-endpoints), then the machine type defaults to `n1-standard-2`. If this field is not specified and you are using the global endpoint (`ml.googleapis.com`), then the machine type defaults to `mls1-c1-m2`.
     */
    public readonly machineType!: pulumi.Output<string>;
    /**
     * Manually select the number of nodes to use for serving the model. You should generally use `auto_scaling` with an appropriate `min_nodes` instead, but this option is available if you want more predictable billing. Beware that latency and error rates will increase if the traffic exceeds that capability of the system to serve it based on the selected number of nodes.
     */
    public readonly manualScaling!: pulumi.Output<outputs.ml.v1.GoogleCloudMlV1__ManualScalingResponse>;
    /**
     * The name specified for the version when it was created. The version name must be unique within the model it is created in.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Optional. Cloud Storage paths (`gs://…`) of packages for [custom prediction routines](/ml-engine/docs/tensorflow/custom-prediction-routines) or [scikit-learn pipelines with custom code](/ml-engine/docs/scikit/exporting-for-prediction#custom-pipeline-code). For a custom prediction routine, one of these packages must contain your Predictor class (see [`predictionClass`](#Version.FIELDS.prediction_class)). Additionally, include any dependencies used by your Predictor or scikit-learn pipeline uses that are not already included in your selected [runtime version](/ml-engine/docs/tensorflow/runtime-version-list). If you specify this field, you must also set [`runtimeVersion`](#Version.FIELDS.runtime_version) to 1.4 or greater.
     */
    public readonly packageUris!: pulumi.Output<string[]>;
    /**
     * Optional. The fully qualified name (module_name.class_name) of a class that implements the Predictor interface described in this reference field. The module containing this class should be included in a package provided to the [`packageUris` field](#Version.FIELDS.package_uris). Specify this field if and only if you are deploying a [custom prediction routine (beta)](/ml-engine/docs/tensorflow/custom-prediction-routines). If you specify this field, you must set [`runtimeVersion`](#Version.FIELDS.runtime_version) to 1.4 or greater and you must set `machineType` to a [legacy (MLS1) machine type](/ml-engine/docs/machine-types-online-prediction). The following code sample provides the Predictor interface: class Predictor(object): """Interface for constructing custom predictors.""" def predict(self, instances, **kwargs): """Performs custom prediction. Instances are the decoded values from the request. They have already been deserialized from JSON. Args: instances: A list of prediction input instances. **kwargs: A dictionary of keyword args provided as additional fields on the predict request body. Returns: A list of outputs containing the prediction results. This list must be JSON serializable. """ raise NotImplementedError() @classmethod def from_path(cls, model_dir): """Creates an instance of Predictor using the given path. Loading of the predictor should be done in this method. Args: model_dir: The local directory that contains the exported model file along with any additional files uploaded when creating the version resource. Returns: An instance implementing this Predictor class. """ raise NotImplementedError() Learn more about [the Predictor interface and custom prediction routines](/ml-engine/docs/tensorflow/custom-prediction-routines).
     */
    public readonly predictionClass!: pulumi.Output<string>;
    /**
     * The version of Python used in prediction. The following Python versions are available: * Python '3.7' is available when `runtime_version` is set to '1.15' or later. * Python '3.5' is available when `runtime_version` is set to a version from '1.4' to '1.14'. * Python '2.7' is available when `runtime_version` is set to '1.15' or earlier. Read more about the Python versions available for [each runtime version](/ml-engine/docs/runtime-version-list).
     */
    public readonly pythonVersion!: pulumi.Output<string>;
    /**
     * Optional. *Only* specify this field in a projects.models.versions.patch request. Specifying it in a projects.models.versions.create request has no effect. Configures the request-response pair logging on predictions from this Version.
     */
    public readonly requestLoggingConfig!: pulumi.Output<outputs.ml.v1.GoogleCloudMlV1__RequestLoggingConfigResponse>;
    /**
     * Optional. Specifies paths on a custom container's HTTP server where AI Platform Prediction sends certain requests. If you specify this field, then you must also specify the `container` field. If you specify the `container` field and do not specify this field, it defaults to the following: ```json { "predict": "/v1/models/MODEL/versions/VERSION:predict", "health": "/v1/models/MODEL/versions/VERSION" } ``` See RouteMap for more details about these default values.
     */
    public readonly routes!: pulumi.Output<outputs.ml.v1.GoogleCloudMlV1__RouteMapResponse>;
    /**
     * The AI Platform runtime version to use for this deployment. For more information, see the [runtime version list](/ml-engine/docs/runtime-version-list) and [how to manage runtime versions](/ml-engine/docs/versioning).
     */
    public readonly runtimeVersion!: pulumi.Output<string>;
    /**
     * Optional. Specifies the service account for resource access control. If you specify this field, then you must also specify either the `containerSpec` or the `predictionClass` field. Learn more about [using a custom service account](/ai-platform/prediction/docs/custom-service-account).
     */
    public readonly serviceAccount!: pulumi.Output<string>;
    /**
     * The state of a version.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a Version resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VersionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.modelId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'modelId'");
            }
            if ((!args || args.pythonVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pythonVersion'");
            }
            if ((!args || args.runtimeVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runtimeVersion'");
            }
            resourceInputs["acceleratorConfig"] = args ? args.acceleratorConfig : undefined;
            resourceInputs["autoScaling"] = args ? args.autoScaling : undefined;
            resourceInputs["container"] = args ? args.container : undefined;
            resourceInputs["deploymentUri"] = args ? args.deploymentUri : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["explanationConfig"] = args ? args.explanationConfig : undefined;
            resourceInputs["framework"] = args ? args.framework : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["machineType"] = args ? args.machineType : undefined;
            resourceInputs["manualScaling"] = args ? args.manualScaling : undefined;
            resourceInputs["modelId"] = args ? args.modelId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["packageUris"] = args ? args.packageUris : undefined;
            resourceInputs["predictionClass"] = args ? args.predictionClass : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["pythonVersion"] = args ? args.pythonVersion : undefined;
            resourceInputs["requestLoggingConfig"] = args ? args.requestLoggingConfig : undefined;
            resourceInputs["routes"] = args ? args.routes : undefined;
            resourceInputs["runtimeVersion"] = args ? args.runtimeVersion : undefined;
            resourceInputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["errorMessage"] = undefined /*out*/;
            resourceInputs["isDefault"] = undefined /*out*/;
            resourceInputs["lastMigrationModelId"] = undefined /*out*/;
            resourceInputs["lastMigrationTime"] = undefined /*out*/;
            resourceInputs["lastUseTime"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["acceleratorConfig"] = undefined /*out*/;
            resourceInputs["autoScaling"] = undefined /*out*/;
            resourceInputs["container"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deploymentUri"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["errorMessage"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["explanationConfig"] = undefined /*out*/;
            resourceInputs["framework"] = undefined /*out*/;
            resourceInputs["isDefault"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["lastMigrationModelId"] = undefined /*out*/;
            resourceInputs["lastMigrationTime"] = undefined /*out*/;
            resourceInputs["lastUseTime"] = undefined /*out*/;
            resourceInputs["machineType"] = undefined /*out*/;
            resourceInputs["manualScaling"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["packageUris"] = undefined /*out*/;
            resourceInputs["predictionClass"] = undefined /*out*/;
            resourceInputs["pythonVersion"] = undefined /*out*/;
            resourceInputs["requestLoggingConfig"] = undefined /*out*/;
            resourceInputs["routes"] = undefined /*out*/;
            resourceInputs["runtimeVersion"] = undefined /*out*/;
            resourceInputs["serviceAccount"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Version.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Version resource.
 */
export interface VersionArgs {
    /**
     * Optional. Accelerator config for using GPUs for online prediction (beta). Only specify this field if you have specified a Compute Engine (N1) machine type in the `machineType` field. Learn more about [using GPUs for online prediction](/ml-engine/docs/machine-types-online-prediction#gpus).
     */
    acceleratorConfig?: pulumi.Input<inputs.ml.v1.GoogleCloudMlV1__AcceleratorConfigArgs>;
    /**
     * Automatically scale the number of nodes used to serve the model in response to increases and decreases in traffic. Care should be taken to ramp up traffic according to the model's ability to scale or you will start seeing increases in latency and 429 response codes.
     */
    autoScaling?: pulumi.Input<inputs.ml.v1.GoogleCloudMlV1__AutoScalingArgs>;
    /**
     * Optional. Specifies a custom container to use for serving predictions. If you specify this field, then `machineType` is required. If you specify this field, then `deploymentUri` is optional. If you specify this field, then you must not specify `runtimeVersion`, `packageUris`, `framework`, `pythonVersion`, or `predictionClass`.
     */
    container?: pulumi.Input<inputs.ml.v1.GoogleCloudMlV1__ContainerSpecArgs>;
    /**
     * The Cloud Storage URI of a directory containing trained model artifacts to be used to create the model version. See the [guide to deploying models](/ai-platform/prediction/docs/deploying-models) for more information. The total number of files under this directory must not exceed 1000. During projects.models.versions.create, AI Platform Prediction copies all files from the specified directory to a location managed by the service. From then on, AI Platform Prediction uses these copies of the model artifacts to serve predictions, not the original files in Cloud Storage, so this location is useful only as a historical record. If you specify container, then this field is optional. Otherwise, it is required. Learn [how to use this field with a custom container](/ai-platform/prediction/docs/custom-container-requirements#artifacts).
     */
    deploymentUri?: pulumi.Input<string>;
    /**
     * Optional. The description specified for the version when it was created.
     */
    description?: pulumi.Input<string>;
    /**
     * `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a model from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform model updates in order to avoid race conditions: An `etag` is returned in the response to `GetVersion`, and systems are expected to put that etag in the request to `UpdateVersion` to ensure that their change will be applied to the model as intended.
     */
    etag?: pulumi.Input<string>;
    /**
     * Optional. Configures explainability features on the model's version. Some explanation features require additional metadata to be loaded as part of the model payload.
     */
    explanationConfig?: pulumi.Input<inputs.ml.v1.GoogleCloudMlV1__ExplanationConfigArgs>;
    /**
     * Optional. The machine learning framework AI Platform uses to train this version of the model. Valid values are `TENSORFLOW`, `SCIKIT_LEARN`, `XGBOOST`. If you do not specify a framework, AI Platform will analyze files in the deployment_uri to determine a framework. If you choose `SCIKIT_LEARN` or `XGBOOST`, you must also set the runtime version of the model to 1.4 or greater. Do **not** specify a framework if you're deploying a [custom prediction routine](/ai-platform/prediction/docs/custom-prediction-routines) or if you're using a [custom container](/ai-platform/prediction/docs/use-custom-container).
     */
    framework?: pulumi.Input<enums.ml.v1.VersionFramework>;
    /**
     * Optional. One or more labels that you can add, to organize your model versions. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels. Note that this field is not updatable for mls1* models.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional. The type of machine on which to serve the model. Currently only applies to online prediction service. To learn about valid values for this field, read [Choosing a machine type for online prediction](/ai-platform/prediction/docs/machine-types-online-prediction). If this field is not specified and you are using a [regional endpoint](/ai-platform/prediction/docs/regional-endpoints), then the machine type defaults to `n1-standard-2`. If this field is not specified and you are using the global endpoint (`ml.googleapis.com`), then the machine type defaults to `mls1-c1-m2`.
     */
    machineType?: pulumi.Input<string>;
    /**
     * Manually select the number of nodes to use for serving the model. You should generally use `auto_scaling` with an appropriate `min_nodes` instead, but this option is available if you want more predictable billing. Beware that latency and error rates will increase if the traffic exceeds that capability of the system to serve it based on the selected number of nodes.
     */
    manualScaling?: pulumi.Input<inputs.ml.v1.GoogleCloudMlV1__ManualScalingArgs>;
    modelId: pulumi.Input<string>;
    /**
     * The name specified for the version when it was created. The version name must be unique within the model it is created in.
     */
    name?: pulumi.Input<string>;
    /**
     * Optional. Cloud Storage paths (`gs://…`) of packages for [custom prediction routines](/ml-engine/docs/tensorflow/custom-prediction-routines) or [scikit-learn pipelines with custom code](/ml-engine/docs/scikit/exporting-for-prediction#custom-pipeline-code). For a custom prediction routine, one of these packages must contain your Predictor class (see [`predictionClass`](#Version.FIELDS.prediction_class)). Additionally, include any dependencies used by your Predictor or scikit-learn pipeline uses that are not already included in your selected [runtime version](/ml-engine/docs/tensorflow/runtime-version-list). If you specify this field, you must also set [`runtimeVersion`](#Version.FIELDS.runtime_version) to 1.4 or greater.
     */
    packageUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional. The fully qualified name (module_name.class_name) of a class that implements the Predictor interface described in this reference field. The module containing this class should be included in a package provided to the [`packageUris` field](#Version.FIELDS.package_uris). Specify this field if and only if you are deploying a [custom prediction routine (beta)](/ml-engine/docs/tensorflow/custom-prediction-routines). If you specify this field, you must set [`runtimeVersion`](#Version.FIELDS.runtime_version) to 1.4 or greater and you must set `machineType` to a [legacy (MLS1) machine type](/ml-engine/docs/machine-types-online-prediction). The following code sample provides the Predictor interface: class Predictor(object): """Interface for constructing custom predictors.""" def predict(self, instances, **kwargs): """Performs custom prediction. Instances are the decoded values from the request. They have already been deserialized from JSON. Args: instances: A list of prediction input instances. **kwargs: A dictionary of keyword args provided as additional fields on the predict request body. Returns: A list of outputs containing the prediction results. This list must be JSON serializable. """ raise NotImplementedError() @classmethod def from_path(cls, model_dir): """Creates an instance of Predictor using the given path. Loading of the predictor should be done in this method. Args: model_dir: The local directory that contains the exported model file along with any additional files uploaded when creating the version resource. Returns: An instance implementing this Predictor class. """ raise NotImplementedError() Learn more about [the Predictor interface and custom prediction routines](/ml-engine/docs/tensorflow/custom-prediction-routines).
     */
    predictionClass?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The version of Python used in prediction. The following Python versions are available: * Python '3.7' is available when `runtime_version` is set to '1.15' or later. * Python '3.5' is available when `runtime_version` is set to a version from '1.4' to '1.14'. * Python '2.7' is available when `runtime_version` is set to '1.15' or earlier. Read more about the Python versions available for [each runtime version](/ml-engine/docs/runtime-version-list).
     */
    pythonVersion: pulumi.Input<string>;
    /**
     * Optional. *Only* specify this field in a projects.models.versions.patch request. Specifying it in a projects.models.versions.create request has no effect. Configures the request-response pair logging on predictions from this Version.
     */
    requestLoggingConfig?: pulumi.Input<inputs.ml.v1.GoogleCloudMlV1__RequestLoggingConfigArgs>;
    /**
     * Optional. Specifies paths on a custom container's HTTP server where AI Platform Prediction sends certain requests. If you specify this field, then you must also specify the `container` field. If you specify the `container` field and do not specify this field, it defaults to the following: ```json { "predict": "/v1/models/MODEL/versions/VERSION:predict", "health": "/v1/models/MODEL/versions/VERSION" } ``` See RouteMap for more details about these default values.
     */
    routes?: pulumi.Input<inputs.ml.v1.GoogleCloudMlV1__RouteMapArgs>;
    /**
     * The AI Platform runtime version to use for this deployment. For more information, see the [runtime version list](/ml-engine/docs/runtime-version-list) and [how to manage runtime versions](/ml-engine/docs/versioning).
     */
    runtimeVersion: pulumi.Input<string>;
    /**
     * Optional. Specifies the service account for resource access control. If you specify this field, then you must also specify either the `containerSpec` or the `predictionClass` field. Learn more about [using a custom service account](/ai-platform/prediction/docs/custom-service-account).
     */
    serviceAccount?: pulumi.Input<string>;
}
