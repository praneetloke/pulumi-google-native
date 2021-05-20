// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates an agent version. The new version points to the agent instance in the "default" environment.
 */
export class AgentVersion extends pulumi.CustomResource {
    /**
     * Get an existing AgentVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AgentVersion {
        return new AgentVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v2beta1:AgentVersion';

    /**
     * Returns true if the given object is an instance of AgentVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AgentVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AgentVersion.__pulumiType;
    }

    /**
     * The creation time of this version. This field is read-only, i.e., it cannot be set by create and update methods.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. The developer-provided description of this version.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The unique identifier of this agent version. Supported formats: - `projects//agent/versions/` - `projects//locations//agent/versions/`
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The status of this version. This field is read-only and cannot be set by create and update methods.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The sequential number of this version. This field is read-only which means it cannot be set by create and update methods.
     */
    public /*out*/ readonly versionNumber!: pulumi.Output<number>;

    /**
     * Create a AgentVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AgentVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.versionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'versionId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["versionId"] = args ? args.versionId : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["versionNumber"] = undefined /*out*/;
        } else {
            inputs["createTime"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["versionNumber"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AgentVersion.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AgentVersion resource.
 */
export interface AgentVersionArgs {
    /**
     * Optional. The developer-provided description of this version.
     */
    readonly description?: pulumi.Input<string>;
    readonly location: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    readonly versionId: pulumi.Input<string>;
}
