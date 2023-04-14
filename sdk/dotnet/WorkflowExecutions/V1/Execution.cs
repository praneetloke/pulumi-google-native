// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.WorkflowExecutions.V1
{
    /// <summary>
    /// Creates a new execution using the latest revision of the given workflow.
    /// Auto-naming is currently not supported for this resource.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:workflowexecutions/v1:Execution")]
    public partial class Execution : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
        /// </summary>
        [Output("argument")]
        public Output<string> Argument { get; private set; } = null!;

        /// <summary>
        /// The call logging level associated to this execution.
        /// </summary>
        [Output("callLogLevel")]
        public Output<string> CallLogLevel { get; private set; } = null!;

        /// <summary>
        /// Measures the duration of the execution.
        /// </summary>
        [Output("duration")]
        public Output<string> Duration { get; private set; } = null!;

        /// <summary>
        /// Marks the end of execution, successful or not.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// The error which caused the execution to finish prematurely. The value is only present if the execution's state is `FAILED` or `CANCELLED`.
        /// </summary>
        [Output("error")]
        public Output<Outputs.ErrorResponse> Error { get; private set; } = null!;

        /// <summary>
        /// Labels associated with this execution. Labels can contain at most 64 entries. Keys and values can be no longer than 63 characters and can only contain lowercase letters, numeric characters, underscores, and dashes. Label keys must start with a letter. International characters are allowed. By default, labels are inherited from the workflow but are overridden by any labels associated with the execution.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name of the execution. Format: projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Output of the execution represented as a JSON string. The value can only be present if the execution's state is `SUCCEEDED`.
        /// </summary>
        [Output("result")]
        public Output<string> Result { get; private set; } = null!;

        /// <summary>
        /// Marks the beginning of execution.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// Current state of the execution.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Error regarding the state of the Execution resource. For example, this field will have error details if the Execution data is unavailable due to revoked KMS key permissions.
        /// </summary>
        [Output("stateError")]
        public Output<Outputs.StateErrorResponse> StateError { get; private set; } = null!;

        /// <summary>
        /// Status tracks the current steps and progress data of this execution.
        /// </summary>
        [Output("status")]
        public Output<Outputs.StatusResponse> Status { get; private set; } = null!;

        [Output("workflowId")]
        public Output<string> WorkflowId { get; private set; } = null!;

        /// <summary>
        /// Revision of the workflow this execution is using.
        /// </summary>
        [Output("workflowRevisionId")]
        public Output<string> WorkflowRevisionId { get; private set; } = null!;


        /// <summary>
        /// Create a Execution resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Execution(string name, ExecutionArgs args, CustomResourceOptions? options = null)
            : base("google-native:workflowexecutions/v1:Execution", name, args ?? new ExecutionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Execution(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:workflowexecutions/v1:Execution", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "location",
                    "project",
                    "workflowId",
                },
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

    public sealed class ExecutionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
        /// </summary>
        [Input("argument")]
        public Input<string>? Argument { get; set; }

        /// <summary>
        /// The call logging level associated to this execution.
        /// </summary>
        [Input("callLogLevel")]
        public Input<Pulumi.GoogleNative.WorkflowExecutions.V1.ExecutionCallLogLevel>? CallLogLevel { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels associated with this execution. Labels can contain at most 64 entries. Keys and values can be no longer than 63 characters and can only contain lowercase letters, numeric characters, underscores, and dashes. Label keys must start with a letter. International characters are allowed. By default, labels are inherited from the workflow but are overridden by any labels associated with the execution.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("workflowId", required: true)]
        public Input<string> WorkflowId { get; set; } = null!;

        public ExecutionArgs()
        {
        }
        public static new ExecutionArgs Empty => new ExecutionArgs();
    }
}
