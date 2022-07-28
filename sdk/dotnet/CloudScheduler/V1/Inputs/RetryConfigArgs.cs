// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudScheduler.V1.Inputs
{

    /// <summary>
    /// Settings that determine the retry behavior. By default, if a job does not complete successfully (meaning that an acknowledgement is not received from the handler, then it will be retried with exponential backoff according to the settings in RetryConfig.
    /// </summary>
    public sealed class RetryConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum amount of time to wait before retrying a job after it fails. The default value of this field is 1 hour.
        /// </summary>
        [Input("maxBackoffDuration")]
        public Input<string>? MaxBackoffDuration { get; set; }

        /// <summary>
        /// The time between retries will double `max_doublings` times. A job's retry interval starts at min_backoff_duration, then doubles `max_doublings` times, then increases linearly, and finally retries at intervals of max_backoff_duration up to retry_count times. For example, if min_backoff_duration is 10s, max_backoff_duration is 300s, and `max_doublings` is 3, then the a job will first be retried in 10s. The retry interval will double three times, and then increase linearly by 2^3 * 10s. Finally, the job will retry at intervals of max_backoff_duration until the job has been attempted retry_count times. Thus, the requests will retry at 10s, 20s, 40s, 80s, 160s, 240s, 300s, 300s, .... The default value of this field is 5.
        /// </summary>
        [Input("maxDoublings")]
        public Input<int>? MaxDoublings { get; set; }

        /// <summary>
        /// The time limit for retrying a failed job, measured from time when an execution was first attempted. If specified with retry_count, the job will be retried until both limits are reached. The default value for max_retry_duration is zero, which means retry duration is unlimited.
        /// </summary>
        [Input("maxRetryDuration")]
        public Input<string>? MaxRetryDuration { get; set; }

        /// <summary>
        /// The minimum amount of time to wait before retrying a job after it fails. The default value of this field is 5 seconds.
        /// </summary>
        [Input("minBackoffDuration")]
        public Input<string>? MinBackoffDuration { get; set; }

        /// <summary>
        /// The number of attempts that the system will make to run a job using the exponential backoff procedure described by max_doublings. The default value of retry_count is zero. If retry_count is zero, a job attempt will *not* be retried if it fails. Instead the Cloud Scheduler system will wait for the next scheduled execution time. If retry_count is set to a non-zero number then Cloud Scheduler will retry failed attempts, using exponential backoff, retry_count times, or until the next scheduled execution time, whichever comes first. Values greater than 5 and negative values are not allowed.
        /// </summary>
        [Input("retryCount")]
        public Input<int>? RetryCount { get; set; }

        public RetryConfigArgs()
        {
        }
        public static new RetryConfigArgs Empty => new RetryConfigArgs();
    }
}
