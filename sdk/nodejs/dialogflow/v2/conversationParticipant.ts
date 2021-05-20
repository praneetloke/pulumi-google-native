// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new participant in a conversation.
 */
export class ConversationParticipant extends pulumi.CustomResource {
    /**
     * Get an existing ConversationParticipant resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConversationParticipant {
        return new ConversationParticipant(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v2:ConversationParticipant';

    /**
     * Returns true if the given object is an instance of ConversationParticipant.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConversationParticipant {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConversationParticipant.__pulumiType;
    }

    /**
     * Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
     */
    public readonly sipRecordingMediaLabel!: pulumi.Output<string>;

    /**
     * Create a ConversationParticipant resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConversationParticipantArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.conversationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'conversationId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.participantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'participantId'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["conversationId"] = args ? args.conversationId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["participantId"] = args ? args.participantId : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["sipRecordingMediaLabel"] = args ? args.sipRecordingMediaLabel : undefined;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["role"] = undefined /*out*/;
            inputs["sipRecordingMediaLabel"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ConversationParticipant.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConversationParticipant resource.
 */
export interface ConversationParticipantArgs {
    readonly conversationId: pulumi.Input<string>;
    readonly location: pulumi.Input<string>;
    /**
     * Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
     */
    readonly name?: pulumi.Input<string>;
    readonly participantId: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
     */
    readonly sipRecordingMediaLabel?: pulumi.Input<string>;
}
