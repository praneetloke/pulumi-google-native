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
    /// TimeSeriesQuery collects the set of supported methods for querying time series data from the Stackdriver metrics API.
    /// </summary>
    public sealed class TimeSeriesQueryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Filter parameters to fetch time series.
        /// </summary>
        [Input("timeSeriesFilter")]
        public Input<Inputs.TimeSeriesFilterArgs>? TimeSeriesFilter { get; set; }

        /// <summary>
        /// Parameters to fetch a ratio between two time series filters.
        /// </summary>
        [Input("timeSeriesFilterRatio")]
        public Input<Inputs.TimeSeriesFilterRatioArgs>? TimeSeriesFilterRatio { get; set; }

        /// <summary>
        /// A query used to fetch time series with MQL.
        /// </summary>
        [Input("timeSeriesQueryLanguage")]
        public Input<string>? TimeSeriesQueryLanguage { get; set; }

        /// <summary>
        /// The unit of data contained in fetched time series. If non-empty, this unit will override any unit that accompanies fetched data. The format is the same as the unit (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors) field in MetricDescriptor.
        /// </summary>
        [Input("unitOverride")]
        public Input<string>? UnitOverride { get; set; }

        public TimeSeriesQueryArgs()
        {
        }
        public static new TimeSeriesQueryArgs Empty => new TimeSeriesQueryArgs();
    }
}
