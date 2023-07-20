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

// Creates a new DebugToken for the specified app. For security reasons, after the creation operation completes, the `token` field cannot be updated or retrieved, but you can revoke the debug token using DeleteDebugToken. Each app can have a maximum of 20 debug tokens.
type DebugToken struct {
	pulumi.CustomResourceState

	AppId pulumi.StringOutput `pulumi:"appId"`
	// A human readable display name used to identify this debug token.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The relative resource name of the debug token, in the format: ```projects/{project_number}/apps/{app_id}/debugTokens/{debug_token_id}```
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Input only. Immutable. The secret token itself. Must be provided during creation, and must be a UUID4, case insensitive. This field is immutable once set, and cannot be provided during an UpdateDebugToken request. You can, however, delete this debug token using DeleteDebugToken to revoke it. For security reasons, this field will never be populated in any response.
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewDebugToken registers a new resource with the given unique name, arguments, and options.
func NewDebugToken(ctx *pulumi.Context,
	name string, args *DebugTokenArgs, opts ...pulumi.ResourceOption) (*DebugToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"appId",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DebugToken
	err := ctx.RegisterResource("google-native:firebaseappcheck/v1beta:DebugToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDebugToken gets an existing DebugToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDebugToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DebugTokenState, opts ...pulumi.ResourceOption) (*DebugToken, error) {
	var resource DebugToken
	err := ctx.ReadResource("google-native:firebaseappcheck/v1beta:DebugToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DebugToken resources.
type debugTokenState struct {
}

type DebugTokenState struct {
}

func (DebugTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*debugTokenState)(nil)).Elem()
}

type debugTokenArgs struct {
	AppId string `pulumi:"appId"`
	// A human readable display name used to identify this debug token.
	DisplayName string `pulumi:"displayName"`
	// The relative resource name of the debug token, in the format: ```projects/{project_number}/apps/{app_id}/debugTokens/{debug_token_id}```
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Input only. Immutable. The secret token itself. Must be provided during creation, and must be a UUID4, case insensitive. This field is immutable once set, and cannot be provided during an UpdateDebugToken request. You can, however, delete this debug token using DeleteDebugToken to revoke it. For security reasons, this field will never be populated in any response.
	Token string `pulumi:"token"`
}

// The set of arguments for constructing a DebugToken resource.
type DebugTokenArgs struct {
	AppId pulumi.StringInput
	// A human readable display name used to identify this debug token.
	DisplayName pulumi.StringInput
	// The relative resource name of the debug token, in the format: ```projects/{project_number}/apps/{app_id}/debugTokens/{debug_token_id}```
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Input only. Immutable. The secret token itself. Must be provided during creation, and must be a UUID4, case insensitive. This field is immutable once set, and cannot be provided during an UpdateDebugToken request. You can, however, delete this debug token using DeleteDebugToken to revoke it. For security reasons, this field will never be populated in any response.
	Token pulumi.StringInput
}

func (DebugTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*debugTokenArgs)(nil)).Elem()
}

type DebugTokenInput interface {
	pulumi.Input

	ToDebugTokenOutput() DebugTokenOutput
	ToDebugTokenOutputWithContext(ctx context.Context) DebugTokenOutput
}

func (*DebugToken) ElementType() reflect.Type {
	return reflect.TypeOf((**DebugToken)(nil)).Elem()
}

func (i *DebugToken) ToDebugTokenOutput() DebugTokenOutput {
	return i.ToDebugTokenOutputWithContext(context.Background())
}

func (i *DebugToken) ToDebugTokenOutputWithContext(ctx context.Context) DebugTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DebugTokenOutput)
}

type DebugTokenOutput struct{ *pulumi.OutputState }

func (DebugTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DebugToken)(nil)).Elem()
}

func (o DebugTokenOutput) ToDebugTokenOutput() DebugTokenOutput {
	return o
}

func (o DebugTokenOutput) ToDebugTokenOutputWithContext(ctx context.Context) DebugTokenOutput {
	return o
}

func (o DebugTokenOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *DebugToken) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// A human readable display name used to identify this debug token.
func (o DebugTokenOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DebugToken) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The relative resource name of the debug token, in the format: ```projects/{project_number}/apps/{app_id}/debugTokens/{debug_token_id}```
func (o DebugTokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DebugToken) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DebugTokenOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DebugToken) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Input only. Immutable. The secret token itself. Must be provided during creation, and must be a UUID4, case insensitive. This field is immutable once set, and cannot be provided during an UpdateDebugToken request. You can, however, delete this debug token using DeleteDebugToken to revoke it. For security reasons, this field will never be populated in any response.
func (o DebugTokenOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *DebugToken) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DebugTokenInput)(nil)).Elem(), &DebugToken{})
	pulumi.RegisterOutputType(DebugTokenOutput{})
}
