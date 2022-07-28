// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1
{
    /// <summary>
    /// Creates a new Consent in the parent consent store.
    /// </summary>
    [GoogleNativeResourceType("google-native:healthcare/v1:Consent")]
    public partial class Consent : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The resource name of the Consent artifact that contains proof of the end user's consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`.
        /// </summary>
        [Output("consentArtifact")]
        public Output<string> ConsentArtifact { get; private set; } = null!;

        [Output("consentStoreId")]
        public Output<string> ConsentStoreId { get; private set; } = null!;

        [Output("datasetId")]
        public Output<string> DatasetId { get; private set; } = null!;

        /// <summary>
        /// Timestamp in UTC of when this Consent is considered expired.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Optional. User-supplied key-value pairs used to organize Consent resources. Metadata keys must: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - begin with a letter - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes Metadata values must be: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes No more than 64 metadata entries can be associated with a given consent.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>> Metadata { get; private set; } = null!;

        /// <summary>
        /// Resource name of the Consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consents/{consent_id}`. Cannot be changed after creation.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. Represents a user's consent in terms of the resources that can be accessed and under what conditions.
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<Outputs.GoogleCloudHealthcareV1ConsentPolicyResponse>> Policies { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The timestamp that the revision was created.
        /// </summary>
        [Output("revisionCreateTime")]
        public Output<string> RevisionCreateTime { get; private set; } = null!;

        /// <summary>
        /// The revision ID of the Consent. The format is an 8-character hexadecimal string. Refer to a specific revision of a Consent by appending `@{revision_id}` to the Consent's resource name.
        /// </summary>
        [Output("revisionId")]
        public Output<string> RevisionId { get; private set; } = null!;

        /// <summary>
        /// Indicates the current state of this Consent.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Input only. The time to live for this Consent from when it is created.
        /// </summary>
        [Output("ttl")]
        public Output<string> Ttl { get; private set; } = null!;

        /// <summary>
        /// User's UUID provided by the client.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a Consent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Consent(string name, ConsentArgs args, CustomResourceOptions? options = null)
            : base("google-native:healthcare/v1:Consent", name, args ?? new ConsentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Consent(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:healthcare/v1:Consent", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "consentStoreId",
                    "datasetId",
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
        /// Get an existing Consent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Consent Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Consent(name, id, options);
        }
    }

    public sealed class ConsentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource name of the Consent artifact that contains proof of the end user's consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`.
        /// </summary>
        [Input("consentArtifact", required: true)]
        public Input<string> ConsentArtifact { get; set; } = null!;

        [Input("consentStoreId", required: true)]
        public Input<string> ConsentStoreId { get; set; } = null!;

        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        /// <summary>
        /// Timestamp in UTC of when this Consent is considered expired.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Optional. User-supplied key-value pairs used to organize Consent resources. Metadata keys must: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - begin with a letter - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes Metadata values must be: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes No more than 64 metadata entries can be associated with a given consent.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// Resource name of the Consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consents/{consent_id}`. Cannot be changed after creation.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policies")]
        private InputList<Inputs.GoogleCloudHealthcareV1ConsentPolicyArgs>? _policies;

        /// <summary>
        /// Optional. Represents a user's consent in terms of the resources that can be accessed and under what conditions.
        /// </summary>
        public InputList<Inputs.GoogleCloudHealthcareV1ConsentPolicyArgs> Policies
        {
            get => _policies ?? (_policies = new InputList<Inputs.GoogleCloudHealthcareV1ConsentPolicyArgs>());
            set => _policies = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Indicates the current state of this Consent.
        /// </summary>
        [Input("state", required: true)]
        public Input<Pulumi.GoogleNative.Healthcare.V1.ConsentState> State { get; set; } = null!;

        /// <summary>
        /// Input only. The time to live for this Consent from when it is created.
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// User's UUID provided by the client.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public ConsentArgs()
        {
        }
        public static new ConsentArgs Empty => new ConsentArgs();
    }
}
