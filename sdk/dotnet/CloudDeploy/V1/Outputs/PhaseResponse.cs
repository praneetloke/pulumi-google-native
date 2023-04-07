// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1.Outputs
{

    /// <summary>
    /// Phase represents a collection of jobs that are logically grouped together for a `Rollout`.
    /// </summary>
    [OutputType]
    public sealed class PhaseResponse
    {
        /// <summary>
        /// ChildRollout job composition.
        /// </summary>
        public readonly Outputs.ChildRolloutJobsResponse ChildRolloutJobs;
        /// <summary>
        /// Deployment job composition.
        /// </summary>
        public readonly Outputs.DeploymentJobsResponse DeploymentJobs;
        /// <summary>
        /// Additional information on why the Phase was skipped, if available.
        /// </summary>
        public readonly string SkipMessage;
        /// <summary>
        /// Current state of the Phase.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private PhaseResponse(
            Outputs.ChildRolloutJobsResponse childRolloutJobs,

            Outputs.DeploymentJobsResponse deploymentJobs,

            string skipMessage,

            string state)
        {
            ChildRolloutJobs = childRolloutJobs;
            DeploymentJobs = deploymentJobs;
            SkipMessage = skipMessage;
            State = state;
        }
    }
}
