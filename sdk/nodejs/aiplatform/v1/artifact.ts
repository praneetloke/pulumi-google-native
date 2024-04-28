// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates an Artifact associated with a MetadataStore.
 * Auto-naming is currently not supported for this resource.
 */
export class Artifact extends pulumi.CustomResource {
    /**
     * Get an existing Artifact resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Artifact {
        return new Artifact(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:aiplatform/v1:Artifact';

    /**
     * Returns true if the given object is an instance of Artifact.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Artifact {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Artifact.__pulumiType;
    }

    /**
     * The {artifact} portion of the resource name with the format: `projects/{project}/locations/{location}/metadataStores/{metadatastore}/artifacts/{artifact}` If not provided, the Artifact's ID will be a UUID generated by the service. Must be 4-128 characters in length. Valid characters are `/a-z-/`. Must be unique across all Artifacts in the parent MetadataStore. (Otherwise the request will fail with ALREADY_EXISTS, or PERMISSION_DENIED if the caller can't view the preexisting Artifact.)
     */
    public readonly artifactId!: pulumi.Output<string | undefined>;
    /**
     * Timestamp when this Artifact was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Description of the Artifact
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * User provided display name of the Artifact. May be up to 128 Unicode characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * An eTag used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * The labels with user-defined metadata to organize your Artifacts. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one Artifact (System labels are excluded).
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Properties of the Artifact. Top level metadata keys' heading and trailing spaces will be trimmed. The size of this field should not exceed 200KB.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: any}>;
    public readonly metadataStoreId!: pulumi.Output<string>;
    /**
     * The resource name of the Artifact.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The title of the schema describing the metadata. Schema title and version is expected to be registered in earlier Create Schema calls. And both are used together as unique identifiers to identify schemas within the local metadata store.
     */
    public readonly schemaTitle!: pulumi.Output<string>;
    /**
     * The version of the schema in schema_name to use. Schema title and version is expected to be registered in earlier Create Schema calls. And both are used together as unique identifiers to identify schemas within the local metadata store.
     */
    public readonly schemaVersion!: pulumi.Output<string>;
    /**
     * The state of this Artifact. This is a property of the Artifact, and does not imply or capture any ongoing process. This property is managed by clients (such as Vertex AI Pipelines), and the system does not prescribe or check the validity of state transitions.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * Timestamp when this Artifact was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The uniform resource identifier of the artifact file. May be empty if there is no actual artifact file.
     */
    public readonly uri!: pulumi.Output<string>;

    /**
     * Create a Artifact resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ArtifactArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.metadataStoreId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metadataStoreId'");
            }
            resourceInputs["artifactId"] = args ? args.artifactId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["metadataStoreId"] = args ? args.metadataStoreId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["schemaTitle"] = args ? args.schemaTitle : undefined;
            resourceInputs["schemaVersion"] = args ? args.schemaVersion : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["uri"] = args ? args.uri : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["artifactId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["metadataStoreId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["schemaTitle"] = undefined /*out*/;
            resourceInputs["schemaVersion"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["uri"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "metadataStoreId", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Artifact.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Artifact resource.
 */
export interface ArtifactArgs {
    /**
     * The {artifact} portion of the resource name with the format: `projects/{project}/locations/{location}/metadataStores/{metadatastore}/artifacts/{artifact}` If not provided, the Artifact's ID will be a UUID generated by the service. Must be 4-128 characters in length. Valid characters are `/a-z-/`. Must be unique across all Artifacts in the parent MetadataStore. (Otherwise the request will fail with ALREADY_EXISTS, or PERMISSION_DENIED if the caller can't view the preexisting Artifact.)
     */
    artifactId?: pulumi.Input<string>;
    /**
     * Description of the Artifact
     */
    description?: pulumi.Input<string>;
    /**
     * User provided display name of the Artifact. May be up to 128 Unicode characters.
     */
    displayName?: pulumi.Input<string>;
    /**
     * An eTag used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
     */
    etag?: pulumi.Input<string>;
    /**
     * The labels with user-defined metadata to organize your Artifacts. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one Artifact (System labels are excluded).
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Properties of the Artifact. Top level metadata keys' heading and trailing spaces will be trimmed. The size of this field should not exceed 200KB.
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    metadataStoreId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The title of the schema describing the metadata. Schema title and version is expected to be registered in earlier Create Schema calls. And both are used together as unique identifiers to identify schemas within the local metadata store.
     */
    schemaTitle?: pulumi.Input<string>;
    /**
     * The version of the schema in schema_name to use. Schema title and version is expected to be registered in earlier Create Schema calls. And both are used together as unique identifiers to identify schemas within the local metadata store.
     */
    schemaVersion?: pulumi.Input<string>;
    /**
     * The state of this Artifact. This is a property of the Artifact, and does not imply or capture any ongoing process. This property is managed by clients (such as Vertex AI Pipelines), and the system does not prescribe or check the validity of state transitions.
     */
    state?: pulumi.Input<enums.aiplatform.v1.ArtifactState>;
    /**
     * The uniform resource identifier of the artifact file. May be empty if there is no actual artifact file.
     */
    uri?: pulumi.Input<string>;
}
