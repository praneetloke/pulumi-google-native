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
    /// Defines a hash object for use in Materials and Products.
    /// </summary>
    public sealed class ArtifactHashesArgs : global::Pulumi.ResourceArgs
    {
        [Input("sha256")]
        public Input<string>? Sha256 { get; set; }

        public ArtifactHashesArgs()
        {
        }
        public static new ArtifactHashesArgs Empty => new ArtifactHashesArgs();
    }
}
