// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets the specified note.
 */
export function getNote(args: GetNoteArgs, opts?: pulumi.InvokeOptions): Promise<GetNoteResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:containeranalysis/v1:getNote", {
        "noteId": args.noteId,
        "project": args.project,
    }, opts);
}

export interface GetNoteArgs {
    noteId: string;
    project?: string;
}

export interface GetNoteResult {
    /**
     * A note describing an attestation role.
     */
    readonly attestation: outputs.containeranalysis.v1.AttestationNoteResponse;
    /**
     * A note describing build provenance for a verifiable build.
     */
    readonly build: outputs.containeranalysis.v1.BuildNoteResponse;
    /**
     * A note describing a compliance check.
     */
    readonly compliance: outputs.containeranalysis.v1.ComplianceNoteResponse;
    /**
     * The time this note was created. This field can be used as a filter in list requests.
     */
    readonly createTime: string;
    /**
     * A note describing something that can be deployed.
     */
    readonly deployment: outputs.containeranalysis.v1.DeploymentNoteResponse;
    /**
     * A note describing the initial analysis of a resource.
     */
    readonly discovery: outputs.containeranalysis.v1.DiscoveryNoteResponse;
    /**
     * A note describing a dsse attestation note.
     */
    readonly dsseAttestation: outputs.containeranalysis.v1.DSSEAttestationNoteResponse;
    /**
     * Time of expiration for this note. Empty if note does not expire.
     */
    readonly expirationTime: string;
    /**
     * A note describing a base image.
     */
    readonly image: outputs.containeranalysis.v1.ImageNoteResponse;
    /**
     * The type of analysis. This field can be used as a filter in list requests.
     */
    readonly kind: string;
    /**
     * A detailed description of this note.
     */
    readonly longDescription: string;
    /**
     * The name of the note in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`.
     */
    readonly name: string;
    /**
     * A note describing a package hosted by various package managers.
     */
    readonly package: outputs.containeranalysis.v1.PackageNoteResponse;
    /**
     * Other notes related to this note.
     */
    readonly relatedNoteNames: string[];
    /**
     * URLs associated with this note.
     */
    readonly relatedUrl: outputs.containeranalysis.v1.RelatedUrlResponse[];
    /**
     * A note describing an SBOM reference.
     */
    readonly sbomReference: outputs.containeranalysis.v1.SBOMReferenceNoteResponse;
    /**
     * A one sentence description of this note.
     */
    readonly shortDescription: string;
    /**
     * The time this note was last updated. This field can be used as a filter in list requests.
     */
    readonly updateTime: string;
    /**
     * A note describing available package upgrades.
     */
    readonly upgrade: outputs.containeranalysis.v1.UpgradeNoteResponse;
    /**
     * A note describing a package vulnerability.
     */
    readonly vulnerability: outputs.containeranalysis.v1.VulnerabilityNoteResponse;
    /**
     * A note describing a vulnerability assessment.
     */
    readonly vulnerabilityAssessment: outputs.containeranalysis.v1.VulnerabilityAssessmentNoteResponse;
}
/**
 * Gets the specified note.
 */
export function getNoteOutput(args: GetNoteOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNoteResult> {
    return pulumi.output(args).apply((a: any) => getNote(a, opts))
}

export interface GetNoteOutputArgs {
    noteId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
