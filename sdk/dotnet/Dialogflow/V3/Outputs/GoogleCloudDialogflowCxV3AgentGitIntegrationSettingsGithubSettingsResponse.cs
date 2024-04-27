// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Outputs
{

    /// <summary>
    /// Settings of integration with GitHub.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3AgentGitIntegrationSettingsGithubSettingsResponse
    {
        /// <summary>
        /// The access token used to authenticate the access to the GitHub repository.
        /// </summary>
        public readonly string AccessToken;
        /// <summary>
        /// A list of branches configured to be used from Dialogflow.
        /// </summary>
        public readonly ImmutableArray<string> Branches;
        /// <summary>
        /// The unique repository display name for the GitHub repository.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The GitHub repository URI related to the agent.
        /// </summary>
        public readonly string RepositoryUri;
        /// <summary>
        /// The branch of the GitHub repository tracked for this agent.
        /// </summary>
        public readonly string TrackingBranch;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3AgentGitIntegrationSettingsGithubSettingsResponse(
            string accessToken,

            ImmutableArray<string> branches,

            string displayName,

            string repositoryUri,

            string trackingBranch)
        {
            AccessToken = accessToken;
            Branches = branches;
            DisplayName = displayName;
            RepositoryUri = repositoryUri;
            TrackingBranch = trackingBranch;
        }
    }
}