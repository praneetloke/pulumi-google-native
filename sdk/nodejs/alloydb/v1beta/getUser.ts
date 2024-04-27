// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets details of a single User.
 */
export function getUser(args: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:alloydb/v1beta:getUser", {
        "clusterId": args.clusterId,
        "location": args.location,
        "project": args.project,
        "userId": args.userId,
    }, opts);
}

export interface GetUserArgs {
    clusterId: string;
    location: string;
    project?: string;
    userId: string;
}

export interface GetUserResult {
    /**
     * Optional. List of database roles this user has. The database role strings are subject to the PostgreSQL naming conventions.
     */
    readonly databaseRoles: string[];
    /**
     * Name of the resource in the form of projects/{project}/locations/{location}/cluster/{cluster}/users/{user}.
     */
    readonly name: string;
    /**
     * Input only. Password for the user.
     */
    readonly password: string;
    /**
     * Optional. Type of this user.
     */
    readonly userType: string;
}
/**
 * Gets details of a single User.
 */
export function getUserOutput(args: GetUserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserResult> {
    return pulumi.output(args).apply((a: any) => getUser(a, opts))
}

export interface GetUserOutputArgs {
    clusterId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    userId: pulumi.Input<string>;
}