// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// Represents a Nat resource. It enables the VMs within the specified subnetworks to access Internet without external IP addresses. It specifies a list of subnetworks (and the ranges within) that want to use NAT. Customers can also provide the external IPs that would be used for NAT. GCP would auto-allocate ephemeral IPs if no external IPs are provided.
    /// </summary>
    public sealed class RouterNatArgs : global::Pulumi.ResourceArgs
    {
        [Input("drainNatIps")]
        private InputList<string>? _drainNatIps;

        /// <summary>
        /// A list of URLs of the IP resources to be drained. These IPs must be valid static external IPs that have been assigned to the NAT. These IPs should be used for updating/patching a NAT only.
        /// </summary>
        public InputList<string> DrainNatIps
        {
            get => _drainNatIps ?? (_drainNatIps = new InputList<string>());
            set => _drainNatIps = value;
        }

        /// <summary>
        /// Enable Dynamic Port Allocation. If not specified, it is disabled by default. If set to true, - Dynamic Port Allocation will be enabled on this NAT config. - enableEndpointIndependentMapping cannot be set to true. - If minPorts is set, minPortsPerVm must be set to a power of two greater than or equal to 32. If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config. 
        /// </summary>
        [Input("enableDynamicPortAllocation")]
        public Input<bool>? EnableDynamicPortAllocation { get; set; }

        [Input("enableEndpointIndependentMapping")]
        public Input<bool>? EnableEndpointIndependentMapping { get; set; }

        [Input("endpointTypes")]
        private InputList<Pulumi.GoogleNative.Compute.V1.RouterNatEndpointTypesItem>? _endpointTypes;

        /// <summary>
        /// List of NAT-ted endpoint types supported by the Nat Gateway. If the list is empty, then it will be equivalent to include ENDPOINT_TYPE_VM
        /// </summary>
        public InputList<Pulumi.GoogleNative.Compute.V1.RouterNatEndpointTypesItem> EndpointTypes
        {
            get => _endpointTypes ?? (_endpointTypes = new InputList<Pulumi.GoogleNative.Compute.V1.RouterNatEndpointTypesItem>());
            set => _endpointTypes = value;
        }

        /// <summary>
        /// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
        /// </summary>
        [Input("icmpIdleTimeoutSec")]
        public Input<int>? IcmpIdleTimeoutSec { get; set; }

        /// <summary>
        /// Configure logging on this NAT.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.RouterNatLogConfigArgs>? LogConfig { get; set; }

        /// <summary>
        /// Maximum number of ports allocated to a VM from this NAT config when Dynamic Port Allocation is enabled. If Dynamic Port Allocation is not enabled, this field has no effect. If Dynamic Port Allocation is enabled, and this field is set, it must be set to a power of two greater than minPortsPerVm, or 64 if minPortsPerVm is not set. If Dynamic Port Allocation is enabled and this field is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
        /// </summary>
        [Input("maxPortsPerVm")]
        public Input<int>? MaxPortsPerVm { get; set; }

        /// <summary>
        /// Minimum number of ports allocated to a VM from this NAT config. If not set, a default number of ports is allocated to a VM. This is rounded up to the nearest power of 2. For example, if the value of this field is 50, at least 64 ports are allocated to a VM.
        /// </summary>
        [Input("minPortsPerVm")]
        public Input<int>? MinPortsPerVm { get; set; }

        /// <summary>
        /// Unique name of this Nat service. The name must be 1-63 characters long and comply with RFC1035.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specify the NatIpAllocateOption, which can take one of the following values: - MANUAL_ONLY: Uses only Nat IP addresses provided by customers. When there are not enough specified Nat IPs, the Nat service fails for new VMs. - AUTO_ONLY: Nat IPs are allocated by Google Cloud Platform; customers can't specify any Nat IPs. When choosing AUTO_ONLY, then nat_ip should be empty. 
        /// </summary>
        [Input("natIpAllocateOption")]
        public Input<Pulumi.GoogleNative.Compute.V1.RouterNatNatIpAllocateOption>? NatIpAllocateOption { get; set; }

        [Input("natIps")]
        private InputList<string>? _natIps;

        /// <summary>
        /// A list of URLs of the IP resources used for this Nat service. These IP addresses must be valid static external IP addresses assigned to the project.
        /// </summary>
        public InputList<string> NatIps
        {
            get => _natIps ?? (_natIps = new InputList<string>());
            set => _natIps = value;
        }

        [Input("rules")]
        private InputList<Inputs.RouterNatRuleArgs>? _rules;

        /// <summary>
        /// A list of rules associated with this NAT.
        /// </summary>
        public InputList<Inputs.RouterNatRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.RouterNatRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Specify the Nat option, which can take one of the following values: - ALL_SUBNETWORKS_ALL_IP_RANGES: All of the IP ranges in every Subnetwork are allowed to Nat. - ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES: All of the primary IP ranges in every Subnetwork are allowed to Nat. - LIST_OF_SUBNETWORKS: A list of Subnetworks are allowed to Nat (specified in the field subnetwork below) The default is SUBNETWORK_IP_RANGE_TO_NAT_OPTION_UNSPECIFIED. Note that if this field contains ALL_SUBNETWORKS_ALL_IP_RANGES then there should not be any other Router.Nat section in any Router for this network in this region.
        /// </summary>
        [Input("sourceSubnetworkIpRangesToNat")]
        public Input<Pulumi.GoogleNative.Compute.V1.RouterNatSourceSubnetworkIpRangesToNat>? SourceSubnetworkIpRangesToNat { get; set; }

        [Input("subnetworks")]
        private InputList<Inputs.RouterNatSubnetworkToNatArgs>? _subnetworks;

        /// <summary>
        /// A list of Subnetwork resources whose traffic should be translated by NAT Gateway. It is used only when LIST_OF_SUBNETWORKS is selected for the SubnetworkIpRangeToNatOption above.
        /// </summary>
        public InputList<Inputs.RouterNatSubnetworkToNatArgs> Subnetworks
        {
            get => _subnetworks ?? (_subnetworks = new InputList<Inputs.RouterNatSubnetworkToNatArgs>());
            set => _subnetworks = value;
        }

        /// <summary>
        /// Timeout (in seconds) for TCP established connections. Defaults to 1200s if not set.
        /// </summary>
        [Input("tcpEstablishedIdleTimeoutSec")]
        public Input<int>? TcpEstablishedIdleTimeoutSec { get; set; }

        /// <summary>
        /// Timeout (in seconds) for TCP connections that are in TIME_WAIT state. Defaults to 120s if not set.
        /// </summary>
        [Input("tcpTimeWaitTimeoutSec")]
        public Input<int>? TcpTimeWaitTimeoutSec { get; set; }

        /// <summary>
        /// Timeout (in seconds) for TCP transitory connections. Defaults to 30s if not set.
        /// </summary>
        [Input("tcpTransitoryIdleTimeoutSec")]
        public Input<int>? TcpTransitoryIdleTimeoutSec { get; set; }

        /// <summary>
        /// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
        /// </summary>
        [Input("udpIdleTimeoutSec")]
        public Input<int>? UdpIdleTimeoutSec { get; set; }

        public RouterNatArgs()
        {
        }
        public static new RouterNatArgs Empty => new RouterNatArgs();
    }
}
