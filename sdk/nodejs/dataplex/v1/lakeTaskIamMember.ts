// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Sets the access control policy on the specified resource. Replaces any existing policy.Can return NOT_FOUND, INVALID_ARGUMENT, and PERMISSION_DENIED errors.
 */
export class LakeTaskIamMember extends pulumi.CustomResource {
    /**
     * Get an existing LakeTaskIamMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LakeTaskIamMember {
        return new LakeTaskIamMember(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dataplex/v1:LakeTaskIamMember';

    /**
     * Returns true if the given object is an instance of LakeTaskIamMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LakeTaskIamMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LakeTaskIamMember.__pulumiType;
    }

    /**
     * An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
     */
    public readonly condition!: pulumi.Output<outputs.iam.v1.Condition | undefined>;
    /**
     * The etag of the resource's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Identity that will be granted the privilege in role. The entry can have one of the following values:
     *
     *  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     *  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     *  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
     *  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     */
    public readonly member!: pulumi.Output<string>;
    /**
     * The name of the resource to manage IAM policies for.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project in which the resource belongs. If it is not provided, a default will be supplied.
     */
    public /*out*/ readonly project!: pulumi.Output<string>;
    /**
     * The role that should be applied.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a LakeTaskIamMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LakeTaskIamMemberArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.member === undefined) && !opts.urn) {
                throw new Error("Missing required property 'member'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["member"] = args ? args.member : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
        } else {
            resourceInputs["condition"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["member"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["role"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LakeTaskIamMember.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a LakeTaskIamMember resource.
 */
export interface LakeTaskIamMemberArgs {
    /**
     * An IAM Condition for a given binding.
     */
    condition?: pulumi.Input<inputs.iam.v1.ConditionArgs>;
    /**
     * Identity that will be granted the privilege in role. The entry can have one of the following values:
     *
     *  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     *  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     *  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
     *  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     */
    member: pulumi.Input<string>;
    /**
     * The name of the resource to manage IAM policies for.
     */
    name: pulumi.Input<string>;
    /**
     * The role that should be applied.
     */
    role: pulumi.Input<string>;
}
