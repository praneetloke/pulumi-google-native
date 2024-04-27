// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1.Inputs
{

    /// <summary>
    /// Filter for text-based data types like varchar.
    /// </summary>
    public sealed class SourceTextFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The filter will match columns with length smaller than or equal to this number.
        /// </summary>
        [Input("sourceMaxLengthFilter")]
        public Input<string>? SourceMaxLengthFilter { get; set; }

        /// <summary>
        /// Optional. The filter will match columns with length greater than or equal to this number.
        /// </summary>
        [Input("sourceMinLengthFilter")]
        public Input<string>? SourceMinLengthFilter { get; set; }

        public SourceTextFilterArgs()
        {
        }
        public static new SourceTextFilterArgs Empty => new SourceTextFilterArgs();
    }
}