// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Outputs
{

    /// <summary>
    /// Policy that defines the task retry logic and failure type. If no FailurePolicy is defined for a task, all its dependent tasks will not be executed (i.e, a `retry_strategy` of NONE will be applied).
    /// </summary>
    [OutputType]
    public sealed class EnterpriseCrmEventbusProtoFailurePolicyResponse
    {
        /// <summary>
        /// Required if retry_strategy is FIXED_INTERVAL or LINEAR/EXPONENTIAL_BACKOFF/RESTART_WORKFLOW_WITH_BACKOFF. Defines the initial interval for backoff.
        /// </summary>
        public readonly string IntervalInSeconds;
        /// <summary>
        /// Required if retry_strategy is FIXED_INTERVAL or LINEAR/EXPONENTIAL_BACKOFF/RESTART_WORKFLOW_WITH_BACKOFF. Defines the number of times the task will be retried if failed.
        /// </summary>
        public readonly int MaxNumRetries;
        /// <summary>
        /// Defines what happens to the task upon failure.
        /// </summary>
        public readonly string RetryStrategy;

        [OutputConstructor]
        private EnterpriseCrmEventbusProtoFailurePolicyResponse(
            string intervalInSeconds,

            int maxNumRetries,

            string retryStrategy)
        {
            IntervalInSeconds = intervalInSeconds;
            MaxNumRetries = maxNumRetries;
            RetryStrategy = retryStrategy;
        }
    }
}