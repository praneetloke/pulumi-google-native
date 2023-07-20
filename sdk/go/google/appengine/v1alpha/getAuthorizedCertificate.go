// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified SSL certificate.
func LookupAuthorizedCertificate(ctx *pulumi.Context, args *LookupAuthorizedCertificateArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizedCertificateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAuthorizedCertificateResult
	err := ctx.Invoke("google-native:appengine/v1alpha:getAuthorizedCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizedCertificateArgs struct {
	AppId                   string  `pulumi:"appId"`
	AuthorizedCertificateId string  `pulumi:"authorizedCertificateId"`
	View                    *string `pulumi:"view"`
}

type LookupAuthorizedCertificateResult struct {
	// The SSL certificate serving the AuthorizedCertificate resource. This must be obtained independently from a certificate authority.
	CertificateRawData CertificateRawDataResponse `pulumi:"certificateRawData"`
	// The user-specified display name of the certificate. This is not guaranteed to be unique. Example: My Certificate.
	DisplayName string `pulumi:"displayName"`
	// Aggregate count of the domain mappings with this certificate mapped. This count includes domain mappings on applications for which the user does not have VIEWER permissions.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.
	DomainMappingsCount int `pulumi:"domainMappingsCount"`
	// Topmost applicable domains of this certificate. This certificate applies to these domains and their subdomains. Example: example.com.
	DomainNames []string `pulumi:"domainNames"`
	// The time when this certificate expires. To update the renewal time on this certificate, upload an SSL certificate with a different expiration time using AuthorizedCertificates.UpdateAuthorizedCertificate.
	ExpireTime string `pulumi:"expireTime"`
	// Only applicable if this certificate is managed by App Engine. Managed certificates are tied to the lifecycle of a DomainMapping and cannot be updated or deleted via the AuthorizedCertificates API. If this certificate is manually administered by the user, this field will be empty.
	ManagedCertificate ManagedCertificateResponse `pulumi:"managedCertificate"`
	// Full path to the AuthorizedCertificate resource in the API. Example: apps/myapp/authorizedCertificates/12345.
	Name string `pulumi:"name"`
	// The full paths to user visible Domain Mapping resources that have this certificate mapped. Example: apps/myapp/domainMappings/example.com.This may not represent the full list of mapped domain mappings if the user does not have VIEWER permissions on all of the applications that have this certificate mapped. See domain_mappings_count for a complete count.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.
	VisibleDomainMappings []string `pulumi:"visibleDomainMappings"`
}

func LookupAuthorizedCertificateOutput(ctx *pulumi.Context, args LookupAuthorizedCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizedCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthorizedCertificateResult, error) {
			args := v.(LookupAuthorizedCertificateArgs)
			r, err := LookupAuthorizedCertificate(ctx, &args, opts...)
			var s LookupAuthorizedCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAuthorizedCertificateResultOutput)
}

type LookupAuthorizedCertificateOutputArgs struct {
	AppId                   pulumi.StringInput    `pulumi:"appId"`
	AuthorizedCertificateId pulumi.StringInput    `pulumi:"authorizedCertificateId"`
	View                    pulumi.StringPtrInput `pulumi:"view"`
}

func (LookupAuthorizedCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizedCertificateArgs)(nil)).Elem()
}

type LookupAuthorizedCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizedCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizedCertificateResult)(nil)).Elem()
}

func (o LookupAuthorizedCertificateResultOutput) ToLookupAuthorizedCertificateResultOutput() LookupAuthorizedCertificateResultOutput {
	return o
}

func (o LookupAuthorizedCertificateResultOutput) ToLookupAuthorizedCertificateResultOutputWithContext(ctx context.Context) LookupAuthorizedCertificateResultOutput {
	return o
}

// The SSL certificate serving the AuthorizedCertificate resource. This must be obtained independently from a certificate authority.
func (o LookupAuthorizedCertificateResultOutput) CertificateRawData() CertificateRawDataResponseOutput {
	return o.ApplyT(func(v LookupAuthorizedCertificateResult) CertificateRawDataResponse { return v.CertificateRawData }).(CertificateRawDataResponseOutput)
}

// The user-specified display name of the certificate. This is not guaranteed to be unique. Example: My Certificate.
func (o LookupAuthorizedCertificateResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizedCertificateResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Aggregate count of the domain mappings with this certificate mapped. This count includes domain mappings on applications for which the user does not have VIEWER permissions.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.
func (o LookupAuthorizedCertificateResultOutput) DomainMappingsCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAuthorizedCertificateResult) int { return v.DomainMappingsCount }).(pulumi.IntOutput)
}

// Topmost applicable domains of this certificate. This certificate applies to these domains and their subdomains. Example: example.com.
func (o LookupAuthorizedCertificateResultOutput) DomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthorizedCertificateResult) []string { return v.DomainNames }).(pulumi.StringArrayOutput)
}

// The time when this certificate expires. To update the renewal time on this certificate, upload an SSL certificate with a different expiration time using AuthorizedCertificates.UpdateAuthorizedCertificate.
func (o LookupAuthorizedCertificateResultOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizedCertificateResult) string { return v.ExpireTime }).(pulumi.StringOutput)
}

// Only applicable if this certificate is managed by App Engine. Managed certificates are tied to the lifecycle of a DomainMapping and cannot be updated or deleted via the AuthorizedCertificates API. If this certificate is manually administered by the user, this field will be empty.
func (o LookupAuthorizedCertificateResultOutput) ManagedCertificate() ManagedCertificateResponseOutput {
	return o.ApplyT(func(v LookupAuthorizedCertificateResult) ManagedCertificateResponse { return v.ManagedCertificate }).(ManagedCertificateResponseOutput)
}

// Full path to the AuthorizedCertificate resource in the API. Example: apps/myapp/authorizedCertificates/12345.
func (o LookupAuthorizedCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizedCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

// The full paths to user visible Domain Mapping resources that have this certificate mapped. Example: apps/myapp/domainMappings/example.com.This may not represent the full list of mapped domain mappings if the user does not have VIEWER permissions on all of the applications that have this certificate mapped. See domain_mappings_count for a complete count.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.
func (o LookupAuthorizedCertificateResultOutput) VisibleDomainMappings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthorizedCertificateResult) []string { return v.VisibleDomainMappings }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizedCertificateResultOutput{})
}
