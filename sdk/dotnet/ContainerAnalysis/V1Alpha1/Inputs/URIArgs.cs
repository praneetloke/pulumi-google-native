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
    /// An URI message.
    /// </summary>
    public sealed class URIArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A label for the URI.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// The unique resource identifier.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public URIArgs()
        {
        }
        public static new URIArgs Empty => new URIArgs();
    }
}