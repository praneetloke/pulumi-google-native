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
    /// Configuration to suppress records whose suppression conditions evaluate to true.
    /// </summary>
    public sealed class GooglePrivacyDlpV2RecordSuppressionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A condition that when it evaluates to true will result in the record being evaluated to be suppressed from the transformed content.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.GooglePrivacyDlpV2RecordConditionArgs>? Condition { get; set; }

        public GooglePrivacyDlpV2RecordSuppressionArgs()
        {
        }
        public static new GooglePrivacyDlpV2RecordSuppressionArgs Empty => new GooglePrivacyDlpV2RecordSuppressionArgs();
    }
}
