// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Inputs
{

    /// <summary>
    /// Hierarchical advanced settings for agent/flow/page/fulfillment/parameter. Settings exposed at lower level overrides the settings exposed at higher level. Overriding occurs at the sub-setting level. For example, the playback_interruption_settings at fulfillment level only overrides the playback_interruption_settings at the agent level, leaving other settings at the agent level unchanged. DTMF settings does not override each other. DTMF settings set at different levels define DTMF detections running in parallel. Hierarchy: Agent-&gt;Flow-&gt;Page-&gt;Fulfillment/Parameter.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3AdvancedSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Settings for logging. Settings for Dialogflow History, Contact Center messages, StackDriver logs, and speech logging. Exposed at the following levels: - Agent level.
        /// </summary>
        [Input("loggingSettings")]
        public Input<Inputs.GoogleCloudDialogflowCxV3AdvancedSettingsLoggingSettingsArgs>? LoggingSettings { get; set; }

        public GoogleCloudDialogflowCxV3AdvancedSettingsArgs()
        {
        }
        public static new GoogleCloudDialogflowCxV3AdvancedSettingsArgs Empty => new GoogleCloudDialogflowCxV3AdvancedSettingsArgs();
    }
}
