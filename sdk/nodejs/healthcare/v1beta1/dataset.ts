// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new health dataset. Results are returned through the Operation interface which returns either an `Operation.response` which contains a Dataset or `Operation.error`. The metadata field type is OperationMetadata.
 */
export class Dataset extends pulumi.CustomResource {
    /**
     * Get an existing Dataset resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Dataset {
        return new Dataset(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:healthcare/v1beta1:Dataset';

    /**
     * Returns true if the given object is an instance of Dataset.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Dataset {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Dataset.__pulumiType;
    }

    /**
     * Resource name of the dataset, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources, such as HL7 messages, where no explicit timezone is specified.
     */
    public readonly timeZone!: pulumi.Output<string>;

    /**
     * Create a Dataset resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DatasetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["datasetId"] = args ? args.datasetId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["timeZone"] = args ? args.timeZone : undefined;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["timeZone"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Dataset.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Dataset resource.
 */
export interface DatasetArgs {
    datasetId?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Resource name of the dataset, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources, such as HL7 messages, where no explicit timezone is specified.
     */
    timeZone?: pulumi.Input<string>;
}
