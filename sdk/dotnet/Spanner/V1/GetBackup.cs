// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Spanner.V1
{
    public static class GetBackup
    {
        /// <summary>
        /// Gets metadata on a pending or completed Backup.
        /// </summary>
        public static Task<GetBackupResult> InvokeAsync(GetBackupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackupResult>("google-native:spanner/v1:getBackup", args ?? new GetBackupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets metadata on a pending or completed Backup.
        /// </summary>
        public static Output<GetBackupResult> Invoke(GetBackupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBackupResult>("google-native:spanner/v1:getBackup", args ?? new GetBackupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackupArgs : global::Pulumi.InvokeArgs
    {
        [Input("backupId", required: true)]
        public string BackupId { get; set; } = null!;

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetBackupArgs()
        {
        }
        public static new GetBackupArgs Empty => new GetBackupArgs();
    }

    public sealed class GetBackupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("backupId", required: true)]
        public Input<string> BackupId { get; set; } = null!;

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetBackupInvokeArgs()
        {
        }
        public static new GetBackupInvokeArgs Empty => new GetBackupInvokeArgs();
    }


    [OutputType]
    public sealed class GetBackupResult
    {
        /// <summary>
        /// The time the CreateBackup request is received. If the request does not specify `version_time`, the `version_time` of the backup will be equivalent to the `create_time`.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Required for the CreateBackup operation. Name of the database from which this backup was created. This needs to be in the same instance as the backup. Values are of the form `projects//instances//databases/`.
        /// </summary>
        public readonly string Database;
        /// <summary>
        /// The database dialect information for the backup.
        /// </summary>
        public readonly string DatabaseDialect;
        /// <summary>
        /// The encryption information for the backup.
        /// </summary>
        public readonly Outputs.EncryptionInfoResponse EncryptionInfo;
        /// <summary>
        /// Required for the CreateBackup operation. The expiration time of the backup, with microseconds granularity that must be at least 6 hours and at most 366 days from the time the CreateBackup request is processed. Once the `expire_time` has passed, the backup is eligible to be automatically deleted by Cloud Spanner to free the resources used by the backup.
        /// </summary>
        public readonly string ExpireTime;
        /// <summary>
        /// The max allowed expiration time of the backup, with microseconds granularity. A backup's expiration time can be configured in multiple APIs: CreateBackup, UpdateBackup, CopyBackup. When updating or copying an existing backup, the expiration time specified must be less than `Backup.max_expire_time`.
        /// </summary>
        public readonly string MaxExpireTime;
        /// <summary>
        /// Output only for the CreateBackup operation. Required for the UpdateBackup operation. A globally unique identifier for the backup which cannot be changed. Values are of the form `projects//instances//backups/a-z*[a-z0-9]` The final segment of the name must be between 2 and 60 characters in length. The backup is stored in the location(s) specified in the instance configuration of the instance containing the backup, identified by the prefix of the backup name of the form `projects//instances/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The names of the destination backups being created by copying this source backup. The backup names are of the form `projects//instances//backups/`. Referencing backups may exist in different instances. The existence of any referencing backup prevents the backup from being deleted. When the copy operation is done (either successfully completed or cancelled or the destination backup is deleted), the reference to the backup is removed.
        /// </summary>
        public readonly ImmutableArray<string> ReferencingBackups;
        /// <summary>
        /// The names of the restored databases that reference the backup. The database names are of the form `projects//instances//databases/`. Referencing databases may exist in different instances. The existence of any referencing database prevents the backup from being deleted. When a restored database from the backup enters the `READY` state, the reference to the backup is removed.
        /// </summary>
        public readonly ImmutableArray<string> ReferencingDatabases;
        /// <summary>
        /// Size of the backup in bytes.
        /// </summary>
        public readonly string SizeBytes;
        /// <summary>
        /// The current state of the backup.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The backup will contain an externally consistent copy of the database at the timestamp specified by `version_time`. If `version_time` is not specified, the system will set `version_time` to the `create_time` of the backup.
        /// </summary>
        public readonly string VersionTime;

        [OutputConstructor]
        private GetBackupResult(
            string createTime,

            string database,

            string databaseDialect,

            Outputs.EncryptionInfoResponse encryptionInfo,

            string expireTime,

            string maxExpireTime,

            string name,

            ImmutableArray<string> referencingBackups,

            ImmutableArray<string> referencingDatabases,

            string sizeBytes,

            string state,

            string versionTime)
        {
            CreateTime = createTime;
            Database = database;
            DatabaseDialect = databaseDialect;
            EncryptionInfo = encryptionInfo;
            ExpireTime = expireTime;
            MaxExpireTime = maxExpireTime;
            Name = name;
            ReferencingBackups = referencingBackups;
            ReferencingDatabases = referencingDatabases;
            SizeBytes = sizeBytes;
            State = state;
            VersionTime = versionTime;
        }
    }
}
