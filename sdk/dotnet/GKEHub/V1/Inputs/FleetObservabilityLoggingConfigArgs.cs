// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1.Inputs
{

    /// <summary>
    /// LoggingConfig defines the configuration for different types of logs.
    /// </summary>
    public sealed class FleetObservabilityLoggingConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specified if applying the default routing config to logs not specified in other configs.
        /// </summary>
        [Input("defaultConfig")]
        public Input<Inputs.FleetObservabilityRoutingConfigArgs>? DefaultConfig { get; set; }

        /// <summary>
        /// Specified if applying the routing config to all logs for all fleet scopes.
        /// </summary>
        [Input("fleetScopeLogsConfig")]
        public Input<Inputs.FleetObservabilityRoutingConfigArgs>? FleetScopeLogsConfig { get; set; }

        public FleetObservabilityLoggingConfigArgs()
        {
        }
        public static new FleetObservabilityLoggingConfigArgs Empty => new FleetObservabilityLoggingConfigArgs();
    }
}