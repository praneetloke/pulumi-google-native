// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new TagKey. If another request with the same parameters is sent while the original request is in process, the second request will receive an error. A maximum of 1000 TagKeys can exist under a parent at any given time.
type TagKey struct {
	pulumi.CustomResourceState

	// Creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. User-assigned description of the TagKey. Must not exceed 256 characters. Read-write.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. Entity tag which users can pass to prevent race conditions. This field is always set in server responses. See UpdateTagKeyRequest for details.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Immutable. The resource name for a TagKey. Must be in the format `tagKeys/{tag_key_id}`, where `tag_key_id` is the generated numeric id for the TagKey.
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. Namespaced name of the TagKey.
	NamespacedName pulumi.StringOutput `pulumi:"namespacedName"`
	// Immutable. The resource name of the TagKey's parent. A TagKey can be parented by an Organization or a Project. For a TagKey parented by an Organization, its parent must be in the form `organizations/{org_id}`. For a TagKey parented by a Project, its parent can be in the form `projects/{project_id}` or `projects/{project_number}`.
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Optional. A purpose denotes that this Tag is intended for use in policies of a specific policy engine, and will involve that policy engine in management operations involving this Tag. A purpose does not grant a policy engine exclusive rights to the Tag, and it may be referenced by other policy engines. A purpose cannot be changed once set.
	Purpose pulumi.StringOutput `pulumi:"purpose"`
	// Optional. Purpose data corresponds to the policy system that the tag is intended for. See documentation for `Purpose` for formatting of this field. Purpose data cannot be changed once set.
	PurposeData pulumi.StringMapOutput `pulumi:"purposeData"`
	// Immutable. The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace. The short name must be 1-63 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	ShortName pulumi.StringOutput `pulumi:"shortName"`
	// Update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewTagKey registers a new resource with the given unique name, arguments, and options.
func NewTagKey(ctx *pulumi.Context,
	name string, args *TagKeyArgs, opts ...pulumi.ResourceOption) (*TagKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ShortName == nil {
		return nil, errors.New("invalid value for required argument 'ShortName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TagKey
	err := ctx.RegisterResource("google-native:cloudresourcemanager/v3:TagKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagKey gets an existing TagKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagKeyState, opts ...pulumi.ResourceOption) (*TagKey, error) {
	var resource TagKey
	err := ctx.ReadResource("google-native:cloudresourcemanager/v3:TagKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagKey resources.
type tagKeyState struct {
}

type TagKeyState struct {
}

func (TagKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagKeyState)(nil)).Elem()
}

type tagKeyArgs struct {
	// Optional. User-assigned description of the TagKey. Must not exceed 256 characters. Read-write.
	Description *string `pulumi:"description"`
	// Optional. Entity tag which users can pass to prevent race conditions. This field is always set in server responses. See UpdateTagKeyRequest for details.
	Etag *string `pulumi:"etag"`
	// Immutable. The resource name for a TagKey. Must be in the format `tagKeys/{tag_key_id}`, where `tag_key_id` is the generated numeric id for the TagKey.
	Name *string `pulumi:"name"`
	// Immutable. The resource name of the TagKey's parent. A TagKey can be parented by an Organization or a Project. For a TagKey parented by an Organization, its parent must be in the form `organizations/{org_id}`. For a TagKey parented by a Project, its parent can be in the form `projects/{project_id}` or `projects/{project_number}`.
	Parent *string `pulumi:"parent"`
	// Optional. A purpose denotes that this Tag is intended for use in policies of a specific policy engine, and will involve that policy engine in management operations involving this Tag. A purpose does not grant a policy engine exclusive rights to the Tag, and it may be referenced by other policy engines. A purpose cannot be changed once set.
	Purpose *TagKeyPurpose `pulumi:"purpose"`
	// Optional. Purpose data corresponds to the policy system that the tag is intended for. See documentation for `Purpose` for formatting of this field. Purpose data cannot be changed once set.
	PurposeData map[string]string `pulumi:"purposeData"`
	// Immutable. The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace. The short name must be 1-63 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	ShortName string `pulumi:"shortName"`
}

// The set of arguments for constructing a TagKey resource.
type TagKeyArgs struct {
	// Optional. User-assigned description of the TagKey. Must not exceed 256 characters. Read-write.
	Description pulumi.StringPtrInput
	// Optional. Entity tag which users can pass to prevent race conditions. This field is always set in server responses. See UpdateTagKeyRequest for details.
	Etag pulumi.StringPtrInput
	// Immutable. The resource name for a TagKey. Must be in the format `tagKeys/{tag_key_id}`, where `tag_key_id` is the generated numeric id for the TagKey.
	Name pulumi.StringPtrInput
	// Immutable. The resource name of the TagKey's parent. A TagKey can be parented by an Organization or a Project. For a TagKey parented by an Organization, its parent must be in the form `organizations/{org_id}`. For a TagKey parented by a Project, its parent can be in the form `projects/{project_id}` or `projects/{project_number}`.
	Parent pulumi.StringPtrInput
	// Optional. A purpose denotes that this Tag is intended for use in policies of a specific policy engine, and will involve that policy engine in management operations involving this Tag. A purpose does not grant a policy engine exclusive rights to the Tag, and it may be referenced by other policy engines. A purpose cannot be changed once set.
	Purpose TagKeyPurposePtrInput
	// Optional. Purpose data corresponds to the policy system that the tag is intended for. See documentation for `Purpose` for formatting of this field. Purpose data cannot be changed once set.
	PurposeData pulumi.StringMapInput
	// Immutable. The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace. The short name must be 1-63 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	ShortName pulumi.StringInput
}

func (TagKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagKeyArgs)(nil)).Elem()
}

type TagKeyInput interface {
	pulumi.Input

	ToTagKeyOutput() TagKeyOutput
	ToTagKeyOutputWithContext(ctx context.Context) TagKeyOutput
}

func (*TagKey) ElementType() reflect.Type {
	return reflect.TypeOf((**TagKey)(nil)).Elem()
}

func (i *TagKey) ToTagKeyOutput() TagKeyOutput {
	return i.ToTagKeyOutputWithContext(context.Background())
}

func (i *TagKey) ToTagKeyOutputWithContext(ctx context.Context) TagKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyOutput)
}

