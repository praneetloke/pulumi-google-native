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
    /// Assigns the input data to training, validation, and test sets as per the given fractions. Any of `training_fraction`, `validation_fraction` and `test_fraction` may optionally be provided, they must sum to up to 1. If the provided ones sum to less than 1, the remainder is assigned to sets as decided by Vertex AI. If none of the fractions are set, by default roughly 80% of data is used for training, 10% for validation, and 10% for test.
    /// </summary>
    public sealed class GoogleCloudAiplatformV1FractionSplitArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The fraction of the input data that is to be used to evaluate the Model.
        /// </summary>
        [Input("testFraction")]
        public Input<double>? TestFraction { get; set; }

        /// <summary>
        /// The fraction of the input data that is to be used to train the Model.
        /// </summary>
        [Input("trainingFraction")]
        public Input<double>? TrainingFraction { get; set; }

        /// <summary>
        /// The fraction of the input data that is to be used to validate the Model.
        /// </summary>
        [Input("validationFraction")]
        public Input<double>? ValidationFraction { get; set; }

        public GoogleCloudAiplatformV1FractionSplitArgs()
        {
        }
        public static new GoogleCloudAiplatformV1FractionSplitArgs Empty => new GoogleCloudAiplatformV1FractionSplitArgs();
    }
}