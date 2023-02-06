// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets details of a single Instance.
 */
export function getInstance(args: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:apigeeregistry/v1:getInstance", {
        "instanceId": args.instanceId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetInstanceArgs {
    instanceId: string;
    location: string;
    project?: string;
}

export interface GetInstanceResult {
    /**
     * Build info of the Instance if it's in `ACTIVE` state.
     */
    readonly build: outputs.apigeeregistry.v1.BuildResponse;
    /**
     * Config of the Instance.
     */
    readonly config: outputs.apigeeregistry.v1.ConfigResponse;
    /**
     * Creation timestamp.
     */
    readonly createTime: string;
    /**
     * Format: `projects/*&#47;locations/*&#47;instance`. Currently only `locations/global` is supported.
     */
    readonly name: string;
    /**
     * The current state of the Instance.
     */
    readonly state: string;
    /**
     * Extra information of Instance.State if the state is `FAILED`.
     */
    readonly stateMessage: string;
    /**
     * Last update timestamp.
     */
    readonly updateTime: string;
}

export function getInstanceOutput(args: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    return pulumi.output(args).apply(a => getInstance(a, opts))
}

export interface GetInstanceOutputArgs {
    instanceId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
