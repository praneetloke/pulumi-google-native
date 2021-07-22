// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a glossary and returns the long-running operation. Returns NOT_FOUND, if the project doesn't exist.
 */
export class Glossary extends pulumi.CustomResource {
    /**
     * Get an existing Glossary resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Glossary {
        return new Glossary(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:translate/v3:Glossary';

    /**
     * Returns true if the given object is an instance of Glossary.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Glossary {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Glossary.__pulumiType;
    }

    /**
     * When the glossary creation was finished.
     */
    public /*out*/ readonly endTime!: pulumi.Output<string>;
    /**
     * The number of entries defined in the glossary.
     */
    public /*out*/ readonly entryCount!: pulumi.Output<number>;
    /**
     * Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
     */
    public readonly inputConfig!: pulumi.Output<outputs.translate.v3.GlossaryInputConfigResponse>;
    /**
     * Used with equivalent term set glossaries.
     */
    public readonly languageCodesSet!: pulumi.Output<outputs.translate.v3.LanguageCodesSetResponse>;
    /**
     * Used with unidirectional glossaries.
     */
    public readonly languagePair!: pulumi.Output<outputs.translate.v3.LanguageCodePairResponse>;
    /**
     * The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * When CreateGlossary was called.
     */
    public /*out*/ readonly submitTime!: pulumi.Output<string>;

    /**
     * Create a Glossary resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GlossaryArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.inputConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'inputConfig'");
            }
            inputs["inputConfig"] = args ? args.inputConfig : undefined;
            inputs["languageCodesSet"] = args ? args.languageCodesSet : undefined;
            inputs["languagePair"] = args ? args.languagePair : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["endTime"] = undefined /*out*/;
            inputs["entryCount"] = undefined /*out*/;
            inputs["submitTime"] = undefined /*out*/;
        } else {
            inputs["endTime"] = undefined /*out*/;
            inputs["entryCount"] = undefined /*out*/;
            inputs["inputConfig"] = undefined /*out*/;
            inputs["languageCodesSet"] = undefined /*out*/;
            inputs["languagePair"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["submitTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Glossary.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Glossary resource.
 */
export interface GlossaryArgs {
    /**
     * Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
     */
    inputConfig: pulumi.Input<inputs.translate.v3.GlossaryInputConfigArgs>;
    /**
     * Used with equivalent term set glossaries.
     */
    languageCodesSet?: pulumi.Input<inputs.translate.v3.LanguageCodesSetArgs>;
    /**
     * Used with unidirectional glossaries.
     */
    languagePair?: pulumi.Input<inputs.translate.v3.LanguageCodePairArgs>;
    location?: pulumi.Input<string>;
    /**
     * The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
