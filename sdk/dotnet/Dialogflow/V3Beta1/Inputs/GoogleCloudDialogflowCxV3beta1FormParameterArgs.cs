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
    /// Represents a form parameter.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3beta1FormParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default value of an optional parameter. If the parameter is required, the default value will be ignored.
        /// </summary>
        [Input("defaultValue")]
        public Input<object>? DefaultValue { get; set; }

        /// <summary>
        /// The human-readable name of the parameter, unique within the form.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The entity type of the parameter. Format: `projects/-/locations/-/agents/-/entityTypes/` for system entity types (for example, `projects/-/locations/-/agents/-/entityTypes/sys.date`), or `projects//locations//agents//entityTypes/` for developer entity types.
        /// </summary>
        [Input("entityType", required: true)]
        public Input<string> EntityType { get; set; } = null!;

        /// <summary>
        /// Defines fill behavior for the parameter.
        /// </summary>
        [Input("fillBehavior", required: true)]
        public Input<Inputs.GoogleCloudDialogflowCxV3beta1FormParameterFillBehaviorArgs> FillBehavior { get; set; } = null!;

        /// <summary>
        /// Indicates whether the parameter represents a list of values.
        /// </summary>
        [Input("isList")]
        public Input<bool>? IsList { get; set; }

        /// <summary>
        /// Indicates whether the parameter content should be redacted in log. If redaction is enabled, the parameter content will be replaced by parameter name during logging. Note: the parameter content is subject to redaction if either parameter level redaction or entity type level redaction is enabled.
        /// </summary>
        [Input("redact")]
        public Input<bool>? Redact { get; set; }

        /// <summary>
        /// Indicates whether the parameter is required. Optional parameters will not trigger prompts; however, they are filled if the user specifies them. Required parameters must be filled before form filling concludes.
        /// </summary>
        [Input("required")]
        public Input<bool>? Required { get; set; }

        public GoogleCloudDialogflowCxV3beta1FormParameterArgs()
        {
        }
        public static new GoogleCloudDialogflowCxV3beta1FormParameterArgs Empty => new GoogleCloudDialogflowCxV3beta1FormParameterArgs();
    }
}
