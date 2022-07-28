// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Beta1.Inputs
{

    /// <summary>
    /// KubernetesResource contains the YAML manifests and configuration for Membership Kubernetes resources in the cluster. After CreateMembership or UpdateMembership, these resources should be re-applied in the cluster.
    /// </summary>
    public sealed class KubernetesResourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Input only. The YAML representation of the Membership CR. This field is ignored for GKE clusters where Hub can read the CR directly. Callers should provide the CR that is currently present in the cluster during CreateMembership or UpdateMembership, or leave this field empty if none exists. The CR manifest is used to validate the cluster has not been registered with another Membership.
        /// </summary>
        [Input("membershipCrManifest")]
        public Input<string>? MembershipCrManifest { get; set; }

        /// <summary>
        /// Optional. Options for Kubernetes resource generation.
        /// </summary>
        [Input("resourceOptions")]
        public Input<Inputs.ResourceOptionsArgs>? ResourceOptions { get; set; }

        public KubernetesResourceArgs()
        {
        }
        public static new KubernetesResourceArgs Empty => new KubernetesResourceArgs();
    }
}
