# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ArtifactState',
    'ExecutionState',
    'FeatureGroupFeatureValueType',
    'FeatureStoreFeatureValueType',
    'GoogleCloudAiplatformV1beta1ExamplesExampleGcsSourceDataFormat',
    'GoogleCloudAiplatformV1beta1ExplanationMetadataInputMetadataEncoding',
    'GoogleCloudAiplatformV1beta1ExplanationMetadataInputMetadataVisualizationColorMap',
    'GoogleCloudAiplatformV1beta1ExplanationMetadataInputMetadataVisualizationOverlayType',
    'GoogleCloudAiplatformV1beta1ExplanationMetadataInputMetadataVisualizationPolarity',
    'GoogleCloudAiplatformV1beta1ExplanationMetadataInputMetadataVisualizationType',
    'GoogleCloudAiplatformV1beta1FeatureViewVectorSearchConfigDistanceMeasureType',
    'GoogleCloudAiplatformV1beta1FeaturestoreMonitoringConfigImportFeaturesAnalysisAnomalyDetectionBaseline',
    'GoogleCloudAiplatformV1beta1FeaturestoreMonitoringConfigImportFeaturesAnalysisState',
    'GoogleCloudAiplatformV1beta1MachineSpecAcceleratorType',
    'GoogleCloudAiplatformV1beta1ModelMonitoringObjectiveConfigExplanationConfigExplanationBaselinePredictionFormat',
    'GoogleCloudAiplatformV1beta1ModelMonitoringStatsAnomaliesObjective',
    'GoogleCloudAiplatformV1beta1NasJobSpecMultiTrialAlgorithmSpecMetricSpecGoal',
    'GoogleCloudAiplatformV1beta1NasJobSpecMultiTrialAlgorithmSpecMultiTrialAlgorithm',
    'GoogleCloudAiplatformV1beta1PipelineJobRuntimeConfigFailurePolicy',
    'GoogleCloudAiplatformV1beta1PresetsModality',
    'GoogleCloudAiplatformV1beta1PresetsQuery',
    'GoogleCloudAiplatformV1beta1SampleConfigSampleStrategy',
    'GoogleCloudAiplatformV1beta1StudySpecAlgorithm',
    'GoogleCloudAiplatformV1beta1StudySpecMeasurementSelectionType',
    'GoogleCloudAiplatformV1beta1StudySpecMetricSpecGoal',
    'GoogleCloudAiplatformV1beta1StudySpecObservationNoise',
    'GoogleCloudAiplatformV1beta1StudySpecParameterSpecScaleType',
    'IndexIndexUpdateMethod',
    'MetadataSchemaSchemaType',
    'NotebookRuntimeTemplateNotebookRuntimeType',
    'TimeSeriesValueType',
]


class ArtifactState(str, Enum):
    """
    The state of this Artifact. This is a property of the Artifact, and does not imply or capture any ongoing process. This property is managed by clients (such as Vertex AI Pipelines), and the system does not prescribe or check the validity of state transitions.
    """
    STATE_UNSPECIFIED = "STATE_UNSPECIFIED"
    """
    Unspecified state for the Artifact.
    """
    PENDING = "PENDING"
    """
    A state used by systems like Vertex AI Pipelines to indicate that the underlying data item represented by this Artifact is being created.
    """
    LIVE = "LIVE"
    """
    A state indicating that the Artifact should exist, unless something external to the system deletes it.
    """


class ExecutionState(str, Enum):
    """
    The state of this Execution. This is a property of the Execution, and does not imply or capture any ongoing process. This property is managed by clients (such as Vertex AI Pipelines) and the system does not prescribe or check the validity of state transitions.
    """
    STATE_UNSPECIFIED = "STATE_UNSPECIFIED"
    """
    Unspecified Execution state
    """
    NEW = "NEW"
    """
    The Execution is new
    """
    RUNNING = "RUNNING"
    """
    The Execution is running
    """
    COMPLETE = "COMPLETE"
    """
    The Execution has finished running
    """
    FAILED = "FAILED"
    """
    The Execution has failed
    """
    CACHED = "CACHED"
    """
    The Execution completed through Cache hit.
    """
    CANCELLED = "CANCELLED"
    """
    The Execution was cancelled.
    """


