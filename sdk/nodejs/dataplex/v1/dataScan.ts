// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a DataScan resource.
 * Auto-naming is currently not supported for this resource.
 */
export class DataScan extends pulumi.CustomResource {
    /**
     * Get an existing DataScan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataScan {
        return new DataScan(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dataplex/v1:DataScan';

    /**
     * Returns true if the given object is an instance of DataScan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataScan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataScan.__pulumiType;
    }

    /**
     * The time when the scan was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The data source for DataScan.
     */
    public readonly data!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1DataSourceResponse>;
    /**
     * The result of the data profile scan.
     */
    public /*out*/ readonly dataProfileResult!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1DataProfileResultResponse>;
    /**
     * DataProfileScan related setting.
     */
    public readonly dataProfileSpec!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1DataProfileSpecResponse>;
    /**
     * The result of the data quality scan.
     */
    public /*out*/ readonly dataQualityResult!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1DataQualityResultResponse>;
    /**
     * DataQualityScan related setting.
     */
    public readonly dataQualitySpec!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1DataQualitySpecResponse>;
    /**
     * Required. DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter. Must be between 1-63 characters. Must be unique within the customer project / location.
     */
    public readonly dataScanId!: pulumi.Output<string>;
    /**
     * Optional. Description of the scan. Must be between 1-1024 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. User friendly display name. Must be between 1-256 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Optional. DataScan execution settings.If not specified, the fields in it will use their default values.
     */
    public readonly executionSpec!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1DataScanExecutionSpecResponse>;
    /**
     * Status of the data scan execution.
     */
    public /*out*/ readonly executionStatus!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1DataScanExecutionStatusResponse>;
    /**
     * Optional. User-defined labels for the scan.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The relative resource name of the scan, of the form: projects/{project}/locations/{location_id}/dataScans/{datascan_id}, where project refers to a project_id or project_number and location_id refers to a GCP region.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Current state of the DataScan.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The type of DataScan.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * System generated globally unique ID for the scan. This ID will be different if the scan is deleted and re-created with the same name.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * The time when the scan was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a DataScan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataScanArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.data === undefined) && !opts.urn) {
                throw new Error("Missing required property 'data'");
            }
            if ((!args || args.dataScanId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataScanId'");
            }
            resourceInputs["data"] = args ? args.data : undefined;
            resourceInputs["dataProfileSpec"] = args ? args.dataProfileSpec : undefined;
            resourceInputs["dataQualitySpec"] = args ? args.dataQualitySpec : undefined;
            resourceInputs["dataScanId"] = args ? args.dataScanId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["executionSpec"] = args ? args.executionSpec : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["dataProfileResult"] = undefined /*out*/;
            resourceInputs["dataQualityResult"] = undefined /*out*/;
            resourceInputs["executionStatus"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["data"] = undefined /*out*/;
            resourceInputs["dataProfileResult"] = undefined /*out*/;
            resourceInputs["dataProfileSpec"] = undefined /*out*/;
            resourceInputs["dataQualityResult"] = undefined /*out*/;
            resourceInputs["dataQualitySpec"] = undefined /*out*/;
            resourceInputs["dataScanId"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["executionSpec"] = undefined /*out*/;
            resourceInputs["executionStatus"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["dataScanId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(DataScan.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataScan resource.
 */
export interface DataScanArgs {
    /**
     * The data source for DataScan.
     */
    data: pulumi.Input<inputs.dataplex.v1.GoogleCloudDataplexV1DataSourceArgs>;
    /**
     * DataProfileScan related setting.
     */
    dataProfileSpec?: pulumi.Input<inputs.dataplex.v1.GoogleCloudDataplexV1DataProfileSpecArgs>;
    /**
     * DataQualityScan related setting.
     */
    dataQualitySpec?: pulumi.Input<inputs.dataplex.v1.GoogleCloudDataplexV1DataQualitySpecArgs>;
    /**
     * Required. DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter. Must be between 1-63 characters. Must be unique within the customer project / location.
     */
    dataScanId: pulumi.Input<string>;
    /**
     * Optional. Description of the scan. Must be between 1-1024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional. User friendly display name. Must be between 1-256 characters.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Optional. DataScan execution settings.If not specified, the fields in it will use their default values.
     */
    executionSpec?: pulumi.Input<inputs.dataplex.v1.GoogleCloudDataplexV1DataScanExecutionSpecArgs>;
    /**
     * Optional. User-defined labels for the scan.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
