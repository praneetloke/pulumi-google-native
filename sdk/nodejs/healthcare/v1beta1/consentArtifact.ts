// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Consent artifact in the parent consent store.
 */
export class ConsentArtifact extends pulumi.CustomResource {
    /**
     * Get an existing ConsentArtifact resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConsentArtifact {
        return new ConsentArtifact(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:healthcare/v1beta1:ConsentArtifact';

    /**
     * Returns true if the given object is an instance of ConsentArtifact.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConsentArtifact {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConsentArtifact.__pulumiType;
    }

    /**
     * Optional. Screenshots, PDFs, or other binary information documenting the user's consent.
     */
    public readonly consentContentScreenshots!: pulumi.Output<outputs.healthcare.v1beta1.ImageResponse[]>;
    /**
     * Optional. An string indicating the version of the consent information shown to the user.
     */
    public readonly consentContentVersion!: pulumi.Output<string>;
    /**
     * Optional. A signature from a guardian.
     */
    public readonly guardianSignature!: pulumi.Output<outputs.healthcare.v1beta1.SignatureResponse>;
    /**
     * Optional. Metadata associated with the Consent artifact. For example, the consent locale or user agent version.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource name of the Consent artifact, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`. Cannot be changed after creation.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * User's UUID provided by the client.
     */
    public readonly userId!: pulumi.Output<string>;
    /**
     * Optional. User's signature.
     */
    public readonly userSignature!: pulumi.Output<outputs.healthcare.v1beta1.SignatureResponse>;
    /**
     * Optional. A signature from a witness.
     */
    public readonly witnessSignature!: pulumi.Output<outputs.healthcare.v1beta1.SignatureResponse>;

    /**
     * Create a ConsentArtifact resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConsentArtifactArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.consentStoreId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'consentStoreId'");
            }
            if ((!args || args.datasetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasetId'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            inputs["consentContentScreenshots"] = args ? args.consentContentScreenshots : undefined;
            inputs["consentContentVersion"] = args ? args.consentContentVersion : undefined;
            inputs["consentStoreId"] = args ? args.consentStoreId : undefined;
            inputs["datasetId"] = args ? args.datasetId : undefined;
            inputs["guardianSignature"] = args ? args.guardianSignature : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["userId"] = args ? args.userId : undefined;
            inputs["userSignature"] = args ? args.userSignature : undefined;
            inputs["witnessSignature"] = args ? args.witnessSignature : undefined;
        } else {
            inputs["consentContentScreenshots"] = undefined /*out*/;
            inputs["consentContentVersion"] = undefined /*out*/;
            inputs["guardianSignature"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["userId"] = undefined /*out*/;
            inputs["userSignature"] = undefined /*out*/;
            inputs["witnessSignature"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ConsentArtifact.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConsentArtifact resource.
 */
export interface ConsentArtifactArgs {
    /**
     * Optional. Screenshots, PDFs, or other binary information documenting the user's consent.
     */
    consentContentScreenshots?: pulumi.Input<pulumi.Input<inputs.healthcare.v1beta1.ImageArgs>[]>;
    /**
     * Optional. An string indicating the version of the consent information shown to the user.
     */
    consentContentVersion?: pulumi.Input<string>;
    consentStoreId: pulumi.Input<string>;
    datasetId: pulumi.Input<string>;
    /**
     * Optional. A signature from a guardian.
     */
    guardianSignature?: pulumi.Input<inputs.healthcare.v1beta1.SignatureArgs>;
    location?: pulumi.Input<string>;
    /**
     * Optional. Metadata associated with the Consent artifact. For example, the consent locale or user agent version.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Resource name of the Consent artifact, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`. Cannot be changed after creation.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * User's UUID provided by the client.
     */
    userId: pulumi.Input<string>;
    /**
     * Optional. User's signature.
     */
    userSignature?: pulumi.Input<inputs.healthcare.v1beta1.SignatureArgs>;
    /**
     * Optional. A signature from a witness.
     */
    witnessSignature?: pulumi.Input<inputs.healthcare.v1beta1.SignatureArgs>;
}
