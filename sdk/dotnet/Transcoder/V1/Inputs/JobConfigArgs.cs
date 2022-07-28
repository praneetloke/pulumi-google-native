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
    /// Job configuration
    /// </summary>
    public sealed class JobConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("adBreaks")]
        private InputList<Inputs.AdBreakArgs>? _adBreaks;

        /// <summary>
        /// List of ad breaks. Specifies where to insert ad break tags in the output manifests.
        /// </summary>
        public InputList<Inputs.AdBreakArgs> AdBreaks
        {
            get => _adBreaks ?? (_adBreaks = new InputList<Inputs.AdBreakArgs>());
            set => _adBreaks = value;
        }

        [Input("editList")]
        private InputList<Inputs.EditAtomArgs>? _editList;

        /// <summary>
        /// List of `Edit atom`s. Defines the ultimate timeline of the resulting file or manifest.
        /// </summary>
        public InputList<Inputs.EditAtomArgs> EditList
        {
            get => _editList ?? (_editList = new InputList<Inputs.EditAtomArgs>());
            set => _editList = value;
        }

        [Input("elementaryStreams")]
        private InputList<Inputs.ElementaryStreamArgs>? _elementaryStreams;

        /// <summary>
        /// List of elementary streams.
        /// </summary>
        public InputList<Inputs.ElementaryStreamArgs> ElementaryStreams
        {
            get => _elementaryStreams ?? (_elementaryStreams = new InputList<Inputs.ElementaryStreamArgs>());
            set => _elementaryStreams = value;
        }

        [Input("inputs")]
        private InputList<Inputs.InputArgs>? _inputs;

        /// <summary>
        /// List of input assets stored in Cloud Storage.
        /// </summary>
        public InputList<Inputs.InputArgs> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<Inputs.InputArgs>());
            set => _inputs = value;
        }

        [Input("manifests")]
        private InputList<Inputs.ManifestArgs>? _manifests;

        /// <summary>
        /// List of output manifests.
        /// </summary>
        public InputList<Inputs.ManifestArgs> Manifests
        {
            get => _manifests ?? (_manifests = new InputList<Inputs.ManifestArgs>());
            set => _manifests = value;
        }

        [Input("muxStreams")]
        private InputList<Inputs.MuxStreamArgs>? _muxStreams;

        /// <summary>
        /// List of multiplexing settings for output streams.
        /// </summary>
        public InputList<Inputs.MuxStreamArgs> MuxStreams
        {
            get => _muxStreams ?? (_muxStreams = new InputList<Inputs.MuxStreamArgs>());
            set => _muxStreams = value;
        }

        /// <summary>
        /// Output configuration.
        /// </summary>
        [Input("output")]
        public Input<Inputs.OutputArgs>? Output { get; set; }

        [Input("overlays")]
        private InputList<Inputs.OverlayArgs>? _overlays;

        /// <summary>
        /// List of overlays on the output video, in descending Z-order.
        /// </summary>
        public InputList<Inputs.OverlayArgs> Overlays
        {
            get => _overlays ?? (_overlays = new InputList<Inputs.OverlayArgs>());
            set => _overlays = value;
        }

        /// <summary>
        /// Destination on Pub/Sub.
        /// </summary>
        [Input("pubsubDestination")]
        public Input<Inputs.PubsubDestinationArgs>? PubsubDestination { get; set; }

        [Input("spriteSheets")]
        private InputList<Inputs.SpriteSheetArgs>? _spriteSheets;

        /// <summary>
        /// List of output sprite sheets. Spritesheets require at least one VideoStream in the Jobconfig.
        /// </summary>
        public InputList<Inputs.SpriteSheetArgs> SpriteSheets
        {
            get => _spriteSheets ?? (_spriteSheets = new InputList<Inputs.SpriteSheetArgs>());
            set => _spriteSheets = value;
        }

        public JobConfigArgs()
        {
        }
        public static new JobConfigArgs Empty => new JobConfigArgs();
    }
}
