// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Batch.V1.Outputs
{

    /// <summary>
    /// Spec of a task
    /// </summary>
    [OutputType]
    public sealed class TaskSpecResponse
    {
        /// <summary>
        /// ComputeResource requirements.
        /// </summary>
        public readonly Outputs.ComputeResourceResponse ComputeResource;
        /// <summary>
        /// Environment variables to set before running the Task.
        /// </summary>
        public readonly Outputs.EnvironmentResponse Environment;
        /// <summary>
        /// Deprecated: please use environment(non-plural) instead.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Environments;
        /// <summary>
        /// Lifecycle management schema when any task in a task group is failed. Currently we only support one lifecycle policy. When the lifecycle policy condition is met, the action in the policy will execute. If task execution result does not meet with the defined lifecycle policy, we consider it as the default policy. Default policy means if the exit code is 0, exit task. If task ends with non-zero exit code, retry the task with max_retry_count.
        /// </summary>
        public readonly ImmutableArray<Outputs.LifecyclePolicyResponse> LifecyclePolicies;
        /// <summary>
        /// Maximum number of retries on failures. The default, 0, which means never retry. The valid value range is [0, 10].
        /// </summary>
        public readonly int MaxRetryCount;
        /// <summary>
        /// Maximum duration the task should run. The task will be killed and marked as FAILED if over this limit.
        /// </summary>
        public readonly string MaxRunDuration;
        /// <summary>
        /// The sequence of scripts or containers to run for this Task. Each Task using this TaskSpec executes its list of runnables in order. The Task succeeds if all of its runnables either exit with a zero status or any that exit with a non-zero status have the ignore_exit_status flag. Background runnables are killed automatically (if they have not already exited) a short time after all foreground runnables have completed. Even though this is likely to result in a non-zero exit status for the background runnable, these automatic kills are not treated as Task failures.
        /// </summary>
        public readonly ImmutableArray<Outputs.RunnableResponse> Runnables;
        /// <summary>
        /// Volumes to mount before running Tasks using this TaskSpec.
        /// </summary>
        public readonly ImmutableArray<Outputs.VolumeResponse> Volumes;

        [OutputConstructor]
        private TaskSpecResponse(
            Outputs.ComputeResourceResponse computeResource,

            Outputs.EnvironmentResponse environment,

            ImmutableDictionary<string, string> environments,

            ImmutableArray<Outputs.LifecyclePolicyResponse> lifecyclePolicies,

            int maxRetryCount,

            string maxRunDuration,

            ImmutableArray<Outputs.RunnableResponse> runnables,

            ImmutableArray<Outputs.VolumeResponse> volumes)
        {
            ComputeResource = computeResource;
            Environment = environment;
            Environments = environments;
            LifecyclePolicies = lifecyclePolicies;
            MaxRetryCount = maxRetryCount;
            MaxRunDuration = maxRunDuration;
            Runnables = runnables;
            Volumes = volumes;
        }
    }
}
