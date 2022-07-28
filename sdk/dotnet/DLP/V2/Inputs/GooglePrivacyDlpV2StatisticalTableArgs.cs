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
    /// An auxiliary table containing statistical information on the relative frequency of different quasi-identifiers values. It has one or several quasi-identifiers columns, and one column that indicates the relative frequency of each quasi-identifier tuple. If a tuple is present in the data but not in the auxiliary table, the corresponding relative frequency is assumed to be zero (and thus, the tuple is highly reidentifiable).
    /// </summary>
    public sealed class GooglePrivacyDlpV2StatisticalTableArgs : global::Pulumi.ResourceArgs
    {
        [Input("quasiIds", required: true)]
        private InputList<Inputs.GooglePrivacyDlpV2QuasiIdentifierFieldArgs>? _quasiIds;

        /// <summary>
        /// Quasi-identifier columns.
        /// </summary>
        public InputList<Inputs.GooglePrivacyDlpV2QuasiIdentifierFieldArgs> QuasiIds
        {
            get => _quasiIds ?? (_quasiIds = new InputList<Inputs.GooglePrivacyDlpV2QuasiIdentifierFieldArgs>());
            set => _quasiIds = value;
        }

        /// <summary>
        /// The relative frequency column must contain a floating-point number between 0 and 1 (inclusive). Null values are assumed to be zero.
        /// </summary>
        [Input("relativeFrequency", required: true)]
        public Input<Inputs.GooglePrivacyDlpV2FieldIdArgs> RelativeFrequency { get; set; } = null!;

        /// <summary>
        /// Auxiliary table location.
        /// </summary>
        [Input("table", required: true)]
        public Input<Inputs.GooglePrivacyDlpV2BigQueryTableArgs> Table { get; set; } = null!;

        public GooglePrivacyDlpV2StatisticalTableArgs()
        {
        }
        public static new GooglePrivacyDlpV2StatisticalTableArgs Empty => new GooglePrivacyDlpV2StatisticalTableArgs();
    }
}
