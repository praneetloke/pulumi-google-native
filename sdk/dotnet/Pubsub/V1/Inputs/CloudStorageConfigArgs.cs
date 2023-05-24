// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Pubsub.V1.Inputs
{

    /// <summary>
    /// Configuration for a Cloud Storage subscription.
    /// </summary>
    public sealed class CloudStorageConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, message data will be written to Cloud Storage in Avro format.
        /// </summary>
        [Input("avroConfig")]
        public Input<Inputs.AvroConfigArgs>? AvroConfig { get; set; }

        /// <summary>
        /// User-provided name for the Cloud Storage bucket. The bucket must be created by the user. The bucket name must be without any prefix like "gs://". See the [bucket naming requirements] (https://cloud.google.com/storage/docs/buckets#naming).
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// User-provided prefix for Cloud Storage filename. See the [object naming requirements](https://cloud.google.com/storage/docs/objects#naming).
        /// </summary>
        [Input("filenamePrefix")]
        public Input<string>? FilenamePrefix { get; set; }

        /// <summary>
        /// User-provided suffix for Cloud Storage filename. See the [object naming requirements](https://cloud.google.com/storage/docs/objects#naming).
        /// </summary>
        [Input("filenameSuffix")]
        public Input<string>? FilenameSuffix { get; set; }

        /// <summary>
        /// The maximum bytes that can be written to a Cloud Storage file before a new file is created. Min 1 KB, max 10 GiB. The max_bytes limit may be exceeded in cases where messages are larger than the limit.
        /// </summary>
        [Input("maxBytes")]
        public Input<string>? MaxBytes { get; set; }

        /// <summary>
        /// The maximum duration that can elapse before a new Cloud Storage file is created. Min 1 minute, max 10 minutes, default 5 minutes. May not exceed the subscription's acknowledgement deadline.
        /// </summary>
        [Input("maxDuration")]
        public Input<string>? MaxDuration { get; set; }

        /// <summary>
        /// If set, message data will be written to Cloud Storage in text format.
        /// </summary>
        [Input("textConfig")]
        public Input<Inputs.TextConfigArgs>? TextConfig { get; set; }

        public CloudStorageConfigArgs()
        {
        }
        public static new CloudStorageConfigArgs Empty => new CloudStorageConfigArgs();
    }
}