// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CertificateManager.V1
{
    public static class GetCertificate
    {
        /// <summary>
        /// Gets details of a single Certificate.
        /// </summary>
        public static Task<GetCertificateResult> InvokeAsync(GetCertificateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCertificateResult>("google-native:certificatemanager/v1:getCertificate", args ?? new GetCertificateArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Certificate.
        /// </summary>
        public static Output<GetCertificateResult> Invoke(GetCertificateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCertificateResult>("google-native:certificatemanager/v1:getCertificate", args ?? new GetCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCertificateArgs : global::Pulumi.InvokeArgs
    {
        [Input("certificateId", required: true)]
        public string CertificateId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetCertificateArgs()
        {
        }
        public static new GetCertificateArgs Empty => new GetCertificateArgs();
    }

    public sealed class GetCertificateInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetCertificateInvokeArgs()
        {
        }
        public static new GetCertificateInvokeArgs Empty => new GetCertificateInvokeArgs();
    }


    [OutputType]
    public sealed class GetCertificateResult
    {
        /// <summary>
        /// The creation timestamp of a Certificate.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// One or more paragraphs of text description of a certificate.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The expiry timestamp of a Certificate.
        /// </summary>
        public readonly string ExpireTime;
        /// <summary>
        /// Set of labels associated with a Certificate.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// If set, contains configuration and state of a managed certificate.
        /// </summary>
        public readonly Outputs.ManagedCertificateResponse Managed;
        /// <summary>
        /// A user-defined name of the certificate. Certificate names must be unique globally and match pattern `projects/*/locations/*/certificates/*`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The PEM-encoded certificate chain.
        /// </summary>
        public readonly string PemCertificate;
        /// <summary>
        /// The list of Subject Alternative Names of dnsName type defined in the certificate (see RFC 5280 4.2.1.6). Managed certificates that haven't been provisioned yet have this field populated with a value of the managed.domains field.
        /// </summary>
        public readonly ImmutableArray<string> SanDnsnames;
        /// <summary>
        /// Immutable. The scope of the certificate.
        /// </summary>
        public readonly string Scope;
        /// <summary>
        /// If set, defines data of a self-managed certificate.
        /// </summary>
        public readonly Outputs.SelfManagedCertificateResponse SelfManaged;
        /// <summary>
        /// The last update timestamp of a Certificate.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetCertificateResult(
            string createTime,

            string description,

            string expireTime,

            ImmutableDictionary<string, string> labels,

            Outputs.ManagedCertificateResponse managed,

            string name,

            string pemCertificate,

            ImmutableArray<string> sanDnsnames,

            string scope,

            Outputs.SelfManagedCertificateResponse selfManaged,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            ExpireTime = expireTime;
            Labels = labels;
            Managed = managed;
            Name = name;
            PemCertificate = pemCertificate;
            SanDnsnames = sanDnsnames;
            Scope = scope;
            SelfManaged = selfManaged;
            UpdateTime = updateTime;
        }
    }
}
