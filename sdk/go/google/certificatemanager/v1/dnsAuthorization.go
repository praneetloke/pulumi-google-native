// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new DnsAuthorization in a given project and location.
type DnsAuthorization struct {
	pulumi.CustomResourceState

	// The creation timestamp of a DnsAuthorization.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// One or more paragraphs of text description of a DnsAuthorization.
	Description pulumi.StringOutput `pulumi:"description"`
	// DNS Resource Record that needs to be added to DNS configuration.
	DnsResourceRecord DnsResourceRecordResponseOutput `pulumi:"dnsResourceRecord"`
	// Immutable. A domain which is being authorized. A DnsAuthorization resource covers a single domain and its wildcard, e.g. authorization for `example.com` can be used to issue certificates for `example.com` and `*.example.com`.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Set of labels associated with a DnsAuthorization.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// A user-defined name of the dns authorization. DnsAuthorization names must be unique globally and match pattern `projects/*/locations/*/dnsAuthorizations/*`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The last update timestamp of a DnsAuthorization.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewDnsAuthorization registers a new resource with the given unique name, arguments, and options.
func NewDnsAuthorization(ctx *pulumi.Context,
	name string, args *DnsAuthorizationArgs, opts ...pulumi.ResourceOption) (*DnsAuthorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DnsAuthorizationId == nil {
		return nil, errors.New("invalid value for required argument 'DnsAuthorizationId'")
	}
	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	var resource DnsAuthorization
	err := ctx.RegisterResource("google-native:certificatemanager/v1:DnsAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsAuthorization gets an existing DnsAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsAuthorizationState, opts ...pulumi.ResourceOption) (*DnsAuthorization, error) {
	var resource DnsAuthorization
	err := ctx.ReadResource("google-native:certificatemanager/v1:DnsAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsAuthorization resources.
type dnsAuthorizationState struct {
}

type DnsAuthorizationState struct {
}

func (DnsAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsAuthorizationState)(nil)).Elem()
}

type dnsAuthorizationArgs struct {
	// One or more paragraphs of text description of a DnsAuthorization.
	Description        *string `pulumi:"description"`
	DnsAuthorizationId string  `pulumi:"dnsAuthorizationId"`
	// Immutable. A domain which is being authorized. A DnsAuthorization resource covers a single domain and its wildcard, e.g. authorization for `example.com` can be used to issue certificates for `example.com` and `*.example.com`.
	Domain string `pulumi:"domain"`
	// Set of labels associated with a DnsAuthorization.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// A user-defined name of the dns authorization. DnsAuthorization names must be unique globally and match pattern `projects/*/locations/*/dnsAuthorizations/*`.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a DnsAuthorization resource.
type DnsAuthorizationArgs struct {
	// One or more paragraphs of text description of a DnsAuthorization.
	Description        pulumi.StringPtrInput
	DnsAuthorizationId pulumi.StringInput
	// Immutable. A domain which is being authorized. A DnsAuthorization resource covers a single domain and its wildcard, e.g. authorization for `example.com` can be used to issue certificates for `example.com` and `*.example.com`.
	Domain pulumi.StringInput
	// Set of labels associated with a DnsAuthorization.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// A user-defined name of the dns authorization. DnsAuthorization names must be unique globally and match pattern `projects/*/locations/*/dnsAuthorizations/*`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (DnsAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsAuthorizationArgs)(nil)).Elem()
}

type DnsAuthorizationInput interface {
	pulumi.Input

	ToDnsAuthorizationOutput() DnsAuthorizationOutput
	ToDnsAuthorizationOutputWithContext(ctx context.Context) DnsAuthorizationOutput
}

func (*DnsAuthorization) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsAuthorization)(nil)).Elem()
}

func (i *DnsAuthorization) ToDnsAuthorizationOutput() DnsAuthorizationOutput {
	return i.ToDnsAuthorizationOutputWithContext(context.Background())
}

func (i *DnsAuthorization) ToDnsAuthorizationOutputWithContext(ctx context.Context) DnsAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsAuthorizationOutput)
}

type DnsAuthorizationOutput struct{ *pulumi.OutputState }

func (DnsAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsAuthorization)(nil)).Elem()
}

func (o DnsAuthorizationOutput) ToDnsAuthorizationOutput() DnsAuthorizationOutput {
	return o
}

func (o DnsAuthorizationOutput) ToDnsAuthorizationOutputWithContext(ctx context.Context) DnsAuthorizationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsAuthorizationInput)(nil)).Elem(), &DnsAuthorization{})
	pulumi.RegisterOutputType(DnsAuthorizationOutput{})
}
