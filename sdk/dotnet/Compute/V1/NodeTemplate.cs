// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    /// <summary>
    /// Creates a NodeTemplate resource in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/v1:NodeTemplate")]
    public partial class NodeTemplate : global::Pulumi.CustomResource
    {
        [Output("accelerators")]
        public Output<ImmutableArray<Outputs.AcceleratorConfigResponse>> Accelerators { get; private set; } = null!;

        /// <summary>
        /// CPU overcommit.
        /// </summary>
        [Output("cpuOvercommitType")]
        public Output<string> CpuOvercommitType { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("disks")]
        public Output<ImmutableArray<Outputs.LocalDiskResponse>> Disks { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Always compute#nodeTemplate for node templates.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Labels to use for node affinity, which will be used in instance scheduling.
        /// </summary>
        [Output("nodeAffinityLabels")]
        public Output<ImmutableDictionary<string, string>> NodeAffinityLabels { get; private set; } = null!;

        /// <summary>
        /// The node type to use for nodes group that are created from this template.
        /// </summary>
        [Output("nodeType")]
        public Output<string> NodeType { get; private set; } = null!;

        /// <summary>
        /// The flexible properties of the desired node type. Node groups that use this node template will create nodes of a type that matches these properties. This field is mutually exclusive with the node_type property; you can only define one or the other, but not both.
        /// </summary>
        [Output("nodeTypeFlexibility")]
        public Output<Outputs.NodeTemplateNodeTypeFlexibilityResponse> NodeTypeFlexibility { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Sets the binding properties for the physical server. Valid values include: - *[Default]* RESTART_NODE_ON_ANY_SERVER: Restarts VMs on any available physical server - RESTART_NODE_ON_MINIMAL_SERVER: Restarts VMs on the same physical server whenever possible See Sole-tenant node options for more information.
        /// </summary>
        [Output("serverBinding")]
        public Output<Outputs.ServerBindingResponse> ServerBinding { get; private set; } = null!;

        /// <summary>
        /// The status of the node template. One of the following values: CREATING, READY, and DELETING.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// An optional, human-readable explanation of the status.
        /// </summary>
        [Output("statusMessage")]
        public Output<string> StatusMessage { get; private set; } = null!;


        /// <summary>
        /// Create a NodeTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NodeTemplate(string name, NodeTemplateArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:NodeTemplate", name, args ?? new NodeTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NodeTemplate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:NodeTemplate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "project",
                    "region",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NodeTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NodeTemplate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NodeTemplate(name, id, options);
        }
    }

    public sealed class NodeTemplateArgs : global::Pulumi.ResourceArgs
    {
        [Input("accelerators")]
        private InputList<Inputs.AcceleratorConfigArgs>? _accelerators;
        public InputList<Inputs.AcceleratorConfigArgs> Accelerators
        {
            get => _accelerators ?? (_accelerators = new InputList<Inputs.AcceleratorConfigArgs>());
            set => _accelerators = value;
        }

        /// <summary>
        /// CPU overcommit.
        /// </summary>
        [Input("cpuOvercommitType")]
        public Input<Pulumi.GoogleNative.Compute.V1.NodeTemplateCpuOvercommitType>? CpuOvercommitType { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("disks")]
        private InputList<Inputs.LocalDiskArgs>? _disks;
        public InputList<Inputs.LocalDiskArgs> Disks
        {
            get => _disks ?? (_disks = new InputList<Inputs.LocalDiskArgs>());
            set => _disks = value;
        }

        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nodeAffinityLabels")]
        private InputMap<string>? _nodeAffinityLabels;

        /// <summary>
        /// Labels to use for node affinity, which will be used in instance scheduling.
        /// </summary>
        public InputMap<string> NodeAffinityLabels
        {
            get => _nodeAffinityLabels ?? (_nodeAffinityLabels = new InputMap<string>());
            set => _nodeAffinityLabels = value;
        }

        /// <summary>
        /// The node type to use for nodes group that are created from this template.
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        /// <summary>
        /// The flexible properties of the desired node type. Node groups that use this node template will create nodes of a type that matches these properties. This field is mutually exclusive with the node_type property; you can only define one or the other, but not both.
        /// </summary>
        [Input("nodeTypeFlexibility")]
        public Input<Inputs.NodeTemplateNodeTypeFlexibilityArgs>? NodeTypeFlexibility { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Sets the binding properties for the physical server. Valid values include: - *[Default]* RESTART_NODE_ON_ANY_SERVER: Restarts VMs on any available physical server - RESTART_NODE_ON_MINIMAL_SERVER: Restarts VMs on the same physical server whenever possible See Sole-tenant node options for more information.
        /// </summary>
        [Input("serverBinding")]
        public Input<Inputs.ServerBindingArgs>? ServerBinding { get; set; }

        public NodeTemplateArgs()
        {
        }
        public static new NodeTemplateArgs Empty => new NodeTemplateArgs();
    }
}
