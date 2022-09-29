// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    [OutputType]
    public sealed class HTTPSHealthCheckResponse
    {
        /// <summary>
        /// The value of the host header in the HTTPS health check request. If left empty (default value), the host header is set to the destination IP address to which health check packets are sent. The destination IP address depends on the type of load balancer. For details, see: https://cloud.google.com/load-balancing/docs/health-check-concepts#hc-packet-dest
        /// </summary>
        public readonly string Host;
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
        /// The request path of the HTTPS health check request. The default value is /.
        /// </summary>
        public readonly string RequestPath;
        /// <summary>
        /// Creates a content-based HTTPS health check. In addition to the required HTTP 200 (OK) status code, you can configure the health check to pass only when the backend sends this specific ASCII response string within the first 1024 bytes of the HTTP response body. For details, see: https://cloud.google.com/load-balancing/docs/health-check-concepts#criteria-protocol-http
        /// </summary>
        public readonly string Response;

        [OutputConstructor]
        private HTTPSHealthCheckResponse(
            string host,

            int port,

            string portName,

            string portSpecification,

            string proxyHeader,

            string requestPath,

            string response)
        {
            Host = host;
            Port = port;
            PortName = portName;
            PortSpecification = portSpecification;
            ProxyHeader = proxyHeader;
            RequestPath = requestPath;
            Response = response;
        }
    }
}
