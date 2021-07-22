// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a policy tag in the specified taxonomy.
 * Auto-naming is currently not supported for this resource.
 */
export class PolicyTag extends pulumi.CustomResource {
    /**
     * Get an existing PolicyTag resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PolicyTag {
        return new PolicyTag(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:datacatalog/v1beta1:PolicyTag';

    /**
     * Returns true if the given object is an instance of PolicyTag.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyTag {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyTag.__pulumiType;
    }

    /**
     * Resource names of child policy tags of this policy tag.
     */
    public /*out*/ readonly childPolicyTags!: pulumi.Output<string[]>;
    /**
     * Description of this policy tag. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description. If not set, defaults to an empty description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * User defined name of this policy tag. It must: be unique within the parent taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Resource name of this policy tag, whose format is: "projects/{project_number}/locations/{location_id}/taxonomies/{taxonomy_id}/policyTags/{id}".
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource name of this policy tag's parent policy tag (e.g. for the "LatLong" policy tag in the example above, this field contains the resource name of the "Geolocation" policy tag). If empty, it means this policy tag is a top level policy tag (e.g. this field is empty for the "Geolocation" policy tag in the example above). If not set, defaults to an empty string.
     */
    public readonly parentPolicyTag!: pulumi.Output<string>;

    /**
     * Create a PolicyTag resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyTagArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.taxonomyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taxonomyId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["parentPolicyTag"] = args ? args.parentPolicyTag : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["taxonomyId"] = args ? args.taxonomyId : undefined;
            inputs["childPolicyTags"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        } else {
            inputs["childPolicyTags"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parentPolicyTag"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(PolicyTag.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PolicyTag resource.
 */
export interface PolicyTagArgs {
    /**
     * Description of this policy tag. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description. If not set, defaults to an empty description.
     */
    description?: pulumi.Input<string>;
    /**
     * User defined name of this policy tag. It must: be unique within the parent taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
     */
    displayName: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Resource name of this policy tag's parent policy tag (e.g. for the "LatLong" policy tag in the example above, this field contains the resource name of the "Geolocation" policy tag). If empty, it means this policy tag is a top level policy tag (e.g. this field is empty for the "Geolocation" policy tag in the example above). If not set, defaults to an empty string.
     */
    parentPolicyTag?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    taxonomyId: pulumi.Input<string>;
}
