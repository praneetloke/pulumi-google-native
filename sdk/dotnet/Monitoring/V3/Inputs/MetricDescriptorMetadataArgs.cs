// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3.Inputs
{

    /// <summary>
    /// Additional annotations that can be used to guide the usage of a metric.
    /// </summary>
    public sealed class MetricDescriptorMetadataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The delay of data points caused by ingestion. Data points older than this age are guaranteed to be ingested and available to be read, excluding data loss due to errors.
        /// </summary>
        [Input("ingestDelay")]
        public Input<string>? IngestDelay { get; set; }

        /// <summary>
        /// Deprecated. Must use the MetricDescriptor.launch_stage instead.
        /// </summary>
        [Input("launchStage")]
        public Input<Pulumi.GoogleNative.Monitoring.V3.MetricDescriptorMetadataLaunchStage>? LaunchStage { get; set; }

        /// <summary>
        /// The sampling period of metric data points. For metrics which are written periodically, consecutive data points are stored at this time interval, excluding data loss due to errors. Metrics with a higher granularity have a smaller sampling period.
        /// </summary>
        [Input("samplePeriod")]
        public Input<string>? SamplePeriod { get; set; }

        public MetricDescriptorMetadataArgs()
        {
        }
        public static new MetricDescriptorMetadataArgs Empty => new MetricDescriptorMetadataArgs();
    }
}
