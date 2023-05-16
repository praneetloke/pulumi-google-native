// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets the specified occurrence.
 */
export function getOccurrence(args: GetOccurrenceArgs, opts?: pulumi.InvokeOptions): Promise<GetOccurrenceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:containeranalysis/v1:getOccurrence", {
        "occurrenceId": args.occurrenceId,
        "project": args.project,
    }, opts);
}

export interface GetOccurrenceArgs {
    occurrenceId: string;
    project?: string;
}

export interface GetOccurrenceResult {
    /**
     * Describes an attestation of an artifact.
     */
    readonly attestation: outputs.containeranalysis.v1.AttestationOccurrenceResponse;
    /**
     * Describes a verifiable build.
     */
    readonly build: outputs.containeranalysis.v1.BuildOccurrenceResponse;
    /**
     * Describes a compliance violation on a linked resource.
     */
    readonly compliance: outputs.containeranalysis.v1.ComplianceOccurrenceResponse;
    /**
     * The time this occurrence was created.
     */
    readonly createTime: string;
    /**
     * Describes the deployment of an artifact on a runtime.
     */
    readonly deployment: outputs.containeranalysis.v1.DeploymentOccurrenceResponse;
    /**
     * Describes when a resource was discovered.
     */
    readonly discovery: outputs.containeranalysis.v1.DiscoveryOccurrenceResponse;
    /**
     * Describes an attestation of an artifact using dsse.
     */
    readonly dsseAttestation: outputs.containeranalysis.v1.DSSEAttestationOccurrenceResponse;
    /**
     * https://github.com/secure-systems-lab/dsse
     */
    readonly envelope: outputs.containeranalysis.v1.EnvelopeResponse;
    /**
     * Describes how this resource derives from the basis in the associated note.
     */
    readonly image: outputs.containeranalysis.v1.ImageOccurrenceResponse;
    /**
     * This explicitly denotes which of the occurrence details are specified. This field can be used as a filter in list requests.
     */
    readonly kind: string;
    /**
     * The name of the occurrence in the form of `projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]`.
     */
    readonly name: string;
    /**
     * Immutable. The analysis note associated with this occurrence, in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`. This field can be used as a filter in list requests.
     */
    readonly noteName: string;
    /**
     * Describes the installation of a package on the linked resource.
     */
    readonly package: outputs.containeranalysis.v1.PackageOccurrenceResponse;
    /**
     * A description of actions that can be taken to remedy the note.
     */
    readonly remediation: string;
    /**
     * Immutable. A URI that represents the resource for which the occurrence applies. For example, `https://gcr.io/project/image@sha256:123abc` for a Docker image.
     */
    readonly resourceUri: string;
    /**
     * Describes a specific SBOM reference occurrences.
     */
    readonly sbomReference: outputs.containeranalysis.v1.SBOMReferenceOccurrenceResponse;
    /**
     * The time this occurrence was last updated.
     */
    readonly updateTime: string;
    /**
     * Describes an available package upgrade on the linked resource.
     */
    readonly upgrade: outputs.containeranalysis.v1.UpgradeOccurrenceResponse;
    /**
     * Describes a security vulnerability.
     */
    readonly vulnerability: outputs.containeranalysis.v1.VulnerabilityOccurrenceResponse;
}
/**
 * Gets the specified occurrence.
 */
export function getOccurrenceOutput(args: GetOccurrenceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOccurrenceResult> {
    return pulumi.output(args).apply((a: any) => getOccurrence(a, opts))
}

export interface GetOccurrenceOutputArgs {
    occurrenceId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
