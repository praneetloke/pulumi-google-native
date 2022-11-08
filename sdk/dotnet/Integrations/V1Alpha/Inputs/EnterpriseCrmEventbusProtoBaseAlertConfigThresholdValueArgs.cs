// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Inputs
{

    /// <summary>
    /// The threshold value of the metric, above or below which the alert should be triggered. See EventAlertConfig or TaskAlertConfig for the different alert metric types in each case. For the *RATE metrics, one or both of these fields may be set. Zero is the default value and can be left at that. For *PERCENTILE_DURATION metrics, one or both of these fields may be set, and also, the duration threshold value should be specified in the threshold_duration_ms member below. For *AVERAGE_DURATION metrics, these fields should not be set at all. A different member, threshold_duration_ms, must be set in the EventAlertConfig or the TaskAlertConfig. See go/eventbus-alert-config-examples
    /// </summary>
    public sealed class EnterpriseCrmEventbusProtoBaseAlertConfigThresholdValueArgs : global::Pulumi.ResourceArgs
    {
        [Input("absolute")]
        public Input<string>? Absolute { get; set; }

        [Input("percentage")]
        public Input<int>? Percentage { get; set; }

        public EnterpriseCrmEventbusProtoBaseAlertConfigThresholdValueArgs()
        {
        }
        public static new EnterpriseCrmEventbusProtoBaseAlertConfigThresholdValueArgs Empty => new EnterpriseCrmEventbusProtoBaseAlertConfigThresholdValueArgs();
    }
}
