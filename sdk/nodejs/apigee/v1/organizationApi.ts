// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an API proxy. The API proxy created will not be accessible at runtime until it is deployed to an environment. Create a new API proxy by setting the `name` query parameter to the name of the API proxy. Import an API proxy configuration bundle stored in zip format on your local machine to your organization by doing the following: * Set the `name` query parameter to the name of the API proxy. * Set the `action` query parameter to `import`. * Set the `Content-Type` header to `multipart/form-data`. * Pass as a file the name of API proxy configuration bundle stored in zip format on your local machine using the `file` form field. **Note**: To validate the API proxy configuration bundle only without importing it, set the `action` query parameter to `validate`. When importing an API proxy configuration bundle, if the API proxy does not exist, it will be created. If the API proxy exists, then a new revision is created. Invalid API proxy configurations are rejected, and a list of validation errors is returned to the client.
 */
export class OrganizationApi extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationApi resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OrganizationApi {
        return new OrganizationApi(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:OrganizationApi';

    /**
     * Returns true if the given object is an instance of OrganizationApi.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationApi {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationApi.__pulumiType;
    }

    /**
     * The id of the most recently created revision for this api proxy.
     */
    public /*out*/ readonly latestRevisionId!: pulumi.Output<string>;
    /**
     * Metadata describing the API proxy.
     */
    public /*out*/ readonly metaData!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1EntityMetadataResponse>;
    /**
     * Name of the API proxy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of revisons defined for the API proxy.
     */
    public /*out*/ readonly revision!: pulumi.Output<string[]>;

    /**
     * Create a OrganizationApi resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationApiArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.apiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiId'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            inputs["action"] = args ? args.action : undefined;
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["contentType"] = args ? args.contentType : undefined;
            inputs["data"] = args ? args.data : undefined;
            inputs["extensions"] = args ? args.extensions : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["validate"] = args ? args.validate : undefined;
            inputs["latestRevisionId"] = undefined /*out*/;
            inputs["metaData"] = undefined /*out*/;
            inputs["revision"] = undefined /*out*/;
        } else {
            inputs["latestRevisionId"] = undefined /*out*/;
            inputs["metaData"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["revision"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OrganizationApi.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a OrganizationApi resource.
 */
export interface OrganizationApiArgs {
    readonly action?: pulumi.Input<string>;
    readonly apiId: pulumi.Input<string>;
    /**
     * The HTTP Content-Type header value specifying the content type of the body.
     */
    readonly contentType?: pulumi.Input<string>;
    /**
     * The HTTP request/response body as raw binary.
     */
    readonly data?: pulumi.Input<string>;
    /**
     * Application specific response metadata. Must be set in the first response for streaming APIs.
     */
    readonly extensions?: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    readonly name?: pulumi.Input<string>;
    readonly organizationId: pulumi.Input<string>;
    readonly validate?: pulumi.Input<string>;
}
