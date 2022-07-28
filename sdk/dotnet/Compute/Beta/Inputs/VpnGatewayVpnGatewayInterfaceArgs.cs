// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// A VPN gateway interface.
    /// </summary>
    public sealed class VpnGatewayVpnGatewayInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// URL of the VLAN attachment (interconnectAttachment) resource for this VPN gateway interface. When the value of this field is present, the VPN gateway is used for IPsec-encrypted Cloud Interconnect; all egress or ingress traffic for this VPN gateway interface goes through the specified VLAN attachment resource. Not currently available publicly. 
        /// </summary>
        [Input("interconnectAttachment")]
        public Input<string>? InterconnectAttachment { get; set; }

        public VpnGatewayVpnGatewayInterfaceArgs()
        {
        }
        public static new VpnGatewayVpnGatewayInterfaceArgs Empty => new VpnGatewayVpnGatewayInterfaceArgs();
    }
}
