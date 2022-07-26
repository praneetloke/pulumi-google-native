// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    /// <summary>
    /// Creates an alias from a key/certificate pair. The structure of the request is controlled by the `format` query parameter: - `keycertfile` - Separate PEM-encoded key and certificate files are uploaded. Set `Content-Type: multipart/form-data` and include the `keyFile`, `certFile`, and `password` (if keys are encrypted) fields in the request body. If uploading to a truststore, omit `keyFile`. - `pkcs12` - A PKCS12 file is uploaded. Set `Content-Type: multipart/form-data`, provide the file in the `file` field, and include the `password` field if the file is encrypted in the request body. - `selfsignedcert` - A new private key and certificate are generated. Set `Content-Type: application/json` and include CertificateGenerationSpec in the request body.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:Alias")]
    public partial class Alias : Pulumi.CustomResource
    {
        /// <summary>
        /// Alias for the key/certificate pair. Values must match the regular expression `[\w\s-.]{1,255}`. This must be provided for all formats except `selfsignedcert`; self-signed certs may specify the alias in either this parameter or the JSON body.
        /// </summary>
        [Output("alias")]
        public Output<string> AliasValue { get; private set; } = null!;

        /// <summary>
        /// Chain of certificates under this alias.
        /// </summary>
        [Output("certsInfo")]
        public Output<Outputs.GoogleCloudApigeeV1CertificateResponse> CertsInfo { get; private set; } = null!;

        [Output("environmentId")]
        public Output<string> EnvironmentId { get; private set; } = null!;

        /// <summary>
        /// Required. Format of the data. Valid values include: `selfsignedcert`, `keycertfile`, or `pkcs12`
        /// </summary>
        [Output("format")]
        public Output<string> Format { get; private set; } = null!;

        /// <summary>
        /// Flag that specifies whether to ignore expiry validation. If set to `true`, no expiry validation will be performed.
        /// </summary>
        [Output("ignoreExpiryValidation")]
        public Output<string?> IgnoreExpiryValidation { get; private set; } = null!;

        /// <summary>
        /// Flag that specifies whether to ignore newline validation. If set to `true`, no error is thrown when the file contains a certificate chain with no newline between each certificate. Defaults to `false`.
        /// </summary>
        [Output("ignoreNewlineValidation")]
        public Output<string?> IgnoreNewlineValidation { get; private set; } = null!;

        [Output("keystoreId")]
        public Output<string> KeystoreId { get; private set; } = null!;

        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// DEPRECATED: For improved security, specify the password in the request body instead of using the query parameter. To specify the password in the request body, set `Content-type: multipart/form-data` part with name `password`. Password for the private key file, if required.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Type of alias.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Alias resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Alias(string name, AliasArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:Alias", name, args ?? new AliasArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Alias(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:Alias", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "environmentId",
                    "format",
                    "keystoreId",
                    "organizationId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Alias resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Alias Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Alias(name, id, options);
        }
    }

    public sealed class AliasArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alias for the key/certificate pair. Values must match the regular expression `[\w\s-.]{1,255}`. This must be provided for all formats except `selfsignedcert`; self-signed certs may specify the alias in either this parameter or the JSON body.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// The HTTP Content-Type header value specifying the content type of the body.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// The HTTP request/response body as raw binary.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        [Input("extensions")]
        private InputList<ImmutableDictionary<string, string>>? _extensions;

        /// <summary>
        /// Application specific response metadata. Must be set in the first response for streaming APIs.
        /// </summary>
        public InputList<ImmutableDictionary<string, string>> Extensions
        {
            get => _extensions ?? (_extensions = new InputList<ImmutableDictionary<string, string>>());
            set => _extensions = value;
        }

        /// <summary>
        /// File to upload.
        /// </summary>
        [Input("file")]
        public Input<AssetOrArchive>? File { get; set; }

        /// <summary>
        /// Required. Format of the data. Valid values include: `selfsignedcert`, `keycertfile`, or `pkcs12`
        /// </summary>
        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        /// <summary>
        /// Flag that specifies whether to ignore expiry validation. If set to `true`, no expiry validation will be performed.
        /// </summary>
        [Input("ignoreExpiryValidation")]
        public Input<string>? IgnoreExpiryValidation { get; set; }

        /// <summary>
        /// Flag that specifies whether to ignore newline validation. If set to `true`, no error is thrown when the file contains a certificate chain with no newline between each certificate. Defaults to `false`.
        /// </summary>
        [Input("ignoreNewlineValidation")]
        public Input<string>? IgnoreNewlineValidation { get; set; }

        [Input("keystoreId", required: true)]
        public Input<string> KeystoreId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        /// <summary>
        /// DEPRECATED: For improved security, specify the password in the request body instead of using the query parameter. To specify the password in the request body, set `Content-type: multipart/form-data` part with name `password`. Password for the private key file, if required.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        public AliasArgs()
        {
        }
    }
}
