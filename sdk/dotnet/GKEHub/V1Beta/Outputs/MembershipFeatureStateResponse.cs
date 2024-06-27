// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Beta.Outputs
{

    /// <summary>
    /// MembershipFeatureState contains Feature status information for a single Membership.
    /// </summary>
    [OutputType]
    public sealed class MembershipFeatureStateResponse
    {
        /// <summary>
        /// Appdevexperience specific state.
        /// </summary>
        public readonly Outputs.AppDevExperienceFeatureStateResponse Appdevexperience;
        /// <summary>
        /// ClusterUpgrade state.
        /// </summary>
        public readonly Outputs.ClusterUpgradeMembershipStateResponse Clusterupgrade;
        /// <summary>
        /// Config Management-specific state.
        /// </summary>
        public readonly Outputs.ConfigManagementMembershipStateResponse Configmanagement;
        /// <summary>
        /// Fleet observability membership state.
        /// </summary>
        public readonly Outputs.FleetObservabilityMembershipStateResponse Fleetobservability;
        /// <summary>
        /// Identity Service-specific state.
        /// </summary>
        public readonly Outputs.IdentityServiceMembershipStateResponse Identityservice;
        /// <summary>
        /// Metering-specific state.
        /// </summary>
        public readonly Outputs.MeteringMembershipStateResponse Metering;
        /// <summary>
        /// Policycontroller-specific state.
        /// </summary>
        public readonly Outputs.PolicyControllerMembershipStateResponse Policycontroller;
        /// <summary>
        /// Service Mesh-specific state.
        /// </summary>
        public readonly Outputs.ServiceMeshMembershipStateResponse Servicemesh;
        /// <summary>
        /// The high-level state of this Feature for a single membership.
        /// </summary>
        public readonly Outputs.FeatureStateResponse State;

        [OutputConstructor]
        private MembershipFeatureStateResponse(
            Outputs.AppDevExperienceFeatureStateResponse appdevexperience,

            Outputs.ClusterUpgradeMembershipStateResponse clusterupgrade,

            Outputs.ConfigManagementMembershipStateResponse configmanagement,

            Outputs.FleetObservabilityMembershipStateResponse fleetobservability,

            Outputs.IdentityServiceMembershipStateResponse identityservice,

            Outputs.MeteringMembershipStateResponse metering,

            Outputs.PolicyControllerMembershipStateResponse policycontroller,

            Outputs.ServiceMeshMembershipStateResponse servicemesh,

            Outputs.FeatureStateResponse state)
        {
            Appdevexperience = appdevexperience;
            Clusterupgrade = clusterupgrade;
            Configmanagement = configmanagement;
            Fleetobservability = fleetobservability;
            Identityservice = identityservice;
            Metering = metering;
            Policycontroller = policycontroller;
            Servicemesh = servicemesh;
            State = state;
        }
    }
}