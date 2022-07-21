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
    /// An AwsS3CompatibleData resource.
    /// </summary>
    public sealed class AwsS3CompatibleDataArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the bucket.
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// Specifies the endpoint of the storage service.
        /// </summary>
        [Input("endpoint", required: true)]
        public Input<string> Endpoint { get; set; } = null!;

        /// <summary>
        /// Specifies the root path to transfer objects. Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should generally not begin with a '/'.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Specifies the region to sign requests with. This can be left blank if requests should be signed with an empty region.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// A S3 compatible metadata.
        /// </summary>
        [Input("s3Metadata")]
        public Input<Inputs.S3CompatibleMetadataArgs>? S3Metadata { get; set; }

        public AwsS3CompatibleDataArgs()
        {
        }
    }
}
