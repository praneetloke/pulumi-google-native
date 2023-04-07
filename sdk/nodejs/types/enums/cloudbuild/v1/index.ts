// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const BuildOptionsDefaultLogsBucketBehavior = {
    /**
     * Unspecified.
     */
    DefaultLogsBucketBehaviorUnspecified: "DEFAULT_LOGS_BUCKET_BEHAVIOR_UNSPECIFIED",
    /**
     * Bucket is located in user-owned project in the same region as the build. The builder service account must have access to create and write to GCS buckets in the build project.
     */
    RegionalUserOwnedBucket: "REGIONAL_USER_OWNED_BUCKET",
} as const;

/**
 * Optional. Option to specify how default logs buckets are setup.
 */
export type BuildOptionsDefaultLogsBucketBehavior = (typeof BuildOptionsDefaultLogsBucketBehavior)[keyof typeof BuildOptionsDefaultLogsBucketBehavior];

export const BuildOptionsLogStreamingOption = {
    /**
     * Service may automatically determine build log streaming behavior.
     */
    StreamDefault: "STREAM_DEFAULT",
    /**
     * Build logs should be streamed to Google Cloud Storage.
     */
    StreamOn: "STREAM_ON",
    /**
     * Build logs should not be streamed to Google Cloud Storage; they will be written when the build is completed.
     */
    StreamOff: "STREAM_OFF",
} as const;

/**
 * Option to define build log streaming behavior to Google Cloud Storage.
 */
export type BuildOptionsLogStreamingOption = (typeof BuildOptionsLogStreamingOption)[keyof typeof BuildOptionsLogStreamingOption];

export const BuildOptionsLogging = {
    /**
     * The service determines the logging mode. The default is `LEGACY`. Do not rely on the default logging behavior as it may change in the future.
     */
    LoggingUnspecified: "LOGGING_UNSPECIFIED",
    /**
     * Build logs are stored in Cloud Logging and Cloud Storage.
     */
    Legacy: "LEGACY",
    /**
     * Build logs are stored in Cloud Storage.
     */
    GcsOnly: "GCS_ONLY",
    /**
     * This option is the same as CLOUD_LOGGING_ONLY.
     */
    StackdriverOnly: "STACKDRIVER_ONLY",
    /**
     * Build logs are stored in Cloud Logging. Selecting this option will not allow [logs streaming](https://cloud.google.com/sdk/gcloud/reference/builds/log).
     */
    CloudLoggingOnly: "CLOUD_LOGGING_ONLY",
    /**
     * Turn off all logging. No build logs will be captured.
     */
    None: "NONE",
} as const;

/**
 * Option to specify the logging mode, which determines if and where build logs are stored.
 */
export type BuildOptionsLogging = (typeof BuildOptionsLogging)[keyof typeof BuildOptionsLogging];

export const BuildOptionsMachineType = {
    /**
     * Standard machine type.
     */
    Unspecified: "UNSPECIFIED",
    /**
     * Highcpu machine with 8 CPUs.
     */
    N1Highcpu8: "N1_HIGHCPU_8",
    /**
     * Highcpu machine with 32 CPUs.
     */
    N1Highcpu32: "N1_HIGHCPU_32",
    /**
     * Highcpu e2 machine with 8 CPUs.
     */
    E2Highcpu8: "E2_HIGHCPU_8",
    /**
     * Highcpu e2 machine with 32 CPUs.
     */
    E2Highcpu32: "E2_HIGHCPU_32",
} as const;

/**
 * Compute Engine machine type on which to run the build.
 */
export type BuildOptionsMachineType = (typeof BuildOptionsMachineType)[keyof typeof BuildOptionsMachineType];

export const BuildOptionsRequestedVerifyOption = {
    /**
     * Not a verifiable build (the default).
     */
    NotVerified: "NOT_VERIFIED",
    /**
     * Build must be verified.
     */
    Verified: "VERIFIED",
} as const;

