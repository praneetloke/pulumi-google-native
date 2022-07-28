// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    /// <summary>
    /// Creates a Interconnect in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/v1:Interconnect")]
    public partial class Interconnect : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Administrative status of the interconnect. When this is set to true, the Interconnect is functional and can carry traffic. When set to false, no packets can be carried over the interconnect and no BGP routes are exchanged over it. By default, the status is set to true.
        /// </summary>
        [Output("adminEnabled")]
        public Output<bool> AdminEnabled { get; private set; } = null!;

        /// <summary>
        /// A list of CircuitInfo objects, that describe the individual circuits in this LAG.
        /// </summary>
        [Output("circuitInfos")]
        public Output<ImmutableArray<Outputs.InterconnectCircuitInfoResponse>> CircuitInfos { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// Customer name, to put in the Letter of Authorization as the party authorized to request a crossconnect.
        /// </summary>
        [Output("customerName")]
        public Output<string> CustomerName { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// A list of outages expected for this Interconnect.
        /// </summary>
        [Output("expectedOutages")]
        public Output<ImmutableArray<Outputs.InterconnectOutageNotificationResponse>> ExpectedOutages { get; private set; } = null!;

        /// <summary>
        /// IP address configured on the Google side of the Interconnect link. This can be used only for ping tests.
        /// </summary>
        [Output("googleIpAddress")]
        public Output<string> GoogleIpAddress { get; private set; } = null!;

        /// <summary>
        /// Google reference ID to be used when raising support tickets with Google or otherwise to debug backend connectivity issues.
        /// </summary>
        [Output("googleReferenceId")]
        public Output<string> GoogleReferenceId { get; private set; } = null!;

        /// <summary>
        /// A list of the URLs of all InterconnectAttachments configured to use this Interconnect.
        /// </summary>
        [Output("interconnectAttachments")]
        public Output<ImmutableArray<string>> InterconnectAttachments { get; private set; } = null!;

        /// <summary>
        /// Type of interconnect, which can take one of the following values: - PARTNER: A partner-managed interconnection shared between customers though a partner. - DEDICATED: A dedicated physical interconnection with the customer. Note that a value IT_PRIVATE has been deprecated in favor of DEDICATED.
        /// </summary>
        [Output("interconnectType")]
        public Output<string> InterconnectType { get; private set; } = null!;

        /// <summary>
        /// Type of the resource. Always compute#interconnect for interconnects.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Type of link requested, which can take one of the following values: - LINK_TYPE_ETHERNET_10G_LR: A 10G Ethernet with LR optics - LINK_TYPE_ETHERNET_100G_LR: A 100G Ethernet with LR optics. Note that this field indicates the speed of each of the links in the bundle, not the speed of the entire bundle.
        /// </summary>
        [Output("linkType")]
        public Output<string> LinkType { get; private set; } = null!;

        /// <summary>
        /// URL of the InterconnectLocation object that represents where this connection is to be provisioned.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Email address to contact the customer NOC for operations and maintenance notifications regarding this Interconnect. If specified, this will be used for notifications in addition to all other forms described, such as Stackdriver logs alerting and Cloud Notifications.
        /// </summary>
        [Output("nocContactEmail")]
        public Output<string> NocContactEmail { get; private set; } = null!;

        /// <summary>
        /// The current status of this Interconnect's functionality, which can take one of the following values: - OS_ACTIVE: A valid Interconnect, which is turned up and is ready to use. Attachments may be provisioned on this Interconnect. - OS_UNPROVISIONED: An Interconnect that has not completed turnup. No attachments may be provisioned on this Interconnect. - OS_UNDER_MAINTENANCE: An Interconnect that is undergoing internal maintenance. No attachments may be provisioned or updated on this Interconnect. 
        /// </summary>
        [Output("operationalStatus")]
        public Output<string> OperationalStatus { get; private set; } = null!;

        /// <summary>
        /// IP address configured on the customer side of the Interconnect link. The customer should configure this IP address during turnup when prompted by Google NOC. This can be used only for ping tests.
        /// </summary>
        [Output("peerIpAddress")]
        public Output<string> PeerIpAddress { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Number of links actually provisioned in this interconnect.
        /// </summary>
        [Output("provisionedLinkCount")]
        public Output<int> ProvisionedLinkCount { get; private set; } = null!;

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// Target number of physical links in the link bundle, as requested by the customer.
        /// </summary>
        [Output("requestedLinkCount")]
        public Output<int> RequestedLinkCount { get; private set; } = null!;

        /// <summary>
        /// Set to true if the resource satisfies the zone separation organization policy constraints and false otherwise. Defaults to false if the field is not present.
        /// </summary>
        [Output("satisfiesPzs")]
        public Output<bool> SatisfiesPzs { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The current state of Interconnect functionality, which can take one of the following values: - ACTIVE: The Interconnect is valid, turned up and ready to use. Attachments may be provisioned on this Interconnect. - UNPROVISIONED: The Interconnect has not completed turnup. No attachments may be provisioned on this Interconnect. - UNDER_MAINTENANCE: The Interconnect is undergoing internal maintenance. No attachments may be provisioned or updated on this Interconnect. 
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a Interconnect resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Interconnect(string name, InterconnectArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:Interconnect", name, args ?? new InterconnectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Interconnect(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:Interconnect", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Interconnect resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Interconnect Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Interconnect(name, id, options);
        }
    }

    public sealed class InterconnectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative status of the interconnect. When this is set to true, the Interconnect is functional and can carry traffic. When set to false, no packets can be carried over the interconnect and no BGP routes are exchanged over it. By default, the status is set to true.
        /// </summary>
        [Input("adminEnabled")]
        public Input<bool>? AdminEnabled { get; set; }

        /// <summary>
        /// Customer name, to put in the Letter of Authorization as the party authorized to request a crossconnect.
        /// </summary>
        [Input("customerName")]
        public Input<string>? CustomerName { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Type of interconnect, which can take one of the following values: - PARTNER: A partner-managed interconnection shared between customers though a partner. - DEDICATED: A dedicated physical interconnection with the customer. Note that a value IT_PRIVATE has been deprecated in favor of DEDICATED.
        /// </summary>
        [Input("interconnectType")]
        public Input<Pulumi.GoogleNative.Compute.V1.InterconnectInterconnectType>? InterconnectType { get; set; }

        /// <summary>
        /// Type of link requested, which can take one of the following values: - LINK_TYPE_ETHERNET_10G_LR: A 10G Ethernet with LR optics - LINK_TYPE_ETHERNET_100G_LR: A 100G Ethernet with LR optics. Note that this field indicates the speed of each of the links in the bundle, not the speed of the entire bundle.
        /// </summary>
        [Input("linkType")]
        public Input<Pulumi.GoogleNative.Compute.V1.InterconnectLinkType>? LinkType { get; set; }

        /// <summary>
        /// URL of the InterconnectLocation object that represents where this connection is to be provisioned.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Email address to contact the customer NOC for operations and maintenance notifications regarding this Interconnect. If specified, this will be used for notifications in addition to all other forms described, such as Stackdriver logs alerting and Cloud Notifications.
        /// </summary>
        [Input("nocContactEmail")]
        public Input<string>? NocContactEmail { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Target number of physical links in the link bundle, as requested by the customer.
        /// </summary>
        [Input("requestedLinkCount")]
        public Input<int>? RequestedLinkCount { get; set; }

        public InterconnectArgs()
        {
        }
        public static new InterconnectArgs Empty => new InterconnectArgs();
    }
}
