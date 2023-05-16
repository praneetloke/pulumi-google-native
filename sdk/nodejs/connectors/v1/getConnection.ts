// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets details of a single Connection.
 */
export function getConnection(args: GetConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:connectors/v1:getConnection", {
        "connectionId": args.connectionId,
        "location": args.location,
        "project": args.project,
        "view": args.view,
    }, opts);
}

export interface GetConnectionArgs {
    connectionId: string;
    location: string;
    project?: string;
    view?: string;
}

export interface GetConnectionResult {
    /**
     * Optional. Configuration for establishing the connection's authentication with an external system.
     */
    readonly authConfig: outputs.connectors.v1.AuthConfigResponse;
    /**
     * Optional. Configuration for configuring the connection with an external system.
     */
    readonly configVariables: outputs.connectors.v1.ConfigVariableResponse[];
    /**
     * Connector version on which the connection is created. The format is: projects/*&#47;locations/*&#47;providers/*&#47;connectors/*&#47;versions/* Only global location is supported for ConnectorVersion resource.
     */
    readonly connectorVersion: string;
    /**
     * Created time.
     */
    readonly createTime: string;
    /**
     * Optional. Description of the resource.
     */
    readonly description: string;
    /**
     * Optional. Configuration of the Connector's destination. Only accepted for Connectors that accepts user defined destination(s).
     */
    readonly destinationConfigs: outputs.connectors.v1.DestinationConfigResponse[];
    /**
     * GCR location where the envoy image is stored. formatted like: gcr.io/{bucketName}/{imageName}
     */
    readonly envoyImageLocation: string;
    /**
     * GCR location where the runtime image is stored. formatted like: gcr.io/{bucketName}/{imageName}
     */
    readonly imageLocation: string;
    /**
     * Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
     */
    readonly labels: {[key: string]: string};
    /**
     * Optional. Configuration that indicates whether or not the Connection can be edited.
     */
    readonly lockConfig: outputs.connectors.v1.LockConfigResponse;
    /**
     * Optional. Log configuration for the connection.
     */
    readonly logConfig: outputs.connectors.v1.ConnectorsLogConfigResponse;
    /**
     * Resource name of the Connection. Format: projects/{project}/locations/{location}/connections/{connection}
     */
    readonly name: string;
    /**
     * Optional. Node configuration for the connection.
     */
    readonly nodeConfig: outputs.connectors.v1.NodeConfigResponse;
    /**
     * Optional. Service account needed for runtime plane to access GCP resources.
     */
    readonly serviceAccount: string;
    /**
     * The name of the Service Directory service name. Used for Private Harpoon to resolve the ILB address. e.g. "projects/cloud-connectors-e2e-testing/locations/us-central1/namespaces/istio-system/services/istio-ingressgateway-connectors"
     */
    readonly serviceDirectory: string;
    /**
     * Optional. Ssl config of a connection
     */
    readonly sslConfig: outputs.connectors.v1.SslConfigResponse;
    /**
     * Current status of the connection.
     */
    readonly status: outputs.connectors.v1.ConnectionStatusResponse;
    /**
     * This subscription type enum states the subscription type of the project.
     */
    readonly subscriptionType: string;
    /**
     * Optional. Suspended indicates if a user has suspended a connection or not.
     */
    readonly suspended: boolean;
    /**
     * Updated time.
     */
    readonly updateTime: string;
}
/**
 * Gets details of a single Connection.
 */
export function getConnectionOutput(args: GetConnectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConnectionResult> {
    return pulumi.output(args).apply((a: any) => getConnection(a, opts))
}

export interface GetConnectionOutputArgs {
    connectionId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    view?: pulumi.Input<string>;
}
