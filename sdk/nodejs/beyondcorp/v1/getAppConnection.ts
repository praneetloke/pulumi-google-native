// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets details of a single AppConnection.
 */
export function getAppConnection(args: GetAppConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetAppConnectionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:beyondcorp/v1:getAppConnection", {
        "appConnectionId": args.appConnectionId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetAppConnectionArgs {
    appConnectionId: string;
    location: string;
    project?: string;
}

export interface GetAppConnectionResult {
    /**
     * Address of the remote application endpoint for the BeyondCorp AppConnection.
     */
    readonly applicationEndpoint: outputs.beyondcorp.v1.GoogleCloudBeyondcorpAppconnectionsV1AppConnectionApplicationEndpointResponse;
    /**
     * Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this AppConnection.
     */
    readonly connectors: string[];
    /**
     * Timestamp when the resource was created.
     */
    readonly createTime: string;
    /**
     * Optional. An arbitrary user-provided name for the AppConnection. Cannot exceed 64 characters.
     */
    readonly displayName: string;
    /**
     * Optional. Gateway used by the AppConnection.
     */
    readonly gateway: outputs.beyondcorp.v1.GoogleCloudBeyondcorpAppconnectionsV1AppConnectionGatewayResponse;
    /**
     * Optional. Resource labels to represent user provided metadata.
     */
    readonly labels: {[key: string]: string};
    /**
     * Unique resource name of the AppConnection. The name is ignored when creating a AppConnection.
     */
    readonly name: string;
    /**
     * The current state of the AppConnection.
     */
    readonly state: string;
    /**
     * The type of network connectivity used by the AppConnection.
     */
    readonly type: string;
    /**
     * A unique identifier for the instance generated by the system.
     */
    readonly uid: string;
    /**
     * Timestamp when the resource was last modified.
     */
    readonly updateTime: string;
}

export function getAppConnectionOutput(args: GetAppConnectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppConnectionResult> {
    return pulumi.output(args).apply(a => getAppConnection(a, opts))
}

export interface GetAppConnectionOutputArgs {
    appConnectionId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
