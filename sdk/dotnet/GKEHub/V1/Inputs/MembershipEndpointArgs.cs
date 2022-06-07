// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1.Inputs
{

    /// <summary>
    /// MembershipEndpoint contains information needed to contact a Kubernetes API, endpoint and any additional Kubernetes metadata.
    /// </summary>
    public sealed class MembershipEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Specific information for a GDC Edge Appliance cluster.
        /// </summary>
        [Input("applianceCluster")]
        public Input<Inputs.ApplianceClusterArgs>? ApplianceCluster { get; set; }

        /// <summary>
        /// Optional. Specific information for a Google Edge cluster.
        /// </summary>
        [Input("edgeCluster")]
        public Input<Inputs.EdgeClusterArgs>? EdgeCluster { get; set; }

        /// <summary>
        /// Optional. Specific information for a GKE-on-GCP cluster.
        /// </summary>
        [Input("gkeCluster")]
        public Input<Inputs.GkeClusterArgs>? GkeCluster { get; set; }

        /// <summary>
        /// Optional. The in-cluster Kubernetes Resources that should be applied for a correctly registered cluster, in the steady state. These resources: * Ensure that the cluster is exclusively registered to one and only one Hub Membership. * Propagate Workload Pool Information available in the Membership Authority field. * Ensure proper initial configuration of default Hub Features.
        /// </summary>
        [Input("kubernetesResource")]
        public Input<Inputs.KubernetesResourceArgs>? KubernetesResource { get; set; }

        /// <summary>
        /// Optional. Specific information for a GKE Multi-Cloud cluster.
        /// </summary>
        [Input("multiCloudCluster")]
        public Input<Inputs.MultiCloudClusterArgs>? MultiCloudCluster { get; set; }

        /// <summary>
        /// Optional. Specific information for a GKE On-Prem cluster. An onprem user-cluster who has no resourceLink is not allowed to use this field, it should have a nil "type" instead.
        /// </summary>
        [Input("onPremCluster")]
        public Input<Inputs.OnPremClusterArgs>? OnPremCluster { get; set; }

        public MembershipEndpointArgs()
        {
        }
    }
}
