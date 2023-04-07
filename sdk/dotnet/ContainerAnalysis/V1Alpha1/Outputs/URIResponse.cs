// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Outputs
{

    /// <summary>
    /// An URI message.
    /// </summary>
    [OutputType]
    public sealed class URIResponse
    {
        /// <summary>
        /// A label for the URI.
        /// </summary>
        public readonly string Label;
        /// <summary>
        /// The unique resource identifier.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private URIResponse(
            string label,

            string uri)
        {
            Label = label;
            Uri = uri;
        }
    }
}
