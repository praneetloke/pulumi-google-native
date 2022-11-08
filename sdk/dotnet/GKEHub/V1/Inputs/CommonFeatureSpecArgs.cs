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
    /// CommonFeatureSpec contains Hub-wide configuration information
    /// </summary>
    public sealed class CommonFeatureSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Appdevexperience specific spec.
        /// </summary>
        [Input("appdevexperience")]
        public Input<Inputs.AppDevExperienceFeatureSpecArgs>? Appdevexperience { get; set; }

        /// <summary>
        /// FleetObservability feature spec.
        /// </summary>
        [Input("fleetobservability")]
        public Input<Inputs.FleetObservabilityFeatureSpecArgs>? Fleetobservability { get; set; }

        /// <summary>
        /// Multicluster Ingress-specific spec.
        /// </summary>
        [Input("multiclusteringress")]
        public Input<Inputs.MultiClusterIngressFeatureSpecArgs>? Multiclusteringress { get; set; }

        public CommonFeatureSpecArgs()
        {
        }
        public static new CommonFeatureSpecArgs Empty => new CommonFeatureSpecArgs();
    }
}
