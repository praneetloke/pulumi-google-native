// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.RemoteBuildExecution.V1Alpha
{
    public static class GetInstance
    {
        /// <summary>
        /// Returns the specified instance.
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("google-native:remotebuildexecution/v1alpha:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified instance.
        /// </summary>
        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("google-native:remotebuildexecution/v1alpha:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetInstanceArgs()
        {
        }
        public static new GetInstanceArgs Empty => new GetInstanceArgs();
    }

    public sealed class GetInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetInstanceInvokeArgs()
        {
        }
        public static new GetInstanceInvokeArgs Empty => new GetInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        /// <summary>
        /// The policy to define whether or not RBE features can be used or how they can be used.
        /// </summary>
        public readonly Outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyResponse FeaturePolicy;
        /// <summary>
        /// The location is a GCP region. Currently only `us-central1` is supported.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Whether stack driver logging is enabled for the instance.
        /// </summary>
        public readonly bool LoggingEnabled;
        /// <summary>
        /// Instance resource name formatted as: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]`. Name should not be populated when creating an instance since it is provided in the `instance_id` field.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// State of the instance.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetInstanceResult(
            Outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyResponse featurePolicy,

            string location,

            bool loggingEnabled,

            string name,

            string state)
        {
            FeaturePolicy = featurePolicy;
            Location = location;
            LoggingEnabled = loggingEnabled;
            Name = name;
            State = state;
        }
    }
}
