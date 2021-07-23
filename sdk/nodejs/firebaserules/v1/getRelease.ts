// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Get a `Release` by name.
 */
export function getRelease(args: GetReleaseArgs, opts?: pulumi.InvokeOptions): Promise<GetReleaseResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:firebaserules/v1:getRelease", {
        "project": args.project,
        "releaseId": args.releaseId,
    }, opts);
}

export interface GetReleaseArgs {
    project?: string;
    releaseId: string;
}

export interface GetReleaseResult {
    /**
     * Time the release was created.
     */
    readonly createTime: string;
    /**
     * Format: `projects/{project_id}/releases/{release_id}`
     */
    readonly name: string;
    /**
     * Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist the `Release` to be created.
     */
    readonly rulesetName: string;
    /**
     * Time the release was updated.
     */
    readonly updateTime: string;
}