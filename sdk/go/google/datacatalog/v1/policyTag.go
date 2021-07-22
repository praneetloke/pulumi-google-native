// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a policy tag in a taxonomy.
// Auto-naming is currently not supported for this resource.
type PolicyTag struct {
	pulumi.CustomResourceState

	// Resource names of child policy tags of this policy tag.
	ChildPolicyTags pulumi.StringArrayOutput `pulumi:"childPolicyTags"`
	// Description of this policy tag. If not set, defaults to empty. The description must contain only Unicode characters, tabs, newlines, carriage returns and page breaks, and be at most 2000 bytes long when encoded in UTF-8.
	Description pulumi.StringOutput `pulumi:"description"`
	// User-defined name of this policy tag. The name can't start or end with spaces and must be unique within the parent taxonomy, contain only Unicode letters, numbers, underscores, dashes and spaces, and be at most 200 bytes long when encoded in UTF-8.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource name of this policy tag in the URL format. The policy tag manager generates unique taxonomy IDs and policy tag IDs.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource name of this policy tag's parent policy tag. If empty, this is a top level tag. If not set, defaults to an empty string. For example, for the "LatLong" policy tag in the example above, this field contains the resource name of the "Geolocation" policy tag, and, for "Geolocation", this field is empty.
	ParentPolicyTag pulumi.StringOutput `pulumi:"parentPolicyTag"`
}

// NewPolicyTag registers a new resource with the given unique name, arguments, and options.
func NewPolicyTag(ctx *pulumi.Context,
	name string, args *PolicyTagArgs, opts ...pulumi.ResourceOption) (*PolicyTag, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.TaxonomyId == nil {
		return nil, errors.New("invalid value for required argument 'TaxonomyId'")
	}
	var resource PolicyTag
	err := ctx.RegisterResource("google-native:datacatalog/v1:PolicyTag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyTag gets an existing PolicyTag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyTagState, opts ...pulumi.ResourceOption) (*PolicyTag, error) {
	var resource PolicyTag
	err := ctx.ReadResource("google-native:datacatalog/v1:PolicyTag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyTag resources.
type policyTagState struct {
}

type PolicyTagState struct {
}

func (PolicyTagState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyTagState)(nil)).Elem()
}

type policyTagArgs struct {
	// Description of this policy tag. If not set, defaults to empty. The description must contain only Unicode characters, tabs, newlines, carriage returns and page breaks, and be at most 2000 bytes long when encoded in UTF-8.
	Description *string `pulumi:"description"`
	// User-defined name of this policy tag. The name can't start or end with spaces and must be unique within the parent taxonomy, contain only Unicode letters, numbers, underscores, dashes and spaces, and be at most 200 bytes long when encoded in UTF-8.
	DisplayName string  `pulumi:"displayName"`
	Location    *string `pulumi:"location"`
	// Resource name of this policy tag's parent policy tag. If empty, this is a top level tag. If not set, defaults to an empty string. For example, for the "LatLong" policy tag in the example above, this field contains the resource name of the "Geolocation" policy tag, and, for "Geolocation", this field is empty.
	ParentPolicyTag *string `pulumi:"parentPolicyTag"`
	Project         *string `pulumi:"project"`
	TaxonomyId      string  `pulumi:"taxonomyId"`
}

// The set of arguments for constructing a PolicyTag resource.
type PolicyTagArgs struct {
	// Description of this policy tag. If not set, defaults to empty. The description must contain only Unicode characters, tabs, newlines, carriage returns and page breaks, and be at most 2000 bytes long when encoded in UTF-8.
	Description pulumi.StringPtrInput
	// User-defined name of this policy tag. The name can't start or end with spaces and must be unique within the parent taxonomy, contain only Unicode letters, numbers, underscores, dashes and spaces, and be at most 200 bytes long when encoded in UTF-8.
	DisplayName pulumi.StringInput
	Location    pulumi.StringPtrInput
	// Resource name of this policy tag's parent policy tag. If empty, this is a top level tag. If not set, defaults to an empty string. For example, for the "LatLong" policy tag in the example above, this field contains the resource name of the "Geolocation" policy tag, and, for "Geolocation", this field is empty.
	ParentPolicyTag pulumi.StringPtrInput
	Project         pulumi.StringPtrInput
	TaxonomyId      pulumi.StringInput
}

func (PolicyTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyTagArgs)(nil)).Elem()
}

type PolicyTagInput interface {
	pulumi.Input

	ToPolicyTagOutput() PolicyTagOutput
	ToPolicyTagOutputWithContext(ctx context.Context) PolicyTagOutput
}

func (*PolicyTag) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyTag)(nil))
}

func (i *PolicyTag) ToPolicyTagOutput() PolicyTagOutput {
	return i.ToPolicyTagOutputWithContext(context.Background())
}

func (i *PolicyTag) ToPolicyTagOutputWithContext(ctx context.Context) PolicyTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagOutput)
}

type PolicyTagOutput struct {
	*pulumi.OutputState
}

func (PolicyTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyTag)(nil))
}

func (o PolicyTagOutput) ToPolicyTagOutput() PolicyTagOutput {
	return o
}

func (o PolicyTagOutput) ToPolicyTagOutputWithContext(ctx context.Context) PolicyTagOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PolicyTagOutput{})
}