class FeatureGroupFeatureValueType(str, Enum):
    """
    Immutable. Only applicable for Vertex AI Feature Store (Legacy). Type of Feature value.
    """
    VALUE_TYPE_UNSPECIFIED = "VALUE_TYPE_UNSPECIFIED"
    """
    The value type is unspecified.
    """
    BOOL = "BOOL"
    """
    Used for Feature that is a boolean.
    """
    BOOL_ARRAY = "BOOL_ARRAY"
    """
    Used for Feature that is a list of boolean.
    """
    DOUBLE = "DOUBLE"
    """
    Used for Feature that is double.
    """
    DOUBLE_ARRAY = "DOUBLE_ARRAY"
    """
    Used for Feature that is a list of double.
    """
    INT64 = "INT64"
    """
    Used for Feature that is INT64.
    """
    INT64_ARRAY = "INT64_ARRAY"
    """
    Used for Feature that is a list of INT64.
    """
    STRING = "STRING"
    """
    Used for Feature that is string.
    """
    STRING_ARRAY = "STRING_ARRAY"
    """
    Used for Feature that is a list of String.
    """
    BYTES = "BYTES"
    """
    Used for Feature that is bytes.
    """


class FeatureStoreFeatureValueType(str, Enum):
    """
    Immutable. Only applicable for Vertex AI Feature Store (Legacy). Type of Feature value.
    """
    VALUE_TYPE_UNSPECIFIED = "VALUE_TYPE_UNSPECIFIED"
    """
    The value type is unspecified.
    """
    BOOL = "BOOL"
    """
    Used for Feature that is a boolean.
    """
    BOOL_ARRAY = "BOOL_ARRAY"
    """
    Used for Feature that is a list of boolean.
    """
    DOUBLE = "DOUBLE"
    """
    Used for Feature that is double.
    """
    DOUBLE_ARRAY = "DOUBLE_ARRAY"
    """
    Used for Feature that is a list of double.
    """
    INT64 = "INT64"
    """
    Used for Feature that is INT64.
    """
    INT64_ARRAY = "INT64_ARRAY"
    """
    Used for Feature that is a list of INT64.
    """
    STRING = "STRING"
    """
    Used for Feature that is string.
    """
    STRING_ARRAY = "STRING_ARRAY"
    """
    Used for Feature that is a list of String.
    """
    BYTES = "BYTES"
    """
    Used for Feature that is bytes.
    """


class GoogleCloudAiplatformV1beta1ExamplesExampleGcsSourceDataFormat(str, Enum):
    """
    The format in which instances are given, if not specified, assume it's JSONL format. Currently only JSONL format is supported.
    """
    DATA_FORMAT_UNSPECIFIED = "DATA_FORMAT_UNSPECIFIED"
    """
    Format unspecified, used when unset.
    """
    JSONL = "JSONL"
    """
    Examples are stored in JSONL files.
    """


