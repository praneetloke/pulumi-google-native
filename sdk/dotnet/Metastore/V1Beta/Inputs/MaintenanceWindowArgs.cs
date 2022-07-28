// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Metastore.V1Beta.Inputs
{

    /// <summary>
    /// Maintenance window. This specifies when Dataproc Metastore may perform system maintenance operation to the service.
    /// </summary>
    public sealed class MaintenanceWindowArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The day of week, when the window starts.
        /// </summary>
        [Input("dayOfWeek")]
        public Input<Pulumi.GoogleNative.Metastore.V1Beta.MaintenanceWindowDayOfWeek>? DayOfWeek { get; set; }

        /// <summary>
        /// The hour of day (0-23) when the window starts.
        /// </summary>
        [Input("hourOfDay")]
        public Input<int>? HourOfDay { get; set; }

        public MaintenanceWindowArgs()
        {
        }
        public static new MaintenanceWindowArgs Empty => new MaintenanceWindowArgs();
    }
}
