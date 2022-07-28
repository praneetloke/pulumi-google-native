// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    public static class GetInterconnect
    {
        /// <summary>
        /// Returns the specified interconnect. Get a list of available interconnects by making a list() request.
        /// </summary>
        public static Task<GetInterconnectResult> InvokeAsync(GetInterconnectArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInterconnectResult>("google-native:compute/v1:getInterconnect", args ?? new GetInterconnectArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified interconnect. Get a list of available interconnects by making a list() request.
        /// </summary>
        public static Output<GetInterconnectResult> Invoke(GetInterconnectInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInterconnectResult>("google-native:compute/v1:getInterconnect", args ?? new GetInterconnectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInterconnectArgs : global::Pulumi.InvokeArgs
    {
        [Input("interconnect", required: true)]
        public string Interconnect { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetInterconnectArgs()
        {
        }
        public static new GetInterconnectArgs Empty => new GetInterconnectArgs();
    }

    public sealed class GetInterconnectInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("interconnect", required: true)]
        public Input<string> Interconnect { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetInterconnectInvokeArgs()
        {
        }
        public static new GetInterconnectInvokeArgs Empty => new GetInterconnectInvokeArgs();
    }


    [OutputType]
    public sealed class GetInterconnectResult
    {
        /// <summary>
        /// Administrative status of the interconnect. When this is set to true, the Interconnect is functional and can carry traffic. When set to false, no packets can be carried over the interconnect and no BGP routes are exchanged over it. By default, the status is set to true.
        /// </summary>
        public readonly bool AdminEnabled;
        /// <summary>
        /// A list of CircuitInfo objects, that describe the individual circuits in this LAG.
        /// </summary>
        public readonly ImmutableArray<Outputs.InterconnectCircuitInfoResponse> CircuitInfos;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// Customer name, to put in the Letter of Authorization as the party authorized to request a crossconnect.
        /// </summary>
        public readonly string CustomerName;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// A list of outages expected for this Interconnect.
        /// </summary>
        public readonly ImmutableArray<Outputs.InterconnectOutageNotificationResponse> ExpectedOutages;
        /// <summary>
        /// IP address configured on the Google side of the Interconnect link. This can be used only for ping tests.
        /// </summary>
        public readonly string GoogleIpAddress;
        /// <summary>
        /// Google reference ID to be used when raising support tickets with Google or otherwise to debug backend connectivity issues.
        /// </summary>
        public readonly string GoogleReferenceId;
        /// <summary>
        /// A list of the URLs of all InterconnectAttachments configured to use this Interconnect.
        /// </summary>
        public readonly ImmutableArray<string> InterconnectAttachments;
        /// <summary>
        /// Type of interconnect, which can take one of the following values: - PARTNER: A partner-managed interconnection shared between customers though a partner. - DEDICATED: A dedicated physical interconnection with the customer. Note that a value IT_PRIVATE has been deprecated in favor of DEDICATED.
        /// </summary>
        public readonly string InterconnectType;
        /// <summary>
        /// Type of the resource. Always compute#interconnect for interconnects.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Type of link requested, which can take one of the following values: - LINK_TYPE_ETHERNET_10G_LR: A 10G Ethernet with LR optics - LINK_TYPE_ETHERNET_100G_LR: A 100G Ethernet with LR optics. Note that this field indicates the speed of each of the links in the bundle, not the speed of the entire bundle.
        /// </summary>
        public readonly string LinkType;
        /// <summary>
        /// URL of the InterconnectLocation object that represents where this connection is to be provisioned.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Email address to contact the customer NOC for operations and maintenance notifications regarding this Interconnect. If specified, this will be used for notifications in addition to all other forms described, such as Stackdriver logs alerting and Cloud Notifications.
        /// </summary>
        public readonly string NocContactEmail;
        /// <summary>
        /// The current status of this Interconnect's functionality, which can take one of the following values: - OS_ACTIVE: A valid Interconnect, which is turned up and is ready to use. Attachments may be provisioned on this Interconnect. - OS_UNPROVISIONED: An Interconnect that has not completed turnup. No attachments may be provisioned on this Interconnect. - OS_UNDER_MAINTENANCE: An Interconnect that is undergoing internal maintenance. No attachments may be provisioned or updated on this Interconnect. 
        /// </summary>
        public readonly string OperationalStatus;
        /// <summary>
        /// IP address configured on the customer side of the Interconnect link. The customer should configure this IP address during turnup when prompted by Google NOC. This can be used only for ping tests.
        /// </summary>
        public readonly string PeerIpAddress;
        /// <summary>
        /// Number of links actually provisioned in this interconnect.
        /// </summary>
        public readonly int ProvisionedLinkCount;
        /// <summary>
        /// Target number of physical links in the link bundle, as requested by the customer.
        /// </summary>
        public readonly int RequestedLinkCount;
        /// <summary>
        /// Set to true if the resource satisfies the zone separation organization policy constraints and false otherwise. Defaults to false if the field is not present.
        /// </summary>
        public readonly bool SatisfiesPzs;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The current state of Interconnect functionality, which can take one of the following values: - ACTIVE: The Interconnect is valid, turned up and ready to use. Attachments may be provisioned on this Interconnect. - UNPROVISIONED: The Interconnect has not completed turnup. No attachments may be provisioned on this Interconnect. - UNDER_MAINTENANCE: The Interconnect is undergoing internal maintenance. No attachments may be provisioned or updated on this Interconnect. 
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetInterconnectResult(
            bool adminEnabled,

            ImmutableArray<Outputs.InterconnectCircuitInfoResponse> circuitInfos,

            string creationTimestamp,

            string customerName,

            string description,

            ImmutableArray<Outputs.InterconnectOutageNotificationResponse> expectedOutages,

            string googleIpAddress,

            string googleReferenceId,

            ImmutableArray<string> interconnectAttachments,

            string interconnectType,

            string kind,

            string linkType,

            string location,

            string name,

            string nocContactEmail,

            string operationalStatus,

            string peerIpAddress,

            int provisionedLinkCount,

            int requestedLinkCount,

            bool satisfiesPzs,

            string selfLink,

            string state)
        {
            AdminEnabled = adminEnabled;
            CircuitInfos = circuitInfos;
            CreationTimestamp = creationTimestamp;
            CustomerName = customerName;
            Description = description;
            ExpectedOutages = expectedOutages;
            GoogleIpAddress = googleIpAddress;
            GoogleReferenceId = googleReferenceId;
            InterconnectAttachments = interconnectAttachments;
            InterconnectType = interconnectType;
            Kind = kind;
            LinkType = linkType;
            Location = location;
            Name = name;
            NocContactEmail = nocContactEmail;
            OperationalStatus = operationalStatus;
            PeerIpAddress = peerIpAddress;
            ProvisionedLinkCount = provisionedLinkCount;
            RequestedLinkCount = requestedLinkCount;
            SatisfiesPzs = satisfiesPzs;
            SelfLink = selfLink;
            State = state;
        }
    }
}
