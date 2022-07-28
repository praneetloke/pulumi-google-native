// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Inputs
{

    /// <summary>
    /// Material is a material used in the generation of the provenance
    /// </summary>
    public sealed class MaterialArgs : global::Pulumi.ResourceArgs
    {
        [Input("digest")]
        private InputMap<string>? _digest;

        /// <summary>
        /// digest is a map from a hash algorithm (e.g. sha256) to the value in the material
        /// </summary>
        public InputMap<string> Digest
        {
            get => _digest ?? (_digest = new InputMap<string>());
            set => _digest = value;
        }

        /// <summary>
        /// uri is the uri of the material
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public MaterialArgs()
        {
        }
        public static new MaterialArgs Empty => new MaterialArgs();
    }
}
