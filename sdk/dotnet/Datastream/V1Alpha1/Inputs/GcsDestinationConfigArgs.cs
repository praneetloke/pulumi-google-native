// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1Alpha1.Inputs
{

    /// <summary>
    /// Google Cloud Storage destination configuration
    /// </summary>
    public sealed class GcsDestinationConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AVRO file format configuration.
        /// </summary>
        [Input("avroFileFormat")]
        public Input<Inputs.AvroFileFormatArgs>? AvroFileFormat { get; set; }

        /// <summary>
        /// The maximum duration for which new events are added before a file is closed and a new file is created.
        /// </summary>
        [Input("fileRotationInterval")]
        public Input<string>? FileRotationInterval { get; set; }

        /// <summary>
        /// The maximum file size to be saved in the bucket.
        /// </summary>
        [Input("fileRotationMb")]
        public Input<int>? FileRotationMb { get; set; }

        /// <summary>
        /// File format that data should be written in. Deprecated field (b/169501737) - use file_format instead.
        /// </summary>
        [Input("gcsFileFormat")]
        public Input<Pulumi.GoogleNative.Datastream.V1Alpha1.GcsDestinationConfigGcsFileFormat>? GcsFileFormat { get; set; }

        /// <summary>
        /// JSON file format configuration.
        /// </summary>
        [Input("jsonFileFormat")]
        public Input<Inputs.JsonFileFormatArgs>? JsonFileFormat { get; set; }

        /// <summary>
        /// Path inside the Cloud Storage bucket to write data to.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public GcsDestinationConfigArgs()
        {
        }
        public static new GcsDestinationConfigArgs Empty => new GcsDestinationConfigArgs();
    }
}
