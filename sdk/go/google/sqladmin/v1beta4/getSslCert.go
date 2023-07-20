// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta4

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a particular SSL certificate. Does not include the private key (required for usage). The private key must be saved from the response to initial creation.
func LookupSslCert(ctx *pulumi.Context, args *LookupSslCertArgs, opts ...pulumi.InvokeOption) (*LookupSslCertResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSslCertResult
	err := ctx.Invoke("google-native:sqladmin/v1beta4:getSslCert", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSslCertArgs struct {
	Instance        string  `pulumi:"instance"`
	Project         *string `pulumi:"project"`
	Sha1Fingerprint string  `pulumi:"sha1Fingerprint"`
}

type LookupSslCertResult struct {
	// PEM representation.
	Cert string `pulumi:"cert"`
	// Serial number, as extracted from the certificate.
	CertSerialNumber string `pulumi:"certSerialNumber"`
	// User supplied name. Constrained to [a-zA-Z.-_ ]+.
	CommonName string `pulumi:"commonName"`
	// The time when the certificate was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
	CreateTime string `pulumi:"createTime"`
	// The time when the certificate expires in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
	ExpirationTime string `pulumi:"expirationTime"`
	// Name of the database instance.
	Instance string `pulumi:"instance"`
	// This is always `sql#sslCert`.
	Kind string `pulumi:"kind"`
	// The URI of this resource.
	SelfLink string `pulumi:"selfLink"`
	// Sha1 Fingerprint.
	Sha1Fingerprint string `pulumi:"sha1Fingerprint"`
}

func LookupSslCertOutput(ctx *pulumi.Context, args LookupSslCertOutputArgs, opts ...pulumi.InvokeOption) LookupSslCertResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSslCertResult, error) {
			args := v.(LookupSslCertArgs)
			r, err := LookupSslCert(ctx, &args, opts...)
			var s LookupSslCertResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSslCertResultOutput)
}

type LookupSslCertOutputArgs struct {
	Instance        pulumi.StringInput    `pulumi:"instance"`
	Project         pulumi.StringPtrInput `pulumi:"project"`
	Sha1Fingerprint pulumi.StringInput    `pulumi:"sha1Fingerprint"`
}

func (LookupSslCertOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSslCertArgs)(nil)).Elem()
}

type LookupSslCertResultOutput struct{ *pulumi.OutputState }

func (LookupSslCertResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSslCertResult)(nil)).Elem()
}

func (o LookupSslCertResultOutput) ToLookupSslCertResultOutput() LookupSslCertResultOutput {
	return o
}

func (o LookupSslCertResultOutput) ToLookupSslCertResultOutputWithContext(ctx context.Context) LookupSslCertResultOutput {
	return o
}

// PEM representation.
func (o LookupSslCertResultOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslCertResult) string { return v.Cert }).(pulumi.StringOutput)
}

// Serial number, as extracted from the certificate.
func (o LookupSslCertResultOutput) CertSerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslCertResult) string { return v.CertSerialNumber }).(pulumi.StringOutput)
}

// User supplied name. Constrained to [a-zA-Z.-_ ]+.
func (o LookupSslCertResultOutput) CommonName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslCertResult) string { return v.CommonName }).(pulumi.StringOutput)
}

// The time when the certificate was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
func (o LookupSslCertResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslCertResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The time when the certificate expires in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
func (o LookupSslCertResultOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslCertResult) string { return v.ExpirationTime }).(pulumi.StringOutput)
}

// Name of the database instance.
func (o LookupSslCertResultOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslCertResult) string { return v.Instance }).(pulumi.StringOutput)
}

// This is always `sql#sslCert`.
func (o LookupSslCertResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslCertResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The URI of this resource.
func (o LookupSslCertResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslCertResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Sha1 Fingerprint.
func (o LookupSslCertResultOutput) Sha1Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslCertResult) string { return v.Sha1Fingerprint }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSslCertResultOutput{})
}
