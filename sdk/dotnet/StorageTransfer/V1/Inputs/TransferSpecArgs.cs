// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.StorageTransfer.V1.Inputs
{

    /// <summary>
    /// Configuration for running a transfer.
    /// </summary>
    public sealed class TransferSpecArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An AWS S3 compatible data source.
        /// </summary>
        [Input("awsS3CompatibleDataSource")]
        public Input<Inputs.AwsS3CompatibleDataArgs>? AwsS3CompatibleDataSource { get; set; }

        /// <summary>
        /// An AWS S3 data source.
        /// </summary>
        [Input("awsS3DataSource")]
        public Input<Inputs.AwsS3DataArgs>? AwsS3DataSource { get; set; }

        /// <summary>
        /// An Azure Blob Storage data source.
        /// </summary>
        [Input("azureBlobStorageDataSource")]
        public Input<Inputs.AzureBlobStorageDataArgs>? AzureBlobStorageDataSource { get; set; }

        /// <summary>
        /// A Cloud Storage data sink.
        /// </summary>
        [Input("gcsDataSink")]
        public Input<Inputs.GcsDataArgs>? GcsDataSink { get; set; }

        /// <summary>
        /// A Cloud Storage data source.
        /// </summary>
        [Input("gcsDataSource")]
        public Input<Inputs.GcsDataArgs>? GcsDataSource { get; set; }

        /// <summary>
        /// Cloud Storage intermediate data location.
        /// </summary>
        [Input("gcsIntermediateDataLocation")]
        public Input<Inputs.GcsDataArgs>? GcsIntermediateDataLocation { get; set; }

        /// <summary>
        /// An HTTP URL data source.
        /// </summary>
        [Input("httpDataSource")]
        public Input<Inputs.HttpDataArgs>? HttpDataSource { get; set; }

        /// <summary>
        /// Only objects that satisfy these object conditions are included in the set of data source and data sink objects. Object conditions based on objects' "last modification time" do not exclude objects in a data sink.
        /// </summary>
        [Input("objectConditions")]
        public Input<Inputs.ObjectConditionsArgs>? ObjectConditions { get; set; }

        /// <summary>
        /// A POSIX Filesystem data sink.
        /// </summary>
        [Input("posixDataSink")]
        public Input<Inputs.PosixFilesystemArgs>? PosixDataSink { get; set; }

        /// <summary>
        /// A POSIX Filesystem data source.
        /// </summary>
        [Input("posixDataSource")]
        public Input<Inputs.PosixFilesystemArgs>? PosixDataSource { get; set; }

        /// <summary>
        /// Specifies the agent pool name associated with the posix data sink. When unspecified, the default name is used.
        /// </summary>
        [Input("sinkAgentPoolName")]
        public Input<string>? SinkAgentPoolName { get; set; }

        /// <summary>
        /// Specifies the agent pool name associated with the posix data source. When unspecified, the default name is used.
        /// </summary>
        [Input("sourceAgentPoolName")]
        public Input<string>? SourceAgentPoolName { get; set; }

        /// <summary>
        /// A manifest file provides a list of objects to be transferred from the data source. This field points to the location of the manifest file. Otherwise, the entire source bucket is used. ObjectConditions still apply.
        /// </summary>
        [Input("transferManifest")]
        public Input<Inputs.TransferManifestArgs>? TransferManifest { get; set; }

        /// <summary>
        /// If the option delete_objects_unique_in_sink is `true` and time-based object conditions such as 'last modification time' are specified, the request fails with an INVALID_ARGUMENT error.
        /// </summary>
        [Input("transferOptions")]
        public Input<Inputs.TransferOptionsArgs>? TransferOptions { get; set; }

        public TransferSpecArgs()
        {
        }
    }
}
