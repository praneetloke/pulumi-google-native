// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a FHIR resource. Implements the FHIR standard create interaction ([DSTU2](http://hl7.org/implement/standards/fhir/DSTU2/http.html#create), [STU3](http://hl7.org/implement/standards/fhir/STU3/http.html#create), [R4](http://hl7.org/implement/standards/fhir/R4/http.html#create)), which creates a new resource with a server-assigned resource ID. The request body must contain a JSON-encoded FHIR resource, and the request headers must contain `Content-Type: application/fhir+json`. On success, the response body contains a JSON-encoded representation of the resource as it was created on the server, including the server-assigned resource ID and version ID. Errors generated by the FHIR store contain a JSON-encoded `OperationOutcome` resource describing the reason for the error. If the request cannot be mapped to a valid API method on a FHIR store, a generic GCP error might be returned instead. For samples that show how to call `create`, see [Creating a FHIR resource](/healthcare/docs/how-tos/fhir-resources#creating_a_fhir_resource).
 */
export class DatasetFhirStoreFhir extends pulumi.CustomResource {
    /**
     * Get an existing DatasetFhirStoreFhir resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatasetFhirStoreFhir {
        return new DatasetFhirStoreFhir(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:healthcare/v1:DatasetFhirStoreFhir';

    /**
     * Returns true if the given object is an instance of DatasetFhirStoreFhir.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatasetFhirStoreFhir {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatasetFhirStoreFhir.__pulumiType;
    }

    /**
     * The HTTP Content-Type header value specifying the content type of the body.
     */
    public readonly contentType!: pulumi.Output<string>;
    /**
     * The HTTP request/response body as raw binary.
     */
    public readonly data!: pulumi.Output<string>;
    /**
     * Application specific response metadata. Must be set in the first response for streaming APIs.
     */
    public readonly extensions!: pulumi.Output<{[key: string]: string}[]>;

    /**
     * Create a DatasetFhirStoreFhir resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetFhirStoreFhirArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.datasetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasetId'");
            }
            if ((!args || args.fhirId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fhirId'");
            }
            if ((!args || args.fhirId1 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fhirId1'");
            }
            if ((!args || args.fhirStoreId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fhirStoreId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["contentType"] = args ? args.contentType : undefined;
            inputs["data"] = args ? args.data : undefined;
            inputs["datasetId"] = args ? args.datasetId : undefined;
            inputs["extensions"] = args ? args.extensions : undefined;
            inputs["fhirId"] = args ? args.fhirId : undefined;
            inputs["fhirId1"] = args ? args.fhirId1 : undefined;
            inputs["fhirStoreId"] = args ? args.fhirStoreId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
        } else {
            inputs["contentType"] = undefined /*out*/;
            inputs["data"] = undefined /*out*/;
            inputs["extensions"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DatasetFhirStoreFhir.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatasetFhirStoreFhir resource.
 */
export interface DatasetFhirStoreFhirArgs {
    /**
     * The HTTP Content-Type header value specifying the content type of the body.
     */
    readonly contentType?: pulumi.Input<string>;
    /**
     * The HTTP request/response body as raw binary.
     */
    readonly data?: pulumi.Input<string>;
    readonly datasetId: pulumi.Input<string>;
    /**
     * Application specific response metadata. Must be set in the first response for streaming APIs.
     */
    readonly extensions?: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    readonly fhirId: pulumi.Input<string>;
    readonly fhirId1: pulumi.Input<string>;
    readonly fhirStoreId: pulumi.Input<string>;
    readonly location: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
}
