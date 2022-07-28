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
    /// Instructions regarding the table content being inspected.
    /// </summary>
    public sealed class GooglePrivacyDlpV2TableOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("identifyingFields")]
        private InputList<Inputs.GooglePrivacyDlpV2FieldIdArgs>? _identifyingFields;

        /// <summary>
        /// The columns that are the primary keys for table objects included in ContentItem. A copy of this cell's value will stored alongside alongside each finding so that the finding can be traced to the specific row it came from. No more than 3 may be provided.
        /// </summary>
        public InputList<Inputs.GooglePrivacyDlpV2FieldIdArgs> IdentifyingFields
        {
            get => _identifyingFields ?? (_identifyingFields = new InputList<Inputs.GooglePrivacyDlpV2FieldIdArgs>());
            set => _identifyingFields = value;
        }

        public GooglePrivacyDlpV2TableOptionsArgs()
        {
        }
        public static new GooglePrivacyDlpV2TableOptionsArgs Empty => new GooglePrivacyDlpV2TableOptionsArgs();
    }
}
