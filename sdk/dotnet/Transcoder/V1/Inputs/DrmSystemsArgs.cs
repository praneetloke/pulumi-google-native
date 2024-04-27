// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1.Inputs
{

    /// <summary>
    /// Defines configuration for DRM systems in use.
    /// </summary>
    public sealed class DrmSystemsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Clearkey configuration.
        /// </summary>
        [Input("clearkey")]
        public Input<Inputs.ClearkeyArgs>? Clearkey { get; set; }

        /// <summary>
        /// Fairplay configuration.
        /// </summary>
        [Input("fairplay")]
        public Input<Inputs.FairplayArgs>? Fairplay { get; set; }

        /// <summary>
        /// Playready configuration.
        /// </summary>
        [Input("playready")]
        public Input<Inputs.PlayreadyArgs>? Playready { get; set; }

        /// <summary>
        /// Widevine configuration.
        /// </summary>
        [Input("widevine")]
        public Input<Inputs.WidevineArgs>? Widevine { get; set; }

        public DrmSystemsArgs()
        {
        }
        public static new DrmSystemsArgs Empty => new DrmSystemsArgs();
    }
}