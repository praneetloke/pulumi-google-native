// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified SslCertificate resource in the specified region. Get a list of available SSL certificates by making a list() request.
func LookupRegionSslCertificate(ctx *pulumi.Context, args *LookupRegionSslCertificateArgs, opts ...pulumi.InvokeOption) (*LookupRegionSslCertificateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRegionSslCertificateResult
	err := ctx.Invoke("google-native:compute/beta:getRegionSslCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegionSslCertificateArgs struct {
	Project        *string `pulumi:"project"`
	Region         string  `pulumi:"region"`
	SslCertificate string  `pulumi:"sslCertificate"`
}

type LookupRegionSslCertificateResult struct {
	// A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
	Certificate string `pulumi:"certificate"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Expire time of the certificate. RFC3339
	ExpireTime string `pulumi:"expireTime"`
	// Type of the resource. Always compute#sslCertificate for SSL certificates.
	Kind string `pulumi:"kind"`
	// Configuration and status of a managed SSL certificate.
	Managed SslCertificateManagedSslCertificateResponse `pulumi:"managed"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
	PrivateKey string `pulumi:"privateKey"`
	// URL of the region where the regional SSL Certificate resides. This field is not applicable to global SSL Certificate.
	Region string `pulumi:"region"`
	// [Output only] Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Configuration and status of a self-managed SSL certificate.
	SelfManaged SslCertificateSelfManagedSslCertificateResponse `pulumi:"selfManaged"`
	// Domains associated with the certificate via Subject Alternative Name.
	SubjectAlternativeNames []string `pulumi:"subjectAlternativeNames"`
	// (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
	Type string `pulumi:"type"`
}

func LookupRegionSslCertificateOutput(ctx *pulumi.Context, args LookupRegionSslCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupRegionSslCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegionSslCertificateResult, error) {
			args := v.(LookupRegionSslCertificateArgs)
			r, err := LookupRegionSslCertificate(ctx, &args, opts...)
			var s LookupRegionSslCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegionSslCertificateResultOutput)
}

type LookupRegionSslCertificateOutputArgs struct {
	Project        pulumi.StringPtrInput `pulumi:"project"`
	Region         pulumi.StringInput    `pulumi:"region"`
	SslCertificate pulumi.StringInput    `pulumi:"sslCertificate"`
}

func (LookupRegionSslCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionSslCertificateArgs)(nil)).Elem()
}

type LookupRegionSslCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupRegionSslCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionSslCertificateResult)(nil)).Elem()
}

func (o LookupRegionSslCertificateResultOutput) ToLookupRegionSslCertificateResultOutput() LookupRegionSslCertificateResultOutput {
	return o
}

func (o LookupRegionSslCertificateResultOutput) ToLookupRegionSslCertificateResultOutputWithContext(ctx context.Context) LookupRegionSslCertificateResultOutput {
	return o
}

// A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
func (o LookupRegionSslCertificateResultOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSslCertificateResult) string { return v.Certificate }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupRegionSslCertificateResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSslCertificateResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupRegionSslCertificateResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSslCertificateResult) string { return v.Description }).(pulumi.StringOutput)
}

// Expire time of the certificate. RFC3339
func (o LookupRegionSslCertificateResultOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSslCertificateResult) string { return v.ExpireTime }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#sslCertificate for SSL certificates.
func (o LookupRegionSslCertificateResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSslCertificateResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Configuration and status of a managed SSL certificate.
func (o LookupRegionSslCertificateResultOutput) Managed() SslCertificateManagedSslCertificateResponseOutput {
	return o.ApplyT(func(v LookupRegionSslCertificateResult) SslCertificateManagedSslCertificateResponse { return v.Managed }).(SslCertificateManagedSslCertificateResponseOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupRegionSslCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSslCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

// A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
func (o LookupRegionSslCertificateResultOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSslCertificateResult) string { return v.PrivateKey }).(pulumi.StringOutput)
}

// URL of the region where the regional SSL Certificate resides. This field is not applicable to global SSL Certificate.
func (o LookupRegionSslCertificateResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSslCertificateResult) string { return v.Region }).(pulumi.StringOutput)
}

// [Output only] Server-defined URL for the resource.
func (o LookupRegionSslCertificateResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSslCertificateResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Configuration and status of a self-managed SSL certificate.
func (o LookupRegionSslCertificateResultOutput) SelfManaged() SslCertificateSelfManagedSslCertificateResponseOutput {
	return o.ApplyT(func(v LookupRegionSslCertificateResult) SslCertificateSelfManagedSslCertificateResponse {
		return v.SelfManaged
	}).(SslCertificateSelfManagedSslCertificateResponseOutput)
}

// Domains associated with the certificate via Subject Alternative Name.
func (o LookupRegionSslCertificateResultOutput) SubjectAlternativeNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRegionSslCertificateResult) []string { return v.SubjectAlternativeNames }).(pulumi.StringArrayOutput)
}

// (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
func (o LookupRegionSslCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSslCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegionSslCertificateResultOutput{})
}
