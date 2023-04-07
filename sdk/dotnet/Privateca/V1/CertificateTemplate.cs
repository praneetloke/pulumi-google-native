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
    /// Create a new CertificateTemplate in a given Project and Location.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:privateca/v1:CertificateTemplate")]
    public partial class CertificateTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Required. It must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
        /// </summary>
        [Output("certificateTemplateId")]
        public Output<string> CertificateTemplateId { get; private set; } = null!;

        /// <summary>
        /// The time at which this CertificateTemplate was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. A human-readable description of scenarios this template is intended for.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
        /// </summary>
        [Output("identityConstraints")]
        public Output<Outputs.CertificateIdentityConstraintsResponse> IdentityConstraints { get; private set; } = null!;

        /// <summary>
        /// Optional. Labels with user-defined metadata.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name for this CertificateTemplate in the format `projects/*/locations/*/certificateTemplates/*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baseline_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
        /// </summary>
        [Output("passthroughExtensions")]
        public Output<Outputs.CertificateExtensionConstraintsResponse> PassthroughExtensions { get; private set; } = null!;

        /// <summary>
        /// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baseline_values for the same properties, the certificate issuance request will fail.
        /// </summary>
        [Output("predefinedValues")]
        public Output<Outputs.X509ParametersResponse> PredefinedValues { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. An ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// The time at which this CertificateTemplate was updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a CertificateTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertificateTemplate(string name, CertificateTemplateArgs args, CustomResourceOptions? options = null)
            : base("google-native:privateca/v1:CertificateTemplate", name, args ?? new CertificateTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertificateTemplate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:privateca/v1:CertificateTemplate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "certificateTemplateId",
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
        /// Get an existing CertificateTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertificateTemplate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CertificateTemplate(name, id, options);
        }
    }

    public sealed class CertificateTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. It must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
        /// </summary>
        [Input("certificateTemplateId", required: true)]
        public Input<string> CertificateTemplateId { get; set; } = null!;

        /// <summary>
        /// Optional. A human-readable description of scenarios this template is intended for.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
        /// </summary>
        [Input("identityConstraints")]
        public Input<Inputs.CertificateIdentityConstraintsArgs>? IdentityConstraints { get; set; }

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

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baseline_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
        /// </summary>
        [Input("passthroughExtensions")]
        public Input<Inputs.CertificateExtensionConstraintsArgs>? PassthroughExtensions { get; set; }

        /// <summary>
        /// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baseline_values for the same properties, the certificate issuance request will fail.
        /// </summary>
        [Input("predefinedValues")]
        public Input<Inputs.X509ParametersArgs>? PredefinedValues { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. An ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        public CertificateTemplateArgs()
        {
        }
        public static new CertificateTemplateArgs Empty => new CertificateTemplateArgs();
    }
}
