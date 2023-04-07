// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetNodeGroup
    {
        /// <summary>
        /// Returns the specified NodeGroup. Get a list of available NodeGroups by making a list() request. Note: the "nodes" field should not be used. Use nodeGroups.listNodes instead.
        /// </summary>
        public static Task<GetNodeGroupResult> InvokeAsync(GetNodeGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNodeGroupResult>("google-native:compute/alpha:getNodeGroup", args ?? new GetNodeGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified NodeGroup. Get a list of available NodeGroups by making a list() request. Note: the "nodes" field should not be used. Use nodeGroups.listNodes instead.
        /// </summary>
        public static Output<GetNodeGroupResult> Invoke(GetNodeGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNodeGroupResult>("google-native:compute/alpha:getNodeGroup", args ?? new GetNodeGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNodeGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("nodeGroup", required: true)]
        public string NodeGroup { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetNodeGroupArgs()
        {
        }
        public static new GetNodeGroupArgs Empty => new GetNodeGroupArgs();
    }

    public sealed class GetNodeGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("nodeGroup", required: true)]
        public Input<string> NodeGroup { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetNodeGroupInvokeArgs()
        {
        }
        public static new GetNodeGroupInvokeArgs Empty => new GetNodeGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetNodeGroupResult
    {
        /// <summary>
        /// Specifies how autoscaling should behave.
        /// </summary>
        public readonly Outputs.NodeGroupAutoscalingPolicyResponse AutoscalingPolicy;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        public readonly string Fingerprint;
        /// <summary>
        /// The type of the resource. Always compute#nodeGroup for node group.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// An opaque location hint used to place the Node close to other resources. This field is for use by internal tools that use the public API. The location hint here on the NodeGroup overrides any location_hint present in the NodeTemplate.
        /// </summary>
        public readonly string LocationHint;
        /// <summary>
        /// Specifies the frequency of planned maintenance events. The accepted values are: `AS_NEEDED` and `RECURRENT`.
        /// </summary>
        public readonly string MaintenanceInterval;
        /// <summary>
        /// Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT. For more information, see Maintenance policies.
        /// </summary>
        public readonly string MaintenancePolicy;
        public readonly Outputs.NodeGroupMaintenanceWindowResponse MaintenanceWindow;
        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// URL of the node template to create the node group from.
        /// </summary>
        public readonly string NodeTemplate;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Server-defined URL for this resource with the resource id.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// Share-settings for the node group
        /// </summary>
        public readonly Outputs.ShareSettingsResponse ShareSettings;
        /// <summary>
        /// The total number of nodes in the node group.
        /// </summary>
        public readonly int Size;
        public readonly string Status;
        /// <summary>
        /// The name of the zone where the node group resides, such as us-central1-a.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetNodeGroupResult(
            Outputs.NodeGroupAutoscalingPolicyResponse autoscalingPolicy,

            string creationTimestamp,

            string description,

            string fingerprint,

            string kind,

            string locationHint,

            string maintenanceInterval,

            string maintenancePolicy,

            Outputs.NodeGroupMaintenanceWindowResponse maintenanceWindow,

            string name,

            string nodeTemplate,

            string selfLink,

            string selfLinkWithId,

            Outputs.ShareSettingsResponse shareSettings,

            int size,

            string status,

            string zone)
        {
            AutoscalingPolicy = autoscalingPolicy;
            CreationTimestamp = creationTimestamp;
            Description = description;
            Fingerprint = fingerprint;
            Kind = kind;
            LocationHint = locationHint;
            MaintenanceInterval = maintenanceInterval;
            MaintenancePolicy = maintenancePolicy;
            MaintenanceWindow = maintenanceWindow;
            Name = name;
            NodeTemplate = nodeTemplate;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            ShareSettings = shareSettings;
            Size = size;
            Status = status;
            Zone = zone;
        }
    }
}
