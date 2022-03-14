// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new channel in the specified site.
 */
export class Channel extends pulumi.CustomResource {
    /**
     * Get an existing Channel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Channel {
        return new Channel(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:firebasehosting/v1beta1:Channel';

    /**
     * Returns true if the given object is an instance of Channel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Channel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Channel.__pulumiType;
    }

    /**
     * The time at which the channel was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted. This field is present in the output whether it's set directly or via the `ttl` field.
     */
    public readonly expireTime!: pulumi.Output<string>;
    /**
     * Text labels used for extra metadata and/or filtering.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The fully-qualified resource name for the channel, in the format: sites/ SITE_ID/channels/CHANNEL_ID
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The current release for the channel, if any.
     */
    public /*out*/ readonly release!: pulumi.Output<outputs.firebasehosting.v1beta1.ReleaseResponse>;
    /**
     * The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100. Defaults to 10 for new channels.
     */
    public readonly retainedReleaseCount!: pulumi.Output<number>;
    /**
     * Input only. A time-to-live for this channel. Sets `expire_time` to the provided duration past the time of the request.
     */
    public readonly ttl!: pulumi.Output<string>;
    /**
     * The time at which the channel was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The URL at which the content of this channel's current release can be viewed. This URL is a Firebase-provided subdomain of `web.app`. The content of this channel's current release can also be viewed at the Firebase-provided subdomain of `firebaseapp.com`. If this channel is the `live` channel for the Hosting site, then the content of this channel's current release can also be viewed at any connected custom domains.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a Channel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ChannelArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.channelId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'channelId'");
            }
            if ((!args || args.siteId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'siteId'");
            }
            resourceInputs["channelId"] = args ? args.channelId : undefined;
            resourceInputs["expireTime"] = args ? args.expireTime : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["retainedReleaseCount"] = args ? args.retainedReleaseCount : undefined;
            resourceInputs["siteId"] = args ? args.siteId : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["release"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["expireTime"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["release"] = undefined /*out*/;
            resourceInputs["retainedReleaseCount"] = undefined /*out*/;
            resourceInputs["ttl"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Channel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Channel resource.
 */
export interface ChannelArgs {
    /**
     * Required. Immutable. A unique ID within the site that identifies the channel.
     */
    channelId: pulumi.Input<string>;
    /**
     * The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted. This field is present in the output whether it's set directly or via the `ttl` field.
     */
    expireTime?: pulumi.Input<string>;
    /**
     * Text labels used for extra metadata and/or filtering.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The fully-qualified resource name for the channel, in the format: sites/ SITE_ID/channels/CHANNEL_ID
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100. Defaults to 10 for new channels.
     */
    retainedReleaseCount?: pulumi.Input<number>;
    siteId: pulumi.Input<string>;
    /**
     * Input only. A time-to-live for this channel. Sets `expire_time` to the provided duration past the time of the request.
     */
    ttl?: pulumi.Input<string>;
}
