// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Maps a domain to an application. A user must be authorized to administer a domain in order to map it to an application. For a list of available authorized domains, see AuthorizedDomains.ListAuthorizedDomains.
type AppDomainMapping struct {
	pulumi.CustomResourceState

	// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.@OutputOnly
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource records required to configure this domain mapping. These records must be added to the domain's DNS configuration in order to serve the application via this domain mapping.@OutputOnly
	ResourceRecords ResourceRecordResponseArrayOutput `pulumi:"resourceRecords"`
	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	SslSettings SslSettingsResponseOutput `pulumi:"sslSettings"`
}

// NewAppDomainMapping registers a new resource with the given unique name, arguments, and options.
func NewAppDomainMapping(ctx *pulumi.Context,
	name string, args *AppDomainMappingArgs, opts ...pulumi.ResourceOption) (*AppDomainMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.DomainMappingId == nil {
		return nil, errors.New("invalid value for required argument 'DomainMappingId'")
	}
	var resource AppDomainMapping
	err := ctx.RegisterResource("google-native:appengine/v1alpha:AppDomainMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppDomainMapping gets an existing AppDomainMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppDomainMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppDomainMappingState, opts ...pulumi.ResourceOption) (*AppDomainMapping, error) {
	var resource AppDomainMapping
	err := ctx.ReadResource("google-native:appengine/v1alpha:AppDomainMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppDomainMapping resources.
type appDomainMappingState struct {
	// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.@OutputOnly
	Name *string `pulumi:"name"`
	// The resource records required to configure this domain mapping. These records must be added to the domain's DNS configuration in order to serve the application via this domain mapping.@OutputOnly
	ResourceRecords []ResourceRecordResponse `pulumi:"resourceRecords"`
	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	SslSettings *SslSettingsResponse `pulumi:"sslSettings"`
}

type AppDomainMappingState struct {
	// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.@OutputOnly
	Name pulumi.StringPtrInput
	// The resource records required to configure this domain mapping. These records must be added to the domain's DNS configuration in order to serve the application via this domain mapping.@OutputOnly
	ResourceRecords ResourceRecordResponseArrayInput
	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	SslSettings SslSettingsResponsePtrInput
}

func (AppDomainMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*appDomainMappingState)(nil)).Elem()
}

type appDomainMappingArgs struct {
	AppId           string `pulumi:"appId"`
	DomainMappingId string `pulumi:"domainMappingId"`
	// Relative name of the domain serving the application. Example: example.com.
	Id *string `pulumi:"id"`
	// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.@OutputOnly
	Name                 *string `pulumi:"name"`
	NoManagedCertificate *string `pulumi:"noManagedCertificate"`
	OverrideStrategy     *string `pulumi:"overrideStrategy"`
	// The resource records required to configure this domain mapping. These records must be added to the domain's DNS configuration in order to serve the application via this domain mapping.@OutputOnly
	ResourceRecords []ResourceRecord `pulumi:"resourceRecords"`
	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	SslSettings *SslSettings `pulumi:"sslSettings"`
}

// The set of arguments for constructing a AppDomainMapping resource.
type AppDomainMappingArgs struct {
	AppId           pulumi.StringInput
	DomainMappingId pulumi.StringInput
	// Relative name of the domain serving the application. Example: example.com.
	Id pulumi.StringPtrInput
	// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.@OutputOnly
	Name                 pulumi.StringPtrInput
	NoManagedCertificate pulumi.StringPtrInput
	OverrideStrategy     pulumi.StringPtrInput
	// The resource records required to configure this domain mapping. These records must be added to the domain's DNS configuration in order to serve the application via this domain mapping.@OutputOnly
	ResourceRecords ResourceRecordArrayInput
	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	SslSettings SslSettingsPtrInput
}

func (AppDomainMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appDomainMappingArgs)(nil)).Elem()
}

type AppDomainMappingInput interface {
	pulumi.Input

	ToAppDomainMappingOutput() AppDomainMappingOutput
	ToAppDomainMappingOutputWithContext(ctx context.Context) AppDomainMappingOutput
}

func (*AppDomainMapping) ElementType() reflect.Type {
	return reflect.TypeOf((*AppDomainMapping)(nil))
}

func (i *AppDomainMapping) ToAppDomainMappingOutput() AppDomainMappingOutput {
	return i.ToAppDomainMappingOutputWithContext(context.Background())
}

func (i *AppDomainMapping) ToAppDomainMappingOutputWithContext(ctx context.Context) AppDomainMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppDomainMappingOutput)
}

type AppDomainMappingOutput struct {
	*pulumi.OutputState
}

func (AppDomainMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppDomainMapping)(nil))
}

func (o AppDomainMappingOutput) ToAppDomainMappingOutput() AppDomainMappingOutput {
	return o
}

func (o AppDomainMappingOutput) ToAppDomainMappingOutputWithContext(ctx context.Context) AppDomainMappingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AppDomainMappingOutput{})
}
