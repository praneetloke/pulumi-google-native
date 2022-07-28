// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GameServices.V1.Inputs
{

    /// <summary>
    /// Increment a streamz counter with the specified metric and field names. Metric names should start with a '/', generally be lowercase-only, and end in "_count". Field names should not contain an initial slash. The actual exported metric names will have "/iam/policy" prepended. Field names correspond to IAM request parameters and field values are their respective values. Supported field names: - "authority", which is "[token]" if IAMContext.token is present, otherwise the value of IAMContext.authority_selector if present, and otherwise a representation of IAMContext.principal; or - "iam_principal", a representation of IAMContext.principal even if a token or authority selector is present; or - "" (empty string), resulting in a counter with no fields. Examples: counter { metric: "/debug_access_count" field: "iam_principal" } ==&gt; increment counter /iam/policy/debug_access_count {iam_principal=[value of IAMContext.principal]}
    /// </summary>
    public sealed class CounterOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("customFields")]
        private InputList<Inputs.CustomFieldArgs>? _customFields;

        /// <summary>
        /// Custom fields.
        /// </summary>
        public InputList<Inputs.CustomFieldArgs> CustomFields
        {
            get => _customFields ?? (_customFields = new InputList<Inputs.CustomFieldArgs>());
            set => _customFields = value;
        }

        /// <summary>
        /// The field value to attribute.
        /// </summary>
        [Input("field")]
        public Input<string>? Field { get; set; }

        /// <summary>
        /// The metric to update.
        /// </summary>
        [Input("metric")]
        public Input<string>? Metric { get; set; }

        public CounterOptionsArgs()
        {
        }
        public static new CounterOptionsArgs Empty => new CounterOptionsArgs();
    }
}
