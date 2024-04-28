// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1
{
    public static class GetBuild
    {
        /// <summary>
        /// Returns information about a previously requested build. The `Build` that is returned includes its status (such as `SUCCESS`, `FAILURE`, or `WORKING`), and timing information.
        /// </summary>
        public static Task<GetBuildResult> InvokeAsync(GetBuildArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBuildResult>("google-native:cloudbuild/v1:getBuild", args ?? new GetBuildArgs(), options.WithDefaults());

        /// <summary>
        /// Returns information about a previously requested build. The `Build` that is returned includes its status (such as `SUCCESS`, `FAILURE`, or `WORKING`), and timing information.
        /// </summary>
        public static Output<GetBuildResult> Invoke(GetBuildInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBuildResult>("google-native:cloudbuild/v1:getBuild", args ?? new GetBuildInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBuildArgs : global::Pulumi.InvokeArgs
    {
        [Input("buildId", required: true)]
        public string BuildId { get; set; } = null!;

        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetBuildArgs()
        {
        }
        public static new GetBuildArgs Empty => new GetBuildArgs();
    }

    public sealed class GetBuildInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("buildId", required: true)]
        public Input<string> BuildId { get; set; } = null!;

        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetBuildInvokeArgs()
        {
        }
        public static new GetBuildInvokeArgs Empty => new GetBuildInvokeArgs();
    }


    [OutputType]
    public sealed class GetBuildResult
    {
        /// <summary>
        /// Describes this build's approval configuration, status, and result.
        /// </summary>
        public readonly Outputs.BuildApprovalResponse Approval;
        /// <summary>
        /// Artifacts produced by the build that should be uploaded upon successful completion of all build steps.
        /// </summary>
        public readonly Outputs.ArtifactsResponse Artifacts;
        /// <summary>
        /// Secrets and secret environment variables.
        /// </summary>
        public readonly Outputs.SecretsResponse AvailableSecrets;
        /// <summary>
        /// The ID of the `BuildTrigger` that triggered this build, if it was triggered automatically.
        /// </summary>
        public readonly string BuildTriggerId;
        /// <summary>
        /// Time at which the request to create the build was received.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Contains information about the build when status=FAILURE.
        /// </summary>
        public readonly Outputs.FailureInfoResponse FailureInfo;
        /// <summary>
        /// Time at which execution of the build was finished. The difference between finish_time and start_time is the duration of the build's execution.
        /// </summary>
        public readonly string FinishTime;
        /// <summary>
        /// A list of images to be pushed upon the successful completion of all build steps. The images are pushed using the builder service account's credentials. The digests of the pushed images will be stored in the `Build` resource's results field. If any of the images fail to be pushed, the build status is marked `FAILURE`.
        /// </summary>
        public readonly ImmutableArray<string> Images;
        /// <summary>
        /// URL to logs for this build in Google Cloud Console.
        /// </summary>
        public readonly string LogUrl;
        /// <summary>
        /// Cloud Storage bucket where logs should be written (see [Bucket Name Requirements](https://cloud.google.com/storage/docs/bucket-naming#requirements)). Logs file names will be of the format `${logs_bucket}/log-${build_id}.txt`.
        /// </summary>
        public readonly string LogsBucket;
        /// <summary>
        /// The 'Build' name with format: `projects/{project}/locations/{location}/builds/{build}`, where {build} is a unique identifier generated by the service.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Special options for this build.
        /// </summary>
        public readonly Outputs.BuildOptionsResponse Options;
        /// <summary>
        /// ID of the project.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// TTL in queue for this build. If provided and the build is enqueued longer than this value, the build will expire and the build status will be `EXPIRED`. The TTL starts ticking from create_time.
        /// </summary>
        public readonly string QueueTtl;
        /// <summary>
        /// Results of the build.
        /// </summary>
        public readonly Outputs.ResultsResponse Results;
        /// <summary>
        /// Secrets to decrypt using Cloud Key Management Service. Note: Secret Manager is the recommended technique for managing sensitive data with Cloud Build. Use `available_secrets` to configure builds to access secrets from Secret Manager. For instructions, see: https://cloud.google.com/cloud-build/docs/securing-builds/use-secrets
        /// </summary>
        public readonly ImmutableArray<Outputs.SecretResponse> Secrets;
        /// <summary>
        /// IAM service account whose credentials will be used at build runtime. Must be of the format `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. ACCOUNT can be email address or uniqueId of the service account. 
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// The location of the source files to build.
        /// </summary>
        public readonly Outputs.SourceResponse Source;
        /// <summary>
        /// A permanent fixed identifier for source.
        /// </summary>
        public readonly Outputs.SourceProvenanceResponse SourceProvenance;
        /// <summary>
        /// Time at which execution of the build was started.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Status of the build.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Customer-readable message about the current status.
        /// </summary>
        public readonly string StatusDetail;
        /// <summary>
        /// The operations to be performed on the workspace.
        /// </summary>
        public readonly ImmutableArray<Outputs.BuildStepResponse> Steps;
        /// <summary>
        /// Substitutions data for `Build` resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Substitutions;
        /// <summary>
        /// Tags for annotation of a `Build`. These are not docker tags.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// Amount of time that this build should be allowed to run, to second granularity. If this amount of time elapses, work on the build will cease and the build status will be `TIMEOUT`. `timeout` starts ticking from `startTime`. Default time is 60 minutes.
        /// </summary>
        public readonly string Timeout;
        /// <summary>
        /// Stores timing information for phases of the build. Valid keys are: * BUILD: time to execute all build steps. * PUSH: time to push all artifacts including docker images and non docker artifacts. * FETCHSOURCE: time to fetch source. * SETUPBUILD: time to set up build. If the build does not specify source or images, these keys will not be included.
        /// </summary>
        public readonly Outputs.TimeSpanResponse Timing;
        /// <summary>
        /// Non-fatal problems encountered during the execution of the build.
        /// </summary>
        public readonly ImmutableArray<Outputs.WarningResponse> Warnings;

        [OutputConstructor]
        private GetBuildResult(
            Outputs.BuildApprovalResponse approval,

            Outputs.ArtifactsResponse artifacts,

            Outputs.SecretsResponse availableSecrets,

            string buildTriggerId,

            string createTime,

            Outputs.FailureInfoResponse failureInfo,

            string finishTime,

            ImmutableArray<string> images,

            string logUrl,

            string logsBucket,

            string name,

            Outputs.BuildOptionsResponse options,

            string project,

            string queueTtl,

            Outputs.ResultsResponse results,

            ImmutableArray<Outputs.SecretResponse> secrets,

            string serviceAccount,

            Outputs.SourceResponse source,

            Outputs.SourceProvenanceResponse sourceProvenance,

            string startTime,

            string status,

            string statusDetail,

            ImmutableArray<Outputs.BuildStepResponse> steps,

            ImmutableDictionary<string, string> substitutions,

            ImmutableArray<string> tags,

            string timeout,

            Outputs.TimeSpanResponse timing,

            ImmutableArray<Outputs.WarningResponse> warnings)
        {
            Approval = approval;
            Artifacts = artifacts;
            AvailableSecrets = availableSecrets;
            BuildTriggerId = buildTriggerId;
            CreateTime = createTime;
            FailureInfo = failureInfo;
            FinishTime = finishTime;
            Images = images;
            LogUrl = logUrl;
            LogsBucket = logsBucket;
            Name = name;
            Options = options;
            Project = project;
            QueueTtl = queueTtl;
            Results = results;
            Secrets = secrets;
            ServiceAccount = serviceAccount;
            Source = source;
            SourceProvenance = sourceProvenance;
            StartTime = startTime;
            Status = status;
            StatusDetail = statusDetail;
            Steps = steps;
            Substitutions = substitutions;
            Tags = tags;
            Timeout = timeout;
            Timing = timing;
            Warnings = warnings;
        }
    }
}