class GoogleCloudAiplatformV1beta1ExplanationMetadataInputMetadataEncoding(str, Enum):
    """
    Defines how the feature is encoded into the input tensor. Defaults to IDENTITY.
    """
    ENCODING_UNSPECIFIED = "ENCODING_UNSPECIFIED"
    """
    Default value. This is the same as IDENTITY.
    """
    IDENTITY = "IDENTITY"
    """
    The tensor represents one feature.
    """
    BAG_OF_FEATURES = "BAG_OF_FEATURES"
    """
    The tensor represents a bag of features where each index maps to a feature. InputMetadata.index_feature_mapping must be provided for this encoding. For example: ``` input = [27, 6.0, 150] index_feature_mapping = ["age", "height", "weight"] ```
    """
    BAG_OF_FEATURES_SPARSE = "BAG_OF_FEATURES_SPARSE"
    """
    The tensor represents a bag of features where each index maps to a feature. Zero values in the tensor indicates feature being non-existent. InputMetadata.index_feature_mapping must be provided for this encoding. For example: ``` input = [2, 0, 5, 0, 1] index_feature_mapping = ["a", "b", "c", "d", "e"] ```
    """
    INDICATOR = "INDICATOR"
    """
    The tensor is a list of binaries representing whether a feature exists or not (1 indicates existence). InputMetadata.index_feature_mapping must be provided for this encoding. For example: ``` input = [1, 0, 1, 0, 1] index_feature_mapping = ["a", "b", "c", "d", "e"] ```
    """
    COMBINED_EMBEDDING = "COMBINED_EMBEDDING"
    """
    The tensor is encoded into a 1-dimensional array represented by an encoded tensor. InputMetadata.encoded_tensor_name must be provided for this encoding. For example: ``` input = ["This", "is", "a", "test", "."] encoded = [0.1, 0.2, 0.3, 0.4, 0.5] ```
    """
    CONCAT_EMBEDDING = "CONCAT_EMBEDDING"
    """
    Select this encoding when the input tensor is encoded into a 2-dimensional array represented by an encoded tensor. InputMetadata.encoded_tensor_name must be provided for this encoding. The first dimension of the encoded tensor's shape is the same as the input tensor's shape. For example: ``` input = ["This", "is", "a", "test", "."] encoded = [[0.1, 0.2, 0.3, 0.4, 0.5], [0.2, 0.1, 0.4, 0.3, 0.5], [0.5, 0.1, 0.3, 0.5, 0.4], [0.5, 0.3, 0.1, 0.2, 0.4], [0.4, 0.3, 0.2, 0.5, 0.1]] ```
    """


class GoogleCloudAiplatformV1beta1ExplanationMetadataInputMetadataVisualizationColorMap(str, Enum):
    """
    The color scheme used for the highlighted areas. Defaults to PINK_GREEN for Integrated Gradients attribution, which shows positive attributions in green and negative in pink. Defaults to VIRIDIS for XRAI attribution, which highlights the most influential regions in yellow and the least influential in blue.
    """
    COLOR_MAP_UNSPECIFIED = "COLOR_MAP_UNSPECIFIED"
    """
    Should not be used.
    """
    PINK_GREEN = "PINK_GREEN"
    """
    Positive: green. Negative: pink.
    """
    VIRIDIS = "VIRIDIS"
    """
    Viridis color map: A perceptually uniform color mapping which is easier to see by those with colorblindness and progresses from yellow to green to blue. Positive: yellow. Negative: blue.
    """
    RED = "RED"
    """
    Positive: red. Negative: red.
    """
    GREEN = "GREEN"
    """
    Positive: green. Negative: green.
    """
    RED_GREEN = "RED_GREEN"
    """
    Positive: green. Negative: red.
    """
    PINK_WHITE_GREEN = "PINK_WHITE_GREEN"
    """
    PiYG palette.
    """


class GoogleCloudAiplatformV1beta1ExplanationMetadataInputMetadataVisualizationOverlayType(str, Enum):
    """
    How the original image is displayed in the visualization. Adjusting the overlay can help increase visual clarity if the original image makes it difficult to view the visualization. Defaults to NONE.
    """
    OVERLAY_TYPE_UNSPECIFIED = "OVERLAY_TYPE_UNSPECIFIED"
    """
    Default value. This is the same as NONE.
    """
    NONE = "NONE"
    """
    No overlay.
    """
    ORIGINAL = "ORIGINAL"
    """
    The attributions are shown on top of the original image.
    """
    GRAYSCALE = "GRAYSCALE"
    """
    The attributions are shown on top of grayscaled version of the original image.
    """
    MASK_BLACK = "MASK_BLACK"
    """
    The attributions are used as a mask to reveal predictive parts of the image and hide the un-predictive parts.
    """