/**
 * Requested verifiability options.
 */
export type BuildOptionsRequestedVerifyOption = (typeof BuildOptionsRequestedVerifyOption)[keyof typeof BuildOptionsRequestedVerifyOption];

export const BuildOptionsSourceProvenanceHashItem = {
    /**
     * No hash requested.
     */
    None: "NONE",
    /**
     * Use a sha256 hash.
     */
    Sha256: "SHA256",
    /**
     * Use a md5 hash.
     */
    Md5: "MD5",
} as const;

export type BuildOptionsSourceProvenanceHashItem = (typeof BuildOptionsSourceProvenanceHashItem)[keyof typeof BuildOptionsSourceProvenanceHashItem];

export const BuildOptionsSubstitutionOption = {
    /**
     * Fails the build if error in substitutions checks, like missing a substitution in the template or in the map.
     */
    MustMatch: "MUST_MATCH",
    /**
     * Do not fail the build if error in substitutions checks.
     */
    AllowLoose: "ALLOW_LOOSE",
} as const;

/**
 * Option to specify behavior when there is an error in the substitution checks. NOTE: this is always set to ALLOW_LOOSE for triggered builds and cannot be overridden in the build configuration file.
 */
export type BuildOptionsSubstitutionOption = (typeof BuildOptionsSubstitutionOption)[keyof typeof BuildOptionsSubstitutionOption];

export const GitFileSourceRepoType = {
    /**
     * The default, unknown repo type. Don't use it, instead use one of the other repo types.
     */
    Unknown: "UNKNOWN",
    /**
     * A Google Cloud Source Repositories-hosted repo.
     */
    CloudSourceRepositories: "CLOUD_SOURCE_REPOSITORIES",
    /**
     * A GitHub-hosted repo not necessarily on "github.com" (i.e. GitHub Enterprise).
     */
    Github: "GITHUB",
    /**
     * A Bitbucket Server-hosted repo.
     */
    BitbucketServer: "BITBUCKET_SERVER",
    /**
     * A GitLab-hosted repo.
     */
    Gitlab: "GITLAB",
} as const;

/**
 * See RepoType above.
 */
export type GitFileSourceRepoType = (typeof GitFileSourceRepoType)[keyof typeof GitFileSourceRepoType];

export const GitRepoSourceRepoType = {
    /**
     * The default, unknown repo type. Don't use it, instead use one of the other repo types.
     */
    Unknown: "UNKNOWN",
    /**
     * A Google Cloud Source Repositories-hosted repo.
     */
    CloudSourceRepositories: "CLOUD_SOURCE_REPOSITORIES",
    /**
     * A GitHub-hosted repo not necessarily on "github.com" (i.e. GitHub Enterprise).
     */
    Github: "GITHUB",
    /**
     * A Bitbucket Server-hosted repo.
     */
    BitbucketServer: "BITBUCKET_SERVER",
    /**
     * A GitLab-hosted repo.
     */
    Gitlab: "GITLAB",
} as const;

/**
 * See RepoType below.
 */
export type GitRepoSourceRepoType = (typeof GitRepoSourceRepoType)[keyof typeof GitRepoSourceRepoType];

export const NetworkConfigEgressOption = {
    /**
     * If set, defaults to PUBLIC_EGRESS.
     */
    EgressOptionUnspecified: "EGRESS_OPTION_UNSPECIFIED",
    /**
     * If set, workers are created without any public address, which prevents network egress to public IPs unless a network proxy is configured.
     */
    NoPublicEgress: "NO_PUBLIC_EGRESS",
    /**
     * If set, workers are created with a public address which allows for public internet egress.
     */
    PublicEgress: "PUBLIC_EGRESS",
} as const;

/**
 * Option to configure network egress for the workers.
 */
export type NetworkConfigEgressOption = (typeof NetworkConfigEgressOption)[keyof typeof NetworkConfigEgressOption];

