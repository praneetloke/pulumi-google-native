// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.PolicySimulator.V1Beta1
{
    public static class GetOrganizationReplay
    {
        /// <summary>
        /// Gets the specified Replay. Each `Replay` is available for at least 7 days.
        /// </summary>
        public static Task<GetOrganizationReplayResult> InvokeAsync(GetOrganizationReplayArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationReplayResult>("google-native:policysimulator/v1beta1:getOrganizationReplay", args ?? new GetOrganizationReplayArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified Replay. Each `Replay` is available for at least 7 days.
        /// </summary>
        public static Output<GetOrganizationReplayResult> Invoke(GetOrganizationReplayInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationReplayResult>("google-native:policysimulator/v1beta1:getOrganizationReplay", args ?? new GetOrganizationReplayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrganizationReplayArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        [Input("replayId", required: true)]
        public string ReplayId { get; set; } = null!;

        public GetOrganizationReplayArgs()
        {
        }
        public static new GetOrganizationReplayArgs Empty => new GetOrganizationReplayArgs();
    }

    public sealed class GetOrganizationReplayInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        [Input("replayId", required: true)]
        public Input<string> ReplayId { get; set; } = null!;

        public GetOrganizationReplayInvokeArgs()
        {
        }
        public static new GetOrganizationReplayInvokeArgs Empty => new GetOrganizationReplayInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrganizationReplayResult
    {
        /// <summary>
        /// The configuration used for the `Replay`.
        /// </summary>
        public readonly Outputs.GoogleCloudPolicysimulatorV1beta1ReplayConfigResponse Config;
        /// <summary>
        /// The resource name of the `Replay`, which has the following format: `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}`, where `{resource-id}` is the ID of the project, folder, or organization that owns the Replay. Example: `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Summary statistics about the replayed log entries.
        /// </summary>
        public readonly Outputs.GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponse ResultsSummary;
        /// <summary>
        /// The current state of the `Replay`.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetOrganizationReplayResult(
            Outputs.GoogleCloudPolicysimulatorV1beta1ReplayConfigResponse config,

            string name,

            Outputs.GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponse resultsSummary,

            string state)
        {
            Config = config;
            Name = name;
            ResultsSummary = resultsSummary;
            State = state;
        }
    }
}
