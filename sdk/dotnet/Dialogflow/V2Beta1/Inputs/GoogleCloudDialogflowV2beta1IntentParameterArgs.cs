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
    /// Represents intent parameters.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1IntentParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The default value to use when the `value` yields an empty result. Default values can be extracted from contexts by using the following syntax: `#context_name.parameter_name`.
        /// </summary>
        [Input("defaultValue")]
        public Input<string>? DefaultValue { get; set; }

        /// <summary>
        /// The name of the parameter.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Optional. The name of the entity type, prefixed with `@`, that describes values of the parameter. If the parameter is required, this must be provided.
        /// </summary>
        [Input("entityTypeDisplayName")]
        public Input<string>? EntityTypeDisplayName { get; set; }

        /// <summary>
        /// Optional. Indicates whether the parameter represents a list of values.
        /// </summary>
        [Input("isList")]
        public Input<bool>? IsList { get; set; }

        /// <summary>
        /// Optional. Indicates whether the parameter is required. That is, whether the intent cannot be completed without collecting the parameter value.
        /// </summary>
        [Input("mandatory")]
        public Input<bool>? Mandatory { get; set; }

        /// <summary>
        /// The unique identifier of this parameter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("prompts")]
        private InputList<string>? _prompts;

        /// <summary>
        /// Optional. The collection of prompts that the agent can present to the user in order to collect a value for the parameter.
        /// </summary>
        public InputList<string> Prompts
        {
            get => _prompts ?? (_prompts = new InputList<string>());
            set => _prompts = value;
        }

        /// <summary>
        /// Optional. The definition of the parameter value. It can be: - a constant string, - a parameter value defined as `$parameter_name`, - an original parameter value defined as `$parameter_name.original`, - a parameter value from some context defined as `#context_name.parameter_name`.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public GoogleCloudDialogflowV2beta1IntentParameterArgs()
        {
        }
        public static new GoogleCloudDialogflowV2beta1IntentParameterArgs Empty => new GoogleCloudDialogflowV2beta1IntentParameterArgs();
    }
}
