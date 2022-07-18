// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkebackup.V1.Inputs
{

    /// <summary>
    /// RetentionPolicy defines a Backup retention policy for a BackupPlan.
    /// </summary>
    public sealed class RetentionPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Minimum age for Backups created via this BackupPlan (in days). This field MUST be an integer value between 0-90 (inclusive). A Backup created under this BackupPlan will NOT be deletable until it reaches Backup's (create_time + backup_delete_lock_days). Updating this field of a BackupPlan does NOT affect existing Backups under it. Backups created AFTER a successful update will inherit the new value. Default: 0 (no delete blocking)
        /// </summary>
        [Input("backupDeleteLockDays")]
        public Input<int>? BackupDeleteLockDays { get; set; }

        /// <summary>
        /// The default maximum age of a Backup created via this BackupPlan. This field MUST be an integer value &gt;= 0 and &lt;= 365. If specified, a Backup created under this BackupPlan will be automatically deleted after its age reaches (create_time + backup_retain_days). If not specified, Backups created under this BackupPlan will NOT be subject to automatic deletion. Updating this field does NOT affect existing Backups under it. Backups created AFTER a successful update will automatically pick up the new value. NOTE: backup_retain_days must be &gt;= backup_delete_lock_days. If cron_schedule is defined, then this must be &lt;= 360 * the creation interval. Default: 0 (no automatic deletion)
        /// </summary>
        [Input("backupRetainDays")]
        public Input<int>? BackupRetainDays { get; set; }

        /// <summary>
        /// This flag denotes whether the retention policy of this BackupPlan is locked. If set to True, no further update is allowed on this policy, including the `locked` field itself. Default: False
        /// </summary>
        [Input("locked")]
        public Input<bool>? Locked { get; set; }

        public RetentionPolicyArgs()
        {
        }
    }
}
