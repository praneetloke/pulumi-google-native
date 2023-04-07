// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Get a group.
 */
export function getGroup(args: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:firebaseappdistribution/v1:getGroup", {
        "groupId": args.groupId,
        "project": args.project,
    }, opts);
}

export interface GetGroupArgs {
    groupId: string;
    project?: string;
}

export interface GetGroupResult {
    /**
     * The display name of the group.
     */
    readonly displayName: string;
    /**
     * The number of invite links for this group.
     */
    readonly inviteLinkCount: number;
    /**
     * The name of the group resource. Format: `projects/{project_number}/groups/{group_alias}`
     */
    readonly name: string;
    /**
     * The number of releases this group is permitted to access.
     */
    readonly releaseCount: number;
    /**
     * The number of testers who are members of this group.
     */
    readonly testerCount: number;
}
/**
 * Get a group.
 */
export function getGroupOutput(args: GetGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupResult> {
    return pulumi.output(args).apply((a: any) => getGroup(a, opts))
}

export interface GetGroupOutputArgs {
    groupId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
