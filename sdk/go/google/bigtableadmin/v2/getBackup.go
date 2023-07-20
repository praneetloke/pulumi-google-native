// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets metadata on a pending or completed Cloud Bigtable Backup.
func LookupBackup(ctx *pulumi.Context, args *LookupBackupArgs, opts ...pulumi.InvokeOption) (*LookupBackupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBackupResult
	err := ctx.Invoke("google-native:bigtableadmin/v2:getBackup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupArgs struct {
	BackupId   string  `pulumi:"backupId"`
	ClusterId  string  `pulumi:"clusterId"`
	InstanceId string  `pulumi:"instanceId"`
	Project    *string `pulumi:"project"`
}

type LookupBackupResult struct {
	// The encryption information for the backup.
	EncryptionInfo EncryptionInfoResponse `pulumi:"encryptionInfo"`
	// `end_time` is the time that the backup was finished. The row data in the backup will be no newer than this timestamp.
	EndTime string `pulumi:"endTime"`
	// The expiration time of the backup, with microseconds granularity that must be at least 6 hours and at most 30 days from the time the request is received. Once the `expire_time` has passed, Cloud Bigtable will delete the backup and free the resources used by the backup.
	ExpireTime string `pulumi:"expireTime"`
	// A globally unique identifier for the backup which cannot be changed. Values are of the form `projects/{project}/instances/{instance}/clusters/{cluster}/ backups/_a-zA-Z0-9*` The final segment of the name must be between 1 and 50 characters in length. The backup is stored in the cluster identified by the prefix of the backup name of the form `projects/{project}/instances/{instance}/clusters/{cluster}`.
	Name string `pulumi:"name"`
	// Size of the backup in bytes.
	SizeBytes string `pulumi:"sizeBytes"`
	// Name of the backup from which this backup was copied. If a backup is not created by copying a backup, this field will be empty. Values are of the form: projects//instances//backups/.
	SourceBackup string `pulumi:"sourceBackup"`
	// Immutable. Name of the table from which this backup was created. This needs to be in the same instance as the backup. Values are of the form `projects/{project}/instances/{instance}/tables/{source_table}`.
	SourceTable string `pulumi:"sourceTable"`
	// `start_time` is the time that the backup was started (i.e. approximately the time the CreateBackup request is received). The row data in this backup will be no older than this timestamp.
	StartTime string `pulumi:"startTime"`
	// The current state of the backup.
	State string `pulumi:"state"`
}

func LookupBackupOutput(ctx *pulumi.Context, args LookupBackupOutputArgs, opts ...pulumi.InvokeOption) LookupBackupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupResult, error) {
			args := v.(LookupBackupArgs)
			r, err := LookupBackup(ctx, &args, opts...)
			var s LookupBackupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupResultOutput)
}

type LookupBackupOutputArgs struct {
	BackupId   pulumi.StringInput    `pulumi:"backupId"`
	ClusterId  pulumi.StringInput    `pulumi:"clusterId"`
	InstanceId pulumi.StringInput    `pulumi:"instanceId"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupBackupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupArgs)(nil)).Elem()
}

type LookupBackupResultOutput struct{ *pulumi.OutputState }

func (LookupBackupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupResult)(nil)).Elem()
}

func (o LookupBackupResultOutput) ToLookupBackupResultOutput() LookupBackupResultOutput {
	return o
}

func (o LookupBackupResultOutput) ToLookupBackupResultOutputWithContext(ctx context.Context) LookupBackupResultOutput {
	return o
}

// The encryption information for the backup.
func (o LookupBackupResultOutput) EncryptionInfo() EncryptionInfoResponseOutput {
	return o.ApplyT(func(v LookupBackupResult) EncryptionInfoResponse { return v.EncryptionInfo }).(EncryptionInfoResponseOutput)
}

// `end_time` is the time that the backup was finished. The row data in the backup will be no newer than this timestamp.
func (o LookupBackupResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// The expiration time of the backup, with microseconds granularity that must be at least 6 hours and at most 30 days from the time the request is received. Once the `expire_time` has passed, Cloud Bigtable will delete the backup and free the resources used by the backup.
func (o LookupBackupResultOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.ExpireTime }).(pulumi.StringOutput)
}

// A globally unique identifier for the backup which cannot be changed. Values are of the form `projects/{project}/instances/{instance}/clusters/{cluster}/ backups/_a-zA-Z0-9*` The final segment of the name must be between 1 and 50 characters in length. The backup is stored in the cluster identified by the prefix of the backup name of the form `projects/{project}/instances/{instance}/clusters/{cluster}`.
func (o LookupBackupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Size of the backup in bytes.
func (o LookupBackupResultOutput) SizeBytes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.SizeBytes }).(pulumi.StringOutput)
}

// Name of the backup from which this backup was copied. If a backup is not created by copying a backup, this field will be empty. Values are of the form: projects//instances//backups/.
func (o LookupBackupResultOutput) SourceBackup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.SourceBackup }).(pulumi.StringOutput)
}

// Immutable. Name of the table from which this backup was created. This needs to be in the same instance as the backup. Values are of the form `projects/{project}/instances/{instance}/tables/{source_table}`.
func (o LookupBackupResultOutput) SourceTable() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.SourceTable }).(pulumi.StringOutput)
}

// `start_time` is the time that the backup was started (i.e. approximately the time the CreateBackup request is received). The row data in this backup will be no older than this timestamp.
func (o LookupBackupResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// The current state of the backup.
func (o LookupBackupResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupResultOutput{})
}
