// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1Beta1.Outputs
{

    /// <summary>
    /// Specifies Vertex AI owned input data to be used for training, and possibly evaluating, the Model.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudAiplatformV1beta1InputDataConfigResponse
    {
        /// <summary>
        /// Applicable only to custom training with Datasets that have DataItems and Annotations. Cloud Storage URI that points to a YAML file describing the annotation schema. The schema is defined as an OpenAPI 3.0.2 [Schema Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject). The schema files that can be used here are found in gs://google-cloud-aiplatform/schema/dataset/annotation/ , note that the chosen schema must be consistent with metadata of the Dataset specified by dataset_id. Only Annotations that both match this schema and belong to DataItems not ignored by the split method are used in respectively training, validation or test role, depending on the role of the DataItem they are on. When used in conjunction with annotations_filter, the Annotations used for training are filtered by both annotations_filter and annotation_schema_uri.
        /// </summary>
        public readonly string AnnotationSchemaUri;
        /// <summary>
        /// Applicable only to Datasets that have DataItems and Annotations. A filter on Annotations of the Dataset. Only Annotations that both match this filter and belong to DataItems not ignored by the split method are used in respectively training, validation or test role, depending on the role of the DataItem they are on (for the auto-assigned that role is decided by Vertex AI). A filter with same syntax as the one used in ListAnnotations may be used, but note here it filters across all Annotations of the Dataset, and not just within a single DataItem.
        /// </summary>
        public readonly string AnnotationsFilter;
        /// <summary>
        /// Only applicable to custom training with tabular Dataset with BigQuery source. The BigQuery project location where the training data is to be written to. In the given project a new dataset is created with name `dataset___` where timestamp is in YYYY_MM_DDThh_mm_ss_sssZ format. All training input data is written into that dataset. In the dataset three tables are created, `training`, `validation` and `test`. * AIP_DATA_FORMAT = "bigquery". * AIP_TRAINING_DATA_URI = "bigquery_destination.dataset___.training" * AIP_VALIDATION_DATA_URI = "bigquery_destination.dataset___.validation" * AIP_TEST_DATA_URI = "bigquery_destination.dataset___.test"
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1beta1BigQueryDestinationResponse BigqueryDestination;
        /// <summary>
        /// The ID of the Dataset in the same Project and Location which data will be used to train the Model. The Dataset must use schema compatible with Model being trained, and what is compatible should be described in the used TrainingPipeline's training_task_definition. For tabular Datasets, all their data is exported to training, to pick and choose from.
        /// </summary>
        public readonly string DatasetId;
        /// <summary>
        /// Split based on the provided filters for each set.
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1beta1FilterSplitResponse FilterSplit;
        /// <summary>
        /// Split based on fractions defining the size of each set.
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1beta1FractionSplitResponse FractionSplit;
        /// <summary>
        /// The Cloud Storage location where the training data is to be written to. In the given directory a new directory is created with name: `dataset---` where timestamp is in YYYY-MM-DDThh:mm:ss.sssZ ISO-8601 format. All training input data is written into that directory. The Vertex AI environment variables representing Cloud Storage data URIs are represented in the Cloud Storage wildcard format to support sharded data. e.g.: "gs://.../training-*.jsonl" * AIP_DATA_FORMAT = "jsonl" for non-tabular data, "csv" for tabular data * AIP_TRAINING_DATA_URI = "gcs_destination/dataset---/training-*.${AIP_DATA_FORMAT}" * AIP_VALIDATION_DATA_URI = "gcs_destination/dataset---/validation-*.${AIP_DATA_FORMAT}" * AIP_TEST_DATA_URI = "gcs_destination/dataset---/test-*.${AIP_DATA_FORMAT}"
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1beta1GcsDestinationResponse GcsDestination;
        /// <summary>
        /// Whether to persist the ML use assignment to data item system labels.
        /// </summary>
        public readonly bool PersistMlUseAssignment;
        /// <summary>
        /// Supported only for tabular Datasets. Split based on a predefined key.
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1beta1PredefinedSplitResponse PredefinedSplit;
        /// <summary>
        /// Only applicable to Datasets that have SavedQueries. The ID of a SavedQuery (annotation set) under the Dataset specified by dataset_id used for filtering Annotations for training. Only Annotations that are associated with this SavedQuery are used in respectively training. When used in conjunction with annotations_filter, the Annotations used for training are filtered by both saved_query_id and annotations_filter. Only one of saved_query_id and annotation_schema_uri should be specified as both of them represent the same thing: problem type.
        /// </summary>
        public readonly string SavedQueryId;
        /// <summary>
        /// Supported only for tabular Datasets. Split based on the distribution of the specified column.
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1beta1StratifiedSplitResponse StratifiedSplit;
        /// <summary>
        /// Supported only for tabular Datasets. Split based on the timestamp of the input data pieces.
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1beta1TimestampSplitResponse TimestampSplit;

        [OutputConstructor]
        private GoogleCloudAiplatformV1beta1InputDataConfigResponse(
            string annotationSchemaUri,

            string annotationsFilter,

            Outputs.GoogleCloudAiplatformV1beta1BigQueryDestinationResponse bigqueryDestination,

            string datasetId,

            Outputs.GoogleCloudAiplatformV1beta1FilterSplitResponse filterSplit,

            Outputs.GoogleCloudAiplatformV1beta1FractionSplitResponse fractionSplit,

            Outputs.GoogleCloudAiplatformV1beta1GcsDestinationResponse gcsDestination,

            bool persistMlUseAssignment,

            Outputs.GoogleCloudAiplatformV1beta1PredefinedSplitResponse predefinedSplit,

            string savedQueryId,

            Outputs.GoogleCloudAiplatformV1beta1StratifiedSplitResponse stratifiedSplit,

            Outputs.GoogleCloudAiplatformV1beta1TimestampSplitResponse timestampSplit)
        {
            AnnotationSchemaUri = annotationSchemaUri;
            AnnotationsFilter = annotationsFilter;
            BigqueryDestination = bigqueryDestination;
            DatasetId = datasetId;
            FilterSplit = filterSplit;
            FractionSplit = fractionSplit;
            GcsDestination = gcsDestination;
            PersistMlUseAssignment = persistMlUseAssignment;
            PredefinedSplit = predefinedSplit;
            SavedQueryId = savedQueryId;
            StratifiedSplit = stratifiedSplit;
            TimestampSplit = timestampSplit;
        }
    }
}