// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new document. Note: The `projects.agent.knowledgeBases.documents` resource is deprecated; only use `projects.knowledgeBases.documents`.
 */
export class KnowledgeBaseDocument extends pulumi.CustomResource {
    /**
     * Get an existing KnowledgeBaseDocument resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): KnowledgeBaseDocument {
        return new KnowledgeBaseDocument(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v2beta1:KnowledgeBaseDocument';

    /**
     * Returns true if the given object is an instance of KnowledgeBaseDocument.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KnowledgeBaseDocument {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KnowledgeBaseDocument.__pulumiType;
    }

    /**
     * The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types. Note: This field is in the process of being deprecated, please use raw_content instead.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * The URI where the file content is located. For documents stored in Google Cloud Storage, these URIs must have the form `gs:///`. NOTE: External URLs must correspond to public webpages, i.e., they must be indexed by Google Search. In particular, URLs for showing documents in Google Cloud Storage (i.e. the URL in your browser) are not supported. Instead use the `gs://` format URI described above.
     */
    public readonly contentUri!: pulumi.Output<string>;
    /**
     * Required. The display name of the document. The name must be 1024 bytes or less; otherwise, the creation request fails.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Optional. If true, we try to automatically reload the document every day (at a time picked by the system). If false or unspecified, we don't try to automatically reload the document. Currently you can only enable automatic reload for documents sourced from a public url, see `source` field for the source types. Reload status can be tracked in `latest_reload_status`. If a reload fails, we will keep the document unchanged. If a reload fails with internal errors, the system will try to reload the document on the next day. If a reload fails with non-retriable errors (e.g. PERMISION_DENIED), the system will not try to reload the document anymore. You need to manually reload the document successfully by calling `ReloadDocument` and clear the errors.
     */
    public readonly enableAutoReload!: pulumi.Output<boolean>;
    /**
     * Required. The knowledge type of document content.
     */
    public readonly knowledgeTypes!: pulumi.Output<string[]>;
    /**
     * The time and status of the latest reload. This reload may have been triggered automatically or manually and may not have succeeded.
     */
    public /*out*/ readonly latestReloadStatus!: pulumi.Output<outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1DocumentReloadStatusResponse>;
    /**
     * Optional. Metadata for the document. The metadata supports arbitrary key-value pairs. Suggested use cases include storing a document's title, an external URL distinct from the document's content_uri, etc. The max size of a `key` or a `value` of the metadata is 1024 bytes.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * Required. The MIME type of this document.
     */
    public readonly mimeType!: pulumi.Output<string>;
    /**
     * Optional. The document resource name. The name must be empty when creating a document. Format: `projects//locations//knowledgeBases//documents/`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types.
     */
    public readonly rawContent!: pulumi.Output<string>;

    /**
     * Create a KnowledgeBaseDocument resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KnowledgeBaseDocumentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.documentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'documentId'");
            }
            if ((!args || args.knowledgeBaseId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'knowledgeBaseId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["content"] = args ? args.content : undefined;
            inputs["contentUri"] = args ? args.contentUri : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["documentId"] = args ? args.documentId : undefined;
            inputs["enableAutoReload"] = args ? args.enableAutoReload : undefined;
            inputs["importGcsCustomMetadata"] = args ? args.importGcsCustomMetadata : undefined;
            inputs["knowledgeBaseId"] = args ? args.knowledgeBaseId : undefined;
            inputs["knowledgeTypes"] = args ? args.knowledgeTypes : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["mimeType"] = args ? args.mimeType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["rawContent"] = args ? args.rawContent : undefined;
            inputs["latestReloadStatus"] = undefined /*out*/;
        } else {
            inputs["content"] = undefined /*out*/;
            inputs["contentUri"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["enableAutoReload"] = undefined /*out*/;
            inputs["knowledgeTypes"] = undefined /*out*/;
            inputs["latestReloadStatus"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["mimeType"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["rawContent"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(KnowledgeBaseDocument.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a KnowledgeBaseDocument resource.
 */
export interface KnowledgeBaseDocumentArgs {
    /**
     * The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types. Note: This field is in the process of being deprecated, please use raw_content instead.
     */
    readonly content?: pulumi.Input<string>;
    /**
     * The URI where the file content is located. For documents stored in Google Cloud Storage, these URIs must have the form `gs:///`. NOTE: External URLs must correspond to public webpages, i.e., they must be indexed by Google Search. In particular, URLs for showing documents in Google Cloud Storage (i.e. the URL in your browser) are not supported. Instead use the `gs://` format URI described above.
     */
    readonly contentUri?: pulumi.Input<string>;
    /**
     * Required. The display name of the document. The name must be 1024 bytes or less; otherwise, the creation request fails.
     */
    readonly displayName?: pulumi.Input<string>;
    readonly documentId: pulumi.Input<string>;
    /**
     * Optional. If true, we try to automatically reload the document every day (at a time picked by the system). If false or unspecified, we don't try to automatically reload the document. Currently you can only enable automatic reload for documents sourced from a public url, see `source` field for the source types. Reload status can be tracked in `latest_reload_status`. If a reload fails, we will keep the document unchanged. If a reload fails with internal errors, the system will try to reload the document on the next day. If a reload fails with non-retriable errors (e.g. PERMISION_DENIED), the system will not try to reload the document anymore. You need to manually reload the document successfully by calling `ReloadDocument` and clear the errors.
     */
    readonly enableAutoReload?: pulumi.Input<boolean>;
    readonly importGcsCustomMetadata?: pulumi.Input<string>;
    readonly knowledgeBaseId: pulumi.Input<string>;
    /**
     * Required. The knowledge type of document content.
     */
    readonly knowledgeTypes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly location: pulumi.Input<string>;
    /**
     * Optional. Metadata for the document. The metadata supports arbitrary key-value pairs. Suggested use cases include storing a document's title, an external URL distinct from the document's content_uri, etc. The max size of a `key` or a `value` of the metadata is 1024 bytes.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Required. The MIME type of this document.
     */
    readonly mimeType?: pulumi.Input<string>;
    /**
     * Optional. The document resource name. The name must be empty when creating a document. Format: `projects//locations//knowledgeBases//documents/`.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types.
     */
    readonly rawContent?: pulumi.Input<string>;
}
