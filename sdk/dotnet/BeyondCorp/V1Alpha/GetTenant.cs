// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BeyondCorp.V1Alpha
{
    public static class GetTenant
    {
        /// <summary>
        /// Gets details of a single Tenant.
        /// </summary>
        public static Task<GetTenantResult> InvokeAsync(GetTenantArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTenantResult>("google-native:beyondcorp/v1alpha:getTenant", args ?? new GetTenantArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Tenant.
        /// </summary>
        public static Output<GetTenantResult> Invoke(GetTenantInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTenantResult>("google-native:beyondcorp/v1alpha:getTenant", args ?? new GetTenantInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTenantArgs : global::Pulumi.InvokeArgs
    {
        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        [Input("tenantId", required: true)]
        public string TenantId { get; set; } = null!;

        public GetTenantArgs()
        {
        }
        public static new GetTenantArgs Empty => new GetTenantArgs();
    }

    public sealed class GetTenantInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public GetTenantInvokeArgs()
        {
        }
        public static new GetTenantInvokeArgs Empty => new GetTenantInvokeArgs();
    }


    [OutputType]
    public sealed class GetTenantResult
    {
        /// <summary>
        /// Timestamp when the resource was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. An arbitrary caller-provided name for the Tenant. Cannot exceed 64 characters.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Optional. Google group email to which the Tenant is enabled.
        /// </summary>
        public readonly string GoogleGroupEmail;
        /// <summary>
        /// Optional. Google group ID to which the Tenant is enabled.
        /// </summary>
        public readonly string GoogleGroupId;
        /// <summary>
        /// Unique resource name of the Tenant. The name is ignored when creating Tenant.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional. Metadata provided by the Partner associated with Tenant.
        /// </summary>
        public readonly Outputs.GoogleCloudBeyondcorpPartnerservicesV1alphaPartnerMetadataResponse PartnerMetadata;
        /// <summary>
        /// Timestamp when the resource was last modified.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetTenantResult(
            string createTime,

            string displayName,

            string googleGroupEmail,

            string googleGroupId,

            string name,

            Outputs.GoogleCloudBeyondcorpPartnerservicesV1alphaPartnerMetadataResponse partnerMetadata,

            string updateTime)
        {
            CreateTime = createTime;
            DisplayName = displayName;
            GoogleGroupEmail = googleGroupEmail;
            GoogleGroupId = googleGroupId;
            Name = name;
            PartnerMetadata = partnerMetadata;
            UpdateTime = updateTime;
        }
    }
}
