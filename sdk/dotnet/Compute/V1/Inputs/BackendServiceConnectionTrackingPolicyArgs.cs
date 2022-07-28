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
    /// Connection Tracking configuration for this BackendService.
    /// </summary>
    public sealed class BackendServiceConnectionTrackingPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies connection persistence when backends are unhealthy. The default value is DEFAULT_FOR_PROTOCOL. If set to DEFAULT_FOR_PROTOCOL, the existing connections persist on unhealthy backends only for connection-oriented protocols (TCP and SCTP) and only if the Tracking Mode is PER_CONNECTION (default tracking mode) or the Session Affinity is configured for 5-tuple. They do not persist for UDP. If set to NEVER_PERSIST, after a backend becomes unhealthy, the existing connections on the unhealthy backend are never persisted on the unhealthy backend. They are always diverted to newly selected healthy backends (unless all backends are unhealthy). If set to ALWAYS_PERSIST, existing connections always persist on unhealthy backends regardless of protocol and session affinity. It is generally not recommended to use this mode overriding the default. For more details, see [Connection Persistence for Network Load Balancing](https://cloud.google.com/load-balancing/docs/network/networklb-backend-service#connection-persistence) and [Connection Persistence for Internal TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/internal#connection-persistence).
        /// </summary>
        [Input("connectionPersistenceOnUnhealthyBackends")]
        public Input<Pulumi.GoogleNative.Compute.V1.BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackends>? ConnectionPersistenceOnUnhealthyBackends { get; set; }

        /// <summary>
        /// Enable Strong Session Affinity for Network Load Balancing. This option is not available publicly.
        /// </summary>
        [Input("enableStrongAffinity")]
        public Input<bool>? EnableStrongAffinity { get; set; }

        /// <summary>
        /// Specifies how long to keep a Connection Tracking entry while there is no matching traffic (in seconds). For Internal TCP/UDP Load Balancing: - The minimum (default) is 10 minutes and the maximum is 16 hours. - It can be set only if Connection Tracking is less than 5-tuple (i.e. Session Affinity is CLIENT_IP_NO_DESTINATION, CLIENT_IP or CLIENT_IP_PROTO, and Tracking Mode is PER_SESSION). For Network Load Balancer the default is 60 seconds. This option is not available publicly.
        /// </summary>
        [Input("idleTimeoutSec")]
        public Input<int>? IdleTimeoutSec { get; set; }

        /// <summary>
        /// Specifies the key used for connection tracking. There are two options: - PER_CONNECTION: This is the default mode. The Connection Tracking is performed as per the Connection Key (default Hash Method) for the specific protocol. - PER_SESSION: The Connection Tracking is performed as per the configured Session Affinity. It matches the configured Session Affinity. For more details, see [Tracking Mode for Network Load Balancing](https://cloud.google.com/load-balancing/docs/network/networklb-backend-service#tracking-mode) and [Tracking Mode for Internal TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/internal#tracking-mode).
        /// </summary>
        [Input("trackingMode")]
        public Input<Pulumi.GoogleNative.Compute.V1.BackendServiceConnectionTrackingPolicyTrackingMode>? TrackingMode { get; set; }

        public BackendServiceConnectionTrackingPolicyArgs()
        {
        }
        public static new BackendServiceConnectionTrackingPolicyArgs Empty => new BackendServiceConnectionTrackingPolicyArgs();
    }
}
