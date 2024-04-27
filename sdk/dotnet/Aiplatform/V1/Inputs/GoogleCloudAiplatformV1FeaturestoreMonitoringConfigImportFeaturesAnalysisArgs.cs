// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1.Inputs
{

    /// <summary>
    /// Configuration of the Featurestore's ImportFeature Analysis Based Monitoring. This type of analysis generates statistics for values of each Feature imported by every ImportFeatureValues operation.
    /// </summary>
    public sealed class GoogleCloudAiplatformV1FeaturestoreMonitoringConfigImportFeaturesAnalysisArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The baseline used to do anomaly detection for the statistics generated by import features analysis.
        /// </summary>
        [Input("anomalyDetectionBaseline")]
        public Input<Pulumi.GoogleNative.Aiplatform.V1.GoogleCloudAiplatformV1FeaturestoreMonitoringConfigImportFeaturesAnalysisAnomalyDetectionBaseline>? AnomalyDetectionBaseline { get; set; }

        /// <summary>
        /// Whether to enable / disable / inherite default hebavior for import features analysis.
        /// </summary>
        [Input("state")]
        public Input<Pulumi.GoogleNative.Aiplatform.V1.GoogleCloudAiplatformV1FeaturestoreMonitoringConfigImportFeaturesAnalysisState>? State { get; set; }

        public GoogleCloudAiplatformV1FeaturestoreMonitoringConfigImportFeaturesAnalysisArgs()
        {
        }
        public static new GoogleCloudAiplatformV1FeaturestoreMonitoringConfigImportFeaturesAnalysisArgs Empty => new GoogleCloudAiplatformV1FeaturestoreMonitoringConfigImportFeaturesAnalysisArgs();
    }
}