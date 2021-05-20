// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1
{
    /// <summary>
    /// Creates an Experiment in the specified Environment.
    /// </summary>
    [GoogleNativeResourceType("google-native:dialogflow/v3beta1:AgentEnvironmentExperiment")]
    public partial class AgentEnvironmentExperiment : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation time of this experiment.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The definition of the experiment.
        /// </summary>
        [Output("definition")]
        public Output<Outputs.GoogleCloudDialogflowCxV3beta1ExperimentDefinitionResponse> Definition { get; private set; } = null!;

        /// <summary>
        /// The human-readable description of the experiment.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Required. The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// End time of this experiment.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// Maximum number of days to run the experiment. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
        /// </summary>
        [Output("experimentLength")]
        public Output<string> ExperimentLength { get; private set; } = null!;

        /// <summary>
        /// Last update time of this experiment.
        /// </summary>
        [Output("lastUpdateTime")]
        public Output<string> LastUpdateTime { get; private set; } = null!;

        /// <summary>
        /// The name of the experiment. Format: projects//locations//agents//environments//experiments/..
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Inference result of the experiment.
        /// </summary>
        [Output("result")]
        public Output<Outputs.GoogleCloudDialogflowCxV3beta1ExperimentResultResponse> Result { get; private set; } = null!;

        /// <summary>
        /// Start time of this experiment.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// The current state of the experiment. Transition triggered by Expriments.StartExperiment: PENDING-&gt;RUNNING. Transition triggered by Expriments.CancelExperiment: PENDING-&gt;CANCELLED or RUNNING-&gt;CANCELLED.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The history of updates to the experiment variants.
        /// </summary>
        [Output("variantsHistory")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1VariantsHistoryResponse>> VariantsHistory { get; private set; } = null!;


        /// <summary>
        /// Create a AgentEnvironmentExperiment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AgentEnvironmentExperiment(string name, AgentEnvironmentExperimentArgs args, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3beta1:AgentEnvironmentExperiment", name, args ?? new AgentEnvironmentExperimentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AgentEnvironmentExperiment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3beta1:AgentEnvironmentExperiment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AgentEnvironmentExperiment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AgentEnvironmentExperiment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AgentEnvironmentExperiment(name, id, options);
        }
    }

    public sealed class AgentEnvironmentExperimentArgs : Pulumi.ResourceArgs
    {
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

        /// <summary>
        /// Creation time of this experiment.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The definition of the experiment.
        /// </summary>
        [Input("definition")]
        public Input<Inputs.GoogleCloudDialogflowCxV3beta1ExperimentDefinitionArgs>? Definition { get; set; }

        /// <summary>
        /// The human-readable description of the experiment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Required. The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// End time of this experiment.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        [Input("experimentId", required: true)]
        public Input<string> ExperimentId { get; set; } = null!;

        /// <summary>
        /// Maximum number of days to run the experiment. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
        /// </summary>
        [Input("experimentLength")]
        public Input<string>? ExperimentLength { get; set; }

        /// <summary>
        /// Last update time of this experiment.
        /// </summary>
        [Input("lastUpdateTime")]
        public Input<string>? LastUpdateTime { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the experiment. Format: projects//locations//agents//environments//experiments/..
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Inference result of the experiment.
        /// </summary>
        [Input("result")]
        public Input<Inputs.GoogleCloudDialogflowCxV3beta1ExperimentResultArgs>? Result { get; set; }

        /// <summary>
        /// Start time of this experiment.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// The current state of the experiment. Transition triggered by Expriments.StartExperiment: PENDING-&gt;RUNNING. Transition triggered by Expriments.CancelExperiment: PENDING-&gt;CANCELLED or RUNNING-&gt;CANCELLED.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("variantsHistory")]
        private InputList<Inputs.GoogleCloudDialogflowCxV3beta1VariantsHistoryArgs>? _variantsHistory;

        /// <summary>
        /// The history of updates to the experiment variants.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowCxV3beta1VariantsHistoryArgs> VariantsHistory
        {
            get => _variantsHistory ?? (_variantsHistory = new InputList<Inputs.GoogleCloudDialogflowCxV3beta1VariantsHistoryArgs>());
            set => _variantsHistory = value;
        }

        public AgentEnvironmentExperimentArgs()
        {
        }
    }
}
