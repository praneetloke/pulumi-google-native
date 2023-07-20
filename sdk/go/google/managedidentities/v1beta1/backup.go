// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Backup for a domain.
// Auto-naming is currently not supported for this resource.
type Backup struct {
	pulumi.CustomResourceState

	// Required. Backup Id, unique name to identify the backups with the following restrictions: * Must be lowercase letters, numbers, and hyphens * Must start with a letter. * Must contain between 1-63 characters. * Must end with a number or a letter. * Must be unique within the domain.
	BackupId pulumi.StringOutput `pulumi:"backupId"`
	// The time the backups was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. A short description of the backup.
	Description pulumi.StringOutput `pulumi:"description"`
	DomainId    pulumi.StringOutput `pulumi:"domainId"`
	// Optional. Resource labels to represent user provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The unique name of the Backup in the form of projects/{project_id}/locations/global/domains/{domain_name}/backups/{name}
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The current state of the backup.
	State pulumi.StringOutput `pulumi:"state"`
	// Additional information about the current status of this backup, if available.
	StatusMessage pulumi.StringOutput `pulumi:"statusMessage"`
	// Indicates whether it’s an on-demand backup or scheduled.
	Type pulumi.StringOutput `pulumi:"type"`
	// Last update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewBackup registers a new resource with the given unique name, arguments, and options.
func NewBackup(ctx *pulumi.Context,
	name string, args *BackupArgs, opts ...pulumi.ResourceOption) (*Backup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupId == nil {
		return nil, errors.New("invalid value for required argument 'BackupId'")
	}
	if args.DomainId == nil {
		return nil, errors.New("invalid value for required argument 'DomainId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"backupId",
		"domainId",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Backup
	err := ctx.RegisterResource("google-native:managedidentities/v1beta1:Backup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackup gets an existing Backup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupState, opts ...pulumi.ResourceOption) (*Backup, error) {
	var resource Backup
	err := ctx.ReadResource("google-native:managedidentities/v1beta1:Backup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Backup resources.
type backupState struct {
}

type BackupState struct {
}

func (BackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupState)(nil)).Elem()
}

type backupArgs struct {
	// Required. Backup Id, unique name to identify the backups with the following restrictions: * Must be lowercase letters, numbers, and hyphens * Must start with a letter. * Must contain between 1-63 characters. * Must end with a number or a letter. * Must be unique within the domain.
	BackupId string `pulumi:"backupId"`
	// Optional. A short description of the backup.
	Description *string `pulumi:"description"`
	DomainId    string  `pulumi:"domainId"`
	// Optional. Resource labels to represent user provided metadata.
	Labels  map[string]string `pulumi:"labels"`
	Project *string           `pulumi:"project"`
}

// The set of arguments for constructing a Backup resource.
type BackupArgs struct {
	// Required. Backup Id, unique name to identify the backups with the following restrictions: * Must be lowercase letters, numbers, and hyphens * Must start with a letter. * Must contain between 1-63 characters. * Must end with a number or a letter. * Must be unique within the domain.
	BackupId pulumi.StringInput
	// Optional. A short description of the backup.
	Description pulumi.StringPtrInput
	DomainId    pulumi.StringInput
	// Optional. Resource labels to represent user provided metadata.
	Labels  pulumi.StringMapInput
	Project pulumi.StringPtrInput
}

func (BackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupArgs)(nil)).Elem()
}

type BackupInput interface {
	pulumi.Input

	ToBackupOutput() BackupOutput
	ToBackupOutputWithContext(ctx context.Context) BackupOutput
}

func (*Backup) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (i *Backup) ToBackupOutput() BackupOutput {
	return i.ToBackupOutputWithContext(context.Background())
}

func (i *Backup) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupOutput)
}

type BackupOutput struct{ *pulumi.OutputState }

func (BackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (o BackupOutput) ToBackupOutput() BackupOutput {
	return o
}

func (o BackupOutput) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return o
}

// Required. Backup Id, unique name to identify the backups with the following restrictions: * Must be lowercase letters, numbers, and hyphens * Must start with a letter. * Must contain between 1-63 characters. * Must end with a number or a letter. * Must be unique within the domain.
func (o BackupOutput) BackupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.BackupId }).(pulumi.StringOutput)
}

// The time the backups was created.
func (o BackupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. A short description of the backup.
func (o BackupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o BackupOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.DomainId }).(pulumi.StringOutput)
}

// Optional. Resource labels to represent user provided metadata.
func (o BackupOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The unique name of the Backup in the form of projects/{project_id}/locations/global/domains/{domain_name}/backups/{name}
func (o BackupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BackupOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The current state of the backup.
func (o BackupOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Additional information about the current status of this backup, if available.
func (o BackupOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.StatusMessage }).(pulumi.StringOutput)
}

// Indicates whether it’s an on-demand backup or scheduled.
func (o BackupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Last update time.
func (o BackupOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupInput)(nil)).Elem(), &Backup{})
	pulumi.RegisterOutputType(BackupOutput{})
}
