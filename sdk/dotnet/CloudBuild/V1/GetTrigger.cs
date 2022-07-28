// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1
{
    public static class GetTrigger
    {
        /// <summary>
        /// Returns information about a `BuildTrigger`. This API is experimental.
        /// </summary>
        public static Task<GetTriggerResult> InvokeAsync(GetTriggerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTriggerResult>("google-native:cloudbuild/v1:getTrigger", args ?? new GetTriggerArgs(), options.WithDefaults());

        /// <summary>
        /// Returns information about a `BuildTrigger`. This API is experimental.
        /// </summary>
        public static Output<GetTriggerResult> Invoke(GetTriggerInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTriggerResult>("google-native:cloudbuild/v1:getTrigger", args ?? new GetTriggerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTriggerArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        [Input("triggerId", required: true)]
        public string TriggerId { get; set; } = null!;

        public GetTriggerArgs()
        {
        }
        public static new GetTriggerArgs Empty => new GetTriggerArgs();
    }

    public sealed class GetTriggerInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("triggerId", required: true)]
        public Input<string> TriggerId { get; set; } = null!;

        public GetTriggerInvokeArgs()
        {
        }
        public static new GetTriggerInvokeArgs Empty => new GetTriggerInvokeArgs();
    }


    [OutputType]
    public sealed class GetTriggerResult
    {
        /// <summary>
        /// Configuration for manual approval to start a build invocation of this BuildTrigger.
        /// </summary>
        public readonly Outputs.ApprovalConfigResponse ApprovalConfig;
        /// <summary>
        /// Autodetect build configuration. The following precedence is used (case insensitive): 1. cloudbuild.yaml 2. cloudbuild.yml 3. cloudbuild.json 4. Dockerfile Currently only available for GitHub App Triggers.
        /// </summary>
        public readonly bool Autodetect;
        /// <summary>
        /// BitbucketServerTriggerConfig describes the configuration of a trigger that creates a build whenever a Bitbucket Server event is received.
        /// </summary>
        public readonly Outputs.BitbucketServerTriggerConfigResponse BitbucketServerTriggerConfig;
        /// <summary>
        /// Contents of the build template.
        /// </summary>
        public readonly Outputs.BuildResponse Build;
        /// <summary>
        /// Time when the trigger was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Human-readable description of this trigger.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// If true, the trigger will never automatically execute a build.
        /// </summary>
        public readonly bool Disabled;
        /// <summary>
        /// EventType allows the user to explicitly set the type of event to which this BuildTrigger should respond. This field will be validated against the rest of the configuration if it is set.
        /// </summary>
        public readonly string EventType;
        /// <summary>
        /// Path, from the source root, to the build configuration file (i.e. cloudbuild.yaml).
        /// </summary>
        public readonly string Filename;
        /// <summary>
        /// A Common Expression Language string.
        /// </summary>
        public readonly string Filter;
        /// <summary>
        /// The file source describing the local or remote Build template.
        /// </summary>
        public readonly Outputs.GitFileSourceResponse GitFileSource;
        /// <summary>
        /// GitHubEventsConfig describes the configuration of a trigger that creates a build whenever a GitHub event is received. Mutually exclusive with `trigger_template`.
        /// </summary>
        public readonly Outputs.GitHubEventsConfigResponse Github;
        /// <summary>
        /// ignored_files and included_files are file glob matches using https://golang.org/pkg/path/filepath/#Match extended with support for "**". If ignored_files and changed files are both empty, then they are not used to determine whether or not to trigger a build. If ignored_files is not empty, then we ignore any files that match any of the ignored_file globs. If the change has no files that are outside of the ignored_files globs, then we do not trigger a build.
        /// </summary>
        public readonly ImmutableArray<string> IgnoredFiles;
        /// <summary>
        /// If set to INCLUDE_BUILD_LOGS_WITH_STATUS, log url will be shown on GitHub page when build status is final. Setting this field to INCLUDE_BUILD_LOGS_WITH_STATUS for non GitHub triggers results in INVALID_ARGUMENT error.
        /// </summary>
        public readonly string IncludeBuildLogs;
        /// <summary>
        /// If any of the files altered in the commit pass the ignored_files filter and included_files is empty, then as far as this filter is concerned, we should trigger the build. If any of the files altered in the commit pass the ignored_files filter and included_files is not empty, then we make sure that at least one of those files matches a included_files glob. If not, then we do not trigger a build.
        /// </summary>
        public readonly ImmutableArray<string> IncludedFiles;
        /// <summary>
        /// User-assigned name of the trigger. Must be unique within the project. Trigger names must meet the following requirements: + They must contain only alphanumeric characters and dashes. + They can be 1-64 characters long. + They must begin and end with an alphanumeric character.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// PubsubConfig describes the configuration of a trigger that creates a build whenever a Pub/Sub message is published.
        /// </summary>
        public readonly Outputs.PubsubConfigResponse PubsubConfig;
        /// <summary>
        /// The `Trigger` name with format: `projects/{project}/locations/{location}/triggers/{trigger}`, where {trigger} is a unique identifier generated by the service.
        /// </summary>
        public readonly string ResourceName;
        /// <summary>
        /// The service account used for all user-controlled operations including UpdateBuildTrigger, RunBuildTrigger, CreateBuild, and CancelBuild. If no service account is set, then the standard Cloud Build service account ([PROJECT_NUM]@system.gserviceaccount.com) will be used instead. Format: `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT_ID_OR_EMAIL}`
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// The repo and ref of the repository from which to build. This field is used only for those triggers that do not respond to SCM events. Triggers that respond to such events build source at whatever commit caused the event. This field is currently only used by Webhook, Pub/Sub, Manual, and Cron triggers.
        /// </summary>
        public readonly Outputs.GitRepoSourceResponse SourceToBuild;
        /// <summary>
        /// Substitutions for Build resource. The keys must match the following regular expression: `^_[A-Z0-9_]+$`.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Substitutions;
        /// <summary>
        /// Tags for annotation of a `BuildTrigger`
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// Template describing the types of source changes to trigger a build. Branch and tag names in trigger templates are interpreted as regular expressions. Any branch or tag change that matches that regular expression will trigger a build. Mutually exclusive with `github`.
        /// </summary>
        public readonly Outputs.RepoSourceResponse TriggerTemplate;
        /// <summary>
        /// WebhookConfig describes the configuration of a trigger that creates a build whenever a webhook is sent to a trigger's webhook URL.
        /// </summary>
        public readonly Outputs.WebhookConfigResponse WebhookConfig;

        [OutputConstructor]
        private GetTriggerResult(
            Outputs.ApprovalConfigResponse approvalConfig,

            bool autodetect,

            Outputs.BitbucketServerTriggerConfigResponse bitbucketServerTriggerConfig,

            Outputs.BuildResponse build,

            string createTime,

            string description,

            bool disabled,

            string eventType,

            string filename,

            string filter,

            Outputs.GitFileSourceResponse gitFileSource,

            Outputs.GitHubEventsConfigResponse github,

            ImmutableArray<string> ignoredFiles,

            string includeBuildLogs,

            ImmutableArray<string> includedFiles,

            string name,

            Outputs.PubsubConfigResponse pubsubConfig,

            string resourceName,

            string serviceAccount,

            Outputs.GitRepoSourceResponse sourceToBuild,

            ImmutableDictionary<string, string> substitutions,

            ImmutableArray<string> tags,

            Outputs.RepoSourceResponse triggerTemplate,

            Outputs.WebhookConfigResponse webhookConfig)
        {
            ApprovalConfig = approvalConfig;
            Autodetect = autodetect;
            BitbucketServerTriggerConfig = bitbucketServerTriggerConfig;
            Build = build;
            CreateTime = createTime;
            Description = description;
            Disabled = disabled;
            EventType = eventType;
            Filename = filename;
            Filter = filter;
            GitFileSource = gitFileSource;
            Github = github;
            IgnoredFiles = ignoredFiles;
            IncludeBuildLogs = includeBuildLogs;
            IncludedFiles = includedFiles;
            Name = name;
            PubsubConfig = pubsubConfig;
            ResourceName = resourceName;
            ServiceAccount = serviceAccount;
            SourceToBuild = sourceToBuild;
            Substitutions = substitutions;
            Tags = tags;
            TriggerTemplate = triggerTemplate;
            WebhookConfig = webhookConfig;
        }
    }
}
