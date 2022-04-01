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
    /// A filter to reduce the amount of data charted in relevant widgets.
    /// </summary>
    public sealed class DashboardFilterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The specified filter type
        /// </summary>
        [Input("filterType")]
        public Input<Pulumi.GoogleNative.Monitoring.V1.DashboardFilterFilterType>? FilterType { get; set; }

        /// <summary>
        /// The key for the label
        /// </summary>
        [Input("labelKey", required: true)]
        public Input<string> LabelKey { get; set; } = null!;

        /// <summary>
        /// A variable-length string value.
        /// </summary>
        [Input("stringValue")]
        public Input<string>? StringValue { get; set; }

        /// <summary>
        /// The placeholder text that can be referenced in a filter string or MQL query. If omitted, the dashboard filter will be applied to all relevant widgets in the dashboard.
        /// </summary>
        [Input("templateVariable")]
        public Input<string>? TemplateVariable { get; set; }

        public DashboardFilterArgs()
        {
        }
    }
}
