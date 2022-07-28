// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1.Inputs
{

    /// <summary>
    /// Configuration specific to Salesforce Live Agent.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigSalesforceLiveAgentConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Live Agent chat button ID.
        /// </summary>
        [Input("buttonId", required: true)]
        public Input<string> ButtonId { get; set; } = null!;

        /// <summary>
        /// Live Agent deployment ID.
        /// </summary>
        [Input("deploymentId", required: true)]
        public Input<string> DeploymentId { get; set; } = null!;

        /// <summary>
        /// Domain of the Live Agent endpoint for this agent. You can find the endpoint URL in the `Live Agent settings` page. For example if URL has the form https://d.la4-c2-phx.salesforceliveagent.com/..., you should fill in d.la4-c2-phx.salesforceliveagent.com.
        /// </summary>
        [Input("endpointDomain", required: true)]
        public Input<string> EndpointDomain { get; set; } = null!;

        /// <summary>
        /// The organization ID of the Salesforce account.
        /// </summary>
        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        public GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigSalesforceLiveAgentConfigArgs()
        {
        }
        public static new GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigSalesforceLiveAgentConfigArgs Empty => new GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigSalesforceLiveAgentConfigArgs();
    }
}
