// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates and starts a Replay using the given ReplayConfig.
 * Auto-naming is currently not supported for this resource.
 */
export class Replay extends pulumi.CustomResource {
    /**
     * Get an existing Replay resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Replay {
        return new Replay(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:policysimulator/v1beta1:Replay';

    /**
     * Returns true if the given object is an instance of Replay.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Replay {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Replay.__pulumiType;
    }

    /**
     * The configuration used for the `Replay`.
     */
    public readonly config!: pulumi.Output<outputs.policysimulator.v1beta1.GoogleCloudPolicysimulatorV1beta1ReplayConfigResponse>;
    /**
     * The resource name of the `Replay`, which has the following format: `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}`, where `{resource-id}` is the ID of the project, folder, or organization that owns the Replay. Example: `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36`
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Summary statistics about the replayed log entries.
     */
    public /*out*/ readonly resultsSummary!: pulumi.Output<outputs.policysimulator.v1beta1.GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponse>;
    /**
     * The current state of the `Replay`.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a Replay resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplayArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            inputs["config"] = args ? args.config : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["resultsSummary"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        } else {
            inputs["config"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["resultsSummary"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Replay.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Replay resource.
 */
export interface ReplayArgs {
    /**
     * The configuration used for the `Replay`.
     */
    config: pulumi.Input<inputs.policysimulator.v1beta1.GoogleCloudPolicysimulatorV1beta1ReplayConfigArgs>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
