// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Translate.V3Beta1
{
    /// <summary>
    /// Creates a glossary and returns the long-running operation. Returns NOT_FOUND, if the project doesn't exist.
    /// </summary>
    [GoogleNativeResourceType("google-native:translate/v3beta1:Glossary")]
    public partial class Glossary : Pulumi.CustomResource
    {
        /// <summary>
        /// When the glossary creation was finished.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// The number of entries defined in the glossary.
        /// </summary>
        [Output("entryCount")]
        public Output<int> EntryCount { get; private set; } = null!;

        /// <summary>
        /// Required. Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
        /// </summary>
        [Output("inputConfig")]
        public Output<Outputs.GlossaryInputConfigResponse> InputConfig { get; private set; } = null!;

        /// <summary>
        /// Used with equivalent term set glossaries.
        /// </summary>
        [Output("languageCodesSet")]
        public Output<Outputs.LanguageCodesSetResponse> LanguageCodesSet { get; private set; } = null!;

        /// <summary>
        /// Used with unidirectional glossaries.
        /// </summary>
        [Output("languagePair")]
        public Output<Outputs.LanguageCodePairResponse> LanguagePair { get; private set; } = null!;

        /// <summary>
        /// Required. The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// When CreateGlossary was called.
        /// </summary>
        [Output("submitTime")]
        public Output<string> SubmitTime { get; private set; } = null!;


        /// <summary>
        /// Create a Glossary resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Glossary(string name, GlossaryArgs args, CustomResourceOptions? options = null)
            : base("google-native:translate/v3beta1:Glossary", name, args ?? new GlossaryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Glossary(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:translate/v3beta1:Glossary", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Glossary resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Glossary Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Glossary(name, id, options);
        }
    }

    public sealed class GlossaryArgs : Pulumi.ResourceArgs
    {
        [Input("glossaryId", required: true)]
        public Input<string> GlossaryId { get; set; } = null!;

        /// <summary>
        /// Required. Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
        /// </summary>
        [Input("inputConfig")]
        public Input<Inputs.GlossaryInputConfigArgs>? InputConfig { get; set; }

        /// <summary>
        /// Used with equivalent term set glossaries.
        /// </summary>
        [Input("languageCodesSet")]
        public Input<Inputs.LanguageCodesSetArgs>? LanguageCodesSet { get; set; }

        /// <summary>
        /// Used with unidirectional glossaries.
        /// </summary>
        [Input("languagePair")]
        public Input<Inputs.LanguageCodePairArgs>? LanguagePair { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Required. The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public GlossaryArgs()
        {
        }
    }
}
