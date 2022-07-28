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
    /// Settings for exporting conversations to [Insights](https://cloud.google.com/contact-center/insights/docs).
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3beta1SecuritySettingsInsightsExportSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If enabled, we will automatically exports conversations to Insights and Insights runs its analyzers.
        /// </summary>
        [Input("enableInsightsExport")]
        public Input<bool>? EnableInsightsExport { get; set; }

        public GoogleCloudDialogflowCxV3beta1SecuritySettingsInsightsExportSettingsArgs()
        {
        }
        public static new GoogleCloudDialogflowCxV3beta1SecuritySettingsInsightsExportSettingsArgs Empty => new GoogleCloudDialogflowCxV3beta1SecuritySettingsInsightsExportSettingsArgs();
    }
}
