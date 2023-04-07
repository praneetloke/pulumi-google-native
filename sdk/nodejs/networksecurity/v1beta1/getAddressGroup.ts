// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets details of a single address group.
 */
export function getAddressGroup(args: GetAddressGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetAddressGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:networksecurity/v1beta1:getAddressGroup", {
        "addressGroupId": args.addressGroupId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetAddressGroupArgs {
    addressGroupId: string;
    location: string;
    project?: string;
}

export interface GetAddressGroupResult {
    /**
     * Capacity of the Address Group
     */
    readonly capacity: number;
    /**
     * The timestamp when the resource was created.
     */
    readonly createTime: string;
    /**
     * Optional. Free-text description of the resource.
     */
    readonly description: string;
    /**
     * Optional. List of items.
     */
    readonly items: string[];
    /**
     * Optional. Set of label tags associated with the AddressGroup resource.
     */
    readonly labels: {[key: string]: string};
    /**
     * Name of the AddressGroup resource. It matches pattern `projects/*&#47;locations/{location}/addressGroups/`.
     */
    readonly name: string;
    /**
     * Server-defined fully-qualified URL for this resource.
     */
    readonly selfLink: string;
    /**
     * The type of the Address Group. Possible values are "IPv4" or "IPV6".
     */
    readonly type: string;
    /**
     * The timestamp when the resource was updated.
     */
    readonly updateTime: string;
}
/**
 * Gets details of a single address group.
 */
export function getAddressGroupOutput(args: GetAddressGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAddressGroupResult> {
    return pulumi.output(args).apply((a: any) => getAddressGroup(a, opts))
}

export interface GetAddressGroupOutputArgs {
    addressGroupId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
