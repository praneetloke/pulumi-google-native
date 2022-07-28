// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha.Inputs
{

    /// <summary>
    /// ResourceOptions represent options for Kubernetes resource generation.
    /// </summary>
    public sealed class ResourceOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The Connect agent version to use for connect_resources. Defaults to the latest GKE Connect version. The version must be a currently supported version, obsolete versions will be rejected.
        /// </summary>
        [Input("connectVersion")]
        public Input<string>? ConnectVersion { get; set; }

        /// <summary>
        /// Optional. Major version of the Kubernetes cluster. This is only used to determine which version to use for the CustomResourceDefinition resources, `apiextensions/v1beta1` or`apiextensions/v1`.
        /// </summary>
        [Input("k8sVersion")]
        public Input<string>? K8sVersion { get; set; }

        /// <summary>
        /// Optional. Use `apiextensions/v1beta1` instead of `apiextensions/v1` for CustomResourceDefinition resources. This option should be set for clusters with Kubernetes apiserver versions &lt;1.16.
        /// </summary>
        [Input("v1beta1Crd")]
        public Input<bool>? V1beta1Crd { get; set; }

        public ResourceOptionsArgs()
        {
        }
        public static new ResourceOptionsArgs Empty => new ResourceOptionsArgs();
    }
}
