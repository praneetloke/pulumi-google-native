// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an endpoint attachment. **Note:** Not supported for Apigee hybrid.
// Auto-naming is currently not supported for this resource.
type EndpointAttachment struct {
	pulumi.CustomResourceState

	// Host that can be used in either the HTTP target endpoint directly or as the host in target server.
	Host pulumi.StringOutput `pulumi:"host"`
	// Location of the endpoint attachment.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the endpoint attachment. Use the following structure in your request: `organizations/{org}/endpointAttachments/{endpoint_attachment}`
	Name pulumi.StringOutput `pulumi:"name"`
	// Format: projects/*/regions/*/serviceAttachments/*
	ServiceAttachment pulumi.StringOutput `pulumi:"serviceAttachment"`
}

// NewEndpointAttachment registers a new resource with the given unique name, arguments, and options.
func NewEndpointAttachment(ctx *pulumi.Context,
	name string, args *EndpointAttachmentArgs, opts ...pulumi.ResourceOption) (*EndpointAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource EndpointAttachment
	err := ctx.RegisterResource("google-native:apigee/v1:EndpointAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointAttachment gets an existing EndpointAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointAttachmentState, opts ...pulumi.ResourceOption) (*EndpointAttachment, error) {
	var resource EndpointAttachment
	err := ctx.ReadResource("google-native:apigee/v1:EndpointAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointAttachment resources.
type endpointAttachmentState struct {
}

type EndpointAttachmentState struct {
}

func (EndpointAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointAttachmentState)(nil)).Elem()
}

type endpointAttachmentArgs struct {
	// ID to use for the endpoint attachment. The ID can contain lowercase letters and numbers, must start with a letter, and must be 1-20 characters in length.
	EndpointAttachmentId *string `pulumi:"endpointAttachmentId"`
	// Location of the endpoint attachment.
	Location *string `pulumi:"location"`
	// Name of the endpoint attachment. Use the following structure in your request: `organizations/{org}/endpointAttachments/{endpoint_attachment}`
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	// Format: projects/*/regions/*/serviceAttachments/*
	ServiceAttachment *string `pulumi:"serviceAttachment"`
}

// The set of arguments for constructing a EndpointAttachment resource.
type EndpointAttachmentArgs struct {
	// ID to use for the endpoint attachment. The ID can contain lowercase letters and numbers, must start with a letter, and must be 1-20 characters in length.
	EndpointAttachmentId pulumi.StringPtrInput
	// Location of the endpoint attachment.
	Location pulumi.StringPtrInput
	// Name of the endpoint attachment. Use the following structure in your request: `organizations/{org}/endpointAttachments/{endpoint_attachment}`
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
	// Format: projects/*/regions/*/serviceAttachments/*
	ServiceAttachment pulumi.StringPtrInput
}

func (EndpointAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointAttachmentArgs)(nil)).Elem()
}

type EndpointAttachmentInput interface {
	pulumi.Input

	ToEndpointAttachmentOutput() EndpointAttachmentOutput
	ToEndpointAttachmentOutputWithContext(ctx context.Context) EndpointAttachmentOutput
}

func (*EndpointAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointAttachment)(nil)).Elem()
}

func (i *EndpointAttachment) ToEndpointAttachmentOutput() EndpointAttachmentOutput {
	return i.ToEndpointAttachmentOutputWithContext(context.Background())
}

func (i *EndpointAttachment) ToEndpointAttachmentOutputWithContext(ctx context.Context) EndpointAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAttachmentOutput)
}

type EndpointAttachmentOutput struct{ *pulumi.OutputState }

func (EndpointAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointAttachment)(nil)).Elem()
}

func (o EndpointAttachmentOutput) ToEndpointAttachmentOutput() EndpointAttachmentOutput {
	return o
}

func (o EndpointAttachmentOutput) ToEndpointAttachmentOutputWithContext(ctx context.Context) EndpointAttachmentOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointAttachmentInput)(nil)).Elem(), &EndpointAttachment{})
	pulumi.RegisterOutputType(EndpointAttachmentOutput{})
}
