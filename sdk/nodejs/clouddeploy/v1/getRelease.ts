// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets details of a single Release.
 */
export function getRelease(args: GetReleaseArgs, opts?: pulumi.InvokeOptions): Promise<GetReleaseResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:clouddeploy/v1:getRelease", {
        "deliveryPipelineId": args.deliveryPipelineId,
        "location": args.location,
        "project": args.project,
        "releaseId": args.releaseId,
    }, opts);
}

export interface GetReleaseArgs {
    deliveryPipelineId: string;
    location: string;
    project?: string;
    releaseId: string;
}

export interface GetReleaseResult {
    /**
     * Indicates whether this is an abandoned release.
     */
    readonly abandoned: boolean;
    /**
     * User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
     */
    readonly annotations: {[key: string]: string};
    /**
     * List of artifacts to pass through to Skaffold command.
     */
    readonly buildArtifacts: outputs.clouddeploy.v1.BuildArtifactResponse[];
    /**
     * Information around the state of the Release.
     */
    readonly condition: outputs.clouddeploy.v1.ReleaseConditionResponse;
    /**
     * Time at which the `Release` was created.
     */
    readonly createTime: string;
    /**
     * Snapshot of the parent pipeline taken at release creation time.
     */
    readonly deliveryPipelineSnapshot: outputs.clouddeploy.v1.DeliveryPipelineResponse;
    /**
     * Optional. The deploy parameters to use for all targets in this release.
     */
    readonly deployParameters: {[key: string]: string};
    /**
     * Description of the `Release`. Max length is 255 characters.
     */
    readonly description: string;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     */
    readonly etag: string;
    /**
     * Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
     */
    readonly labels: {[key: string]: string};
    /**
     * Optional. Name of the `Release`. Format is `projects/{project}/locations/{location}/deliveryPipelines/{deliveryPipeline}/releases/a-z{0,62}`.
     */
    readonly name: string;
    /**
     * Time at which the render completed.
     */
    readonly renderEndTime: string;
    /**
     * Time at which the render began.
     */
    readonly renderStartTime: string;
    /**
     * Current state of the render operation.
     */
    readonly renderState: string;
    /**
     * Filepath of the Skaffold config inside of the config URI.
     */
    readonly skaffoldConfigPath: string;
    /**
     * Cloud Storage URI of tar.gz archive containing Skaffold configuration.
     */
    readonly skaffoldConfigUri: string;
    /**
     * The Skaffold version to use when operating on this release, such as "1.20.0". Not all versions are valid; Cloud Deploy supports a specific set of versions. If unset, the most recent supported Skaffold version will be used.
     */
    readonly skaffoldVersion: string;
    /**
     * Map from target ID to the target artifacts created during the render operation.
     */
    readonly targetArtifacts: outputs.clouddeploy.v1.TargetArtifactResponse;
    /**
     * Map from target ID to details of the render operation for that target.
     */
    readonly targetRenders: outputs.clouddeploy.v1.TargetRenderResponse;
    /**
     * Snapshot of the targets taken at release creation time.
     */
    readonly targetSnapshots: outputs.clouddeploy.v1.TargetResponse[];
    /**
     * Unique identifier of the `Release`.
     */
    readonly uid: string;
}
/**
 * Gets details of a single Release.
 */
export function getReleaseOutput(args: GetReleaseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReleaseResult> {
    return pulumi.output(args).apply((a: any) => getRelease(a, opts))
}

export interface GetReleaseOutputArgs {
    deliveryPipelineId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    releaseId: pulumi.Input<string>;
}
