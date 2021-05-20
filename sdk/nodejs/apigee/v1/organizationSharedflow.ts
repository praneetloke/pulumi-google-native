// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Uploads a ZIP-formatted shared flow configuration bundle to an organization. If the shared flow already exists, this creates a new revision of it. If the shared flow does not exist, this creates it. Once imported, the shared flow revision must be deployed before it can be accessed at runtime. The size limit of a shared flow bundle is 15 MB.
 */
export class OrganizationSharedflow extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationSharedflow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OrganizationSharedflow {
        return new OrganizationSharedflow(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:OrganizationSharedflow';

    /**
     * Returns true if the given object is an instance of OrganizationSharedflow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationSharedflow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationSharedflow.__pulumiType;
    }

    /**
     * The id of the most recently created revision for this shared flow.
     */
    public /*out*/ readonly latestRevisionId!: pulumi.Output<string>;
    /**
     * Metadata describing the shared flow.
     */
    public /*out*/ readonly metaData!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1EntityMetadataResponse>;
    /**
     * The ID of the shared flow.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of revisions of this shared flow.
     */
    public /*out*/ readonly revision!: pulumi.Output<string[]>;

    /**
     * Create a OrganizationSharedflow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationSharedflowArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            if ((!args || args.sharedflowId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sharedflowId'");
            }
            inputs["action"] = args ? args.action : undefined;
            inputs["contentType"] = args ? args.contentType : undefined;
            inputs["data"] = args ? args.data : undefined;
            inputs["extensions"] = args ? args.extensions : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["sharedflowId"] = args ? args.sharedflowId : undefined;
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
        super(OrganizationSharedflow.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a OrganizationSharedflow resource.
 */
export interface OrganizationSharedflowArgs {
    readonly action: pulumi.Input<string>;
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
    readonly name: pulumi.Input<string>;
    readonly organizationId: pulumi.Input<string>;
    readonly sharedflowId: pulumi.Input<string>;
}
