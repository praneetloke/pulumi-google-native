// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a taxonomy.
 */
export function getTaxonomy(args: GetTaxonomyArgs, opts?: pulumi.InvokeOptions): Promise<GetTaxonomyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:datacatalog/v1:getTaxonomy", {
        "location": args.location,
        "project": args.project,
        "taxonomyId": args.taxonomyId,
    }, opts);
}

export interface GetTaxonomyArgs {
    location: string;
    project?: string;
    taxonomyId: string;
}

export interface GetTaxonomyResult {
    /**
     * Optional. A list of policy types that are activated for this taxonomy. If not set, defaults to an empty list.
     */
    readonly activatedPolicyTypes: string[];
    /**
     * Optional. Description of this taxonomy. If not set, defaults to empty. The description must contain only Unicode characters, tabs, newlines, carriage returns, and page breaks, and be at most 2000 bytes long when encoded in UTF-8.
     */
    readonly description: string;
    /**
     * User-defined name of this taxonomy. The name can't start or end with spaces, must contain only Unicode letters, numbers, underscores, dashes, and spaces, and be at most 200 bytes long when encoded in UTF-8. The taxonomy display name must be unique within an organization.
     */
    readonly displayName: string;
    /**
     * Resource name of this taxonomy in URL format. Note: Policy tag manager generates unique taxonomy IDs.
     */
    readonly name: string;
    /**
     * Number of policy tags in this taxonomy.
     */
    readonly policyTagCount: number;
    /**
     * Identity of the service which owns the Taxonomy. This field is only populated when the taxonomy is created by a Google Cloud service. Currently only 'DATAPLEX' is supported.
     */
    readonly service: outputs.datacatalog.v1.GoogleCloudDatacatalogV1TaxonomyServiceResponse;
    /**
     * Creation and modification timestamps of this taxonomy.
     */
    readonly taxonomyTimestamps: outputs.datacatalog.v1.GoogleCloudDatacatalogV1SystemTimestampsResponse;
}
/**
 * Gets a taxonomy.
 */
export function getTaxonomyOutput(args: GetTaxonomyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTaxonomyResult> {
    return pulumi.output(args).apply((a: any) => getTaxonomy(a, opts))
}

export interface GetTaxonomyOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    taxonomyId: pulumi.Input<string>;
}