type TagKeyOutput struct{ *pulumi.OutputState }

func (TagKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagKey)(nil)).Elem()
}

func (o TagKeyOutput) ToTagKeyOutput() TagKeyOutput {
	return o
}

func (o TagKeyOutput) ToTagKeyOutputWithContext(ctx context.Context) TagKeyOutput {
	return o
}

// Creation time.
func (o TagKeyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TagKey) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. User-assigned description of the TagKey. Must not exceed 256 characters. Read-write.
func (o TagKeyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *TagKey) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Optional. Entity tag which users can pass to prevent race conditions. This field is always set in server responses. See UpdateTagKeyRequest for details.
func (o TagKeyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TagKey) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Immutable. The resource name for a TagKey. Must be in the format `tagKeys/{tag_key_id}`, where `tag_key_id` is the generated numeric id for the TagKey.
func (o TagKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TagKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Immutable. Namespaced name of the TagKey.
func (o TagKeyOutput) NamespacedName() pulumi.StringOutput {
	return o.ApplyT(func(v *TagKey) pulumi.StringOutput { return v.NamespacedName }).(pulumi.StringOutput)
}

// Immutable. The resource name of the TagKey's parent. A TagKey can be parented by an Organization or a Project. For a TagKey parented by an Organization, its parent must be in the form `organizations/{org_id}`. For a TagKey parented by a Project, its parent can be in the form `projects/{project_id}` or `projects/{project_number}`.
func (o TagKeyOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *TagKey) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

// Optional. A purpose denotes that this Tag is intended for use in policies of a specific policy engine, and will involve that policy engine in management operations involving this Tag. A purpose does not grant a policy engine exclusive rights to the Tag, and it may be referenced by other policy engines. A purpose cannot be changed once set.
func (o TagKeyOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v *TagKey) pulumi.StringOutput { return v.Purpose }).(pulumi.StringOutput)
}

// Optional. Purpose data corresponds to the policy system that the tag is intended for. See documentation for `Purpose` for formatting of this field. Purpose data cannot be changed once set.
func (o TagKeyOutput) PurposeData() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TagKey) pulumi.StringMapOutput { return v.PurposeData }).(pulumi.StringMapOutput)
}

// Immutable. The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace. The short name must be 1-63 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
func (o TagKeyOutput) ShortName() pulumi.StringOutput {
	return o.ApplyT(func(v *TagKey) pulumi.StringOutput { return v.ShortName }).(pulumi.StringOutput)
}

// Update time.
func (o TagKeyOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TagKey) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagKeyInput)(nil)).Elem(), &TagKey{})
	pulumi.RegisterOutputType(TagKeyOutput{})
}
