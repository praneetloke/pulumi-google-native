// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1
{
    /// <summary>
    /// Creates a new User data mapping in the parent consent store.
    /// </summary>
    [GoogleNativeResourceType("google-native:healthcare/v1beta1:UserDataMapping")]
    public partial class UserDataMapping : Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates the time when this mapping was archived.
        /// </summary>
        [Output("archiveTime")]
        public Output<string> ArchiveTime { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this mapping is archived.
        /// </summary>
        [Output("archived")]
        public Output<bool> Archived { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for the mapped resource.
        /// </summary>
        [Output("dataId")]
        public Output<string> DataId { get; private set; } = null!;

        /// <summary>
        /// Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
        /// </summary>
        [Output("resourceAttributes")]
        public Output<ImmutableArray<Outputs.AttributeResponse>> ResourceAttributes { get; private set; } = null!;

        /// <summary>
        /// User's UUID provided by the client.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a UserDataMapping resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserDataMapping(string name, UserDataMappingArgs args, CustomResourceOptions? options = null)
            : base("google-native:healthcare/v1beta1:UserDataMapping", name, args ?? new UserDataMappingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserDataMapping(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:healthcare/v1beta1:UserDataMapping", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing UserDataMapping resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserDataMapping Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new UserDataMapping(name, id, options);
        }
    }

    public sealed class UserDataMappingArgs : Pulumi.ResourceArgs
    {
        [Input("consentStoreId", required: true)]
        public Input<string> ConsentStoreId { get; set; } = null!;

        /// <summary>
        /// A unique identifier for the mapped resource.
        /// </summary>
        [Input("dataId", required: true)]
        public Input<string> DataId { get; set; } = null!;

        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("resourceAttributes")]
        private InputList<Inputs.AttributeArgs>? _resourceAttributes;

        /// <summary>
        /// Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
        /// </summary>
        public InputList<Inputs.AttributeArgs> ResourceAttributes
        {
            get => _resourceAttributes ?? (_resourceAttributes = new InputList<Inputs.AttributeArgs>());
            set => _resourceAttributes = value;
        }

        /// <summary>
        /// User's UUID provided by the client.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public UserDataMappingArgs()
        {
        }
    }
}
