// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1.Outputs
{

    /// <summary>
    /// **Service Mesh**: State for a single Membership, as analyzed by the Service Mesh Hub Controller.
    /// </summary>
    [OutputType]
    public sealed class ServiceMeshMembershipStateResponse
    {
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
            Outputs.ServiceMeshControlPlaneManagementResponse controlPlaneManagement,

            Outputs.ServiceMeshDataPlaneManagementResponse dataPlaneManagement)
        {
            ControlPlaneManagement = controlPlaneManagement;
            DataPlaneManagement = dataPlaneManagement;
        }
    }
}
