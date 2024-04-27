// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Backupdr.V1
{
    /// <summary>
    /// Creates a new ManagementServer in a given project and location.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:backupdr/v1:ManagementServer")]
    public partial class ManagementServer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The time when the instance was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. The description of the ManagementServer instance (2048 characters or less).
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Optional. Server specified ETag for the ManagementServer resource to prevent simultaneous updates from overwiting each other.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Optional. Resource labels to represent user provided metadata. Labels currently defined: 1. migrate_from_go= If set to true, the MS is created in migration ready mode.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Required. The name of the management server to create. The name must be unique for the specified project and location.
        /// </summary>
        [Output("managementServerId")]
        public Output<string> ManagementServerId { get; private set; } = null!;

        /// <summary>
        /// The hostname or ip address of the exposed AGM endpoints, used by clients to connect to AGM/RD graphical user interface and APIs.
        /// </summary>
        [Output("managementUri")]
        public Output<Outputs.ManagementURIResponse> ManagementUri { get; private set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// VPC networks to which the ManagementServer instance is connected. For this version, only a single network is supported.
        /// </summary>
        [Output("networks")]
        public Output<ImmutableArray<Outputs.NetworkConfigResponse>> Networks { get; private set; } = null!;

        /// <summary>
        /// The OAuth 2.0 client id is required to make API calls to the BackupDR instance API of this ManagementServer. This is the value that should be provided in the ‘aud’ field of the OIDC ID Token (see openid specification https://openid.net/specs/openid-connect-core-1_0.html#IDToken).
        /// </summary>
        [Output("oauth2ClientId")]
        public Output<string> Oauth2ClientId { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// The ManagementServer state.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The type of the ManagementServer resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The time when the instance was updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// The hostnames of the exposed AGM endpoints for both types of user i.e. 1p and 3p, used to connect AGM/RM UI.
        /// </summary>
        [Output("workforceIdentityBasedManagementUri")]
        public Output<Outputs.WorkforceIdentityBasedManagementURIResponse> WorkforceIdentityBasedManagementUri { get; private set; } = null!;

        /// <summary>
        /// The OAuth client IDs for both types of user i.e. 1p and 3p.
        /// </summary>
        [Output("workforceIdentityBasedOauth2ClientId")]
        public Output<Outputs.WorkforceIdentityBasedOAuth2ClientIDResponse> WorkforceIdentityBasedOauth2ClientId { get; private set; } = null!;


        /// <summary>
        /// Create a ManagementServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagementServer(string name, ManagementServerArgs args, CustomResourceOptions? options = null)
            : base("google-native:backupdr/v1:ManagementServer", name, args ?? new ManagementServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagementServer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:backupdr/v1:ManagementServer", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "location",
                    "managementServerId",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagementServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagementServer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagementServer(name, id, options);
        }
    }

    public sealed class ManagementServerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The description of the ManagementServer instance (2048 characters or less).
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. Server specified ETag for the ManagementServer resource to prevent simultaneous updates from overwiting each other.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Resource labels to represent user provided metadata. Labels currently defined: 1. migrate_from_go= If set to true, the MS is created in migration ready mode.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Required. The name of the management server to create. The name must be unique for the specified project and location.
        /// </summary>
        [Input("managementServerId", required: true)]
        public Input<string> ManagementServerId { get; set; } = null!;

        [Input("networks", required: true)]
        private InputList<Inputs.NetworkConfigArgs>? _networks;

        /// <summary>
        /// VPC networks to which the ManagementServer instance is connected. For this version, only a single network is supported.
        /// </summary>
        public InputList<Inputs.NetworkConfigArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.NetworkConfigArgs>());
            set => _networks = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// The type of the ManagementServer resource.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.Backupdr.V1.ManagementServerType> Type { get; set; } = null!;

        public ManagementServerArgs()
        {
        }
        public static new ManagementServerArgs Empty => new ManagementServerArgs();
    }
}