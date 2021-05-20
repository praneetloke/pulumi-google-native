// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a DeidentifyTemplate for re-using frequently used configuration for de-identifying content, images, and storage. See https://cloud.google.com/dlp/docs/creating-templates-deid to learn more.
 */
export class DeidentifyTemplate extends pulumi.CustomResource {
    /**
     * Get an existing DeidentifyTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DeidentifyTemplate {
        return new DeidentifyTemplate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dlp/v2:DeidentifyTemplate';

    /**
     * Returns true if the given object is an instance of DeidentifyTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeidentifyTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeidentifyTemplate.__pulumiType;
    }

    /**
     * The creation timestamp of an inspectTemplate.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The core content of the template.
     */
    public readonly deidentifyConfig!: pulumi.Output<outputs.dlp.v2.GooglePrivacyDlpV2DeidentifyConfigResponse>;
    /**
     * Short description (max 256 chars).
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Display name (max 256 chars).
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The template name. The template will have one of the following formats: `projects/PROJECT_ID/deidentifyTemplates/TEMPLATE_ID` OR `organizations/ORGANIZATION_ID/deidentifyTemplates/TEMPLATE_ID`
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The last update timestamp of an inspectTemplate.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a DeidentifyTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeidentifyTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.deidentifyTemplateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deidentifyTemplateId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["deidentifyConfig"] = args ? args.deidentifyConfig : undefined;
            inputs["deidentifyTemplateId"] = args ? args.deidentifyTemplateId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["templateId"] = args ? args.templateId : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["createTime"] = undefined /*out*/;
            inputs["deidentifyConfig"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DeidentifyTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DeidentifyTemplate resource.
 */
export interface DeidentifyTemplateArgs {
    /**
     * The core content of the template.
     */
    readonly deidentifyConfig?: pulumi.Input<inputs.dlp.v2.GooglePrivacyDlpV2DeidentifyConfigArgs>;
    readonly deidentifyTemplateId: pulumi.Input<string>;
    /**
     * Short description (max 256 chars).
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Display name (max 256 chars).
     */
    readonly displayName?: pulumi.Input<string>;
    readonly location: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * The template id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
     */
    readonly templateId?: pulumi.Input<string>;
}
