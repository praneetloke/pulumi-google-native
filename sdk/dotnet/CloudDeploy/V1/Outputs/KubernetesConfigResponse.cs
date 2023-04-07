// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1.Outputs
{

    /// <summary>
    /// KubernetesConfig contains the Kubernetes runtime configuration.
    /// </summary>
    [OutputType]
    public sealed class KubernetesConfigResponse
    {
        /// <summary>
        /// Kubernetes Gateway API service mesh configuration.
        /// </summary>
        public readonly Outputs.GatewayServiceMeshResponse GatewayServiceMesh;
        /// <summary>
        /// Kubernetes Service networking configuration.
        /// </summary>
        public readonly Outputs.ServiceNetworkingResponse ServiceNetworking;

        [OutputConstructor]
        private KubernetesConfigResponse(
            Outputs.GatewayServiceMeshResponse gatewayServiceMesh,

            Outputs.ServiceNetworkingResponse serviceNetworking)
        {
            GatewayServiceMesh = gatewayServiceMesh;
            ServiceNetworking = serviceNetworking;
        }
    }
}
