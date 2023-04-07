// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified commitment resource.
 */
export function getRegionCommitment(args: GetRegionCommitmentArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionCommitmentResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/beta:getRegionCommitment", {
        "commitment": args.commitment,
        "project": args.project,
        "region": args.region,
    }, opts);
}

export interface GetRegionCommitmentArgs {
    commitment: string;
    project?: string;
    region: string;
}

export interface GetRegionCommitmentResult {
    /**
     * Specifies whether to enable automatic renewal for the commitment. The default value is false if not specified. The field can be updated until the day of the commitment expiration at 12:00am PST. If the field is set to true, the commitment will be automatically renewed for either one or three years according to the terms of the existing commitment.
     */
    readonly autoRenew: boolean;
    /**
     * The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
     */
    readonly category: string;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Commitment end time in RFC3339 text format.
     */
    readonly endTimestamp: string;
    /**
     * Type of the resource. Always compute#commitment for commitments.
     */
    readonly kind: string;
    /**
     * The license specification required as part of a license commitment.
     */
    readonly licenseResource: outputs.compute.beta.LicenseResourceCommitmentResponse;
    /**
     * List of source commitments to be merged into a new commitment.
     */
    readonly mergeSourceCommitments: string[];
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
     */
    readonly plan: string;
    /**
     * URL of the region where this commitment may be used.
     */
    readonly region: string;
    /**
     * List of reservations in this commitment.
     */
    readonly reservations: outputs.compute.beta.ReservationResponse[];
    /**
     * A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
     */
    readonly resources: outputs.compute.beta.ResourceCommitmentResponse[];
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * Source commitment to be splitted into a new commitment.
     */
    readonly splitSourceCommitment: string;
    /**
     * Commitment start time in RFC3339 text format.
     */
    readonly startTimestamp: string;
    /**
     * Status of the commitment with regards to eventual expiration (each commitment has an end date defined). One of the following values: NOT_YET_ACTIVE, ACTIVE, EXPIRED.
     */
    readonly status: string;
    /**
     * An optional, human-readable explanation of the status.
     */
    readonly statusMessage: string;
    /**
     * The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
     */
    readonly type: string;
}
/**
 * Returns the specified commitment resource.
 */
export function getRegionCommitmentOutput(args: GetRegionCommitmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionCommitmentResult> {
    return pulumi.output(args).apply((a: any) => getRegionCommitment(a, opts))
}

export interface GetRegionCommitmentOutputArgs {
    commitment: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
}
