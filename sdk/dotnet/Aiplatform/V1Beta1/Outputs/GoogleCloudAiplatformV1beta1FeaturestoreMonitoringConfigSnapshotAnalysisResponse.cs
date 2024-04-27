// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1Beta1.Outputs
{

    /// <summary>
    /// Configuration of the Featurestore's Snapshot Analysis Based Monitoring. This type of analysis generates statistics for each Feature based on a snapshot of the latest feature value of each entities every monitoring_interval.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudAiplatformV1beta1FeaturestoreMonitoringConfigSnapshotAnalysisResponse
    {
        /// <summary>
        /// The monitoring schedule for snapshot analysis. For EntityType-level config: unset / disabled = true indicates disabled by default for Features under it; otherwise by default enable snapshot analysis monitoring with monitoring_interval for Features under it. Feature-level config: disabled = true indicates disabled regardless of the EntityType-level config; unset monitoring_interval indicates going with EntityType-level config; otherwise run snapshot analysis monitoring with monitoring_interval regardless of the EntityType-level config. Explicitly Disable the snapshot analysis based monitoring.
        /// </summary>
        public readonly bool Disabled;
        /// <summary>
        /// Configuration of the snapshot analysis based monitoring pipeline running interval. The value is rolled up to full day. If both monitoring_interval_days and the deprecated `monitoring_interval` field are set when creating/updating EntityTypes/Features, monitoring_interval_days will be used.
        /// </summary>
        public readonly string MonitoringInterval;
        /// <summary>
        /// Configuration of the snapshot analysis based monitoring pipeline running interval. The value indicates number of days.
        /// </summary>
        public readonly int MonitoringIntervalDays;
        /// <summary>
        /// Customized export features time window for snapshot analysis. Unit is one day. Default value is 3 weeks. Minimum value is 1 day. Maximum value is 4000 days.
        /// </summary>
        public readonly int StalenessDays;

        [OutputConstructor]
        private GoogleCloudAiplatformV1beta1FeaturestoreMonitoringConfigSnapshotAnalysisResponse(
            bool disabled,

            string monitoringInterval,

            int monitoringIntervalDays,

            int stalenessDays)
        {
            Disabled = disabled;
            MonitoringInterval = monitoringInterval;
            MonitoringIntervalDays = monitoringIntervalDays;
            StalenessDays = stalenessDays;
        }
    }
}