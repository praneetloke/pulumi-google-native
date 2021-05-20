// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a resource file. Specify the `Content-Type` as `application/octet-stream` or `multipart/form-data`. For more information about resource files, see [Resource files](https://cloud.google.com/apigee/docs/api-platform/develop/resource-files).
 */
export class OrganizationEnvironmentResourcefile extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationEnvironmentResourcefile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OrganizationEnvironmentResourcefile {
        return new OrganizationEnvironmentResourcefile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:OrganizationEnvironmentResourcefile';

    /**
     * Returns true if the given object is an instance of OrganizationEnvironmentResourcefile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationEnvironmentResourcefile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationEnvironmentResourcefile.__pulumiType;
    }

    /**
     * The HTTP Content-Type header value specifying the content type of the body.
     */
    public readonly contentType!: pulumi.Output<string>;
    /**
     * The HTTP request/response body as raw binary.
     */
    public readonly data!: pulumi.Output<string>;
    /**
     * Application specific response metadata. Must be set in the first response for streaming APIs.
     */
    public readonly extensions!: pulumi.Output<{[key: string]: string}[]>;

    /**
     * Create a OrganizationEnvironmentResourcefile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationEnvironmentResourcefileArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["contentType"] = args ? args.contentType : undefined;
            inputs["data"] = args ? args.data : undefined;
            inputs["environmentId"] = args ? args.environmentId : undefined;
            inputs["extensions"] = args ? args.extensions : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["type"] = args ? args.type : undefined;
        } else {
            inputs["contentType"] = undefined /*out*/;
            inputs["data"] = undefined /*out*/;
            inputs["extensions"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OrganizationEnvironmentResourcefile.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a OrganizationEnvironmentResourcefile resource.
 */
export interface OrganizationEnvironmentResourcefileArgs {
    /**
     * The HTTP Content-Type header value specifying the content type of the body.
     */
    readonly contentType?: pulumi.Input<string>;
    /**
     * The HTTP request/response body as raw binary.
     */
    readonly data?: pulumi.Input<string>;
    readonly environmentId: pulumi.Input<string>;
    /**
     * Application specific response metadata. Must be set in the first response for streaming APIs.
     */
    readonly extensions?: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    readonly name: pulumi.Input<string>;
    readonly organizationId: pulumi.Input<string>;
    readonly type: pulumi.Input<string>;
}
