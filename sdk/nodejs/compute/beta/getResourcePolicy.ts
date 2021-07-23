// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves all information of the specified resource policy.
 */
export function getResourcePolicy(args: GetResourcePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetResourcePolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:compute/beta:getResourcePolicy", {
        "project": args.project,
        "region": args.region,
        "resourcePolicy": args.resourcePolicy,
    }, opts);
}

export interface GetResourcePolicyArgs {
    project?: string;
    region: string;
    resourcePolicy: string;
}

export interface GetResourcePolicyResult {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    readonly description: string;
    /**
     * Resource policy for instances for placement configuration.
     */
    readonly groupPlacementPolicy: outputs.compute.beta.ResourcePolicyGroupPlacementPolicyResponse;
    /**
     * Resource policy for scheduling instance operations.
     */
    readonly instanceSchedulePolicy: outputs.compute.beta.ResourcePolicyInstanceSchedulePolicyResponse;
    /**
     * Type of the resource. Always compute#resource_policies for resource policies.
     */
    readonly kind: string;
    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    readonly region: string;
    /**
     * The system status of the resource policy.
     */
    readonly resourceStatus: outputs.compute.beta.ResourcePolicyResourceStatusResponse;
    /**
     * Server-defined fully-qualified URL for this resource.
     */
    readonly selfLink: string;
    /**
     * Resource policy for persistent disks for creating snapshots.
     */
    readonly snapshotSchedulePolicy: outputs.compute.beta.ResourcePolicySnapshotSchedulePolicyResponse;
    /**
     * The status of resource policy creation.
     */
    readonly status: string;
}