class GoogleCloudAiplatformV1beta1ExplanationMetadataInputMetadataVisualizationPolarity(str, Enum):
    """
    Whether to only highlight pixels with positive contributions, negative or both. Defaults to POSITIVE.
    """
    POLARITY_UNSPECIFIED = "POLARITY_UNSPECIFIED"
    """
    Default value. This is the same as POSITIVE.
    """
    POSITIVE = "POSITIVE"
    """
    Highlights the pixels/outlines that were most influential to the model's prediction.
    """
    NEGATIVE = "NEGATIVE"
    """
    Setting polarity to negative highlights areas that does not lead to the models's current prediction.
    """
    BOTH = "BOTH"
    """
    Shows both positive and negative attributions.
    """


class GoogleCloudAiplatformV1beta1ExplanationMetadataInputMetadataVisualizationType(str, Enum):
    """
    Type of the image visualization. Only applicable to Integrated Gradients attribution. OUTLINES shows regions of attribution, while PIXELS shows per-pixel attribution. Defaults to OUTLINES.
    """
    TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED"
    """
    Should not be used.
    """
    PIXELS = "PIXELS"
    """
    Shows which pixel contributed to the image prediction.
    """
    OUTLINES = "OUTLINES"
    """
    Shows which region contributed to the image prediction by outlining the region.
    """


class GoogleCloudAiplatformV1beta1FeatureViewVectorSearchConfigDistanceMeasureType(str, Enum):
    """
    Optional. The distance measure used in nearest neighbor search.
    """
    DISTANCE_MEASURE_TYPE_UNSPECIFIED = "DISTANCE_MEASURE_TYPE_UNSPECIFIED"
    """
    Should not be set.
    """
    SQUARED_L2_DISTANCE = "SQUARED_L2_DISTANCE"
    """
    Euclidean (L_2) Distance.
    """
    COSINE_DISTANCE = "COSINE_DISTANCE"
    """
    Cosine Distance. Defined as 1 - cosine similarity. We strongly suggest using DOT_PRODUCT_DISTANCE + UNIT_L2_NORM instead of COSINE distance. Our algorithms have been more optimized for DOT_PRODUCT distance which, when combined with UNIT_L2_NORM, is mathematically equivalent to COSINE distance and results in the same ranking.
    """
    DOT_PRODUCT_DISTANCE = "DOT_PRODUCT_DISTANCE"
    """
    Dot Product Distance. Defined as a negative of the dot product.
    """


class GoogleCloudAiplatformV1beta1FeaturestoreMonitoringConfigImportFeaturesAnalysisAnomalyDetectionBaseline(str, Enum):
    """
    The baseline used to do anomaly detection for the statistics generated by import features analysis.
    """
    BASELINE_UNSPECIFIED = "BASELINE_UNSPECIFIED"
    """
    Should not be used.
    """
    LATEST_STATS = "LATEST_STATS"
    """
    Choose the later one statistics generated by either most recent snapshot analysis or previous import features analysis. If non of them exists, skip anomaly detection and only generate a statistics.
    """
    MOST_RECENT_SNAPSHOT_STATS = "MOST_RECENT_SNAPSHOT_STATS"
    """
    Use the statistics generated by the most recent snapshot analysis if exists.
    """
    PREVIOUS_IMPORT_FEATURES_STATS = "PREVIOUS_IMPORT_FEATURES_STATS"
    """
    Use the statistics generated by the previous import features analysis if exists.
    """


