// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// The parameters of the raw disk image.
    /// </summary>
    public sealed class ImageRawDiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The format used to encode and transmit the block device, which should be TAR. This is just a container and transmission format and not a runtime format. Provided by the client when the disk image is created.
        /// </summary>
        [Input("containerType")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.ImageRawDiskContainerType>? ContainerType { get; set; }

        /// <summary>
        /// [Deprecated] This field is deprecated. An optional SHA1 checksum of the disk image before unpackaging provided by the client when the disk image is created.
        /// </summary>
        [Input("sha1Checksum")]
        public Input<string>? Sha1Checksum { get; set; }

        /// <summary>
        /// The full Google Cloud Storage URL where the raw disk image archive is stored. The following are valid formats for the URL: - https://storage.googleapis.com/bucket_name/image_archive_name - https://storage.googleapis.com/bucket_name/folder_name/ image_archive_name In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        public ImageRawDiskArgs()
        {
        }
        public static new ImageRawDiskArgs Empty => new ImageRawDiskArgs();
    }
}
