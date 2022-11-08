// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a repository.
 */
export function getRepository(args: GetRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:artifactregistry/v1:getRepository", {
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
     * The time when the repository was created.
     */
    readonly createTime: string;
    /**
     * The user-provided description of the repository.
     */
    readonly description: string;
    /**
     * The format of packages that are stored in the repository.
     */
    readonly format: string;
    /**
     * The Cloud KMS resource name of the customer managed encryption key that's used to encrypt the contents of the Repository. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. This value may not be changed after the Repository has been created.
     */
    readonly kmsKeyName: string;
    /**
     * Labels with user-defined metadata. This field may contain up to 64 entries. Label keys and values may be no longer than 63 characters. Label keys must begin with a lowercase letter and may only contain lowercase letters, numeric characters, underscores, and dashes.
     */
    readonly labels: {[key: string]: string};
    /**
     * Maven repository config contains repository level configuration for the repositories of maven type.
     */
    readonly mavenConfig: outputs.artifactregistry.v1.MavenRepositoryConfigResponse;
    /**
     * The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1".
     */
    readonly name: string;
    /**
     * If set, the repository satisfies physical zone separation.
     */
    readonly satisfiesPzs: boolean;
    /**
     * The size, in bytes, of all artifact storage in this repository. Repositories that are generally available or in public preview use this to calculate storage costs.
     */
    readonly sizeBytes: string;
    /**
     * The time when the repository was last updated.
     */
    readonly updateTime: string;
}

export function getRepositoryOutput(args: GetRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRepositoryResult> {
    return pulumi.output(args).apply(a => getRepository(a, opts))
}

export interface GetRepositoryOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    repositoryId: pulumi.Input<string>;
}
