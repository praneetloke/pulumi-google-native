// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Fetches the representation of an existing ResourceRecordSet.
 */
export function getResourceRecordSet(args: GetResourceRecordSetArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceRecordSetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:dns/v1:getResourceRecordSet", {
        "clientOperationId": args.clientOperationId,
        "managedZone": args.managedZone,
        "name": args.name,
        "project": args.project,
        "type": args.type,
    }, opts);
}

export interface GetResourceRecordSetArgs {
    clientOperationId?: string;
    managedZone: string;
    name: string;
    project?: string;
    type: string;
}

export interface GetResourceRecordSetResult {
    readonly kind: string;
    /**
     * For example, www.example.com.
     */
    readonly name: string;
    /**
     * Configures dynamic query responses based on geo location of querying user or a weighted round robin based routing policy. A ResourceRecordSet should only have either rrdata (static) or routing_policy (dynamic). An error is returned otherwise.
     */
    readonly routingPolicy: outputs.dns.v1.RRSetRoutingPolicyResponse;
    /**
     * As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
     */
    readonly rrdatas: string[];
    /**
     * As defined in RFC 4034 (section 3.2).
     */
    readonly signatureRrdatas: string[];
    /**
     * Number of seconds that this ResourceRecordSet can be cached by resolvers.
     */
    readonly ttl: number;
    /**
     * The identifier of a supported record type. See the list of Supported DNS record types.
     */
    readonly type: string;
}

export function getResourceRecordSetOutput(args: GetResourceRecordSetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourceRecordSetResult> {
    return pulumi.output(args).apply(a => getResourceRecordSet(a, opts))
}

export interface GetResourceRecordSetOutputArgs {
    clientOperationId?: pulumi.Input<string>;
    managedZone: pulumi.Input<string>;
    name: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    type: pulumi.Input<string>;
}
