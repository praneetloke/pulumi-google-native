// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1
{
    public static class GetNote
    {
        /// <summary>
        /// Returns the requested `Note`.
        /// </summary>
        public static Task<GetNoteResult> InvokeAsync(GetNoteArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNoteResult>("google-native:containeranalysis/v1alpha1:getNote", args ?? new GetNoteArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the requested `Note`.
        /// </summary>
        public static Output<GetNoteResult> Invoke(GetNoteInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNoteResult>("google-native:containeranalysis/v1alpha1:getNote", args ?? new GetNoteInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNoteArgs : global::Pulumi.InvokeArgs
    {
        [Input("noteId", required: true)]
        public string NoteId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetNoteArgs()
        {
        }
        public static new GetNoteArgs Empty => new GetNoteArgs();
    }

    public sealed class GetNoteInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("noteId", required: true)]
        public Input<string> NoteId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetNoteInvokeArgs()
        {
        }
        public static new GetNoteInvokeArgs Empty => new GetNoteInvokeArgs();
    }


    [OutputType]
    public sealed class GetNoteResult
    {
        /// <summary>
        /// A note describing an attestation role.
        /// </summary>
        public readonly Outputs.AttestationAuthorityResponse AttestationAuthority;
        /// <summary>
        /// A note describing a base image.
        /// </summary>
        public readonly Outputs.BasisResponse BaseImage;
        /// <summary>
        /// Build provenance type for a verifiable build.
        /// </summary>
        public readonly Outputs.BuildTypeResponse BuildType;
        /// <summary>
        /// A note describing a compliance check.
        /// </summary>
        public readonly Outputs.ComplianceNoteResponse Compliance;
        /// <summary>
        /// The time this note was created. This field can be used as a filter in list requests.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// A note describing something that can be deployed.
        /// </summary>
        public readonly Outputs.DeployableResponse Deployable;
        /// <summary>
        /// A note describing a provider/analysis type.
        /// </summary>
        public readonly Outputs.DiscoveryResponse Discovery;
        /// <summary>
        /// A note describing a dsse attestation note.
        /// </summary>
        public readonly Outputs.DSSEAttestationNoteResponse DsseAttestation;
        /// <summary>
        /// Time of expiration for this note, null if note does not expire.
        /// </summary>
        public readonly string ExpirationTime;
        /// <summary>
        /// This explicitly denotes which kind of note is specified. This field can be used as a filter in list requests.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// A detailed description of this `Note`.
        /// </summary>
        public readonly string LongDescription;
        /// <summary>
        /// The name of the note in the form "projects/{provider_project_id}/notes/{NOTE_ID}"
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A note describing a package hosted by various package managers.
        /// </summary>
        public readonly Outputs.PackageResponse Package;
        /// <summary>
        /// URLs associated with this note
        /// </summary>
        public readonly ImmutableArray<Outputs.RelatedUrlResponse> RelatedUrl;
        /// <summary>
        /// A note describing a software bill of materials.
        /// </summary>
        public readonly Outputs.DocumentNoteResponse Sbom;
        /// <summary>
        /// A one sentence description of this `Note`.
        /// </summary>
        public readonly string ShortDescription;
        /// <summary>
        /// A note describing an SPDX File.
        /// </summary>
        public readonly Outputs.FileNoteResponse SpdxFile;
        /// <summary>
        /// A note describing an SPDX Package.
        /// </summary>
        public readonly Outputs.PackageInfoNoteResponse SpdxPackage;
        /// <summary>
        /// A note describing a relationship between SPDX elements.
        /// </summary>
        public readonly Outputs.RelationshipNoteResponse SpdxRelationship;
        /// <summary>
        /// The time this note was last updated. This field can be used as a filter in list requests.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// A note describing an upgrade.
        /// </summary>
        public readonly Outputs.UpgradeNoteResponse Upgrade;
        /// <summary>
        /// A package vulnerability type of note.
        /// </summary>
        public readonly Outputs.VulnerabilityTypeResponse VulnerabilityType;

        [OutputConstructor]
        private GetNoteResult(
            Outputs.AttestationAuthorityResponse attestationAuthority,

            Outputs.BasisResponse baseImage,

            Outputs.BuildTypeResponse buildType,

            Outputs.ComplianceNoteResponse compliance,

            string createTime,

            Outputs.DeployableResponse deployable,

            Outputs.DiscoveryResponse discovery,

            Outputs.DSSEAttestationNoteResponse dsseAttestation,

            string expirationTime,

            string kind,

            string longDescription,

            string name,

            Outputs.PackageResponse package,

            ImmutableArray<Outputs.RelatedUrlResponse> relatedUrl,

            Outputs.DocumentNoteResponse sbom,

            string shortDescription,

            Outputs.FileNoteResponse spdxFile,

            Outputs.PackageInfoNoteResponse spdxPackage,

            Outputs.RelationshipNoteResponse spdxRelationship,

            string updateTime,

            Outputs.UpgradeNoteResponse upgrade,

            Outputs.VulnerabilityTypeResponse vulnerabilityType)
        {
            AttestationAuthority = attestationAuthority;
            BaseImage = baseImage;
            BuildType = buildType;
            Compliance = compliance;
            CreateTime = createTime;
            Deployable = deployable;
            Discovery = discovery;
            DsseAttestation = dsseAttestation;
            ExpirationTime = expirationTime;
            Kind = kind;
            LongDescription = longDescription;
            Name = name;
            Package = package;
            RelatedUrl = relatedUrl;
            Sbom = sbom;
            ShortDescription = shortDescription;
            SpdxFile = spdxFile;
            SpdxPackage = spdxPackage;
            SpdxRelationship = spdxRelationship;
            UpdateTime = updateTime;
            Upgrade = upgrade;
            VulnerabilityType = vulnerabilityType;
        }
    }
}
