// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.WorkflowExecutions.V1.Outputs
{

    /// <summary>
    /// Represents the current status of this execution.
    /// </summary>
    [OutputType]
    public sealed class StatusResponse
    {
        /// <summary>
        /// A list of currently executing or last executed step names for the workflow execution currently running. If the workflow has succeeded or failed, this is the last attempted or executed step. Presently, if the current step is inside a subworkflow, the list only includes that step. In the future, the list will contain items for each step in the call stack, starting with the outermost step in the `main` subworkflow, and ending with the most deeply nested step.
        /// </summary>
        public readonly ImmutableArray<Outputs.StepResponse> CurrentSteps;

        [OutputConstructor]
        private StatusResponse(ImmutableArray<Outputs.StepResponse> currentSteps)
        {
            CurrentSteps = currentSteps;
        }
    }
}
