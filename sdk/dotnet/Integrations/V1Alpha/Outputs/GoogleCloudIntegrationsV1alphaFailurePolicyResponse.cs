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
    public sealed class GoogleCloudIntegrationsV1alphaFailurePolicyResponse
    {
        /// <summary>
        /// Required if retry_strategy is FIXED_INTERVAL or LINEAR/EXPONENTIAL_BACKOFF/RESTART_INTEGRATION_WITH_BACKOFF. Defines the initial interval in seconds for backoff.
        /// </summary>
        public readonly string IntervalTime;
        /// <summary>
        /// Required if retry_strategy is FIXED_INTERVAL or LINEAR/EXPONENTIAL_BACKOFF/RESTART_INTEGRATION_WITH_BACKOFF. Defines the number of times the task will be retried if failed.
        /// </summary>
        public readonly int MaxRetries;
        /// <summary>
        /// Defines what happens to the task upon failure.
        /// </summary>
        public readonly string RetryStrategy;

        [OutputConstructor]
        private GoogleCloudIntegrationsV1alphaFailurePolicyResponse(
            string intervalTime,

            int maxRetries,

            string retryStrategy)
        {
            IntervalTime = intervalTime;
            MaxRetries = maxRetries;
            RetryStrategy = retryStrategy;
        }
    }
}
