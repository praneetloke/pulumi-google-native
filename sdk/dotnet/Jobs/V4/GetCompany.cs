// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Jobs.V4
{
    public static class GetCompany
    {
        /// <summary>
        /// Retrieves specified company.
        /// </summary>
        public static Task<GetCompanyResult> InvokeAsync(GetCompanyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCompanyResult>("google-native:jobs/v4:getCompany", args ?? new GetCompanyArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves specified company.
        /// </summary>
        public static Output<GetCompanyResult> Invoke(GetCompanyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCompanyResult>("google-native:jobs/v4:getCompany", args ?? new GetCompanyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCompanyArgs : global::Pulumi.InvokeArgs
    {
        [Input("companyId", required: true)]
        public string CompanyId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("tenantId", required: true)]
        public string TenantId { get; set; } = null!;

        public GetCompanyArgs()
        {
        }
        public static new GetCompanyArgs Empty => new GetCompanyArgs();
    }

    public sealed class GetCompanyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("companyId", required: true)]
        public Input<string> CompanyId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public GetCompanyInvokeArgs()
        {
        }
        public static new GetCompanyInvokeArgs Empty => new GetCompanyInvokeArgs();
    }


    [OutputType]
    public sealed class GetCompanyResult
    {
        /// <summary>
        /// The URI to employer's career site or careers page on the employer's web site, for example, "https://careers.google.com".
        /// </summary>
        public readonly string CareerSiteUri;
        /// <summary>
        /// Derived details about the company.
        /// </summary>
        public readonly Outputs.CompanyDerivedInfoResponse DerivedInfo;
        /// <summary>
        /// The display name of the company, for example, "Google LLC".
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Equal Employment Opportunity legal disclaimer text to be associated with all jobs, and typically to be displayed in all roles. The maximum number of allowed characters is 500.
        /// </summary>
        public readonly string EeoText;
        /// <summary>
        /// Client side company identifier, used to uniquely identify the company. The maximum number of allowed characters is 255.
        /// </summary>
        public readonly string ExternalId;
        /// <summary>
        /// The street address of the company's main headquarters, which may be different from the job location. The service attempts to geolocate the provided address, and populates a more specific location wherever possible in DerivedInfo.headquarters_location.
        /// </summary>
        public readonly string HeadquartersAddress;
        /// <summary>
        /// Set to true if it is the hiring agency that post jobs for other employers. Defaults to false if not provided.
        /// </summary>
        public readonly bool HiringAgency;
        /// <summary>
        /// A URI that hosts the employer's company logo.
        /// </summary>
        public readonly string ImageUri;
        /// <summary>
        /// This field is deprecated. Please set the searchability of the custom attribute in the Job.custom_attributes going forward. A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword searches. Jobs with `string_values` under these specified field keys are returned if any of the values match the search keyword. Custom field values with parenthesis, brackets and special symbols are not searchable as-is, and those keyword queries must be surrounded by quotes.
        /// </summary>
        public readonly ImmutableArray<string> KeywordSearchableJobCustomAttributes;
        /// <summary>
        /// Required during company update. The resource name for a company. This is generated by the service when a company is created. The format is "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for example, "projects/foo/tenants/bar/companies/baz".
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The employer's company size.
        /// </summary>
        public readonly string Size;
        /// <summary>
        /// Indicates whether a company is flagged to be suspended from public availability by the service when job content appears suspicious, abusive, or spammy.
        /// </summary>
        public readonly bool Suspended;
        /// <summary>
        /// The URI representing the company's primary web site or home page, for example, "https://www.google.com". The maximum number of allowed characters is 255.
        /// </summary>
        public readonly string WebsiteUri;

        [OutputConstructor]
        private GetCompanyResult(
            string careerSiteUri,

            Outputs.CompanyDerivedInfoResponse derivedInfo,

            string displayName,

            string eeoText,

            string externalId,

            string headquartersAddress,

            bool hiringAgency,

            string imageUri,

            ImmutableArray<string> keywordSearchableJobCustomAttributes,

            string name,

            string size,

            bool suspended,

            string websiteUri)
        {
            CareerSiteUri = careerSiteUri;
            DerivedInfo = derivedInfo;
            DisplayName = displayName;
            EeoText = eeoText;
            ExternalId = externalId;
            HeadquartersAddress = headquartersAddress;
            HiringAgency = hiringAgency;
            ImageUri = imageUri;
            KeywordSearchableJobCustomAttributes = keywordSearchableJobCustomAttributes;
            Name = name;
            Size = size;
            Suspended = suspended;
            WebsiteUri = websiteUri;
        }
    }
}
