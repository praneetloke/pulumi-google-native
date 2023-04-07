// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a new CertificateAuthority in a given Project and Location.
// Auto-naming is currently not supported for this resource.
type CertificateAuthority struct {
	pulumi.CustomResourceState

	// URLs for accessing content published by this CA, such as the CA certificate and CRLs.
	AccessUrls AccessUrlsResponseOutput `pulumi:"accessUrls"`
	// A structured description of this CertificateAuthority's CA certificate and its issuers. Ordered as self-to-root.
	CaCertificateDescriptions CertificateDescriptionResponseArrayOutput `pulumi:"caCertificateDescriptions"`
	CaPoolId                  pulumi.StringOutput                       `pulumi:"caPoolId"`
	// Required. It must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
	CertificateAuthorityId pulumi.StringOutput `pulumi:"certificateAuthorityId"`
	// Immutable. The config used to create a self-signed X.509 certificate or CSR.
	Config CertificateConfigResponseOutput `pulumi:"config"`
	// The time at which this CertificateAuthority was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The time at which this CertificateAuthority was soft deleted, if it is in the DELETED state.
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// The time at which this CertificateAuthority will be permanently purged, if it is in the DELETED state.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// Immutable. The name of a Cloud Storage bucket where this CertificateAuthority will publish content, such as the CA certificate and CRLs. This must be a bucket name, without any prefixes (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named `my-bucket`, you would simply specify `my-bucket`. If not specified, a managed bucket will be created.
	GcsBucket pulumi.StringOutput `pulumi:"gcsBucket"`
	// Immutable. Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA certificate. Otherwise, it is used to sign a CSR.
	KeySpec KeyVersionSpecResponseOutput `pulumi:"keySpec"`
	// Optional. Labels with user-defined metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Immutable. The desired lifetime of the CA certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate.
	Lifetime pulumi.StringOutput `pulumi:"lifetime"`
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name for this CertificateAuthority in the format `projects/*/locations/*/caPools/*/certificateAuthorities/*`.
	Name pulumi.StringOutput `pulumi:"name"`
	// This CertificateAuthority's certificate chain, including the current CertificateAuthority's certificate. Ordered such that the root issuer is the final element (consistent with RFC 5246). For a self-signed CA, this will only list the current CertificateAuthority's certificate.
	PemCaCertificates pulumi.StringArrayOutput `pulumi:"pemCaCertificates"`
	Project           pulumi.StringOutput      `pulumi:"project"`
	// Optional. An ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// The State for this CertificateAuthority.
	State pulumi.StringOutput `pulumi:"state"`
	// Optional. If this is a subordinate CertificateAuthority, this field will be set with the subordinate configuration, which describes its issuers. This may be updated, but this CertificateAuthority must continue to validate.
	SubordinateConfig SubordinateConfigResponseOutput `pulumi:"subordinateConfig"`
	// The CaPool.Tier of the CaPool that includes this CertificateAuthority.
	Tier pulumi.StringOutput `pulumi:"tier"`
	// Immutable. The Type of this CertificateAuthority.
	Type pulumi.StringOutput `pulumi:"type"`
	// The time at which this CertificateAuthority was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewCertificateAuthority registers a new resource with the given unique name, arguments, and options.
func NewCertificateAuthority(ctx *pulumi.Context,
	name string, args *CertificateAuthorityArgs, opts ...pulumi.ResourceOption) (*CertificateAuthority, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaPoolId == nil {
		return nil, errors.New("invalid value for required argument 'CaPoolId'")
	}
	if args.CertificateAuthorityId == nil {
		return nil, errors.New("invalid value for required argument 'CertificateAuthorityId'")
	}
	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.KeySpec == nil {
		return nil, errors.New("invalid value for required argument 'KeySpec'")
	}
	if args.Lifetime == nil {
		return nil, errors.New("invalid value for required argument 'Lifetime'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"caPoolId",
		"certificateAuthorityId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource CertificateAuthority
	err := ctx.RegisterResource("google-native:privateca/v1:CertificateAuthority", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateAuthority gets an existing CertificateAuthority resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateAuthority(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateAuthorityState, opts ...pulumi.ResourceOption) (*CertificateAuthority, error) {
	var resource CertificateAuthority
	err := ctx.ReadResource("google-native:privateca/v1:CertificateAuthority", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateAuthority resources.
type certificateAuthorityState struct {
}

type CertificateAuthorityState struct {
}

func (CertificateAuthorityState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateAuthorityState)(nil)).Elem()
}

type certificateAuthorityArgs struct {
	CaPoolId string `pulumi:"caPoolId"`
	// Required. It must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
	CertificateAuthorityId string `pulumi:"certificateAuthorityId"`
	// Immutable. The config used to create a self-signed X.509 certificate or CSR.
	Config CertificateConfig `pulumi:"config"`
	// Immutable. The name of a Cloud Storage bucket where this CertificateAuthority will publish content, such as the CA certificate and CRLs. This must be a bucket name, without any prefixes (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named `my-bucket`, you would simply specify `my-bucket`. If not specified, a managed bucket will be created.
	GcsBucket *string `pulumi:"gcsBucket"`
	// Immutable. Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA certificate. Otherwise, it is used to sign a CSR.
	KeySpec KeyVersionSpec `pulumi:"keySpec"`
	// Optional. Labels with user-defined metadata.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. The desired lifetime of the CA certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate.
	Lifetime string  `pulumi:"lifetime"`
	Location *string `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// Optional. An ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Optional. If this is a subordinate CertificateAuthority, this field will be set with the subordinate configuration, which describes its issuers. This may be updated, but this CertificateAuthority must continue to validate.
	SubordinateConfig *SubordinateConfig `pulumi:"subordinateConfig"`
	// Immutable. The Type of this CertificateAuthority.
	Type CertificateAuthorityType `pulumi:"type"`
}

// The set of arguments for constructing a CertificateAuthority resource.
type CertificateAuthorityArgs struct {
	CaPoolId pulumi.StringInput
	// Required. It must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
	CertificateAuthorityId pulumi.StringInput
	// Immutable. The config used to create a self-signed X.509 certificate or CSR.
	Config CertificateConfigInput
	// Immutable. The name of a Cloud Storage bucket where this CertificateAuthority will publish content, such as the CA certificate and CRLs. This must be a bucket name, without any prefixes (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named `my-bucket`, you would simply specify `my-bucket`. If not specified, a managed bucket will be created.
	GcsBucket pulumi.StringPtrInput
	// Immutable. Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA certificate. Otherwise, it is used to sign a CSR.
	KeySpec KeyVersionSpecInput
	// Optional. Labels with user-defined metadata.
	Labels pulumi.StringMapInput
	// Immutable. The desired lifetime of the CA certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate.
	Lifetime pulumi.StringInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// Optional. An ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Optional. If this is a subordinate CertificateAuthority, this field will be set with the subordinate configuration, which describes its issuers. This may be updated, but this CertificateAuthority must continue to validate.
	SubordinateConfig SubordinateConfigPtrInput
	// Immutable. The Type of this CertificateAuthority.
	Type CertificateAuthorityTypeInput
}

func (CertificateAuthorityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateAuthorityArgs)(nil)).Elem()
}

type CertificateAuthorityInput interface {
	pulumi.Input

	ToCertificateAuthorityOutput() CertificateAuthorityOutput
	ToCertificateAuthorityOutputWithContext(ctx context.Context) CertificateAuthorityOutput
}

func (*CertificateAuthority) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateAuthority)(nil)).Elem()
}

func (i *CertificateAuthority) ToCertificateAuthorityOutput() CertificateAuthorityOutput {
	return i.ToCertificateAuthorityOutputWithContext(context.Background())
}

func (i *CertificateAuthority) ToCertificateAuthorityOutputWithContext(ctx context.Context) CertificateAuthorityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateAuthorityOutput)
}

type CertificateAuthorityOutput struct{ *pulumi.OutputState }

func (CertificateAuthorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateAuthority)(nil)).Elem()
}

func (o CertificateAuthorityOutput) ToCertificateAuthorityOutput() CertificateAuthorityOutput {
	return o
}

func (o CertificateAuthorityOutput) ToCertificateAuthorityOutputWithContext(ctx context.Context) CertificateAuthorityOutput {
	return o
}

// URLs for accessing content published by this CA, such as the CA certificate and CRLs.
func (o CertificateAuthorityOutput) AccessUrls() AccessUrlsResponseOutput {
	return o.ApplyT(func(v *CertificateAuthority) AccessUrlsResponseOutput { return v.AccessUrls }).(AccessUrlsResponseOutput)
}

// A structured description of this CertificateAuthority's CA certificate and its issuers. Ordered as self-to-root.
func (o CertificateAuthorityOutput) CaCertificateDescriptions() CertificateDescriptionResponseArrayOutput {
	return o.ApplyT(func(v *CertificateAuthority) CertificateDescriptionResponseArrayOutput {
		return v.CaCertificateDescriptions
	}).(CertificateDescriptionResponseArrayOutput)
}

func (o CertificateAuthorityOutput) CaPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.CaPoolId }).(pulumi.StringOutput)
}

// Required. It must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
func (o CertificateAuthorityOutput) CertificateAuthorityId() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.CertificateAuthorityId }).(pulumi.StringOutput)
}

// Immutable. The config used to create a self-signed X.509 certificate or CSR.
func (o CertificateAuthorityOutput) Config() CertificateConfigResponseOutput {
	return o.ApplyT(func(v *CertificateAuthority) CertificateConfigResponseOutput { return v.Config }).(CertificateConfigResponseOutput)
}

// The time at which this CertificateAuthority was created.
func (o CertificateAuthorityOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The time at which this CertificateAuthority was soft deleted, if it is in the DELETED state.
func (o CertificateAuthorityOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.DeleteTime }).(pulumi.StringOutput)
}

// The time at which this CertificateAuthority will be permanently purged, if it is in the DELETED state.
func (o CertificateAuthorityOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.ExpireTime }).(pulumi.StringOutput)
}

// Immutable. The name of a Cloud Storage bucket where this CertificateAuthority will publish content, such as the CA certificate and CRLs. This must be a bucket name, without any prefixes (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named `my-bucket`, you would simply specify `my-bucket`. If not specified, a managed bucket will be created.
func (o CertificateAuthorityOutput) GcsBucket() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.GcsBucket }).(pulumi.StringOutput)
}

// Immutable. Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA certificate. Otherwise, it is used to sign a CSR.
func (o CertificateAuthorityOutput) KeySpec() KeyVersionSpecResponseOutput {
	return o.ApplyT(func(v *CertificateAuthority) KeyVersionSpecResponseOutput { return v.KeySpec }).(KeyVersionSpecResponseOutput)
}

// Optional. Labels with user-defined metadata.
func (o CertificateAuthorityOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Immutable. The desired lifetime of the CA certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate.
func (o CertificateAuthorityOutput) Lifetime() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.Lifetime }).(pulumi.StringOutput)
}

func (o CertificateAuthorityOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name for this CertificateAuthority in the format `projects/*/locations/*/caPools/*/certificateAuthorities/*`.
func (o CertificateAuthorityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// This CertificateAuthority's certificate chain, including the current CertificateAuthority's certificate. Ordered such that the root issuer is the final element (consistent with RFC 5246). For a self-signed CA, this will only list the current CertificateAuthority's certificate.
func (o CertificateAuthorityOutput) PemCaCertificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringArrayOutput { return v.PemCaCertificates }).(pulumi.StringArrayOutput)
}

func (o CertificateAuthorityOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. An ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
func (o CertificateAuthorityOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// The State for this CertificateAuthority.
func (o CertificateAuthorityOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Optional. If this is a subordinate CertificateAuthority, this field will be set with the subordinate configuration, which describes its issuers. This may be updated, but this CertificateAuthority must continue to validate.
func (o CertificateAuthorityOutput) SubordinateConfig() SubordinateConfigResponseOutput {
	return o.ApplyT(func(v *CertificateAuthority) SubordinateConfigResponseOutput { return v.SubordinateConfig }).(SubordinateConfigResponseOutput)
}

// The CaPool.Tier of the CaPool that includes this CertificateAuthority.
func (o CertificateAuthorityOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.Tier }).(pulumi.StringOutput)
}

// Immutable. The Type of this CertificateAuthority.
func (o CertificateAuthorityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The time at which this CertificateAuthority was last updated.
func (o CertificateAuthorityOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateAuthorityInput)(nil)).Elem(), &CertificateAuthority{})
	pulumi.RegisterOutputType(CertificateAuthorityOutput{})
}
