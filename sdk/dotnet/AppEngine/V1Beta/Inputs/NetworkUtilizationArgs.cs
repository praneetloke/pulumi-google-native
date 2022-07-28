// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta.Inputs
{

    /// <summary>
    /// Target scaling by network usage. Only applicable in the App Engine flexible environment.
    /// </summary>
    public sealed class NetworkUtilizationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Target bytes received per second.
        /// </summary>
        [Input("targetReceivedBytesPerSecond")]
        public Input<int>? TargetReceivedBytesPerSecond { get; set; }

        /// <summary>
        /// Target packets received per second.
        /// </summary>
        [Input("targetReceivedPacketsPerSecond")]
        public Input<int>? TargetReceivedPacketsPerSecond { get; set; }

        /// <summary>
        /// Target bytes sent per second.
        /// </summary>
        [Input("targetSentBytesPerSecond")]
        public Input<int>? TargetSentBytesPerSecond { get; set; }

        /// <summary>
        /// Target packets sent per second.
        /// </summary>
        [Input("targetSentPacketsPerSecond")]
        public Input<int>? TargetSentPacketsPerSecond { get; set; }

        public NetworkUtilizationArgs()
        {
        }
        public static new NetworkUtilizationArgs Empty => new NetworkUtilizationArgs();
    }
}
