// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1.Inputs
{

    /// <summary>
    /// Utilization metrics values for a single VM.
    /// </summary>
    public sealed class VmUtilizationMetricsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Average CPU usage, percent.
        /// </summary>
        [Input("cpuAveragePercent")]
        public Input<int>? CpuAveragePercent { get; set; }

        /// <summary>
        /// Max CPU usage, percent.
        /// </summary>
        [Input("cpuMaxPercent")]
        public Input<int>? CpuMaxPercent { get; set; }

        /// <summary>
        /// Average disk IO rate, in kilobytes per second.
        /// </summary>
        [Input("diskIoRateAverageKbps")]
        public Input<string>? DiskIoRateAverageKbps { get; set; }

        /// <summary>
        /// Max disk IO rate, in kilobytes per second.
        /// </summary>
        [Input("diskIoRateMaxKbps")]
        public Input<string>? DiskIoRateMaxKbps { get; set; }

        /// <summary>
        /// Average memory usage, percent.
        /// </summary>
        [Input("memoryAveragePercent")]
        public Input<int>? MemoryAveragePercent { get; set; }

        /// <summary>
        /// Max memory usage, percent.
        /// </summary>
        [Input("memoryMaxPercent")]
        public Input<int>? MemoryMaxPercent { get; set; }

        /// <summary>
        /// Average network throughput (combined transmit-rates and receive-rates), in kilobytes per second.
        /// </summary>
        [Input("networkThroughputAverageKbps")]
        public Input<string>? NetworkThroughputAverageKbps { get; set; }

        /// <summary>
        /// Max network throughput (combined transmit-rates and receive-rates), in kilobytes per second.
        /// </summary>
        [Input("networkThroughputMaxKbps")]
        public Input<string>? NetworkThroughputMaxKbps { get; set; }

        public VmUtilizationMetricsArgs()
        {
        }
        public static new VmUtilizationMetricsArgs Empty => new VmUtilizationMetricsArgs();
    }
}
