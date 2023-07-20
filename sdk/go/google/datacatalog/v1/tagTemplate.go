// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a tag template. You must enable the Data Catalog API in the project identified by the `parent` parameter. For more information, see [Data Catalog resource project] (https://cloud.google.com/data-catalog/docs/concepts/resource-project).
type TagTemplate struct {
	pulumi.CustomResourceState

	// Display name for this template. Defaults to an empty string. The name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and can't start or end with spaces. The maximum length is 200 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Map of tag template field IDs to the settings for the field. This map is an exhaustive list of the allowed fields. The map must contain at least one field and at most 500 fields. The keys to this map are tag template field IDs. The IDs have the following limitations: * Can contain uppercase and lowercase letters, numbers (0-9) and underscores (_). * Must be at least 1 character and at most 64 characters long. * Must start with a letter or underscore.
	Fields pulumi.StringMapOutput `pulumi:"fields"`
	// Indicates whether tags created with this template are public. Public tags do not require tag template access to appear in ListTags API response. Additionally, you can search for a public tag by value with a simple search query in addition to using a ``tag:`` predicate.
	IsPubliclyReadable pulumi.BoolOutput   `pulumi:"isPubliclyReadable"`
	Location           pulumi.StringOutput `pulumi:"location"`
	// The resource name of the tag template in URL format. Note: The tag template itself and its child resources might not be stored in the location specified in its name.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Required. The ID of the tag template to create. The ID must contain only lowercase letters (a-z), numbers (0-9), or underscores (_), and must start with a letter or underscore. The maximum size is 64 bytes when encoded in UTF-8.
	TagTemplateId pulumi.StringOutput `pulumi:"tagTemplateId"`
}

// NewTagTemplate registers a new resource with the given unique name, arguments, and options.
func NewTagTemplate(ctx *pulumi.Context,
	name string, args *TagTemplateArgs, opts ...pulumi.ResourceOption) (*TagTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fields == nil {
		return nil, errors.New("invalid value for required argument 'Fields'")
	}
	if args.TagTemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TagTemplateId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"tagTemplateId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TagTemplate
	err := ctx.RegisterResource("google-native:datacatalog/v1:TagTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagTemplate gets an existing TagTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagTemplateState, opts ...pulumi.ResourceOption) (*TagTemplate, error) {
	var resource TagTemplate
	err := ctx.ReadResource("google-native:datacatalog/v1:TagTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagTemplate resources.
type tagTemplateState struct {
}

type TagTemplateState struct {
}

func (TagTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagTemplateState)(nil)).Elem()
}

type tagTemplateArgs struct {
	// Display name for this template. Defaults to an empty string. The name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and can't start or end with spaces. The maximum length is 200 characters.
	DisplayName *string `pulumi:"displayName"`
	// Map of tag template field IDs to the settings for the field. This map is an exhaustive list of the allowed fields. The map must contain at least one field and at most 500 fields. The keys to this map are tag template field IDs. The IDs have the following limitations: * Can contain uppercase and lowercase letters, numbers (0-9) and underscores (_). * Must be at least 1 character and at most 64 characters long. * Must start with a letter or underscore.
	Fields map[string]string `pulumi:"fields"`
	// Indicates whether tags created with this template are public. Public tags do not require tag template access to appear in ListTags API response. Additionally, you can search for a public tag by value with a simple search query in addition to using a ``tag:`` predicate.
	IsPubliclyReadable *bool   `pulumi:"isPubliclyReadable"`
	Location           *string `pulumi:"location"`
	// The resource name of the tag template in URL format. Note: The tag template itself and its child resources might not be stored in the location specified in its name.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Required. The ID of the tag template to create. The ID must contain only lowercase letters (a-z), numbers (0-9), or underscores (_), and must start with a letter or underscore. The maximum size is 64 bytes when encoded in UTF-8.
	TagTemplateId string `pulumi:"tagTemplateId"`
}

// The set of arguments for constructing a TagTemplate resource.
type TagTemplateArgs struct {
	// Display name for this template. Defaults to an empty string. The name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and can't start or end with spaces. The maximum length is 200 characters.
	DisplayName pulumi.StringPtrInput
	// Map of tag template field IDs to the settings for the field. This map is an exhaustive list of the allowed fields. The map must contain at least one field and at most 500 fields. The keys to this map are tag template field IDs. The IDs have the following limitations: * Can contain uppercase and lowercase letters, numbers (0-9) and underscores (_). * Must be at least 1 character and at most 64 characters long. * Must start with a letter or underscore.
	Fields pulumi.StringMapInput
	// Indicates whether tags created with this template are public. Public tags do not require tag template access to appear in ListTags API response. Additionally, you can search for a public tag by value with a simple search query in addition to using a ``tag:`` predicate.
	IsPubliclyReadable pulumi.BoolPtrInput
	Location           pulumi.StringPtrInput
	// The resource name of the tag template in URL format. Note: The tag template itself and its child resources might not be stored in the location specified in its name.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Required. The ID of the tag template to create. The ID must contain only lowercase letters (a-z), numbers (0-9), or underscores (_), and must start with a letter or underscore. The maximum size is 64 bytes when encoded in UTF-8.
	TagTemplateId pulumi.StringInput
}

func (TagTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagTemplateArgs)(nil)).Elem()
}

