// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets information about a particular instance.
 */
export function getInstance(args: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:spanner/v1:getInstance", {
        "fieldMask": args.fieldMask,
        "instanceId": args.instanceId,
        "project": args.project,
    }, opts);
}

export interface GetInstanceArgs {
    fieldMask?: string;
    instanceId: string;
    project?: string;
}

export interface GetInstanceResult {
    /**
     * The name of the instance's configuration. Values are of the form `projects//instanceConfigs/`. See also InstanceConfig and ListInstanceConfigs.
     */
    readonly config: string;
    /**
     * The time at which the instance was created.
     */
    readonly createTime: string;
    /**
     * The descriptive name for this instance as it appears in UIs. Must be unique per project and between 4 and 30 characters in length.
     */
    readonly displayName: string;
    /**
     * Deprecated. This field is not populated.
     *
     * @deprecated Deprecated. This field is not populated.
     */
    readonly endpointUris: string[];
    /**
     * Free instance metadata. Only populated for free instances.
     */
    readonly freeInstanceMetadata: outputs.spanner.v1.FreeInstanceMetadataResponse;
    /**
     * The `InstanceType` of the current instance.
     */
    readonly instanceType: string;
    /**
     * Cloud Labels are a flexible and lightweight mechanism for organizing cloud resources into groups that reflect a customer's organizational needs and deployment strategies. Cloud Labels can be used to filter collections of resources. They can be used to control how resource metrics are aggregated. And they can be used as arguments to policy management rules (e.g. route, firewall, load balancing, etc.). * Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `a-z{0,62}`. * Label values must be between 0 and 63 characters long and must conform to the regular expression `[a-z0-9_-]{0,63}`. * No more than 64 labels can be associated with a given resource. See https://goo.gl/xmQnxf for more information on and examples of labels. If you plan to use labels in your own code, please note that additional characters may be allowed in the future. And so you are advised to use an internal label representation, such as JSON, which doesn't rely upon specific characters being disallowed. For example, representing labels as the string: name + "_" + value would prove problematic if we were to allow "_" in a future release.
     */
    readonly labels: {[key: string]: string};
    /**
     * A unique identifier for the instance, which cannot be changed after the instance is created. Values are of the form `projects//instances/a-z*[a-z0-9]`. The final segment of the name must be between 2 and 64 characters in length.
     */
    readonly name: string;
    /**
     * The number of nodes allocated to this instance. At most one of either node_count or processing_units should be present in the message. Users can set the node_count field to specify the target number of nodes allocated to the instance. This may be zero in API responses for instances that are not yet in state `READY`. See [the documentation](https://cloud.google.com/spanner/docs/compute-capacity) for more information about nodes and processing units.
     */
    readonly nodeCount: number;
    /**
     * The number of processing units allocated to this instance. At most one of processing_units or node_count should be present in the message. Users can set the processing_units field to specify the target number of processing units allocated to the instance. This may be zero in API responses for instances that are not yet in state `READY`. See [the documentation](https://cloud.google.com/spanner/docs/compute-capacity) for more information about nodes and processing units.
     */
    readonly processingUnits: number;
    /**
     * The current instance state. For CreateInstance, the state must be either omitted or set to `CREATING`. For UpdateInstance, the state must be either omitted or set to `READY`.
     */
    readonly state: string;
    /**
     * The time at which the instance was most recently updated.
     */
    readonly updateTime: string;
}
/**
 * Gets information about a particular instance.
 */
export function getInstanceOutput(args: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    return pulumi.output(args).apply((a: any) => getInstance(a, opts))
}

export interface GetInstanceOutputArgs {
    fieldMask?: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
