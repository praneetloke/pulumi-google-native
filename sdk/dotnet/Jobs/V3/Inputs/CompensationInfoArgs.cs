// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Jobs.V3.Inputs
{

    /// <summary>
    /// Job compensation details.
    /// </summary>
    public sealed class CompensationInfoArgs : global::Pulumi.ResourceArgs
    {
        [Input("entries")]
        private InputList<Inputs.CompensationEntryArgs>? _entries;

        /// <summary>
        /// Optional. Job compensation information. At most one entry can be of type CompensationInfo.CompensationType.BASE, which is referred as ** base compensation entry ** for the job.
        /// </summary>
        public InputList<Inputs.CompensationEntryArgs> Entries
        {
            get => _entries ?? (_entries = new InputList<Inputs.CompensationEntryArgs>());
            set => _entries = value;
        }

        public CompensationInfoArgs()
        {
        }
        public static new CompensationInfoArgs Empty => new CompensationInfoArgs();
    }
}
