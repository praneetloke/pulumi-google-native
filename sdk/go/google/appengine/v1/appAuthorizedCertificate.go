// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Uploads the specified SSL certificate.
type AppAuthorizedCertificate struct {
	pulumi.CustomResourceState

	// The SSL certificate serving the AuthorizedCertificate resource. This must be obtained independently from a certificate authority.
	CertificateRawData CertificateRawDataResponseOutput `pulumi:"certificateRawData"`
	// The user-specified display name of the certificate. This is not guaranteed to be unique. Example: My Certificate.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Aggregate count of the domain mappings with this certificate mapped. This count includes domain mappings on applications for which the user does not have VIEWER permissions.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.@OutputOnly
	DomainMappingsCount pulumi.IntOutput `pulumi:"domainMappingsCount"`
	// Topmost applicable domains of this certificate. This certificate applies to these domains and their subdomains. Example: example.com.@OutputOnly
	DomainNames pulumi.StringArrayOutput `pulumi:"domainNames"`
	// The time when this certificate expires. To update the renewal time on this certificate, upload an SSL certificate with a different expiration time using AuthorizedCertificates.UpdateAuthorizedCertificate.@OutputOnly
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// Only applicable if this certificate is managed by App Engine. Managed certificates are tied to the lifecycle of a DomainMapping and cannot be updated or deleted via the AuthorizedCertificates API. If this certificate is manually administered by the user, this field will be empty.@OutputOnly
	ManagedCertificate ManagedCertificateResponseOutput `pulumi:"managedCertificate"`
	// Full path to the AuthorizedCertificate resource in the API. Example: apps/myapp/authorizedCertificates/12345.@OutputOnly
	Name pulumi.StringOutput `pulumi:"name"`
	// The full paths to user visible Domain Mapping resources that have this certificate mapped. Example: apps/myapp/domainMappings/example.com.This may not represent the full list of mapped domain mappings if the user does not have VIEWER permissions on all of the applications that have this certificate mapped. See domain_mappings_count for a complete count.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.@OutputOnly
	VisibleDomainMappings pulumi.StringArrayOutput `pulumi:"visibleDomainMappings"`
}

