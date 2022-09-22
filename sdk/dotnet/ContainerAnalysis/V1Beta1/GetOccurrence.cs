// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1
{
    public static class GetOccurrence
    {
        /// <summary>
        /// Gets the specified occurrence.
        /// </summary>
        public static Task<GetOccurrenceResult> InvokeAsync(GetOccurrenceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOccurrenceResult>("google-native:containeranalysis/v1beta1:getOccurrence", args ?? new GetOccurrenceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified occurrence.
        /// </summary>
        public static Output<GetOccurrenceResult> Invoke(GetOccurrenceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOccurrenceResult>("google-native:containeranalysis/v1beta1:getOccurrence", args ?? new GetOccurrenceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOccurrenceArgs : global::Pulumi.InvokeArgs
    {
        [Input("occurrenceId", required: true)]
        public string OccurrenceId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetOccurrenceArgs()
        {
        }
        public static new GetOccurrenceArgs Empty => new GetOccurrenceArgs();
    }

    public sealed class GetOccurrenceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("occurrenceId", required: true)]
        public Input<string> OccurrenceId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetOccurrenceInvokeArgs()
        {
        }
        public static new GetOccurrenceInvokeArgs Empty => new GetOccurrenceInvokeArgs();
    }


    [OutputType]
    public sealed class GetOccurrenceResult
    {
        /// <summary>
        /// Describes an attestation of an artifact.
        /// </summary>
        public readonly Outputs.DetailsResponse Attestation;
        /// <summary>
        /// Describes a verifiable build.
        /// </summary>
        public readonly Outputs.GrafeasV1beta1BuildDetailsResponse Build;
        /// <summary>
        /// The time this occurrence was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Describes the deployment of an artifact on a runtime.
        /// </summary>
        public readonly Outputs.GrafeasV1beta1DeploymentDetailsResponse Deployment;
        /// <summary>
        /// Describes how this resource derives from the basis in the associated note.
        /// </summary>
        public readonly Outputs.GrafeasV1beta1ImageDetailsResponse DerivedImage;
        /// <summary>
        /// Describes when a resource was discovered.
        /// </summary>
        public readonly Outputs.GrafeasV1beta1DiscoveryDetailsResponse Discovered;
        /// <summary>
        /// https://github.com/secure-systems-lab/dsse
        /// </summary>
        public readonly Outputs.EnvelopeResponse Envelope;
        /// <summary>
        /// Describes the installation of a package on the linked resource.
        /// </summary>
        public readonly Outputs.GrafeasV1beta1PackageDetailsResponse Installation;
        /// <summary>
        /// Describes a specific in-toto link.
        /// </summary>
        public readonly Outputs.GrafeasV1beta1IntotoDetailsResponse Intoto;
        /// <summary>
        /// This explicitly denotes which of the occurrence details are specified. This field can be used as a filter in list requests.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The name of the occurrence in the form of `projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Immutable. The analysis note associated with this occurrence, in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`. This field can be used as a filter in list requests.
        /// </summary>
        public readonly string NoteName;
        /// <summary>
        /// A description of actions that can be taken to remedy the note.
        /// </summary>
        public readonly string Remediation;
        /// <summary>
        /// Immutable. The resource for which the occurrence applies.
        /// </summary>
        public readonly Outputs.ResourceResponse Resource;
        /// <summary>
        /// Describes a specific software bill of materials document.
        /// </summary>
        public readonly Outputs.DocumentOccurrenceResponse Sbom;
        /// <summary>
        /// Describes a specific SPDX File.
        /// </summary>
        public readonly Outputs.FileOccurrenceResponse SpdxFile;
        /// <summary>
        /// Describes a specific SPDX Package.
        /// </summary>
        public readonly Outputs.PackageInfoOccurrenceResponse SpdxPackage;
        /// <summary>
        /// Describes a specific SPDX Relationship.
        /// </summary>
        public readonly Outputs.RelationshipOccurrenceResponse SpdxRelationship;
        /// <summary>
        /// The time this occurrence was last updated.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// Describes a security vulnerability.
        /// </summary>
        public readonly Outputs.GrafeasV1beta1VulnerabilityDetailsResponse Vulnerability;

        [OutputConstructor]
        private GetOccurrenceResult(
            Outputs.DetailsResponse attestation,

            Outputs.GrafeasV1beta1BuildDetailsResponse build,

            string createTime,

            Outputs.GrafeasV1beta1DeploymentDetailsResponse deployment,

            Outputs.GrafeasV1beta1ImageDetailsResponse derivedImage,

            Outputs.GrafeasV1beta1DiscoveryDetailsResponse discovered,

            Outputs.EnvelopeResponse envelope,

            Outputs.GrafeasV1beta1PackageDetailsResponse installation,

            Outputs.GrafeasV1beta1IntotoDetailsResponse intoto,

            string kind,

            string name,

            string noteName,

            string remediation,

            Outputs.ResourceResponse resource,

            Outputs.DocumentOccurrenceResponse sbom,

            Outputs.FileOccurrenceResponse spdxFile,

            Outputs.PackageInfoOccurrenceResponse spdxPackage,

            Outputs.RelationshipOccurrenceResponse spdxRelationship,

            string updateTime,

            Outputs.GrafeasV1beta1VulnerabilityDetailsResponse vulnerability)
        {
            Attestation = attestation;
            Build = build;
            CreateTime = createTime;
            Deployment = deployment;
            DerivedImage = derivedImage;
            Discovered = discovered;
            Envelope = envelope;
            Installation = installation;
            Intoto = intoto;
            Kind = kind;
            Name = name;
            NoteName = noteName;
            Remediation = remediation;
            Resource = resource;
            Sbom = sbom;
            SpdxFile = spdxFile;
            SpdxPackage = spdxPackage;
            SpdxRelationship = spdxRelationship;
            UpdateTime = updateTime;
            Vulnerability = vulnerability;
        }
    }
}
