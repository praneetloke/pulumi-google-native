// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.IDS.V1
{
    public static class GetEndpoint
    {
        /// <summary>
        /// Gets details of a single Endpoint.
        /// </summary>
        public static Task<GetEndpointResult> InvokeAsync(GetEndpointArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEndpointResult>("google-native:ids/v1:getEndpoint", args ?? new GetEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Endpoint.
        /// </summary>
        public static Output<GetEndpointResult> Invoke(GetEndpointInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEndpointResult>("google-native:ids/v1:getEndpoint", args ?? new GetEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEndpointArgs : global::Pulumi.InvokeArgs
    {
        [Input("endpointId", required: true)]
        public string EndpointId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetEndpointArgs()
        {
        }
        public static new GetEndpointArgs Empty => new GetEndpointArgs();
    }

    public sealed class GetEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("endpointId", required: true)]
        public Input<string> EndpointId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetEndpointInvokeArgs()
        {
        }
        public static new GetEndpointInvokeArgs Empty => new GetEndpointInvokeArgs();
    }


    [OutputType]
    public sealed class GetEndpointResult
    {
        /// <summary>
        /// The create time timestamp.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// User-provided description of the endpoint
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The fully qualified URL of the endpoint's ILB Forwarding Rule.
        /// </summary>
        public readonly string EndpointForwardingRule;
        /// <summary>
        /// The IP address of the IDS Endpoint's ILB.
        /// </summary>
        public readonly string EndpointIp;
        /// <summary>
        /// The labels of the endpoint.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The fully qualified URL of the network to which the IDS Endpoint is attached.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Lowest threat severity that this endpoint will alert on.
        /// </summary>
        public readonly string Severity;
        /// <summary>
        /// Current state of the endpoint.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Whether the endpoint should report traffic logs in addition to threat logs.
        /// </summary>
        public readonly bool TrafficLogs;
        /// <summary>
        /// The update time timestamp.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetEndpointResult(
            string createTime,

            string description,

            string endpointForwardingRule,

            string endpointIp,

            ImmutableDictionary<string, string> labels,

            string name,

            string network,

            string severity,

            string state,

            bool trafficLogs,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            EndpointForwardingRule = endpointForwardingRule;
            EndpointIp = endpointIp;
            Labels = labels;
            Name = name;
            Network = network;
            Severity = severity;
            State = state;
            TrafficLogs = trafficLogs;
            UpdateTime = updateTime;
        }
    }
}
