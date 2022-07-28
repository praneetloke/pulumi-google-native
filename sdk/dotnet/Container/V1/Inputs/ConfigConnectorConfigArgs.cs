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
    /// Configuration options for the Config Connector add-on.
    /// </summary>
    public sealed class ConfigConnectorConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether Cloud Connector is enabled for this cluster.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public ConfigConnectorConfigArgs()
        {
        }
        public static new ConfigConnectorConfigArgs Empty => new ConfigConnectorConfigArgs();
    }
}
