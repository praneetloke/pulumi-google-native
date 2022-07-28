// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.FirebaseHosting.V1Beta1
{
    public static class GetDomain
    {
        /// <summary>
        /// Gets a domain mapping on the specified site.
        /// </summary>
        public static Task<GetDomainResult> InvokeAsync(GetDomainArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainResult>("google-native:firebasehosting/v1beta1:getDomain", args ?? new GetDomainArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a domain mapping on the specified site.
        /// </summary>
        public static Output<GetDomainResult> Invoke(GetDomainInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDomainResult>("google-native:firebasehosting/v1beta1:getDomain", args ?? new GetDomainInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainArgs : global::Pulumi.InvokeArgs
    {
        [Input("domainId", required: true)]
        public string DomainId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("siteId", required: true)]
        public string SiteId { get; set; } = null!;

        public GetDomainArgs()
        {
        }
        public static new GetDomainArgs Empty => new GetDomainArgs();
    }

    public sealed class GetDomainInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("domainId", required: true)]
        public Input<string> DomainId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("siteId", required: true)]
        public Input<string> SiteId { get; set; } = null!;

        public GetDomainInvokeArgs()
        {
        }
        public static new GetDomainInvokeArgs Empty => new GetDomainInvokeArgs();
    }


    [OutputType]
    public sealed class GetDomainResult
    {
        /// <summary>
        /// The domain name of the association.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// If set, the domain should redirect with the provided parameters.
        /// </summary>
        public readonly Outputs.DomainRedirectResponse DomainRedirect;
        /// <summary>
        /// Information about the provisioning of certificates and the health of the DNS resolution for the domain.
        /// </summary>
        public readonly Outputs.DomainProvisioningResponse Provisioning;
        /// <summary>
        /// The site name of the association.
        /// </summary>
        public readonly string Site;
        /// <summary>
        /// Additional status of the domain association.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The time at which the domain was last updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetDomainResult(
            string domainName,

            Outputs.DomainRedirectResponse domainRedirect,

            Outputs.DomainProvisioningResponse provisioning,

            string site,

            string status,

            string updateTime)
        {
            DomainName = domainName;
            DomainRedirect = domainRedirect;
            Provisioning = provisioning;
            Site = site;
            Status = status;
            UpdateTime = updateTime;
        }
    }
}
