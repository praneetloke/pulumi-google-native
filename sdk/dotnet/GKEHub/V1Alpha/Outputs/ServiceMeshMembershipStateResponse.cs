// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha.Outputs
{

    /// <summary>
    /// **Service Mesh**: State for a single Membership, as analyzed by the Service Mesh Hub Controller.
    /// </summary>
    [OutputType]
    public sealed class ServiceMeshMembershipStateResponse
    {
        /// <summary>
        /// Results of running Service Mesh analyzers.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceMeshAnalysisMessageResponse> AnalysisMessages;
        /// <summary>
        /// The API version (i.e. Istio CRD version) for configuring service mesh in this cluster. This version is influenced by the `default_channel` field.
        /// </summary>
        public readonly string ConfigApiVersion;
        /// <summary>
        /// Status of control plane management
        /// </summary>
        public readonly Outputs.ServiceMeshControlPlaneManagementResponse ControlPlaneManagement;
        /// <summary>
        /// Status of data plane management.
        /// </summary>
        public readonly Outputs.ServiceMeshDataPlaneManagementResponse DataPlaneManagement;

        [OutputConstructor]
        private ServiceMeshMembershipStateResponse(
            ImmutableArray<Outputs.ServiceMeshAnalysisMessageResponse> analysisMessages,

            string configApiVersion,

            Outputs.ServiceMeshControlPlaneManagementResponse controlPlaneManagement,

            Outputs.ServiceMeshDataPlaneManagementResponse dataPlaneManagement)
        {
            AnalysisMessages = analysisMessages;
            ConfigApiVersion = configApiVersion;
            ControlPlaneManagement = controlPlaneManagement;
            DataPlaneManagement = dataPlaneManagement;
        }
    }
}
