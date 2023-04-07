// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Returns the specified TargetInstance resource.
 */
export function getTargetInstance(args: GetTargetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetTargetInstanceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/v1:getTargetInstance", {
        "project": args.project,
        "targetInstance": args.targetInstance,
        "zone": args.zone,
    }, opts);
}

export interface GetTargetInstanceArgs {
    project?: string;
    targetInstance: string;
    zone: string;
}

export interface GetTargetInstanceResult {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * A URL to the virtual machine instance that handles traffic for this target instance. When creating a target instance, you can provide the fully-qualified URL or a valid partial URL to the desired virtual machine. For example, the following are all valid URLs: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance - zones/zone/instances/instance 
     */
    readonly instance: string;
    /**
     * The type of the resource. Always compute#targetInstance for target instances.
     */
    readonly kind: string;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * Must have a value of NO_NAT. Protocol forwarding delivers packets while preserving the destination IP address of the forwarding rule referencing the target instance.
     */
    readonly natPolicy: string;
    /**
     * The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
     */
    readonly network: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * URL of the zone where the target instance resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly zone: string;
}
/**
 * Returns the specified TargetInstance resource.
 */
export function getTargetInstanceOutput(args: GetTargetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTargetInstanceResult> {
    return pulumi.output(args).apply((a: any) => getTargetInstance(a, opts))
}

export interface GetTargetInstanceOutputArgs {
    project?: pulumi.Input<string>;
    targetInstance: pulumi.Input<string>;
    zone: pulumi.Input<string>;
}
