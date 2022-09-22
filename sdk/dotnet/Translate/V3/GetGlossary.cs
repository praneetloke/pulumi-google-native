// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Translate.V3
{
    public static class GetGlossary
    {
        /// <summary>
        /// Gets a glossary. Returns NOT_FOUND, if the glossary doesn't exist.
        /// </summary>
        public static Task<GetGlossaryResult> InvokeAsync(GetGlossaryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGlossaryResult>("google-native:translate/v3:getGlossary", args ?? new GetGlossaryArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a glossary. Returns NOT_FOUND, if the glossary doesn't exist.
        /// </summary>
        public static Output<GetGlossaryResult> Invoke(GetGlossaryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGlossaryResult>("google-native:translate/v3:getGlossary", args ?? new GetGlossaryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGlossaryArgs : global::Pulumi.InvokeArgs
    {
        [Input("glossaryId", required: true)]
        public string GlossaryId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetGlossaryArgs()
        {
        }
        public static new GetGlossaryArgs Empty => new GetGlossaryArgs();
    }

    public sealed class GetGlossaryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("glossaryId", required: true)]
        public Input<string> GlossaryId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetGlossaryInvokeArgs()
        {
        }
        public static new GetGlossaryInvokeArgs Empty => new GetGlossaryInvokeArgs();
    }


    [OutputType]
    public sealed class GetGlossaryResult
    {
        /// <summary>
        /// Optional. The display name of the glossary.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// When the glossary creation was finished.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// The number of entries defined in the glossary.
        /// </summary>
        public readonly int EntryCount;
        /// <summary>
        /// Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
        /// </summary>
        public readonly Outputs.GlossaryInputConfigResponse InputConfig;
        /// <summary>
        /// Used with equivalent term set glossaries.
        /// </summary>
        public readonly Outputs.LanguageCodesSetResponse LanguageCodesSet;
        /// <summary>
        /// Used with unidirectional glossaries.
        /// </summary>
        public readonly Outputs.LanguageCodePairResponse LanguagePair;
        /// <summary>
        /// The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// When CreateGlossary was called.
        /// </summary>
        public readonly string SubmitTime;

        [OutputConstructor]
        private GetGlossaryResult(
            string displayName,

            string endTime,

            int entryCount,

            Outputs.GlossaryInputConfigResponse inputConfig,

            Outputs.LanguageCodesSetResponse languageCodesSet,

            Outputs.LanguageCodePairResponse languagePair,

            string name,

            string submitTime)
        {
            DisplayName = displayName;
            EndTime = endTime;
            EntryCount = entryCount;
            InputConfig = inputConfig;
            LanguageCodesSet = languageCodesSet;
            LanguagePair = languagePair;
            Name = name;
            SubmitTime = submitTime;
        }
    }
}
