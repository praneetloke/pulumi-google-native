// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create an SSH public key
// Auto-naming is currently not supported for this resource.
type SshPublicKey struct {
	pulumi.CustomResourceState

	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec pulumi.StringOutput `pulumi:"expirationTimeUsec"`
	// The SHA-256 fingerprint of the SSH public key.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Public key text in SSH format, defined by RFC4253 section 6.6.
	Key pulumi.StringOutput `pulumi:"key"`
	// The canonical resource name.
	Name   pulumi.StringOutput `pulumi:"name"`
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewSshPublicKey registers a new resource with the given unique name, arguments, and options.
func NewSshPublicKey(ctx *pulumi.Context,
	name string, args *SshPublicKeyArgs, opts ...pulumi.ResourceOption) (*SshPublicKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"userId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SshPublicKey
	err := ctx.RegisterResource("google-native:oslogin/v1beta:SshPublicKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSshPublicKey gets an existing SshPublicKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSshPublicKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SshPublicKeyState, opts ...pulumi.ResourceOption) (*SshPublicKey, error) {
	var resource SshPublicKey
	err := ctx.ReadResource("google-native:oslogin/v1beta:SshPublicKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SshPublicKey resources.
type sshPublicKeyState struct {
}

type SshPublicKeyState struct {
}

func (SshPublicKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sshPublicKeyState)(nil)).Elem()
}

type sshPublicKeyArgs struct {
	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec *string `pulumi:"expirationTimeUsec"`
	// Public key text in SSH format, defined by RFC4253 section 6.6.
	Key    *string `pulumi:"key"`
	UserId string  `pulumi:"userId"`
}

// The set of arguments for constructing a SshPublicKey resource.
type SshPublicKeyArgs struct {
	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec pulumi.StringPtrInput
	// Public key text in SSH format, defined by RFC4253 section 6.6.
	Key    pulumi.StringPtrInput
	UserId pulumi.StringInput
}

func (SshPublicKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sshPublicKeyArgs)(nil)).Elem()
}

type SshPublicKeyInput interface {
	pulumi.Input

	ToSshPublicKeyOutput() SshPublicKeyOutput
	ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput
}

func (*SshPublicKey) ElementType() reflect.Type {
	return reflect.TypeOf((**SshPublicKey)(nil)).Elem()
}

func (i *SshPublicKey) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return i.ToSshPublicKeyOutputWithContext(context.Background())
}

func (i *SshPublicKey) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyOutput)
}

type SshPublicKeyOutput struct{ *pulumi.OutputState }

func (SshPublicKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshPublicKey)(nil)).Elem()
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return o
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return o
}

// An expiration time in microseconds since epoch.
func (o SshPublicKeyOutput) ExpirationTimeUsec() pulumi.StringOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringOutput { return v.ExpirationTimeUsec }).(pulumi.StringOutput)
}

// The SHA-256 fingerprint of the SSH public key.
func (o SshPublicKeyOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// Public key text in SSH format, defined by RFC4253 section 6.6.
func (o SshPublicKeyOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The canonical resource name.
func (o SshPublicKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SshPublicKeyOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SshPublicKeyInput)(nil)).Elem(), &SshPublicKey{})
	pulumi.RegisterOutputType(SshPublicKeyOutput{})
}
