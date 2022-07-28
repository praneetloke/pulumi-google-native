// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ServiceManagement.V1.Inputs
{

    /// <summary>
    /// Configuration of a specific monitoring destination (the producer project or the consumer project).
    /// </summary>
    public sealed class MonitoringDestinationArgs : global::Pulumi.ResourceArgs
    {
        [Input("metrics")]
        private InputList<string>? _metrics;

        /// <summary>
        /// Types of the metrics to report to this monitoring destination. Each type must be defined in Service.metrics section.
        /// </summary>
        public InputList<string> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<string>());
            set => _metrics = value;
        }

        /// <summary>
        /// The monitored resource type. The type must be defined in Service.monitored_resources section.
        /// </summary>
        [Input("monitoredResource")]
        public Input<string>? MonitoredResource { get; set; }

        public MonitoringDestinationArgs()
        {
        }
        public static new MonitoringDestinationArgs Empty => new MonitoringDestinationArgs();
    }
}
