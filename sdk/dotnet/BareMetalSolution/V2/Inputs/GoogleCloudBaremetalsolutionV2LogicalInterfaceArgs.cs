// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BareMetalSolution.V2.Inputs
{

    /// <summary>
    /// Each logical interface represents a logical abstraction of the underlying physical interface (for eg. bond, nic) of the instance. Each logical interface can effectively map to multiple network-IP pairs and still be mapped to one underlying physical interface.
    /// </summary>
    public sealed class GoogleCloudBaremetalsolutionV2LogicalInterfaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The index of the logical interface mapping to the index of the hardware bond or nic on the chosen network template.
        /// </summary>
        [Input("interfaceIndex")]
        public Input<int>? InterfaceIndex { get; set; }

        [Input("logicalNetworkInterfaces")]
        private InputList<Inputs.LogicalNetworkInterfaceArgs>? _logicalNetworkInterfaces;

        /// <summary>
        /// List of logical network interfaces within a logical interface.
        /// </summary>
        public InputList<Inputs.LogicalNetworkInterfaceArgs> LogicalNetworkInterfaces
        {
            get => _logicalNetworkInterfaces ?? (_logicalNetworkInterfaces = new InputList<Inputs.LogicalNetworkInterfaceArgs>());
            set => _logicalNetworkInterfaces = value;
        }

        /// <summary>
        /// Interface name. This is of syntax or and forms part of the network template name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GoogleCloudBaremetalsolutionV2LogicalInterfaceArgs()
        {
        }
    }
}
