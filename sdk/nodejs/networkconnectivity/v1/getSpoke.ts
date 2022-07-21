// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets details about a Network Connectivity Center spoke.
 */
export function getSpoke(args: GetSpokeArgs, opts?: pulumi.InvokeOptions): Promise<GetSpokeResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:networkconnectivity/v1:getSpoke", {
        "location": args.location,
        "project": args.project,
        "spokeId": args.spokeId,
    }, opts);
}

export interface GetSpokeArgs {
    location: string;
    project?: string;
    spokeId: string;
}

export interface GetSpokeResult {
    /**
     * The time the spoke was created.
     */
    readonly createTime: string;
    /**
     * An optional description of the spoke.
     */
    readonly description: string;
    /**
     * Immutable. The name of the hub that this spoke is attached to.
     */
    readonly hub: string;
    /**
     * Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
     */
    readonly labels: {[key: string]: string};
    /**
     * VLAN attachments that are associated with the spoke.
     */
    readonly linkedInterconnectAttachments: outputs.networkconnectivity.v1.LinkedInterconnectAttachmentsResponse;
    /**
     * Router appliance instances that are associated with the spoke.
     */
    readonly linkedRouterApplianceInstances: outputs.networkconnectivity.v1.LinkedRouterApplianceInstancesResponse;
    /**
     * VPN tunnels that are associated with the spoke.
     */
    readonly linkedVpnTunnels: outputs.networkconnectivity.v1.LinkedVpnTunnelsResponse;
    /**
     * Immutable. The name of the spoke. Spoke names must be unique. They use the following form: `projects/{project_number}/locations/{region}/spokes/{spoke_id}`
     */
    readonly name: string;
    /**
     * The current lifecycle state of this spoke.
     */
    readonly state: string;
    /**
     * The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is deleted and another with the same name is created, the new spoke is assigned a different unique_id.
     */
    readonly uniqueId: string;
    /**
     * The time the spoke was last updated.
     */
    readonly updateTime: string;
}

export function getSpokeOutput(args: GetSpokeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSpokeResult> {
    return pulumi.output(args).apply(a => getSpoke(a, opts))
}

export interface GetSpokeOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    spokeId: pulumi.Input<string>;
}
