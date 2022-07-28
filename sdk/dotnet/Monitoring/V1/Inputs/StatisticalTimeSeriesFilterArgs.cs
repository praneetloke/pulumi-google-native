// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V1.Inputs
{

    /// <summary>
    /// A filter that ranks streams based on their statistical relation to other streams in a request. Note: This field is deprecated and completely ignored by the API.
    /// </summary>
    public sealed class StatisticalTimeSeriesFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// How many time series to output.
        /// </summary>
        [Input("numTimeSeries")]
        public Input<int>? NumTimeSeries { get; set; }

        /// <summary>
        /// rankingMethod is applied to a set of time series, and then the produced value for each individual time series is used to compare a given time series to others. These are methods that cannot be applied stream-by-stream, but rather require the full context of a request to evaluate time series.
        /// </summary>
        [Input("rankingMethod")]
        public Input<Pulumi.GoogleNative.Monitoring.V1.StatisticalTimeSeriesFilterRankingMethod>? RankingMethod { get; set; }

        public StatisticalTimeSeriesFilterArgs()
        {
        }
        public static new StatisticalTimeSeriesFilterArgs Empty => new StatisticalTimeSeriesFilterArgs();
    }
}
