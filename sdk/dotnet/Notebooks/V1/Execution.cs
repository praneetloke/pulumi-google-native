// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Notebooks.V1
{
    /// <summary>
    /// Creates a new Scheduled Notebook in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:notebooks/v1:Execution")]
    public partial class Execution : Pulumi.CustomResource
    {
        /// <summary>
        /// Time the Execution was instantiated.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// A brief description of this execution.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Name used for UI purposes. Name can only contain alphanumeric characters and underscores '_'.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// execute metadata including name, hardware spec, region, labels, etc.
        /// </summary>
        [Output("executionTemplate")]
        public Output<Outputs.ExecutionTemplateResponse> ExecutionTemplate { get; private set; } = null!;

        /// <summary>
        /// The resource name of the execute. Format: `projects/{project_id}/locations/{location}/execution/{execution_id}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Output notebook file generated by this execution
        /// </summary>
        [Output("outputNotebookFile")]
        public Output<string> OutputNotebookFile { get; private set; } = null!;

        /// <summary>
        /// State of the underlying AI Platform job.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Time the Execution was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Execution resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Execution(string name, ExecutionArgs args, CustomResourceOptions? options = null)
            : base("google-native:notebooks/v1:Execution", name, args ?? new ExecutionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Execution(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:notebooks/v1:Execution", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Execution resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Execution Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Execution(name, id, options);
        }
    }

    public sealed class ExecutionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A brief description of this execution.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("executionId", required: true)]
        public Input<string> ExecutionId { get; set; } = null!;

        /// <summary>
        /// execute metadata including name, hardware spec, region, labels, etc.
        /// </summary>
        [Input("executionTemplate")]
        public Input<Inputs.ExecutionTemplateArgs>? ExecutionTemplate { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Output notebook file generated by this execution
        /// </summary>
        [Input("outputNotebookFile")]
        public Input<string>? OutputNotebookFile { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public ExecutionArgs()
        {
        }
    }
}
