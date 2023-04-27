// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new AppConnection in a given project and location.
 */
export class AppConnection extends pulumi.CustomResource {
    /**
     * Get an existing AppConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AppConnection {
        return new AppConnection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:beyondcorp/v1alpha:AppConnection';

    /**
     * Returns true if the given object is an instance of AppConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppConnection.__pulumiType;
    }

    /**
     * Optional. User-settable AppConnection resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
     */
    public readonly appConnectionId!: pulumi.Output<string | undefined>;
    /**
     * Address of the remote application endpoint for the BeyondCorp AppConnection.
     */
    public readonly applicationEndpoint!: pulumi.Output<outputs.beyondcorp.v1alpha.GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionApplicationEndpointResponse>;
    /**
     * Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this AppConnection.
     */
    public readonly connectors!: pulumi.Output<string[]>;
    /**
     * Timestamp when the resource was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. An arbitrary user-provided name for the AppConnection. Cannot exceed 64 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Optional. Gateway used by the AppConnection.
     */
    public readonly gateway!: pulumi.Output<outputs.beyondcorp.v1alpha.GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionGatewayResponse>;
    /**
     * Optional. Resource labels to represent user provided metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Unique resource name of the AppConnection. The name is ignored when creating a AppConnection.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * The current state of the AppConnection.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The type of network connectivity used by the AppConnection.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * A unique identifier for the instance generated by the system.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * Timestamp when the resource was last modified.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a AppConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.applicationEndpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationEndpoint'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["appConnectionId"] = args ? args.appConnectionId : undefined;
            resourceInputs["applicationEndpoint"] = args ? args.applicationEndpoint : undefined;
            resourceInputs["connectors"] = args ? args.connectors : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["gateway"] = args ? args.gateway : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["appConnectionId"] = undefined /*out*/;
            resourceInputs["applicationEndpoint"] = undefined /*out*/;
            resourceInputs["connectors"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["gateway"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(AppConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AppConnection resource.
 */
export interface AppConnectionArgs {
    /**
     * Optional. User-settable AppConnection resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
     */
    appConnectionId?: pulumi.Input<string>;
    /**
     * Address of the remote application endpoint for the BeyondCorp AppConnection.
     */
    applicationEndpoint: pulumi.Input<inputs.beyondcorp.v1alpha.GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionApplicationEndpointArgs>;
    /**
     * Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this AppConnection.
     */
    connectors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional. An arbitrary user-provided name for the AppConnection. Cannot exceed 64 characters.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Optional. Gateway used by the AppConnection.
     */
    gateway?: pulumi.Input<inputs.beyondcorp.v1alpha.GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionGatewayArgs>;
    /**
     * Optional. Resource labels to represent user provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Unique resource name of the AppConnection. The name is ignored when creating a AppConnection.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * The type of network connectivity used by the AppConnection.
     */
    type: pulumi.Input<enums.beyondcorp.v1alpha.AppConnectionType>;
}
