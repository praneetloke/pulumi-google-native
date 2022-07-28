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
    /// Configuration specific to LivePerson (https://www.liveperson.com).
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigLivePersonConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Account number of the LivePerson account to connect. This is the account number you input at the login page.
        /// </summary>
        [Input("accountNumber", required: true)]
        public Input<string> AccountNumber { get; set; } = null!;

        public GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigLivePersonConfigArgs()
        {
        }
        public static new GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigLivePersonConfigArgs Empty => new GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigLivePersonConfigArgs();
    }
}
