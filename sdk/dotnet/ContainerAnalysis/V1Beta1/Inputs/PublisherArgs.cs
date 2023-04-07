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
    /// Publisher contains information about the publisher of this Note.
    /// </summary>
    public sealed class PublisherArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Provides information about the authority of the issuing party to release the document, in particular, the party's constituency and responsibilities or other obligations.
        /// </summary>
        [Input("issuingAuthority")]
        public Input<string>? IssuingAuthority { get; set; }

        /// <summary>
        /// Name of the publisher. Examples: 'Google', 'Google Cloud Platform'.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The context or namespace. Contains a URL which is under control of the issuing party and can be used as a globally unique identifier for that issuing party. Example: https://csaf.io
        /// </summary>
        [Input("publisherNamespace")]
        public Input<string>? PublisherNamespace { get; set; }

        public PublisherArgs()
        {
        }
        public static new PublisherArgs Empty => new PublisherArgs();
    }
}
