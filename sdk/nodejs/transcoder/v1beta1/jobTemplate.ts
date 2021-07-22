// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a job template in the specified region.
 */
export class JobTemplate extends pulumi.CustomResource {
    /**
     * Get an existing JobTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): JobTemplate {
        return new JobTemplate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:transcoder/v1beta1:JobTemplate';

    /**
     * Returns true if the given object is an instance of JobTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is JobTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === JobTemplate.__pulumiType;
    }

    /**
     * The configuration for this template.
     */
    public readonly config!: pulumi.Output<outputs.transcoder.v1beta1.JobConfigResponse>;
    /**
     * The resource name of the job template. Format: `projects/{project}/locations/{location}/jobTemplates/{job_template}`
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a JobTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.jobTemplateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobTemplateId'");
            }
            inputs["config"] = args ? args.config : undefined;
            inputs["jobTemplateId"] = args ? args.jobTemplateId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
        } else {
            inputs["config"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(JobTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a JobTemplate resource.
 */
export interface JobTemplateArgs {
    /**
     * The configuration for this template.
     */
    config?: pulumi.Input<inputs.transcoder.v1beta1.JobConfigArgs>;
    jobTemplateId: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * The resource name of the job template. Format: `projects/{project}/locations/{location}/jobTemplates/{job_template}`
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
