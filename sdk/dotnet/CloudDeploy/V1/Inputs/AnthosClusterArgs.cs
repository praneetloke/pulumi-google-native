// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1.Inputs
{

    /// <summary>
    /// Information specifying an Anthos Cluster.
    /// </summary>
    public sealed class AnthosClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Membership of the GKE Hub-registered cluster to which to apply the Skaffold configuration. Format is `projects/{project}/locations/{location}/memberships/{membership_name}`.
        /// </summary>
        [Input("membership")]
        public Input<string>? Membership { get; set; }

        public AnthosClusterArgs()
        {
        }
        public static new AnthosClusterArgs Empty => new AnthosClusterArgs();
    }
}