class GoogleCloudAiplatformV1beta1FeaturestoreMonitoringConfigImportFeaturesAnalysisState(str, Enum):
    """
    Whether to enable / disable / inherite default hebavior for import features analysis.
    """
    STATE_UNSPECIFIED = "STATE_UNSPECIFIED"
    """
    Should not be used.
    """
    DEFAULT = "DEFAULT"
    """
    The default behavior of whether to enable the monitoring. EntityType-level config: disabled. Feature-level config: inherited from the configuration of EntityType this Feature belongs to.
    """
    ENABLED = "ENABLED"
    """
    Explicitly enables import features analysis. EntityType-level config: by default enables import features analysis for all Features under it. Feature-level config: enables import features analysis regardless of the EntityType-level config.
    """
    DISABLED = "DISABLED"
    """
    Explicitly disables import features analysis. EntityType-level config: by default disables import features analysis for all Features under it. Feature-level config: disables import features analysis regardless of the EntityType-level config.
    """


class GoogleCloudAiplatformV1beta1MachineSpecAcceleratorType(str, Enum):
    """
    Immutable. The type of accelerator(s) that may be attached to the machine as per accelerator_count.
    """
    ACCELERATOR_TYPE_UNSPECIFIED = "ACCELERATOR_TYPE_UNSPECIFIED"
    """
    Unspecified accelerator type, which means no accelerator.
    """
    NVIDIA_TESLA_K80 = "NVIDIA_TESLA_K80"
    """
    Nvidia Tesla K80 GPU.
    """
    NVIDIA_TESLA_P100 = "NVIDIA_TESLA_P100"
    """
    Nvidia Tesla P100 GPU.
    """
    NVIDIA_TESLA_V100 = "NVIDIA_TESLA_V100"
    """
    Nvidia Tesla V100 GPU.
    """
    NVIDIA_TESLA_P4 = "NVIDIA_TESLA_P4"
    """
    Nvidia Tesla P4 GPU.
    """
    NVIDIA_TESLA_T4 = "NVIDIA_TESLA_T4"
    """
    Nvidia Tesla T4 GPU.
    """
    NVIDIA_TESLA_A100 = "NVIDIA_TESLA_A100"
    """
    Nvidia Tesla A100 GPU.
    """
    NVIDIA_A10080GB = "NVIDIA_A100_80GB"
    """
    Nvidia A100 80GB GPU.
    """
    NVIDIA_L4 = "NVIDIA_L4"
    """
    Nvidia L4 GPU.
    """
    NVIDIA_H10080GB = "NVIDIA_H100_80GB"
    """
    Nvidia H100 80Gb GPU.
    """
    TPU_V2 = "TPU_V2"
    """
    TPU v2.
    """
    TPU_V3 = "TPU_V3"
    """
    TPU v3.
    """
    TPU_V4_POD = "TPU_V4_POD"
    """
    TPU v4.
    """
    TPU_V5_LITEPOD = "TPU_V5_LITEPOD"
    """
    TPU v5.
    """


class GoogleCloudAiplatformV1beta1ModelMonitoringObjectiveConfigExplanationConfigExplanationBaselinePredictionFormat(str, Enum):
    """
    The storage format of the predictions generated BatchPrediction job.
    """
    PREDICTION_FORMAT_UNSPECIFIED = "PREDICTION_FORMAT_UNSPECIFIED"
    """
    Should not be set.
    """
    JSONL = "JSONL"
    """
    Predictions are in JSONL files.
    """
    BIGQUERY = "BIGQUERY"
    """
    Predictions are in BigQuery.
    """


