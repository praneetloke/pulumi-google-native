// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Starts a build with the specified configuration. This method returns a long-running `Operation`, which includes the build ID. Pass the build ID to `GetBuild` to determine the build status (such as `SUCCESS` or `FAILURE`).
 * Auto-naming is currently not supported for this resource.
 */
export class Build extends pulumi.CustomResource {
    /**
     * Get an existing Build resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Build {
        return new Build(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudbuild/v1:Build';

    /**
     * Returns true if the given object is an instance of Build.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Build {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Build.__pulumiType;
    }

    /**
     * Artifacts produced by the build that should be uploaded upon successful completion of all build steps.
     */
    public readonly artifacts!: pulumi.Output<outputs.cloudbuild.v1.ArtifactsResponse>;
    /**
     * Secrets and secret environment variables.
     */
    public readonly availableSecrets!: pulumi.Output<outputs.cloudbuild.v1.SecretsResponse>;
    /**
     * The ID of the `BuildTrigger` that triggered this build, if it was triggered automatically.
     */
    public /*out*/ readonly buildTriggerId!: pulumi.Output<string>;
    /**
     * Time at which the request to create the build was received.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Time at which execution of the build was finished. The difference between finish_time and start_time is the duration of the build's execution.
     */
    public /*out*/ readonly finishTime!: pulumi.Output<string>;
    /**
     * A list of images to be pushed upon the successful completion of all build steps. The images are pushed using the builder service account's credentials. The digests of the pushed images will be stored in the `Build` resource's results field. If any of the images fail to be pushed, the build status is marked `FAILURE`.
     */
    public readonly images!: pulumi.Output<string[]>;
    /**
     * URL to logs for this build in Google Cloud Console.
     */
    public /*out*/ readonly logUrl!: pulumi.Output<string>;
    /**
     * Google Cloud Storage bucket where logs should be written (see [Bucket Name Requirements](https://cloud.google.com/storage/docs/bucket-naming#requirements)). Logs file names will be of the format `${logs_bucket}/log-${build_id}.txt`.
     */
    public readonly logsBucket!: pulumi.Output<string>;
    /**
     * The 'Build' name with format: `projects/{project}/locations/{location}/builds/{build}`, where {build} is a unique identifier generated by the service.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Special options for this build.
     */
    public readonly options!: pulumi.Output<outputs.cloudbuild.v1.BuildOptionsResponse>;
    /**
     * ID of the project.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * TTL in queue for this build. If provided and the build is enqueued longer than this value, the build will expire and the build status will be `EXPIRED`. The TTL starts ticking from create_time.
     */
    public readonly queueTtl!: pulumi.Output<string>;
    /**
     * Results of the build.
     */
    public /*out*/ readonly results!: pulumi.Output<outputs.cloudbuild.v1.ResultsResponse>;
    /**
     * Secrets to decrypt using Cloud Key Management Service. Note: Secret Manager is the recommended technique for managing sensitive data with Cloud Build. Use `available_secrets` to configure builds to access secrets from Secret Manager. For instructions, see: https://cloud.google.com/cloud-build/docs/securing-builds/use-secrets
     */
    public readonly secrets!: pulumi.Output<outputs.cloudbuild.v1.SecretResponse[]>;
    /**
     * IAM service account whose credentials will be used at build runtime. Must be of the format `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. ACCOUNT can be email address or uniqueId of the service account. 
     */
    public readonly serviceAccount!: pulumi.Output<string>;
    /**
     * The location of the source files to build.
     */
    public readonly source!: pulumi.Output<outputs.cloudbuild.v1.SourceResponse>;
    /**
     * A permanent fixed identifier for source.
     */
    public /*out*/ readonly sourceProvenance!: pulumi.Output<outputs.cloudbuild.v1.SourceProvenanceResponse>;
    /**
     * Time at which execution of the build was started.
     */
    public /*out*/ readonly startTime!: pulumi.Output<string>;
    /**
     * Status of the build.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Customer-readable message about the current status.
     */
    public /*out*/ readonly statusDetail!: pulumi.Output<string>;
    /**
     * The operations to be performed on the workspace.
     */
    public readonly steps!: pulumi.Output<outputs.cloudbuild.v1.BuildStepResponse[]>;
    /**
     * Substitutions data for `Build` resource.
     */
    public readonly substitutions!: pulumi.Output<{[key: string]: string}>;
    /**
     * Tags for annotation of a `Build`. These are not docker tags.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * Amount of time that this build should be allowed to run, to second granularity. If this amount of time elapses, work on the build will cease and the build status will be `TIMEOUT`. `timeout` starts ticking from `startTime`. Default time is ten minutes.
     */
    public readonly timeout!: pulumi.Output<string>;
    /**
     * Stores timing information for phases of the build. Valid keys are: * BUILD: time to execute all build steps * PUSH: time to push all specified images. * FETCHSOURCE: time to fetch source. If the build does not specify source or images, these keys will not be included.
     */
    public /*out*/ readonly timing!: pulumi.Output<{[key: string]: string}>;
    /**
     * Non-fatal problems encountered during the execution of the build.
     */
    public /*out*/ readonly warnings!: pulumi.Output<outputs.cloudbuild.v1.WarningResponse[]>;

