// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Inputs
{

    /// <summary>
    /// Options defining BigQuery table and row identifiers.
    /// </summary>
    public sealed class GooglePrivacyDlpV2BigQueryOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("excludedFields")]
        private InputList<Inputs.GooglePrivacyDlpV2FieldIdArgs>? _excludedFields;

        /// <summary>
        /// References to fields excluded from scanning. This allows you to skip inspection of entire columns which you know have no findings. When inspecting a table, we recommend that you inspect all columns. Otherwise, findings might be affected because hints from excluded columns will not be used.
        /// </summary>
        public InputList<Inputs.GooglePrivacyDlpV2FieldIdArgs> ExcludedFields
        {
            get => _excludedFields ?? (_excludedFields = new InputList<Inputs.GooglePrivacyDlpV2FieldIdArgs>());
            set => _excludedFields = value;
        }

        [Input("identifyingFields")]
        private InputList<Inputs.GooglePrivacyDlpV2FieldIdArgs>? _identifyingFields;

        /// <summary>
        /// Table fields that may uniquely identify a row within the table. When `actions.saveFindings.outputConfig.table` is specified, the values of columns specified here are available in the output table under `location.content_locations.record_location.record_key.id_values`. Nested fields such as `person.birthdate.year` are allowed.
        /// </summary>
        public InputList<Inputs.GooglePrivacyDlpV2FieldIdArgs> IdentifyingFields
        {
            get => _identifyingFields ?? (_identifyingFields = new InputList<Inputs.GooglePrivacyDlpV2FieldIdArgs>());
            set => _identifyingFields = value;
        }

        [Input("includedFields")]
        private InputList<Inputs.GooglePrivacyDlpV2FieldIdArgs>? _includedFields;

        /// <summary>
        /// Limit scanning only to these fields. When inspecting a table, we recommend that you inspect all columns. Otherwise, findings might be affected because hints from excluded columns will not be used.
        /// </summary>
        public InputList<Inputs.GooglePrivacyDlpV2FieldIdArgs> IncludedFields
        {
            get => _includedFields ?? (_includedFields = new InputList<Inputs.GooglePrivacyDlpV2FieldIdArgs>());
            set => _includedFields = value;
        }

        /// <summary>
        /// Max number of rows to scan. If the table has more rows than this value, the rest of the rows are omitted. If not set, or if set to 0, all rows will be scanned. Only one of rows_limit and rows_limit_percent can be specified. Cannot be used in conjunction with TimespanConfig.
        /// </summary>
        [Input("rowsLimit")]
        public Input<string>? RowsLimit { get; set; }

        /// <summary>
        /// Max percentage of rows to scan. The rest are omitted. The number of rows scanned is rounded down. Must be between 0 and 100, inclusively. Both 0 and 100 means no limit. Defaults to 0. Only one of rows_limit and rows_limit_percent can be specified. Cannot be used in conjunction with TimespanConfig. Caution: A [known issue](https://cloud.google.com/dlp/docs/known-issues#bq-sampling) is causing the `rowsLimitPercent` field to behave unexpectedly. We recommend using `rowsLimit` instead.
        /// </summary>
        [Input("rowsLimitPercent")]
        public Input<int>? RowsLimitPercent { get; set; }

        [Input("sampleMethod")]
        public Input<Pulumi.GoogleNative.DLP.V2.GooglePrivacyDlpV2BigQueryOptionsSampleMethod>? SampleMethod { get; set; }

        /// <summary>
        /// Complete BigQuery table reference.
        /// </summary>
        [Input("tableReference")]
        public Input<Inputs.GooglePrivacyDlpV2BigQueryTableArgs>? TableReference { get; set; }

        public GooglePrivacyDlpV2BigQueryOptionsArgs()
        {
        }
        public static new GooglePrivacyDlpV2BigQueryOptionsArgs Empty => new GooglePrivacyDlpV2BigQueryOptionsArgs();
    }
}
