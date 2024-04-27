// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1.Inputs
{

    /// <summary>
    /// Specifies Vertex AI owned input data to be used for training, and possibly evaluating, the Model.
    /// </summary>
    public sealed class GoogleCloudAiplatformV1InputDataConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Applicable only to custom training with Datasets that have DataItems and Annotations. Cloud Storage URI that points to a YAML file describing the annotation schema. The schema is defined as an OpenAPI 3.0.2 [Schema Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject). The schema files that can be used here are found in gs://google-cloud-aiplatform/schema/dataset/annotation/ , note that the chosen schema must be consistent with metadata of the Dataset specified by dataset_id. Only Annotations that both match this schema and belong to DataItems not ignored by the split method are used in respectively training, validation or test role, depending on the role of the DataItem they are on. When used in conjunction with annotations_filter, the Annotations used for training are filtered by both annotations_filter and annotation_schema_uri.
        /// </summary>
        [Input("annotationSchemaUri")]
        public Input<string>? AnnotationSchemaUri { get; set; }

        /// <summary>
        /// Applicable only to Datasets that have DataItems and Annotations. A filter on Annotations of the Dataset. Only Annotations that both match this filter and belong to DataItems not ignored by the split method are used in respectively training, validation or test role, depending on the role of the DataItem they are on (for the auto-assigned that role is decided by Vertex AI). A filter with same syntax as the one used in ListAnnotations may be used, but note here it filters across all Annotations of the Dataset, and not just within a single DataItem.
        /// </summary>
        [Input("annotationsFilter")]
        public Input<string>? AnnotationsFilter { get; set; }

        /// <summary>
        /// Only applicable to custom training with tabular Dataset with BigQuery source. The BigQuery project location where the training data is to be written to. In the given project a new dataset is created with name `dataset___` where timestamp is in YYYY_MM_DDThh_mm_ss_sssZ format. All training input data is written into that dataset. In the dataset three tables are created, `training`, `validation` and `test`. * AIP_DATA_FORMAT = "bigquery". * AIP_TRAINING_DATA_URI = "bigquery_destination.dataset___.training" * AIP_VALIDATION_DATA_URI = "bigquery_destination.dataset___.validation" * AIP_TEST_DATA_URI = "bigquery_destination.dataset___.test"
        /// </summary>
        [Input("bigqueryDestination")]
        public Input<Inputs.GoogleCloudAiplatformV1BigQueryDestinationArgs>? BigqueryDestination { get; set; }

        /// <summary>
        /// The ID of the Dataset in the same Project and Location which data will be used to train the Model. The Dataset must use schema compatible with Model being trained, and what is compatible should be described in the used TrainingPipeline's training_task_definition. For tabular Datasets, all their data is exported to training, to pick and choose from.
        /// </summary>
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        /// <summary>
        /// Split based on the provided filters for each set.
        /// </summary>
        [Input("filterSplit")]
        public Input<Inputs.GoogleCloudAiplatformV1FilterSplitArgs>? FilterSplit { get; set; }

        /// <summary>
        /// Split based on fractions defining the size of each set.
        /// </summary>
        [Input("fractionSplit")]
        public Input<Inputs.GoogleCloudAiplatformV1FractionSplitArgs>? FractionSplit { get; set; }

        /// <summary>
        /// The Cloud Storage location where the training data is to be written to. In the given directory a new directory is created with name: `dataset---` where timestamp is in YYYY-MM-DDThh:mm:ss.sssZ ISO-8601 format. All training input data is written into that directory. The Vertex AI environment variables representing Cloud Storage data URIs are represented in the Cloud Storage wildcard format to support sharded data. e.g.: "gs://.../training-*.jsonl" * AIP_DATA_FORMAT = "jsonl" for non-tabular data, "csv" for tabular data * AIP_TRAINING_DATA_URI = "gcs_destination/dataset---/training-*.${AIP_DATA_FORMAT}" * AIP_VALIDATION_DATA_URI = "gcs_destination/dataset---/validation-*.${AIP_DATA_FORMAT}" * AIP_TEST_DATA_URI = "gcs_destination/dataset---/test-*.${AIP_DATA_FORMAT}"
        /// </summary>
        [Input("gcsDestination")]
        public Input<Inputs.GoogleCloudAiplatformV1GcsDestinationArgs>? GcsDestination { get; set; }

        /// <summary>
        /// Whether to persist the ML use assignment to data item system labels.
        /// </summary>
        [Input("persistMlUseAssignment")]
        public Input<bool>? PersistMlUseAssignment { get; set; }

        /// <summary>
        /// Supported only for tabular Datasets. Split based on a predefined key.
        /// </summary>
        [Input("predefinedSplit")]
        public Input<Inputs.GoogleCloudAiplatformV1PredefinedSplitArgs>? PredefinedSplit { get; set; }

        /// <summary>
        /// Only applicable to Datasets that have SavedQueries. The ID of a SavedQuery (annotation set) under the Dataset specified by dataset_id used for filtering Annotations for training. Only Annotations that are associated with this SavedQuery are used in respectively training. When used in conjunction with annotations_filter, the Annotations used for training are filtered by both saved_query_id and annotations_filter. Only one of saved_query_id and annotation_schema_uri should be specified as both of them represent the same thing: problem type.
        /// </summary>
        [Input("savedQueryId")]
        public Input<string>? SavedQueryId { get; set; }

        /// <summary>
        /// Supported only for tabular Datasets. Split based on the distribution of the specified column.
        /// </summary>
        [Input("stratifiedSplit")]
        public Input<Inputs.GoogleCloudAiplatformV1StratifiedSplitArgs>? StratifiedSplit { get; set; }

        /// <summary>
        /// Supported only for tabular Datasets. Split based on the timestamp of the input data pieces.
        /// </summary>
        [Input("timestampSplit")]
        public Input<Inputs.GoogleCloudAiplatformV1TimestampSplitArgs>? TimestampSplit { get; set; }

        public GoogleCloudAiplatformV1InputDataConfigArgs()
        {
        }
        public static new GoogleCloudAiplatformV1InputDataConfigArgs Empty => new GoogleCloudAiplatformV1InputDataConfigArgs();
    }
}