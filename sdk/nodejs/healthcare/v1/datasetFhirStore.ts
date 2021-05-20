// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new FHIR store within the parent dataset.
 */
export class DatasetFhirStore extends pulumi.CustomResource {
    /**
     * Get an existing DatasetFhirStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatasetFhirStore {
        return new DatasetFhirStore(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:healthcare/v1:DatasetFhirStore';

    /**
     * Returns true if the given object is an instance of DatasetFhirStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatasetFhirStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatasetFhirStore.__pulumiType;
    }

    /**
     * If true, overrides the default search behavior for this FHIR store to `handling=strict` which returns an error for unrecognized search parameters. If false, uses the FHIR specification default `handling=lenient` which ignores unrecognized search parameters. The handling can always be changed from the default on an individual API call by setting the HTTP header `Prefer: handling=strict` or `Prefer: handling=lenient`.
     */
    public readonly defaultSearchHandlingStrict!: pulumi.Output<boolean>;
    /**
     * Immutable. Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The default value is false, meaning that the API enforces referential integrity and fails the requests that result in inconsistent state in the FHIR store. When this field is set to true, the API skips referential integrity checks. Consequently, operations that rely on references, such as GetPatientEverything, do not return all the results if broken references exist.
     */
    public readonly disableReferentialIntegrity!: pulumi.Output<boolean>;
    /**
     * Immutable. Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR store. If set to false, which is the default behavior, all write operations cause historical versions to be recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set to true, no historical versions are kept. The server sends errors for attempts to read the historical versions.
     */
    public readonly disableResourceVersioning!: pulumi.Output<boolean>;
    /**
     * Whether this FHIR store has the [updateCreate capability](https://www.hl7.org/fhir/capabilitystatement-definitions.html#CapabilityStatement.rest.resource.updateCreate). This determines if the client can use an Update operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and attempts to update a non-existent resource return errors. It is strongly advised not to include or encode any sensitive data such as patient identifiers in client-specified resource IDs. Those IDs are part of the FHIR resource path recorded in Cloud audit logs and Pub/Sub notifications. Those IDs can also be contained in reference fields within other resources.
     */
    public readonly enableUpdateCreate!: pulumi.Output<boolean>;
    /**
     * User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource name of the FHIR store, of the form `projects/{project_id}/datasets/{dataset_id}/fhirStores/{fhir_store_id}`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource".
     */
    public readonly notificationConfig!: pulumi.Output<outputs.healthcare.v1.NotificationConfigResponse>;
    /**
     * A list of streaming configs that configure the destinations of streaming export for every resource mutation in this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next resource mutation is streamed to the new location in addition to the existing ones. When a location is removed from the list, the server stops streaming to that location. Before adding a new config, you must add the required [`bigquery.dataEditor`](https://cloud.google.com/bigquery/docs/access-control#bigquery.dataEditor) role to your project's **Cloud Healthcare Service Agent** [service account](https://cloud.google.com/iam/docs/service-accounts). Some lag (typically on the order of dozens of seconds) is expected before the results show up in the streaming destination.
     */
    public readonly streamConfigs!: pulumi.Output<outputs.healthcare.v1.StreamConfigResponse[]>;
    /**
     * Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a DatasetFhirStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetFhirStoreArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.datasetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasetId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["datasetId"] = args ? args.datasetId : undefined;
            inputs["defaultSearchHandlingStrict"] = args ? args.defaultSearchHandlingStrict : undefined;
            inputs["disableReferentialIntegrity"] = args ? args.disableReferentialIntegrity : undefined;
            inputs["disableResourceVersioning"] = args ? args.disableResourceVersioning : undefined;
            inputs["enableUpdateCreate"] = args ? args.enableUpdateCreate : undefined;
            inputs["fhirStoreId"] = args ? args.fhirStoreId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notificationConfig"] = args ? args.notificationConfig : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["streamConfigs"] = args ? args.streamConfigs : undefined;
            inputs["version"] = args ? args.version : undefined;
        } else {
            inputs["defaultSearchHandlingStrict"] = undefined /*out*/;
            inputs["disableReferentialIntegrity"] = undefined /*out*/;
            inputs["disableResourceVersioning"] = undefined /*out*/;
            inputs["enableUpdateCreate"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["notificationConfig"] = undefined /*out*/;
            inputs["streamConfigs"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DatasetFhirStore.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatasetFhirStore resource.
 */
export interface DatasetFhirStoreArgs {
    readonly datasetId: pulumi.Input<string>;
    /**
     * If true, overrides the default search behavior for this FHIR store to `handling=strict` which returns an error for unrecognized search parameters. If false, uses the FHIR specification default `handling=lenient` which ignores unrecognized search parameters. The handling can always be changed from the default on an individual API call by setting the HTTP header `Prefer: handling=strict` or `Prefer: handling=lenient`.
     */
    readonly defaultSearchHandlingStrict?: pulumi.Input<boolean>;
    /**
     * Immutable. Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The default value is false, meaning that the API enforces referential integrity and fails the requests that result in inconsistent state in the FHIR store. When this field is set to true, the API skips referential integrity checks. Consequently, operations that rely on references, such as GetPatientEverything, do not return all the results if broken references exist.
     */
    readonly disableReferentialIntegrity?: pulumi.Input<boolean>;
    /**
     * Immutable. Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR store. If set to false, which is the default behavior, all write operations cause historical versions to be recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set to true, no historical versions are kept. The server sends errors for attempts to read the historical versions.
     */
    readonly disableResourceVersioning?: pulumi.Input<boolean>;
    /**
     * Whether this FHIR store has the [updateCreate capability](https://www.hl7.org/fhir/capabilitystatement-definitions.html#CapabilityStatement.rest.resource.updateCreate). This determines if the client can use an Update operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and attempts to update a non-existent resource return errors. It is strongly advised not to include or encode any sensitive data such as patient identifiers in client-specified resource IDs. Those IDs are part of the FHIR resource path recorded in Cloud audit logs and Pub/Sub notifications. Those IDs can also be contained in reference fields within other resources.
     */
    readonly enableUpdateCreate?: pulumi.Input<boolean>;
    readonly fhirStoreId?: pulumi.Input<string>;
    /**
     * User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly location: pulumi.Input<string>;
    /**
     * Resource name of the FHIR store, of the form `projects/{project_id}/datasets/{dataset_id}/fhirStores/{fhir_store_id}`.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource".
     */
    readonly notificationConfig?: pulumi.Input<inputs.healthcare.v1.NotificationConfigArgs>;
    readonly project: pulumi.Input<string>;
    /**
     * A list of streaming configs that configure the destinations of streaming export for every resource mutation in this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next resource mutation is streamed to the new location in addition to the existing ones. When a location is removed from the list, the server stops streaming to that location. Before adding a new config, you must add the required [`bigquery.dataEditor`](https://cloud.google.com/bigquery/docs/access-control#bigquery.dataEditor) role to your project's **Cloud Healthcare Service Agent** [service account](https://cloud.google.com/iam/docs/service-accounts). Some lag (typically on the order of dozens of seconds) is expected before the results show up in the streaming destination.
     */
    readonly streamConfigs?: pulumi.Input<pulumi.Input<inputs.healthcare.v1.StreamConfigArgs>[]>;
    /**
     * Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store.
     */
    readonly version?: pulumi.Input<string>;
}
