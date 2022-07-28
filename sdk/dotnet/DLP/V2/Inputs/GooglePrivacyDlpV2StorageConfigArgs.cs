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
    /// Shared message indicating Cloud storage type.
    /// </summary>
    public sealed class GooglePrivacyDlpV2StorageConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// BigQuery options.
        /// </summary>
        [Input("bigQueryOptions")]
        public Input<Inputs.GooglePrivacyDlpV2BigQueryOptionsArgs>? BigQueryOptions { get; set; }

        /// <summary>
        /// Google Cloud Storage options.
        /// </summary>
        [Input("cloudStorageOptions")]
        public Input<Inputs.GooglePrivacyDlpV2CloudStorageOptionsArgs>? CloudStorageOptions { get; set; }

        /// <summary>
        /// Google Cloud Datastore options.
        /// </summary>
        [Input("datastoreOptions")]
        public Input<Inputs.GooglePrivacyDlpV2DatastoreOptionsArgs>? DatastoreOptions { get; set; }

        /// <summary>
        /// Hybrid inspection options.
        /// </summary>
        [Input("hybridOptions")]
        public Input<Inputs.GooglePrivacyDlpV2HybridOptionsArgs>? HybridOptions { get; set; }

        [Input("timespanConfig")]
        public Input<Inputs.GooglePrivacyDlpV2TimespanConfigArgs>? TimespanConfig { get; set; }

        public GooglePrivacyDlpV2StorageConfigArgs()
        {
        }
        public static new GooglePrivacyDlpV2StorageConfigArgs Empty => new GooglePrivacyDlpV2StorageConfigArgs();
    }
}
