// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.TPU.V2
{
    public static class GetNode
    {
        /// <summary>
        /// Gets the details of a node.
        /// </summary>
        public static Task<GetNodeResult> InvokeAsync(GetNodeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNodeResult>("google-native:tpu/v2:getNode", args ?? new GetNodeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of a node.
        /// </summary>
        public static Output<GetNodeResult> Invoke(GetNodeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNodeResult>("google-native:tpu/v2:getNode", args ?? new GetNodeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNodeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("nodeId", required: true)]
        public string NodeId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetNodeArgs()
        {
        }
        public static new GetNodeArgs Empty => new GetNodeArgs();
    }

    public sealed class GetNodeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("nodeId", required: true)]
        public Input<string> NodeId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetNodeInvokeArgs()
        {
        }
        public static new GetNodeInvokeArgs Empty => new GetNodeInvokeArgs();
    }


    [OutputType]
    public sealed class GetNodeResult
    {
        /// <summary>
        /// The type of hardware accelerators associated with this node.
        /// </summary>
        public readonly string AcceleratorType;
        /// <summary>
        /// The API version that created this Node.
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
        /// </summary>
        public readonly string CidrBlock;
        /// <summary>
        /// The time when the node was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The additional data disks for the Node.
        /// </summary>
        public readonly ImmutableArray<Outputs.AttachedDiskResponse> DataDisks;
        /// <summary>
        /// The user-supplied description of the TPU. Maximum of 512 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The health status of the TPU node.
        /// </summary>
        public readonly string Health;
        /// <summary>
        /// If this field is populated, it contains a description of why the TPU Node is unhealthy.
        /// </summary>
        public readonly string HealthDescription;
        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Custom metadata to apply to the TPU Node. Can set startup-script and shutdown-script
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// Immutable. The name of the TPU.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Network configurations for the TPU node.
        /// </summary>
        public readonly Outputs.NetworkConfigResponse NetworkConfig;
        /// <summary>
        /// The network endpoints where TPU workers can be accessed and sent work. It is recommended that runtime clients of the node reach out to the 0th entry in this map first.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkEndpointResponse> NetworkEndpoints;
        /// <summary>
        /// The runtime version running in the Node.
        /// </summary>
        public readonly string RuntimeVersion;
        /// <summary>
        /// The scheduling options for this node.
        /// </summary>
        public readonly Outputs.SchedulingConfigResponse SchedulingConfig;
        /// <summary>
        /// The Google Cloud Platform Service Account to be used by the TPU node VMs. If None is specified, the default compute service account will be used.
        /// </summary>
        public readonly Outputs.ServiceAccountResponse ServiceAccount;
        /// <summary>
        /// Shielded Instance options.
        /// </summary>
        public readonly Outputs.ShieldedInstanceConfigResponse ShieldedInstanceConfig;
        /// <summary>
        /// The current state for the TPU Node.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The Symptoms that have occurred to the TPU Node.
        /// </summary>
        public readonly ImmutableArray<Outputs.SymptomResponse> Symptoms;
        /// <summary>
        /// Tags to apply to the TPU Node. Tags are used to identify valid sources or targets for network firewalls.
        /// </summary>
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private GetNodeResult(
            string acceleratorType,

            string apiVersion,

            string cidrBlock,

            string createTime,

            ImmutableArray<Outputs.AttachedDiskResponse> dataDisks,

            string description,

            string health,

            string healthDescription,

            ImmutableDictionary<string, string> labels,

            ImmutableDictionary<string, string> metadata,

            string name,

            Outputs.NetworkConfigResponse networkConfig,

            ImmutableArray<Outputs.NetworkEndpointResponse> networkEndpoints,

            string runtimeVersion,

            Outputs.SchedulingConfigResponse schedulingConfig,

            Outputs.ServiceAccountResponse serviceAccount,

            Outputs.ShieldedInstanceConfigResponse shieldedInstanceConfig,

            string state,

            ImmutableArray<Outputs.SymptomResponse> symptoms,

            ImmutableArray<string> tags)
        {
            AcceleratorType = acceleratorType;
            ApiVersion = apiVersion;
            CidrBlock = cidrBlock;
            CreateTime = createTime;
            DataDisks = dataDisks;
            Description = description;
            Health = health;
            HealthDescription = healthDescription;
            Labels = labels;
            Metadata = metadata;
            Name = name;
            NetworkConfig = networkConfig;
            NetworkEndpoints = networkEndpoints;
            RuntimeVersion = runtimeVersion;
            SchedulingConfig = schedulingConfig;
            ServiceAccount = serviceAccount;
            ShieldedInstanceConfig = shieldedInstanceConfig;
            State = state;
            Symptoms = symptoms;
            Tags = tags;
        }
    }
}
