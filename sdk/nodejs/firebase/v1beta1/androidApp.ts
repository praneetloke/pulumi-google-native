// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Requests the creation of a new AndroidApp in the specified FirebaseProject. The result of this call is an `Operation` which can be used to track the provisioning process. The `Operation` is automatically deleted after completion, so there is no need to call `DeleteOperation`.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class AndroidApp extends pulumi.CustomResource {
    /**
     * Get an existing AndroidApp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AndroidApp {
        return new AndroidApp(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:firebase/v1beta1:AndroidApp';

    /**
     * Returns true if the given object is an instance of AndroidApp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AndroidApp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AndroidApp.__pulumiType;
    }

    /**
     * The key_id of the GCP ApiKey associated with this App. If set must have no restrictions, or only have restrictions that are valid for the associated Firebase App. Cannot be set in create requests, instead an existing valid API Key will be chosen, or if no valid API Keys exist, one will be provisioned for you. Cannot be set to an empty value in update requests.
     */
    public readonly apiKeyId!: pulumi.Output<string>;
    /**
     * Immutable. The globally unique, Firebase-assigned identifier for the `AndroidApp`. This identifier should be treated as an opaque token, as the data format is not specified.
     */
    public /*out*/ readonly appId!: pulumi.Output<string>;
    /**
     * The user-assigned display name for the `AndroidApp`.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The resource name of the AndroidApp, in the format: projects/ PROJECT_IDENTIFIER/androidApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.androidApps#AndroidApp.FIELDS.app_id)).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Immutable. The canonical package name of the Android app as would appear in the Google Play Developer Console.
     */
    public readonly packageName!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The lifecycle state of the App.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a AndroidApp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AndroidAppArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["apiKeyId"] = args ? args.apiKeyId : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["packageName"] = args ? args.packageName : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["apiKeyId"] = undefined /*out*/;
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["packageName"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AndroidApp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AndroidApp resource.
 */
export interface AndroidAppArgs {
    /**
     * The key_id of the GCP ApiKey associated with this App. If set must have no restrictions, or only have restrictions that are valid for the associated Firebase App. Cannot be set in create requests, instead an existing valid API Key will be chosen, or if no valid API Keys exist, one will be provisioned for you. Cannot be set to an empty value in update requests.
     */
    apiKeyId?: pulumi.Input<string>;
    /**
     * The user-assigned display name for the `AndroidApp`.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The resource name of the AndroidApp, in the format: projects/ PROJECT_IDENTIFIER/androidApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.androidApps#AndroidApp.FIELDS.app_id)).
     */
    name?: pulumi.Input<string>;
    /**
     * Immutable. The canonical package name of the Android app as would appear in the Google Play Developer Console.
     */
    packageName?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
