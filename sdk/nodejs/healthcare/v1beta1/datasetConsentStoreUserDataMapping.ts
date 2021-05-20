// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new User data mapping in the parent consent store.
 */
export class DatasetConsentStoreUserDataMapping extends pulumi.CustomResource {
    /**
     * Get an existing DatasetConsentStoreUserDataMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatasetConsentStoreUserDataMapping {
        return new DatasetConsentStoreUserDataMapping(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:healthcare/v1beta1:DatasetConsentStoreUserDataMapping';

    /**
     * Returns true if the given object is an instance of DatasetConsentStoreUserDataMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatasetConsentStoreUserDataMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatasetConsentStoreUserDataMapping.__pulumiType;
    }

    /**
     * Indicates the time when this mapping was archived.
     */
    public /*out*/ readonly archiveTime!: pulumi.Output<string>;
    /**
     * Indicates whether this mapping is archived.
     */
    public /*out*/ readonly archived!: pulumi.Output<boolean>;
    /**
     * Required. A unique identifier for the mapped resource.
     */
    public readonly dataId!: pulumi.Output<string>;
    /**
     * Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
     */
    public readonly resourceAttributes!: pulumi.Output<outputs.healthcare.v1beta1.AttributeResponse[]>;
    /**
     * Required. User's UUID provided by the client.
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a DatasetConsentStoreUserDataMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetConsentStoreUserDataMappingArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.consentStoreId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'consentStoreId'");
            }
            if ((!args || args.datasetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasetId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.userDataMappingId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userDataMappingId'");
            }
            inputs["consentStoreId"] = args ? args.consentStoreId : undefined;
            inputs["dataId"] = args ? args.dataId : undefined;
            inputs["datasetId"] = args ? args.datasetId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["resourceAttributes"] = args ? args.resourceAttributes : undefined;
            inputs["userDataMappingId"] = args ? args.userDataMappingId : undefined;
            inputs["userId"] = args ? args.userId : undefined;
            inputs["archiveTime"] = undefined /*out*/;
            inputs["archived"] = undefined /*out*/;
        } else {
            inputs["archiveTime"] = undefined /*out*/;
            inputs["archived"] = undefined /*out*/;
            inputs["dataId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["resourceAttributes"] = undefined /*out*/;
            inputs["userId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DatasetConsentStoreUserDataMapping.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatasetConsentStoreUserDataMapping resource.
 */
export interface DatasetConsentStoreUserDataMappingArgs {
    readonly consentStoreId: pulumi.Input<string>;
    /**
     * Required. A unique identifier for the mapped resource.
     */
    readonly dataId?: pulumi.Input<string>;
    readonly datasetId: pulumi.Input<string>;
    readonly location: pulumi.Input<string>;
    /**
     * Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
     */
    readonly resourceAttributes?: pulumi.Input<pulumi.Input<inputs.healthcare.v1beta1.AttributeArgs>[]>;
    readonly userDataMappingId: pulumi.Input<string>;
    /**
     * Required. User's UUID provided by the client.
     */
    readonly userId?: pulumi.Input<string>;
}
