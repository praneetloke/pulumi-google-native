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
    /// Overlaid jpeg image.
    /// </summary>
    public sealed class ImageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Target image opacity. Valid values are from `1.0` (solid, default) to `0.0` (transparent), exclusive. Set this to a value greater than `0.0`.
        /// </summary>
        [Input("alpha")]
        public Input<double>? Alpha { get; set; }

        /// <summary>
        /// Normalized image resolution, based on output video resolution. Valid values: `0.0`–`1.0`. To respect the original image aspect ratio, set either `x` or `y` to `0.0`. To use the original image resolution, set both `x` and `y` to `0.0`.
        /// </summary>
        [Input("resolution")]
        public Input<Inputs.NormalizedCoordinateArgs>? Resolution { get; set; }

        /// <summary>
        /// URI of the JPEG image in Cloud Storage. For example, `gs://bucket/inputs/image.jpeg`. JPEG is the only supported image type.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public ImageArgs()
        {
        }
        public static new ImageArgs Empty => new ImageArgs();
    }
}
