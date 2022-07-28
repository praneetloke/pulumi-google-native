// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1Alpha2.Inputs
{

    /// <summary>
    /// WorkerConfig defines the configuration to be used for a creating workers in the pool.
    /// </summary>
    public sealed class WorkerConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Size of the disk attached to the worker, in GB. See https://cloud.google.com/compute/docs/disks/ If `0` is specified, Cloud Build will use a standard disk size.
        /// </summary>
        [Input("diskSizeGb")]
        public Input<string>? DiskSizeGb { get; set; }

        /// <summary>
        /// Machine Type of the worker, such as n1-standard-1. See https://cloud.google.com/compute/docs/machine-types. If left blank, Cloud Build will use a standard unspecified machine to create the worker pool.
        /// </summary>
        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        public WorkerConfigArgs()
        {
        }
        public static new WorkerConfigArgs Empty => new WorkerConfigArgs();
    }
}
