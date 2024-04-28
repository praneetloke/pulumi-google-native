// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a PipelineJob.
 */
export function getPipelineJob(args: GetPipelineJobArgs, opts?: pulumi.InvokeOptions): Promise<GetPipelineJobResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:aiplatform/v1beta1:getPipelineJob", {
        "location": args.location,
        "pipelineJobId": args.pipelineJobId,
        "project": args.project,
    }, opts);
}

export interface GetPipelineJobArgs {
    location: string;
    pipelineJobId: string;
    project?: string;
}

export interface GetPipelineJobResult {
    /**
     * Pipeline creation time.
     */
    readonly createTime: string;
    /**
     * The display name of the Pipeline. The name can be up to 128 characters long and can consist of any UTF-8 characters.
     */
    readonly displayName: string;
    /**
     * Customer-managed encryption key spec for a pipelineJob. If set, this PipelineJob and all of its sub-resources will be secured by this key.
     */
    readonly encryptionSpec: outputs.aiplatform.v1beta1.GoogleCloudAiplatformV1beta1EncryptionSpecResponse;
    /**
     * Pipeline end time.
     */
    readonly endTime: string;
    /**
     * The error that occurred during pipeline execution. Only populated when the pipeline's state is FAILED or CANCELLED.
     */
    readonly error: outputs.aiplatform.v1beta1.GoogleRpcStatusResponse;
    /**
     * The details of pipeline run. Not available in the list view.
     */
    readonly jobDetail: outputs.aiplatform.v1beta1.GoogleCloudAiplatformV1beta1PipelineJobDetailResponse;
    /**
     * The labels with user-defined metadata to organize PipelineJob. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels. Note there is some reserved label key for Vertex AI Pipelines. - `vertex-ai-pipelines-run-billing-id`, user set value will get overrided.
     */
    readonly labels: {[key: string]: string};
    /**
     * The resource name of the PipelineJob.
     */
    readonly name: string;
    /**
     * The full name of the Compute Engine [network](/compute/docs/networks-and-firewalls#networks) to which the Pipeline Job's workload should be peered. For example, `projects/12345/global/networks/myVPC`. [Format](/compute/docs/reference/rest/v1/networks/insert) is of the form `projects/{project}/global/networks/{network}`. Where {project} is a project number, as in `12345`, and {network} is a network name. Private services access must already be configured for the network. Pipeline job will apply the network configuration to the Google Cloud resources being launched, if applied, such as Vertex AI Training or Dataflow job. If left unspecified, the workload is not peered with any network.
     */
    readonly network: string;
    /**
     * The spec of the pipeline.
     */
    readonly pipelineSpec: {[key: string]: any};
    /**
     * A list of names for the reserved ip ranges under the VPC network that can be used for this Pipeline Job's workload. If set, we will deploy the Pipeline Job's workload within the provided ip ranges. Otherwise, the job will be deployed to any ip ranges under the provided VPC network. Example: ['vertex-ai-ip-range'].
     */
    readonly reservedIpRanges: string[];
    /**
     * Runtime config of the pipeline.
     */
    readonly runtimeConfig: outputs.aiplatform.v1beta1.GoogleCloudAiplatformV1beta1PipelineJobRuntimeConfigResponse;
    /**
     * The schedule resource name. Only returned if the Pipeline is created by Schedule API.
     */
    readonly scheduleName: string;
    /**
     * The service account that the pipeline workload runs as. If not specified, the Compute Engine default service account in the project will be used. See https://cloud.google.com/compute/docs/access/service-accounts#default_service_account Users starting the pipeline must have the `iam.serviceAccounts.actAs` permission on this service account.
     */
    readonly serviceAccount: string;
    /**
     * Pipeline start time.
     */
    readonly startTime: string;
    /**
     * The detailed state of the job.
     */
    readonly state: string;
    /**
     * Pipeline template metadata. Will fill up fields if PipelineJob.template_uri is from supported template registry.
     */
    readonly templateMetadata: outputs.aiplatform.v1beta1.GoogleCloudAiplatformV1beta1PipelineTemplateMetadataResponse;
    /**
     * A template uri from where the PipelineJob.pipeline_spec, if empty, will be downloaded. Currently, only uri from Vertex Template Registry & Gallery is supported. Reference to https://cloud.google.com/vertex-ai/docs/pipelines/create-pipeline-template.
     */
    readonly templateUri: string;
    /**
     * Timestamp when this PipelineJob was most recently updated.
     */
    readonly updateTime: string;
}
/**
 * Gets a PipelineJob.
 */
export function getPipelineJobOutput(args: GetPipelineJobOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPipelineJobResult> {
    return pulumi.output(args).apply((a: any) => getPipelineJob(a, opts))
}

export interface GetPipelineJobOutputArgs {
    location: pulumi.Input<string>;
    pipelineJobId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
