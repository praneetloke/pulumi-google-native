// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Inputs
{

    /// <summary>
    /// Settings to manage the metadata discovery and publishing in a zone.
    /// </summary>
    public sealed class GoogleCloudDataplexV1ZoneDiscoverySpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Configuration for CSV data.
        /// </summary>
        [Input("csvOptions")]
        public Input<Inputs.GoogleCloudDataplexV1ZoneDiscoverySpecCsvOptionsArgs>? CsvOptions { get; set; }

        /// <summary>
        /// Whether discovery is enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("excludePatterns")]
        private InputList<string>? _excludePatterns;

        /// <summary>
        /// Optional. The list of patterns to apply for selecting data to exclude during discovery. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
        /// </summary>
        public InputList<string> ExcludePatterns
        {
            get => _excludePatterns ?? (_excludePatterns = new InputList<string>());
            set => _excludePatterns = value;
        }

        [Input("includePatterns")]
        private InputList<string>? _includePatterns;

        /// <summary>
        /// Optional. The list of patterns to apply for selecting data to include during discovery if only a subset of the data should considered. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
        /// </summary>
        public InputList<string> IncludePatterns
        {
            get => _includePatterns ?? (_includePatterns = new InputList<string>());
            set => _includePatterns = value;
        }

        /// <summary>
        /// Optional. Configuration for Json data.
        /// </summary>
        [Input("jsonOptions")]
        public Input<Inputs.GoogleCloudDataplexV1ZoneDiscoverySpecJsonOptionsArgs>? JsonOptions { get; set; }

        /// <summary>
        /// Optional. Cron schedule (https://en.wikipedia.org/wiki/Cron) for running discovery periodically. Successive discovery runs must be scheduled at least 60 minutes apart. The default value is to run discovery every 60 minutes. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or TZ=${IANA_TIME_ZONE}". The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone database. For example, "CRON_TZ=America/New_York 1 * * * *", or "TZ=America/New_York 1 * * * *".
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        public GoogleCloudDataplexV1ZoneDiscoverySpecArgs()
        {
        }
        public static new GoogleCloudDataplexV1ZoneDiscoverySpecArgs Empty => new GoogleCloudDataplexV1ZoneDiscoverySpecArgs();
    }
}