type TagTemplateInput interface {
	pulumi.Input

	ToTagTemplateOutput() TagTemplateOutput
	ToTagTemplateOutputWithContext(ctx context.Context) TagTemplateOutput
}

func (*TagTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**TagTemplate)(nil)).Elem()
}

func (i *TagTemplate) ToTagTemplateOutput() TagTemplateOutput {
	return i.ToTagTemplateOutputWithContext(context.Background())
}

func (i *TagTemplate) ToTagTemplateOutputWithContext(ctx context.Context) TagTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagTemplateOutput)
}

type TagTemplateOutput struct{ *pulumi.OutputState }

func (TagTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagTemplate)(nil)).Elem()
}

func (o TagTemplateOutput) ToTagTemplateOutput() TagTemplateOutput {
	return o
}

func (o TagTemplateOutput) ToTagTemplateOutputWithContext(ctx context.Context) TagTemplateOutput {
	return o
}

// Display name for this template. Defaults to an empty string. The name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and can't start or end with spaces. The maximum length is 200 characters.
func (o TagTemplateOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *TagTemplate) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Map of tag template field IDs to the settings for the field. This map is an exhaustive list of the allowed fields. The map must contain at least one field and at most 500 fields. The keys to this map are tag template field IDs. The IDs have the following limitations: * Can contain uppercase and lowercase letters, numbers (0-9) and underscores (_). * Must be at least 1 character and at most 64 characters long. * Must start with a letter or underscore.
func (o TagTemplateOutput) Fields() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TagTemplate) pulumi.StringMapOutput { return v.Fields }).(pulumi.StringMapOutput)
}

// Indicates whether tags created with this template are public. Public tags do not require tag template access to appear in ListTags API response. Additionally, you can search for a public tag by value with a simple search query in addition to using a “tag:“ predicate.
func (o TagTemplateOutput) IsPubliclyReadable() pulumi.BoolOutput {
	return o.ApplyT(func(v *TagTemplate) pulumi.BoolOutput { return v.IsPubliclyReadable }).(pulumi.BoolOutput)
}

func (o TagTemplateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TagTemplate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the tag template in URL format. Note: The tag template itself and its child resources might not be stored in the location specified in its name.
func (o TagTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TagTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TagTemplateOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TagTemplate) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Required. The ID of the tag template to create. The ID must contain only lowercase letters (a-z), numbers (0-9), or underscores (_), and must start with a letter or underscore. The maximum size is 64 bytes when encoded in UTF-8.
func (o TagTemplateOutput) TagTemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *TagTemplate) pulumi.StringOutput { return v.TagTemplateId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagTemplateInput)(nil)).Elem(), &TagTemplate{})
	pulumi.RegisterOutputType(TagTemplateOutput{})
}
