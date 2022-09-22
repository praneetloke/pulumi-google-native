// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3
{
    public static class GetSessionEntityType
    {
        /// <summary>
        /// Retrieves the specified session entity type.
        /// </summary>
        public static Task<GetSessionEntityTypeResult> InvokeAsync(GetSessionEntityTypeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSessionEntityTypeResult>("google-native:dialogflow/v3:getSessionEntityType", args ?? new GetSessionEntityTypeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the specified session entity type.
        /// </summary>
        public static Output<GetSessionEntityTypeResult> Invoke(GetSessionEntityTypeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSessionEntityTypeResult>("google-native:dialogflow/v3:getSessionEntityType", args ?? new GetSessionEntityTypeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSessionEntityTypeArgs : global::Pulumi.InvokeArgs
    {
        [Input("agentId", required: true)]
        public string AgentId { get; set; } = null!;

        [Input("entityTypeId", required: true)]
        public string EntityTypeId { get; set; } = null!;

        [Input("environmentId", required: true)]
        public string EnvironmentId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("sessionId", required: true)]
        public string SessionId { get; set; } = null!;

        public GetSessionEntityTypeArgs()
        {
        }
        public static new GetSessionEntityTypeArgs Empty => new GetSessionEntityTypeArgs();
    }

    public sealed class GetSessionEntityTypeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

        [Input("entityTypeId", required: true)]
        public Input<string> EntityTypeId { get; set; } = null!;

        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sessionId", required: true)]
        public Input<string> SessionId { get; set; } = null!;

        public GetSessionEntityTypeInvokeArgs()
        {
        }
        public static new GetSessionEntityTypeInvokeArgs Empty => new GetSessionEntityTypeInvokeArgs();
    }


    [OutputType]
    public sealed class GetSessionEntityTypeResult
    {
        /// <summary>
        /// The collection of entities to override or supplement the custom entity type.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowCxV3EntityTypeEntityResponse> Entities;
        /// <summary>
        /// Indicates whether the additional data should override or supplement the custom entity type definition.
        /// </summary>
        public readonly string EntityOverrideMode;
        /// <summary>
        /// The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetSessionEntityTypeResult(
            ImmutableArray<Outputs.GoogleCloudDialogflowCxV3EntityTypeEntityResponse> entities,

            string entityOverrideMode,

            string name)
        {
            Entities = entities;
            EntityOverrideMode = entityOverrideMode;
            Name = name;
        }
    }
}
