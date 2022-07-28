// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V1.Inputs
{

    /// <summary>
    /// A simplified layout that divides the available space into rows and arranges a set of widgets horizontally in each row.
    /// </summary>
    public sealed class RowLayoutArgs : global::Pulumi.ResourceArgs
    {
        [Input("rows")]
        private InputList<Inputs.RowArgs>? _rows;

        /// <summary>
        /// The rows of content to display.
        /// </summary>
        public InputList<Inputs.RowArgs> Rows
        {
            get => _rows ?? (_rows = new InputList<Inputs.RowArgs>());
            set => _rows = value;
        }

        public RowLayoutArgs()
        {
        }
        public static new RowLayoutArgs Empty => new RowLayoutArgs();
    }
}
