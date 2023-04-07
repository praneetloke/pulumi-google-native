// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the details of a Scope.
 */
export function getScope(args: GetScopeArgs, opts?: pulumi.InvokeOptions): Promise<GetScopeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:gkehub/v1beta:getScope", {
        "location": args.location,
        "project": args.project,
        "scopeId": args.scopeId,
    }, opts);
}

export interface GetScopeArgs {
    location: string;
    project?: string;
    scopeId: string;
}

export interface GetScopeResult {
    /**
     * When the scope was created.
     */
    readonly createTime: string;
    /**
     * When the scope was deleted.
     */
    readonly deleteTime: string;
    /**
     * The resource name for the scope `projects/{project}/locations/{location}/scopes/{scope}`
     */
    readonly name: string;
    /**
     * State of the scope resource.
     */
    readonly state: outputs.gkehub.v1beta.ScopeLifecycleStateResponse;
    /**
     * Google-generated UUID for this resource. This is unique across all scope resources. If a scope resource is deleted and another resource with the same name is created, it gets a different uid.
     */
    readonly uid: string;
    /**
     * When the scope was last updated.
     */
    readonly updateTime: string;
}
/**
 * Returns the details of a Scope.
 */
export function getScopeOutput(args: GetScopeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetScopeResult> {
    return pulumi.output(args).apply((a: any) => getScope(a, opts))
}

export interface GetScopeOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    scopeId: pulumi.Input<string>;
}
