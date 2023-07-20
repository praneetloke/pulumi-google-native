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

// Creates key value entries in a key value map scoped to an organization, environment, or API proxy. **Note**: Supported for Apigee hybrid 1.8.x and higher.
type Entry struct {
	pulumi.CustomResourceState

	ApiId         pulumi.StringOutput `pulumi:"apiId"`
	KeyvaluemapId pulumi.StringOutput `pulumi:"keyvaluemapId"`
	// Resource URI that can be used to identify the scope of the key value map entries.
	Name           pulumi.StringOutput `pulumi:"name"`
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Data or payload that is being retrieved and associated with the unique key.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewEntry registers a new resource with the given unique name, arguments, and options.
func NewEntry(ctx *pulumi.Context,
	name string, args *EntryArgs, opts ...pulumi.ResourceOption) (*Entry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.KeyvaluemapId == nil {
		return nil, errors.New("invalid value for required argument 'KeyvaluemapId'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"apiId",
		"keyvaluemapId",
		"organizationId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Entry
	err := ctx.RegisterResource("google-native:apigee/v1:Entry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntry gets an existing Entry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntryState, opts ...pulumi.ResourceOption) (*Entry, error) {
	var resource Entry
	err := ctx.ReadResource("google-native:apigee/v1:Entry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Entry resources.
type entryState struct {
}

type EntryState struct {
}

func (EntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*entryState)(nil)).Elem()
}

type entryArgs struct {
	ApiId         string `pulumi:"apiId"`
	KeyvaluemapId string `pulumi:"keyvaluemapId"`
	// Resource URI that can be used to identify the scope of the key value map entries.
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	// Data or payload that is being retrieved and associated with the unique key.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a Entry resource.
type EntryArgs struct {
	ApiId         pulumi.StringInput
	KeyvaluemapId pulumi.StringInput
	// Resource URI that can be used to identify the scope of the key value map entries.
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
	// Data or payload that is being retrieved and associated with the unique key.
	Value pulumi.StringInput
}

func (EntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entryArgs)(nil)).Elem()
}

type EntryInput interface {
	pulumi.Input

	ToEntryOutput() EntryOutput
	ToEntryOutputWithContext(ctx context.Context) EntryOutput
}

func (*Entry) ElementType() reflect.Type {
	return reflect.TypeOf((**Entry)(nil)).Elem()
}

func (i *Entry) ToEntryOutput() EntryOutput {
	return i.ToEntryOutputWithContext(context.Background())
}

func (i *Entry) ToEntryOutputWithContext(ctx context.Context) EntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryOutput)
}

type EntryOutput struct{ *pulumi.OutputState }

func (EntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Entry)(nil)).Elem()
}

func (o EntryOutput) ToEntryOutput() EntryOutput {
	return o
}

func (o EntryOutput) ToEntryOutputWithContext(ctx context.Context) EntryOutput {
	return o
}

func (o EntryOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *Entry) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

func (o EntryOutput) KeyvaluemapId() pulumi.StringOutput {
	return o.ApplyT(func(v *Entry) pulumi.StringOutput { return v.KeyvaluemapId }).(pulumi.StringOutput)
}

// Resource URI that can be used to identify the scope of the key value map entries.
func (o EntryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Entry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EntryOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Entry) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Data or payload that is being retrieved and associated with the unique key.
func (o EntryOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *Entry) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EntryInput)(nil)).Elem(), &Entry{})
	pulumi.RegisterOutputType(EntryOutput{})
}
