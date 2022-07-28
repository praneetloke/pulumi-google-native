// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// Message containing information of one individual backend.
    /// </summary>
    public sealed class BackendArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies how to determine whether the backend of a load balancer can handle additional traffic or is fully loaded. For usage guidelines, see Connection balancing mode. Backends must use compatible balancing modes. For more information, see Supported balancing modes and target capacity settings and Restrictions and guidance for instance groups. Note: Currently, if you use the API to configure incompatible balancing modes, the configuration might be accepted even though it has no impact and is ignored. Specifically, Backend.maxUtilization is ignored when Backend.balancingMode is RATE. In the future, this incompatible combination will be rejected.
        /// </summary>
        [Input("balancingMode")]
        public Input<Pulumi.GoogleNative.Compute.V1.BackendBalancingMode>? BalancingMode { get; set; }

        /// <summary>
        /// A multiplier applied to the backend's target capacity of its balancing mode. The default value is 1, which means the group serves up to 100% of its configured capacity (depending on balancingMode). A setting of 0 means the group is completely drained, offering 0% of its available capacity. The valid ranges are 0.0 and [0.1,1.0]. You cannot configure a setting larger than 0 and smaller than 0.1. You cannot configure a setting of 0 when there is only one backend attached to the backend service.
        /// </summary>
        [Input("capacityScaler")]
        public Input<double>? CapacityScaler { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// This field designates whether this is a failover backend. More than one failover backend can be configured for a given BackendService.
        /// </summary>
        [Input("failover")]
        public Input<bool>? Failover { get; set; }

        /// <summary>
        /// The fully-qualified URL of an instance group or network endpoint group (NEG) resource. To determine what types of backends a load balancer supports, see the [Backend services overview](https://cloud.google.com/load-balancing/docs/backend-service#backends). You must use the *fully-qualified* URL (starting with https://www.googleapis.com/) to specify the instance group or NEG. Partial URLs are not supported.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// Defines a target maximum number of simultaneous connections. For usage guidelines, see Connection balancing mode and Utilization balancing mode. Not available if the backend's balancingMode is RATE.
        /// </summary>
        [Input("maxConnections")]
        public Input<int>? MaxConnections { get; set; }

        /// <summary>
        /// Defines a target maximum number of simultaneous connections. For usage guidelines, see Connection balancing mode and Utilization balancing mode. Not available if the backend's balancingMode is RATE.
        /// </summary>
        [Input("maxConnectionsPerEndpoint")]
        public Input<int>? MaxConnectionsPerEndpoint { get; set; }

        /// <summary>
        /// Defines a target maximum number of simultaneous connections. For usage guidelines, see Connection balancing mode and Utilization balancing mode. Not available if the backend's balancingMode is RATE.
        /// </summary>
        [Input("maxConnectionsPerInstance")]
        public Input<int>? MaxConnectionsPerInstance { get; set; }

        /// <summary>
        /// Defines a maximum number of HTTP requests per second (RPS). For usage guidelines, see Rate balancing mode and Utilization balancing mode. Not available if the backend's balancingMode is CONNECTION.
        /// </summary>
        [Input("maxRate")]
        public Input<int>? MaxRate { get; set; }

        /// <summary>
        /// Defines a maximum target for requests per second (RPS). For usage guidelines, see Rate balancing mode and Utilization balancing mode. Not available if the backend's balancingMode is CONNECTION.
        /// </summary>
        [Input("maxRatePerEndpoint")]
        public Input<double>? MaxRatePerEndpoint { get; set; }

        /// <summary>
        /// Defines a maximum target for requests per second (RPS). For usage guidelines, see Rate balancing mode and Utilization balancing mode. Not available if the backend's balancingMode is CONNECTION.
        /// </summary>
        [Input("maxRatePerInstance")]
        public Input<double>? MaxRatePerInstance { get; set; }

        /// <summary>
        /// Optional parameter to define a target capacity for the UTILIZATION balancing mode. The valid range is [0.0, 1.0]. For usage guidelines, see Utilization balancing mode.
        /// </summary>
        [Input("maxUtilization")]
        public Input<double>? MaxUtilization { get; set; }

        public BackendArgs()
        {
        }
        public static new BackendArgs Empty => new BackendArgs();
    }
}
