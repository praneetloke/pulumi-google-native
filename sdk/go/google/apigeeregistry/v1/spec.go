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

// Creates a specified spec.
type Spec struct {
	pulumi.CustomResourceState

	// Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	ApiId       pulumi.StringOutput    `pulumi:"apiId"`
	// Required. The ID to use for the spec, which will become the final component of the spec's resource name. This value should be 4-63 characters, and valid characters are /a-z-/. Following AIP-162, IDs must not have the form of a UUID.
	ApiSpecId pulumi.StringOutput `pulumi:"apiSpecId"`
	// Input only. The contents of the spec. Provided by API callers when specs are created or updated. To access the contents of a spec, use GetApiSpecContents.
	Contents pulumi.StringOutput `pulumi:"contents"`
	// Creation timestamp; when the spec resource was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A detailed description.
	Description pulumi.StringOutput `pulumi:"description"`
	// A possibly-hierarchical name used to refer to the spec from other specs.
	Filename pulumi.StringOutput `pulumi:"filename"`
	// A SHA-256 hash of the spec's contents. If the spec is gzipped, this is the hash of the uncompressed spec.
	Hash pulumi.StringOutput `pulumi:"hash"`
	// Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with `apigeeregistry.googleapis.com/` and cannot be changed.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// A style (format) descriptor for this spec that is specified as a Media Type (https://en.wikipedia.org/wiki/Media_type). Possible values include `application/vnd.apigee.proto`, `application/vnd.apigee.openapi`, and `application/vnd.apigee.graphql`, with possible suffixes representing compression types. These hypothetical names are defined in the vendor tree defined in RFC6838 (https://tools.ietf.org/html/rfc6838) and are not final. Content types can specify compression. Currently only GZip compression is supported (indicated with "+gzip").
	MimeType pulumi.StringOutput `pulumi:"mimeType"`
	// Resource name.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Revision creation timestamp; when the represented revision was created.
	RevisionCreateTime pulumi.StringOutput `pulumi:"revisionCreateTime"`
	// Immutable. The revision ID of the spec. A new revision is committed whenever the spec contents are changed. The format is an 8-character hexadecimal string.
	RevisionId pulumi.StringOutput `pulumi:"revisionId"`
	// Last update timestamp: when the represented revision was last modified.
	RevisionUpdateTime pulumi.StringOutput `pulumi:"revisionUpdateTime"`
	// The size of the spec file in bytes. If the spec is gzipped, this is the size of the uncompressed spec.
	SizeBytes pulumi.IntOutput `pulumi:"sizeBytes"`
	// The original source URI of the spec (if one exists). This is an external location that can be used for reference purposes but which may not be authoritative since this external resource may change after the spec is retrieved.
	SourceUri pulumi.StringOutput `pulumi:"sourceUri"`
	VersionId pulumi.StringOutput `pulumi:"versionId"`
}

// NewSpec registers a new resource with the given unique name, arguments, and options.
func NewSpec(ctx *pulumi.Context,
	name string, args *SpecArgs, opts ...pulumi.ResourceOption) (*Spec, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.ApiSpecId == nil {
		return nil, errors.New("invalid value for required argument 'ApiSpecId'")
	}
	if args.VersionId == nil {
		return nil, errors.New("invalid value for required argument 'VersionId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"apiId",
		"apiSpecId",
		"location",
		"project",
		"versionId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Spec
	err := ctx.RegisterResource("google-native:apigeeregistry/v1:Spec", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpec gets an existing Spec resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpec(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpecState, opts ...pulumi.ResourceOption) (*Spec, error) {
	var resource Spec
	err := ctx.ReadResource("google-native:apigeeregistry/v1:Spec", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Spec resources.
type specState struct {
}

type SpecState struct {
}

func (SpecState) ElementType() reflect.Type {
	return reflect.TypeOf((*specState)(nil)).Elem()
}

type specArgs struct {
	// Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Annotations map[string]string `pulumi:"annotations"`
	ApiId       string            `pulumi:"apiId"`
	// Required. The ID to use for the spec, which will become the final component of the spec's resource name. This value should be 4-63 characters, and valid characters are /a-z-/. Following AIP-162, IDs must not have the form of a UUID.
	ApiSpecId string `pulumi:"apiSpecId"`
	// Input only. The contents of the spec. Provided by API callers when specs are created or updated. To access the contents of a spec, use GetApiSpecContents.
	Contents *string `pulumi:"contents"`
	// A detailed description.
	Description *string `pulumi:"description"`
	// A possibly-hierarchical name used to refer to the spec from other specs.
	Filename *string `pulumi:"filename"`
	// Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with `apigeeregistry.googleapis.com/` and cannot be changed.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// A style (format) descriptor for this spec that is specified as a Media Type (https://en.wikipedia.org/wiki/Media_type). Possible values include `application/vnd.apigee.proto`, `application/vnd.apigee.openapi`, and `application/vnd.apigee.graphql`, with possible suffixes representing compression types. These hypothetical names are defined in the vendor tree defined in RFC6838 (https://tools.ietf.org/html/rfc6838) and are not final. Content types can specify compression. Currently only GZip compression is supported (indicated with "+gzip").
	MimeType *string `pulumi:"mimeType"`
	// Resource name.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The original source URI of the spec (if one exists). This is an external location that can be used for reference purposes but which may not be authoritative since this external resource may change after the spec is retrieved.
	SourceUri *string `pulumi:"sourceUri"`
	VersionId string  `pulumi:"versionId"`
}

// The set of arguments for constructing a Spec resource.
type SpecArgs struct {
	// Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Annotations pulumi.StringMapInput
	ApiId       pulumi.StringInput
	// Required. The ID to use for the spec, which will become the final component of the spec's resource name. This value should be 4-63 characters, and valid characters are /a-z-/. Following AIP-162, IDs must not have the form of a UUID.
	ApiSpecId pulumi.StringInput
	// Input only. The contents of the spec. Provided by API callers when specs are created or updated. To access the contents of a spec, use GetApiSpecContents.
	Contents pulumi.StringPtrInput
	// A detailed description.
	Description pulumi.StringPtrInput
	// A possibly-hierarchical name used to refer to the spec from other specs.
	Filename pulumi.StringPtrInput
	// Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with `apigeeregistry.googleapis.com/` and cannot be changed.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// A style (format) descriptor for this spec that is specified as a Media Type (https://en.wikipedia.org/wiki/Media_type). Possible values include `application/vnd.apigee.proto`, `application/vnd.apigee.openapi`, and `application/vnd.apigee.graphql`, with possible suffixes representing compression types. These hypothetical names are defined in the vendor tree defined in RFC6838 (https://tools.ietf.org/html/rfc6838) and are not final. Content types can specify compression. Currently only GZip compression is supported (indicated with "+gzip").
	MimeType pulumi.StringPtrInput
	// Resource name.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The original source URI of the spec (if one exists). This is an external location that can be used for reference purposes but which may not be authoritative since this external resource may change after the spec is retrieved.
	SourceUri pulumi.StringPtrInput
	VersionId pulumi.StringInput
}

func (SpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*specArgs)(nil)).Elem()
}

type SpecInput interface {
	pulumi.Input

	ToSpecOutput() SpecOutput
	ToSpecOutputWithContext(ctx context.Context) SpecOutput
}

func (*Spec) ElementType() reflect.Type {
	return reflect.TypeOf((**Spec)(nil)).Elem()
}

func (i *Spec) ToSpecOutput() SpecOutput {
	return i.ToSpecOutputWithContext(context.Background())
}

func (i *Spec) ToSpecOutputWithContext(ctx context.Context) SpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpecOutput)
}

type SpecOutput struct{ *pulumi.OutputState }

func (SpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Spec)(nil)).Elem()
}

