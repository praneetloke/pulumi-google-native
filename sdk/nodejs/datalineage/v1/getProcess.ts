// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets the details of the specified process.
 */
export function getProcess(args: GetProcessArgs, opts?: pulumi.InvokeOptions): Promise<GetProcessResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:datalineage/v1:getProcess", {
        "location": args.location,
        "processId": args.processId,
        "project": args.project,
    }, opts);
}

export interface GetProcessArgs {
    location: string;
    processId: string;
    project?: string;
}

export interface GetProcessResult {
    /**
     * Optional. The attributes of the process. Should only be used for the purpose of non-semantic management (classifying, describing or labeling the process). Up to 100 attributes are allowed.
     */
    readonly attributes: {[key: string]: string};
    /**
     * Optional. A human-readable name you can set to display in a user interface. Must be not longer than 200 characters and only contain UTF-8 letters or numbers, spaces or characters like `_-:&.`
     */
    readonly displayName: string;
    /**
     * Immutable. The resource name of the lineage process. Format: `projects/{project}/locations/{location}/processes/{process}`. Can be specified or auto-assigned. {process} must be not longer than 200 characters and only contain characters in a set: `a-zA-Z0-9_-:.`
     */
    readonly name: string;
    /**
     * Optional. The origin of this process and its runs and lineage events.
     */
    readonly origin: outputs.datalineage.v1.GoogleCloudDatacatalogLineageV1OriginResponse;
}
/**
 * Gets the details of the specified process.
 */
export function getProcessOutput(args: GetProcessOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProcessResult> {
    return pulumi.output(args).apply((a: any) => getProcess(a, opts))
}

export interface GetProcessOutputArgs {
    location: pulumi.Input<string>;
    processId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}