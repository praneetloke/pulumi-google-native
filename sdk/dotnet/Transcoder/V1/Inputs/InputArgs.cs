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
    /// Input asset.
    /// </summary>
    public sealed class InputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique key for this input. Must be specified when using advanced mapping and edit lists.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Preprocessing configurations.
        /// </summary>
        [Input("preprocessingConfig")]
        public Input<Inputs.PreprocessingConfigArgs>? PreprocessingConfig { get; set; }

        /// <summary>
        /// URI of the media. Input files must be at least 5 seconds in duration and stored in Cloud Storage (for example, `gs://bucket/inputs/file.mp4`). If empty, the value is populated from `Job.input_uri`. See [Supported input and output formats](https://cloud.google.com/transcoder/docs/concepts/supported-input-and-output-formats).
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public InputArgs()
        {
        }
        public static new InputArgs Empty => new InputArgs();
    }
}
