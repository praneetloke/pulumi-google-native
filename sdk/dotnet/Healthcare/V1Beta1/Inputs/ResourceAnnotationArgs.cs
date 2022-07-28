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
    /// Resource level annotation.
    /// </summary>
    public sealed class ResourceAnnotationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the annotation record.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        public ResourceAnnotationArgs()
        {
        }
        public static new ResourceAnnotationArgs Empty => new ResourceAnnotationArgs();
    }
}
