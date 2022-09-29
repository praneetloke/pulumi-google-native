// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    [OutputType]
    public sealed class SSLHealthCheckResponse
    {
        /// <summary>
        /// The TCP port number to which the health check prober sends packets. The default value is 443. Valid values are 1 through 65535.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Not supported.
        /// </summary>
        public readonly string PortName;
        /// <summary>
        /// Specifies how a port is selected for health checking. Can be one of the following values: USE_FIXED_PORT: Specifies a port number explicitly using the port field in the health check. Supported by backend services for pass-through load balancers and backend services for proxy load balancers. Not supported by target pools. The health check supports all backends supported by the backend service provided the backend can be health checked. For example, GCE_VM_IP network endpoint groups, GCE_VM_IP_PORT network endpoint groups, and instance group backends. USE_NAMED_PORT: Not supported. USE_SERVING_PORT: Provides an indirect method of specifying the health check port by referring to the backend service. Only supported by backend services for proxy load balancers. Not supported by target pools. Not supported by backend services for pass-through load balancers. Supports all backends that can be health checked; for example, GCE_VM_IP_PORT network endpoint groups and instance group backends. For GCE_VM_IP_PORT network endpoint group backends, the health check uses the port number specified for each endpoint in the network endpoint group. For instance group backends, the health check uses the port number determined by looking up the backend service's named port in the instance group's list of named ports.
        /// </summary>
        public readonly string PortSpecification;
        /// <summary>
        /// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
        /// </summary>
        public readonly string ProxyHeader;
        /// <summary>
        /// Instructs the health check prober to send this exact ASCII string, up to 1024 bytes in length, after establishing the TCP connection and SSL handshake.
        /// </summary>
        public readonly string Request;
        /// <summary>
        /// Creates a content-based SSL health check. In addition to establishing a TCP connection and the TLS handshake, you can configure the health check to pass only when the backend sends this exact response ASCII string, up to 1024 bytes in length. For details, see: https://cloud.google.com/load-balancing/docs/health-check-concepts#criteria-protocol-ssl-tcp
        /// </summary>
        public readonly string Response;

        [OutputConstructor]
        private SSLHealthCheckResponse(
            int port,

            string portName,

            string portSpecification,

            string proxyHeader,

            string request,

            string response)
        {
            Port = port;
            PortName = portName;
            PortSpecification = portSpecification;
            ProxyHeader = proxyHeader;
            Request = request;
            Response = response;
        }
    }
}
