// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a page in the specified flow. Note: You should always train a flow prior to sending it queries. See the [training documentation](https://cloud.google.com/dialogflow/cx/docs/concept/training).
 */
export class Page extends pulumi.CustomResource {
    /**
     * Get an existing Page resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Page {
        return new Page(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v3:Page';

    /**
     * Returns true if the given object is an instance of Page.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Page {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Page.__pulumiType;
    }

    /**
     * The human-readable name of the page, unique within the agent.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The fulfillment to call when the session is entering the page.
     */
    public readonly entryFulfillment!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3FulfillmentResponse>;
    /**
     * Handlers associated with the page to handle events such as webhook errors, no match or no input.
     */
    public readonly eventHandlers!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3EventHandlerResponse[]>;
    /**
     * The form associated with the page, used for collecting parameters relevant to the page.
     */
    public readonly form!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3FormResponse>;
    /**
     * The unique identifier of the page. Required for the Pages.UpdatePage method. Pages.CreatePage populates the name automatically. Format: `projects//locations//agents//flows//pages/`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Ordered list of `TransitionRouteGroups` associated with the page. Transition route groups must be unique within a page. * If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes. * If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
     */
    public readonly transitionRouteGroups!: pulumi.Output<string[]>;
    /**
     * A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order: * TransitionRoutes defined in the page with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in flow with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in the page with only condition specified. * TransitionRoutes defined in the transition route groups with only condition specified.
     */
    public readonly transitionRoutes!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3TransitionRouteResponse[]>;

    /**
     * Create a Page resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PageArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.agentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentId'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.flowId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flowId'");
            }
            inputs["agentId"] = args ? args.agentId : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["entryFulfillment"] = args ? args.entryFulfillment : undefined;
            inputs["eventHandlers"] = args ? args.eventHandlers : undefined;
            inputs["flowId"] = args ? args.flowId : undefined;
            inputs["form"] = args ? args.form : undefined;
            inputs["languageCode"] = args ? args.languageCode : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["transitionRouteGroups"] = args ? args.transitionRouteGroups : undefined;
            inputs["transitionRoutes"] = args ? args.transitionRoutes : undefined;
        } else {
            inputs["displayName"] = undefined /*out*/;
            inputs["entryFulfillment"] = undefined /*out*/;
            inputs["eventHandlers"] = undefined /*out*/;
            inputs["form"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["transitionRouteGroups"] = undefined /*out*/;
            inputs["transitionRoutes"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Page.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Page resource.
 */
export interface PageArgs {
    agentId: pulumi.Input<string>;
    /**
     * The human-readable name of the page, unique within the agent.
     */
    displayName: pulumi.Input<string>;
    /**
     * The fulfillment to call when the session is entering the page.
     */
    entryFulfillment?: pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3FulfillmentArgs>;
    /**
     * Handlers associated with the page to handle events such as webhook errors, no match or no input.
     */
    eventHandlers?: pulumi.Input<pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3EventHandlerArgs>[]>;
    flowId: pulumi.Input<string>;
    /**
     * The form associated with the page, used for collecting parameters relevant to the page.
     */
    form?: pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3FormArgs>;
    languageCode?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * The unique identifier of the page. Required for the Pages.UpdatePage method. Pages.CreatePage populates the name automatically. Format: `projects//locations//agents//flows//pages/`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Ordered list of `TransitionRouteGroups` associated with the page. Transition route groups must be unique within a page. * If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes. * If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
     */
    transitionRouteGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order: * TransitionRoutes defined in the page with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in flow with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in the page with only condition specified. * TransitionRoutes defined in the transition route groups with only condition specified.
     */
    transitionRoutes?: pulumi.Input<pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3TransitionRouteArgs>[]>;
}
