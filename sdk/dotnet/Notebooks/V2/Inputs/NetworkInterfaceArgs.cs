// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Notebooks.V2.Inputs
{

    /// <summary>
    /// The definition of a network interface resource attached to a VM.
    /// </summary>
    public sealed class NetworkInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The name of the VPC that this VM instance is in. Format: `projects/{project_id}/global/networks/{network_id}`
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// Optional. The type of vNIC to be used on this interface. This may be gVNIC or VirtioNet.
        /// </summary>
        [Input("nicType")]
        public Input<Pulumi.GoogleNative.Notebooks.V2.NetworkInterfaceNicType>? NicType { get; set; }

        /// <summary>
        /// Optional. The name of the subnet that this VM instance is in. Format: `projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}`
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        public NetworkInterfaceArgs()
        {
        }
        public static new NetworkInterfaceArgs Empty => new NetworkInterfaceArgs();
    }
}
