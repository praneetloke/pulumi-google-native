// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3
{
    public static class GetSecuritySetting
    {
        /// <summary>
        /// Retrieves the specified SecuritySettings. The returned settings may be stale by up to 1 minute.
        /// </summary>
        public static Task<GetSecuritySettingResult> InvokeAsync(GetSecuritySettingArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecuritySettingResult>("google-native:dialogflow/v3:getSecuritySetting", args ?? new GetSecuritySettingArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the specified SecuritySettings. The returned settings may be stale by up to 1 minute.
        /// </summary>
        public static Output<GetSecuritySettingResult> Invoke(GetSecuritySettingInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSecuritySettingResult>("google-native:dialogflow/v3:getSecuritySetting", args ?? new GetSecuritySettingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecuritySettingArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("securitySettingId", required: true)]
        public string SecuritySettingId { get; set; } = null!;

        public GetSecuritySettingArgs()
        {
        }
        public static new GetSecuritySettingArgs Empty => new GetSecuritySettingArgs();
    }

    public sealed class GetSecuritySettingInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("securitySettingId", required: true)]
        public Input<string> SecuritySettingId { get; set; } = null!;

        public GetSecuritySettingInvokeArgs()
        {
        }
        public static new GetSecuritySettingInvokeArgs Empty => new GetSecuritySettingInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecuritySettingResult
    {
        /// <summary>
        /// Controls audio export settings for post-conversation analytics when ingesting audio to conversations via Participants.AnalyzeContent or Participants.StreamingAnalyzeContent. If retention_strategy is set to REMOVE_AFTER_CONVERSATION or audio_export_settings.gcs_bucket is empty, audio export is disabled. If audio export is enabled, audio is recorded and saved to audio_export_settings.gcs_bucket, subject to retention policy of audio_export_settings.gcs_bucket. This setting won't effect audio input for implicit sessions via Sessions.DetectIntent or Sessions.StreamingDetectIntent.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowCxV3SecuritySettingsAudioExportSettingsResponse AudioExportSettings;
        /// <summary>
        /// [DLP](https://cloud.google.com/dlp/docs) deidentify template name. Use this template to define de-identification configuration for the content. The `DLP De-identify Templates Reader` role is needed on the Dialogflow service identity service account (has the form `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`) for your agent's project. If empty, Dialogflow replaces sensitive info with `[redacted]` text. The template name will have one of the following formats: `projects//locations//deidentifyTemplates/` OR `organizations//locations//deidentifyTemplates/` Note: `deidentify_template` must be located in the same region as the `SecuritySettings`.
        /// </summary>
        public readonly string DeidentifyTemplate;
        /// <summary>
        /// The human-readable name of the security settings, unique within the location.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Controls conversation exporting settings to Insights after conversation is completed. If retention_strategy is set to REMOVE_AFTER_CONVERSATION, Insights export is disabled no matter what you configure here.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowCxV3SecuritySettingsInsightsExportSettingsResponse InsightsExportSettings;
        /// <summary>
        /// [DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this template to define inspect base settings. The `DLP Inspect Templates Reader` role is needed on the Dialogflow service identity service account (has the form `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`) for your agent's project. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects//locations//inspectTemplates/` OR `organizations//locations//inspectTemplates/` Note: `inspect_template` must be located in the same region as the `SecuritySettings`.
        /// </summary>
        public readonly string InspectTemplate;
        /// <summary>
        /// Resource name of the settings. Required for the SecuritySettingsService.UpdateSecuritySettings method. SecuritySettingsService.CreateSecuritySettings populates the name automatically. Format: `projects//locations//securitySettings/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of types of data to remove when retention settings triggers purge.
        /// </summary>
        public readonly ImmutableArray<string> PurgeDataTypes;
        /// <summary>
        /// Defines the data for which Dialogflow applies redaction. Dialogflow does not redact data that it does not have access to – for example, Cloud logging.
        /// </summary>
        public readonly string RedactionScope;
        /// <summary>
        /// Strategy that defines how we do redaction.
        /// </summary>
        public readonly string RedactionStrategy;
        /// <summary>
        /// Retains data in interaction logging for the specified number of days. This does not apply to Cloud logging, which is owned by the user - not Dialogflow. User must set a value lower than Dialogflow's default 365d TTL. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use Dialogflow's default TTL. Note: Interaction logging is a limited access feature. Talk to your Google representative to check availability for you.
        /// </summary>
        public readonly int RetentionWindowDays;

        [OutputConstructor]
        private GetSecuritySettingResult(
            Outputs.GoogleCloudDialogflowCxV3SecuritySettingsAudioExportSettingsResponse audioExportSettings,

            string deidentifyTemplate,

            string displayName,

            Outputs.GoogleCloudDialogflowCxV3SecuritySettingsInsightsExportSettingsResponse insightsExportSettings,

            string inspectTemplate,

            string name,

            ImmutableArray<string> purgeDataTypes,

            string redactionScope,

            string redactionStrategy,

            int retentionWindowDays)
        {
            AudioExportSettings = audioExportSettings;
            DeidentifyTemplate = deidentifyTemplate;
            DisplayName = displayName;
            InsightsExportSettings = insightsExportSettings;
            InspectTemplate = inspectTemplate;
            Name = name;
            PurgeDataTypes = purgeDataTypes;
            RedactionScope = redactionScope;
            RedactionStrategy = redactionStrategy;
            RetentionWindowDays = retentionWindowDays;
        }
    }
}
