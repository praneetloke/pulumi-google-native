// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a Serverless VPC Access connector, returns an operation.
 */
export class Connector extends pulumi.CustomResource {
    /**
     * Get an existing Connector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Connector {
        return new Connector(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:vpcaccess/v1:Connector';

    /**
     * Returns true if the given object is an instance of Connector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connector.__pulumiType;
    }

    /**
     * List of projects using the connector.
     */
    public /*out*/ readonly connectedProjects!: pulumi.Output<string[]>;
    /**
     * The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
     */
    public readonly ipCidrRange!: pulumi.Output<string>;
    /**
     * Machine type of VM Instance underlying connector. Default is e2-micro
     */
    public readonly machineType!: pulumi.Output<string>;
    /**
     * Maximum value of instances in autoscaling group underlying the connector.
     */
    public readonly maxInstances!: pulumi.Output<number>;
    /**
     * Maximum throughput of the connector in Mbps. Default is 300, max is 1000.
     */
    public readonly maxThroughput!: pulumi.Output<number>;
    /**
     * Minimum value of instances in autoscaling group underlying the connector.
     */
    public readonly minInstances!: pulumi.Output<number>;
    /**
     * Minimum throughput of the connector in Mbps. Default and min is 200.
     */
    public readonly minThroughput!: pulumi.Output<number>;
    /**
     * The resource name in the format `projects/*&#47;locations/*&#47;connectors/*`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Name of a VPC network.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * State of the VPC access connector.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The subnet in which to house the VPC Access Connector.
     */
    public readonly subnet!: pulumi.Output<outputs.vpcaccess.v1.SubnetResponse>;

    /**
     * Create a Connector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connectorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectorId'");
            }
            resourceInputs["connectorId"] = args ? args.connectorId : undefined;
            resourceInputs["ipCidrRange"] = args ? args.ipCidrRange : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["machineType"] = args ? args.machineType : undefined;
            resourceInputs["maxInstances"] = args ? args.maxInstances : undefined;
            resourceInputs["maxThroughput"] = args ? args.maxThroughput : undefined;
            resourceInputs["minInstances"] = args ? args.minInstances : undefined;
            resourceInputs["minThroughput"] = args ? args.minThroughput : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["subnet"] = args ? args.subnet : undefined;
            resourceInputs["connectedProjects"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["connectedProjects"] = undefined /*out*/;
            resourceInputs["ipCidrRange"] = undefined /*out*/;
            resourceInputs["machineType"] = undefined /*out*/;
            resourceInputs["maxInstances"] = undefined /*out*/;
            resourceInputs["maxThroughput"] = undefined /*out*/;
            resourceInputs["minInstances"] = undefined /*out*/;
            resourceInputs["minThroughput"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["network"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["subnet"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Connector.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Connector resource.
 */
export interface ConnectorArgs {
    /**
     * Required. The ID to use for this connector.
     */
    connectorId: pulumi.Input<string>;
    /**
     * The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
     */
    ipCidrRange?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Machine type of VM Instance underlying connector. Default is e2-micro
     */
    machineType?: pulumi.Input<string>;
    /**
     * Maximum value of instances in autoscaling group underlying the connector.
     */
    maxInstances?: pulumi.Input<number>;
    /**
     * Maximum throughput of the connector in Mbps. Default is 300, max is 1000.
     */
    maxThroughput?: pulumi.Input<number>;
    /**
     * Minimum value of instances in autoscaling group underlying the connector.
     */
    minInstances?: pulumi.Input<number>;
    /**
     * Minimum throughput of the connector in Mbps. Default and min is 200.
     */
    minThroughput?: pulumi.Input<number>;
    /**
     * The resource name in the format `projects/*&#47;locations/*&#47;connectors/*`.
     */
    name?: pulumi.Input<string>;
    /**
     * Name of a VPC network.
     */
    network?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The subnet in which to house the VPC Access Connector.
     */
    subnet?: pulumi.Input<inputs.vpcaccess.v1.SubnetArgs>;
}
