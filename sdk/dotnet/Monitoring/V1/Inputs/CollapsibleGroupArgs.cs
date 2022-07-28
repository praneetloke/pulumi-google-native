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
    /// A widget that groups the other widgets. All widgets that are within the area spanned by the grouping widget are considered member widgets.
    /// </summary>
    public sealed class CollapsibleGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The collapsed state of the widget on first page load.
        /// </summary>
        [Input("collapsed")]
        public Input<bool>? Collapsed { get; set; }

        public CollapsibleGroupArgs()
        {
        }
        public static new CollapsibleGroupArgs Empty => new CollapsibleGroupArgs();
    }
}
