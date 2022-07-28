// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1.Inputs
{

    /// <summary>
    /// Location of output file(s) in a Cloud Storage bucket.
    /// </summary>
    public sealed class OutputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// URI for the output file(s). For example, `gs://my-bucket/outputs/`. If empty, the value is populated from `Job.output_uri`. See [Supported input and output formats](https://cloud.google.com/transcoder/docs/concepts/supported-input-and-output-formats).
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public OutputArgs()
        {
        }
        public static new OutputArgs Empty => new OutputArgs();
    }
}
