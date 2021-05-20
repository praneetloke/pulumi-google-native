// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new note.
 */
export class Note extends pulumi.CustomResource {
    /**
     * Get an existing Note resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Note {
        return new Note(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:containeranalysis/v1beta1:Note';

    /**
     * Returns true if the given object is an instance of Note.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Note {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Note.__pulumiType;
    }

    /**
     * A note describing an attestation role.
     */
    public readonly attestationAuthority!: pulumi.Output<outputs.containeranalysis.v1beta1.AuthorityResponse>;
    /**
     * A note describing a base image.
     */
    public readonly baseImage!: pulumi.Output<outputs.containeranalysis.v1beta1.BasisResponse>;
    /**
     * A note describing build provenance for a verifiable build.
     */
    public readonly build!: pulumi.Output<outputs.containeranalysis.v1beta1.BuildResponse>;
    /**
     * The time this note was created. This field can be used as a filter in list requests.
     */
    public readonly createTime!: pulumi.Output<string>;
    /**
     * A note describing something that can be deployed.
     */
    public readonly deployable!: pulumi.Output<outputs.containeranalysis.v1beta1.DeployableResponse>;
    /**
     * A note describing the initial analysis of a resource.
     */
    public readonly discovery!: pulumi.Output<outputs.containeranalysis.v1beta1.DiscoveryResponse>;
    /**
     * Time of expiration for this note. Empty if note does not expire.
     */
    public readonly expirationTime!: pulumi.Output<string>;
    /**
     * A note describing an in-toto link.
     */
    public readonly intoto!: pulumi.Output<outputs.containeranalysis.v1beta1.InTotoResponse>;
    /**
     * The type of analysis. This field can be used as a filter in list requests.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * A detailed description of this note.
     */
    public readonly longDescription!: pulumi.Output<string>;
    /**
     * The name of the note in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A note describing a package hosted by various package managers.
     */
    public readonly package!: pulumi.Output<outputs.containeranalysis.v1beta1.PackageResponse>;
    /**
     * Other notes related to this note.
     */
    public readonly relatedNoteNames!: pulumi.Output<string[]>;
    /**
     * URLs associated with this note.
     */
    public readonly relatedUrl!: pulumi.Output<outputs.containeranalysis.v1beta1.RelatedUrlResponse[]>;
    /**
     * A one sentence description of this note.
     */
    public readonly shortDescription!: pulumi.Output<string>;
    /**
     * The time this note was last updated. This field can be used as a filter in list requests.
     */
    public readonly updateTime!: pulumi.Output<string>;
    /**
     * A note describing a package vulnerability.
     */
    public readonly vulnerability!: pulumi.Output<outputs.containeranalysis.v1beta1.VulnerabilityResponse>;

    /**
     * Create a Note resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NoteArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.noteId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'noteId'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["attestationAuthority"] = args ? args.attestationAuthority : undefined;
            inputs["baseImage"] = args ? args.baseImage : undefined;
            inputs["build"] = args ? args.build : undefined;
            inputs["createTime"] = args ? args.createTime : undefined;
            inputs["deployable"] = args ? args.deployable : undefined;
            inputs["discovery"] = args ? args.discovery : undefined;
            inputs["expirationTime"] = args ? args.expirationTime : undefined;
            inputs["intoto"] = args ? args.intoto : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["longDescription"] = args ? args.longDescription : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["noteId"] = args ? args.noteId : undefined;
            inputs["package"] = args ? args.package : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["relatedNoteNames"] = args ? args.relatedNoteNames : undefined;
            inputs["relatedUrl"] = args ? args.relatedUrl : undefined;
            inputs["shortDescription"] = args ? args.shortDescription : undefined;
            inputs["updateTime"] = args ? args.updateTime : undefined;
            inputs["vulnerability"] = args ? args.vulnerability : undefined;
        } else {
            inputs["attestationAuthority"] = undefined /*out*/;
            inputs["baseImage"] = undefined /*out*/;
            inputs["build"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["deployable"] = undefined /*out*/;
            inputs["discovery"] = undefined /*out*/;
            inputs["expirationTime"] = undefined /*out*/;
            inputs["intoto"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["longDescription"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["package"] = undefined /*out*/;
            inputs["relatedNoteNames"] = undefined /*out*/;
            inputs["relatedUrl"] = undefined /*out*/;
            inputs["shortDescription"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
            inputs["vulnerability"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Note.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Note resource.
 */
export interface NoteArgs {
    /**
     * A note describing an attestation role.
     */
    readonly attestationAuthority?: pulumi.Input<inputs.containeranalysis.v1beta1.AuthorityArgs>;
    /**
     * A note describing a base image.
     */
    readonly baseImage?: pulumi.Input<inputs.containeranalysis.v1beta1.BasisArgs>;
    /**
     * A note describing build provenance for a verifiable build.
     */
    readonly build?: pulumi.Input<inputs.containeranalysis.v1beta1.BuildArgs>;
    /**
     * The time this note was created. This field can be used as a filter in list requests.
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * A note describing something that can be deployed.
     */
    readonly deployable?: pulumi.Input<inputs.containeranalysis.v1beta1.DeployableArgs>;
    /**
     * A note describing the initial analysis of a resource.
     */
    readonly discovery?: pulumi.Input<inputs.containeranalysis.v1beta1.DiscoveryArgs>;
    /**
     * Time of expiration for this note. Empty if note does not expire.
     */
    readonly expirationTime?: pulumi.Input<string>;
    /**
     * A note describing an in-toto link.
     */
    readonly intoto?: pulumi.Input<inputs.containeranalysis.v1beta1.InTotoArgs>;
    /**
     * The type of analysis. This field can be used as a filter in list requests.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * A detailed description of this note.
     */
    readonly longDescription?: pulumi.Input<string>;
    /**
     * The name of the note in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`.
     */
    readonly name?: pulumi.Input<string>;
    readonly noteId: pulumi.Input<string>;
    /**
     * A note describing a package hosted by various package managers.
     */
    readonly package?: pulumi.Input<inputs.containeranalysis.v1beta1.PackageArgs>;
    readonly project: pulumi.Input<string>;
    /**
     * Other notes related to this note.
     */
    readonly relatedNoteNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * URLs associated with this note.
     */
    readonly relatedUrl?: pulumi.Input<pulumi.Input<inputs.containeranalysis.v1beta1.RelatedUrlArgs>[]>;
    /**
     * A one sentence description of this note.
     */
    readonly shortDescription?: pulumi.Input<string>;
    /**
     * The time this note was last updated. This field can be used as a filter in list requests.
     */
    readonly updateTime?: pulumi.Input<string>;
    /**
     * A note describing a package vulnerability.
     */
    readonly vulnerability?: pulumi.Input<inputs.containeranalysis.v1beta1.VulnerabilityArgs>;
}
