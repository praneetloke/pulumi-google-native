// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkConnectivity.V1.Outputs
{

    /// <summary>
    /// Filter matches L4 traffic.
    /// </summary>
    [OutputType]
    public sealed class FilterResponse
    {
        /// <summary>
        /// Optional. The destination IP range of outgoing packets that this policy-based route applies to. Default is "0.0.0.0/0" if protocol version is IPv4.
        /// </summary>
        public readonly string DestRange;
        /// <summary>
        /// Optional. The IP protocol that this policy-based route applies to. Valid values are 'TCP', 'UDP', and 'ALL'. Default is 'ALL'.
        /// </summary>
        public readonly string IpProtocol;
        /// <summary>
        /// Internet protocol versions this policy-based route applies to. For this version, only IPV4 is supported.
        /// </summary>
        public readonly string ProtocolVersion;
        /// <summary>
        /// Optional. The source IP range of outgoing packets that this policy-based route applies to. Default is "0.0.0.0/0" if protocol version is IPv4.
        /// </summary>
        public readonly string SrcRange;

        [OutputConstructor]
        private FilterResponse(
            string destRange,

            string ipProtocol,

            string protocolVersion,

            string srcRange)
        {
            DestRange = destRange;
            IpProtocol = ipProtocol;
            ProtocolVersion = protocolVersion;
            SrcRange = srcRange;
        }
    }
}