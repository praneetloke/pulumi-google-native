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
    /// A filter that defines a subset of time series data that is displayed in a widget. Time series data is fetched using the ListTimeSeries (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list) method.
    /// </summary>
    public sealed class TimeSeriesFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// By default, the raw time series data is returned. Use this field to combine multiple time series for different views of the data.
        /// </summary>
        [Input("aggregation")]
        public Input<Inputs.AggregationArgs>? Aggregation { get; set; }

        /// <summary>
        /// The monitoring filter (https://cloud.google.com/monitoring/api/v3/filters) that identifies the metric types, resources, and projects to query.
        /// </summary>
        [Input("filter", required: true)]
        public Input<string> Filter { get; set; } = null!;

        /// <summary>
        /// Ranking based time series filter.
        /// </summary>
        [Input("pickTimeSeriesFilter")]
        public Input<Inputs.PickTimeSeriesFilterArgs>? PickTimeSeriesFilter { get; set; }

        /// <summary>
        /// Apply a second aggregation after aggregation is applied.
        /// </summary>
        [Input("secondaryAggregation")]
        public Input<Inputs.AggregationArgs>? SecondaryAggregation { get; set; }

        /// <summary>
        /// Statistics based time series filter. Note: This field is deprecated and completely ignored by the API.
        /// </summary>
        [Input("statisticalTimeSeriesFilter")]
        public Input<Inputs.StatisticalTimeSeriesFilterArgs>? StatisticalTimeSeriesFilter { get; set; }

        public TimeSeriesFilterArgs()
        {
        }
        public static new TimeSeriesFilterArgs Empty => new TimeSeriesFilterArgs();
    }
}
