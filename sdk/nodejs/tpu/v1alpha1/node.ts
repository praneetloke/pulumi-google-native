// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a node.
 * Auto-naming is currently not supported for this resource.
 */
export class Node extends pulumi.CustomResource {
    /**
     * Get an existing Node resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Node {
        return new Node(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:tpu/v1alpha1:Node';

    /**
     * Returns true if the given object is an instance of Node.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Node {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Node.__pulumiType;
    }

    /**
     * The type of hardware accelerators associated with this node.
     */
    public readonly acceleratorType!: pulumi.Output<string>;
    /**
     * The API version that created this Node.
     */
    public /*out*/ readonly apiVersion!: pulumi.Output<string>;
    /**
     * The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
     */
    public readonly cidrBlock!: pulumi.Output<string>;
    /**
     * The time when the node was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The user-supplied description of the TPU. Maximum of 512 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The health status of the TPU node.
     */
    public readonly health!: pulumi.Output<string>;
    /**
     * If this field is populated, it contains a description of why the TPU Node is unhealthy.
     */
    public /*out*/ readonly healthDescription!: pulumi.Output<string>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Immutable. The name of the TPU
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The name of a network they wish to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on which this API has been activated. If none is provided, "default" will be used.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the node reach out to the 0th entry in this map first.
     */
    public /*out*/ readonly networkEndpoints!: pulumi.Output<outputs.tpu.v1alpha1.NetworkEndpointResponse[]>;
    /**
     * The scheduling options for this node.
     */
    public readonly schedulingConfig!: pulumi.Output<outputs.tpu.v1alpha1.SchedulingConfigResponse>;
    /**
     * The service account used to run the tensor flow services within the node. To share resources, including Google Cloud Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
     */
    public /*out*/ readonly serviceAccount!: pulumi.Output<string>;
    /**
     * The current state for the TPU Node.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The Symptoms that have occurred to the TPU Node.
     */
    public /*out*/ readonly symptoms!: pulumi.Output<outputs.tpu.v1alpha1.SymptomResponse[]>;
    /**
     * The version of Tensorflow running in the Node.
     */
    public readonly tensorflowVersion!: pulumi.Output<string>;
    /**
     * Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before provisioning the node. If this field is set, cidr_block field should not be specified. If the network, that you want to peer the TPU Node to, is Shared VPC networks, the node must be created with this this field enabled.
     */
    public readonly useServiceNetworking!: pulumi.Output<boolean>;

    /**
     * Create a Node resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.acceleratorType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acceleratorType'");
            }
            if ((!args || args.tensorflowVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tensorflowVersion'");
            }
            inputs["acceleratorType"] = args ? args.acceleratorType : undefined;
            inputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["health"] = args ? args.health : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["nodeId"] = args ? args.nodeId : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["schedulingConfig"] = args ? args.schedulingConfig : undefined;
            inputs["tensorflowVersion"] = args ? args.tensorflowVersion : undefined;
            inputs["useServiceNetworking"] = args ? args.useServiceNetworking : undefined;
            inputs["apiVersion"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["healthDescription"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkEndpoints"] = undefined /*out*/;
            inputs["serviceAccount"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["symptoms"] = undefined /*out*/;
        } else {
            inputs["acceleratorType"] = undefined /*out*/;
            inputs["apiVersion"] = undefined /*out*/;
            inputs["cidrBlock"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["health"] = undefined /*out*/;
            inputs["healthDescription"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["network"] = undefined /*out*/;
            inputs["networkEndpoints"] = undefined /*out*/;
            inputs["schedulingConfig"] = undefined /*out*/;
            inputs["serviceAccount"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["symptoms"] = undefined /*out*/;
            inputs["tensorflowVersion"] = undefined /*out*/;
            inputs["useServiceNetworking"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Node.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Node resource.
 */
export interface NodeArgs {
    /**
     * The type of hardware accelerators associated with this node.
     */
    acceleratorType: pulumi.Input<string>;
    /**
     * The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * The user-supplied description of the TPU. Maximum of 512 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * The health status of the TPU node.
     */
    health?: pulumi.Input<enums.tpu.v1alpha1.NodeHealth>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * The name of a network they wish to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on which this API has been activated. If none is provided, "default" will be used.
     */
    network?: pulumi.Input<string>;
    nodeId?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The scheduling options for this node.
     */
    schedulingConfig?: pulumi.Input<inputs.tpu.v1alpha1.SchedulingConfigArgs>;
    /**
     * The version of Tensorflow running in the Node.
     */
    tensorflowVersion: pulumi.Input<string>;
    /**
     * Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before provisioning the node. If this field is set, cidr_block field should not be specified. If the network, that you want to peer the TPU Node to, is Shared VPC networks, the node must be created with this this field enabled.
     */
    useServiceNetworking?: pulumi.Input<boolean>;
}
