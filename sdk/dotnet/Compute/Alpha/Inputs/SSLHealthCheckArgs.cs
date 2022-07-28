// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    public sealed class SSLHealthCheckArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The TCP port number for the health check request. The default value is 443. Valid values are 1 through 65535.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Port name as defined in InstanceGroup#NamedPort#name. If both port and port_name are defined, port takes precedence.
        /// </summary>
        [Input("portName")]
        public Input<string>? PortName { get; set; }

        /// <summary>
        /// Specifies how port is selected for health checking, can be one of following values: USE_FIXED_PORT: The port number in port is used for health checking. USE_NAMED_PORT: The portName is used for health checking. USE_SERVING_PORT: For NetworkEndpointGroup, the port specified for each network endpoint is used for health checking. For other backends, the port or named port specified in the Backend Service is used for health checking. If not specified, SSL health check follows behavior specified in port and portName fields.
        /// </summary>
        [Input("portSpecification")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SSLHealthCheckPortSpecification>? PortSpecification { get; set; }

        /// <summary>
        /// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
        /// </summary>
        [Input("proxyHeader")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SSLHealthCheckProxyHeader>? ProxyHeader { get; set; }

        /// <summary>
        /// The application data to send once the SSL connection has been established (default value is empty). If both request and response are empty, the connection establishment alone will indicate health. The request data can only be ASCII.
        /// </summary>
        [Input("request")]
        public Input<string>? Request { get; set; }

        /// <summary>
        /// The bytes to match against the beginning of the response data. If left empty (the default value), any response will indicate health. The response data can only be ASCII.
        /// </summary>
        [Input("response")]
        public Input<string>? Response { get; set; }

        public SSLHealthCheckArgs()
        {
        }
        public static new SSLHealthCheckArgs Empty => new SSLHealthCheckArgs();
    }
}
