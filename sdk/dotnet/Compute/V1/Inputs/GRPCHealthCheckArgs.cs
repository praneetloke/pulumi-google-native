// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    public sealed class GRPCHealthCheckArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The gRPC service name for the health check. This field is optional. The value of grpc_service_name has the following meanings by convention: - Empty service_name means the overall status of all services at the backend. - Non-empty service_name means the health of that gRPC service, as defined by the owner of the service. The grpc_service_name can only be ASCII.
        /// </summary>
        [Input("grpcServiceName")]
        public Input<string>? GrpcServiceName { get; set; }

        /// <summary>
        /// The port number for the health check request. Must be specified if port_name and port_specification are not set or if port_specification is USE_FIXED_PORT. Valid values are 1 through 65535.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Port name as defined in InstanceGroup#NamedPort#name. If both port and port_name are defined, port takes precedence. The port_name should conform to RFC1035.
        /// </summary>
        [Input("portName")]
        public Input<string>? PortName { get; set; }

        /// <summary>
        /// Specifies how port is selected for health checking, can be one of following values: USE_FIXED_PORT: The port number in port is used for health checking. USE_NAMED_PORT: The portName is used for health checking. USE_SERVING_PORT: For NetworkEndpointGroup, the port specified for each network endpoint is used for health checking. For other backends, the port or named port specified in the Backend Service is used for health checking. If not specified, gRPC health check follows behavior specified in port and portName fields.
        /// </summary>
        [Input("portSpecification")]
        public Input<Pulumi.GoogleNative.Compute.V1.GRPCHealthCheckPortSpecification>? PortSpecification { get; set; }

        public GRPCHealthCheckArgs()
        {
        }
        public static new GRPCHealthCheckArgs Empty => new GRPCHealthCheckArgs();
    }
}
