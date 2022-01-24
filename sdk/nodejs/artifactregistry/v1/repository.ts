// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a repository. The returned Operation will finish once the repository has been created. Its response will be the created Repository.
 */
export class Repository extends pulumi.CustomResource {
    /**
     * Get an existing Repository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Repository {
        return new Repository(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:artifactregistry/v1:Repository';

    /**
     * Returns true if the given object is an instance of Repository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Repository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Repository.__pulumiType;
    }

    /**
     * The time when the repository was created.
     */
    public readonly createTime!: pulumi.Output<string>;
    /**
     * The user-provided description of the repository.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The format of packages that are stored in the repository.
     */
    public readonly format!: pulumi.Output<string>;
    /**
     * The Cloud KMS resource name of the customer managed encryption key that’s used to encrypt the contents of the Repository. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. This value may not be changed after the Repository has been created.
     */
    public readonly kmsKeyName!: pulumi.Output<string>;
    /**
     * Labels with user-defined metadata. This field may contain up to 64 entries. Label keys and values may be no longer than 63 characters. Label keys must begin with a lowercase letter and may only contain lowercase letters, numeric characters, underscores, and dashes.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Maven repository config contains repository level configuration for the repositories of maven type.
     */
    public readonly mavenConfig!: pulumi.Output<outputs.artifactregistry.v1.MavenRepositoryConfigResponse>;
    /**
     * The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1".
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The time when the repository was last updated.
     */
    public readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Repository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RepositoryArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["createTime"] = args ? args.createTime : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["kmsKeyName"] = args ? args.kmsKeyName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["mavenConfig"] = args ? args.mavenConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["repositoryId"] = args ? args.repositoryId : undefined;
            resourceInputs["updateTime"] = args ? args.updateTime : undefined;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["format"] = undefined /*out*/;
            resourceInputs["kmsKeyName"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["mavenConfig"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Repository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Repository resource.
 */
export interface RepositoryArgs {
    /**
     * The time when the repository was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The user-provided description of the repository.
     */
    description?: pulumi.Input<string>;
    /**
     * The format of packages that are stored in the repository.
     */
    format?: pulumi.Input<enums.artifactregistry.v1.RepositoryFormat>;
    /**
     * The Cloud KMS resource name of the customer managed encryption key that’s used to encrypt the contents of the Repository. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. This value may not be changed after the Repository has been created.
     */
    kmsKeyName?: pulumi.Input<string>;
    /**
     * Labels with user-defined metadata. This field may contain up to 64 entries. Label keys and values may be no longer than 63 characters. Label keys must begin with a lowercase letter and may only contain lowercase letters, numeric characters, underscores, and dashes.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Maven repository config contains repository level configuration for the repositories of maven type.
     */
    mavenConfig?: pulumi.Input<inputs.artifactregistry.v1.MavenRepositoryConfigArgs>;
    /**
     * The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1".
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    repositoryId?: pulumi.Input<string>;
    /**
     * The time when the repository was last updated.
     */
    updateTime?: pulumi.Input<string>;
}
