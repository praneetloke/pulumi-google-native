// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ServiceManagement.V1.Outputs
{

    /// <summary>
    /// Bind API methods to metrics. Binding a method to a metric causes that metric's configured quota behaviors to apply to the method call.
    /// </summary>
    [OutputType]
    public sealed class MetricRuleResponse
    {
        /// <summary>
        /// Metrics to update when the selected methods are called. The key of the map is the metric name, the value is the DynamicCostType to specify how to calculate the cost from the request. The cost amount will be increased for the metric against which the quota limits are defined. It is only implemented in CloudESF(go/cloudesf)
        /// </summary>
        public readonly ImmutableDictionary<string, string> DynamicMetricCosts;
        /// <summary>
        /// Metrics to update when the selected methods are called, and the associated cost applied to each metric. The key of the map is the metric name, and the values are the amount increased for the metric against which the quota limits are defined. The value must not be negative.
        /// </summary>
        public readonly ImmutableDictionary<string, string> MetricCosts;
        /// <summary>
        /// Selects the methods to which this rule applies. Refer to selector for syntax details.
        /// </summary>
        public readonly string Selector;

        [OutputConstructor]
        private MetricRuleResponse(
            ImmutableDictionary<string, string> dynamicMetricCosts,

            ImmutableDictionary<string, string> metricCosts,

            string selector)
        {
            DynamicMetricCosts = dynamicMetricCosts;
            MetricCosts = metricCosts;
            Selector = selector;
        }
    }
}