// NewAppAuthorizedCertificate registers a new resource with the given unique name, arguments, and options.
func NewAppAuthorizedCertificate(ctx *pulumi.Context,
	name string, args *AppAuthorizedCertificateArgs, opts ...pulumi.ResourceOption) (*AppAuthorizedCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.AuthorizedCertificateId == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizedCertificateId'")
	}
	var resource AppAuthorizedCertificate
	err := ctx.RegisterResource("google-native:appengine/v1:AppAuthorizedCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppAuthorizedCertificate gets an existing AppAuthorizedCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppAuthorizedCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppAuthorizedCertificateState, opts ...pulumi.ResourceOption) (*AppAuthorizedCertificate, error) {
	var resource AppAuthorizedCertificate
	err := ctx.ReadResource("google-native:appengine/v1:AppAuthorizedCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppAuthorizedCertificate resources.
type appAuthorizedCertificateState struct {
	// The SSL certificate serving the AuthorizedCertificate resource. This must be obtained independently from a certificate authority.
	CertificateRawData *CertificateRawDataResponse `pulumi:"certificateRawData"`
	// The user-specified display name of the certificate. This is not guaranteed to be unique. Example: My Certificate.
	DisplayName *string `pulumi:"displayName"`
	// Aggregate count of the domain mappings with this certificate mapped. This count includes domain mappings on applications for which the user does not have VIEWER permissions.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.@OutputOnly
	DomainMappingsCount *int `pulumi:"domainMappingsCount"`
	// Topmost applicable domains of this certificate. This certificate applies to these domains and their subdomains. Example: example.com.@OutputOnly
	DomainNames []string `pulumi:"domainNames"`
	// The time when this certificate expires. To update the renewal time on this certificate, upload an SSL certificate with a different expiration time using AuthorizedCertificates.UpdateAuthorizedCertificate.@OutputOnly
	ExpireTime *string `pulumi:"expireTime"`
	// Only applicable if this certificate is managed by App Engine. Managed certificates are tied to the lifecycle of a DomainMapping and cannot be updated or deleted via the AuthorizedCertificates API. If this certificate is manually administered by the user, this field will be empty.@OutputOnly
	ManagedCertificate *ManagedCertificateResponse `pulumi:"managedCertificate"`
	// Full path to the AuthorizedCertificate resource in the API. Example: apps/myapp/authorizedCertificates/12345.@OutputOnly
	Name *string `pulumi:"name"`
	// The full paths to user visible Domain Mapping resources that have this certificate mapped. Example: apps/myapp/domainMappings/example.com.This may not represent the full list of mapped domain mappings if the user does not have VIEWER permissions on all of the applications that have this certificate mapped. See domain_mappings_count for a complete count.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.@OutputOnly
	VisibleDomainMappings []string `pulumi:"visibleDomainMappings"`
}

type AppAuthorizedCertificateState struct {
	// The SSL certificate serving the AuthorizedCertificate resource. This must be obtained independently from a certificate authority.
	CertificateRawData CertificateRawDataResponsePtrInput
	// The user-specified display name of the certificate. This is not guaranteed to be unique. Example: My Certificate.
	DisplayName pulumi.StringPtrInput
	// Aggregate count of the domain mappings with this certificate mapped. This count includes domain mappings on applications for which the user does not have VIEWER permissions.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.@OutputOnly
	DomainMappingsCount pulumi.IntPtrInput
	// Topmost applicable domains of this certificate. This certificate applies to these domains and their subdomains. Example: example.com.@OutputOnly
	DomainNames pulumi.StringArrayInput
	// The time when this certificate expires. To update the renewal time on this certificate, upload an SSL certificate with a different expiration time using AuthorizedCertificates.UpdateAuthorizedCertificate.@OutputOnly
	ExpireTime pulumi.StringPtrInput
	// Only applicable if this certificate is managed by App Engine. Managed certificates are tied to the lifecycle of a DomainMapping and cannot be updated or deleted via the AuthorizedCertificates API. If this certificate is manually administered by the user, this field will be empty.@OutputOnly
	ManagedCertificate ManagedCertificateResponsePtrInput
	// Full path to the AuthorizedCertificate resource in the API. Example: apps/myapp/authorizedCertificates/12345.@OutputOnly
	Name pulumi.StringPtrInput
	// The full paths to user visible Domain Mapping resources that have this certificate mapped. Example: apps/myapp/domainMappings/example.com.This may not represent the full list of mapped domain mappings if the user does not have VIEWER permissions on all of the applications that have this certificate mapped. See domain_mappings_count for a complete count.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.@OutputOnly
	VisibleDomainMappings pulumi.StringArrayInput
}

func (AppAuthorizedCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*appAuthorizedCertificateState)(nil)).Elem()
}

type appAuthorizedCertificateArgs struct {
	AppId                   string `pulumi:"appId"`
	AuthorizedCertificateId string `pulumi:"authorizedCertificateId"`
	// The SSL certificate serving the AuthorizedCertificate resource. This must be obtained independently from a certificate authority.
	CertificateRawData *CertificateRawData `pulumi:"certificateRawData"`
	// The user-specified display name of the certificate. This is not guaranteed to be unique. Example: My Certificate.
	DisplayName *string `pulumi:"displayName"`
	// Aggregate count of the domain mappings with this certificate mapped. This count includes domain mappings on applications for which the user does not have VIEWER permissions.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.@OutputOnly
	DomainMappingsCount *int `pulumi:"domainMappingsCount"`
	// Topmost applicable domains of this certificate. This certificate applies to these domains and their subdomains. Example: example.com.@OutputOnly
	DomainNames []string `pulumi:"domainNames"`
	// The time when this certificate expires. To update the renewal time on this certificate, upload an SSL certificate with a different expiration time using AuthorizedCertificates.UpdateAuthorizedCertificate.@OutputOnly
	ExpireTime *string `pulumi:"expireTime"`
	// Relative name of the certificate. This is a unique value autogenerated on AuthorizedCertificate resource creation. Example: 12345.@OutputOnly
	Id *string `pulumi:"id"`
	// Only applicable if this certificate is managed by App Engine. Managed certificates are tied to the lifecycle of a DomainMapping and cannot be updated or deleted via the AuthorizedCertificates API. If this certificate is manually administered by the user, this field will be empty.@OutputOnly
	ManagedCertificate *ManagedCertificate `pulumi:"managedCertificate"`
	// Full path to the AuthorizedCertificate resource in the API. Example: apps/myapp/authorizedCertificates/12345.@OutputOnly
	Name *string `pulumi:"name"`
	// The full paths to user visible Domain Mapping resources that have this certificate mapped. Example: apps/myapp/domainMappings/example.com.This may not represent the full list of mapped domain mappings if the user does not have VIEWER permissions on all of the applications that have this certificate mapped. See domain_mappings_count for a complete count.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.@OutputOnly
	VisibleDomainMappings []string `pulumi:"visibleDomainMappings"`
}

