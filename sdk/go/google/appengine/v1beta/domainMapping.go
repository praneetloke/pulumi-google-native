// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Maps a domain to an application. A user must be authorized to administer a domain in order to map it to an application. For a list of available authorized domains, see AuthorizedDomains.ListAuthorizedDomains.
// Auto-naming is currently not supported for this resource.
type DomainMapping struct {
	pulumi.CustomResourceState

	AppId pulumi.StringOutput `pulumi:"appId"`
	// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether the domain creation should override any existing mappings for this domain. By default, overrides are rejected.
	OverrideStrategy pulumi.StringPtrOutput `pulumi:"overrideStrategy"`
	// The resource records required to configure this domain mapping. These records must be added to the domain's DNS configuration in order to serve the application via this domain mapping.
	ResourceRecords ResourceRecordResponseArrayOutput `pulumi:"resourceRecords"`
	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	SslSettings SslSettingsResponseOutput `pulumi:"sslSettings"`
}

// NewDomainMapping registers a new resource with the given unique name, arguments, and options.
func NewDomainMapping(ctx *pulumi.Context,
	name string, args *DomainMappingArgs, opts ...pulumi.ResourceOption) (*DomainMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"appId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DomainMapping
	err := ctx.RegisterResource("google-native:appengine/v1beta:DomainMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainMapping gets an existing DomainMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainMappingState, opts ...pulumi.ResourceOption) (*DomainMapping, error) {
	var resource DomainMapping
	err := ctx.ReadResource("google-native:appengine/v1beta:DomainMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainMapping resources.
type domainMappingState struct {
}

type DomainMappingState struct {
}

func (DomainMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainMappingState)(nil)).Elem()
}

type domainMappingArgs struct {
	AppId string `pulumi:"appId"`
	// Relative name of the domain serving the application. Example: example.com.
	Id *string `pulumi:"id"`
	// Whether the domain creation should override any existing mappings for this domain. By default, overrides are rejected.
	OverrideStrategy *string `pulumi:"overrideStrategy"`
	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	SslSettings *SslSettings `pulumi:"sslSettings"`
}

// The set of arguments for constructing a DomainMapping resource.
type DomainMappingArgs struct {
	AppId pulumi.StringInput
	// Relative name of the domain serving the application. Example: example.com.
	Id pulumi.StringPtrInput
	// Whether the domain creation should override any existing mappings for this domain. By default, overrides are rejected.
	OverrideStrategy pulumi.StringPtrInput
	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	SslSettings SslSettingsPtrInput
}

func (DomainMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainMappingArgs)(nil)).Elem()
}

type DomainMappingInput interface {
	pulumi.Input

	ToDomainMappingOutput() DomainMappingOutput
	ToDomainMappingOutputWithContext(ctx context.Context) DomainMappingOutput
}

func (*DomainMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainMapping)(nil)).Elem()
}

func (i *DomainMapping) ToDomainMappingOutput() DomainMappingOutput {
	return i.ToDomainMappingOutputWithContext(context.Background())
}

func (i *DomainMapping) ToDomainMappingOutputWithContext(ctx context.Context) DomainMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMappingOutput)
}

type DomainMappingOutput struct{ *pulumi.OutputState }

func (DomainMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainMapping)(nil)).Elem()
}

func (o DomainMappingOutput) ToDomainMappingOutput() DomainMappingOutput {
	return o
}

func (o DomainMappingOutput) ToDomainMappingOutputWithContext(ctx context.Context) DomainMappingOutput {
	return o
}

func (o DomainMappingOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainMapping) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.
func (o DomainMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether the domain creation should override any existing mappings for this domain. By default, overrides are rejected.
func (o DomainMappingOutput) OverrideStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainMapping) pulumi.StringPtrOutput { return v.OverrideStrategy }).(pulumi.StringPtrOutput)
}

// The resource records required to configure this domain mapping. These records must be added to the domain's DNS configuration in order to serve the application via this domain mapping.
func (o DomainMappingOutput) ResourceRecords() ResourceRecordResponseArrayOutput {
	return o.ApplyT(func(v *DomainMapping) ResourceRecordResponseArrayOutput { return v.ResourceRecords }).(ResourceRecordResponseArrayOutput)
}

// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
func (o DomainMappingOutput) SslSettings() SslSettingsResponseOutput {
	return o.ApplyT(func(v *DomainMapping) SslSettingsResponseOutput { return v.SslSettings }).(SslSettingsResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainMappingInput)(nil)).Elem(), &DomainMapping{})
	pulumi.RegisterOutputType(DomainMappingOutput{})
}
