// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Retrieves a `GitLabConfig`. This API is experimental
 */
export function getGitLabConfig(args: GetGitLabConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetGitLabConfigResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:cloudbuild/v1:getGitLabConfig", {
        "gitLabConfigId": args.gitLabConfigId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetGitLabConfigArgs {
    gitLabConfigId: string;
    location: string;
    project?: string;
}

export interface GetGitLabConfigResult {
    /**
     * Connected GitLab.com or GitLabEnterprise repositories for this config.
     */
    readonly connectedRepositories: outputs.cloudbuild.v1.GitLabRepositoryIdResponse[];
    /**
     * Time when the config was created.
     */
    readonly createTime: string;
    /**
     * Optional. GitLabEnterprise config.
     */
    readonly enterpriseConfig: outputs.cloudbuild.v1.GitLabEnterpriseConfigResponse;
    /**
     * The resource name for the config.
     */
    readonly name: string;
    /**
     * Secret Manager secrets needed by the config.
     */
    readonly secrets: outputs.cloudbuild.v1.GitLabSecretsResponse;
    /**
     * Username of the GitLab.com or GitLab Enterprise account Cloud Build will use.
     */
    readonly username: string;
    /**
     * UUID included in webhook requests. The UUID is used to look up the corresponding config.
     */
    readonly webhookKey: string;
}

export function getGitLabConfigOutput(args: GetGitLabConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGitLabConfigResult> {
    return pulumi.output(args).apply(a => getGitLabConfig(a, opts))
}

export interface GetGitLabConfigOutputArgs {
    gitLabConfigId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
