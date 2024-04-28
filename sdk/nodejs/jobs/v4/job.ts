// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new job. Typically, the job becomes searchable within 10 seconds, but it may take up to 5 minutes.
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:jobs/v4:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * Strongly recommended for the best service experience. Location(s) where the employer is looking to hire for this job posting. Specifying the full street address(es) of the hiring location enables better API results, especially job searches by commute time. At most 50 locations are allowed for best search performance. If a job has more locations, it is suggested to split it into multiple jobs with unique requisition_ids (e.g. 'ReqA' becomes 'ReqA-1', 'ReqA-2', and so on.) as multiple jobs with the same company, language_code and requisition_id are not allowed. If the original requisition_id must be preserved, a custom field should be used for storage. It is also suggested to group the locations that close to each other in the same job for better search experience. Jobs with multiple addresses must have their addresses with the same LocationType to allow location filtering to work properly. (For example, a Job with addresses "1600 Amphitheatre Parkway, Mountain View, CA, USA" and "London, UK" may not have location filters applied correctly at search time since the first is a LocationType.STREET_ADDRESS and the second is a LocationType.LOCALITY.) If a job needs to have multiple addresses, it is suggested to split it into multiple jobs with same LocationTypes. The maximum number of allowed characters is 500.
     */
    public readonly addresses!: pulumi.Output<string[]>;
    /**
     * Job application information.
     */
    public readonly applicationInfo!: pulumi.Output<outputs.jobs.v4.ApplicationInfoResponse>;
    /**
     * The resource name of the company listing the job. The format is "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}". For example, "projects/foo/tenants/bar/companies/baz".
     */
    public readonly company!: pulumi.Output<string>;
    /**
     * Display name of the company listing the job.
     */
    public /*out*/ readonly companyDisplayName!: pulumi.Output<string>;
    /**
     * Job compensation information (a.k.a. "pay rate") i.e., the compensation that will paid to the employee.
     */
    public readonly compensationInfo!: pulumi.Output<outputs.jobs.v4.CompensationInfoResponse>;
    /**
     * A map of fields to hold both filterable and non-filterable custom job attributes that are not covered by the provided structured fields. The keys of the map are strings up to 64 bytes and must match the pattern: `a-zA-Z*`. For example, key0LikeThis or KEY_1_LIKE_THIS. At most 100 filterable and at most 100 unfilterable keys are supported. For filterable `string_values`, across all keys at most 200 values are allowed, with each string no more than 255 characters. For unfilterable `string_values`, the maximum total size of `string_values` across all keys is 50KB.
     */
    public readonly customAttributes!: pulumi.Output<outputs.jobs.v4.CustomAttributeResponse>;
    /**
     * The desired education degrees for the job, such as Bachelors, Masters.
     */
    public readonly degreeTypes!: pulumi.Output<string[]>;
    /**
     * The department or functional area within the company with the open position. The maximum number of allowed characters is 255.
     */
    public readonly department!: pulumi.Output<string>;
    /**
     * Derived details about the job posting.
     */
    public /*out*/ readonly derivedInfo!: pulumi.Output<outputs.jobs.v4.JobDerivedInfoResponse>;
    /**
     * The description of the job, which typically includes a multi-paragraph description of the company and related information. Separate fields are provided on the job object for responsibilities, qualifications, and other job characteristics. Use of these separate job fields is recommended. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 100,000.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The employment type(s) of a job, for example, full time or part time.
     */
    public readonly employmentTypes!: pulumi.Output<string[]>;
    /**
     * A description of bonus, commission, and other compensation incentives associated with the job not including salary or pay. The maximum number of allowed characters is 10,000.
     */
    public readonly incentives!: pulumi.Output<string>;
    /**
     * The benefits included with the job.
     */
    public readonly jobBenefits!: pulumi.Output<string[]>;
    /**
     * The end timestamp of the job. Typically this field is used for contracting engagements. Invalid timestamps are ignored.
     */
    public readonly jobEndTime!: pulumi.Output<string>;
    /**
     * The experience level associated with the job, such as "Entry Level".
     */
    public readonly jobLevel!: pulumi.Output<string>;
    /**
     * The start timestamp of the job in UTC time zone. Typically this field is used for contracting engagements. Invalid timestamps are ignored.
     */
    public readonly jobStartTime!: pulumi.Output<string>;
    /**
     * The language of the posting. This field is distinct from any requirements for fluency that are associated with the job. Language codes must be in BCP-47 format, such as "en-US" or "sr-Latn". For more information, see [Tags for Identifying Languages](https://tools.ietf.org/html/bcp47){: class="external" target="_blank" }. If this field is unspecified and Job.description is present, detected language code based on Job.description is assigned, otherwise defaults to 'en_US'.
     */
    public readonly languageCode!: pulumi.Output<string>;
    /**
     * Required during job update. The resource name for the job. This is generated by the service when a job is created. The format is "projects/{project_id}/tenants/{tenant_id}/jobs/{job_id}". For example, "projects/foo/tenants/bar/jobs/baz". Use of this field in job queries and API calls is preferred over the use of requisition_id since this value is unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The timestamp when this job posting was created.
     */
    public /*out*/ readonly postingCreateTime!: pulumi.Output<string>;
    /**
     * Strongly recommended for the best service experience. The expiration timestamp of the job. After this timestamp, the job is marked as expired, and it no longer appears in search results. The expired job can't be listed by the ListJobs API, but it can be retrieved with the GetJob API or updated with the UpdateJob API or deleted with the DeleteJob API. An expired job can be updated and opened again by using a future expiration timestamp. Updating an expired job fails if there is another existing open job with same company, language_code and requisition_id. The expired jobs are retained in our system for 90 days. However, the overall expired job count cannot exceed 3 times the maximum number of open jobs over previous 7 days. If this threshold is exceeded, expired jobs are cleaned out in order of earliest expire time. Expired jobs are no longer accessible after they are cleaned out. Invalid timestamps are ignored, and treated as expire time not provided. If the timestamp is before the instant request is made, the job is treated as expired immediately on creation. This kind of job can not be updated. And when creating a job with past timestamp, the posting_publish_time must be set before posting_expire_time. The purpose of this feature is to allow other objects, such as Application, to refer a job that didn't exist in the system prior to becoming expired. If you want to modify a job that was expired on creation, delete it and create a new one. If this value isn't provided at the time of job creation or is invalid, the job posting expires after 30 days from the job's creation time. For example, if the job was created on 2017/01/01 13:00AM UTC with an unspecified expiration date, the job expires after 2017/01/31 13:00AM UTC. If this value isn't provided on job update, it depends on the field masks set by UpdateJobRequest.update_mask. If the field masks include job_end_time, or the masks are empty meaning that every field is updated, the job posting expires after 30 days from the job's last update time. Otherwise the expiration date isn't updated.
     */
    public readonly postingExpireTime!: pulumi.Output<string>;
    /**
     * The timestamp this job posting was most recently published. The default value is the time the request arrives at the server. Invalid timestamps are ignored.
     */
    public readonly postingPublishTime!: pulumi.Output<string>;
    /**
     * The job PostingRegion (for example, state, country) throughout which the job is available. If this field is set, a LocationFilter in a search query within the job region finds this job posting if an exact location match isn't specified. If this field is set to PostingRegion.NATION or PostingRegion.ADMINISTRATIVE_AREA, setting job Job.addresses to the same location level as this field is strongly recommended.
     */
    public readonly postingRegion!: pulumi.Output<string>;
    /**
     * The timestamp when this job posting was last updated.
     */
    public /*out*/ readonly postingUpdateTime!: pulumi.Output<string>;
    /**
     * Options for job processing.
     */
    public readonly processingOptions!: pulumi.Output<outputs.jobs.v4.ProcessingOptionsResponse>;
    public readonly project!: pulumi.Output<string>;
    /**
     * A promotion value of the job, as determined by the client. The value determines the sort order of the jobs returned when searching for jobs using the featured jobs search call, with higher promotional values being returned first and ties being resolved by relevance sort. Only the jobs with a promotionValue >0 are returned in a FEATURED_JOB_SEARCH. Default value is 0, and negative values are treated as 0.
     */
    public readonly promotionValue!: pulumi.Output<number>;
    /**
     * A description of the qualifications required to perform the job. The use of this field is recommended as an alternative to using the more general description field. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 10,000.
     */
    public readonly qualifications!: pulumi.Output<string>;
    /**
     * The requisition ID, also referred to as the posting ID, is assigned by the client to identify a job. This field is intended to be used by clients for client identification and tracking of postings. A job isn't allowed to be created if there is another job with the same company, language_code and requisition_id. The maximum number of allowed characters is 255.
     */
    public readonly requisitionId!: pulumi.Output<string>;
    /**
     * A description of job responsibilities. The use of this field is recommended as an alternative to using the more general description field. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 10,000.
     */
    public readonly responsibilities!: pulumi.Output<string>;
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * The title of the job, such as "Software Engineer" The maximum number of allowed characters is 500.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * Deprecated. The job is only visible to the owner. The visibility of the job. Defaults to Visibility.ACCOUNT_ONLY if not specified.
     *
     * @deprecated Deprecated. The job is only visible to the owner. The visibility of the job. Defaults to Visibility.ACCOUNT_ONLY if not specified.
     */
    public readonly visibility!: pulumi.Output<string>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.company === undefined) && !opts.urn) {
                throw new Error("Missing required property 'company'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.requisitionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requisitionId'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            resourceInputs["addresses"] = args ? args.addresses : undefined;
            resourceInputs["applicationInfo"] = args ? args.applicationInfo : undefined;
            resourceInputs["company"] = args ? args.company : undefined;
            resourceInputs["compensationInfo"] = args ? args.compensationInfo : undefined;
            resourceInputs["customAttributes"] = args ? args.customAttributes : undefined;
            resourceInputs["degreeTypes"] = args ? args.degreeTypes : undefined;
            resourceInputs["department"] = args ? args.department : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["employmentTypes"] = args ? args.employmentTypes : undefined;
            resourceInputs["incentives"] = args ? args.incentives : undefined;
            resourceInputs["jobBenefits"] = args ? args.jobBenefits : undefined;
            resourceInputs["jobEndTime"] = args ? args.jobEndTime : undefined;
            resourceInputs["jobLevel"] = args ? args.jobLevel : undefined;
            resourceInputs["jobStartTime"] = args ? args.jobStartTime : undefined;
            resourceInputs["languageCode"] = args ? args.languageCode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["postingExpireTime"] = args ? args.postingExpireTime : undefined;
            resourceInputs["postingPublishTime"] = args ? args.postingPublishTime : undefined;
            resourceInputs["postingRegion"] = args ? args.postingRegion : undefined;
            resourceInputs["processingOptions"] = args ? args.processingOptions : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["promotionValue"] = args ? args.promotionValue : undefined;
            resourceInputs["qualifications"] = args ? args.qualifications : undefined;
            resourceInputs["requisitionId"] = args ? args.requisitionId : undefined;
            resourceInputs["responsibilities"] = args ? args.responsibilities : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
            resourceInputs["companyDisplayName"] = undefined /*out*/;
            resourceInputs["derivedInfo"] = undefined /*out*/;
            resourceInputs["postingCreateTime"] = undefined /*out*/;
            resourceInputs["postingUpdateTime"] = undefined /*out*/;
        } else {
            resourceInputs["addresses"] = undefined /*out*/;
            resourceInputs["applicationInfo"] = undefined /*out*/;
            resourceInputs["company"] = undefined /*out*/;
            resourceInputs["companyDisplayName"] = undefined /*out*/;
            resourceInputs["compensationInfo"] = undefined /*out*/;
            resourceInputs["customAttributes"] = undefined /*out*/;
            resourceInputs["degreeTypes"] = undefined /*out*/;
            resourceInputs["department"] = undefined /*out*/;
            resourceInputs["derivedInfo"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["employmentTypes"] = undefined /*out*/;
            resourceInputs["incentives"] = undefined /*out*/;
            resourceInputs["jobBenefits"] = undefined /*out*/;
            resourceInputs["jobEndTime"] = undefined /*out*/;
            resourceInputs["jobLevel"] = undefined /*out*/;
            resourceInputs["jobStartTime"] = undefined /*out*/;
            resourceInputs["languageCode"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["postingCreateTime"] = undefined /*out*/;
            resourceInputs["postingExpireTime"] = undefined /*out*/;
            resourceInputs["postingPublishTime"] = undefined /*out*/;
            resourceInputs["postingRegion"] = undefined /*out*/;
            resourceInputs["postingUpdateTime"] = undefined /*out*/;
            resourceInputs["processingOptions"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["promotionValue"] = undefined /*out*/;
            resourceInputs["qualifications"] = undefined /*out*/;
            resourceInputs["requisitionId"] = undefined /*out*/;
            resourceInputs["responsibilities"] = undefined /*out*/;
            resourceInputs["tenantId"] = undefined /*out*/;
            resourceInputs["title"] = undefined /*out*/;
            resourceInputs["visibility"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["project", "tenantId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Job.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * Strongly recommended for the best service experience. Location(s) where the employer is looking to hire for this job posting. Specifying the full street address(es) of the hiring location enables better API results, especially job searches by commute time. At most 50 locations are allowed for best search performance. If a job has more locations, it is suggested to split it into multiple jobs with unique requisition_ids (e.g. 'ReqA' becomes 'ReqA-1', 'ReqA-2', and so on.) as multiple jobs with the same company, language_code and requisition_id are not allowed. If the original requisition_id must be preserved, a custom field should be used for storage. It is also suggested to group the locations that close to each other in the same job for better search experience. Jobs with multiple addresses must have their addresses with the same LocationType to allow location filtering to work properly. (For example, a Job with addresses "1600 Amphitheatre Parkway, Mountain View, CA, USA" and "London, UK" may not have location filters applied correctly at search time since the first is a LocationType.STREET_ADDRESS and the second is a LocationType.LOCALITY.) If a job needs to have multiple addresses, it is suggested to split it into multiple jobs with same LocationTypes. The maximum number of allowed characters is 500.
     */
    addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Job application information.
     */
    applicationInfo?: pulumi.Input<inputs.jobs.v4.ApplicationInfoArgs>;
    /**
     * The resource name of the company listing the job. The format is "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}". For example, "projects/foo/tenants/bar/companies/baz".
     */
    company: pulumi.Input<string>;
    /**
     * Job compensation information (a.k.a. "pay rate") i.e., the compensation that will paid to the employee.
     */
    compensationInfo?: pulumi.Input<inputs.jobs.v4.CompensationInfoArgs>;
    /**
     * A map of fields to hold both filterable and non-filterable custom job attributes that are not covered by the provided structured fields. The keys of the map are strings up to 64 bytes and must match the pattern: `a-zA-Z*`. For example, key0LikeThis or KEY_1_LIKE_THIS. At most 100 filterable and at most 100 unfilterable keys are supported. For filterable `string_values`, across all keys at most 200 values are allowed, with each string no more than 255 characters. For unfilterable `string_values`, the maximum total size of `string_values` across all keys is 50KB.
     */
    customAttributes?: pulumi.Input<inputs.jobs.v4.CustomAttributeArgs>;
    /**
     * The desired education degrees for the job, such as Bachelors, Masters.
     */
    degreeTypes?: pulumi.Input<pulumi.Input<enums.jobs.v4.JobDegreeTypesItem>[]>;
    /**
     * The department or functional area within the company with the open position. The maximum number of allowed characters is 255.
     */
    department?: pulumi.Input<string>;
    /**
     * The description of the job, which typically includes a multi-paragraph description of the company and related information. Separate fields are provided on the job object for responsibilities, qualifications, and other job characteristics. Use of these separate job fields is recommended. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 100,000.
     */
    description: pulumi.Input<string>;
    /**
     * The employment type(s) of a job, for example, full time or part time.
     */
    employmentTypes?: pulumi.Input<pulumi.Input<enums.jobs.v4.JobEmploymentTypesItem>[]>;
    /**
     * A description of bonus, commission, and other compensation incentives associated with the job not including salary or pay. The maximum number of allowed characters is 10,000.
     */
    incentives?: pulumi.Input<string>;
    /**
     * The benefits included with the job.
     */
    jobBenefits?: pulumi.Input<pulumi.Input<enums.jobs.v4.JobJobBenefitsItem>[]>;
    /**
     * The end timestamp of the job. Typically this field is used for contracting engagements. Invalid timestamps are ignored.
     */
    jobEndTime?: pulumi.Input<string>;
    /**
     * The experience level associated with the job, such as "Entry Level".
     */
    jobLevel?: pulumi.Input<enums.jobs.v4.JobJobLevel>;
    /**
     * The start timestamp of the job in UTC time zone. Typically this field is used for contracting engagements. Invalid timestamps are ignored.
     */
    jobStartTime?: pulumi.Input<string>;
    /**
     * The language of the posting. This field is distinct from any requirements for fluency that are associated with the job. Language codes must be in BCP-47 format, such as "en-US" or "sr-Latn". For more information, see [Tags for Identifying Languages](https://tools.ietf.org/html/bcp47){: class="external" target="_blank" }. If this field is unspecified and Job.description is present, detected language code based on Job.description is assigned, otherwise defaults to 'en_US'.
     */
    languageCode?: pulumi.Input<string>;
    /**
     * Required during job update. The resource name for the job. This is generated by the service when a job is created. The format is "projects/{project_id}/tenants/{tenant_id}/jobs/{job_id}". For example, "projects/foo/tenants/bar/jobs/baz". Use of this field in job queries and API calls is preferred over the use of requisition_id since this value is unique.
     */
    name?: pulumi.Input<string>;
    /**
     * Strongly recommended for the best service experience. The expiration timestamp of the job. After this timestamp, the job is marked as expired, and it no longer appears in search results. The expired job can't be listed by the ListJobs API, but it can be retrieved with the GetJob API or updated with the UpdateJob API or deleted with the DeleteJob API. An expired job can be updated and opened again by using a future expiration timestamp. Updating an expired job fails if there is another existing open job with same company, language_code and requisition_id. The expired jobs are retained in our system for 90 days. However, the overall expired job count cannot exceed 3 times the maximum number of open jobs over previous 7 days. If this threshold is exceeded, expired jobs are cleaned out in order of earliest expire time. Expired jobs are no longer accessible after they are cleaned out. Invalid timestamps are ignored, and treated as expire time not provided. If the timestamp is before the instant request is made, the job is treated as expired immediately on creation. This kind of job can not be updated. And when creating a job with past timestamp, the posting_publish_time must be set before posting_expire_time. The purpose of this feature is to allow other objects, such as Application, to refer a job that didn't exist in the system prior to becoming expired. If you want to modify a job that was expired on creation, delete it and create a new one. If this value isn't provided at the time of job creation or is invalid, the job posting expires after 30 days from the job's creation time. For example, if the job was created on 2017/01/01 13:00AM UTC with an unspecified expiration date, the job expires after 2017/01/31 13:00AM UTC. If this value isn't provided on job update, it depends on the field masks set by UpdateJobRequest.update_mask. If the field masks include job_end_time, or the masks are empty meaning that every field is updated, the job posting expires after 30 days from the job's last update time. Otherwise the expiration date isn't updated.
     */
    postingExpireTime?: pulumi.Input<string>;
    /**
     * The timestamp this job posting was most recently published. The default value is the time the request arrives at the server. Invalid timestamps are ignored.
     */
    postingPublishTime?: pulumi.Input<string>;
    /**
     * The job PostingRegion (for example, state, country) throughout which the job is available. If this field is set, a LocationFilter in a search query within the job region finds this job posting if an exact location match isn't specified. If this field is set to PostingRegion.NATION or PostingRegion.ADMINISTRATIVE_AREA, setting job Job.addresses to the same location level as this field is strongly recommended.
     */
    postingRegion?: pulumi.Input<enums.jobs.v4.JobPostingRegion>;
    /**
     * Options for job processing.
     */
    processingOptions?: pulumi.Input<inputs.jobs.v4.ProcessingOptionsArgs>;
    project?: pulumi.Input<string>;
    /**
     * A promotion value of the job, as determined by the client. The value determines the sort order of the jobs returned when searching for jobs using the featured jobs search call, with higher promotional values being returned first and ties being resolved by relevance sort. Only the jobs with a promotionValue >0 are returned in a FEATURED_JOB_SEARCH. Default value is 0, and negative values are treated as 0.
     */
    promotionValue?: pulumi.Input<number>;
    /**
     * A description of the qualifications required to perform the job. The use of this field is recommended as an alternative to using the more general description field. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 10,000.
     */
    qualifications?: pulumi.Input<string>;
    /**
     * The requisition ID, also referred to as the posting ID, is assigned by the client to identify a job. This field is intended to be used by clients for client identification and tracking of postings. A job isn't allowed to be created if there is another job with the same company, language_code and requisition_id. The maximum number of allowed characters is 255.
     */
    requisitionId: pulumi.Input<string>;
    /**
     * A description of job responsibilities. The use of this field is recommended as an alternative to using the more general description field. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 10,000.
     */
    responsibilities?: pulumi.Input<string>;
    tenantId: pulumi.Input<string>;
    /**
     * The title of the job, such as "Software Engineer" The maximum number of allowed characters is 500.
     */
    title: pulumi.Input<string>;
    /**
     * Deprecated. The job is only visible to the owner. The visibility of the job. Defaults to Visibility.ACCOUNT_ONLY if not specified.
     *
     * @deprecated Deprecated. The job is only visible to the owner. The visibility of the job. Defaults to Visibility.ACCOUNT_ONLY if not specified.
     */
    visibility?: pulumi.Input<enums.jobs.v4.JobVisibility>;
}