class GoogleCloudAiplatformV1beta1ModelMonitoringStatsAnomaliesObjective(str, Enum):
    """
    Model Monitoring Objective those stats and anomalies belonging to.
    """
    MODEL_DEPLOYMENT_MONITORING_OBJECTIVE_TYPE_UNSPECIFIED = "MODEL_DEPLOYMENT_MONITORING_OBJECTIVE_TYPE_UNSPECIFIED"
    """
    Default value, should not be set.
    """
    RAW_FEATURE_SKEW = "RAW_FEATURE_SKEW"
    """
    Raw feature values' stats to detect skew between Training-Prediction datasets.
    """
    RAW_FEATURE_DRIFT = "RAW_FEATURE_DRIFT"
    """
    Raw feature values' stats to detect drift between Serving-Prediction datasets.
    """
    FEATURE_ATTRIBUTION_SKEW = "FEATURE_ATTRIBUTION_SKEW"
    """
    Feature attribution scores to detect skew between Training-Prediction datasets.
    """
    FEATURE_ATTRIBUTION_DRIFT = "FEATURE_ATTRIBUTION_DRIFT"
    """
    Feature attribution scores to detect skew between Prediction datasets collected within different time windows.
    """


class GoogleCloudAiplatformV1beta1NasJobSpecMultiTrialAlgorithmSpecMetricSpecGoal(str, Enum):
    """
    Required. The optimization goal of the metric.
    """
    GOAL_TYPE_UNSPECIFIED = "GOAL_TYPE_UNSPECIFIED"
    """
    Goal Type will default to maximize.
    """
    MAXIMIZE = "MAXIMIZE"
    """
    Maximize the goal metric.
    """
    MINIMIZE = "MINIMIZE"
    """
    Minimize the goal metric.
    """


class GoogleCloudAiplatformV1beta1NasJobSpecMultiTrialAlgorithmSpecMultiTrialAlgorithm(str, Enum):
    """
    The multi-trial Neural Architecture Search (NAS) algorithm type. Defaults to `REINFORCEMENT_LEARNING`.
    """
    MULTI_TRIAL_ALGORITHM_UNSPECIFIED = "MULTI_TRIAL_ALGORITHM_UNSPECIFIED"
    """
    Defaults to `REINFORCEMENT_LEARNING`.
    """
    REINFORCEMENT_LEARNING = "REINFORCEMENT_LEARNING"
    """
    The Reinforcement Learning Algorithm for Multi-trial Neural Architecture Search (NAS).
    """
    GRID_SEARCH = "GRID_SEARCH"
    """
    The Grid Search Algorithm for Multi-trial Neural Architecture Search (NAS).
    """


class GoogleCloudAiplatformV1beta1PipelineJobRuntimeConfigFailurePolicy(str, Enum):
    """
    Represents the failure policy of a pipeline. Currently, the default of a pipeline is that the pipeline will continue to run until no more tasks can be executed, also known as PIPELINE_FAILURE_POLICY_FAIL_SLOW. However, if a pipeline is set to PIPELINE_FAILURE_POLICY_FAIL_FAST, it will stop scheduling any new tasks when a task has failed. Any scheduled tasks will continue to completion.
    """
    PIPELINE_FAILURE_POLICY_UNSPECIFIED = "PIPELINE_FAILURE_POLICY_UNSPECIFIED"
    """
    Default value, and follows fail slow behavior.
    """
    PIPELINE_FAILURE_POLICY_FAIL_SLOW = "PIPELINE_FAILURE_POLICY_FAIL_SLOW"
    """
    Indicates that the pipeline should continue to run until all possible tasks have been scheduled and completed.
    """
    PIPELINE_FAILURE_POLICY_FAIL_FAST = "PIPELINE_FAILURE_POLICY_FAIL_FAST"
    """
    Indicates that the pipeline should stop scheduling new tasks after a task has failed.
    """


class GoogleCloudAiplatformV1beta1PresetsModality(str, Enum):
    """
    The modality of the uploaded model, which automatically configures the distance measurement and feature normalization for the underlying example index and queries. If your model does not precisely fit one of these types, it is okay to choose the closest type.
    """
    MODALITY_UNSPECIFIED = "MODALITY_UNSPECIFIED"
    """
    Should not be set. Added as a recommended best practice for enums
    """
    IMAGE = "IMAGE"
    """
    IMAGE modality
    """
    TEXT = "TEXT"
    """
    TEXT modality
    """
    TABULAR = "TABULAR"
    """
    TABULAR modality
    """


