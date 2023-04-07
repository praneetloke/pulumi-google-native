// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Retrieves a DataAttributeBinding resource.
 */
export function getDataAttributeBinding(args: GetDataAttributeBindingArgs, opts?: pulumi.InvokeOptions): Promise<GetDataAttributeBindingResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:dataplex/v1:getDataAttributeBinding", {
        "dataAttributeBindingId": args.dataAttributeBindingId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetDataAttributeBindingArgs {
    dataAttributeBindingId: string;
    location: string;
    project?: string;
}

export interface GetDataAttributeBindingResult {
    /**
     * Optional. List of attributes to be associated with the resource, provided in the form: projects/{project}/locations/{location}/dataTaxonomies/{dataTaxonomy}/attributes/{data_attribute_id}
     */
    readonly attributes: string[];
    /**
     * The time when the DataAttributeBinding was created.
     */
    readonly createTime: string;
    /**
     * Optional. Description of the DataAttributeBinding.
     */
    readonly description: string;
    /**
     * Optional. User friendly display name.
     */
    readonly displayName: string;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Etags must be used when calling the DeleteDataAttributeBinding and the UpdateDataAttributeBinding method.
     */
    readonly etag: string;
    /**
     * Optional. User-defined labels for the DataAttributeBinding.
     */
    readonly labels: {[key: string]: string};
    /**
     * The relative resource name of the Data Attribute Binding, of the form: projects/{project_number}/locations/{location}/dataAttributeBindings/{data_attribute_binding_id}
     */
    readonly name: string;
    /**
     * Optional. The list of paths for items within the associated resource (eg. columns and partitions within a table) along with attribute bindings.
     */
    readonly paths: outputs.dataplex.v1.GoogleCloudDataplexV1DataAttributeBindingPathResponse[];
    /**
     * Optional. Immutable. The resource name of the resource that is associated to attributes. Presently, only entity resource is supported in the form: projects/{project}/locations/{location}/lakes/{lake}/zones/{zone}/entities/{entity_id} Must belong in the same project and region as the attribute binding, and there can only exist one active binding for a resource.
     */
    readonly resource: string;
    /**
     * System generated globally unique ID for the DataAttributeBinding. This ID will be different if the DataAttributeBinding is deleted and re-created with the same name.
     */
    readonly uid: string;
    /**
     * The time when the DataAttributeBinding was last updated.
     */
    readonly updateTime: string;
}
/**
 * Retrieves a DataAttributeBinding resource.
 */
export function getDataAttributeBindingOutput(args: GetDataAttributeBindingOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDataAttributeBindingResult> {
    return pulumi.output(args).apply((a: any) => getDataAttributeBinding(a, opts))
}

export interface GetDataAttributeBindingOutputArgs {
    dataAttributeBindingId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
