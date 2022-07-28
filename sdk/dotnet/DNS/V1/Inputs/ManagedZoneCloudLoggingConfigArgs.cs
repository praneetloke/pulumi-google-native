// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1.Inputs
{

    /// <summary>
    /// Cloud Logging configurations for publicly visible zones.
    /// </summary>
    public sealed class ManagedZoneCloudLoggingConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, enable query logging for this ManagedZone. False by default, making logging opt-in.
        /// </summary>
        [Input("enableLogging")]
        public Input<bool>? EnableLogging { get; set; }

        [Input("kind")]
        public Input<string>? Kind { get; set; }

        public ManagedZoneCloudLoggingConfigArgs()
        {
        }
        public static new ManagedZoneCloudLoggingConfigArgs Empty => new ManagedZoneCloudLoggingConfigArgs();
    }
}
