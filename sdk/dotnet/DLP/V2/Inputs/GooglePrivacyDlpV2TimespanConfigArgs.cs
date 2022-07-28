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
    /// Configuration of the timespan of the items to include in scanning. Currently only supported when inspecting Google Cloud Storage and BigQuery.
    /// </summary>
    public sealed class GooglePrivacyDlpV2TimespanConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When the job is started by a JobTrigger we will automatically figure out a valid start_time to avoid scanning files that have not been modified since the last time the JobTrigger executed. This will be based on the time of the execution of the last run of the JobTrigger.
        /// </summary>
        [Input("enableAutoPopulationOfTimespanConfig")]
        public Input<bool>? EnableAutoPopulationOfTimespanConfig { get; set; }

        /// <summary>
        /// Exclude files, tables, or rows newer than this value. If not set, no upper time limit is applied.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Exclude files, tables, or rows older than this value. If not set, no lower time limit is applied.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// Specification of the field containing the timestamp of scanned items. Used for data sources like Datastore and BigQuery. For BigQuery: If this value is not specified and the table was modified between the given start and end times, the entire table will be scanned. If this value is specified, then rows are filtered based on the given start and end times. Rows with a `NULL` value in the provided BigQuery column are skipped. Valid data types of the provided BigQuery column are: `INTEGER`, `DATE`, `TIMESTAMP`, and `DATETIME`. For Datastore: If this value is specified, then entities are filtered based on the given start and end times. If an entity does not contain the provided timestamp property or contains empty or invalid values, then it is included. Valid data types of the provided timestamp property are: `TIMESTAMP`.
        /// </summary>
        [Input("timestampField")]
        public Input<Inputs.GooglePrivacyDlpV2FieldIdArgs>? TimestampField { get; set; }

        public GooglePrivacyDlpV2TimespanConfigArgs()
        {
        }
        public static new GooglePrivacyDlpV2TimespanConfigArgs Empty => new GooglePrivacyDlpV2TimespanConfigArgs();
    }
}
