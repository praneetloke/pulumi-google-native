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
export class FederationIamMember extends pulumi.CustomResource {
    /**
     * Get an existing FederationIamMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FederationIamMember {
        return new FederationIamMember(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:metastore/v1beta:FederationIamMember';

    /**
     * Returns true if the given object is an instance of FederationIamMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FederationIamMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FederationIamMember.__pulumiType;
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
     * Specifies the principals requesting access for a Google Cloud resource. members can have the following values: allUsers: A special identifier that represents anyone who is on the internet; with or without a Google account. allAuthenticatedUsers: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. user:{emailid}: An email address that represents a specific Google account. For example, alice@example.com . serviceAccount:{emailid}: An email address that represents a Google service account. For example, my-other-app@appspot.gserviceaccount.com. serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]: An identifier for a Kubernetes service account (https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, my-project.svc.id.goog[my-namespace/my-kubernetes-sa]. group:{emailid}: An email address that represents a Google group. For example, admins@example.com. deleted:user:{emailid}?uid={uniqueid}: An email address (plus unique identifier) representing a user that has been recently deleted. For example, alice@example.com?uid=123456789012345678901. If the user is recovered, this value reverts to user:{emailid} and the recovered user retains the role in the binding. deleted:serviceAccount:{emailid}?uid={uniqueid}: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901. If the service account is undeleted, this value reverts to serviceAccount:{emailid} and the undeleted service account retains the role in the binding. deleted:group:{emailid}?uid={uniqueid}: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, admins@example.com?uid=123456789012345678901. If the group is recovered, this value reverts to group:{emailid} and the recovered group retains the role in the binding. domain:{domain}: The G Suite domain (primary) that represents all the users of that domain. For example, google.com or example.com.
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
     * Role that is assigned to the list of members, or principals. For example, roles/viewer, roles/editor, or roles/owner.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a FederationIamMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FederationIamMemberArgs, opts?: pulumi.CustomResourceOptions) {
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
        super(FederationIamMember.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a FederationIamMember resource.
 */
export interface FederationIamMemberArgs {
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
