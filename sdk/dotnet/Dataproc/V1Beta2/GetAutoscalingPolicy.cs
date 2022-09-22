// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1Beta2
{
    public static class GetAutoscalingPolicy
    {
        /// <summary>
        /// Retrieves autoscaling policy.
        /// </summary>
        public static Task<GetAutoscalingPolicyResult> InvokeAsync(GetAutoscalingPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAutoscalingPolicyResult>("google-native:dataproc/v1beta2:getAutoscalingPolicy", args ?? new GetAutoscalingPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves autoscaling policy.
        /// </summary>
        public static Output<GetAutoscalingPolicyResult> Invoke(GetAutoscalingPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAutoscalingPolicyResult>("google-native:dataproc/v1beta2:getAutoscalingPolicy", args ?? new GetAutoscalingPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAutoscalingPolicyArgs : global::Pulumi.InvokeArgs
    {
        [Input("autoscalingPolicyId", required: true)]
        public string AutoscalingPolicyId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetAutoscalingPolicyArgs()
        {
        }
        public static new GetAutoscalingPolicyArgs Empty => new GetAutoscalingPolicyArgs();
    }

    public sealed class GetAutoscalingPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("autoscalingPolicyId", required: true)]
        public Input<string> AutoscalingPolicyId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetAutoscalingPolicyInvokeArgs()
        {
        }
        public static new GetAutoscalingPolicyInvokeArgs Empty => new GetAutoscalingPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetAutoscalingPolicyResult
    {
        public readonly Outputs.BasicAutoscalingAlgorithmResponse BasicAlgorithm;
        /// <summary>
        /// The "resource name" of the autoscaling policy, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.autoscalingPolicies, the resource name of the policy has the following format: projects/{project_id}/regions/{region}/autoscalingPolicies/{policy_id} For projects.locations.autoscalingPolicies, the resource name of the policy has the following format: projects/{project_id}/locations/{location}/autoscalingPolicies/{policy_id}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional. Describes how the autoscaler will operate for secondary workers.
        /// </summary>
        public readonly Outputs.InstanceGroupAutoscalingPolicyConfigResponse SecondaryWorkerConfig;
        /// <summary>
        /// Describes how the autoscaler will operate for primary workers.
        /// </summary>
        public readonly Outputs.InstanceGroupAutoscalingPolicyConfigResponse WorkerConfig;

        [OutputConstructor]
        private GetAutoscalingPolicyResult(
            Outputs.BasicAutoscalingAlgorithmResponse basicAlgorithm,

            string name,

            Outputs.InstanceGroupAutoscalingPolicyConfigResponse secondaryWorkerConfig,

            Outputs.InstanceGroupAutoscalingPolicyConfigResponse workerConfig)
        {
            BasicAlgorithm = basicAlgorithm;
            Name = name;
            SecondaryWorkerConfig = secondaryWorkerConfig;
            WorkerConfig = workerConfig;
        }
    }
}
