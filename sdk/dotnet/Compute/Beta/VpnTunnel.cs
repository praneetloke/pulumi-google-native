// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    /// <summary>
    /// Creates a VpnTunnel resource in the specified project and region using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/beta:VpnTunnel")]
    public partial class VpnTunnel : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Detailed status message for the VPN tunnel.
        /// </summary>
        [Output("detailedStatus")]
        public Output<string> DetailedStatus { get; private set; } = null!;

        /// <summary>
        /// IKE protocol version to use when establishing the VPN tunnel with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default version is 2.
        /// </summary>
        [Output("ikeVersion")]
        public Output<int> IkeVersion { get; private set; } = null!;

        /// <summary>
        /// Type of resource. Always compute#vpnTunnel for VPN tunnels.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// A fingerprint for the labels being applied to this VpnTunnel, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a VpnTunnel.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Local traffic selector to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is supported.
        /// </summary>
        [Output("localTrafficSelector")]
        public Output<ImmutableArray<string>> LocalTrafficSelector { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// URL of the peer side external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field is exclusive with the field peerGcpGateway.
        /// </summary>
        [Output("peerExternalGateway")]
        public Output<string> PeerExternalGateway { get; private set; } = null!;

        /// <summary>
        /// The interface ID of the external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created.
        /// </summary>
        [Output("peerExternalGatewayInterface")]
        public Output<int> PeerExternalGatewayInterface { get; private set; } = null!;

        /// <summary>
        /// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field can be used when creating highly available VPN from VPC network to VPC network, the field is exclusive with the field peerExternalGateway. If provided, the VPN tunnel will automatically use the same vpnGatewayInterface ID in the peer GCP VPN gateway.
        /// </summary>
        [Output("peerGcpGateway")]
        public Output<string> PeerGcpGateway { get; private set; } = null!;

        /// <summary>
        /// IP address of the peer VPN gateway. Only IPv4 is supported.
        /// </summary>
        [Output("peerIp")]
        public Output<string> PeerIp { get; private set; } = null!;

        /// <summary>
        /// URL of the region where the VPN tunnel resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Remote traffic selectors to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.
        /// </summary>
        [Output("remoteTrafficSelector")]
        public Output<ImmutableArray<string>> RemoteTrafficSelector { get; private set; } = null!;

        /// <summary>
        /// URL of the router resource to be used for dynamic routing.
        /// </summary>
        [Output("router")]
        public Output<string> Router { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
        /// </summary>
        [Output("sharedSecret")]
        public Output<string> SharedSecret { get; private set; } = null!;

        /// <summary>
        /// Hash of the shared secret.
        /// </summary>
        [Output("sharedSecretHash")]
        public Output<string> SharedSecretHash { get; private set; } = null!;

        /// <summary>
        /// The status of the VPN tunnel, which can be one of the following: - PROVISIONING: Resource is being allocated for the VPN tunnel. - WAITING_FOR_FULL_CONFIG: Waiting to receive all VPN-related configs from the user. Network, TargetVpnGateway, VpnTunnel, ForwardingRule, and Route resources are needed to setup the VPN tunnel. - FIRST_HANDSHAKE: Successful first handshake with the peer VPN. - ESTABLISHED: Secure session is successfully established with the peer VPN. - NETWORK_ERROR: Deprecated, replaced by NO_INCOMING_PACKETS - AUTHORIZATION_ERROR: Auth error (for example, bad shared secret). - NEGOTIATION_FAILURE: Handshake failed. - DEPROVISIONING: Resources are being deallocated for the VPN tunnel. - FAILED: Tunnel creation has failed and the tunnel is not ready to be used. - NO_INCOMING_PACKETS: No incoming packets from peer. - REJECTED: Tunnel configuration was rejected, can be result of being denied access. - ALLOCATING_RESOURCES: Cloud VPN is in the process of allocating all required resources. - STOPPED: Tunnel is stopped due to its Forwarding Rules being deleted for Classic VPN tunnels or the project is in frozen state. - PEER_IDENTITY_MISMATCH: Peer identity does not match peer IP, probably behind NAT. - TS_NARROWING_NOT_ALLOWED: Traffic selector narrowing not allowed for an HA-VPN tunnel. 
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// URL of the Target VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created.
        /// </summary>
        [Output("targetVpnGateway")]
        public Output<string> TargetVpnGateway { get; private set; } = null!;

        /// <summary>
        /// URL of the VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created. This must be used (instead of target_vpn_gateway) if a High Availability VPN gateway resource is created.
        /// </summary>
        [Output("vpnGateway")]
        public Output<string> VpnGateway { get; private set; } = null!;

        /// <summary>
        /// The interface ID of the VPN gateway with which this VPN tunnel is associated.
        /// </summary>
        [Output("vpnGatewayInterface")]
        public Output<int> VpnGatewayInterface { get; private set; } = null!;


        /// <summary>
        /// Create a VpnTunnel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpnTunnel(string name, VpnTunnelArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:VpnTunnel", name, args ?? new VpnTunnelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpnTunnel(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:VpnTunnel", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpnTunnel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpnTunnel Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VpnTunnel(name, id, options);
        }
    }

    public sealed class VpnTunnelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// IKE protocol version to use when establishing the VPN tunnel with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default version is 2.
        /// </summary>
        [Input("ikeVersion")]
        public Input<int>? IkeVersion { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("localTrafficSelector")]
        private InputList<string>? _localTrafficSelector;

        /// <summary>
        /// Local traffic selector to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is supported.
        /// </summary>
        public InputList<string> LocalTrafficSelector
        {
            get => _localTrafficSelector ?? (_localTrafficSelector = new InputList<string>());
            set => _localTrafficSelector = value;
        }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// URL of the peer side external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field is exclusive with the field peerGcpGateway.
        /// </summary>
        [Input("peerExternalGateway")]
        public Input<string>? PeerExternalGateway { get; set; }

        /// <summary>
        /// The interface ID of the external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created.
        /// </summary>
        [Input("peerExternalGatewayInterface")]
        public Input<int>? PeerExternalGatewayInterface { get; set; }

        /// <summary>
        /// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field can be used when creating highly available VPN from VPC network to VPC network, the field is exclusive with the field peerExternalGateway. If provided, the VPN tunnel will automatically use the same vpnGatewayInterface ID in the peer GCP VPN gateway.
        /// </summary>
        [Input("peerGcpGateway")]
        public Input<string>? PeerGcpGateway { get; set; }

        /// <summary>
        /// IP address of the peer VPN gateway. Only IPv4 is supported.
        /// </summary>
        [Input("peerIp")]
        public Input<string>? PeerIp { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("remoteTrafficSelector")]
        private InputList<string>? _remoteTrafficSelector;

        /// <summary>
        /// Remote traffic selectors to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.
        /// </summary>
        public InputList<string> RemoteTrafficSelector
        {
            get => _remoteTrafficSelector ?? (_remoteTrafficSelector = new InputList<string>());
            set => _remoteTrafficSelector = value;
        }

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// URL of the router resource to be used for dynamic routing.
        /// </summary>
        [Input("router")]
        public Input<string>? Router { get; set; }

        /// <summary>
        /// Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
        /// </summary>
        [Input("sharedSecret")]
        public Input<string>? SharedSecret { get; set; }

        /// <summary>
        /// Hash of the shared secret.
        /// </summary>
        [Input("sharedSecretHash")]
        public Input<string>? SharedSecretHash { get; set; }

        /// <summary>
        /// URL of the Target VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created.
        /// </summary>
        [Input("targetVpnGateway")]
        public Input<string>? TargetVpnGateway { get; set; }

        /// <summary>
        /// URL of the VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created. This must be used (instead of target_vpn_gateway) if a High Availability VPN gateway resource is created.
        /// </summary>
        [Input("vpnGateway")]
        public Input<string>? VpnGateway { get; set; }

        /// <summary>
        /// The interface ID of the VPN gateway with which this VPN tunnel is associated.
        /// </summary>
        [Input("vpnGatewayInterface")]
        public Input<int>? VpnGatewayInterface { get; set; }

        public VpnTunnelArgs()
        {
        }
    }
}
