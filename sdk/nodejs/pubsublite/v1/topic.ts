// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new topic.
 */
export class Topic extends pulumi.CustomResource {
    /**
     * Get an existing Topic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Topic {
        return new Topic(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:pubsublite/v1:Topic';

    /**
     * Returns true if the given object is an instance of Topic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Topic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Topic.__pulumiType;
    }

    /**
     * The name of the topic. Structured like: projects/{project_number}/locations/{location}/topics/{topic_id}
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The settings for this topic's partitions.
     */
    public readonly partitionConfig!: pulumi.Output<outputs.pubsublite.v1.PartitionConfigResponse>;
    /**
     * The settings for this topic's message retention.
     */
    public readonly retentionConfig!: pulumi.Output<outputs.pubsublite.v1.RetentionConfigResponse>;

    /**
     * Create a Topic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TopicArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.topicId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicId'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["partitionConfig"] = args ? args.partitionConfig : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["retentionConfig"] = args ? args.retentionConfig : undefined;
            inputs["topicId"] = args ? args.topicId : undefined;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["partitionConfig"] = undefined /*out*/;
            inputs["retentionConfig"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Topic.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Topic resource.
 */
export interface TopicArgs {
    location?: pulumi.Input<string>;
    /**
     * The name of the topic. Structured like: projects/{project_number}/locations/{location}/topics/{topic_id}
     */
    name?: pulumi.Input<string>;
    /**
     * The settings for this topic's partitions.
     */
    partitionConfig?: pulumi.Input<inputs.pubsublite.v1.PartitionConfigArgs>;
    project?: pulumi.Input<string>;
    /**
     * The settings for this topic's message retention.
     */
    retentionConfig?: pulumi.Input<inputs.pubsublite.v1.RetentionConfigArgs>;
    topicId: pulumi.Input<string>;
}
