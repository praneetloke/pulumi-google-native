// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified Interconnect. Get a list of available Interconnects by making a list() request.
 */
export function getInterconnect(args: GetInterconnectArgs, opts?: pulumi.InvokeOptions): Promise<GetInterconnectResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/v1:getInterconnect", {
        "interconnect": args.interconnect,
        "project": args.project,
    }, opts);
}

export interface GetInterconnectArgs {
    interconnect: string;
    project?: string;
}

export interface GetInterconnectResult {
    /**
     * Administrative status of the interconnect. When this is set to true, the Interconnect is functional and can carry traffic. When set to false, no packets can be carried over the interconnect and no BGP routes are exchanged over it. By default, the status is set to true.
     */
    readonly adminEnabled: boolean;
    /**
     * A list of CircuitInfo objects, that describe the individual circuits in this LAG.
     */
    readonly circuitInfos: outputs.compute.v1.InterconnectCircuitInfoResponse[];
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * Customer name, to put in the Letter of Authorization as the party authorized to request a crossconnect.
     */
    readonly customerName: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * A list of outages expected for this Interconnect.
     */
    readonly expectedOutages: outputs.compute.v1.InterconnectOutageNotificationResponse[];
    /**
     * IP address configured on the Google side of the Interconnect link. This can be used only for ping tests.
     */
    readonly googleIpAddress: string;
    /**
     * Google reference ID to be used when raising support tickets with Google or otherwise to debug backend connectivity issues.
     */
    readonly googleReferenceId: string;
    /**
     * A list of the URLs of all InterconnectAttachments configured to use this Interconnect.
     */
    readonly interconnectAttachments: string[];
    /**
     * Type of interconnect, which can take one of the following values: - PARTNER: A partner-managed interconnection shared between customers though a partner. - DEDICATED: A dedicated physical interconnection with the customer. Note that a value IT_PRIVATE has been deprecated in favor of DEDICATED.
     */
    readonly interconnectType: string;
    /**
     * Type of the resource. Always compute#interconnect for interconnects.
     */
    readonly kind: string;
    /**
     * Type of link requested, which can take one of the following values: - LINK_TYPE_ETHERNET_10G_LR: A 10G Ethernet with LR optics - LINK_TYPE_ETHERNET_100G_LR: A 100G Ethernet with LR optics. Note that this field indicates the speed of each of the links in the bundle, not the speed of the entire bundle.
     */
    readonly linkType: string;
    /**
     * URL of the InterconnectLocation object that represents where this connection is to be provisioned.
     */
    readonly location: string;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * Email address to contact the customer NOC for operations and maintenance notifications regarding this Interconnect. If specified, this will be used for notifications in addition to all other forms described, such as Cloud Monitoring logs alerting and Cloud Notifications. This field is required for users who sign up for Cloud Interconnect using workforce identity federation.
     */
    readonly nocContactEmail: string;
    /**
     * The current status of this Interconnect's functionality, which can take one of the following values: - OS_ACTIVE: A valid Interconnect, which is turned up and is ready to use. Attachments may be provisioned on this Interconnect. - OS_UNPROVISIONED: An Interconnect that has not completed turnup. No attachments may be provisioned on this Interconnect. - OS_UNDER_MAINTENANCE: An Interconnect that is undergoing internal maintenance. No attachments may be provisioned or updated on this Interconnect. 
     */
    readonly operationalStatus: string;
    /**
     * IP address configured on the customer side of the Interconnect link. The customer should configure this IP address during turnup when prompted by Google NOC. This can be used only for ping tests.
     */
    readonly peerIpAddress: string;
    /**
     * Number of links actually provisioned in this interconnect.
     */
    readonly provisionedLinkCount: number;
    /**
     * Target number of physical links in the link bundle, as requested by the customer.
     */
    readonly requestedLinkCount: number;
    /**
     * Reserved for future use.
     */
    readonly satisfiesPzs: boolean;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * The current state of Interconnect functionality, which can take one of the following values: - ACTIVE: The Interconnect is valid, turned up and ready to use. Attachments may be provisioned on this Interconnect. - UNPROVISIONED: The Interconnect has not completed turnup. No attachments may be provisioned on this Interconnect. - UNDER_MAINTENANCE: The Interconnect is undergoing internal maintenance. No attachments may be provisioned or updated on this Interconnect. 
     */
    readonly state: string;
}
/**
 * Returns the specified Interconnect. Get a list of available Interconnects by making a list() request.
 */
export function getInterconnectOutput(args: GetInterconnectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInterconnectResult> {
    return pulumi.output(args).apply((a: any) => getInterconnect(a, opts))
}

export interface GetInterconnectOutputArgs {
    interconnect: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
