// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VPCAccess.V1Beta1
{
    public static class GetConnector
    {
        /// <summary>
        /// Gets a Serverless VPC Access connector. Returns NOT_FOUND if the resource does not exist.
        /// </summary>
        public static Task<GetConnectorResult> InvokeAsync(GetConnectorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConnectorResult>("google-native:vpcaccess/v1beta1:getConnector", args ?? new GetConnectorArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a Serverless VPC Access connector. Returns NOT_FOUND if the resource does not exist.
        /// </summary>
        public static Output<GetConnectorResult> Invoke(GetConnectorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConnectorResult>("google-native:vpcaccess/v1beta1:getConnector", args ?? new GetConnectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConnectorArgs : global::Pulumi.InvokeArgs
    {
        [Input("connectorId", required: true)]
        public string ConnectorId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetConnectorArgs()
        {
        }
        public static new GetConnectorArgs Empty => new GetConnectorArgs();
    }

    public sealed class GetConnectorInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("connectorId", required: true)]
        public Input<string> ConnectorId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetConnectorInvokeArgs()
        {
        }
        public static new GetConnectorInvokeArgs Empty => new GetConnectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetConnectorResult
    {
        /// <summary>
        /// List of projects using the connector.
        /// </summary>
        public readonly ImmutableArray<string> ConnectedProjects;
        /// <summary>
        /// The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
        /// </summary>
        public readonly string IpCidrRange;
        /// <summary>
        /// Machine type of VM Instance underlying connector. Default is e2-micro
        /// </summary>
        public readonly string MachineType;
        /// <summary>
        /// Maximum value of instances in autoscaling group underlying the connector.
        /// </summary>
        public readonly int MaxInstances;
        /// <summary>
        /// Maximum throughput of the connector in Mbps. Refers to the expected throughput when using an `e2-micro` machine type. Value must be a multiple of 100 from 300 through 1000. Must be higher than the value specified by --min-throughput. If both max-throughput and max-instances are provided, max-instances takes precedence over max-throughput. The use of `max-throughput` is discouraged in favor of `max-instances`.
        /// </summary>
        public readonly int MaxThroughput;
        /// <summary>
        /// Minimum value of instances in autoscaling group underlying the connector.
        /// </summary>
        public readonly int MinInstances;
        /// <summary>
        /// Minimum throughput of the connector in Mbps. Refers to the expected throughput when using an `e2-micro` machine type. Value must be a multiple of 100 from 200 through 900. Must be lower than the value specified by --max-throughput. If both min-throughput and min-instances are provided, min-instances takes precedence over min-throughput. The use of `min-throughput` is discouraged in favor of `min-instances`.
        /// </summary>
        public readonly int MinThroughput;
        /// <summary>
        /// The resource name in the format `projects/*/locations/*/connectors/*`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Name of a VPC network.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// State of the VPC access connector.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The subnet in which to house the VPC Access Connector.
        /// </summary>
        public readonly Outputs.SubnetResponse Subnet;

        [OutputConstructor]
        private GetConnectorResult(
            ImmutableArray<string> connectedProjects,

            string ipCidrRange,

            string machineType,

            int maxInstances,

            int maxThroughput,

            int minInstances,

            int minThroughput,

            string name,

            string network,

            string state,

            Outputs.SubnetResponse subnet)
        {
            ConnectedProjects = connectedProjects;
            IpCidrRange = ipCidrRange;
            MachineType = machineType;
            MaxInstances = maxInstances;
            MaxThroughput = maxThroughput;
            MinInstances = minInstances;
            MinThroughput = minThroughput;
            Name = name;
            Network = network;
            State = state;
            Subnet = subnet;
        }
    }
}
