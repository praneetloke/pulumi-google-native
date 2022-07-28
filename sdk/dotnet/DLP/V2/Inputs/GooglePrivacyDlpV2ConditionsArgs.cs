// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Inputs
{

    /// <summary>
    /// A collection of conditions.
    /// </summary>
    public sealed class GooglePrivacyDlpV2ConditionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("conditions")]
        private InputList<Inputs.GooglePrivacyDlpV2ConditionArgs>? _conditions;

        /// <summary>
        /// A collection of conditions.
        /// </summary>
        public InputList<Inputs.GooglePrivacyDlpV2ConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.GooglePrivacyDlpV2ConditionArgs>());
            set => _conditions = value;
        }

        public GooglePrivacyDlpV2ConditionsArgs()
        {
        }
        public static new GooglePrivacyDlpV2ConditionsArgs Empty => new GooglePrivacyDlpV2ConditionsArgs();
    }
}
