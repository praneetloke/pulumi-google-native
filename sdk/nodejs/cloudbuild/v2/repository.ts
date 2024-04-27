// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a Repository.
 */
export class Repository extends pulumi.CustomResource {
    /**
     * Get an existing Repository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Repository {
        return new Repository(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudbuild/v2:Repository';

    /**
     * Returns true if the given object is an instance of Repository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Repository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Repository.__pulumiType;
    }

    /**
     * Allows clients to store small amounts of arbitrary data.
     */
    public readonly annotations!: pulumi.Output<{[key: string]: string}>;
    public readonly connectionId!: pulumi.Output<string>;
    /**
     * Server assigned timestamp for when the connection was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     */
    public readonly etag!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Immutable. Resource name of the repository, in the format `projects/*&#47;locations/*&#47;connections/*&#47;repositories/*`.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Git Clone HTTPS URI.
     */
    public readonly remoteUri!: pulumi.Output<string>;
    /**
     * Required. The ID to use for the repository, which will become the final component of the repository's resource name. This ID should be unique in the connection. Allows alphanumeric characters and any of -._~%!$&'()*+,;=@.
     */
    public readonly repositoryId!: pulumi.Output<string>;
    /**
     * Server assigned timestamp for when the connection was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * External ID of the webhook created for the repository.
     */
    public /*out*/ readonly webhookId!: pulumi.Output<string>;

    /**
     * Create a Repository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionId'");
            }
            if ((!args || args.remoteUri === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remoteUri'");
            }
            if ((!args || args.repositoryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repositoryId'");
            }
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["connectionId"] = args ? args.connectionId : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["remoteUri"] = args ? args.remoteUri : undefined;
            resourceInputs["repositoryId"] = args ? args.repositoryId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["webhookId"] = undefined /*out*/;
        } else {
            resourceInputs["annotations"] = undefined /*out*/;
            resourceInputs["connectionId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["remoteUri"] = undefined /*out*/;
            resourceInputs["repositoryId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["webhookId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["connectionId", "location", "project", "repositoryId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Repository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Repository resource.
 */
export interface RepositoryArgs {
    /**
     * Allows clients to store small amounts of arbitrary data.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    connectionId: pulumi.Input<string>;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     */
    etag?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Immutable. Resource name of the repository, in the format `projects/*&#47;locations/*&#47;connections/*&#47;repositories/*`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Git Clone HTTPS URI.
     */
    remoteUri: pulumi.Input<string>;
    /**
     * Required. The ID to use for the repository, which will become the final component of the repository's resource name. This ID should be unique in the connection. Allows alphanumeric characters and any of -._~%!$&'()*+,;=@.
     */
    repositoryId: pulumi.Input<string>;
}