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
    /// KubernetesConfig contains the Kubernetes runtime configuration.
    /// </summary>
    public sealed class KubernetesConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Kubernetes Gateway API service mesh configuration.
        /// </summary>
        [Input("gatewayServiceMesh")]
        public Input<Inputs.GatewayServiceMeshArgs>? GatewayServiceMesh { get; set; }

        /// <summary>
        /// Kubernetes Service networking configuration.
        /// </summary>
        [Input("serviceNetworking")]
        public Input<Inputs.ServiceNetworkingArgs>? ServiceNetworking { get; set; }

        public KubernetesConfigArgs()
        {
        }
        public static new KubernetesConfigArgs Empty => new KubernetesConfigArgs();
    }
}
