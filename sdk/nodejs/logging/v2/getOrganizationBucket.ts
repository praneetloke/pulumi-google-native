// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a log bucket.
 */
export function getOrganizationBucket(args: GetOrganizationBucketArgs, opts?: pulumi.InvokeOptions): Promise<GetOrganizationBucketResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:logging/v2:getOrganizationBucket", {
        "bucketId": args.bucketId,
        "location": args.location,
        "organizationId": args.organizationId,
    }, opts);
}

export interface GetOrganizationBucketArgs {
    bucketId: string;
    location: string;
    organizationId: string;
}

export interface GetOrganizationBucketResult {
    /**
     * Whether log analytics is enabled for this bucket.Once enabled, log analytics features cannot be disabled.
     */
    readonly analyticsEnabled: boolean;
    /**
     * The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.
     */
    readonly cmekSettings: outputs.logging.v2.CmekSettingsResponse;
    /**
     * The creation timestamp of the bucket. This is not set for any of the default buckets.
     */
    readonly createTime: string;
    /**
     * Describes this bucket.
     */
    readonly description: string;
    /**
     * A list of indexed fields and related configuration data.
     */
    readonly indexConfigs: outputs.logging.v2.IndexConfigResponse[];
    /**
     * The bucket lifecycle state.
     */
    readonly lifecycleState: string;
    /**
     * Whether the bucket is locked.The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
     */
    readonly locked: boolean;
    /**
     * The resource name of the bucket.For example:projects/my-project/locations/global/buckets/my-bucketFor a list of supported locations, see Supported Regions (https://cloud.google.com/logging/docs/region-support)For the location of global it is unspecified where log entries are actually stored.After a bucket has been created, the location cannot be changed.
     */
    readonly name: string;
    /**
     * Log entry field paths that are denied access in this bucket.The following fields and their children are eligible: textPayload, jsonPayload, protoPayload, httpRequest, labels, sourceLocation.Restricting a repeated field will restrict all values. Adding a parent will block all child fields. (e.g. foo.bar will block foo.bar.baz)
     */
    readonly restrictedFields: string[];
    /**
     * Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
     */
    readonly retentionDays: number;
    /**
     * The last update timestamp of the bucket.
     */
    readonly updateTime: string;
}
/**
 * Gets a log bucket.
 */
export function getOrganizationBucketOutput(args: GetOrganizationBucketOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrganizationBucketResult> {
    return pulumi.output(args).apply((a: any) => getOrganizationBucket(a, opts))
}

export interface GetOrganizationBucketOutputArgs {
    bucketId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
}
