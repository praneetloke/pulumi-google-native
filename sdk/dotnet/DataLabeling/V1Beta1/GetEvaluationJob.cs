// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataLabeling.V1Beta1
{
    public static class GetEvaluationJob
    {
        /// <summary>
        /// Gets an evaluation job by resource name.
        /// </summary>
        public static Task<GetEvaluationJobResult> InvokeAsync(GetEvaluationJobArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEvaluationJobResult>("google-native:datalabeling/v1beta1:getEvaluationJob", args ?? new GetEvaluationJobArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an evaluation job by resource name.
        /// </summary>
        public static Output<GetEvaluationJobResult> Invoke(GetEvaluationJobInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEvaluationJobResult>("google-native:datalabeling/v1beta1:getEvaluationJob", args ?? new GetEvaluationJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEvaluationJobArgs : global::Pulumi.InvokeArgs
    {
        [Input("evaluationJobId", required: true)]
        public string EvaluationJobId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetEvaluationJobArgs()
        {
        }
        public static new GetEvaluationJobArgs Empty => new GetEvaluationJobArgs();
    }

    public sealed class GetEvaluationJobInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("evaluationJobId", required: true)]
        public Input<string> EvaluationJobId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetEvaluationJobInvokeArgs()
        {
        }
        public static new GetEvaluationJobInvokeArgs Empty => new GetEvaluationJobInvokeArgs();
    }


    [OutputType]
    public sealed class GetEvaluationJobResult
    {
        /// <summary>
        /// Name of the AnnotationSpecSet describing all the labels that your machine learning model outputs. You must create this resource before you create an evaluation job and provide its name in the following format: "projects/{project_id}/annotationSpecSets/{annotation_spec_set_id}"
        /// </summary>
        public readonly string AnnotationSpecSet;
        /// <summary>
        /// Every time the evaluation job runs and an error occurs, the failed attempt is appended to this array.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDatalabelingV1beta1AttemptResponse> Attempts;
        /// <summary>
        /// Timestamp of when this evaluation job was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Description of the job. The description can be up to 25,000 characters long.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Configuration details for the evaluation job.
        /// </summary>
        public readonly Outputs.GoogleCloudDatalabelingV1beta1EvaluationJobConfigResponse EvaluationJobConfig;
        /// <summary>
        /// Whether you want Data Labeling Service to provide ground truth labels for prediction input. If you want the service to assign human labelers to annotate your data, set this to `true`. If you want to provide your own ground truth labels in the evaluation job's BigQuery table, set this to `false`.
        /// </summary>
        public readonly bool LabelMissingGroundTruth;
        /// <summary>
        /// The [AI Platform Prediction model version](/ml-engine/docs/prediction-overview) to be evaluated. Prediction input and output is sampled from this model version. When creating an evaluation job, specify the model version in the following format: "projects/{project_id}/models/{model_name}/versions/{version_name}" There can only be one evaluation job per model version.
        /// </summary>
        public readonly string ModelVersion;
        /// <summary>
        /// After you create a job, Data Labeling Service assigns a name to the job with the following format: "projects/{project_id}/evaluationJobs/ {evaluation_job_id}"
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Describes the interval at which the job runs. This interval must be at least 1 day, and it is rounded to the nearest day. For example, if you specify a 50-hour interval, the job runs every 2 days. You can provide the schedule in [crontab format](/scheduler/docs/configuring/cron-job-schedules) or in an [English-like format](/appengine/docs/standard/python/config/cronref#schedule_format). Regardless of what you specify, the job will run at 10:00 AM UTC. Only the interval from this schedule is used, not the specific time of day.
        /// </summary>
        public readonly string Schedule;
        /// <summary>
        /// Describes the current state of the job.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetEvaluationJobResult(
            string annotationSpecSet,

            ImmutableArray<Outputs.GoogleCloudDatalabelingV1beta1AttemptResponse> attempts,

            string createTime,

            string description,

            Outputs.GoogleCloudDatalabelingV1beta1EvaluationJobConfigResponse evaluationJobConfig,

            bool labelMissingGroundTruth,

            string modelVersion,

            string name,

            string schedule,

            string state)
        {
            AnnotationSpecSet = annotationSpecSet;
            Attempts = attempts;
            CreateTime = createTime;
            Description = description;
            EvaluationJobConfig = evaluationJobConfig;
            LabelMissingGroundTruth = labelMissingGroundTruth;
            ModelVersion = modelVersion;
            Name = name;
            Schedule = schedule;
            State = state;
        }
    }
}
