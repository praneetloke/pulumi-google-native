// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Outputs
{

    [OutputType]
    public sealed class EnterpriseCrmEventbusStatsDimensionsResponse
    {
        public readonly string ClientId;
        /// <summary>
        /// Whether to include or exclude the enums matching the regex.
        /// </summary>
        public readonly string EnumFilterType;
        public readonly string ErrorEnumString;
        public readonly string RetryAttempt;
        public readonly string TaskName;
        public readonly string TaskNumber;
        /// <summary>
        /// Stats have been or will be aggregated on set fields for any semantically-meaningful combination.
        /// </summary>
        public readonly string TriggerId;
        public readonly string WarningEnumString;
        public readonly string WorkflowId;
        public readonly string WorkflowName;

        [OutputConstructor]
        private EnterpriseCrmEventbusStatsDimensionsResponse(
            string clientId,

            string enumFilterType,

            string errorEnumString,

            string retryAttempt,

            string taskName,

            string taskNumber,

            string triggerId,

            string warningEnumString,

            string workflowId,

            string workflowName)
        {
            ClientId = clientId;
            EnumFilterType = enumFilterType;
            ErrorEnumString = errorEnumString;
            RetryAttempt = retryAttempt;
            TaskName = taskName;
            TaskNumber = taskNumber;
            TriggerId = triggerId;
            WarningEnumString = warningEnumString;
            WorkflowId = workflowId;
            WorkflowName = workflowName;
        }
    }
}
