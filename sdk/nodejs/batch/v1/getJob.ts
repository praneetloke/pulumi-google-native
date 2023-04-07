// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Get a Job specified by its resource name.
 */
export function getJob(args: GetJobArgs, opts?: pulumi.InvokeOptions): Promise<GetJobResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:batch/v1:getJob", {
        "jobId": args.jobId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetJobArgs {
    jobId: string;
    location: string;
    project?: string;
}

export interface GetJobResult {
    /**
     * Compute resource allocation for all TaskGroups in the Job.
     */
    readonly allocationPolicy: outputs.batch.v1.AllocationPolicyResponse;
    /**
     * When the Job was created.
     */
    readonly createTime: string;
    /**
     * Labels for the Job. Labels could be user provided or system generated. For example, "labels": { "department": "finance", "environment": "test" } You can assign up to 64 labels. [Google Compute Engine label restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) apply. Label names that start with "goog-" or "google-" are reserved.
     */
    readonly labels: {[key: string]: string};
    /**
     * Log preservation policy for the Job.
     */
    readonly logsPolicy: outputs.batch.v1.LogsPolicyResponse;
    /**
     * Job name. For example: "projects/123456/locations/us-central1/jobs/job01".
     */
    readonly name: string;
    /**
     * Notification configurations.
     */
    readonly notifications: outputs.batch.v1.JobNotificationResponse[];
    /**
     * Priority of the Job. The valid value range is [0, 100). Default value is 0. Higher value indicates higher priority. A job with higher priority value is more likely to run earlier if all other requirements are satisfied.
     */
    readonly priority: string;
    /**
     * Job status. It is read only for users.
     */
    readonly status: outputs.batch.v1.JobStatusResponse;
    /**
     * TaskGroups in the Job. Only one TaskGroup is supported now.
     */
    readonly taskGroups: outputs.batch.v1.TaskGroupResponse[];
    /**
     * A system generated unique ID (in UUID4 format) for the Job.
     */
    readonly uid: string;
    /**
     * The last time the Job was updated.
     */
    readonly updateTime: string;
}
/**
 * Get a Job specified by its resource name.
 */
export function getJobOutput(args: GetJobOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetJobResult> {
    return pulumi.output(args).apply((a: any) => getJob(a, opts))
}

export interface GetJobOutputArgs {
    jobId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
