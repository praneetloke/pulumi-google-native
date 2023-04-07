// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns all of the details about the specified resize request.
 */
export function getInstanceGroupManagerResizeRequest(args: GetInstanceGroupManagerResizeRequestArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceGroupManagerResizeRequestResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/alpha:getInstanceGroupManagerResizeRequest", {
        "instanceGroupManager": args.instanceGroupManager,
        "project": args.project,
        "resizeRequest": args.resizeRequest,
        "zone": args.zone,
    }, opts);
}

export interface GetInstanceGroupManagerResizeRequestArgs {
    instanceGroupManager: string;
    project?: string;
    resizeRequest: string;
    zone: string;
}

export interface GetInstanceGroupManagerResizeRequestResult {
    /**
     * The count of instances to create as part of this resize request.
     */
    readonly count: number;
    /**
     * The creation timestamp for this resize request in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource.
     */
    readonly description: string;
    /**
     * The resource type, which is always compute#instanceGroupManagerResizeRequest for resize requests.
     */
    readonly kind: string;
    /**
     * The name of this resize request. The name must be 1-63 characters long, and comply with RFC1035.
     */
    readonly name: string;
    /**
     * When set, defines queing parameters for the requested deferred capacity. When unset, the request starts provisioning immediately, or fails if immediate provisioning is not possible.
     */
    readonly queuingPolicy: outputs.compute.alpha.QueuingPolicyResponse;
    /**
     * The URL for this resize request. The server defines this URL.
     */
    readonly selfLink: string;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    readonly selfLinkWithId: string;
    /**
     * [Output only] Current state of the request.
     */
    readonly state: string;
    /**
     * [Output only] Status of the request. The Status message is aligned with QueuedResource.status. ResizeRequest.queuing_policy contains the queuing policy as provided by the user; it could have either valid_until_time or valid_until_duration. ResizeRequest.status.queuing_policy always contains absolute time as calculated by the server when the request is queued.
     */
    readonly status: outputs.compute.alpha.InstanceGroupManagerResizeRequestStatusResponse;
    /**
     * The URL of a zone where the resize request is located. Populated only for zonal resize requests.
     */
    readonly zone: string;
}
/**
 * Returns all of the details about the specified resize request.
 */
export function getInstanceGroupManagerResizeRequestOutput(args: GetInstanceGroupManagerResizeRequestOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceGroupManagerResizeRequestResult> {
    return pulumi.output(args).apply((a: any) => getInstanceGroupManagerResizeRequest(a, opts))
}

export interface GetInstanceGroupManagerResizeRequestOutputArgs {
    instanceGroupManager: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    resizeRequest: pulumi.Input<string>;
    zone: pulumi.Input<string>;
}