func (o SpecOutput) ToSpecOutput() SpecOutput {
	return o
}

func (o SpecOutput) ToSpecOutputWithContext(ctx context.Context) SpecOutput {
	return o
}

// Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
func (o SpecOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

func (o SpecOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// Required. The ID to use for the spec, which will become the final component of the spec's resource name. This value should be 4-63 characters, and valid characters are /a-z-/. Following AIP-162, IDs must not have the form of a UUID.
func (o SpecOutput) ApiSpecId() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.ApiSpecId }).(pulumi.StringOutput)
}

// Input only. The contents of the spec. Provided by API callers when specs are created or updated. To access the contents of a spec, use GetApiSpecContents.
func (o SpecOutput) Contents() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.Contents }).(pulumi.StringOutput)
}

// Creation timestamp; when the spec resource was created.
func (o SpecOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A detailed description.
func (o SpecOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// A possibly-hierarchical name used to refer to the spec from other specs.
func (o SpecOutput) Filename() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.Filename }).(pulumi.StringOutput)
}

// A SHA-256 hash of the spec's contents. If the spec is gzipped, this is the hash of the uncompressed spec.
func (o SpecOutput) Hash() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.Hash }).(pulumi.StringOutput)
}

// Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with `apigeeregistry.googleapis.com/` and cannot be changed.
func (o SpecOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o SpecOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// A style (format) descriptor for this spec that is specified as a Media Type (https://en.wikipedia.org/wiki/Media_type). Possible values include `application/vnd.apigee.proto`, `application/vnd.apigee.openapi`, and `application/vnd.apigee.graphql`, with possible suffixes representing compression types. These hypothetical names are defined in the vendor tree defined in RFC6838 (https://tools.ietf.org/html/rfc6838) and are not final. Content types can specify compression. Currently only GZip compression is supported (indicated with "+gzip").
func (o SpecOutput) MimeType() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.MimeType }).(pulumi.StringOutput)
}

// Resource name.
func (o SpecOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SpecOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Revision creation timestamp; when the represented revision was created.
func (o SpecOutput) RevisionCreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.RevisionCreateTime }).(pulumi.StringOutput)
}

// Immutable. The revision ID of the spec. A new revision is committed whenever the spec contents are changed. The format is an 8-character hexadecimal string.
func (o SpecOutput) RevisionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.RevisionId }).(pulumi.StringOutput)
}

// Last update timestamp: when the represented revision was last modified.
func (o SpecOutput) RevisionUpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.RevisionUpdateTime }).(pulumi.StringOutput)
}

// The size of the spec file in bytes. If the spec is gzipped, this is the size of the uncompressed spec.
func (o SpecOutput) SizeBytes() pulumi.IntOutput {
	return o.ApplyT(func(v *Spec) pulumi.IntOutput { return v.SizeBytes }).(pulumi.IntOutput)
}

// The original source URI of the spec (if one exists). This is an external location that can be used for reference purposes but which may not be authoritative since this external resource may change after the spec is retrieved.
func (o SpecOutput) SourceUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.SourceUri }).(pulumi.StringOutput)
}

func (o SpecOutput) VersionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.VersionId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpecInput)(nil)).Elem(), &Spec{})
	pulumi.RegisterOutputType(SpecOutput{})
}
