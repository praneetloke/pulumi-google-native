// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets details of a single Featurestore.
 */
export function getFeaturestore(args: GetFeaturestoreArgs, opts?: pulumi.InvokeOptions): Promise<GetFeaturestoreResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:aiplatform/v1beta1:getFeaturestore", {
        "featurestoreId": args.featurestoreId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetFeaturestoreArgs {
    featurestoreId: string;
    location: string;
    project?: string;
}

export interface GetFeaturestoreResult {
    /**
     * Timestamp when this Featurestore was created.
     */
    readonly createTime: string;
    /**
     * Optional. Customer-managed encryption key spec for data storage. If set, both of the online and offline data storage will be secured by this key.
     */
    readonly encryptionSpec: outputs.aiplatform.v1beta1.GoogleCloudAiplatformV1beta1EncryptionSpecResponse;
    /**
     * Optional. Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
     */
    readonly etag: string;
    /**
     * Optional. The labels with user-defined metadata to organize your Featurestore. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information on and examples of labels. No more than 64 user labels can be associated with one Featurestore(System labels are excluded)." System reserved label keys are prefixed with "aiplatform.googleapis.com/" and are immutable.
     */
    readonly labels: {[key: string]: string};
    /**
     * Name of the Featurestore. Format: `projects/{project}/locations/{location}/featurestores/{featurestore}`
     */
    readonly name: string;
    /**
     * Optional. Config for online storage resources. The field should not co-exist with the field of `OnlineStoreReplicationConfig`. If both of it and OnlineStoreReplicationConfig are unset, the feature store will not have an online store and cannot be used for online serving.
     */
    readonly onlineServingConfig: outputs.aiplatform.v1beta1.GoogleCloudAiplatformV1beta1FeaturestoreOnlineServingConfigResponse;
    /**
     * Optional. TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage periodically removes obsolete feature values older than `online_storage_ttl_days` since the feature generation time. Note that `online_storage_ttl_days` should be less than or equal to `offline_storage_ttl_days` for each EntityType under a featurestore. If not set, default to 4000 days
     */
    readonly onlineStorageTtlDays: number;
    /**
     * State of the featurestore.
     */
    readonly state: string;
    /**
     * Timestamp when this Featurestore was last updated.
     */
    readonly updateTime: string;
}
/**
 * Gets details of a single Featurestore.
 */
export function getFeaturestoreOutput(args: GetFeaturestoreOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeaturestoreResult> {
    return pulumi.output(args).apply((a: any) => getFeaturestore(a, opts))
}

export interface GetFeaturestoreOutputArgs {
    featurestoreId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}