// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Fetches a single Repository.
 */
export function getRepository(args: GetRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:dataform/v1beta1:getRepository", {
        "location": args.location,
        "project": args.project,
        "repositoryId": args.repositoryId,
    }, opts);
}

export interface GetRepositoryArgs {
    location: string;
    project?: string;
    repositoryId: string;
}

export interface GetRepositoryResult {
    /**
     * Optional. If set, configures this repository to be linked to a Git remote.
     */
    readonly gitRemoteSettings: outputs.dataform.v1beta1.GitRemoteSettingsResponse;
    /**
     * Optional. Input only. The initial commit file contents. Represented as map from file path to contents. The path is the full file path to commit including filename, from repository root.
     */
    readonly initialCommitFileContents: {[key: string]: string};
    /**
     * Optional. Input only. An optional initial commit metadata for the Repository. The Repository must not have a value for `git_remote_settings.url`.
     */
    readonly initialCommitMetadata: outputs.dataform.v1beta1.CommitMetadataResponse;
    /**
     * Optional. Repository user labels.
     */
    readonly labels: {[key: string]: string};
    /**
     * The repository's name.
     */
    readonly name: string;
    /**
     * Optional. The name of the Secret Manager secret version to be used to interpolate variables into the .npmrc file for package installation operations. Must be in the format `projects/*&#47;secrets/*&#47;versions/*`. The file itself must be in a JSON format.
     */
    readonly npmrcEnvironmentVariablesSecretVersion: string;
    /**
     * Optional. If set, fields of `workspace_compilation_overrides` override the default compilation settings that are specified in dataform.json when creating workspace-scoped compilation results. See documentation for `WorkspaceCompilationOverrides` for more information.
     */
    readonly workspaceCompilationOverrides: outputs.dataform.v1beta1.WorkspaceCompilationOverridesResponse;
}
/**
 * Fetches a single Repository.
 */
export function getRepositoryOutput(args: GetRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRepositoryResult> {
    return pulumi.output(args).apply((a: any) => getRepository(a, opts))
}

export interface GetRepositoryOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    repositoryId: pulumi.Input<string>;
}
