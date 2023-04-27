// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new group.
 * Auto-naming is currently not supported for this resource.
 */
export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:monitoring/v3:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    /**
     * A user-assigned name for this group, used only for display purposes.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The filter used to determine which monitored resources belong to this group.
     */
    public readonly filter!: pulumi.Output<string>;
    /**
     * If true, the members of this group are considered to be a cluster. The system can perform additional analysis on groups that are clusters.
     */
    public readonly isCluster!: pulumi.Output<boolean>;
    /**
     * The name of this group. The format is: projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID] When creating a group, this field is ignored and a new name is created consisting of the project specified in the call to CreateGroup and a unique [GROUP_ID] that is generated automatically.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The name of the group's parent, if it has one. The format is: projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID] For groups with no parent, parent_name is the empty string, "".
     */
    public readonly parentName!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["isCluster"] = args ? args.isCluster : undefined;
            resourceInputs["parentName"] = args ? args.parentName : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["filter"] = undefined /*out*/;
            resourceInputs["isCluster"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parentName"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Group.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * A user-assigned name for this group, used only for display purposes.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The filter used to determine which monitored resources belong to this group.
     */
    filter?: pulumi.Input<string>;
    /**
     * If true, the members of this group are considered to be a cluster. The system can perform additional analysis on groups that are clusters.
     */
    isCluster?: pulumi.Input<boolean>;
    /**
     * The name of the group's parent, if it has one. The format is: projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID] For groups with no parent, parent_name is the empty string, "".
     */
    parentName?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
