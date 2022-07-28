// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    /// <summary>
    /// Creates a NodeGroup resource in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/alpha:NodeGroup")]
    public partial class NodeGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies how autoscaling should behave.
        /// </summary>
        [Output("autoscalingPolicy")]
        public Output<Outputs.NodeGroupAutoscalingPolicyResponse> AutoscalingPolicy { get; private set; } = null!;

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

        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// Initial count of nodes in the node group.
        /// </summary>
        [Output("initialNodeCount")]
        public Output<string> InitialNodeCount { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Always compute#nodeGroup for node group.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// An opaque location hint used to place the Node close to other resources. This field is for use by internal tools that use the public API. The location hint here on the NodeGroup overrides any location_hint present in the NodeTemplate.
        /// </summary>
        [Output("locationHint")]
        public Output<string> LocationHint { get; private set; } = null!;

        /// <summary>
        /// Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT. For more information, see Maintenance policies.
        /// </summary>
        [Output("maintenancePolicy")]
        public Output<string> MaintenancePolicy { get; private set; } = null!;

        [Output("maintenanceWindow")]
        public Output<Outputs.NodeGroupMaintenanceWindowResponse> MaintenanceWindow { get; private set; } = null!;

        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// URL of the node template to create the node group from.
        /// </summary>
        [Output("nodeTemplate")]
        public Output<string> NodeTemplate { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

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
        /// Server-defined URL for this resource with the resource id.
        /// </summary>
        [Output("selfLinkWithId")]
        public Output<string> SelfLinkWithId { get; private set; } = null!;

        /// <summary>
        /// Share-settings for the node group
        /// </summary>
        [Output("shareSettings")]
        public Output<Outputs.ShareSettingsResponse> ShareSettings { get; private set; } = null!;

        /// <summary>
        /// The total number of nodes in the node group.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a NodeGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NodeGroup(string name, NodeGroupArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:NodeGroup", name, args ?? new NodeGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NodeGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:NodeGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "initialNodeCount",
                    "project",
                    "zone",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NodeGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NodeGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NodeGroup(name, id, options);
        }
    }

    public sealed class NodeGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies how autoscaling should behave.
        /// </summary>
        [Input("autoscalingPolicy")]
        public Input<Inputs.NodeGroupAutoscalingPolicyArgs>? AutoscalingPolicy { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Initial count of nodes in the node group.
        /// </summary>
        [Input("initialNodeCount", required: true)]
        public Input<string> InitialNodeCount { get; set; } = null!;

        /// <summary>
        /// An opaque location hint used to place the Node close to other resources. This field is for use by internal tools that use the public API. The location hint here on the NodeGroup overrides any location_hint present in the NodeTemplate.
        /// </summary>
        [Input("locationHint")]
        public Input<string>? LocationHint { get; set; }

        /// <summary>
        /// Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT. For more information, see Maintenance policies.
        /// </summary>
        [Input("maintenancePolicy")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.NodeGroupMaintenancePolicy>? MaintenancePolicy { get; set; }

        [Input("maintenanceWindow")]
        public Input<Inputs.NodeGroupMaintenanceWindowArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// URL of the node template to create the node group from.
        /// </summary>
        [Input("nodeTemplate")]
        public Input<string>? NodeTemplate { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Share-settings for the node group
        /// </summary>
        [Input("shareSettings")]
        public Input<Inputs.ShareSettingsArgs>? ShareSettings { get; set; }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public NodeGroupArgs()
        {
        }
        public static new NodeGroupArgs Empty => new NodeGroupArgs();
    }
}
