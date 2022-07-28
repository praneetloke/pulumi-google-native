// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.WorkflowExecutions.V1
{
    public static class GetExecution
    {
        /// <summary>
        /// Returns an execution of the given name.
        /// </summary>
        public static Task<GetExecutionResult> InvokeAsync(GetExecutionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetExecutionResult>("google-native:workflowexecutions/v1:getExecution", args ?? new GetExecutionArgs(), options.WithDefaults());

        /// <summary>
        /// Returns an execution of the given name.
        /// </summary>
        public static Output<GetExecutionResult> Invoke(GetExecutionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetExecutionResult>("google-native:workflowexecutions/v1:getExecution", args ?? new GetExecutionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetExecutionArgs : global::Pulumi.InvokeArgs
    {
        [Input("executionId", required: true)]
        public string ExecutionId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("view")]
        public string? View { get; set; }

        [Input("workflowId", required: true)]
        public string WorkflowId { get; set; } = null!;

        public GetExecutionArgs()
        {
        }
        public static new GetExecutionArgs Empty => new GetExecutionArgs();
    }

    public sealed class GetExecutionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("executionId", required: true)]
        public Input<string> ExecutionId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("view")]
        public Input<string>? View { get; set; }

        [Input("workflowId", required: true)]
        public Input<string> WorkflowId { get; set; } = null!;

        public GetExecutionInvokeArgs()
        {
        }
        public static new GetExecutionInvokeArgs Empty => new GetExecutionInvokeArgs();
    }


    [OutputType]
    public sealed class GetExecutionResult
    {
        /// <summary>
        /// Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
        /// </summary>
        public readonly string Argument;
        /// <summary>
        /// The call logging level associated to this execution.
        /// </summary>
        public readonly string CallLogLevel;
        /// <summary>
        /// Marks the end of execution, successful or not.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// The error which caused the execution to finish prematurely. The value is only present if the execution's state is `FAILED` or `CANCELLED`.
        /// </summary>
        public readonly Outputs.ErrorResponse Error;
        /// <summary>
        /// The resource name of the execution. Format: projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Output of the execution represented as a JSON string. The value can only be present if the execution's state is `SUCCEEDED`.
        /// </summary>
        public readonly string Result;
        /// <summary>
        /// Marks the beginning of execution.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Current state of the execution.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Revision of the workflow this execution is using.
        /// </summary>
        public readonly string WorkflowRevisionId;

        [OutputConstructor]
        private GetExecutionResult(
            string argument,

            string callLogLevel,

            string endTime,

            Outputs.ErrorResponse error,

            string name,

            string result,

            string startTime,

            string state,

            string workflowRevisionId)
        {
            Argument = argument;
            CallLogLevel = callLogLevel;
            EndTime = endTime;
            Error = error;
            Name = name;
            Result = result;
            StartTime = startTime;
            State = state;
            WorkflowRevisionId = workflowRevisionId;
        }
    }
}
