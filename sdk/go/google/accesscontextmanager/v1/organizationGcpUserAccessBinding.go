// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a GcpUserAccessBinding. If the client specifies a name, the server will ignore it. Fails if a resource already exists with the same group_key. Completion of this long-running operation does not necessarily signify that the new binding is deployed onto all affected users, which may take more time.
type OrganizationGcpUserAccessBinding struct {
	pulumi.CustomResourceState

	// Required. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	AccessLevels pulumi.StringArrayOutput `pulumi:"accessLevels"`
	// Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the [G Suite Directory API's Groups resource] (https://developers.google.com/admin-sdk/directory/v1/reference/groups#resource). If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
	GroupKey pulumi.StringOutput `pulumi:"groupKey"`
	// Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by [RFC 3986 Section 2.3](https://tools.ietf.org/html/rfc3986#section-2.3)). Should not be specified by the client during creation. Example: "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewOrganizationGcpUserAccessBinding registers a new resource with the given unique name, arguments, and options.
func NewOrganizationGcpUserAccessBinding(ctx *pulumi.Context,
	name string, args *OrganizationGcpUserAccessBindingArgs, opts ...pulumi.ResourceOption) (*OrganizationGcpUserAccessBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GcpUserAccessBindingId == nil {
		return nil, errors.New("invalid value for required argument 'GcpUserAccessBindingId'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource OrganizationGcpUserAccessBinding
	err := ctx.RegisterResource("google-native:accesscontextmanager/v1:OrganizationGcpUserAccessBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationGcpUserAccessBinding gets an existing OrganizationGcpUserAccessBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationGcpUserAccessBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationGcpUserAccessBindingState, opts ...pulumi.ResourceOption) (*OrganizationGcpUserAccessBinding, error) {
	var resource OrganizationGcpUserAccessBinding
	err := ctx.ReadResource("google-native:accesscontextmanager/v1:OrganizationGcpUserAccessBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationGcpUserAccessBinding resources.
type organizationGcpUserAccessBindingState struct {
	// Required. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	AccessLevels []string `pulumi:"accessLevels"`
	// Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the [G Suite Directory API's Groups resource] (https://developers.google.com/admin-sdk/directory/v1/reference/groups#resource). If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
	GroupKey *string `pulumi:"groupKey"`
	// Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by [RFC 3986 Section 2.3](https://tools.ietf.org/html/rfc3986#section-2.3)). Should not be specified by the client during creation. Example: "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
	Name *string `pulumi:"name"`
}

type OrganizationGcpUserAccessBindingState struct {
	// Required. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	AccessLevels pulumi.StringArrayInput
	// Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the [G Suite Directory API's Groups resource] (https://developers.google.com/admin-sdk/directory/v1/reference/groups#resource). If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
	GroupKey pulumi.StringPtrInput
	// Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by [RFC 3986 Section 2.3](https://tools.ietf.org/html/rfc3986#section-2.3)). Should not be specified by the client during creation. Example: "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
	Name pulumi.StringPtrInput
}

func (OrganizationGcpUserAccessBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationGcpUserAccessBindingState)(nil)).Elem()
}

type organizationGcpUserAccessBindingArgs struct {
	// Required. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	AccessLevels           []string `pulumi:"accessLevels"`
	GcpUserAccessBindingId string   `pulumi:"gcpUserAccessBindingId"`
	// Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the [G Suite Directory API's Groups resource] (https://developers.google.com/admin-sdk/directory/v1/reference/groups#resource). If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
	GroupKey *string `pulumi:"groupKey"`
	// Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by [RFC 3986 Section 2.3](https://tools.ietf.org/html/rfc3986#section-2.3)). Should not be specified by the client during creation. Example: "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
}

// The set of arguments for constructing a OrganizationGcpUserAccessBinding resource.
type OrganizationGcpUserAccessBindingArgs struct {
	// Required. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	AccessLevels           pulumi.StringArrayInput
	GcpUserAccessBindingId pulumi.StringInput
	// Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the [G Suite Directory API's Groups resource] (https://developers.google.com/admin-sdk/directory/v1/reference/groups#resource). If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
	GroupKey pulumi.StringPtrInput
	// Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by [RFC 3986 Section 2.3](https://tools.ietf.org/html/rfc3986#section-2.3)). Should not be specified by the client during creation. Example: "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
}

func (OrganizationGcpUserAccessBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationGcpUserAccessBindingArgs)(nil)).Elem()
}

type OrganizationGcpUserAccessBindingInput interface {
	pulumi.Input

	ToOrganizationGcpUserAccessBindingOutput() OrganizationGcpUserAccessBindingOutput
	ToOrganizationGcpUserAccessBindingOutputWithContext(ctx context.Context) OrganizationGcpUserAccessBindingOutput
}

func (*OrganizationGcpUserAccessBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationGcpUserAccessBinding)(nil))
}

func (i *OrganizationGcpUserAccessBinding) ToOrganizationGcpUserAccessBindingOutput() OrganizationGcpUserAccessBindingOutput {
	return i.ToOrganizationGcpUserAccessBindingOutputWithContext(context.Background())
}

func (i *OrganizationGcpUserAccessBinding) ToOrganizationGcpUserAccessBindingOutputWithContext(ctx context.Context) OrganizationGcpUserAccessBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationGcpUserAccessBindingOutput)
}

type OrganizationGcpUserAccessBindingOutput struct {
	*pulumi.OutputState
}

func (OrganizationGcpUserAccessBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationGcpUserAccessBinding)(nil))
}

func (o OrganizationGcpUserAccessBindingOutput) ToOrganizationGcpUserAccessBindingOutput() OrganizationGcpUserAccessBindingOutput {
	return o
}

func (o OrganizationGcpUserAccessBindingOutput) ToOrganizationGcpUserAccessBindingOutputWithContext(ctx context.Context) OrganizationGcpUserAccessBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationGcpUserAccessBindingOutput{})
}
