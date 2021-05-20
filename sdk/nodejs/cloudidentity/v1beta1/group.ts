// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a `Group`.
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
    public static readonly __pulumiType = 'google-native:cloudidentity/v1beta1:Group';

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
     * Additional entity key aliases for a Group.
     */
    public readonly additionalGroupKeys!: pulumi.Output<outputs.cloudidentity.v1beta1.EntityKeyResponse[]>;
    /**
     * The time when the `Group` was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * An extended description to help users determine the purpose of a `Group`. Must not be longer than 4,096 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The display name of the `Group`.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Optional. Dynamic group metadata like queries and status.
     */
    public readonly dynamicGroupMetadata!: pulumi.Output<outputs.cloudidentity.v1beta1.DynamicGroupMetadataResponse>;
    /**
     * Required. Immutable. The `EntityKey` of the `Group`.
     */
    public readonly groupKey!: pulumi.Output<outputs.cloudidentity.v1beta1.EntityKeyResponse>;
    /**
     * Required. One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value. Google Groups are the default type of group and have a label with a key of `cloudidentity.googleapis.com/groups.discussion_forum` and an empty value. Existing Google Groups can have an additional label with a key of `cloudidentity.googleapis.com/groups.security` and an empty value added to them. **This is an immutable change and the security label cannot be removed once added.** Dynamic groups have a label with a key of `cloudidentity.googleapis.com/groups.dynamic`. Identity-mapped groups for Cloud Search have a label with a key of `system/groups/external` and an empty value. Examples: {"cloudidentity.googleapis.com/groups.discussion_forum": ""} or {"system/groups/external": ""}.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Group`. Shall be of the form `groups/{group_id}`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Required. Immutable. The resource name of the entity under which this `Group` resides in the Cloud Identity resource hierarchy. Must be of the form `identitysources/{identity_source_id}` for external- identity-mapped groups or `customers/{customer_id}` for Google Groups.
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * The time when the `Group` was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.initialGroupConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'initialGroupConfig'");
            }
            inputs["additionalGroupKeys"] = args ? args.additionalGroupKeys : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["dynamicGroupMetadata"] = args ? args.dynamicGroupMetadata : undefined;
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["groupKey"] = args ? args.groupKey : undefined;
            inputs["initialGroupConfig"] = args ? args.initialGroupConfig : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["additionalGroupKeys"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["dynamicGroupMetadata"] = undefined /*out*/;
            inputs["groupKey"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parent"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Group.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * Additional entity key aliases for a Group.
     */
    readonly additionalGroupKeys?: pulumi.Input<pulumi.Input<inputs.cloudidentity.v1beta1.EntityKeyArgs>[]>;
    /**
     * An extended description to help users determine the purpose of a `Group`. Must not be longer than 4,096 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The display name of the `Group`.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Optional. Dynamic group metadata like queries and status.
     */
    readonly dynamicGroupMetadata?: pulumi.Input<inputs.cloudidentity.v1beta1.DynamicGroupMetadataArgs>;
    readonly groupId: pulumi.Input<string>;
    /**
     * Required. Immutable. The `EntityKey` of the `Group`.
     */
    readonly groupKey?: pulumi.Input<inputs.cloudidentity.v1beta1.EntityKeyArgs>;
    readonly initialGroupConfig: pulumi.Input<string>;
    /**
     * Required. One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value. Google Groups are the default type of group and have a label with a key of `cloudidentity.googleapis.com/groups.discussion_forum` and an empty value. Existing Google Groups can have an additional label with a key of `cloudidentity.googleapis.com/groups.security` and an empty value added to them. **This is an immutable change and the security label cannot be removed once added.** Dynamic groups have a label with a key of `cloudidentity.googleapis.com/groups.dynamic`. Identity-mapped groups for Cloud Search have a label with a key of `system/groups/external` and an empty value. Examples: {"cloudidentity.googleapis.com/groups.discussion_forum": ""} or {"system/groups/external": ""}.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Required. Immutable. The resource name of the entity under which this `Group` resides in the Cloud Identity resource hierarchy. Must be of the form `identitysources/{identity_source_id}` for external- identity-mapped groups or `customers/{customer_id}` for Google Groups.
     */
    readonly parent?: pulumi.Input<string>;
}
