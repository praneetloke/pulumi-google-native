// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new Api in a given project and location.
 * Auto-naming is currently not supported for this resource.
 */
export class Api extends pulumi.CustomResource {
    /**
     * Get an existing Api resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Api {
        return new Api(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigateway/v1:Api';

    /**
     * Returns true if the given object is an instance of Api.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Api {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Api.__pulumiType;
    }

    /**
     * Created time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. Display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Optional. Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed). If not specified, a new Service will automatically be created in the same project as this API.
     */
    public readonly managedService!: pulumi.Output<string>;
    /**
     * Resource name of the API. Format: projects/{project}/locations/global/apis/{api}
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * State of the API.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Updated time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Api resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.apiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiId'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managedService"] = args ? args.managedService : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["createTime"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["managedService"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Api.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Api resource.
 */
export interface ApiArgs {
    apiId: pulumi.Input<string>;
    /**
     * Optional. Display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Optional. Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed). If not specified, a new Service will automatically be created in the same project as this API.
     */
    managedService?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
