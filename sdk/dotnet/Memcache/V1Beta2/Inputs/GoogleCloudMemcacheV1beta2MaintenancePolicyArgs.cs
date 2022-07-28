// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Memcache.V1Beta2.Inputs
{

    /// <summary>
    /// Maintenance policy per instance.
    /// </summary>
    public sealed class GoogleCloudMemcacheV1beta2MaintenancePolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of what this policy is for. Create/Update methods return INVALID_ARGUMENT if the length is greater than 512.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("weeklyMaintenanceWindow", required: true)]
        private InputList<Inputs.WeeklyMaintenanceWindowArgs>? _weeklyMaintenanceWindow;

        /// <summary>
        /// Maintenance window that is applied to resources covered by this policy. Minimum 1. For the current version, the maximum number of weekly_maintenance_windows is expected to be one.
        /// </summary>
        public InputList<Inputs.WeeklyMaintenanceWindowArgs> WeeklyMaintenanceWindow
        {
            get => _weeklyMaintenanceWindow ?? (_weeklyMaintenanceWindow = new InputList<Inputs.WeeklyMaintenanceWindowArgs>());
            set => _weeklyMaintenanceWindow = value;
        }

        public GoogleCloudMemcacheV1beta2MaintenancePolicyArgs()
        {
        }
        public static new GoogleCloudMemcacheV1beta2MaintenancePolicyArgs Empty => new GoogleCloudMemcacheV1beta2MaintenancePolicyArgs();
    }
}
