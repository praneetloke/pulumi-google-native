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
    /// A quasi-identifier column has a custom_tag, used to know which column in the data corresponds to which column in the statistical model.
    /// </summary>
    public sealed class GooglePrivacyDlpV2QuasiIdentifierFieldArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A column can be tagged with a custom tag. In this case, the user must indicate an auxiliary table that contains statistical information on the possible values of this column (below).
        /// </summary>
        [Input("customTag")]
        public Input<string>? CustomTag { get; set; }

        /// <summary>
        /// Identifies the column.
        /// </summary>
        [Input("field")]
        public Input<Inputs.GooglePrivacyDlpV2FieldIdArgs>? Field { get; set; }

        public GooglePrivacyDlpV2QuasiIdentifierFieldArgs()
        {
        }
        public static new GooglePrivacyDlpV2QuasiIdentifierFieldArgs Empty => new GooglePrivacyDlpV2QuasiIdentifierFieldArgs();
    }
}
