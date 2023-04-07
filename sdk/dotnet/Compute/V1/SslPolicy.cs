// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    /// <summary>
    /// Returns the specified SSL policy resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/v1:SslPolicy")]
    public partial class SslPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// A list of features enabled when the selected profile is CUSTOM. The method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
        /// </summary>
        [Output("customFeatures")]
        public Output<ImmutableArray<string>> CustomFeatures { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The list of features enabled in the SSL policy.
        /// </summary>
        [Output("enabledFeatures")]
        public Output<ImmutableArray<string>> EnabledFeatures { get; private set; } = null!;

        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a SslPolicy. An up-to-date fingerprint must be provided in order to update the SslPolicy, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an SslPolicy.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// [Output only] Type of the resource. Always compute#sslPolicyfor SSL policies.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
        /// </summary>
        [Output("minTlsVersion")]
        public Output<string> MinTlsVersion { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
        /// </summary>
        [Output("profile")]
        public Output<string> Profile { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// URL of the region where the regional SSL policy resides. This field is not applicable to global SSL policies.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// If potential misconfigurations are detected for this SSL policy, this field will be populated with warning messages.
        /// </summary>
        [Output("warnings")]
        public Output<ImmutableArray<Outputs.SslPolicyWarningsItemResponse>> Warnings { get; private set; } = null!;


        /// <summary>
        /// Create a SslPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SslPolicy(string name, SslPolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:SslPolicy", name, args ?? new SslPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SslPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:SslPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SslPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SslPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SslPolicy(name, id, options);
        }
    }

    public sealed class SslPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("customFeatures")]
        private InputList<string>? _customFeatures;

        /// <summary>
        /// A list of features enabled when the selected profile is CUSTOM. The method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
        /// </summary>
        public InputList<string> CustomFeatures
        {
            get => _customFeatures ?? (_customFeatures = new InputList<string>());
            set => _customFeatures = value;
        }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
        /// </summary>
        [Input("minTlsVersion")]
        public Input<Pulumi.GoogleNative.Compute.V1.SslPolicyMinTlsVersion>? MinTlsVersion { get; set; }

        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
        /// </summary>
        [Input("profile")]
        public Input<Pulumi.GoogleNative.Compute.V1.SslPolicyProfile>? Profile { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        public SslPolicyArgs()
        {
        }
        public static new SslPolicyArgs Empty => new SslPolicyArgs();
    }
}
