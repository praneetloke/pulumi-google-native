// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1
{
    /// <summary>
    /// Creates a new note.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:containeranalysis/v1:Note")]
    public partial class Note : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A note describing an attestation role.
        /// </summary>
        [Output("attestation")]
        public Output<Outputs.AttestationNoteResponse> Attestation { get; private set; } = null!;

        /// <summary>
        /// A note describing build provenance for a verifiable build.
        /// </summary>
        [Output("build")]
        public Output<Outputs.BuildNoteResponse> Build { get; private set; } = null!;

        /// <summary>
        /// A note describing a compliance check.
        /// </summary>
        [Output("compliance")]
        public Output<Outputs.ComplianceNoteResponse> Compliance { get; private set; } = null!;

        /// <summary>
        /// The time this note was created. This field can be used as a filter in list requests.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// A note describing something that can be deployed.
        /// </summary>
        [Output("deployment")]
        public Output<Outputs.DeploymentNoteResponse> Deployment { get; private set; } = null!;

        /// <summary>
        /// A note describing the initial analysis of a resource.
        /// </summary>
        [Output("discovery")]
        public Output<Outputs.DiscoveryNoteResponse> Discovery { get; private set; } = null!;

        /// <summary>
        /// A note describing a dsse attestation note.
        /// </summary>
        [Output("dsseAttestation")]
        public Output<Outputs.DSSEAttestationNoteResponse> DsseAttestation { get; private set; } = null!;

        /// <summary>
        /// Time of expiration for this note. Empty if note does not expire.
        /// </summary>
        [Output("expirationTime")]
        public Output<string> ExpirationTime { get; private set; } = null!;

        /// <summary>
        /// A note describing a base image.
        /// </summary>
        [Output("image")]
        public Output<Outputs.ImageNoteResponse> Image { get; private set; } = null!;

        /// <summary>
        /// The type of analysis. This field can be used as a filter in list requests.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// A detailed description of this note.
        /// </summary>
        [Output("longDescription")]
        public Output<string> LongDescription { get; private set; } = null!;

        /// <summary>
        /// The name of the note in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Required. The ID to use for this note.
        /// </summary>
        [Output("noteId")]
        public Output<string> NoteId { get; private set; } = null!;

        /// <summary>
        /// A note describing a package hosted by various package managers.
        /// </summary>
        [Output("package")]
        public Output<Outputs.PackageNoteResponse> Package { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Other notes related to this note.
        /// </summary>
        [Output("relatedNoteNames")]
        public Output<ImmutableArray<string>> RelatedNoteNames { get; private set; } = null!;

        /// <summary>
        /// URLs associated with this note.
        /// </summary>
        [Output("relatedUrl")]
        public Output<ImmutableArray<Outputs.RelatedUrlResponse>> RelatedUrl { get; private set; } = null!;

        /// <summary>
        /// A one sentence description of this note.
        /// </summary>
        [Output("shortDescription")]
        public Output<string> ShortDescription { get; private set; } = null!;

        /// <summary>
        /// The time this note was last updated. This field can be used as a filter in list requests.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// A note describing available package upgrades.
        /// </summary>
        [Output("upgrade")]
        public Output<Outputs.UpgradeNoteResponse> Upgrade { get; private set; } = null!;

        /// <summary>
        /// A note describing a package vulnerability.
        /// </summary>
        [Output("vulnerability")]
        public Output<Outputs.VulnerabilityNoteResponse> Vulnerability { get; private set; } = null!;


        /// <summary>
        /// Create a Note resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Note(string name, NoteArgs args, CustomResourceOptions? options = null)
            : base("google-native:containeranalysis/v1:Note", name, args ?? new NoteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Note(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:containeranalysis/v1:Note", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "noteId",
                    "project",
                },
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

    public sealed class NoteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A note describing an attestation role.
        /// </summary>
        [Input("attestation")]
        public Input<Inputs.AttestationNoteArgs>? Attestation { get; set; }

        /// <summary>
        /// A note describing build provenance for a verifiable build.
        /// </summary>
        [Input("build")]
        public Input<Inputs.BuildNoteArgs>? Build { get; set; }

        /// <summary>
        /// A note describing a compliance check.
        /// </summary>
        [Input("compliance")]
        public Input<Inputs.ComplianceNoteArgs>? Compliance { get; set; }

        /// <summary>
        /// A note describing something that can be deployed.
        /// </summary>
        [Input("deployment")]
        public Input<Inputs.DeploymentNoteArgs>? Deployment { get; set; }

        /// <summary>
        /// A note describing the initial analysis of a resource.
        /// </summary>
        [Input("discovery")]
        public Input<Inputs.DiscoveryNoteArgs>? Discovery { get; set; }

        /// <summary>
        /// A note describing a dsse attestation note.
        /// </summary>
        [Input("dsseAttestation")]
        public Input<Inputs.DSSEAttestationNoteArgs>? DsseAttestation { get; set; }

        /// <summary>
        /// Time of expiration for this note. Empty if note does not expire.
        /// </summary>
        [Input("expirationTime")]
        public Input<string>? ExpirationTime { get; set; }

        /// <summary>
        /// A note describing a base image.
        /// </summary>
        [Input("image")]
        public Input<Inputs.ImageNoteArgs>? Image { get; set; }

        /// <summary>
        /// A detailed description of this note.
        /// </summary>
        [Input("longDescription")]
        public Input<string>? LongDescription { get; set; }

        /// <summary>
        /// Required. The ID to use for this note.
        /// </summary>
        [Input("noteId", required: true)]
        public Input<string> NoteId { get; set; } = null!;

        /// <summary>
        /// A note describing a package hosted by various package managers.
        /// </summary>
        [Input("package")]
        public Input<Inputs.PackageNoteArgs>? Package { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("relatedNoteNames")]
        private InputList<string>? _relatedNoteNames;

        /// <summary>
        /// Other notes related to this note.
        /// </summary>
        public InputList<string> RelatedNoteNames
        {
            get => _relatedNoteNames ?? (_relatedNoteNames = new InputList<string>());
            set => _relatedNoteNames = value;
        }

        [Input("relatedUrl")]
        private InputList<Inputs.RelatedUrlArgs>? _relatedUrl;

        /// <summary>
        /// URLs associated with this note.
        /// </summary>
        public InputList<Inputs.RelatedUrlArgs> RelatedUrl
        {
            get => _relatedUrl ?? (_relatedUrl = new InputList<Inputs.RelatedUrlArgs>());
            set => _relatedUrl = value;
        }

        /// <summary>
        /// A one sentence description of this note.
        /// </summary>
        [Input("shortDescription")]
        public Input<string>? ShortDescription { get; set; }

        /// <summary>
        /// A note describing available package upgrades.
        /// </summary>
        [Input("upgrade")]
        public Input<Inputs.UpgradeNoteArgs>? Upgrade { get; set; }

        /// <summary>
        /// A note describing a package vulnerability.
        /// </summary>
        [Input("vulnerability")]
        public Input<Inputs.VulnerabilityNoteArgs>? Vulnerability { get; set; }

        public NoteArgs()
        {
        }
        public static new NoteArgs Empty => new NoteArgs();
    }
}
