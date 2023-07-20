// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new certificate. The certificate will be registered to the trawler service and will be encrypted using cloud KMS and stored in Spanner Returns the certificate.
// Auto-naming is currently not supported for this resource.
type Certificate struct {
	pulumi.CustomResourceState

	// Status of the certificate
	CertificateStatus pulumi.StringOutput `pulumi:"certificateStatus"`
	// Immutable. Credential id that will be used to register with trawler INTERNAL_ONLY
	CredentialId pulumi.StringOutput `pulumi:"credentialId"`
	// Description of the certificate
	Description pulumi.StringOutput `pulumi:"description"`
	// Name of the certificate
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	Location    pulumi.StringOutput `pulumi:"location"`
	// Auto generated primary key
	Name      pulumi.StringOutput `pulumi:"name"`
	ProductId pulumi.StringOutput `pulumi:"productId"`
	Project   pulumi.StringOutput `pulumi:"project"`
	// Input only. Raw client certificate which would be registered with trawler
	RawCertificate GoogleCloudIntegrationsV1alphaClientCertificateResponseOutput `pulumi:"rawCertificate"`
	// Immutable. Requestor ID to be used to register certificate with trawler
	RequestorId pulumi.StringOutput `pulumi:"requestorId"`
	// The timestamp after which certificate will expire
	ValidEndTime pulumi.StringOutput `pulumi:"validEndTime"`
	// The timestamp after which certificate will be valid
	ValidStartTime pulumi.StringOutput `pulumi:"validStartTime"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"productId",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Certificate
	err := ctx.RegisterResource("google-native:integrations/v1alpha:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("google-native:integrations/v1alpha:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
}

type CertificateState struct {
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// Status of the certificate
	CertificateStatus *CertificateCertificateStatus `pulumi:"certificateStatus"`
	// Immutable. Credential id that will be used to register with trawler INTERNAL_ONLY
	CredentialId *string `pulumi:"credentialId"`
	// Description of the certificate
	Description *string `pulumi:"description"`
	// Name of the certificate
	DisplayName *string `pulumi:"displayName"`
	Location    *string `pulumi:"location"`
	ProductId   string  `pulumi:"productId"`
	Project     *string `pulumi:"project"`
	// Input only. Raw client certificate which would be registered with trawler
	RawCertificate *GoogleCloudIntegrationsV1alphaClientCertificate `pulumi:"rawCertificate"`
	// Immutable. Requestor ID to be used to register certificate with trawler
	RequestorId *string `pulumi:"requestorId"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// Status of the certificate
	CertificateStatus CertificateCertificateStatusPtrInput
	// Immutable. Credential id that will be used to register with trawler INTERNAL_ONLY
	CredentialId pulumi.StringPtrInput
	// Description of the certificate
	Description pulumi.StringPtrInput
	// Name of the certificate
	DisplayName pulumi.StringPtrInput
	Location    pulumi.StringPtrInput
	ProductId   pulumi.StringInput
	Project     pulumi.StringPtrInput
	// Input only. Raw client certificate which would be registered with trawler
	RawCertificate GoogleCloudIntegrationsV1alphaClientCertificatePtrInput
	// Immutable. Requestor ID to be used to register certificate with trawler
	RequestorId pulumi.StringPtrInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

// Status of the certificate
func (o CertificateOutput) CertificateStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CertificateStatus }).(pulumi.StringOutput)
}

// Immutable. Credential id that will be used to register with trawler INTERNAL_ONLY
func (o CertificateOutput) CredentialId() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CredentialId }).(pulumi.StringOutput)
}

// Description of the certificate
func (o CertificateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Name of the certificate
func (o CertificateOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o CertificateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Auto generated primary key
func (o CertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CertificateOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ProductId }).(pulumi.StringOutput)
}

func (o CertificateOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Input only. Raw client certificate which would be registered with trawler
func (o CertificateOutput) RawCertificate() GoogleCloudIntegrationsV1alphaClientCertificateResponseOutput {
	return o.ApplyT(func(v *Certificate) GoogleCloudIntegrationsV1alphaClientCertificateResponseOutput {
		return v.RawCertificate
	}).(GoogleCloudIntegrationsV1alphaClientCertificateResponseOutput)
}

// Immutable. Requestor ID to be used to register certificate with trawler
func (o CertificateOutput) RequestorId() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.RequestorId }).(pulumi.StringOutput)
}

// The timestamp after which certificate will expire
func (o CertificateOutput) ValidEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ValidEndTime }).(pulumi.StringOutput)
}

// The timestamp after which certificate will be valid
func (o CertificateOutput) ValidStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ValidStartTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterOutputType(CertificateOutput{})
}
