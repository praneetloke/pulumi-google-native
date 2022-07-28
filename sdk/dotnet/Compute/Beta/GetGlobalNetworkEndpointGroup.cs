// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetGlobalNetworkEndpointGroup
    {
        /// <summary>
        /// Returns the specified network endpoint group. Gets a list of available network endpoint groups by making a list() request.
        /// </summary>
        public static Task<GetGlobalNetworkEndpointGroupResult> InvokeAsync(GetGlobalNetworkEndpointGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGlobalNetworkEndpointGroupResult>("google-native:compute/beta:getGlobalNetworkEndpointGroup", args ?? new GetGlobalNetworkEndpointGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified network endpoint group. Gets a list of available network endpoint groups by making a list() request.
        /// </summary>
        public static Output<GetGlobalNetworkEndpointGroupResult> Invoke(GetGlobalNetworkEndpointGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGlobalNetworkEndpointGroupResult>("google-native:compute/beta:getGlobalNetworkEndpointGroup", args ?? new GetGlobalNetworkEndpointGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGlobalNetworkEndpointGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("networkEndpointGroup", required: true)]
        public string NetworkEndpointGroup { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetGlobalNetworkEndpointGroupArgs()
        {
        }
        public static new GetGlobalNetworkEndpointGroupArgs Empty => new GetGlobalNetworkEndpointGroupArgs();
    }

    public sealed class GetGlobalNetworkEndpointGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("networkEndpointGroup", required: true)]
        public Input<string> NetworkEndpointGroup { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetGlobalNetworkEndpointGroupInvokeArgs()
        {
        }
        public static new GetGlobalNetworkEndpointGroupInvokeArgs Empty => new GetGlobalNetworkEndpointGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetGlobalNetworkEndpointGroupResult
    {
        /// <summary>
        /// Metadata defined as annotations on the network endpoint group.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Annotations;
        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        /// </summary>
        public readonly Outputs.NetworkEndpointGroupAppEngineResponse AppEngine;
        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        /// </summary>
        public readonly Outputs.NetworkEndpointGroupCloudFunctionResponse CloudFunction;
        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        /// </summary>
        public readonly Outputs.NetworkEndpointGroupCloudRunResponse CloudRun;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// The default port used if the port number is not specified in the network endpoint.
        /// </summary>
        public readonly int DefaultPort;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// This field is only valid when the network endpoint group is used for load balancing. [Deprecated] This field is deprecated.
        /// </summary>
        public readonly Outputs.NetworkEndpointGroupLbNetworkEndpointGroupResponse LoadBalancer;
        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
        /// </summary>
        public readonly string NetworkEndpointType;
        /// <summary>
        /// The target service url used to set up private service connection to a Google API or a PSC Producer Service Attachment. An example value is: "asia-northeast3-cloudkms.googleapis.com"
        /// </summary>
        public readonly string PscTargetService;
        /// <summary>
        /// The URL of the region where the network endpoint group is located.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine, cloudFunction or serverlessDeployment may be set.
        /// </summary>
        public readonly Outputs.NetworkEndpointGroupServerlessDeploymentResponse ServerlessDeployment;
        /// <summary>
        /// [Output only] Number of network endpoints in the network endpoint group.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
        /// </summary>
        public readonly string Subnetwork;
        /// <summary>
        /// The URL of the zone where the network endpoint group is located.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetGlobalNetworkEndpointGroupResult(
            ImmutableDictionary<string, string> annotations,

            Outputs.NetworkEndpointGroupAppEngineResponse appEngine,

            Outputs.NetworkEndpointGroupCloudFunctionResponse cloudFunction,

            Outputs.NetworkEndpointGroupCloudRunResponse cloudRun,

            string creationTimestamp,

            int defaultPort,

            string description,

            string kind,

            Outputs.NetworkEndpointGroupLbNetworkEndpointGroupResponse loadBalancer,

            string name,

            string network,

            string networkEndpointType,

            string pscTargetService,

            string region,

            string selfLink,

            Outputs.NetworkEndpointGroupServerlessDeploymentResponse serverlessDeployment,

            int size,

            string subnetwork,

            string zone)
        {
            Annotations = annotations;
            AppEngine = appEngine;
            CloudFunction = cloudFunction;
            CloudRun = cloudRun;
            CreationTimestamp = creationTimestamp;
            DefaultPort = defaultPort;
            Description = description;
            Kind = kind;
            LoadBalancer = loadBalancer;
            Name = name;
            Network = network;
            NetworkEndpointType = networkEndpointType;
            PscTargetService = pscTargetService;
            Region = region;
            SelfLink = selfLink;
            ServerlessDeployment = serverlessDeployment;
            Size = size;
            Subnetwork = subnetwork;
            Zone = zone;
        }
    }
}
