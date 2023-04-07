// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetNodeTemplate
    {
        /// <summary>
        /// Returns the specified node template.
        /// </summary>
        public static Task<GetNodeTemplateResult> InvokeAsync(GetNodeTemplateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNodeTemplateResult>("google-native:compute/beta:getNodeTemplate", args ?? new GetNodeTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified node template.
        /// </summary>
        public static Output<GetNodeTemplateResult> Invoke(GetNodeTemplateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNodeTemplateResult>("google-native:compute/beta:getNodeTemplate", args ?? new GetNodeTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNodeTemplateArgs : global::Pulumi.InvokeArgs
    {
        [Input("nodeTemplate", required: true)]
        public string NodeTemplate { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        public GetNodeTemplateArgs()
        {
        }
        public static new GetNodeTemplateArgs Empty => new GetNodeTemplateArgs();
    }

    public sealed class GetNodeTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("nodeTemplate", required: true)]
        public Input<string> NodeTemplate { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public GetNodeTemplateInvokeArgs()
        {
        }
        public static new GetNodeTemplateInvokeArgs Empty => new GetNodeTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetNodeTemplateResult
    {
        public readonly ImmutableArray<Outputs.AcceleratorConfigResponse> Accelerators;
        /// <summary>
        /// CPU overcommit.
        /// </summary>
        public readonly string CpuOvercommitType;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        public readonly ImmutableArray<Outputs.LocalDiskResponse> Disks;
        /// <summary>
        /// The type of the resource. Always compute#nodeTemplate for node templates.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Labels to use for node affinity, which will be used in instance scheduling.
        /// </summary>
        public readonly ImmutableDictionary<string, string> NodeAffinityLabels;
        /// <summary>
        /// The node type to use for nodes group that are created from this template.
        /// </summary>
        public readonly string NodeType;
        /// <summary>
        /// Do not use. Instead, use the node_type property.
        /// </summary>
        public readonly Outputs.NodeTemplateNodeTypeFlexibilityResponse NodeTypeFlexibility;
        /// <summary>
        /// The name of the region where the node template resides, such as us-central1.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Sets the binding properties for the physical server. Valid values include: - *[Default]* RESTART_NODE_ON_ANY_SERVER: Restarts VMs on any available physical server - RESTART_NODE_ON_MINIMAL_SERVER: Restarts VMs on the same physical server whenever possible See Sole-tenant node options for more information.
        /// </summary>
        public readonly Outputs.ServerBindingResponse ServerBinding;
        /// <summary>
        /// The status of the node template. One of the following values: CREATING, READY, and DELETING.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// An optional, human-readable explanation of the status.
        /// </summary>
        public readonly string StatusMessage;

        [OutputConstructor]
        private GetNodeTemplateResult(
            ImmutableArray<Outputs.AcceleratorConfigResponse> accelerators,

            string cpuOvercommitType,

            string creationTimestamp,

            string description,

            ImmutableArray<Outputs.LocalDiskResponse> disks,

            string kind,

            string name,

            ImmutableDictionary<string, string> nodeAffinityLabels,

            string nodeType,

            Outputs.NodeTemplateNodeTypeFlexibilityResponse nodeTypeFlexibility,

            string region,

            string selfLink,

            Outputs.ServerBindingResponse serverBinding,

            string status,

            string statusMessage)
        {
            Accelerators = accelerators;
            CpuOvercommitType = cpuOvercommitType;
            CreationTimestamp = creationTimestamp;
            Description = description;
            Disks = disks;
            Kind = kind;
            Name = name;
            NodeAffinityLabels = nodeAffinityLabels;
            NodeType = nodeType;
            NodeTypeFlexibility = nodeTypeFlexibility;
            Region = region;
            SelfLink = selfLink;
            ServerBinding = serverBinding;
            Status = status;
            StatusMessage = statusMessage;
        }
    }
}
