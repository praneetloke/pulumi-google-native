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
    /// Cloud Scheduler Trigger configuration
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudIntegrationsV1alphaCloudSchedulerConfigResponse
    {
        /// <summary>
        /// The cron tab of cloud scheduler trigger.
        /// </summary>
        public readonly string CronTab;
        /// <summary>
        /// Optional. When the job was deleted from Pantheon UI, error_message will be populated when Get/List integrations
        /// </summary>
        public readonly string ErrorMessage;
        /// <summary>
        /// The location where associated cloud scheduler job will be created
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Service account used by Cloud Scheduler to trigger the integration at scheduled time
        /// </summary>
        public readonly string ServiceAccountEmail;

        [OutputConstructor]
        private GoogleCloudIntegrationsV1alphaCloudSchedulerConfigResponse(
            string cronTab,

            string errorMessage,

            string location,

            string serviceAccountEmail)
        {
            CronTab = cronTab;
            ErrorMessage = errorMessage;
            Location = location;
            ServiceAccountEmail = serviceAccountEmail;
        }
    }
}
