// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Returns the specified VpnTunnel resource.
 */
export function getVpnTunnel(args: GetVpnTunnelArgs, opts?: pulumi.InvokeOptions): Promise<GetVpnTunnelResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/alpha:getVpnTunnel", {
        "project": args.project,
        "region": args.region,
        "vpnTunnel": args.vpnTunnel,
    }, opts);
}

export interface GetVpnTunnelArgs {
    project?: string;
    region: string;
    vpnTunnel: string;
}

export interface GetVpnTunnelResult {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Detailed status message for the VPN tunnel.
     */
    readonly detailedStatus: string;
    /**
     * IKE protocol version to use when establishing the VPN tunnel with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default version is 2.
     */
    readonly ikeVersion: number;
    /**
     * Type of resource. Always compute#vpnTunnel for VPN tunnels.
     */
    readonly kind: string;
    /**
     * A fingerprint for the labels being applied to this VpnTunnel, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a VpnTunnel.
     */
    readonly labelFingerprint: string;
    /**
     * Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
     */
    readonly labels: {[key: string]: string};
    /**
     * Local traffic selector to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is supported.
     */
    readonly localTrafficSelector: string[];
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * URL of the peer side external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field is exclusive with the field peerGcpGateway.
     */
    readonly peerExternalGateway: string;
    /**
     * The interface ID of the external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. Possible values are: `0`, `1`, `2`, `3`. The number of IDs in use depends on the external VPN gateway redundancy type.
     */
    readonly peerExternalGatewayInterface: number;
    /**
     * URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field can be used when creating highly available VPN from VPC network to VPC network, the field is exclusive with the field peerExternalGateway. If provided, the VPN tunnel will automatically use the same vpnGatewayInterface ID in the peer GCP VPN gateway.
     */
    readonly peerGcpGateway: string;
    /**
     * IP address of the peer VPN gateway. Only IPv4 is supported.
     */
    readonly peerIp: string;
    /**
     * URL of the region where the VPN tunnel resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region: string;
    /**
     * Remote traffic selectors to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.
     */
    readonly remoteTrafficSelector: string[];
    /**
     * URL of the router resource to be used for dynamic routing.
     */
    readonly router: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
     */
    readonly sharedSecret: string;
    /**
     * Hash of the shared secret.
     */
    readonly sharedSecretHash: string;
    /**
     * The status of the VPN tunnel, which can be one of the following: - PROVISIONING: Resource is being allocated for the VPN tunnel. - WAITING_FOR_FULL_CONFIG: Waiting to receive all VPN-related configs from the user. Network, TargetVpnGateway, VpnTunnel, ForwardingRule, and Route resources are needed to setup the VPN tunnel. - FIRST_HANDSHAKE: Successful first handshake with the peer VPN. - ESTABLISHED: Secure session is successfully established with the peer VPN. - NETWORK_ERROR: Deprecated, replaced by NO_INCOMING_PACKETS - AUTHORIZATION_ERROR: Auth error (for example, bad shared secret). - NEGOTIATION_FAILURE: Handshake failed. - DEPROVISIONING: Resources are being deallocated for the VPN tunnel. - FAILED: Tunnel creation has failed and the tunnel is not ready to be used. - NO_INCOMING_PACKETS: No incoming packets from peer. - REJECTED: Tunnel configuration was rejected, can be result of being denied access. - ALLOCATING_RESOURCES: Cloud VPN is in the process of allocating all required resources. - STOPPED: Tunnel is stopped due to its Forwarding Rules being deleted for Classic VPN tunnels or the project is in frozen state. - PEER_IDENTITY_MISMATCH: Peer identity does not match peer IP, probably behind NAT. - TS_NARROWING_NOT_ALLOWED: Traffic selector narrowing not allowed for an HA-VPN tunnel. 
     */
    readonly status: string;
    /**
     * URL of the Target VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created.
     */
    readonly targetVpnGateway: string;
    /**
     * URL of the VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created. This must be used (instead of target_vpn_gateway) if a High Availability VPN gateway resource is created.
     */
    readonly vpnGateway: string;
    /**
     * The interface ID of the VPN gateway with which this VPN tunnel is associated. Possible values are: `0`, `1`.
     */
    readonly vpnGatewayInterface: number;
}
/**
 * Returns the specified VpnTunnel resource.
 */
export function getVpnTunnelOutput(args: GetVpnTunnelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpnTunnelResult> {
    return pulumi.output(args).apply((a: any) => getVpnTunnel(a, opts))
}

export interface GetVpnTunnelOutputArgs {
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    vpnTunnel: pulumi.Input<string>;
}
