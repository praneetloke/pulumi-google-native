// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Inputs
{

    /// <summary>
    /// Task scheduling and trigger settings.
    /// </summary>
    public sealed class GoogleCloudDataplexV1TaskTriggerSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Prevent the task from executing. This does not cancel already running tasks. It is intended to temporarily disable RECURRING tasks.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Optional. Number of retry attempts before aborting. Set to zero to never attempt to retry a failed task.
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// Optional. Cron schedule (https://en.wikipedia.org/wiki/Cron) for running tasks periodically. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or "TZ=${IANA_TIME_ZONE}". The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone database. For example, "CRON_TZ=America/New_York 1 * * * *", or "TZ=America/New_York 1 * * * *". This field is required for RECURRING tasks.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Optional. The first run of the task will be after this time. If not specified, the task will run shortly after being submitted if ON_DEMAND and based on the schedule if RECURRING.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// Immutable. Trigger type of the user-specified Task.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.Dataplex.V1.GoogleCloudDataplexV1TaskTriggerSpecType> Type { get; set; } = null!;

        public GoogleCloudDataplexV1TaskTriggerSpecArgs()
        {
        }
        public static new GoogleCloudDataplexV1TaskTriggerSpecArgs Empty => new GoogleCloudDataplexV1TaskTriggerSpecArgs();
    }
}
