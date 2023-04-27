// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1
{
    /// <summary>
    /// Create a new Certificate in a given Project, Location from a particular CaPool.
    /// Auto-naming is currently not supported for this resource.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:privateca/v1:Certificate")]
    public partial class Certificate : global::Pulumi.CustomResource
    {
        [Output("caPoolId")]
        public Output<string> CaPoolId { get; private set; } = null!;

        /// <summary>
        /// A structured description of the issued X.509 certificate.
        /// </summary>
        [Output("certificateDescription")]
        public Output<Outputs.CertificateDescriptionResponse> CertificateDescription { get; private set; } = null!;

        /// <summary>
        /// Optional. It must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`. This field is required when using a CertificateAuthority in the Enterprise CertificateAuthority.Tier, but is optional and its value is ignored otherwise.
        /// </summary>
        [Output("certificateId")]
        public Output<string?> CertificateId { get; private set; } = null!;

        /// <summary>
        /// Immutable. The resource name for a CertificateTemplate used to issue this certificate, in the format `projects/*/locations/*/certificateTemplates/*`. If this is specified, the caller must have the necessary permission to use this template. If this is omitted, no template will be used. This template must be in the same location as the Certificate.
        /// </summary>
        [Output("certificateTemplate")]
        public Output<string> CertificateTemplate { get; private set; } = null!;

        /// <summary>
        /// Immutable. A description of the certificate and key that does not require X.509 or ASN.1.
        /// </summary>
        [Output("config")]
        public Output<Outputs.CertificateConfigResponse> Config { get; private set; } = null!;

        /// <summary>
        /// The time at which this Certificate was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The resource name of the issuing CertificateAuthority in the format `projects/*/locations/*/caPools/*/certificateAuthorities/*`.
        /// </summary>
        [Output("issuerCertificateAuthority")]
        public Output<string> IssuerCertificateAuthority { get; private set; } = null!;

        /// <summary>
        /// Optional. The resource ID of the CertificateAuthority that should issue the certificate. This optional field will ignore the load-balancing scheme of the Pool and directly issue the certificate from the CA with the specified ID, contained in the same CaPool referenced by `parent`. Per-CA quota rules apply. If left empty, a CertificateAuthority will be chosen from the CaPool by the service. For example, to issue a Certificate from a Certificate Authority with resource name "projects/my-project/locations/us-central1/caPools/my-pool/certificateAuthorities/my-ca", you can set the parent to "projects/my-project/locations/us-central1/caPools/my-pool" and the issuing_certificate_authority_id to "my-ca".
        /// </summary>
        [Output("issuingCertificateAuthorityId")]
        public Output<string?> IssuingCertificateAuthorityId { get; private set; } = null!;

        /// <summary>
        /// Optional. Labels with user-defined metadata.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain.
        /// </summary>
        [Output("lifetime")]
        public Output<string> Lifetime { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name for this Certificate in the format `projects/*/locations/*/caPools/*/certificates/*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The pem-encoded, signed X.509 certificate.
        /// </summary>
        [Output("pemCertificate")]
        public Output<string> PemCertificate { get; private set; } = null!;

        /// <summary>
        /// The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
        /// </summary>
        [Output("pemCertificateChain")]
        public Output<ImmutableArray<string>> PemCertificateChain { get; private set; } = null!;

        /// <summary>
        /// Immutable. A pem-encoded X.509 certificate signing request (CSR).
        /// </summary>
        [Output("pemCsr")]
        public Output<string> PemCsr { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. An ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
        /// </summary>
        [Output("revocationDetails")]
        public Output<Outputs.RevocationDetailsResponse> RevocationDetails { get; private set; } = null!;

        /// <summary>
        /// Immutable. Specifies how the Certificate's identity fields are to be decided. If this is omitted, the `DEFAULT` subject mode will be used.
        /// </summary>
        [Output("subjectMode")]
        public Output<string> SubjectMode { get; private set; } = null!;

        /// <summary>
        /// The time at which this Certificate was updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs args, CustomResourceOptions? options = null)
            : base("google-native:privateca/v1:Certificate", name, args ?? new CertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Certificate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:privateca/v1:Certificate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "caPoolId",
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Certificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Certificate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Certificate(name, id, options);
        }
    }

    public sealed class CertificateArgs : global::Pulumi.ResourceArgs
    {
        [Input("caPoolId", required: true)]
        public Input<string> CaPoolId { get; set; } = null!;

        /// <summary>
        /// Optional. It must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`. This field is required when using a CertificateAuthority in the Enterprise CertificateAuthority.Tier, but is optional and its value is ignored otherwise.
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// Immutable. The resource name for a CertificateTemplate used to issue this certificate, in the format `projects/*/locations/*/certificateTemplates/*`. If this is specified, the caller must have the necessary permission to use this template. If this is omitted, no template will be used. This template must be in the same location as the Certificate.
        /// </summary>
        [Input("certificateTemplate")]
        public Input<string>? CertificateTemplate { get; set; }

        /// <summary>
        /// Immutable. A description of the certificate and key that does not require X.509 or ASN.1.
        /// </summary>
        [Input("config")]
        public Input<Inputs.CertificateConfigArgs>? Config { get; set; }

        /// <summary>
        /// Optional. The resource ID of the CertificateAuthority that should issue the certificate. This optional field will ignore the load-balancing scheme of the Pool and directly issue the certificate from the CA with the specified ID, contained in the same CaPool referenced by `parent`. Per-CA quota rules apply. If left empty, a CertificateAuthority will be chosen from the CaPool by the service. For example, to issue a Certificate from a Certificate Authority with resource name "projects/my-project/locations/us-central1/caPools/my-pool/certificateAuthorities/my-ca", you can set the parent to "projects/my-project/locations/us-central1/caPools/my-pool" and the issuing_certificate_authority_id to "my-ca".
        /// </summary>
        [Input("issuingCertificateAuthorityId")]
        public Input<string>? IssuingCertificateAuthorityId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Labels with user-defined metadata.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain.
        /// </summary>
        [Input("lifetime", required: true)]
        public Input<string> Lifetime { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Immutable. A pem-encoded X.509 certificate signing request (CSR).
        /// </summary>
        [Input("pemCsr")]
        public Input<string>? PemCsr { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. An ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Immutable. Specifies how the Certificate's identity fields are to be decided. If this is omitted, the `DEFAULT` subject mode will be used.
        /// </summary>
        [Input("subjectMode")]
        public Input<Pulumi.GoogleNative.Privateca.V1.CertificateSubjectMode>? SubjectMode { get; set; }

        public CertificateArgs()
        {
        }
        public static new CertificateArgs Empty => new CertificateArgs();
    }
}
