// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// A set of Display Device options
    /// </summary>
    public sealed class DisplayDeviceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines whether the instance has Display enabled.
        /// </summary>
        [Input("enableDisplay")]
        public Input<bool>? EnableDisplay { get; set; }

        public DisplayDeviceArgs()
        {
        }
        public static new DisplayDeviceArgs Empty => new DisplayDeviceArgs();
    }
}
