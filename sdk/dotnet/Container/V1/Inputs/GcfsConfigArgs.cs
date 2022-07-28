// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Inputs
{

    /// <summary>
    /// GcfsConfig contains configurations of Google Container File System (image streaming).
    /// </summary>
    public sealed class GcfsConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to use GCFS.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public GcfsConfigArgs()
        {
        }
        public static new GcfsConfigArgs Empty => new GcfsConfigArgs();
    }
}