class GoogleCloudAiplatformV1beta1PresetsQuery(str, Enum):
    """
    Preset option controlling parameters for speed-precision trade-off when querying for examples. If omitted, defaults to `PRECISE`.
    """
    PRECISE = "PRECISE"
    """
    More precise neighbors as a trade-off against slower response.
    """
    FAST = "FAST"
    """
    Faster response as a trade-off against less precise neighbors.
    """


class GoogleCloudAiplatformV1beta1SampleConfigSampleStrategy(str, Enum):
    """
    Field to choose sampling strategy. Sampling strategy will decide which data should be selected for human labeling in every batch.
    """
    SAMPLE_STRATEGY_UNSPECIFIED = "SAMPLE_STRATEGY_UNSPECIFIED"
    """
    Default will be treated as UNCERTAINTY.
    """
    UNCERTAINTY = "UNCERTAINTY"
    """
    Sample the most uncertain data to label.
    """


class GoogleCloudAiplatformV1beta1StudySpecAlgorithm(str, Enum):
    """
    The search algorithm specified for the Study.
    """
    ALGORITHM_UNSPECIFIED = "ALGORITHM_UNSPECIFIED"
    """
    The default algorithm used by Vertex AI for [hyperparameter tuning](https://cloud.google.com/vertex-ai/docs/training/hyperparameter-tuning-overview) and [Vertex AI Vizier](https://cloud.google.com/vertex-ai/docs/vizier).
    """
    GRID_SEARCH = "GRID_SEARCH"
    """
    Simple grid search within the feasible space. To use grid search, all parameters must be `INTEGER`, `CATEGORICAL`, or `DISCRETE`.
    """
    RANDOM_SEARCH = "RANDOM_SEARCH"
    """
    Simple random search within the feasible space.
    """


class GoogleCloudAiplatformV1beta1StudySpecMeasurementSelectionType(str, Enum):
    """
    Describe which measurement selection type will be used
    """
    MEASUREMENT_SELECTION_TYPE_UNSPECIFIED = "MEASUREMENT_SELECTION_TYPE_UNSPECIFIED"
    """
    Will be treated as LAST_MEASUREMENT.
    """
    LAST_MEASUREMENT = "LAST_MEASUREMENT"
    """
    Use the last measurement reported.
    """
    BEST_MEASUREMENT = "BEST_MEASUREMENT"
    """
    Use the best measurement reported.
    """


class GoogleCloudAiplatformV1beta1StudySpecMetricSpecGoal(str, Enum):
    """
    Required. The optimization goal of the metric.
    """
    GOAL_TYPE_UNSPECIFIED = "GOAL_TYPE_UNSPECIFIED"
    """
    Goal Type will default to maximize.
    """
    MAXIMIZE = "MAXIMIZE"
    """
    Maximize the goal metric.
    """
    MINIMIZE = "MINIMIZE"
    """
    Minimize the goal metric.
    """


class GoogleCloudAiplatformV1beta1StudySpecObservationNoise(str, Enum):
    """
    The observation noise level of the study. Currently only supported by the Vertex AI Vizier service. Not supported by HyperparameterTuningJob or TrainingPipeline.
    """
    OBSERVATION_NOISE_UNSPECIFIED = "OBSERVATION_NOISE_UNSPECIFIED"
    """
    The default noise level chosen by Vertex AI.
    """
    LOW = "LOW"
    """
    Vertex AI assumes that the objective function is (nearly) perfectly reproducible, and will never repeat the same Trial parameters.
    """
    HIGH = "HIGH"
    """
    Vertex AI will estimate the amount of noise in metric evaluations, it may repeat the same Trial parameters more than once.
    """


