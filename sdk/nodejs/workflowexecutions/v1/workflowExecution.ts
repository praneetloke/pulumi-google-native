// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new execution using the latest revision of the given workflow.
 */
export class WorkflowExecution extends pulumi.CustomResource {
    /**
     * Get an existing WorkflowExecution resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WorkflowExecution {
        return new WorkflowExecution(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:workflowexecutions/v1:WorkflowExecution';

    /**
     * Returns true if the given object is an instance of WorkflowExecution.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkflowExecution {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkflowExecution.__pulumiType;
    }

    /**
     * Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
     */
    public readonly argument!: pulumi.Output<string>;
    /**
     * Marks the end of execution, successful or not.
     */
    public /*out*/ readonly endTime!: pulumi.Output<string>;
    /**
     * The error which caused the execution to finish prematurely. The value is only present if the execution's state is `FAILED` or `CANCELLED`.
     */
    public /*out*/ readonly error!: pulumi.Output<outputs.workflowexecutions.v1.ErrorResponse>;
    /**
     * The resource name of the execution. Format: projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Output of the execution represented as a JSON string. The value can only be present if the execution's state is `SUCCEEDED`.
     */
    public /*out*/ readonly result!: pulumi.Output<string>;
    /**
     * Marks the beginning of execution.
     */
    public /*out*/ readonly startTime!: pulumi.Output<string>;
    /**
     * Current state of the execution.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Revision of the workflow this execution is using.
     */
    public /*out*/ readonly workflowRevisionId!: pulumi.Output<string>;

    /**
     * Create a WorkflowExecution resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkflowExecutionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.executionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'executionId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.workflowId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workflowId'");
            }
            inputs["argument"] = args ? args.argument : undefined;
            inputs["executionId"] = args ? args.executionId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["workflowId"] = args ? args.workflowId : undefined;
            inputs["endTime"] = undefined /*out*/;
            inputs["error"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["result"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["workflowRevisionId"] = undefined /*out*/;
        } else {
            inputs["argument"] = undefined /*out*/;
            inputs["endTime"] = undefined /*out*/;
            inputs["error"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["result"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["workflowRevisionId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WorkflowExecution.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WorkflowExecution resource.
 */
export interface WorkflowExecutionArgs {
    /**
     * Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
     */
    readonly argument?: pulumi.Input<string>;
    readonly executionId: pulumi.Input<string>;
    readonly location: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    readonly workflowId: pulumi.Input<string>;
}
