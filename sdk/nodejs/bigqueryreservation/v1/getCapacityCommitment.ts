// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns information about the capacity commitment.
 */
export function getCapacityCommitment(args: GetCapacityCommitmentArgs, opts?: pulumi.InvokeOptions): Promise<GetCapacityCommitmentResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:bigqueryreservation/v1:getCapacityCommitment", {
        "capacityCommitmentId": args.capacityCommitmentId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetCapacityCommitmentArgs {
    capacityCommitmentId: string;
    location: string;
    project?: string;
}

export interface GetCapacityCommitmentResult {
    /**
     * The end of the current commitment period. It is applicable only for ACTIVE capacity commitments.
     */
    readonly commitmentEndTime: string;
    /**
     * The start of the current commitment period. It is applicable only for ACTIVE capacity commitments.
     */
    readonly commitmentStartTime: string;
    /**
     * Edition of the capacity commitment.
     */
    readonly edition: string;
    /**
     * For FAILED commitment plan, provides the reason of failure.
     */
    readonly failureStatus: outputs.bigqueryreservation.v1.StatusResponse;
    /**
     * Applicable only for commitments located within one of the BigQuery multi-regions (US or EU). If set to true, this commitment is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this commitment is placed in the organization's default region. NOTE: this is a preview feature. Project must be allow-listed in order to set this field.
     */
    readonly multiRegionAuxiliary: boolean;
    /**
     * The resource name of the capacity commitment, e.g., `projects/myproject/locations/US/capacityCommitments/123` The commitment_id must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
     */
    readonly name: string;
    /**
     * Capacity commitment commitment plan.
     */
    readonly plan: string;
    /**
     * The plan this capacity commitment is converted to after commitment_end_time passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable for ANNUAL and TRIAL commitments.
     */
    readonly renewalPlan: string;
    /**
     * Number of slots in this commitment.
     */
    readonly slotCount: string;
    /**
     * State of the commitment.
     */
    readonly state: string;
}
/**
 * Returns information about the capacity commitment.
 */
export function getCapacityCommitmentOutput(args: GetCapacityCommitmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCapacityCommitmentResult> {
    return pulumi.output(args).apply((a: any) => getCapacityCommitment(a, opts))
}

export interface GetCapacityCommitmentOutputArgs {
    capacityCommitmentId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
