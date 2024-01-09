// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets the state of the specified Cloud Dataflow job. To get the state of a job, we recommend using `projects.locations.jobs.get` with a [regional endpoint] (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints). Using `projects.jobs.get` is not recommended, as you can only get the state of jobs that are running in `us-central1`.
 */
export function getJob(args: GetJobArgs, opts?: pulumi.InvokeOptions): Promise<GetJobResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:dataflow/v1b3:getJob", {
        "jobId": args.jobId,
        "location": args.location,
        "project": args.project,
        "view": args.view,
    }, opts);
}

export interface GetJobArgs {
    jobId: string;
    location: string;
    project?: string;
    view?: string;
}

export interface GetJobResult {
    /**
     * The client's unique identifier of the job, re-used across retried attempts. If this field is set, the service will ensure its uniqueness. The request to create a job will fail if the service has knowledge of a previously submitted job with the same client's ID and job name. The caller may use this field to ensure idempotence of job creation across retried attempts to create a job. By default, the field is empty and, in that case, the service ignores it.
     */
    readonly clientRequestId: string;
    /**
     * The timestamp when the job was initially created. Immutable and set by the Cloud Dataflow service.
     */
    readonly createTime: string;
    /**
     * If this is specified, the job's initial state is populated from the given snapshot.
     */
    readonly createdFromSnapshotId: string;
    /**
     * The current state of the job. Jobs are created in the `JOB_STATE_STOPPED` state unless otherwise specified. A job in the `JOB_STATE_RUNNING` state may asynchronously enter a terminal state. After a job has reached a terminal state, no further state updates may be made. This field might be mutated by the Dataflow service; callers cannot mutate it.
     */
    readonly currentState: string;
    /**
     * The timestamp associated with the current state.
     */
    readonly currentStateTime: string;
    /**
     * The environment for the job.
     */
    readonly environment: outputs.dataflow.v1b3.EnvironmentResponse;
    /**
     * Deprecated.
     *
     * @deprecated Deprecated.
     */
    readonly executionInfo: outputs.dataflow.v1b3.JobExecutionInfoResponse;
    /**
     * This field is populated by the Dataflow service to support filtering jobs by the metadata values provided here. Populated for ListJobs and all GetJob views SUMMARY and higher.
     */
    readonly jobMetadata: outputs.dataflow.v1b3.JobMetadataResponse;
    /**
     * User-defined labels for this job. The labels map can contain no more than 64 entries. Entries of the labels map are UTF8 strings that comply with the following restrictions: * Keys must conform to regexp: \p{Ll}\p{Lo}{0,62} * Values must conform to regexp: [\p{Ll}\p{Lo}\p{N}_-]{0,63} * Both keys and values are additionally constrained to be <= 128 bytes in size.
     */
    readonly labels: {[key: string]: string};
    /**
     * The [regional endpoint] (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints) that contains this job.
     */
    readonly location: string;
    /**
     * The user-specified Dataflow job name. Only one active job with a given name can exist in a project within one region at any given time. Jobs in different regions can have the same name. If a caller attempts to create a job with the same name as an active job that already exists, the attempt returns the existing job. The name must match the regular expression `[a-z]([-a-z0-9]{0,1022}[a-z0-9])?`
     */
    readonly name: string;
    /**
     * Preliminary field: The format of this data may change at any time. A description of the user pipeline and stages through which it is executed. Created by Cloud Dataflow service. Only retrieved with JOB_VIEW_DESCRIPTION or JOB_VIEW_ALL.
     */
    readonly pipelineDescription: outputs.dataflow.v1b3.PipelineDescriptionResponse;
    /**
     * The ID of the Google Cloud project that the job belongs to.
     */
    readonly project: string;
    /**
     * If this job is an update of an existing job, this field is the job ID of the job it replaced. When sending a `CreateJobRequest`, you can update a job by specifying it here. The job named here is stopped, and its intermediate state is transferred to this job.
     */
    readonly replaceJobId: string;
    /**
     * If another job is an update of this job (and thus, this job is in `JOB_STATE_UPDATED`), this field contains the ID of that job.
     */
    readonly replacedByJobId: string;
    /**
     * The job's requested state. Applies to `UpdateJob` requests. Set `requested_state` with `UpdateJob` requests to switch between the states `JOB_STATE_STOPPED` and `JOB_STATE_RUNNING`. You can also use `UpdateJob` requests to change a job's state from `JOB_STATE_RUNNING` to `JOB_STATE_CANCELLED`, `JOB_STATE_DONE`, or `JOB_STATE_DRAINED`. These states irrevocably terminate the job if it hasn't already reached a terminal state. This field has no effect on `CreateJob` requests.
     */
    readonly requestedState: string;
    /**
     * This field may ONLY be modified at runtime using the projects.jobs.update method to adjust job behavior. This field has no effect when specified at job creation.
     */
    readonly runtimeUpdatableParams: outputs.dataflow.v1b3.RuntimeUpdatableParamsResponse;
    /**
     * Reserved for future use. This field is set only in responses from the server; it is ignored if it is set in any requests.
     */
    readonly satisfiesPzi: boolean;
    /**
     * Reserved for future use. This field is set only in responses from the server; it is ignored if it is set in any requests.
     */
    readonly satisfiesPzs: boolean;
    /**
     * This field may be mutated by the Cloud Dataflow service; callers cannot mutate it.
     */
    readonly stageStates: outputs.dataflow.v1b3.ExecutionStageStateResponse[];
    /**
     * The timestamp when the job was started (transitioned to JOB_STATE_PENDING). Flexible resource scheduling jobs are started with some delay after job creation, so start_time is unset before start and is updated when the job is started by the Cloud Dataflow service. For other jobs, start_time always equals to create_time and is immutable and set by the Cloud Dataflow service.
     */
    readonly startTime: string;
    /**
     * Exactly one of step or steps_location should be specified. The top-level steps that constitute the entire job. Only retrieved with JOB_VIEW_ALL.
     */
    readonly steps: outputs.dataflow.v1b3.StepResponse[];
    /**
     * The Cloud Storage location where the steps are stored.
     */
    readonly stepsLocation: string;
    /**
     * A set of files the system should be aware of that are used for temporary storage. These temporary files will be removed on job completion. No duplicates are allowed. No file patterns are supported. The supported files are: Google Cloud Storage: storage.googleapis.com/{bucket}/{object} bucket.storage.googleapis.com/{object}
     */
    readonly tempFiles: string[];
    /**
     * The map of transform name prefixes of the job to be replaced to the corresponding name prefixes of the new job.
     */
    readonly transformNameMapping: {[key: string]: string};
    /**
     * The type of Dataflow job.
     */
    readonly type: string;
}
/**
 * Gets the state of the specified Cloud Dataflow job. To get the state of a job, we recommend using `projects.locations.jobs.get` with a [regional endpoint] (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints). Using `projects.jobs.get` is not recommended, as you can only get the state of jobs that are running in `us-central1`.
 */
export function getJobOutput(args: GetJobOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetJobResult> {
    return pulumi.output(args).apply((a: any) => getJob(a, opts))
}

export interface GetJobOutputArgs {
    jobId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    view?: pulumi.Input<string>;
}
