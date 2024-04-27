// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Retrieves a `NetworkPeering` resource by its resource name. The resource contains details of the network peering, such as peered networks, import and export custom route configurations, and peering state.
 */
export function getNetworkPeering(args: GetNetworkPeeringArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkPeeringResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:vmwareengine/v1:getNetworkPeering", {
        "networkPeeringId": args.networkPeeringId,
        "project": args.project,
    }, opts);
}

export interface GetNetworkPeeringArgs {
    networkPeeringId: string;
    project?: string;
}

export interface GetNetworkPeeringResult {
    /**
     * Creation time of this resource.
     */
    readonly createTime: string;
    /**
     * Optional. User-provided description for this network peering.
     */
    readonly description: string;
    /**
     * Optional. True if full mesh connectivity is created and managed automatically between peered networks; false otherwise. Currently this field is always true because Google Compute Engine automatically creates and manages subnetwork routes between two VPC networks when peering state is 'ACTIVE'.
     */
    readonly exchangeSubnetRoutes: boolean;
    /**
     * Optional. True if custom routes are exported to the peered network; false otherwise. The default value is true.
     */
    readonly exportCustomRoutes: boolean;
    /**
     * Optional. True if all subnet routes with a public IP address range are exported; false otherwise. The default value is true. IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
     */
    readonly exportCustomRoutesWithPublicIp: boolean;
    /**
     * Optional. True if custom routes are imported from the peered network; false otherwise. The default value is true.
     */
    readonly importCustomRoutes: boolean;
    /**
     * Optional. True if all subnet routes with public IP address range are imported; false otherwise. The default value is true. IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported to peers and are not controlled by this field.
     */
    readonly importCustomRoutesWithPublicIp: boolean;
    /**
     * The resource name of the network peering. Resource names are scheme-less URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. For example: `projects/my-project/locations/global/networkPeerings/my-peering`
     */
    readonly name: string;
    /**
     * Optional. Maximum transmission unit (MTU) in bytes. The default value is `1500`. If a value of `0` is provided for this field, VMware Engine uses the default value instead.
     */
    readonly peerMtu: number;
    /**
     * The relative resource name of the network to peer with a standard VMware Engine network. The provided network can be a consumer VPC network or another standard VMware Engine network. If the `peer_network_type` is VMWARE_ENGINE_NETWORK, specify the name in the form: `projects/{project}/locations/global/vmwareEngineNetworks/{vmware_engine_network_id}`. Otherwise specify the name in the form: `projects/{project}/global/networks/{network_id}`, where `{project}` can either be a project number or a project ID.
     */
    readonly peerNetwork: string;
    /**
     * The type of the network to peer with the VMware Engine network.
     */
    readonly peerNetworkType: string;
    /**
     * State of the network peering. This field has a value of 'ACTIVE' when there's a matching configuration in the peer network. New values may be added to this enum when appropriate.
     */
    readonly state: string;
    /**
     * Output Only. Details about the current state of the network peering.
     */
    readonly stateDetails: string;
    /**
     * System-generated unique identifier for the resource.
     */
    readonly uid: string;
    /**
     * Last update time of this resource.
     */
    readonly updateTime: string;
    /**
     * The relative resource name of the VMware Engine network. Specify the name in the following form: `projects/{project}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}` where `{project}` can either be a project number or a project ID.
     */
    readonly vmwareEngineNetwork: string;
}
/**
 * Retrieves a `NetworkPeering` resource by its resource name. The resource contains details of the network peering, such as peered networks, import and export custom route configurations, and peering state.
 */
export function getNetworkPeeringOutput(args: GetNetworkPeeringOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkPeeringResult> {
    return pulumi.output(args).apply((a: any) => getNetworkPeering(a, opts))
}

export interface GetNetworkPeeringOutputArgs {
    networkPeeringId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}