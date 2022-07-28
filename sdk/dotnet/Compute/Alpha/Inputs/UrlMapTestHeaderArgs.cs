// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// HTTP headers used in UrlMapTests.
    /// </summary>
    public sealed class UrlMapTestHeaderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Header name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Header value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public UrlMapTestHeaderArgs()
        {
        }
        public static new UrlMapTestHeaderArgs Empty => new UrlMapTestHeaderArgs();
    }
}
