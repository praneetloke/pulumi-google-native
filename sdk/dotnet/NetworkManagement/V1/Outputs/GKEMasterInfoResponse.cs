// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkManagement.V1.Outputs
{

    [OutputType]
    public sealed class GKEMasterInfoResponse
    {
        /// <summary>
        /// URI of a Google Kubernetes Engine cluster network.
        /// </summary>
        public readonly string ClusterNetworkUri;
        /// <summary>
        /// URI of a Google Kubernetes Engine cluster.
        /// </summary>
        public readonly string ClusterUri;
        /// <summary>
        /// External IP address of a Google Kubernetes Engine cluster master.
        /// </summary>
        public readonly string ExternalIp;
        /// <summary>
        /// Internal IP address of a Google Kubernetes Engine cluster master.
        /// </summary>
        public readonly string InternalIp;

        [OutputConstructor]
        private GKEMasterInfoResponse(
            string clusterNetworkUri,

            string clusterUri,

            string externalIp,

            string internalIp)
        {
            ClusterNetworkUri = clusterNetworkUri;
            ClusterUri = clusterUri;
            ExternalIp = externalIp;
            InternalIp = internalIp;
        }
    }
}
