// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an annotation spec set by providing a set of labels.
 */
export class AnnotationSpecSet extends pulumi.CustomResource {
    /**
     * Get an existing AnnotationSpecSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AnnotationSpecSet {
        return new AnnotationSpecSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:datalabeling/v1beta1:AnnotationSpecSet';

    /**
     * Returns true if the given object is an instance of AnnotationSpecSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AnnotationSpecSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AnnotationSpecSet.__pulumiType;
    }

    /**
     * Required. The array of AnnotationSpecs that you define when you create the AnnotationSpecSet. These are the possible labels for the labeling task.
     */
    public readonly annotationSpecs!: pulumi.Output<outputs.datalabeling.v1beta1.GoogleCloudDatalabelingV1beta1AnnotationSpecResponse[]>;
    /**
     * The names of any related resources that are blocking changes to the annotation spec set.
     */
    public readonly blockingResources!: pulumi.Output<string[]>;
    /**
     * Optional. User-provided description of the annotation specification set. The description can be up to 10,000 characters long.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Required. The display name for AnnotationSpecSet that you define when you create it. Maximum of 64 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The AnnotationSpecSet resource name in the following format: "projects/{project_id}/annotationSpecSets/{annotation_spec_set_id}"
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a AnnotationSpecSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AnnotationSpecSetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.annotationSpecSetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'annotationSpecSetId'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["annotationSpecSetId"] = args ? args.annotationSpecSetId : undefined;
            inputs["annotationSpecs"] = args ? args.annotationSpecs : undefined;
            inputs["blockingResources"] = args ? args.blockingResources : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
        } else {
            inputs["annotationSpecs"] = undefined /*out*/;
            inputs["blockingResources"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AnnotationSpecSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AnnotationSpecSet resource.
 */
export interface AnnotationSpecSetArgs {
    readonly annotationSpecSetId: pulumi.Input<string>;
    /**
     * Required. The array of AnnotationSpecs that you define when you create the AnnotationSpecSet. These are the possible labels for the labeling task.
     */
    readonly annotationSpecs?: pulumi.Input<pulumi.Input<inputs.datalabeling.v1beta1.GoogleCloudDatalabelingV1beta1AnnotationSpecArgs>[]>;
    /**
     * The names of any related resources that are blocking changes to the annotation spec set.
     */
    readonly blockingResources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional. User-provided description of the annotation specification set. The description can be up to 10,000 characters long.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Required. The display name for AnnotationSpecSet that you define when you create it. Maximum of 64 characters.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The AnnotationSpecSet resource name in the following format: "projects/{project_id}/annotationSpecSets/{annotation_spec_set_id}"
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
}
