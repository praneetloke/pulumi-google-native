// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1.Inputs
{

    /// <summary>
    /// Extra network settings. Only applicable in the App Engine flexible environment.
    /// </summary>
    public sealed class NetworkArgs : global::Pulumi.ResourceArgs
    {
        [Input("forwardedPorts")]
        private InputList<string>? _forwardedPorts;

        /// <summary>
        /// List of ports, or port pairs, to forward from the virtual machine to the application container. Only applicable in the App Engine flexible environment.
        /// </summary>
        public InputList<string> ForwardedPorts
        {
            get => _forwardedPorts ?? (_forwardedPorts = new InputList<string>());
            set => _forwardedPorts = value;
        }

        /// <summary>
        /// The IP mode for instances. Only applicable in the App Engine flexible environment.
        /// </summary>
        [Input("instanceIpMode")]
        public Input<Pulumi.GoogleNative.AppEngine.V1.NetworkInstanceIpMode>? InstanceIpMode { get; set; }

        /// <summary>
        /// Tag to apply to the instance during creation. Only applicable in the App Engine flexible environment.
        /// </summary>
        [Input("instanceTag")]
        public Input<string>? InstanceTag { get; set; }

        /// <summary>
        /// Google Compute Engine network where the virtual machines are created. Specify the short name, not the resource path.Defaults to default.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable session affinity. Only applicable in the App Engine flexible environment.
        /// </summary>
        [Input("sessionAffinity")]
        public Input<bool>? SessionAffinity { get; set; }

        /// <summary>
        /// Google Cloud Platform sub-network where the virtual machines are created. Specify the short name, not the resource path.If a subnetwork name is specified, a network name will also be required unless it is for the default network. If the network that the instance is being created in is a Legacy network, then the IP address is allocated from the IPv4Range. If the network that the instance is being created in is an auto Subnet Mode Network, then only network name should be specified (not the subnetwork_name) and the IP address is created from the IPCidrRange of the subnetwork that exists in that zone for that network. If the network that the instance is being created in is a custom Subnet Mode Network, then the subnetwork_name must be specified and the IP address is created from the IPCidrRange of the subnetwork.If specified, the subnetwork must exist in the same region as the App Engine flexible environment application.
        /// </summary>
        [Input("subnetworkName")]
        public Input<string>? SubnetworkName { get; set; }

        public NetworkArgs()
        {
        }
        public static new NetworkArgs Empty => new NetworkArgs();
    }
}
