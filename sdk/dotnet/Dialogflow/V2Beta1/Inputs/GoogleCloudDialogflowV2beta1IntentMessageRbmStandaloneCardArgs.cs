// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1.Inputs
{

    /// <summary>
    /// Standalone Rich Business Messaging (RBM) rich card. Rich cards allow you to respond to users with more vivid content, e.g. with media and suggestions. You can group multiple rich cards into one using RbmCarouselCard but carousel cards will give you less control over the card layout.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageRbmStandaloneCardArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Card content.
        /// </summary>
        [Input("cardContent", required: true)]
        public Input<Inputs.GoogleCloudDialogflowV2beta1IntentMessageRbmCardContentArgs> CardContent { get; set; } = null!;

        /// <summary>
        /// Orientation of the card.
        /// </summary>
        [Input("cardOrientation", required: true)]
        public Input<Pulumi.GoogleNative.Dialogflow.V2Beta1.GoogleCloudDialogflowV2beta1IntentMessageRbmStandaloneCardCardOrientation> CardOrientation { get; set; } = null!;

        /// <summary>
        /// Required if orientation is horizontal. Image preview alignment for standalone cards with horizontal layout.
        /// </summary>
        [Input("thumbnailImageAlignment")]
        public Input<Pulumi.GoogleNative.Dialogflow.V2Beta1.GoogleCloudDialogflowV2beta1IntentMessageRbmStandaloneCardThumbnailImageAlignment>? ThumbnailImageAlignment { get; set; }

        public GoogleCloudDialogflowV2beta1IntentMessageRbmStandaloneCardArgs()
        {
        }
        public static new GoogleCloudDialogflowV2beta1IntentMessageRbmStandaloneCardArgs Empty => new GoogleCloudDialogflowV2beta1IntentMessageRbmStandaloneCardArgs();
    }
}
