// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Registers a new domain name and creates a corresponding `Registration` resource. Call `RetrieveRegisterParameters` first to check availability of the domain name and determine parameters like price that are needed to build a call to this method. A successful call creates a `Registration` resource in state `REGISTRATION_PENDING`, which resolves to `ACTIVE` within 1-2 minutes, indicating that the domain was successfully registered. If the resource ends up in state `REGISTRATION_FAILED`, it indicates that the domain was not registered successfully, and you can safely delete the resource and retry registration.
 */
export class Registration extends pulumi.CustomResource {
    /**
     * Get an existing Registration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Registration {
        return new Registration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:domains/v1beta1:Registration';

    /**
     * Returns true if the given object is an instance of Registration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Registration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Registration.__pulumiType;
    }

    /**
     * Required. Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
     */
    public readonly contactSettings!: pulumi.Output<outputs.domains.v1beta1.ContactSettingsResponse>;
    /**
     * The creation timestamp of the `Registration` resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
     */
    public readonly dnsSettings!: pulumi.Output<outputs.domains.v1beta1.DnsSettingsResponse>;
    /**
     * Required. Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The expiration timestamp of the `Registration`.
     */
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    /**
     * The set of issues with the `Registration` that require attention.
     */
    public /*out*/ readonly issues!: pulumi.Output<string[]>;
    /**
     * Set of labels associated with the `Registration`.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
     */
    public readonly managementSettings!: pulumi.Output<outputs.domains.v1beta1.ManagementSettingsResponse>;
    /**
     * Name of the `Registration` resource, in the format `projects/*&#47;locations/*&#47;registrations/`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Pending contact settings for the `Registration`. Updates to the `contact_settings` field that change its `registrant_contact` or `privacy` fields require email confirmation by the `registrant_contact` before taking effect. This field is set only if there are pending updates to the `contact_settings` that have not yet been confirmed. To confirm the changes, the `registrant_contact` must follow the instructions in the email they receive.
     */
    public /*out*/ readonly pendingContactSettings!: pulumi.Output<outputs.domains.v1beta1.ContactSettingsResponse>;
    /**
     * The state of the `Registration`
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Set of options for the `contact_settings.privacy` field that this `Registration` supports.
     */
    public /*out*/ readonly supportedPrivacy!: pulumi.Output<string[]>;

    /**
     * Create a Registration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistrationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.registrationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registrationId'");
            }
            inputs["contactNotices"] = args ? args.contactNotices : undefined;
            inputs["contactSettings"] = args ? args.contactSettings : undefined;
            inputs["dnsSettings"] = args ? args.dnsSettings : undefined;
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["domainNotices"] = args ? args.domainNotices : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managementSettings"] = args ? args.managementSettings : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["registrationId"] = args ? args.registrationId : undefined;
            inputs["validateOnly"] = args ? args.validateOnly : undefined;
            inputs["yearlyPrice"] = args ? args.yearlyPrice : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["expireTime"] = undefined /*out*/;
            inputs["issues"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["pendingContactSettings"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["supportedPrivacy"] = undefined /*out*/;
        } else {
            inputs["contactSettings"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["dnsSettings"] = undefined /*out*/;
            inputs["domainName"] = undefined /*out*/;
            inputs["expireTime"] = undefined /*out*/;
            inputs["issues"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["managementSettings"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["pendingContactSettings"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["supportedPrivacy"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Registration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Registration resource.
 */
export interface RegistrationArgs {
    /**
     * The list of contact notices that the caller acknowledges. The notices needed here depend on the values specified in `registration.contact_settings`.
     */
    readonly contactNotices?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required. Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
     */
    readonly contactSettings?: pulumi.Input<inputs.domains.v1beta1.ContactSettingsArgs>;
    /**
     * Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
     */
    readonly dnsSettings?: pulumi.Input<inputs.domains.v1beta1.DnsSettingsArgs>;
    /**
     * Required. Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
     */
    readonly domainName?: pulumi.Input<string>;
    /**
     * The list of domain notices that you acknowledge. Call `RetrieveRegisterParameters` to see the notices that need acknowledgement.
     */
    readonly domainNotices?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set of labels associated with the `Registration`.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly location: pulumi.Input<string>;
    /**
     * Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
     */
    readonly managementSettings?: pulumi.Input<inputs.domains.v1beta1.ManagementSettingsArgs>;
    readonly project: pulumi.Input<string>;
    readonly registrationId: pulumi.Input<string>;
    /**
     * When true, only validation will be performed, without actually registering the domain. Follows: https://cloud.google.com/apis/design/design_patterns#request_validation
     */
    readonly validateOnly?: pulumi.Input<boolean>;
    /**
     * Required. Yearly price to register or renew the domain. The value that should be put here can be obtained from RetrieveRegisterParameters or SearchDomains calls.
     */
    readonly yearlyPrice?: pulumi.Input<inputs.domains.v1beta1.MoneyArgs>;
}
