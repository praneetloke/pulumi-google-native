// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Jobs.V4
{
    public static class GetTenant
    {
        /// <summary>
        /// Retrieves specified tenant.
        /// </summary>
        public static Task<GetTenantResult> InvokeAsync(GetTenantArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTenantResult>("google-native:jobs/v4:getTenant", args ?? new GetTenantArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves specified tenant.
        /// </summary>
        public static Output<GetTenantResult> Invoke(GetTenantInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTenantResult>("google-native:jobs/v4:getTenant", args ?? new GetTenantInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTenantArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("tenantId", required: true)]
        public string TenantId { get; set; } = null!;

        public GetTenantArgs()
        {
        }
        public static new GetTenantArgs Empty => new GetTenantArgs();
    }

    public sealed class GetTenantInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

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
        /// Client side tenant identifier, used to uniquely identify the tenant. The maximum number of allowed characters is 255.
        /// </summary>
        public readonly string ExternalId;
        /// <summary>
        /// Required during tenant update. The resource name for a tenant. This is generated by the service when a tenant is created. The format is "projects/{project_id}/tenants/{tenant_id}", for example, "projects/foo/tenants/bar".
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetTenantResult(
            string externalId,

            string name)
        {
            ExternalId = externalId;
            Name = name;
        }
    }
}
