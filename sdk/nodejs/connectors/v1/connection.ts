// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new Connection in a given project and location.
 * Auto-naming is currently not supported for this resource.
 */
export class Connection extends pulumi.CustomResource {
    /**
     * Get an existing Connection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Connection {
        return new Connection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:connectors/v1:Connection';

    /**
     * Returns true if the given object is an instance of Connection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connection.__pulumiType;
    }

    /**
     * Optional. Configuration for establishing the connection's authentication with an external system.
     */
    public readonly authConfig!: pulumi.Output<outputs.connectors.v1.AuthConfigResponse>;
    /**
     * Optional. Configuration for configuring the connection with an external system.
     */
    public readonly configVariables!: pulumi.Output<outputs.connectors.v1.ConfigVariableResponse[]>;
    /**
     * Required. Identifier to assign to the Connection. Must be unique within scope of the parent resource.
     */
    public readonly connectionId!: pulumi.Output<string>;
    /**
     * Connector version on which the connection is created. The format is: projects/*&#47;locations/*&#47;providers/*&#47;connectors/*&#47;versions/* Only global location is supported for ConnectorVersion resource.
     */
    public readonly connectorVersion!: pulumi.Output<string>;
    /**
     * Created time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. Description of the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. Configuration of the Connector's destination. Only accepted for Connectors that accepts user defined destination(s).
     */
    public readonly destinationConfigs!: pulumi.Output<outputs.connectors.v1.DestinationConfigResponse[]>;
    /**
     * GCR location where the envoy image is stored. formatted like: gcr.io/{bucketName}/{imageName}
     */
    public /*out*/ readonly envoyImageLocation!: pulumi.Output<string>;
    /**
     * GCR location where the runtime image is stored. formatted like: gcr.io/{bucketName}/{imageName}
     */
    public /*out*/ readonly imageLocation!: pulumi.Output<string>;
    /**
     * Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Optional. Configuration that indicates whether or not the Connection can be edited.
     */
    public readonly lockConfig!: pulumi.Output<outputs.connectors.v1.LockConfigResponse>;
    /**
     * Optional. Log configuration for the connection.
     */
    public readonly logConfig!: pulumi.Output<outputs.connectors.v1.ConnectorsLogConfigResponse>;
    /**
     * Resource name of the Connection. Format: projects/{project}/locations/{location}/connections/{connection}
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Optional. Node configuration for the connection.
     */
    public readonly nodeConfig!: pulumi.Output<outputs.connectors.v1.NodeConfigResponse>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Optional. Service account needed for runtime plane to access GCP resources.
     */
    public readonly serviceAccount!: pulumi.Output<string>;
    /**
     * The name of the Service Directory service name. Used for Private Harpoon to resolve the ILB address. e.g. "projects/cloud-connectors-e2e-testing/locations/us-central1/namespaces/istio-system/services/istio-ingressgateway-connectors"
     */
    public /*out*/ readonly serviceDirectory!: pulumi.Output<string>;
    /**
     * Optional. Ssl config of a connection
     */
    public readonly sslConfig!: pulumi.Output<outputs.connectors.v1.SslConfigResponse>;
    /**
     * Current status of the connection.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.connectors.v1.ConnectionStatusResponse>;
    /**
     * This subscription type enum states the subscription type of the project.
     */
    public /*out*/ readonly subscriptionType!: pulumi.Output<string>;
    /**
     * Optional. Suspended indicates if a user has suspended a connection or not.
     */
    public readonly suspended!: pulumi.Output<boolean>;
    /**
     * Updated time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Connection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionId'");
            }
            if ((!args || args.connectorVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectorVersion'");
            }
            resourceInputs["authConfig"] = args ? args.authConfig : undefined;
            resourceInputs["configVariables"] = args ? args.configVariables : undefined;
            resourceInputs["connectionId"] = args ? args.connectionId : undefined;
            resourceInputs["connectorVersion"] = args ? args.connectorVersion : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destinationConfigs"] = args ? args.destinationConfigs : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["lockConfig"] = args ? args.lockConfig : undefined;
            resourceInputs["logConfig"] = args ? args.logConfig : undefined;
            resourceInputs["nodeConfig"] = args ? args.nodeConfig : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            resourceInputs["sslConfig"] = args ? args.sslConfig : undefined;
            resourceInputs["suspended"] = args ? args.suspended : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["envoyImageLocation"] = undefined /*out*/;
            resourceInputs["imageLocation"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serviceDirectory"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["subscriptionType"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["authConfig"] = undefined /*out*/;
            resourceInputs["configVariables"] = undefined /*out*/;
            resourceInputs["connectionId"] = undefined /*out*/;
            resourceInputs["connectorVersion"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["destinationConfigs"] = undefined /*out*/;
            resourceInputs["envoyImageLocation"] = undefined /*out*/;
            resourceInputs["imageLocation"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["lockConfig"] = undefined /*out*/;
            resourceInputs["logConfig"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["nodeConfig"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["serviceAccount"] = undefined /*out*/;
            resourceInputs["serviceDirectory"] = undefined /*out*/;
            resourceInputs["sslConfig"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["subscriptionType"] = undefined /*out*/;
            resourceInputs["suspended"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["connectionId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Connection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Connection resource.
 */
export interface ConnectionArgs {
    /**
     * Optional. Configuration for establishing the connection's authentication with an external system.
     */
    authConfig?: pulumi.Input<inputs.connectors.v1.AuthConfigArgs>;
    /**
     * Optional. Configuration for configuring the connection with an external system.
     */
    configVariables?: pulumi.Input<pulumi.Input<inputs.connectors.v1.ConfigVariableArgs>[]>;
    /**
     * Required. Identifier to assign to the Connection. Must be unique within scope of the parent resource.
     */
    connectionId: pulumi.Input<string>;
    /**
     * Connector version on which the connection is created. The format is: projects/*&#47;locations/*&#47;providers/*&#47;connectors/*&#47;versions/* Only global location is supported for ConnectorVersion resource.
     */
    connectorVersion: pulumi.Input<string>;
    /**
     * Optional. Description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional. Configuration of the Connector's destination. Only accepted for Connectors that accepts user defined destination(s).
     */
    destinationConfigs?: pulumi.Input<pulumi.Input<inputs.connectors.v1.DestinationConfigArgs>[]>;
    /**
     * Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Optional. Configuration that indicates whether or not the Connection can be edited.
     */
    lockConfig?: pulumi.Input<inputs.connectors.v1.LockConfigArgs>;
    /**
     * Optional. Log configuration for the connection.
     */
    logConfig?: pulumi.Input<inputs.connectors.v1.ConnectorsLogConfigArgs>;
    /**
     * Optional. Node configuration for the connection.
     */
    nodeConfig?: pulumi.Input<inputs.connectors.v1.NodeConfigArgs>;
    project?: pulumi.Input<string>;
    /**
     * Optional. Service account needed for runtime plane to access GCP resources.
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * Optional. Ssl config of a connection
     */
    sslConfig?: pulumi.Input<inputs.connectors.v1.SslConfigArgs>;
    /**
     * Optional. Suspended indicates if a user has suspended a connection or not.
     */
    suspended?: pulumi.Input<boolean>;
}
