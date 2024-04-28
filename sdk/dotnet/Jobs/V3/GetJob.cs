// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Jobs.V3
{
    public static class GetJob
    {
        /// <summary>
        /// Retrieves the specified job, whose status is OPEN or recently EXPIRED within the last 90 days.
        /// </summary>
        public static Task<GetJobResult> InvokeAsync(GetJobArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetJobResult>("google-native:jobs/v3:getJob", args ?? new GetJobArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the specified job, whose status is OPEN or recently EXPIRED within the last 90 days.
        /// </summary>
        public static Output<GetJobResult> Invoke(GetJobInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetJobResult>("google-native:jobs/v3:getJob", args ?? new GetJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetJobArgs : global::Pulumi.InvokeArgs
    {
        [Input("jobId", required: true)]
        public string JobId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetJobArgs()
        {
        }
        public static new GetJobArgs Empty => new GetJobArgs();
    }

    public sealed class GetJobInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("jobId", required: true)]
        public Input<string> JobId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetJobInvokeArgs()
        {
        }
        public static new GetJobInvokeArgs Empty => new GetJobInvokeArgs();
    }


    [OutputType]
    public sealed class GetJobResult
    {
        /// <summary>
        /// Optional but strongly recommended for the best service experience. Location(s) where the employer is looking to hire for this job posting. Specifying the full street address(es) of the hiring location enables better API results, especially job searches by commute time. At most 50 locations are allowed for best search performance. If a job has more locations, it is suggested to split it into multiple jobs with unique requisition_ids (e.g. 'ReqA' becomes 'ReqA-1', 'ReqA-2', etc.) as multiple jobs with the same company_name, language_code and requisition_id are not allowed. If the original requisition_id must be preserved, a custom field should be used for storage. It is also suggested to group the locations that close to each other in the same job for better search experience. Jobs with multiple addresses must have their addresses with the same LocationType to allow location filtering to work properly. (For example, a Job with addresses "1600 Amphitheatre Parkway, Mountain View, CA, USA" and "London, UK" may not have location filters applied correctly at search time since the first is a LocationType.STREET_ADDRESS and the second is a LocationType.LOCALITY.) If a job needs to have multiple addresses, it is suggested to split it into multiple jobs with same LocationTypes. The maximum number of allowed characters is 500.
        /// </summary>
        public readonly ImmutableArray<string> Addresses;
        /// <summary>
        /// At least one field within ApplicationInfo must be specified. Job application information.
        /// </summary>
        public readonly Outputs.ApplicationInfoResponse ApplicationInfo;
        /// <summary>
        /// Display name of the company listing the job.
        /// </summary>
        public readonly string CompanyDisplayName;
        /// <summary>
        /// The resource name of the company listing the job, such as "projects/api-test-project/companies/foo".
        /// </summary>
        public readonly string CompanyName;
        /// <summary>
        /// Optional. Job compensation information.
        /// </summary>
        public readonly Outputs.CompensationInfoResponse CompensationInfo;
        /// <summary>
        /// Optional. A map of fields to hold both filterable and non-filterable custom job attributes that are not covered by the provided structured fields. The keys of the map are strings up to 64 bytes and must match the pattern: a-zA-Z*. For example, key0LikeThis or KEY_1_LIKE_THIS. At most 100 filterable and at most 100 unfilterable keys are supported. For filterable `string_values`, across all keys at most 200 values are allowed, with each string no more than 255 characters. For unfilterable `string_values`, the maximum total size of `string_values` across all keys is 50KB.
        /// </summary>
        public readonly Outputs.CustomAttributeResponse CustomAttributes;
        /// <summary>
        /// Optional. The desired education degrees for the job, such as Bachelors, Masters.
        /// </summary>
        public readonly ImmutableArray<string> DegreeTypes;
        /// <summary>
        /// Optional. The department or functional area within the company with the open position. The maximum number of allowed characters is 255.
        /// </summary>
        public readonly string Department;
        /// <summary>
        /// Derived details about the job posting.
        /// </summary>
        public readonly Outputs.JobDerivedInfoResponse DerivedInfo;
        /// <summary>
        /// The description of the job, which typically includes a multi-paragraph description of the company and related information. Separate fields are provided on the job object for responsibilities, qualifications, and other job characteristics. Use of these separate job fields is recommended. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 100,000.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. The employment type(s) of a job, for example, full time or part time.
        /// </summary>
        public readonly ImmutableArray<string> EmploymentTypes;
        /// <summary>
        /// Optional. A description of bonus, commission, and other compensation incentives associated with the job not including salary or pay. The maximum number of allowed characters is 10,000.
        /// </summary>
        public readonly string Incentives;
        /// <summary>
        /// Optional. The benefits included with the job.
        /// </summary>
        public readonly ImmutableArray<string> JobBenefits;
        /// <summary>
        /// Optional. The end timestamp of the job. Typically this field is used for contracting engagements. Invalid timestamps are ignored.
        /// </summary>
        public readonly string JobEndTime;
        /// <summary>
        /// Optional. The experience level associated with the job, such as "Entry Level".
        /// </summary>
        public readonly string JobLevel;
        /// <summary>
        /// Optional. The start timestamp of the job in UTC time zone. Typically this field is used for contracting engagements. Invalid timestamps are ignored.
        /// </summary>
        public readonly string JobStartTime;
        /// <summary>
        /// Optional. The language of the posting. This field is distinct from any requirements for fluency that are associated with the job. Language codes must be in BCP-47 format, such as "en-US" or "sr-Latn". For more information, see [Tags for Identifying Languages](https://tools.ietf.org/html/bcp47){: class="external" target="_blank" }. If this field is unspecified and Job.description is present, detected language code based on Job.description is assigned, otherwise defaults to 'en_US'.
        /// </summary>
        public readonly string LanguageCode;
        /// <summary>
        /// Required during job update. The resource name for the job. This is generated by the service when a job is created. The format is "projects/{project_id}/jobs/{job_id}", for example, "projects/api-test-project/jobs/1234". Use of this field in job queries and API calls is preferred over the use of requisition_id since this value is unique.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The timestamp when this job posting was created.
        /// </summary>
        public readonly string PostingCreateTime;
        /// <summary>
        /// Optional but strongly recommended for the best service experience. The expiration timestamp of the job. After this timestamp, the job is marked as expired, and it no longer appears in search results. The expired job can't be deleted or listed by the DeleteJob and ListJobs APIs, but it can be retrieved with the GetJob API or updated with the UpdateJob API. An expired job can be updated and opened again by using a future expiration timestamp. Updating an expired job fails if there is another existing open job with same company_name, language_code and requisition_id. The expired jobs are retained in our system for 90 days. However, the overall expired job count cannot exceed 3 times the maximum of open jobs count over the past week, otherwise jobs with earlier expire time are cleaned first. Expired jobs are no longer accessible after they are cleaned out. Invalid timestamps are ignored, and treated as expire time not provided. Timestamp before the instant request is made is considered valid, the job will be treated as expired immediately. If this value is not provided at the time of job creation or is invalid, the job posting expires after 30 days from the job's creation time. For example, if the job was created on 2017/01/01 13:00AM UTC with an unspecified expiration date, the job expires after 2017/01/31 13:00AM UTC. If this value is not provided on job update, it depends on the field masks set by UpdateJobRequest.update_mask. If the field masks include expiry_time, or the masks are empty meaning that every field is updated, the job posting expires after 30 days from the job's last update time. Otherwise the expiration date isn't updated.
        /// </summary>
        public readonly string PostingExpireTime;
        /// <summary>
        /// Optional. The timestamp this job posting was most recently published. The default value is the time the request arrives at the server. Invalid timestamps are ignored.
        /// </summary>
        public readonly string PostingPublishTime;
        /// <summary>
        /// Optional. The job PostingRegion (for example, state, country) throughout which the job is available. If this field is set, a LocationFilter in a search query within the job region finds this job posting if an exact location match isn't specified. If this field is set to PostingRegion.NATION or PostingRegion.ADMINISTRATIVE_AREA, setting job Job.addresses to the same location level as this field is strongly recommended.
        /// </summary>
        public readonly string PostingRegion;
        /// <summary>
        /// The timestamp when this job posting was last updated.
        /// </summary>
        public readonly string PostingUpdateTime;
        /// <summary>
        /// Optional. Options for job processing.
        /// </summary>
        public readonly Outputs.ProcessingOptionsResponse ProcessingOptions;
        /// <summary>
        /// Optional. A promotion value of the job, as determined by the client. The value determines the sort order of the jobs returned when searching for jobs using the featured jobs search call, with higher promotional values being returned first and ties being resolved by relevance sort. Only the jobs with a promotionValue &gt;0 are returned in a FEATURED_JOB_SEARCH. Default value is 0, and negative values are treated as 0.
        /// </summary>
        public readonly int PromotionValue;
        /// <summary>
        /// Optional. A description of the qualifications required to perform the job. The use of this field is recommended as an alternative to using the more general description field. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 10,000.
        /// </summary>
        public readonly string Qualifications;
        /// <summary>
        /// The requisition ID, also referred to as the posting ID, assigned by the client to identify a job. This field is intended to be used by clients for client identification and tracking of postings. A job is not allowed to be created if there is another job with the same [company_name], language_code and requisition_id. The maximum number of allowed characters is 255.
        /// </summary>
        public readonly string RequisitionId;
        /// <summary>
        /// Optional. A description of job responsibilities. The use of this field is recommended as an alternative to using the more general description field. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 10,000.
        /// </summary>
        public readonly string Responsibilities;
        /// <summary>
        /// The title of the job, such as "Software Engineer" The maximum number of allowed characters is 500.
        /// </summary>
        public readonly string Title;
        /// <summary>
        /// Deprecated. The job is only visible to the owner. The visibility of the job. Defaults to Visibility.ACCOUNT_ONLY if not specified.
        /// </summary>
        public readonly string Visibility;

        [OutputConstructor]
        private GetJobResult(
            ImmutableArray<string> addresses,

            Outputs.ApplicationInfoResponse applicationInfo,

            string companyDisplayName,

            string companyName,

            Outputs.CompensationInfoResponse compensationInfo,

            Outputs.CustomAttributeResponse customAttributes,

            ImmutableArray<string> degreeTypes,

            string department,

            Outputs.JobDerivedInfoResponse derivedInfo,

            string description,

            ImmutableArray<string> employmentTypes,

            string incentives,

            ImmutableArray<string> jobBenefits,

            string jobEndTime,

            string jobLevel,

            string jobStartTime,

            string languageCode,

            string name,

            string postingCreateTime,

            string postingExpireTime,

            string postingPublishTime,

            string postingRegion,

            string postingUpdateTime,

            Outputs.ProcessingOptionsResponse processingOptions,

            int promotionValue,

            string qualifications,

            string requisitionId,

            string responsibilities,

            string title,

            string visibility)
        {
            Addresses = addresses;
            ApplicationInfo = applicationInfo;
            CompanyDisplayName = companyDisplayName;
            CompanyName = companyName;
            CompensationInfo = compensationInfo;
            CustomAttributes = customAttributes;
            DegreeTypes = degreeTypes;
            Department = department;
            DerivedInfo = derivedInfo;
            Description = description;
            EmploymentTypes = employmentTypes;
            Incentives = incentives;
            JobBenefits = jobBenefits;
            JobEndTime = jobEndTime;
            JobLevel = jobLevel;
            JobStartTime = jobStartTime;
            LanguageCode = languageCode;
            Name = name;
            PostingCreateTime = postingCreateTime;
            PostingExpireTime = postingExpireTime;
            PostingPublishTime = postingPublishTime;
            PostingRegion = postingRegion;
            PostingUpdateTime = postingUpdateTime;
            ProcessingOptions = processingOptions;
            PromotionValue = promotionValue;
            Qualifications = qualifications;
            RequisitionId = requisitionId;
            Responsibilities = responsibilities;
            Title = title;
            Visibility = visibility;
        }
    }
}
