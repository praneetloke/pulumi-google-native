// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1.Inputs
{

    /// <summary>
    /// An excluded entity phrase that should not be matched.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The word or phrase to be excluded.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseArgs()
        {
        }
        public static new GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseArgs Empty => new GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseArgs();
    }
}