class GoogleCloudAiplatformV1beta1StudySpecParameterSpecScaleType(str, Enum):
    """
    How the parameter should be scaled. Leave unset for `CATEGORICAL` parameters.
    """
    SCALE_TYPE_UNSPECIFIED = "SCALE_TYPE_UNSPECIFIED"
    """
    By default, no scaling is applied.
    """
    UNIT_LINEAR_SCALE = "UNIT_LINEAR_SCALE"
    """
    Scales the feasible space to (0, 1) linearly.
    """
    UNIT_LOG_SCALE = "UNIT_LOG_SCALE"
    """
    Scales the feasible space logarithmically to (0, 1). The entire feasible space must be strictly positive.
    """
    UNIT_REVERSE_LOG_SCALE = "UNIT_REVERSE_LOG_SCALE"
    """
    Scales the feasible space "reverse" logarithmically to (0, 1). The result is that values close to the top of the feasible space are spread out more than points near the bottom. The entire feasible space must be strictly positive.
    """


class IndexIndexUpdateMethod(str, Enum):
    """
    Immutable. The update method to use with this Index. If not set, BATCH_UPDATE will be used by default.
    """
    INDEX_UPDATE_METHOD_UNSPECIFIED = "INDEX_UPDATE_METHOD_UNSPECIFIED"
    """
    Should not be used.
    """
    BATCH_UPDATE = "BATCH_UPDATE"
    """
    BatchUpdate: user can call UpdateIndex with files on Cloud Storage of Datapoints to update.
    """
    STREAM_UPDATE = "STREAM_UPDATE"
    """
    StreamUpdate: user can call UpsertDatapoints/DeleteDatapoints to update the Index and the updates will be applied in corresponding DeployedIndexes in nearly real-time.
    """


class MetadataSchemaSchemaType(str, Enum):
    """
    The type of the MetadataSchema. This is a property that identifies which metadata types will use the MetadataSchema.
    """
    METADATA_SCHEMA_TYPE_UNSPECIFIED = "METADATA_SCHEMA_TYPE_UNSPECIFIED"
    """
    Unspecified type for the MetadataSchema.
    """
    ARTIFACT_TYPE = "ARTIFACT_TYPE"
    """
    A type indicating that the MetadataSchema will be used by Artifacts.
    """
    EXECUTION_TYPE = "EXECUTION_TYPE"
    """
    A typee indicating that the MetadataSchema will be used by Executions.
    """
    CONTEXT_TYPE = "CONTEXT_TYPE"
    """
    A state indicating that the MetadataSchema will be used by Contexts.
    """


class NotebookRuntimeTemplateNotebookRuntimeType(str, Enum):
    """
    Optional. Immutable. The type of the notebook runtime template.
    """
    NOTEBOOK_RUNTIME_TYPE_UNSPECIFIED = "NOTEBOOK_RUNTIME_TYPE_UNSPECIFIED"
    """
    Unspecified notebook runtime type, NotebookRuntimeType will default to USER_DEFINED.
    """
    USER_DEFINED = "USER_DEFINED"
    """
    runtime or template with coustomized configurations from user.
    """
    ONE_CLICK = "ONE_CLICK"
    """
    runtime or template with system defined configurations.
    """


class TimeSeriesValueType(str, Enum):
    """
    Required. Immutable. Type of TensorboardTimeSeries value.
    """
    VALUE_TYPE_UNSPECIFIED = "VALUE_TYPE_UNSPECIFIED"
    """
    The value type is unspecified.
    """
    SCALAR = "SCALAR"
    """
    Used for TensorboardTimeSeries that is a list of scalars. E.g. accuracy of a model over epochs/time.
    """
    TENSOR = "TENSOR"
    """
    Used for TensorboardTimeSeries that is a list of tensors. E.g. histograms of weights of layer in a model over epoch/time.
    """
    BLOB_SEQUENCE = "BLOB_SEQUENCE"
    """
    Used for TensorboardTimeSeries that is a list of blob sequences. E.g. set of sample images with labels over epochs/time.
    """
