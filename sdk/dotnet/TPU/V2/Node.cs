// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.TPU.V2
{
    /// <summary>
    /// Creates a node.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:tpu/v2:Node")]
    public partial class Node : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The type of hardware accelerators associated with this node.
        /// </summary>
        [Output("acceleratorType")]
        public Output<string> AcceleratorType { get; private set; } = null!;

        /// <summary>
        /// The API version that created this Node.
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
        /// </summary>
        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The time when the node was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The additional data disks for the Node.
        /// </summary>
        [Output("dataDisks")]
        public Output<ImmutableArray<Outputs.AttachedDiskResponse>> DataDisks { get; private set; } = null!;

        /// <summary>
        /// The user-supplied description of the TPU. Maximum of 512 characters.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The health status of the TPU node.
        /// </summary>
        [Output("health")]
        public Output<string> Health { get; private set; } = null!;

        /// <summary>
        /// If this field is populated, it contains a description of why the TPU Node is unhealthy.
        /// </summary>
        [Output("healthDescription")]
        public Output<string> HealthDescription { get; private set; } = null!;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Custom metadata to apply to the TPU Node. Can set startup-script and shutdown-script
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>> Metadata { get; private set; } = null!;

        /// <summary>
        /// Immutable. The name of the TPU.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network configurations for the TPU node.
        /// </summary>
        [Output("networkConfig")]
        public Output<Outputs.NetworkConfigResponse> NetworkConfig { get; private set; } = null!;

        /// <summary>
        /// The network endpoints where TPU workers can be accessed and sent work. It is recommended that runtime clients of the node reach out to the 0th entry in this map first.
        /// </summary>
        [Output("networkEndpoints")]
        public Output<ImmutableArray<Outputs.NetworkEndpointResponse>> NetworkEndpoints { get; private set; } = null!;

        /// <summary>
        /// The unqualified resource name.
        /// </summary>
        [Output("nodeId")]
        public Output<string?> NodeId { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The runtime version running in the Node.
        /// </summary>
        [Output("runtimeVersion")]
        public Output<string> RuntimeVersion { get; private set; } = null!;

        /// <summary>
        /// The scheduling options for this node.
        /// </summary>
        [Output("schedulingConfig")]
        public Output<Outputs.SchedulingConfigResponse> SchedulingConfig { get; private set; } = null!;

        /// <summary>
        /// The Google Cloud Platform Service Account to be used by the TPU node VMs. If None is specified, the default compute service account will be used.
        /// </summary>
        [Output("serviceAccount")]
        public Output<Outputs.ServiceAccountResponse> ServiceAccount { get; private set; } = null!;

        /// <summary>
        /// Shielded Instance options.
        /// </summary>
        [Output("shieldedInstanceConfig")]
        public Output<Outputs.ShieldedInstanceConfigResponse> ShieldedInstanceConfig { get; private set; } = null!;

        /// <summary>
        /// The current state for the TPU Node.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The Symptoms that have occurred to the TPU Node.
        /// </summary>
        [Output("symptoms")]
        public Output<ImmutableArray<Outputs.SymptomResponse>> Symptoms { get; private set; } = null!;

        /// <summary>
        /// Tags to apply to the TPU Node. Tags are used to identify valid sources or targets for network firewalls.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Node resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Node(string name, NodeArgs args, CustomResourceOptions? options = null)
            : base("google-native:tpu/v2:Node", name, args ?? new NodeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Node(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:tpu/v2:Node", name, null, MakeResourceOptions(options, id))
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
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Node resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Node Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Node(name, id, options);
        }
    }

    public sealed class NodeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of hardware accelerators associated with this node.
        /// </summary>
        [Input("acceleratorType", required: true)]
        public Input<string> AcceleratorType { get; set; } = null!;

        /// <summary>
        /// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        [Input("dataDisks")]
        private InputList<Inputs.AttachedDiskArgs>? _dataDisks;

        /// <summary>
        /// The additional data disks for the Node.
        /// </summary>
        public InputList<Inputs.AttachedDiskArgs> DataDisks
        {
            get => _dataDisks ?? (_dataDisks = new InputList<Inputs.AttachedDiskArgs>());
            set => _dataDisks = value;
        }

        /// <summary>
        /// The user-supplied description of the TPU. Maximum of 512 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The health status of the TPU node.
        /// </summary>
        [Input("health")]
        public Input<Pulumi.GoogleNative.TPU.V2.NodeHealth>? Health { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Custom metadata to apply to the TPU Node. Can set startup-script and shutdown-script
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// Network configurations for the TPU node.
        /// </summary>
        [Input("networkConfig")]
        public Input<Inputs.NetworkConfigArgs>? NetworkConfig { get; set; }

        /// <summary>
        /// The unqualified resource name.
        /// </summary>
        [Input("nodeId")]
        public Input<string>? NodeId { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The runtime version running in the Node.
        /// </summary>
        [Input("runtimeVersion", required: true)]
        public Input<string> RuntimeVersion { get; set; } = null!;

        /// <summary>
        /// The scheduling options for this node.
        /// </summary>
        [Input("schedulingConfig")]
        public Input<Inputs.SchedulingConfigArgs>? SchedulingConfig { get; set; }

        /// <summary>
        /// The Google Cloud Platform Service Account to be used by the TPU node VMs. If None is specified, the default compute service account will be used.
        /// </summary>
        [Input("serviceAccount")]
        public Input<Inputs.ServiceAccountArgs>? ServiceAccount { get; set; }

        /// <summary>
        /// Shielded Instance options.
        /// </summary>
        [Input("shieldedInstanceConfig")]
        public Input<Inputs.ShieldedInstanceConfigArgs>? ShieldedInstanceConfig { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags to apply to the TPU Node. Tags are used to identify valid sources or targets for network firewalls.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public NodeArgs()
        {
        }
        public static new NodeArgs Empty => new NodeArgs();
    }
}
