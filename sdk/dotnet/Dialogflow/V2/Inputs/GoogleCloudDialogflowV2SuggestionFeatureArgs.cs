// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2.Inputs
{

    /// <summary>
    /// The type of Human Agent Assistant API suggestion to perform, and the maximum number of results to return for that type. Multiple `Feature` objects can be specified in the `features` list.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2SuggestionFeatureArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of Human Agent Assistant API feature to request.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Dialogflow.V2.GoogleCloudDialogflowV2SuggestionFeatureType>? Type { get; set; }

        public GoogleCloudDialogflowV2SuggestionFeatureArgs()
        {
        }
        public static new GoogleCloudDialogflowV2SuggestionFeatureArgs Empty => new GoogleCloudDialogflowV2SuggestionFeatureArgs();
    }
}