// The set of arguments for constructing a AppAuthorizedCertificate resource.
type AppAuthorizedCertificateArgs struct {
	AppId                   pulumi.StringInput
	AuthorizedCertificateId pulumi.StringInput
	// The SSL certificate serving the AuthorizedCertificate resource. This must be obtained independently from a certificate authority.
	CertificateRawData CertificateRawDataPtrInput
	// The user-specified display name of the certificate. This is not guaranteed to be unique. Example: My Certificate.
	DisplayName pulumi.StringPtrInput
	// Aggregate count of the domain mappings with this certificate mapped. This count includes domain mappings on applications for which the user does not have VIEWER permissions.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.@OutputOnly
	DomainMappingsCount pulumi.IntPtrInput
	// Topmost applicable domains of this certificate. This certificate applies to these domains and their subdomains. Example: example.com.@OutputOnly
	DomainNames pulumi.StringArrayInput
	// The time when this certificate expires. To update the renewal time on this certificate, upload an SSL certificate with a different expiration time using AuthorizedCertificates.UpdateAuthorizedCertificate.@OutputOnly
	ExpireTime pulumi.StringPtrInput
	// Relative name of the certificate. This is a unique value autogenerated on AuthorizedCertificate resource creation. Example: 12345.@OutputOnly
	Id pulumi.StringPtrInput
	// Only applicable if this certificate is managed by App Engine. Managed certificates are tied to the lifecycle of a DomainMapping and cannot be updated or deleted via the AuthorizedCertificates API. If this certificate is manually administered by the user, this field will be empty.@OutputOnly
	ManagedCertificate ManagedCertificatePtrInput
	// Full path to the AuthorizedCertificate resource in the API. Example: apps/myapp/authorizedCertificates/12345.@OutputOnly
	Name pulumi.StringPtrInput
	// The full paths to user visible Domain Mapping resources that have this certificate mapped. Example: apps/myapp/domainMappings/example.com.This may not represent the full list of mapped domain mappings if the user does not have VIEWER permissions on all of the applications that have this certificate mapped. See domain_mappings_count for a complete count.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.@OutputOnly
	VisibleDomainMappings pulumi.StringArrayInput
}

func (AppAuthorizedCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appAuthorizedCertificateArgs)(nil)).Elem()
}

type AppAuthorizedCertificateInput interface {
	pulumi.Input

	ToAppAuthorizedCertificateOutput() AppAuthorizedCertificateOutput
	ToAppAuthorizedCertificateOutputWithContext(ctx context.Context) AppAuthorizedCertificateOutput
}

func (*AppAuthorizedCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((*AppAuthorizedCertificate)(nil))
}

func (i *AppAuthorizedCertificate) ToAppAuthorizedCertificateOutput() AppAuthorizedCertificateOutput {
	return i.ToAppAuthorizedCertificateOutputWithContext(context.Background())
}

func (i *AppAuthorizedCertificate) ToAppAuthorizedCertificateOutputWithContext(ctx context.Context) AppAuthorizedCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppAuthorizedCertificateOutput)
}

type AppAuthorizedCertificateOutput struct {
	*pulumi.OutputState
}

func (AppAuthorizedCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppAuthorizedCertificate)(nil))
}

func (o AppAuthorizedCertificateOutput) ToAppAuthorizedCertificateOutput() AppAuthorizedCertificateOutput {
	return o
}

func (o AppAuthorizedCertificateOutput) ToAppAuthorizedCertificateOutputWithContext(ctx context.Context) AppAuthorizedCertificateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AppAuthorizedCertificateOutput{})
}