    /**
     * Create a Build resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BuildArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.steps === undefined) && !opts.urn) {
                throw new Error("Missing required property 'steps'");
            }
            inputs["artifacts"] = args ? args.artifacts : undefined;
            inputs["availableSecrets"] = args ? args.availableSecrets : undefined;
            inputs["images"] = args ? args.images : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["logsBucket"] = args ? args.logsBucket : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["queueTtl"] = args ? args.queueTtl : undefined;
            inputs["secrets"] = args ? args.secrets : undefined;
            inputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["steps"] = args ? args.steps : undefined;
            inputs["substitutions"] = args ? args.substitutions : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
            inputs["buildTriggerId"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["finishTime"] = undefined /*out*/;
            inputs["logUrl"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["results"] = undefined /*out*/;
            inputs["sourceProvenance"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusDetail"] = undefined /*out*/;
            inputs["timing"] = undefined /*out*/;
            inputs["warnings"] = undefined /*out*/;
        } else {
            inputs["artifacts"] = undefined /*out*/;
            inputs["availableSecrets"] = undefined /*out*/;
            inputs["buildTriggerId"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["finishTime"] = undefined /*out*/;
            inputs["images"] = undefined /*out*/;
            inputs["logUrl"] = undefined /*out*/;
            inputs["logsBucket"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["options"] = undefined /*out*/;
            inputs["project"] = undefined /*out*/;
            inputs["queueTtl"] = undefined /*out*/;
            inputs["results"] = undefined /*out*/;
            inputs["secrets"] = undefined /*out*/;
            inputs["serviceAccount"] = undefined /*out*/;
            inputs["source"] = undefined /*out*/;
            inputs["sourceProvenance"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusDetail"] = undefined /*out*/;
            inputs["steps"] = undefined /*out*/;
            inputs["substitutions"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["timeout"] = undefined /*out*/;
            inputs["timing"] = undefined /*out*/;
            inputs["warnings"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Build.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Build resource.
 */
export interface BuildArgs {
    /**
     * Artifacts produced by the build that should be uploaded upon successful completion of all build steps.
     */
    artifacts?: pulumi.Input<inputs.cloudbuild.v1.ArtifactsArgs>;
    /**
     * Secrets and secret environment variables.
     */
    availableSecrets?: pulumi.Input<inputs.cloudbuild.v1.SecretsArgs>;
    /**
     * A list of images to be pushed upon the successful completion of all build steps. The images are pushed using the builder service account's credentials. The digests of the pushed images will be stored in the `Build` resource's results field. If any of the images fail to be pushed, the build status is marked `FAILURE`.
     */
    images?: pulumi.Input<pulumi.Input<string>[]>;
    location?: pulumi.Input<string>;
    /**
     * Google Cloud Storage bucket where logs should be written (see [Bucket Name Requirements](https://cloud.google.com/storage/docs/bucket-naming#requirements)). Logs file names will be of the format `${logs_bucket}/log-${build_id}.txt`.
     */
    logsBucket?: pulumi.Input<string>;
    /**
     * Special options for this build.
     */
    options?: pulumi.Input<inputs.cloudbuild.v1.BuildOptionsArgs>;
    project?: pulumi.Input<string>;
    projectId: pulumi.Input<string>;
    /**
     * TTL in queue for this build. If provided and the build is enqueued longer than this value, the build will expire and the build status will be `EXPIRED`. The TTL starts ticking from create_time.
     */
    queueTtl?: pulumi.Input<string>;
    /**
     * Secrets to decrypt using Cloud Key Management Service. Note: Secret Manager is the recommended technique for managing sensitive data with Cloud Build. Use `available_secrets` to configure builds to access secrets from Secret Manager. For instructions, see: https://cloud.google.com/cloud-build/docs/securing-builds/use-secrets
     */
    secrets?: pulumi.Input<pulumi.Input<inputs.cloudbuild.v1.SecretArgs>[]>;
    /**
     * IAM service account whose credentials will be used at build runtime. Must be of the format `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. ACCOUNT can be email address or uniqueId of the service account. 
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * The location of the source files to build.
     */
    source?: pulumi.Input<inputs.cloudbuild.v1.SourceArgs>;
    /**
     * The operations to be performed on the workspace.
     */
    steps: pulumi.Input<pulumi.Input<inputs.cloudbuild.v1.BuildStepArgs>[]>;
    /**
     * Substitutions data for `Build` resource.
     */
    substitutions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Tags for annotation of a `Build`. These are not docker tags.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Amount of time that this build should be allowed to run, to second granularity. If this amount of time elapses, work on the build will cease and the build status will be `TIMEOUT`. `timeout` starts ticking from `startTime`. Default time is ten minutes.
     */
    timeout?: pulumi.Input<string>;
}
