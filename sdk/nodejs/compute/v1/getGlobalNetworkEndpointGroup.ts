// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified network endpoint group. Gets a list of available network endpoint groups by making a list() request.
 */
export function getGlobalNetworkEndpointGroup(args: GetGlobalNetworkEndpointGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGlobalNetworkEndpointGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:compute/v1:getGlobalNetworkEndpointGroup", {
        "networkEndpointGroup": args.networkEndpointGroup,
        "project": args.project,
    }, opts);
}

export interface GetGlobalNetworkEndpointGroupArgs {
    networkEndpointGroup: string;
    project?: string;
}

export interface GetGlobalNetworkEndpointGroupResult {
    /**
     * Metadata defined as annotations on the network endpoint group.
     */
    readonly annotations: {[key: string]: string};
    /**
     * Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
     */
    readonly appEngine: outputs.compute.v1.NetworkEndpointGroupAppEngineResponse;
    /**
     * Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
     */
    readonly cloudFunction: outputs.compute.v1.NetworkEndpointGroupCloudFunctionResponse;
    /**
     * Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
     */
    readonly cloudRun: outputs.compute.v1.NetworkEndpointGroupCloudRunResponse;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * The default port used if the port number is not specified in the network endpoint.
     */
    readonly defaultPort: number;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
     */
    readonly kind: string;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
     */
    readonly network: string;
    /**
     * Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
     */
    readonly networkEndpointType: string;
    readonly pscData: outputs.compute.v1.NetworkEndpointGroupPscDataResponse;
    /**
     * The target service url used to set up private service connection to a Google API or a PSC Producer Service Attachment. An example value is: "asia-northeast3-cloudkms.googleapis.com"
     */
    readonly pscTargetService: string;
    /**
     * The URL of the region where the network endpoint group is located.
     */
    readonly region: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * [Output only] Number of network endpoints in the network endpoint group.
     */
    readonly size: number;
    /**
     * Optional URL of the subnetwork to which all network endpoints in the NEG belong.
     */
    readonly subnetwork: string;
    /**
     * The URL of the zone where the network endpoint group is located.
     */
    readonly zone: string;
}

export function getGlobalNetworkEndpointGroupOutput(args: GetGlobalNetworkEndpointGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGlobalNetworkEndpointGroupResult> {
    return pulumi.output(args).apply(a => getGlobalNetworkEndpointGroup(a, opts))
}

export interface GetGlobalNetworkEndpointGroupOutputArgs {
    networkEndpointGroup: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
