// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an attestor, and returns a copy of the new attestor. Returns NOT_FOUND if the project does not exist, INVALID_ARGUMENT if the request is malformed, ALREADY_EXISTS if the attestor already exists.
 */
export class Attestor extends pulumi.CustomResource {
    /**
     * Get an existing Attestor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Attestor {
        return new Attestor(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:binaryauthorization/v1:Attestor';

    /**
     * Returns true if the given object is an instance of Attestor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Attestor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Attestor.__pulumiType;
    }

    /**
     * Optional. A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Required. The resource name, in the format: `projects/*&#47;attestors/*`. This field may not be updated.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Time when the attestor was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * This specifies how an attestation will be read, and how it will be used during policy enforcement.
     */
    public readonly userOwnedGrafeasNote!: pulumi.Output<outputs.binaryauthorization.v1.UserOwnedGrafeasNoteResponse>;

    /**
     * Create a Attestor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AttestorArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.attestorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'attestorId'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["attestorId"] = args ? args.attestorId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["userOwnedGrafeasNote"] = args ? args.userOwnedGrafeasNote : undefined;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
            inputs["userOwnedGrafeasNote"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Attestor.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Attestor resource.
 */
export interface AttestorArgs {
    readonly attestorId: pulumi.Input<string>;
    /**
     * Optional. A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Required. The resource name, in the format: `projects/*&#47;attestors/*`. This field may not be updated.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * This specifies how an attestation will be read, and how it will be used during policy enforcement.
     */
    readonly userOwnedGrafeasNote?: pulumi.Input<inputs.binaryauthorization.v1.UserOwnedGrafeasNoteArgs>;
}