export const PubsubConfigState = {
    /**
     * The subscription configuration has not been checked.
     */
    StateUnspecified: "STATE_UNSPECIFIED",
    /**
     * The Pub/Sub subscription is properly configured.
     */
    Ok: "OK",
    /**
     * The subscription has been deleted.
     */
    SubscriptionDeleted: "SUBSCRIPTION_DELETED",
    /**
     * The topic has been deleted.
     */
    TopicDeleted: "TOPIC_DELETED",
    /**
     * Some of the subscription's field are misconfigured.
     */
    SubscriptionMisconfigured: "SUBSCRIPTION_MISCONFIGURED",
} as const;

/**
 * Potential issues with the underlying Pub/Sub subscription configuration. Only populated on get requests.
 */
export type PubsubConfigState = (typeof PubsubConfigState)[keyof typeof PubsubConfigState];

export const PullRequestFilterCommentControl = {
    /**
     * Do not require comments on Pull Requests before builds are triggered.
     */
    CommentsDisabled: "COMMENTS_DISABLED",
    /**
     * Enforce that repository owners or collaborators must comment on Pull Requests before builds are triggered.
     */
    CommentsEnabled: "COMMENTS_ENABLED",
    /**
     * Enforce that repository owners or collaborators must comment on external contributors' Pull Requests before builds are triggered.
     */
    CommentsEnabledForExternalContributorsOnly: "COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY",
} as const;

/**
 * Configure builds to run whether a repository owner or collaborator need to comment `/gcbrun`.
 */
export type PullRequestFilterCommentControl = (typeof PullRequestFilterCommentControl)[keyof typeof PullRequestFilterCommentControl];

export const TriggerEventType = {
    /**
     * EVENT_TYPE_UNSPECIFIED event_types are ignored.
     */
    EventTypeUnspecified: "EVENT_TYPE_UNSPECIFIED",
    /**
     * REPO corresponds to the supported VCS integrations.
     */
    Repo: "REPO",
    /**
     * WEBHOOK corresponds to webhook triggers.
     */
    Webhook: "WEBHOOK",
    /**
     * PUBSUB corresponds to pubsub triggers.
     */
    Pubsub: "PUBSUB",
    /**
     * MANUAL corresponds to manual-only invoked triggers.
     */
    Manual: "MANUAL",
} as const;

/**
 * EventType allows the user to explicitly set the type of event to which this BuildTrigger should respond. This field will be validated against the rest of the configuration if it is set.
 */
export type TriggerEventType = (typeof TriggerEventType)[keyof typeof TriggerEventType];

export const TriggerIncludeBuildLogs = {
    /**
     * Build logs will not be shown on GitHub.
     */
    IncludeBuildLogsUnspecified: "INCLUDE_BUILD_LOGS_UNSPECIFIED",
    /**
     * Build logs will be shown on GitHub.
     */
    IncludeBuildLogsWithStatus: "INCLUDE_BUILD_LOGS_WITH_STATUS",
} as const;

/**
 * If set to INCLUDE_BUILD_LOGS_WITH_STATUS, log url will be shown on GitHub page when build status is final. Setting this field to INCLUDE_BUILD_LOGS_WITH_STATUS for non GitHub triggers results in INVALID_ARGUMENT error.
 */
export type TriggerIncludeBuildLogs = (typeof TriggerIncludeBuildLogs)[keyof typeof TriggerIncludeBuildLogs];

export const WebhookConfigState = {
    /**
     * The webhook auth configuration not been checked.
     */
    StateUnspecified: "STATE_UNSPECIFIED",
    /**
     * The auth configuration is properly setup.
     */
    Ok: "OK",
    /**
     * The secret provided in auth_method has been deleted.
     */
    SecretDeleted: "SECRET_DELETED",
} as const;

/**
 * Potential issues with the underlying Pub/Sub subscription configuration. Only populated on get requests.
 */
export type WebhookConfigState = (typeof WebhookConfigState)[keyof typeof WebhookConfigState];
