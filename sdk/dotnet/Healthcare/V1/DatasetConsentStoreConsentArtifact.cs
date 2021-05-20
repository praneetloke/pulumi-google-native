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
    /// Creates a new Consent artifact in the parent consent store.
    /// </summary>
    [GoogleNativeResourceType("google-native:healthcare/v1:DatasetConsentStoreConsentArtifact")]
    public partial class DatasetConsentStoreConsentArtifact : Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. Screenshots, PDFs, or other binary information documenting the user's consent.
        /// </summary>
        [Output("consentContentScreenshots")]
        public Output<ImmutableArray<Outputs.ImageResponse>> ConsentContentScreenshots { get; private set; } = null!;

        /// <summary>
        /// Optional. An string indicating the version of the consent information shown to the user.
        /// </summary>
        [Output("consentContentVersion")]
        public Output<string> ConsentContentVersion { get; private set; } = null!;

        /// <summary>
        /// Optional. A signature from a guardian.
        /// </summary>
        [Output("guardianSignature")]
        public Output<Outputs.SignatureResponse> GuardianSignature { get; private set; } = null!;

        /// <summary>
        /// Optional. Metadata associated with the Consent artifact. For example, the consent locale or user agent version.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>> Metadata { get; private set; } = null!;

        /// <summary>
        /// Resource name of the Consent artifact, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`. Cannot be changed after creation.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Required. User's UUID provided by the client.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;

        /// <summary>
        /// Optional. User's signature.
        /// </summary>
        [Output("userSignature")]
        public Output<Outputs.SignatureResponse> UserSignature { get; private set; } = null!;

        /// <summary>
        /// Optional. A signature from a witness.
        /// </summary>
        [Output("witnessSignature")]
        public Output<Outputs.SignatureResponse> WitnessSignature { get; private set; } = null!;


        /// <summary>
        /// Create a DatasetConsentStoreConsentArtifact resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatasetConsentStoreConsentArtifact(string name, DatasetConsentStoreConsentArtifactArgs args, CustomResourceOptions? options = null)
            : base("google-native:healthcare/v1:DatasetConsentStoreConsentArtifact", name, args ?? new DatasetConsentStoreConsentArtifactArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatasetConsentStoreConsentArtifact(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:healthcare/v1:DatasetConsentStoreConsentArtifact", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DatasetConsentStoreConsentArtifact resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatasetConsentStoreConsentArtifact Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DatasetConsentStoreConsentArtifact(name, id, options);
        }
    }

    public sealed class DatasetConsentStoreConsentArtifactArgs : Pulumi.ResourceArgs
    {
        [Input("consentArtifactId", required: true)]
        public Input<string> ConsentArtifactId { get; set; } = null!;

        [Input("consentContentScreenshots")]
        private InputList<Inputs.ImageArgs>? _consentContentScreenshots;

        /// <summary>
        /// Optional. Screenshots, PDFs, or other binary information documenting the user's consent.
        /// </summary>
        public InputList<Inputs.ImageArgs> ConsentContentScreenshots
        {
            get => _consentContentScreenshots ?? (_consentContentScreenshots = new InputList<Inputs.ImageArgs>());
            set => _consentContentScreenshots = value;
        }

        /// <summary>
        /// Optional. An string indicating the version of the consent information shown to the user.
        /// </summary>
        [Input("consentContentVersion")]
        public Input<string>? ConsentContentVersion { get; set; }

        [Input("consentStoreId", required: true)]
        public Input<string> ConsentStoreId { get; set; } = null!;

        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        /// <summary>
        /// Optional. A signature from a guardian.
        /// </summary>
        [Input("guardianSignature")]
        public Input<Inputs.SignatureArgs>? GuardianSignature { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Optional. Metadata associated with the Consent artifact. For example, the consent locale or user agent version.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// Resource name of the Consent artifact, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`. Cannot be changed after creation.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Required. User's UUID provided by the client.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        /// <summary>
        /// Optional. User's signature.
        /// </summary>
        [Input("userSignature")]
        public Input<Inputs.SignatureArgs>? UserSignature { get; set; }

        /// <summary>
        /// Optional. A signature from a witness.
        /// </summary>
        [Input("witnessSignature")]
        public Input<Inputs.SignatureArgs>? WitnessSignature { get; set; }

        public DatasetConsentStoreConsentArtifactArgs()
        {
        }
    }
}
