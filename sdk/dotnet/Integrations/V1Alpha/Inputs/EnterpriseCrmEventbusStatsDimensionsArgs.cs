// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Inputs
{

    public sealed class EnterpriseCrmEventbusStatsDimensionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// Whether to include or exclude the enums matching the regex.
        /// </summary>
        [Input("enumFilterType")]
        public Input<Pulumi.GoogleNative.Integrations.V1Alpha.EnterpriseCrmEventbusStatsDimensionsEnumFilterType>? EnumFilterType { get; set; }

        [Input("errorEnumString")]
        public Input<string>? ErrorEnumString { get; set; }

        [Input("retryAttempt")]
        public Input<Pulumi.GoogleNative.Integrations.V1Alpha.EnterpriseCrmEventbusStatsDimensionsRetryAttempt>? RetryAttempt { get; set; }

        [Input("taskName")]
        public Input<string>? TaskName { get; set; }

        [Input("taskNumber")]
        public Input<string>? TaskNumber { get; set; }

        /// <summary>
        /// Stats have been or will be aggregated on set fields for any semantically-meaningful combination.
        /// </summary>
        [Input("triggerId")]
        public Input<string>? TriggerId { get; set; }

        [Input("warningEnumString")]
        public Input<string>? WarningEnumString { get; set; }

        [Input("workflowId")]
        public Input<string>? WorkflowId { get; set; }

        [Input("workflowName")]
        public Input<string>? WorkflowName { get; set; }

        public EnterpriseCrmEventbusStatsDimensionsArgs()
        {
        }
        public static new EnterpriseCrmEventbusStatsDimensionsArgs Empty => new EnterpriseCrmEventbusStatsDimensionsArgs();
    }
}
