// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1.Inputs
{

    /// <summary>
    /// Machine resources for a version.
    /// </summary>
    public sealed class ResourcesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of CPU cores needed.
        /// </summary>
        [Input("cpu")]
        public Input<double>? Cpu { get; set; }

        /// <summary>
        /// Disk size (GB) needed.
        /// </summary>
        [Input("diskGb")]
        public Input<double>? DiskGb { get; set; }

        /// <summary>
        /// The name of the encryption key that is stored in Google Cloud KMS. Only should be used by Cloud Composer to encrypt the vm disk
        /// </summary>
        [Input("kmsKeyReference")]
        public Input<string>? KmsKeyReference { get; set; }

        /// <summary>
        /// Memory (GB) needed.
        /// </summary>
        [Input("memoryGb")]
        public Input<double>? MemoryGb { get; set; }

        [Input("volumes")]
        private InputList<Inputs.VolumeArgs>? _volumes;

        /// <summary>
        /// User specified volumes.
        /// </summary>
        public InputList<Inputs.VolumeArgs> Volumes
        {
            get => _volumes ?? (_volumes = new InputList<Inputs.VolumeArgs>());
            set => _volumes = value;
        }

        public ResourcesArgs()
        {
        }
        public static new ResourcesArgs Empty => new ResourcesArgs();
    }
}
