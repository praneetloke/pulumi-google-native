// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a taxonomy in the specified project.
 * Auto-naming is currently not supported for this resource.
 */
export class Taxonomy extends pulumi.CustomResource {
    /**
     * Get an existing Taxonomy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Taxonomy {
        return new Taxonomy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:datacatalog/v1beta1:Taxonomy';

    /**
     * Returns true if the given object is an instance of Taxonomy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Taxonomy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Taxonomy.__pulumiType;
    }

    /**
     * Optional. A list of policy types that are activated for this taxonomy. If not set, defaults to an empty list.
     */
    public readonly activatedPolicyTypes!: pulumi.Output<string[]>;
    /**
     * Optional. Description of this taxonomy. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * User defined name of this taxonomy. It must: contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Resource name of this taxonomy, whose format is: "projects/{project_number}/locations/{location_id}/taxonomies/{id}".
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Number of policy tags contained in this taxonomy.
     */
    public /*out*/ readonly policyTagCount!: pulumi.Output<number>;
    /**
     * Timestamps about this taxonomy. Only create_time and update_time are used.
     */
    public /*out*/ readonly taxonomyTimestamps!: pulumi.Output<outputs.datacatalog.v1beta1.GoogleCloudDatacatalogV1beta1SystemTimestampsResponse>;

    /**
     * Create a Taxonomy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TaxonomyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            inputs["activatedPolicyTypes"] = args ? args.activatedPolicyTypes : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["policyTagCount"] = undefined /*out*/;
            inputs["taxonomyTimestamps"] = undefined /*out*/;
        } else {
            inputs["activatedPolicyTypes"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["policyTagCount"] = undefined /*out*/;
            inputs["taxonomyTimestamps"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Taxonomy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Taxonomy resource.
 */
export interface TaxonomyArgs {
    /**
     * Optional. A list of policy types that are activated for this taxonomy. If not set, defaults to an empty list.
     */
    activatedPolicyTypes?: pulumi.Input<pulumi.Input<enums.datacatalog.v1beta1.TaxonomyActivatedPolicyTypesItem>[]>;
    /**
     * Optional. Description of this taxonomy. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description.
     */
    description?: pulumi.Input<string>;
    /**
     * User defined name of this taxonomy. It must: contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
     */
    displayName: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
