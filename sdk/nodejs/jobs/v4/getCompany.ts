// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Retrieves specified company.
 */
export function getCompany(args: GetCompanyArgs, opts?: pulumi.InvokeOptions): Promise<GetCompanyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:jobs/v4:getCompany", {
        "companyId": args.companyId,
        "project": args.project,
        "tenantId": args.tenantId,
    }, opts);
}

export interface GetCompanyArgs {
    companyId: string;
    project?: string;
    tenantId: string;
}

export interface GetCompanyResult {
    /**
     * The URI to employer's career site or careers page on the employer's web site, for example, "https://careers.google.com".
     */
    readonly careerSiteUri: string;
    /**
     * Derived details about the company.
     */
    readonly derivedInfo: outputs.jobs.v4.CompanyDerivedInfoResponse;
    /**
     * The display name of the company, for example, "Google LLC".
     */
    readonly displayName: string;
    /**
     * Equal Employment Opportunity legal disclaimer text to be associated with all jobs, and typically to be displayed in all roles. The maximum number of allowed characters is 500.
     */
    readonly eeoText: string;
    /**
     * Client side company identifier, used to uniquely identify the company. The maximum number of allowed characters is 255.
     */
    readonly externalId: string;
    /**
     * The street address of the company's main headquarters, which may be different from the job location. The service attempts to geolocate the provided address, and populates a more specific location wherever possible in DerivedInfo.headquarters_location.
     */
    readonly headquartersAddress: string;
    /**
     * Set to true if it is the hiring agency that post jobs for other employers. Defaults to false if not provided.
     */
    readonly hiringAgency: boolean;
    /**
     * A URI that hosts the employer's company logo.
     */
    readonly imageUri: string;
    /**
     * This field is deprecated. Please set the searchability of the custom attribute in the Job.custom_attributes going forward. A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword searches. Jobs with `string_values` under these specified field keys are returned if any of the values match the search keyword. Custom field values with parenthesis, brackets and special symbols are not searchable as-is, and those keyword queries must be surrounded by quotes.
     *
     * @deprecated This field is deprecated. Please set the searchability of the custom attribute in the Job.custom_attributes going forward. A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword searches. Jobs with `string_values` under these specified field keys are returned if any of the values match the search keyword. Custom field values with parenthesis, brackets and special symbols are not searchable as-is, and those keyword queries must be surrounded by quotes.
     */
    readonly keywordSearchableJobCustomAttributes: string[];
    /**
     * Required during company update. The resource name for a company. This is generated by the service when a company is created. The format is "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for example, "projects/foo/tenants/bar/companies/baz".
     */
    readonly name: string;
    /**
     * The employer's company size.
     */
    readonly size: string;
    /**
     * Indicates whether a company is flagged to be suspended from public availability by the service when job content appears suspicious, abusive, or spammy.
     */
    readonly suspended: boolean;
    /**
     * The URI representing the company's primary web site or home page, for example, "https://www.google.com". The maximum number of allowed characters is 255.
     */
    readonly websiteUri: string;
}

export function getCompanyOutput(args: GetCompanyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCompanyResult> {
    return pulumi.output(args).apply(a => getCompany(a, opts))
}

export interface GetCompanyOutputArgs {
    companyId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    tenantId: pulumi.Input<string>;
}
