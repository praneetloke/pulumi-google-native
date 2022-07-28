// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Inputs
{

    /// <summary>
    /// An artifact that can be deployed in some runtime.
    /// </summary>
    public sealed class DeployableArgs : global::Pulumi.ResourceArgs
    {
        [Input("resourceUri", required: true)]
        private InputList<string>? _resourceUri;

        /// <summary>
        /// Resource URI for the artifact being deployed.
        /// </summary>
        public InputList<string> ResourceUri
        {
            get => _resourceUri ?? (_resourceUri = new InputList<string>());
            set => _resourceUri = value;
        }

        public DeployableArgs()
        {
        }
        public static new DeployableArgs Empty => new DeployableArgs();
    }
}
