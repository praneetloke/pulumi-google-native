// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1.Outputs
{

    /// <summary>
    /// Defines configuration for DRM systems in use.
    /// </summary>
    [OutputType]
    public sealed class DrmSystemsResponse
    {
        /// <summary>
        /// Clearkey configuration.
        /// </summary>
        public readonly Outputs.ClearkeyResponse Clearkey;
        /// <summary>
        /// Fairplay configuration.
        /// </summary>
        public readonly Outputs.FairplayResponse Fairplay;
        /// <summary>
        /// Playready configuration.
        /// </summary>
        public readonly Outputs.PlayreadyResponse Playready;
        /// <summary>
        /// Widevine configuration.
        /// </summary>
        public readonly Outputs.WidevineResponse Widevine;

        [OutputConstructor]
        private DrmSystemsResponse(
            Outputs.ClearkeyResponse clearkey,

            Outputs.FairplayResponse fairplay,

            Outputs.PlayreadyResponse playready,

            Outputs.WidevineResponse widevine)
        {
            Clearkey = clearkey;
            Fairplay = fairplay;
            Playready = playready;
            Widevine = widevine;
        }
    }
}