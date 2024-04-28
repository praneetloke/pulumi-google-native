// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified autoscaler.
 */
export function getRegionAutoscaler(args: GetRegionAutoscalerArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionAutoscalerResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/alpha:getRegionAutoscaler", {
        "autoscaler": args.autoscaler,
        "project": args.project,
        "region": args.region,
    }, opts);
}

export interface GetRegionAutoscalerArgs {
    autoscaler: string;
    project?: string;
    region: string;
}

export interface GetRegionAutoscalerResult {
    /**
     * The configuration parameters for the autoscaling algorithm. You can define one or more signals for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization. If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
     */
    readonly autoscalingPolicy: outputs.compute.alpha.AutoscalingPolicyResponse;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Type of the resource. Always compute#autoscaler for autoscalers.
     */
    readonly kind: string;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * Target recommended MIG size (number of instances) computed by autoscaler. Autoscaler calculates the recommended MIG size even when the autoscaling policy mode is different from ON. This field is empty when autoscaler is not connected to an existing managed instance group or autoscaler did not generate its prediction.
     */
    readonly recommendedSize: number;
    /**
     * URL of the region where the instance group resides (for autoscalers living in regional scope).
     */
    readonly region: string;
    /**
     * Status information of existing scaling schedules.
     */
    readonly scalingScheduleStatus: outputs.compute.alpha.ScalingScheduleStatusResponse;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    readonly selfLinkWithId: string;
    /**
     * The status of the autoscaler configuration. Current set of possible values: - PENDING: Autoscaler backend hasn't read new/updated configuration. - DELETING: Configuration is being deleted. - ACTIVE: Configuration is acknowledged to be effective. Some warnings might be present in the statusDetails field. - ERROR: Configuration has errors. Actionable for users. Details are present in the statusDetails field. New values might be added in the future.
     */
    readonly status: string;
    /**
     * Human-readable details about the current state of the autoscaler. Read the documentation for Commonly returned status messages for examples of status messages you might encounter.
     */
    readonly statusDetails: outputs.compute.alpha.AutoscalerStatusDetailsResponse[];
    /**
     * URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
     */
    readonly target: string;
    /**
     * URL of the zone where the instance group resides (for autoscalers living in zonal scope).
     */
    readonly zone: string;
}
/**
 * Returns the specified autoscaler.
 */
export function getRegionAutoscalerOutput(args: GetRegionAutoscalerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionAutoscalerResult> {
    return pulumi.output(args).apply((a: any) => getRegionAutoscaler(a, opts))
}

export interface GetRegionAutoscalerOutputArgs {
    autoscaler: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
}
