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
    /// Cloud Scheduler Trigger configuration
    /// </summary>
    public sealed class EnterpriseCrmEventbusProtoCloudSchedulerConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cron tab of cloud scheduler trigger.
        /// </summary>
        [Input("cronTab", required: true)]
        public Input<string> CronTab { get; set; } = null!;

        /// <summary>
        /// Optional. When the job was deleted from Pantheon UI, error_message will be populated when Get/List integrations
        /// </summary>
        [Input("errorMessage")]
        public Input<string>? ErrorMessage { get; set; }

        /// <summary>
        /// The location where associated cloud scheduler job will be created
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Service account used by Cloud Scheduler to trigger the integration at scheduled time
        /// </summary>
        [Input("serviceAccountEmail", required: true)]
        public Input<string> ServiceAccountEmail { get; set; } = null!;

        public EnterpriseCrmEventbusProtoCloudSchedulerConfigArgs()
        {
        }
        public static new EnterpriseCrmEventbusProtoCloudSchedulerConfigArgs Empty => new EnterpriseCrmEventbusProtoCloudSchedulerConfigArgs();
    }
}
