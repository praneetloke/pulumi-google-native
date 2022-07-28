// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Storage.V1.Inputs
{

    /// <summary>
    /// The bucket's Autoclass configuration.
    /// </summary>
    public sealed class BucketAutoclassArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not Autoclass is enabled on this bucket
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// A date and time in RFC 3339 format representing the instant at which "enabled" was last toggled.
        /// </summary>
        [Input("toggleTime")]
        public Input<string>? ToggleTime { get; set; }

        public BucketAutoclassArgs()
        {
        }
        public static new BucketAutoclassArgs Empty => new BucketAutoclassArgs();
    }
}
