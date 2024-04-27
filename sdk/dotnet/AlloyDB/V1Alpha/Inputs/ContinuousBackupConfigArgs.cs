// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AlloyDB.V1Alpha.Inputs
{

    /// <summary>
    /// ContinuousBackupConfig describes the continuous backups recovery configurations of a cluster.
    /// </summary>
    public sealed class ContinuousBackupConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether ContinuousBackup is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The encryption config can be specified to encrypt the backups with a customer-managed encryption key (CMEK). When this field is not specified, the backup will then use default encryption scheme to protect the user data.
        /// </summary>
        [Input("encryptionConfig")]
        public Input<Inputs.EncryptionConfigArgs>? EncryptionConfig { get; set; }

        /// <summary>
        /// The number of days that are eligible to restore from using PITR. To support the entire recovery window, backups and logs are retained for one day more than the recovery window. If not set, defaults to 14 days.
        /// </summary>
        [Input("recoveryWindowDays")]
        public Input<int>? RecoveryWindowDays { get; set; }

        public ContinuousBackupConfigArgs()
        {
        }
        public static new ContinuousBackupConfigArgs Empty => new ContinuousBackupConfigArgs();
    }
}