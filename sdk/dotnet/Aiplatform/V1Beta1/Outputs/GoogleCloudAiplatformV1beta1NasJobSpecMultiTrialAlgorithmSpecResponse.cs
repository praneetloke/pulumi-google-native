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
    /// The spec of multi-trial Neural Architecture Search (NAS).
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudAiplatformV1beta1NasJobSpecMultiTrialAlgorithmSpecResponse
    {
        /// <summary>
        /// Metric specs for the NAS job. Validation for this field is done at `multi_trial_algorithm_spec` field.
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1beta1NasJobSpecMultiTrialAlgorithmSpecMetricSpecResponse Metric;
        /// <summary>
        /// The multi-trial Neural Architecture Search (NAS) algorithm type. Defaults to `REINFORCEMENT_LEARNING`.
        /// </summary>
        public readonly string MultiTrialAlgorithm;
        /// <summary>
        /// Spec for search trials.
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1beta1NasJobSpecMultiTrialAlgorithmSpecSearchTrialSpecResponse SearchTrialSpec;
        /// <summary>
        /// Spec for train trials. Top N [TrainTrialSpec.max_parallel_trial_count] search trials will be trained for every M [TrainTrialSpec.frequency] trials searched.
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1beta1NasJobSpecMultiTrialAlgorithmSpecTrainTrialSpecResponse TrainTrialSpec;

        [OutputConstructor]
        private GoogleCloudAiplatformV1beta1NasJobSpecMultiTrialAlgorithmSpecResponse(
            Outputs.GoogleCloudAiplatformV1beta1NasJobSpecMultiTrialAlgorithmSpecMetricSpecResponse metric,

            string multiTrialAlgorithm,

            Outputs.GoogleCloudAiplatformV1beta1NasJobSpecMultiTrialAlgorithmSpecSearchTrialSpecResponse searchTrialSpec,

            Outputs.GoogleCloudAiplatformV1beta1NasJobSpecMultiTrialAlgorithmSpecTrainTrialSpecResponse trainTrialSpec)
        {
            Metric = metric;
            MultiTrialAlgorithm = multiTrialAlgorithm;
            SearchTrialSpec = searchTrialSpec;
            TrainTrialSpec = trainTrialSpec;
        }
    }
}