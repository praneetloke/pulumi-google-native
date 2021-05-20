// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new TagKey. If another request with the same parameters is sent while the original request is in process, the second request will receive an error. A maximum of 300 TagKeys can exist under a parent at any given time.
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
	// Immutable. The resource name of the new TagKey's parent. Must be of the form `organizations/{org_id}`.
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Required. Immutable. The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace. The short name must be 1-63 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
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

	if args.TagKeyId == nil {
		return nil, errors.New("invalid value for required argument 'TagKeyId'")
	}
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
	// Creation time.
	CreateTime *string `pulumi:"createTime"`
	// Optional. User-assigned description of the TagKey. Must not exceed 256 characters. Read-write.
	Description *string `pulumi:"description"`
	// Optional. Entity tag which users can pass to prevent race conditions. This field is always set in server responses. See UpdateTagKeyRequest for details.
	Etag *string `pulumi:"etag"`
	// Immutable. The resource name for a TagKey. Must be in the format `tagKeys/{tag_key_id}`, where `tag_key_id` is the generated numeric id for the TagKey.
	Name *string `pulumi:"name"`
	// Immutable. Namespaced name of the TagKey.
	NamespacedName *string `pulumi:"namespacedName"`
	// Immutable. The resource name of the new TagKey's parent. Must be of the form `organizations/{org_id}`.
	Parent *string `pulumi:"parent"`
	// Required. Immutable. The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace. The short name must be 1-63 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	ShortName *string `pulumi:"shortName"`
	// Update time.
	UpdateTime *string `pulumi:"updateTime"`
}

type TagKeyState struct {
	// Creation time.
	CreateTime pulumi.StringPtrInput
	// Optional. User-assigned description of the TagKey. Must not exceed 256 characters. Read-write.
	Description pulumi.StringPtrInput
	// Optional. Entity tag which users can pass to prevent race conditions. This field is always set in server responses. See UpdateTagKeyRequest for details.
	Etag pulumi.StringPtrInput
	// Immutable. The resource name for a TagKey. Must be in the format `tagKeys/{tag_key_id}`, where `tag_key_id` is the generated numeric id for the TagKey.
	Name pulumi.StringPtrInput
	// Immutable. Namespaced name of the TagKey.
	NamespacedName pulumi.StringPtrInput
	// Immutable. The resource name of the new TagKey's parent. Must be of the form `organizations/{org_id}`.
	Parent pulumi.StringPtrInput
	// Required. Immutable. The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace. The short name must be 1-63 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	ShortName pulumi.StringPtrInput
	// Update time.
	UpdateTime pulumi.StringPtrInput
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
	// Immutable. The resource name of the new TagKey's parent. Must be of the form `organizations/{org_id}`.
	Parent *string `pulumi:"parent"`
	// Required. Immutable. The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace. The short name must be 1-63 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	ShortName    *string `pulumi:"shortName"`
	TagKeyId     string  `pulumi:"tagKeyId"`
	ValidateOnly *string `pulumi:"validateOnly"`
}

// The set of arguments for constructing a TagKey resource.
type TagKeyArgs struct {
	// Optional. User-assigned description of the TagKey. Must not exceed 256 characters. Read-write.
	Description pulumi.StringPtrInput
	// Optional. Entity tag which users can pass to prevent race conditions. This field is always set in server responses. See UpdateTagKeyRequest for details.
	Etag pulumi.StringPtrInput
	// Immutable. The resource name for a TagKey. Must be in the format `tagKeys/{tag_key_id}`, where `tag_key_id` is the generated numeric id for the TagKey.
	Name pulumi.StringPtrInput
	// Immutable. The resource name of the new TagKey's parent. Must be of the form `organizations/{org_id}`.
	Parent pulumi.StringPtrInput
	// Required. Immutable. The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace. The short name must be 1-63 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	ShortName    pulumi.StringPtrInput
	TagKeyId     pulumi.StringInput
	ValidateOnly pulumi.StringPtrInput
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
	return reflect.TypeOf((*TagKey)(nil))
}

func (i *TagKey) ToTagKeyOutput() TagKeyOutput {
	return i.ToTagKeyOutputWithContext(context.Background())
}

func (i *TagKey) ToTagKeyOutputWithContext(ctx context.Context) TagKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyOutput)
}

type TagKeyOutput struct {
	*pulumi.OutputState
}

func (TagKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagKey)(nil))
}

func (o TagKeyOutput) ToTagKeyOutput() TagKeyOutput {
	return o
}

func (o TagKeyOutput) ToTagKeyOutputWithContext(ctx context.Context) TagKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TagKeyOutput{})
}
