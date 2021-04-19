// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudDialogflowV2beta1FulfillmentResponse
    {
        /// <summary>
        /// The human-readable name of the fulfillment, unique within the agent. This field is not used for Fulfillment in an Environment.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Whether fulfillment is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The field defines whether the fulfillment is enabled for certain features.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowV2beta1FulfillmentFeatureResponse> Features;
        /// <summary>
        /// Configuration for a generic web service.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2beta1FulfillmentGenericWebServiceResponse GenericWebService;
        /// <summary>
        /// Required. The unique identifier of the fulfillment. Supported formats: - `projects//agent/fulfillment` - `projects//locations//agent/fulfillment` This field is not used for Fulfillment in an Environment.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GoogleCloudDialogflowV2beta1FulfillmentResponse(
            string displayName,

            bool enabled,

            ImmutableArray<Outputs.GoogleCloudDialogflowV2beta1FulfillmentFeatureResponse> features,

            Outputs.GoogleCloudDialogflowV2beta1FulfillmentGenericWebServiceResponse genericWebService,

            string name)
        {
            DisplayName = displayName;
            Enabled = enabled;
            Features = features;
            GenericWebService = genericWebService;
            Name = name;
        }
    }
}
