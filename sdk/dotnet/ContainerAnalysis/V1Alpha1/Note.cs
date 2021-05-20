// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1
{
    /// <summary>
    /// Creates a new `Note`.
    /// </summary>
    [GoogleNativeResourceType("google-native:containeranalysis/v1alpha1:Note")]
    public partial class Note : Pulumi.CustomResource
    {
        /// <summary>
        /// A note describing an attestation role.
        /// </summary>
        [Output("attestationAuthority")]
        public Output<Outputs.AttestationAuthorityResponse> AttestationAuthority { get; private set; } = null!;

        /// <summary>
        /// A note describing a base image.
        /// </summary>
        [Output("baseImage")]
        public Output<Outputs.BasisResponse> BaseImage { get; private set; } = null!;

        /// <summary>
        /// Build provenance type for a verifiable build.
        /// </summary>
        [Output("buildType")]
        public Output<Outputs.BuildTypeResponse> BuildType { get; private set; } = null!;

        /// <summary>
        /// The time this note was created. This field can be used as a filter in list requests.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// A note describing something that can be deployed.
        /// </summary>
        [Output("deployable")]
        public Output<Outputs.DeployableResponse> Deployable { get; private set; } = null!;

        /// <summary>
        /// A note describing a provider/analysis type.
        /// </summary>
        [Output("discovery")]
        public Output<Outputs.DiscoveryResponse> Discovery { get; private set; } = null!;

        /// <summary>
        /// Time of expiration for this note, null if note does not expire.
        /// </summary>
        [Output("expirationTime")]
        public Output<string> ExpirationTime { get; private set; } = null!;

        /// <summary>
        /// This explicitly denotes which kind of note is specified. This field can be used as a filter in list requests.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// A detailed description of this `Note`.
        /// </summary>
        [Output("longDescription")]
        public Output<string> LongDescription { get; private set; } = null!;

        /// <summary>
        /// The name of the note in the form "projects/{provider_project_id}/notes/{NOTE_ID}"
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A note describing a package hosted by various package managers.
        /// </summary>
        [Output("package")]
        public Output<Outputs.PackageResponse> Package { get; private set; } = null!;

        /// <summary>
        /// URLs associated with this note
        /// </summary>
        [Output("relatedUrl")]
        public Output<ImmutableArray<Outputs.RelatedUrlResponse>> RelatedUrl { get; private set; } = null!;

        /// <summary>
        /// A one sentence description of this `Note`.
        /// </summary>
        [Output("shortDescription")]
        public Output<string> ShortDescription { get; private set; } = null!;

        /// <summary>
        /// The time this note was last updated. This field can be used as a filter in list requests.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// A note describing an upgrade.
        /// </summary>
        [Output("upgrade")]
        public Output<Outputs.UpgradeNoteResponse> Upgrade { get; private set; } = null!;

        /// <summary>
        /// A package vulnerability type of note.
        /// </summary>
        [Output("vulnerabilityType")]
        public Output<Outputs.VulnerabilityTypeResponse> VulnerabilityType { get; private set; } = null!;


        /// <summary>
        /// Create a Note resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Note(string name, NoteArgs args, CustomResourceOptions? options = null)
            : base("google-native:containeranalysis/v1alpha1:Note", name, args ?? new NoteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Note(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:containeranalysis/v1alpha1:Note", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Note resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Note Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Note(name, id, options);
        }
    }

    public sealed class NoteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A note describing an attestation role.
        /// </summary>
        [Input("attestationAuthority")]
        public Input<Inputs.AttestationAuthorityArgs>? AttestationAuthority { get; set; }

        /// <summary>
        /// A note describing a base image.
        /// </summary>
        [Input("baseImage")]
        public Input<Inputs.BasisArgs>? BaseImage { get; set; }

        /// <summary>
        /// Build provenance type for a verifiable build.
        /// </summary>
        [Input("buildType")]
        public Input<Inputs.BuildTypeArgs>? BuildType { get; set; }

        /// <summary>
        /// The time this note was created. This field can be used as a filter in list requests.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// A note describing something that can be deployed.
        /// </summary>
        [Input("deployable")]
        public Input<Inputs.DeployableArgs>? Deployable { get; set; }

        /// <summary>
        /// A note describing a provider/analysis type.
        /// </summary>
        [Input("discovery")]
        public Input<Inputs.DiscoveryArgs>? Discovery { get; set; }

        /// <summary>
        /// Time of expiration for this note, null if note does not expire.
        /// </summary>
        [Input("expirationTime")]
        public Input<string>? ExpirationTime { get; set; }

        /// <summary>
        /// This explicitly denotes which kind of note is specified. This field can be used as a filter in list requests.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// A detailed description of this `Note`.
        /// </summary>
        [Input("longDescription")]
        public Input<string>? LongDescription { get; set; }

        /// <summary>
        /// The name of the note in the form "projects/{provider_project_id}/notes/{NOTE_ID}"
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("noteId")]
        public Input<string>? NoteId { get; set; }

        /// <summary>
        /// A note describing a package hosted by various package managers.
        /// </summary>
        [Input("package")]
        public Input<Inputs.PackageArgs>? Package { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("relatedUrl")]
        private InputList<Inputs.RelatedUrlArgs>? _relatedUrl;

        /// <summary>
        /// URLs associated with this note
        /// </summary>
        public InputList<Inputs.RelatedUrlArgs> RelatedUrl
        {
            get => _relatedUrl ?? (_relatedUrl = new InputList<Inputs.RelatedUrlArgs>());
            set => _relatedUrl = value;
        }

        /// <summary>
        /// A one sentence description of this `Note`.
        /// </summary>
        [Input("shortDescription")]
        public Input<string>? ShortDescription { get; set; }

        /// <summary>
        /// The time this note was last updated. This field can be used as a filter in list requests.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// A note describing an upgrade.
        /// </summary>
        [Input("upgrade")]
        public Input<Inputs.UpgradeNoteArgs>? Upgrade { get; set; }

        /// <summary>
        /// A package vulnerability type of note.
        /// </summary>
        [Input("vulnerabilityType")]
        public Input<Inputs.VulnerabilityTypeArgs>? VulnerabilityType { get; set; }

        public NoteArgs()
        {
        }
    }
}
