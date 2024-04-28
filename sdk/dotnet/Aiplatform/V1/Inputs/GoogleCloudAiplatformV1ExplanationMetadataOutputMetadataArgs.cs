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
    /// Metadata of the prediction output to be explained.
    /// </summary>
    public sealed class GoogleCloudAiplatformV1ExplanationMetadataOutputMetadataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify a field name in the prediction to look for the display name. Use this if the prediction contains the display names for the outputs. The display names in the prediction must have the same shape of the outputs, so that it can be located by Attribution.output_index for a specific output.
        /// </summary>
        [Input("displayNameMappingKey")]
        public Input<string>? DisplayNameMappingKey { get; set; }

        /// <summary>
        /// Static mapping between the index and display name. Use this if the outputs are a deterministic n-dimensional array, e.g. a list of scores of all the classes in a pre-defined order for a multi-classification Model. It's not feasible if the outputs are non-deterministic, e.g. the Model produces top-k classes or sort the outputs by their values. The shape of the value must be an n-dimensional array of strings. The number of dimensions must match that of the outputs to be explained. The Attribution.output_display_name is populated by locating in the mapping with Attribution.output_index.
        /// </summary>
        [Input("indexDisplayNameMapping")]
        public Input<object>? IndexDisplayNameMapping { get; set; }

        /// <summary>
        /// Name of the output tensor. Required and is only applicable to Vertex AI provided images for Tensorflow.
        /// </summary>
        [Input("outputTensorName")]
        public Input<string>? OutputTensorName { get; set; }

        public GoogleCloudAiplatformV1ExplanationMetadataOutputMetadataArgs()
        {
        }
        public static new GoogleCloudAiplatformV1ExplanationMetadataOutputMetadataArgs Empty => new GoogleCloudAiplatformV1ExplanationMetadataOutputMetadataArgs();
    }
}
