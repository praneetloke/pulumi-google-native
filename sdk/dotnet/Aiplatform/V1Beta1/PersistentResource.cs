// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1Beta1
{
    /// <summary>
    /// Creates a PersistentResource.
    /// </summary>
    [GoogleNativeResourceType("google-native:aiplatform/v1beta1:PersistentResource")]
    public partial class PersistentResource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Time when the PersistentResource was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. The display name of the PersistentResource. The name can be up to 128 characters long and can consist of any UTF-8 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Optional. Customer-managed encryption key spec for a PersistentResource. If set, this PersistentResource and all sub-resources of this PersistentResource will be secured by this key.
        /// </summary>
        [Output("encryptionSpec")]
        public Output<Outputs.GoogleCloudAiplatformV1beta1EncryptionSpecResponse> EncryptionSpec { get; private set; } = null!;

        /// <summary>
        /// Only populated when persistent resource's state is `STOPPING` or `ERROR`.
        /// </summary>
        [Output("error")]
        public Output<Outputs.GoogleRpcStatusResponse> Error { get; private set; } = null!;

        /// <summary>
        /// Optional. The labels with user-defined metadata to organize PersistentResource. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Immutable. Resource name of a PersistentResource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. The full name of the Compute Engine [network](/compute/docs/networks-and-firewalls#networks) to peered with Vertex AI to host the persistent resources. For example, `projects/12345/global/networks/myVPC`. [Format](/compute/docs/reference/rest/v1/networks/insert) is of the form `projects/{project}/global/networks/{network}`. Where {project} is a project number, as in `12345`, and {network} is a network name. To specify this field, you must have already [configured VPC Network Peering for Vertex AI](https://cloud.google.com/vertex-ai/docs/general/vpc-peering). If this field is left unspecified, the resources aren't peered with any network.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// Required. The ID to use for the PersistentResource, which become the final component of the PersistentResource's resource name. The maximum length is 63 characters, and valid characters are `/^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$/`.
        /// </summary>
        [Output("persistentResourceId")]
        public Output<string> PersistentResourceId { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. A list of names for the reserved IP ranges under the VPC network that can be used for this persistent resource. If set, we will deploy the persistent resource within the provided IP ranges. Otherwise, the persistent resource is deployed to any IP ranges under the provided VPC network. Example: ['vertex-ai-ip-range'].
        /// </summary>
        [Output("reservedIpRanges")]
        public Output<ImmutableArray<string>> ReservedIpRanges { get; private set; } = null!;

        /// <summary>
        /// The spec of the pools of different resources.
        /// </summary>
        [Output("resourcePools")]
        public Output<ImmutableArray<Outputs.GoogleCloudAiplatformV1beta1ResourcePoolResponse>> ResourcePools { get; private set; } = null!;

        /// <summary>
        /// Runtime information of the Persistent Resource.
        /// </summary>
        [Output("resourceRuntime")]
        public Output<Outputs.GoogleCloudAiplatformV1beta1ResourceRuntimeResponse> ResourceRuntime { get; private set; } = null!;

        /// <summary>
        /// Optional. Persistent Resource runtime spec. For example, used for Ray cluster configuration.
        /// </summary>
        [Output("resourceRuntimeSpec")]
        public Output<Outputs.GoogleCloudAiplatformV1beta1ResourceRuntimeSpecResponse> ResourceRuntimeSpec { get; private set; } = null!;

        /// <summary>
        /// Time when the PersistentResource for the first time entered the `RUNNING` state.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// The detailed state of a Study.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Time when the PersistentResource was most recently updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a PersistentResource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PersistentResource(string name, PersistentResourceArgs args, CustomResourceOptions? options = null)
            : base("google-native:aiplatform/v1beta1:PersistentResource", name, args ?? new PersistentResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PersistentResource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:aiplatform/v1beta1:PersistentResource", name, null, MakeResourceOptions(options, id))
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
                    "persistentResourceId",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PersistentResource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PersistentResource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PersistentResource(name, id, options);
        }
    }

    public sealed class PersistentResourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The display name of the PersistentResource. The name can be up to 128 characters long and can consist of any UTF-8 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Optional. Customer-managed encryption key spec for a PersistentResource. If set, this PersistentResource and all sub-resources of this PersistentResource will be secured by this key.
        /// </summary>
        [Input("encryptionSpec")]
        public Input<Inputs.GoogleCloudAiplatformV1beta1EncryptionSpecArgs>? EncryptionSpec { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. The labels with user-defined metadata to organize PersistentResource. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Immutable. Resource name of a PersistentResource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Optional. The full name of the Compute Engine [network](/compute/docs/networks-and-firewalls#networks) to peered with Vertex AI to host the persistent resources. For example, `projects/12345/global/networks/myVPC`. [Format](/compute/docs/reference/rest/v1/networks/insert) is of the form `projects/{project}/global/networks/{network}`. Where {project} is a project number, as in `12345`, and {network} is a network name. To specify this field, you must have already [configured VPC Network Peering for Vertex AI](https://cloud.google.com/vertex-ai/docs/general/vpc-peering). If this field is left unspecified, the resources aren't peered with any network.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// Required. The ID to use for the PersistentResource, which become the final component of the PersistentResource's resource name. The maximum length is 63 characters, and valid characters are `/^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$/`.
        /// </summary>
        [Input("persistentResourceId", required: true)]
        public Input<string> PersistentResourceId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("reservedIpRanges")]
        private InputList<string>? _reservedIpRanges;

        /// <summary>
        /// Optional. A list of names for the reserved IP ranges under the VPC network that can be used for this persistent resource. If set, we will deploy the persistent resource within the provided IP ranges. Otherwise, the persistent resource is deployed to any IP ranges under the provided VPC network. Example: ['vertex-ai-ip-range'].
        /// </summary>
        public InputList<string> ReservedIpRanges
        {
            get => _reservedIpRanges ?? (_reservedIpRanges = new InputList<string>());
            set => _reservedIpRanges = value;
        }

        [Input("resourcePools", required: true)]
        private InputList<Inputs.GoogleCloudAiplatformV1beta1ResourcePoolArgs>? _resourcePools;

        /// <summary>
        /// The spec of the pools of different resources.
        /// </summary>
        public InputList<Inputs.GoogleCloudAiplatformV1beta1ResourcePoolArgs> ResourcePools
        {
            get => _resourcePools ?? (_resourcePools = new InputList<Inputs.GoogleCloudAiplatformV1beta1ResourcePoolArgs>());
            set => _resourcePools = value;
        }

        /// <summary>
        /// Optional. Persistent Resource runtime spec. For example, used for Ray cluster configuration.
        /// </summary>
        [Input("resourceRuntimeSpec")]
        public Input<Inputs.GoogleCloudAiplatformV1beta1ResourceRuntimeSpecArgs>? ResourceRuntimeSpec { get; set; }

        public PersistentResourceArgs()
        {
        }
        public static new PersistentResourceArgs Empty => new PersistentResourceArgs();
    }
}