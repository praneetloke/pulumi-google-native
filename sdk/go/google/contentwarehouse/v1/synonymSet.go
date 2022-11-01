// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a SynonymSet for a single context. Throws an ALREADY_EXISTS exception if a synonymset already exists for the context.
type SynonymSet struct {
	pulumi.CustomResourceState

	// This is a freeform field. Example contexts can be "sales," "engineering," "real estate," "accounting," etc. The context can be supplied during search requests.
	Context  pulumi.StringOutput `pulumi:"context"`
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the SynonymSet This is mandatory for google.api.resource. Format: projects/{project_number}/locations/{location}/synonymSets/{context}.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// List of Synonyms for the context.
	Synonyms GoogleCloudContentwarehouseV1SynonymSetSynonymResponseArrayOutput `pulumi:"synonyms"`
}

// NewSynonymSet registers a new resource with the given unique name, arguments, and options.
func NewSynonymSet(ctx *pulumi.Context,
	name string, args *SynonymSetArgs, opts ...pulumi.ResourceOption) (*SynonymSet, error) {
	if args == nil {
		args = &SynonymSetArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource SynonymSet
	err := ctx.RegisterResource("google-native:contentwarehouse/v1:SynonymSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSynonymSet gets an existing SynonymSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSynonymSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SynonymSetState, opts ...pulumi.ResourceOption) (*SynonymSet, error) {
	var resource SynonymSet
	err := ctx.ReadResource("google-native:contentwarehouse/v1:SynonymSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SynonymSet resources.
type synonymSetState struct {
}

type SynonymSetState struct {
}

func (SynonymSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*synonymSetState)(nil)).Elem()
}

type synonymSetArgs struct {
	// This is a freeform field. Example contexts can be "sales," "engineering," "real estate," "accounting," etc. The context can be supplied during search requests.
	Context  *string `pulumi:"context"`
	Location *string `pulumi:"location"`
	// The resource name of the SynonymSet This is mandatory for google.api.resource. Format: projects/{project_number}/locations/{location}/synonymSets/{context}.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// List of Synonyms for the context.
	Synonyms []GoogleCloudContentwarehouseV1SynonymSetSynonym `pulumi:"synonyms"`
}

// The set of arguments for constructing a SynonymSet resource.
type SynonymSetArgs struct {
	// This is a freeform field. Example contexts can be "sales," "engineering," "real estate," "accounting," etc. The context can be supplied during search requests.
	Context  pulumi.StringPtrInput
	Location pulumi.StringPtrInput
	// The resource name of the SynonymSet This is mandatory for google.api.resource. Format: projects/{project_number}/locations/{location}/synonymSets/{context}.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// List of Synonyms for the context.
	Synonyms GoogleCloudContentwarehouseV1SynonymSetSynonymArrayInput
}

func (SynonymSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*synonymSetArgs)(nil)).Elem()
}

type SynonymSetInput interface {
	pulumi.Input

	ToSynonymSetOutput() SynonymSetOutput
	ToSynonymSetOutputWithContext(ctx context.Context) SynonymSetOutput
}

func (*SynonymSet) ElementType() reflect.Type {
	return reflect.TypeOf((**SynonymSet)(nil)).Elem()
}

func (i *SynonymSet) ToSynonymSetOutput() SynonymSetOutput {
	return i.ToSynonymSetOutputWithContext(context.Background())
}

func (i *SynonymSet) ToSynonymSetOutputWithContext(ctx context.Context) SynonymSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynonymSetOutput)
}

type SynonymSetOutput struct{ *pulumi.OutputState }

func (SynonymSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SynonymSet)(nil)).Elem()
}

func (o SynonymSetOutput) ToSynonymSetOutput() SynonymSetOutput {
	return o
}

func (o SynonymSetOutput) ToSynonymSetOutputWithContext(ctx context.Context) SynonymSetOutput {
	return o
}

// This is a freeform field. Example contexts can be "sales," "engineering," "real estate," "accounting," etc. The context can be supplied during search requests.
func (o SynonymSetOutput) Context() pulumi.StringOutput {
	return o.ApplyT(func(v *SynonymSet) pulumi.StringOutput { return v.Context }).(pulumi.StringOutput)
}

func (o SynonymSetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SynonymSet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the SynonymSet This is mandatory for google.api.resource. Format: projects/{project_number}/locations/{location}/synonymSets/{context}.
func (o SynonymSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SynonymSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SynonymSetOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SynonymSet) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// List of Synonyms for the context.
func (o SynonymSetOutput) Synonyms() GoogleCloudContentwarehouseV1SynonymSetSynonymResponseArrayOutput {
	return o.ApplyT(func(v *SynonymSet) GoogleCloudContentwarehouseV1SynonymSetSynonymResponseArrayOutput {
		return v.Synonyms
	}).(GoogleCloudContentwarehouseV1SynonymSetSynonymResponseArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SynonymSetInput)(nil)).Elem(), &SynonymSet{})
	pulumi.RegisterOutputType(SynonymSetOutput{})
}