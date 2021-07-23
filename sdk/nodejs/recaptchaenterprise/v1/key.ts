// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new reCAPTCHA Enterprise key.
 */
export class Key extends pulumi.CustomResource {
    /**
     * Get an existing Key resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Key {
        return new Key(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:recaptchaenterprise/v1:Key';

    /**
     * Returns true if the given object is an instance of Key.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Key {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Key.__pulumiType;
    }

    /**
     * Settings for keys that can be used by Android apps.
     */
    public readonly androidSettings!: pulumi.Output<outputs.recaptchaenterprise.v1.GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsResponse>;
    /**
     * The timestamp corresponding to the creation of this Key.
     */
    public readonly createTime!: pulumi.Output<string>;
    /**
     * Human-readable display name of this key. Modifiable by user.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Settings for keys that can be used by iOS apps.
     */
    public readonly iosSettings!: pulumi.Output<outputs.recaptchaenterprise.v1.GoogleCloudRecaptchaenterpriseV1IOSKeySettingsResponse>;
    /**
     * See Creating and managing labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The resource name for the Key in the format "projects/{project}/keys/{key}".
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Options for user acceptance testing.
     */
    public readonly testingOptions!: pulumi.Output<outputs.recaptchaenterprise.v1.GoogleCloudRecaptchaenterpriseV1TestingOptionsResponse>;
    /**
     * Settings for keys that can be used by websites.
     */
    public readonly webSettings!: pulumi.Output<outputs.recaptchaenterprise.v1.GoogleCloudRecaptchaenterpriseV1WebKeySettingsResponse>;

    /**
     * Create a Key resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: KeyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["androidSettings"] = args ? args.androidSettings : undefined;
            inputs["createTime"] = args ? args.createTime : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["iosSettings"] = args ? args.iosSettings : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["testingOptions"] = args ? args.testingOptions : undefined;
            inputs["webSettings"] = args ? args.webSettings : undefined;
        } else {
            inputs["androidSettings"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["iosSettings"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["testingOptions"] = undefined /*out*/;
            inputs["webSettings"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Key.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Key resource.
 */
export interface KeyArgs {
    /**
     * Settings for keys that can be used by Android apps.
     */
    androidSettings?: pulumi.Input<inputs.recaptchaenterprise.v1.GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsArgs>;
    /**
     * The timestamp corresponding to the creation of this Key.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Human-readable display name of this key. Modifiable by user.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Settings for keys that can be used by iOS apps.
     */
    iosSettings?: pulumi.Input<inputs.recaptchaenterprise.v1.GoogleCloudRecaptchaenterpriseV1IOSKeySettingsArgs>;
    /**
     * See Creating and managing labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name for the Key in the format "projects/{project}/keys/{key}".
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Options for user acceptance testing.
     */
    testingOptions?: pulumi.Input<inputs.recaptchaenterprise.v1.GoogleCloudRecaptchaenterpriseV1TestingOptionsArgs>;
    /**
     * Settings for keys that can be used by websites.
     */
    webSettings?: pulumi.Input<inputs.recaptchaenterprise.v1.GoogleCloudRecaptchaenterpriseV1WebKeySettingsArgs>;
}