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

// Creates the specified index. A newly created index's initial state is `CREATING`. On completion of the returned google.longrunning.Operation, the state will be `READY`. If the index already exists, the call will return an `ALREADY_EXISTS` status. During creation, the process could result in an error, in which case the index will move to the `ERROR` state. The process can be recovered by fixing the data that caused the error, removing the index with delete, then re-creating the index with create. Indexes with a single field cannot be created.
type Index struct {
	pulumi.CustomResourceState

	// The collection ID to which this index applies. Required.
	CollectionId pulumi.StringOutput `pulumi:"collectionId"`
	DatabaseId   pulumi.StringOutput `pulumi:"databaseId"`
	// The fields to index.
	Fields GoogleFirestoreAdminV1beta1IndexFieldResponseArrayOutput `pulumi:"fields"`
	// The resource name of the index. Output only.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The state of the index. Output only.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewIndex registers a new resource with the given unique name, arguments, and options.
func NewIndex(ctx *pulumi.Context,
	name string, args *IndexArgs, opts ...pulumi.ResourceOption) (*Index, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseId == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"databaseId",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Index
	err := ctx.RegisterResource("google-native:firestore/v1beta1:Index", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIndex gets an existing Index resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IndexState, opts ...pulumi.ResourceOption) (*Index, error) {
	var resource Index
	err := ctx.ReadResource("google-native:firestore/v1beta1:Index", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Index resources.
type indexState struct {
}

type IndexState struct {
}

func (IndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*indexState)(nil)).Elem()
}

type indexArgs struct {
	// The collection ID to which this index applies. Required.
	CollectionId *string `pulumi:"collectionId"`
	DatabaseId   string  `pulumi:"databaseId"`
	// The fields to index.
	Fields []GoogleFirestoreAdminV1beta1IndexField `pulumi:"fields"`
	// The resource name of the index. Output only.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The state of the index. Output only.
	State *IndexStateEnum `pulumi:"state"`
}

// The set of arguments for constructing a Index resource.
type IndexArgs struct {
	// The collection ID to which this index applies. Required.
	CollectionId pulumi.StringPtrInput
	DatabaseId   pulumi.StringInput
	// The fields to index.
	Fields GoogleFirestoreAdminV1beta1IndexFieldArrayInput
	// The resource name of the index. Output only.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The state of the index. Output only.
	State IndexStateEnumPtrInput
}

func (IndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*indexArgs)(nil)).Elem()
}

type IndexInput interface {
	pulumi.Input

	ToIndexOutput() IndexOutput
	ToIndexOutputWithContext(ctx context.Context) IndexOutput
}

func (*Index) ElementType() reflect.Type {
	return reflect.TypeOf((**Index)(nil)).Elem()
}

func (i *Index) ToIndexOutput() IndexOutput {
	return i.ToIndexOutputWithContext(context.Background())
}

func (i *Index) ToIndexOutputWithContext(ctx context.Context) IndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexOutput)
}

type IndexOutput struct{ *pulumi.OutputState }

func (IndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Index)(nil)).Elem()
}

func (o IndexOutput) ToIndexOutput() IndexOutput {
	return o
}

func (o IndexOutput) ToIndexOutputWithContext(ctx context.Context) IndexOutput {
	return o
}

// The collection ID to which this index applies. Required.
func (o IndexOutput) CollectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.CollectionId }).(pulumi.StringOutput)
}

func (o IndexOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.DatabaseId }).(pulumi.StringOutput)
}

// The fields to index.
func (o IndexOutput) Fields() GoogleFirestoreAdminV1beta1IndexFieldResponseArrayOutput {
	return o.ApplyT(func(v *Index) GoogleFirestoreAdminV1beta1IndexFieldResponseArrayOutput { return v.Fields }).(GoogleFirestoreAdminV1beta1IndexFieldResponseArrayOutput)
}

// The resource name of the index. Output only.
func (o IndexOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IndexOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The state of the index. Output only.
func (o IndexOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IndexInput)(nil)).Elem(), &Index{})
	pulumi.RegisterOutputType(IndexOutput{})
}
