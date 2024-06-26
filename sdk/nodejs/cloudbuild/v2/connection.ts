// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a Connection.
 */
export class Connection extends pulumi.CustomResource {
    /**
     * Get an existing Connection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Connection {
        return new Connection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudbuild/v2:Connection';

    /**
     * Returns true if the given object is an instance of Connection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connection.__pulumiType;
    }

    /**
     * Allows clients to store small amounts of arbitrary data.
     */
    public readonly annotations!: pulumi.Output<{[key: string]: string}>;
    /**
     * Required. The ID to use for the Connection, which will become the final component of the Connection's resource name. Names must be unique per-project per-location. Allows alphanumeric characters and any of -._~%!$&'()*+,;=@.
     */
    public readonly connectionId!: pulumi.Output<string>;
    /**
     * Server assigned timestamp for when the connection was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.
     */
    public readonly disabled!: pulumi.Output<boolean>;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * Configuration for connections to github.com.
     */
    public readonly githubConfig!: pulumi.Output<outputs.cloudbuild.v2.GitHubConfigResponse>;
    /**
     * Configuration for connections to an instance of GitHub Enterprise.
     */
    public readonly githubEnterpriseConfig!: pulumi.Output<outputs.cloudbuild.v2.GoogleDevtoolsCloudbuildV2GitHubEnterpriseConfigResponse>;
    /**
     * Configuration for connections to gitlab.com or an instance of GitLab Enterprise.
     */
    public readonly gitlabConfig!: pulumi.Output<outputs.cloudbuild.v2.GoogleDevtoolsCloudbuildV2GitLabConfigResponse>;
    /**
     * Installation state of the Connection.
     */
    public /*out*/ readonly installationState!: pulumi.Output<outputs.cloudbuild.v2.InstallationStateResponse>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Immutable. The resource name of the connection, in the format `projects/{project}/locations/{location}/connections/{connection_id}`.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Set to true when the connection is being set up or updated in the background.
     */
    public /*out*/ readonly reconciling!: pulumi.Output<boolean>;
    /**
     * Server assigned timestamp for when the connection was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Connection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionId'");
            }
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["connectionId"] = args ? args.connectionId : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["githubConfig"] = args ? args.githubConfig : undefined;
            resourceInputs["githubEnterpriseConfig"] = args ? args.githubEnterpriseConfig : undefined;
            resourceInputs["gitlabConfig"] = args ? args.gitlabConfig : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["installationState"] = undefined /*out*/;
            resourceInputs["reconciling"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["annotations"] = undefined /*out*/;
            resourceInputs["connectionId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["disabled"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["githubConfig"] = undefined /*out*/;
            resourceInputs["githubEnterpriseConfig"] = undefined /*out*/;
            resourceInputs["gitlabConfig"] = undefined /*out*/;
            resourceInputs["installationState"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["reconciling"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["connectionId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Connection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Connection resource.
 */
export interface ConnectionArgs {
    /**
     * Allows clients to store small amounts of arbitrary data.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Required. The ID to use for the Connection, which will become the final component of the Connection's resource name. Names must be unique per-project per-location. Allows alphanumeric characters and any of -._~%!$&'()*+,;=@.
     */
    connectionId: pulumi.Input<string>;
    /**
     * If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     */
    etag?: pulumi.Input<string>;
    /**
     * Configuration for connections to github.com.
     */
    githubConfig?: pulumi.Input<inputs.cloudbuild.v2.GitHubConfigArgs>;
    /**
     * Configuration for connections to an instance of GitHub Enterprise.
     */
    githubEnterpriseConfig?: pulumi.Input<inputs.cloudbuild.v2.GoogleDevtoolsCloudbuildV2GitHubEnterpriseConfigArgs>;
    /**
     * Configuration for connections to gitlab.com or an instance of GitLab Enterprise.
     */
    gitlabConfig?: pulumi.Input<inputs.cloudbuild.v2.GoogleDevtoolsCloudbuildV2GitLabConfigArgs>;
    location?: pulumi.Input<string>;
    /**
     * Immutable. The resource name of the connection, in the format `projects/{project}/locations/{location}/connections/{connection_id}`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
