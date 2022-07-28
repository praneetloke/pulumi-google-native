// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1.Inputs
{

    /// <summary>
    /// A bounding polygon for the detected image annotation.
    /// </summary>
    public sealed class BoundingPolyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of this polygon.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        [Input("vertices")]
        private InputList<Inputs.VertexArgs>? _vertices;

        /// <summary>
        /// List of the vertices of this polygon.
        /// </summary>
        public InputList<Inputs.VertexArgs> Vertices
        {
            get => _vertices ?? (_vertices = new InputList<Inputs.VertexArgs>());
            set => _vertices = value;
        }

        public BoundingPolyArgs()
        {
        }
        public static new BoundingPolyArgs Empty => new BoundingPolyArgs();
    }
}
