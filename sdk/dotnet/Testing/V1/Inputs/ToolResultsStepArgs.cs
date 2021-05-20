// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Testing.V1.Inputs
{

    /// <summary>
    /// Represents a tool results step resource. This has the results of a TestExecution.
    /// </summary>
    public sealed class ToolResultsStepArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A tool results execution ID.
        /// </summary>
        [Input("executionId")]
        public Input<string>? ExecutionId { get; set; }

        /// <summary>
        /// A tool results history ID.
        /// </summary>
        [Input("historyId")]
        public Input<string>? HistoryId { get; set; }

        /// <summary>
        /// The cloud project that owns the tool results step.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A tool results step ID.
        /// </summary>
        [Input("stepId")]
        public Input<string>? StepId { get; set; }

        public ToolResultsStepArgs()
        {
        }
    }
}
