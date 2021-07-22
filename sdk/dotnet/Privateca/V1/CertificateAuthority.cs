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
    /// Create a new CertificateAuthority in a given Project and Location.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:privateca/v1:CertificateAuthority")]
    public partial class CertificateAuthority : Pulumi.CustomResource
    {
        /// <summary>
        /// URLs for accessing content published by this CA, such as the CA certificate and CRLs.
        /// </summary>
        [Output("accessUrls")]
        public Output<Outputs.AccessUrlsResponse> AccessUrls { get; private set; } = null!;

        /// <summary>
        /// A structured description of this CertificateAuthority's CA certificate and its issuers. Ordered as self-to-root.
        /// </summary>
        [Output("caCertificateDescriptions")]
        public Output<ImmutableArray<Outputs.CertificateDescriptionResponse>> CaCertificateDescriptions { get; private set; } = null!;

        /// <summary>
        /// Immutable. The config used to create a self-signed X.509 certificate or CSR.
        /// </summary>
        [Output("config")]
        public Output<Outputs.CertificateConfigResponse> Config { get; private set; } = null!;

        /// <summary>
        /// The time at which this CertificateAuthority was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The time at which this CertificateAuthority was soft deleted, if it is in the DELETED state.
        /// </summary>
        [Output("deleteTime")]
        public Output<string> DeleteTime { get; private set; } = null!;

        /// <summary>
        /// The time at which this CertificateAuthority will be permanently purged, if it is in the DELETED state.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// Immutable. The name of a Cloud Storage bucket where this CertificateAuthority will publish content, such as the CA certificate and CRLs. This must be a bucket name, without any prefixes (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named `my-bucket`, you would simply specify `my-bucket`. If not specified, a managed bucket will be created.
        /// </summary>
        [Output("gcsBucket")]
        public Output<string> GcsBucket { get; private set; } = null!;

        /// <summary>
        /// Immutable. Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA certificate. Otherwise, it is used to sign a CSR.
        /// </summary>
        [Output("keySpec")]
        public Output<Outputs.KeyVersionSpecResponse> KeySpec { get; private set; } = null!;

        /// <summary>
        /// Optional. Labels with user-defined metadata.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The desired lifetime of the CA certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate.
        /// </summary>
        [Output("lifetime")]
        public Output<string> Lifetime { get; private set; } = null!;

        /// <summary>
        /// The resource name for this CertificateAuthority in the format `projects/*/locations/*/caPools/*/certificateAuthorities/*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// This CertificateAuthority's certificate chain, including the current CertificateAuthority's certificate. Ordered such that the root issuer is the final element (consistent with RFC 5246). For a self-signed CA, this will only list the current CertificateAuthority's certificate.
        /// </summary>
        [Output("pemCaCertificates")]
        public Output<ImmutableArray<string>> PemCaCertificates { get; private set; } = null!;

        /// <summary>
        /// The State for this CertificateAuthority.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Optional. If this is a subordinate CertificateAuthority, this field will be set with the subordinate configuration, which describes its issuers. This may be updated, but this CertificateAuthority must continue to validate.
        /// </summary>
        [Output("subordinateConfig")]
        public Output<Outputs.SubordinateConfigResponse> SubordinateConfig { get; private set; } = null!;

        /// <summary>
        /// The CaPool.Tier of the CaPool that includes this CertificateAuthority.
        /// </summary>
        [Output("tier")]
        public Output<string> Tier { get; private set; } = null!;

        /// <summary>
        /// Immutable. The Type of this CertificateAuthority.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The time at which this CertificateAuthority was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a CertificateAuthority resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertificateAuthority(string name, CertificateAuthorityArgs args, CustomResourceOptions? options = null)
            : base("google-native:privateca/v1:CertificateAuthority", name, args ?? new CertificateAuthorityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertificateAuthority(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:privateca/v1:CertificateAuthority", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CertificateAuthority resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertificateAuthority Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CertificateAuthority(name, id, options);
        }
    }

    public sealed class CertificateAuthorityArgs : Pulumi.ResourceArgs
    {
        [Input("caPoolId", required: true)]
        public Input<string> CaPoolId { get; set; } = null!;

        [Input("certificateAuthorityId", required: true)]
        public Input<string> CertificateAuthorityId { get; set; } = null!;

        /// <summary>
        /// Immutable. The config used to create a self-signed X.509 certificate or CSR.
        /// </summary>
        [Input("config", required: true)]
        public Input<Inputs.CertificateConfigArgs> Config { get; set; } = null!;

        /// <summary>
        /// Immutable. The name of a Cloud Storage bucket where this CertificateAuthority will publish content, such as the CA certificate and CRLs. This must be a bucket name, without any prefixes (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named `my-bucket`, you would simply specify `my-bucket`. If not specified, a managed bucket will be created.
        /// </summary>
        [Input("gcsBucket")]
        public Input<string>? GcsBucket { get; set; }

        /// <summary>
        /// Immutable. Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA certificate. Otherwise, it is used to sign a CSR.
        /// </summary>
        [Input("keySpec", required: true)]
        public Input<Inputs.KeyVersionSpecArgs> KeySpec { get; set; } = null!;

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
        /// The desired lifetime of the CA certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate.
        /// </summary>
        [Input("lifetime", required: true)]
        public Input<string> Lifetime { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Optional. If this is a subordinate CertificateAuthority, this field will be set with the subordinate configuration, which describes its issuers. This may be updated, but this CertificateAuthority must continue to validate.
        /// </summary>
        [Input("subordinateConfig")]
        public Input<Inputs.SubordinateConfigArgs>? SubordinateConfig { get; set; }

        /// <summary>
        /// Immutable. The Type of this CertificateAuthority.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.Privateca.V1.CertificateAuthorityType> Type { get; set; } = null!;

        public CertificateAuthorityArgs()
        {
        }
    }
}
