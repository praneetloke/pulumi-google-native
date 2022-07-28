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
    /// Metadata for any related URL information
    /// </summary>
    public sealed class RelatedUrlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Label to describe usage of the URL
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Specific URL to associate with the note
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public RelatedUrlArgs()
        {
        }
        public static new RelatedUrlArgs Empty => new RelatedUrlArgs();
    }
}
