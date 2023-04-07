// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.TPU.V2.Inputs
{

    /// <summary>
    /// A TPU accelerator configuration.
    /// </summary>
    public sealed class AcceleratorConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Topology of TPU in chips.
        /// </summary>
        [Input("topology", required: true)]
        public Input<string> Topology { get; set; } = null!;

        /// <summary>
        /// Type of TPU.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.TPU.V2.AcceleratorConfigType> Type { get; set; } = null!;

        public AcceleratorConfigArgs()
        {
        }
        public static new AcceleratorConfigArgs Empty => new AcceleratorConfigArgs();
    }
}
