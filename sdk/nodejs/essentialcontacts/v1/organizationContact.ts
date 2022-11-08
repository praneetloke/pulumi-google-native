// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Adds a new contact for a resource.
 * Auto-naming is currently not supported for this resource.
 */
export class OrganizationContact extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationContact resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OrganizationContact {
        return new OrganizationContact(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:essentialcontacts/v1:OrganizationContact';

    /**
     * Returns true if the given object is an instance of OrganizationContact.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationContact {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationContact.__pulumiType;
    }

    /**
     * The email address to send notifications to. The email address does not need to be a Google Account.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
     */
    public readonly languageTag!: pulumi.Output<string>;
    /**
     * The identifier for the contact. Format: {resource_type}/{resource_id}/contacts/{contact_id}
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The categories of notifications that the contact will receive communications for.
     */
    public readonly notificationCategorySubscriptions!: pulumi.Output<string[]>;
    public readonly organizationId!: pulumi.Output<string>;
    /**
     * The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
     */
    public readonly validateTime!: pulumi.Output<string>;
    /**
     * The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
     */
    public readonly validationState!: pulumi.Output<string>;

    /**
     * Create a OrganizationContact resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationContactArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.languageTag === undefined) && !opts.urn) {
                throw new Error("Missing required property 'languageTag'");
            }
            if ((!args || args.notificationCategorySubscriptions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notificationCategorySubscriptions'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["languageTag"] = args ? args.languageTag : undefined;
            resourceInputs["notificationCategorySubscriptions"] = args ? args.notificationCategorySubscriptions : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["validateTime"] = args ? args.validateTime : undefined;
            resourceInputs["validationState"] = args ? args.validationState : undefined;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["email"] = undefined /*out*/;
            resourceInputs["languageTag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["notificationCategorySubscriptions"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["validateTime"] = undefined /*out*/;
            resourceInputs["validationState"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["organizationId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(OrganizationContact.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a OrganizationContact resource.
 */
export interface OrganizationContactArgs {
    /**
     * The email address to send notifications to. The email address does not need to be a Google Account.
     */
    email: pulumi.Input<string>;
    /**
     * The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
     */
    languageTag: pulumi.Input<string>;
    /**
     * The categories of notifications that the contact will receive communications for.
     */
    notificationCategorySubscriptions: pulumi.Input<pulumi.Input<enums.essentialcontacts.v1.OrganizationContactNotificationCategorySubscriptionsItem>[]>;
    organizationId: pulumi.Input<string>;
    /**
     * The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
     */
    validateTime?: pulumi.Input<string>;
    /**
     * The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
     */
    validationState?: pulumi.Input<enums.essentialcontacts.v1.OrganizationContactValidationState>;
}
