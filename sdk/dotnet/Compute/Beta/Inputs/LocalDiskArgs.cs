// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    public sealed class LocalDiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the number of such disks.
        /// </summary>
        [Input("diskCount")]
        public Input<int>? DiskCount { get; set; }

        /// <summary>
        /// Specifies the size of the disk in base-2 GB.
        /// </summary>
        [Input("diskSizeGb")]
        public Input<int>? DiskSizeGb { get; set; }

        /// <summary>
        /// Specifies the desired disk type on the node. This disk type must be a local storage type (e.g.: local-ssd). Note that for nodeTemplates, this should be the name of the disk type and not its URL.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        public LocalDiskArgs()
        {
        }
        public static new LocalDiskArgs Empty => new LocalDiskArgs();
    }
}
