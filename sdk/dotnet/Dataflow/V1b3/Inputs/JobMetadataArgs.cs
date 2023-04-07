// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataflow.V1b3.Inputs
{

    /// <summary>
    /// Metadata available primarily for filtering jobs. Will be included in the ListJob response and Job SUMMARY view.
    /// </summary>
    public sealed class JobMetadataArgs : global::Pulumi.ResourceArgs
    {
        [Input("bigTableDetails")]
        private InputList<Inputs.BigTableIODetailsArgs>? _bigTableDetails;

        /// <summary>
        /// Identification of a Cloud Bigtable source used in the Dataflow job.
        /// </summary>
        public InputList<Inputs.BigTableIODetailsArgs> BigTableDetails
        {
            get => _bigTableDetails ?? (_bigTableDetails = new InputList<Inputs.BigTableIODetailsArgs>());
            set => _bigTableDetails = value;
        }

        [Input("bigqueryDetails")]
        private InputList<Inputs.BigQueryIODetailsArgs>? _bigqueryDetails;

        /// <summary>
        /// Identification of a BigQuery source used in the Dataflow job.
        /// </summary>
        public InputList<Inputs.BigQueryIODetailsArgs> BigqueryDetails
        {
            get => _bigqueryDetails ?? (_bigqueryDetails = new InputList<Inputs.BigQueryIODetailsArgs>());
            set => _bigqueryDetails = value;
        }

        [Input("datastoreDetails")]
        private InputList<Inputs.DatastoreIODetailsArgs>? _datastoreDetails;

        /// <summary>
        /// Identification of a Datastore source used in the Dataflow job.
        /// </summary>
        public InputList<Inputs.DatastoreIODetailsArgs> DatastoreDetails
        {
            get => _datastoreDetails ?? (_datastoreDetails = new InputList<Inputs.DatastoreIODetailsArgs>());
            set => _datastoreDetails = value;
        }

        [Input("fileDetails")]
        private InputList<Inputs.FileIODetailsArgs>? _fileDetails;

        /// <summary>
        /// Identification of a File source used in the Dataflow job.
        /// </summary>
        public InputList<Inputs.FileIODetailsArgs> FileDetails
        {
            get => _fileDetails ?? (_fileDetails = new InputList<Inputs.FileIODetailsArgs>());
            set => _fileDetails = value;
        }

        [Input("pubsubDetails")]
        private InputList<Inputs.PubSubIODetailsArgs>? _pubsubDetails;

        /// <summary>
        /// Identification of a Pub/Sub source used in the Dataflow job.
        /// </summary>
        public InputList<Inputs.PubSubIODetailsArgs> PubsubDetails
        {
            get => _pubsubDetails ?? (_pubsubDetails = new InputList<Inputs.PubSubIODetailsArgs>());
            set => _pubsubDetails = value;
        }

        /// <summary>
        /// The SDK version used to run the job.
        /// </summary>
        [Input("sdkVersion")]
        public Input<Inputs.SdkVersionArgs>? SdkVersion { get; set; }

        [Input("spannerDetails")]
        private InputList<Inputs.SpannerIODetailsArgs>? _spannerDetails;

        /// <summary>
        /// Identification of a Spanner source used in the Dataflow job.
        /// </summary>
        public InputList<Inputs.SpannerIODetailsArgs> SpannerDetails
        {
            get => _spannerDetails ?? (_spannerDetails = new InputList<Inputs.SpannerIODetailsArgs>());
            set => _spannerDetails = value;
        }

        [Input("userDisplayProperties")]
        private InputMap<string>? _userDisplayProperties;

        /// <summary>
        /// List of display properties to help UI filter jobs.
        /// </summary>
        public InputMap<string> UserDisplayProperties
        {
            get => _userDisplayProperties ?? (_userDisplayProperties = new InputMap<string>());
            set => _userDisplayProperties = value;
        }

        public JobMetadataArgs()
        {
        }
        public static new JobMetadataArgs Empty => new JobMetadataArgs();
    }
}
