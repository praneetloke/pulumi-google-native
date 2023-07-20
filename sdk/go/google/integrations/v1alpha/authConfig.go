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

// Creates an auth config record. Fetch corresponding credentials for specific auth types, e.g. access token for OAuth 2.0, JWT token for JWT. Encrypt the auth config with Cloud KMS and store the encrypted credentials in Spanner. Returns the encrypted auth config.
// Auto-naming is currently not supported for this resource.
type AuthConfig struct {
	pulumi.CustomResourceState

	// Certificate id for client certificate
	CertificateId pulumi.StringOutput `pulumi:"certificateId"`
	// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines. For example, -----BEGIN CERTIFICATE----- MIICTTCCAbagAwIBAgIJAPT0tSKNxan/MA0GCSqGSIb3DQEBCwUAMCoxFzAVBgNV BAoTDkdvb2dsZSBURVNUSU5HMQ8wDQYDVQQDEwZ0ZXN0Q0EwHhcNMTUwMTAxMDAw MDAwWhcNMjUwMTAxMDAwMDAwWjAuMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElORzET MBEGA1UEAwwKam9lQGJhbmFuYTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA vDYFgMgxi5W488d9J7UpCInl0NXmZQpJDEHE4hvkaRlH7pnC71H0DLt0/3zATRP1 JzY2+eqBmbGl4/sgZKYv8UrLnNyQNUTsNx1iZAfPUflf5FwgVsai8BM0pUciq1NB xD429VFcrGZNucvFLh72RuRFIKH8WUpiK/iZNFkWhZ0CAwEAAaN3MHUwDgYDVR0P AQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMB Af8EAjAAMBkGA1UdDgQSBBCVgnFBCWgL/iwCqnGrhTPQMBsGA1UdIwQUMBKAEKey Um2o4k2WiEVA0ldQvNYwDQYJKoZIhvcNAQELBQADgYEAYK986R4E3L1v+Q6esBtW JrUwA9UmJRSQr0N5w3o9XzarU37/bkjOP0Fw0k/A6Vv1n3vlciYfBFaBIam1qRHr 5dMsYf4CZS6w50r7hyzqyrwDoyNxkLnd2PdcHT/sym1QmflsjEs7pejtnohO6N2H wQW6M0H7Zt8claGRla4fKkg= -----END CERTIFICATE-----
	ClientCertificateEncryptedPrivateKey pulumi.StringPtrOutput `pulumi:"clientCertificateEncryptedPrivateKey"`
	// 'passphrase' should be left unset if private key is not encrypted. Note that 'passphrase' is not the password for web server, but an extra layer of security to protected private key.
	ClientCertificatePassphrase pulumi.StringPtrOutput `pulumi:"clientCertificatePassphrase"`
	// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines. For example, -----BEGIN CERTIFICATE----- MIICTTCCAbagAwIBAgIJAPT0tSKNxan/MA0GCSqGSIb3DQEBCwUAMCoxFzAVBgNV BAoTDkdvb2dsZSBURVNUSU5HMQ8wDQYDVQQDEwZ0ZXN0Q0EwHhcNMTUwMTAxMDAw MDAwWhcNMjUwMTAxMDAwMDAwWjAuMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElORzET MBEGA1UEAwwKam9lQGJhbmFuYTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA vDYFgMgxi5W488d9J7UpCInl0NXmZQpJDEHE4hvkaRlH7pnC71H0DLt0/3zATRP1 JzY2+eqBmbGl4/sgZKYv8UrLnNyQNUTsNx1iZAfPUflf5FwgVsai8BM0pUciq1NB xD429VFcrGZNucvFLh72RuRFIKH8WUpiK/iZNFkWhZ0CAwEAAaN3MHUwDgYDVR0P AQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMB Af8EAjAAMBkGA1UdDgQSBBCVgnFBCWgL/iwCqnGrhTPQMBsGA1UdIwQUMBKAEKey Um2o4k2WiEVA0ldQvNYwDQYJKoZIhvcNAQELBQADgYEAYK986R4E3L1v+Q6esBtW JrUwA9UmJRSQr0N5w3o9XzarU37/bkjOP0Fw0k/A6Vv1n3vlciYfBFaBIam1qRHr 5dMsYf4CZS6w50r7hyzqyrwDoyNxkLnd2PdcHT/sym1QmflsjEs7pejtnohO6N2H wQW6M0H7Zt8claGRla4fKkg= -----END CERTIFICATE-----
	ClientCertificateSslCertificate pulumi.StringPtrOutput `pulumi:"clientCertificateSslCertificate"`
	// The timestamp when the auth config is created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The creator's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
	CreatorEmail pulumi.StringOutput `pulumi:"creatorEmail"`
	// Credential type of the encrypted credential.
	CredentialType pulumi.StringOutput `pulumi:"credentialType"`
	// Raw auth credentials.
	DecryptedCredential GoogleCloudIntegrationsV1alphaCredentialResponseOutput `pulumi:"decryptedCredential"`
	// A description of the auth config.
	Description pulumi.StringOutput `pulumi:"description"`
	// The name of the auth config.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Auth credential encrypted by Cloud KMS. Can be decrypted as Credential with proper KMS key.
	EncryptedCredential pulumi.StringOutput `pulumi:"encryptedCredential"`
	// User can define the time to receive notification after which the auth config becomes invalid. Support up to 30 days. Support granularity in hours.
	ExpiryNotificationDuration pulumi.StringArrayOutput `pulumi:"expiryNotificationDuration"`
	// The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
	LastModifierEmail pulumi.StringOutput `pulumi:"lastModifierEmail"`
	Location          pulumi.StringOutput `pulumi:"location"`
	// Resource name of the SFDC instance projects/{project}/locations/{location}/authConfigs/{authConfig}.
	Name pulumi.StringOutput `pulumi:"name"`
	// User provided expiry time to override. For the example of Salesforce, username/password credentials can be valid for 6 months depending on the instance settings.
	OverrideValidTime pulumi.StringOutput `pulumi:"overrideValidTime"`
	ProductId         pulumi.StringOutput `pulumi:"productId"`
	Project           pulumi.StringOutput `pulumi:"project"`
	// The reason / details of the current status.
	Reason pulumi.StringOutput `pulumi:"reason"`
	// The status of the auth config.
	State pulumi.StringOutput `pulumi:"state"`
	// The timestamp when the auth config is modified.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The time until the auth config is valid. Empty or max value is considered the auth config won't expire.
	ValidTime pulumi.StringOutput `pulumi:"validTime"`
	// The visibility of the auth config.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewAuthConfig registers a new resource with the given unique name, arguments, and options.
func NewAuthConfig(ctx *pulumi.Context,
	name string, args *AuthConfigArgs, opts ...pulumi.ResourceOption) (*AuthConfig, error) {
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
	var resource AuthConfig
	err := ctx.RegisterResource("google-native:integrations/v1alpha:AuthConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthConfig gets an existing AuthConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthConfigState, opts ...pulumi.ResourceOption) (*AuthConfig, error) {
	var resource AuthConfig
	err := ctx.ReadResource("google-native:integrations/v1alpha:AuthConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthConfig resources.
type authConfigState struct {
}

type AuthConfigState struct {
}

func (AuthConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*authConfigState)(nil)).Elem()
}

type authConfigArgs struct {
	// Certificate id for client certificate
	CertificateId *string `pulumi:"certificateId"`
	// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines. For example, -----BEGIN CERTIFICATE----- MIICTTCCAbagAwIBAgIJAPT0tSKNxan/MA0GCSqGSIb3DQEBCwUAMCoxFzAVBgNV BAoTDkdvb2dsZSBURVNUSU5HMQ8wDQYDVQQDEwZ0ZXN0Q0EwHhcNMTUwMTAxMDAw MDAwWhcNMjUwMTAxMDAwMDAwWjAuMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElORzET MBEGA1UEAwwKam9lQGJhbmFuYTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA vDYFgMgxi5W488d9J7UpCInl0NXmZQpJDEHE4hvkaRlH7pnC71H0DLt0/3zATRP1 JzY2+eqBmbGl4/sgZKYv8UrLnNyQNUTsNx1iZAfPUflf5FwgVsai8BM0pUciq1NB xD429VFcrGZNucvFLh72RuRFIKH8WUpiK/iZNFkWhZ0CAwEAAaN3MHUwDgYDVR0P AQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMB Af8EAjAAMBkGA1UdDgQSBBCVgnFBCWgL/iwCqnGrhTPQMBsGA1UdIwQUMBKAEKey Um2o4k2WiEVA0ldQvNYwDQYJKoZIhvcNAQELBQADgYEAYK986R4E3L1v+Q6esBtW JrUwA9UmJRSQr0N5w3o9XzarU37/bkjOP0Fw0k/A6Vv1n3vlciYfBFaBIam1qRHr 5dMsYf4CZS6w50r7hyzqyrwDoyNxkLnd2PdcHT/sym1QmflsjEs7pejtnohO6N2H wQW6M0H7Zt8claGRla4fKkg= -----END CERTIFICATE-----
	ClientCertificateEncryptedPrivateKey *string `pulumi:"clientCertificateEncryptedPrivateKey"`
	// 'passphrase' should be left unset if private key is not encrypted. Note that 'passphrase' is not the password for web server, but an extra layer of security to protected private key.
	ClientCertificatePassphrase *string `pulumi:"clientCertificatePassphrase"`
	// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines. For example, -----BEGIN CERTIFICATE----- MIICTTCCAbagAwIBAgIJAPT0tSKNxan/MA0GCSqGSIb3DQEBCwUAMCoxFzAVBgNV BAoTDkdvb2dsZSBURVNUSU5HMQ8wDQYDVQQDEwZ0ZXN0Q0EwHhcNMTUwMTAxMDAw MDAwWhcNMjUwMTAxMDAwMDAwWjAuMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElORzET MBEGA1UEAwwKam9lQGJhbmFuYTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA vDYFgMgxi5W488d9J7UpCInl0NXmZQpJDEHE4hvkaRlH7pnC71H0DLt0/3zATRP1 JzY2+eqBmbGl4/sgZKYv8UrLnNyQNUTsNx1iZAfPUflf5FwgVsai8BM0pUciq1NB xD429VFcrGZNucvFLh72RuRFIKH8WUpiK/iZNFkWhZ0CAwEAAaN3MHUwDgYDVR0P AQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMB Af8EAjAAMBkGA1UdDgQSBBCVgnFBCWgL/iwCqnGrhTPQMBsGA1UdIwQUMBKAEKey Um2o4k2WiEVA0ldQvNYwDQYJKoZIhvcNAQELBQADgYEAYK986R4E3L1v+Q6esBtW JrUwA9UmJRSQr0N5w3o9XzarU37/bkjOP0Fw0k/A6Vv1n3vlciYfBFaBIam1qRHr 5dMsYf4CZS6w50r7hyzqyrwDoyNxkLnd2PdcHT/sym1QmflsjEs7pejtnohO6N2H wQW6M0H7Zt8claGRla4fKkg= -----END CERTIFICATE-----
	ClientCertificateSslCertificate *string `pulumi:"clientCertificateSslCertificate"`
	// The creator's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
	CreatorEmail *string `pulumi:"creatorEmail"`
	// Credential type of the encrypted credential.
	CredentialType *AuthConfigCredentialType `pulumi:"credentialType"`
	// Raw auth credentials.
	DecryptedCredential *GoogleCloudIntegrationsV1alphaCredential `pulumi:"decryptedCredential"`
	// A description of the auth config.
	Description *string `pulumi:"description"`
	// The name of the auth config.
	DisplayName *string `pulumi:"displayName"`
	// Auth credential encrypted by Cloud KMS. Can be decrypted as Credential with proper KMS key.
	EncryptedCredential *string `pulumi:"encryptedCredential"`
	// User can define the time to receive notification after which the auth config becomes invalid. Support up to 30 days. Support granularity in hours.
	ExpiryNotificationDuration []string `pulumi:"expiryNotificationDuration"`
	// The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
	LastModifierEmail *string `pulumi:"lastModifierEmail"`
	Location          *string `pulumi:"location"`
	// Resource name of the SFDC instance projects/{project}/locations/{location}/authConfigs/{authConfig}.
	Name *string `pulumi:"name"`
	// User provided expiry time to override. For the example of Salesforce, username/password credentials can be valid for 6 months depending on the instance settings.
	OverrideValidTime *string `pulumi:"overrideValidTime"`
	ProductId         string  `pulumi:"productId"`
	Project           *string `pulumi:"project"`
	// The reason / details of the current status.
	Reason *string `pulumi:"reason"`
	// The status of the auth config.
	State *AuthConfigStateEnum `pulumi:"state"`
	// The time until the auth config is valid. Empty or max value is considered the auth config won't expire.
	ValidTime *string `pulumi:"validTime"`
	// The visibility of the auth config.
	Visibility *AuthConfigVisibility `pulumi:"visibility"`
}

// The set of arguments for constructing a AuthConfig resource.
type AuthConfigArgs struct {
	// Certificate id for client certificate
	CertificateId pulumi.StringPtrInput
	// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines. For example, -----BEGIN CERTIFICATE----- MIICTTCCAbagAwIBAgIJAPT0tSKNxan/MA0GCSqGSIb3DQEBCwUAMCoxFzAVBgNV BAoTDkdvb2dsZSBURVNUSU5HMQ8wDQYDVQQDEwZ0ZXN0Q0EwHhcNMTUwMTAxMDAw MDAwWhcNMjUwMTAxMDAwMDAwWjAuMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElORzET MBEGA1UEAwwKam9lQGJhbmFuYTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA vDYFgMgxi5W488d9J7UpCInl0NXmZQpJDEHE4hvkaRlH7pnC71H0DLt0/3zATRP1 JzY2+eqBmbGl4/sgZKYv8UrLnNyQNUTsNx1iZAfPUflf5FwgVsai8BM0pUciq1NB xD429VFcrGZNucvFLh72RuRFIKH8WUpiK/iZNFkWhZ0CAwEAAaN3MHUwDgYDVR0P AQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMB Af8EAjAAMBkGA1UdDgQSBBCVgnFBCWgL/iwCqnGrhTPQMBsGA1UdIwQUMBKAEKey Um2o4k2WiEVA0ldQvNYwDQYJKoZIhvcNAQELBQADgYEAYK986R4E3L1v+Q6esBtW JrUwA9UmJRSQr0N5w3o9XzarU37/bkjOP0Fw0k/A6Vv1n3vlciYfBFaBIam1qRHr 5dMsYf4CZS6w50r7hyzqyrwDoyNxkLnd2PdcHT/sym1QmflsjEs7pejtnohO6N2H wQW6M0H7Zt8claGRla4fKkg= -----END CERTIFICATE-----
	ClientCertificateEncryptedPrivateKey pulumi.StringPtrInput
	// 'passphrase' should be left unset if private key is not encrypted. Note that 'passphrase' is not the password for web server, but an extra layer of security to protected private key.
	ClientCertificatePassphrase pulumi.StringPtrInput
	// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines. For example, -----BEGIN CERTIFICATE----- MIICTTCCAbagAwIBAgIJAPT0tSKNxan/MA0GCSqGSIb3DQEBCwUAMCoxFzAVBgNV BAoTDkdvb2dsZSBURVNUSU5HMQ8wDQYDVQQDEwZ0ZXN0Q0EwHhcNMTUwMTAxMDAw MDAwWhcNMjUwMTAxMDAwMDAwWjAuMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElORzET MBEGA1UEAwwKam9lQGJhbmFuYTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA vDYFgMgxi5W488d9J7UpCInl0NXmZQpJDEHE4hvkaRlH7pnC71H0DLt0/3zATRP1 JzY2+eqBmbGl4/sgZKYv8UrLnNyQNUTsNx1iZAfPUflf5FwgVsai8BM0pUciq1NB xD429VFcrGZNucvFLh72RuRFIKH8WUpiK/iZNFkWhZ0CAwEAAaN3MHUwDgYDVR0P AQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMB Af8EAjAAMBkGA1UdDgQSBBCVgnFBCWgL/iwCqnGrhTPQMBsGA1UdIwQUMBKAEKey Um2o4k2WiEVA0ldQvNYwDQYJKoZIhvcNAQELBQADgYEAYK986R4E3L1v+Q6esBtW JrUwA9UmJRSQr0N5w3o9XzarU37/bkjOP0Fw0k/A6Vv1n3vlciYfBFaBIam1qRHr 5dMsYf4CZS6w50r7hyzqyrwDoyNxkLnd2PdcHT/sym1QmflsjEs7pejtnohO6N2H wQW6M0H7Zt8claGRla4fKkg= -----END CERTIFICATE-----
	ClientCertificateSslCertificate pulumi.StringPtrInput
	// The creator's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
	CreatorEmail pulumi.StringPtrInput
	// Credential type of the encrypted credential.
	CredentialType AuthConfigCredentialTypePtrInput
	// Raw auth credentials.
	DecryptedCredential GoogleCloudIntegrationsV1alphaCredentialPtrInput
	// A description of the auth config.
	Description pulumi.StringPtrInput
	// The name of the auth config.
	DisplayName pulumi.StringPtrInput
	// Auth credential encrypted by Cloud KMS. Can be decrypted as Credential with proper KMS key.
	EncryptedCredential pulumi.StringPtrInput
	// User can define the time to receive notification after which the auth config becomes invalid. Support up to 30 days. Support granularity in hours.
	ExpiryNotificationDuration pulumi.StringArrayInput
	// The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
	LastModifierEmail pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	// Resource name of the SFDC instance projects/{project}/locations/{location}/authConfigs/{authConfig}.
	Name pulumi.StringPtrInput
	// User provided expiry time to override. For the example of Salesforce, username/password credentials can be valid for 6 months depending on the instance settings.
	OverrideValidTime pulumi.StringPtrInput
	ProductId         pulumi.StringInput
	Project           pulumi.StringPtrInput
	// The reason / details of the current status.
	Reason pulumi.StringPtrInput
	// The status of the auth config.
	State AuthConfigStateEnumPtrInput
	// The time until the auth config is valid. Empty or max value is considered the auth config won't expire.
	ValidTime pulumi.StringPtrInput
	// The visibility of the auth config.
	Visibility AuthConfigVisibilityPtrInput
}

func (AuthConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authConfigArgs)(nil)).Elem()
}

type AuthConfigInput interface {
	pulumi.Input

	ToAuthConfigOutput() AuthConfigOutput
	ToAuthConfigOutputWithContext(ctx context.Context) AuthConfigOutput
}

func (*AuthConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthConfig)(nil)).Elem()
}

func (i *AuthConfig) ToAuthConfigOutput() AuthConfigOutput {
	return i.ToAuthConfigOutputWithContext(context.Background())
}

func (i *AuthConfig) ToAuthConfigOutputWithContext(ctx context.Context) AuthConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthConfigOutput)
}

type AuthConfigOutput struct{ *pulumi.OutputState }

func (AuthConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthConfig)(nil)).Elem()
}

func (o AuthConfigOutput) ToAuthConfigOutput() AuthConfigOutput {
	return o
}

func (o AuthConfigOutput) ToAuthConfigOutputWithContext(ctx context.Context) AuthConfigOutput {
	return o
}

// Certificate id for client certificate
func (o AuthConfigOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.CertificateId }).(pulumi.StringOutput)
}

// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines. For example, -----BEGIN CERTIFICATE----- MIICTTCCAbagAwIBAgIJAPT0tSKNxan/MA0GCSqGSIb3DQEBCwUAMCoxFzAVBgNV BAoTDkdvb2dsZSBURVNUSU5HMQ8wDQYDVQQDEwZ0ZXN0Q0EwHhcNMTUwMTAxMDAw MDAwWhcNMjUwMTAxMDAwMDAwWjAuMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElORzET MBEGA1UEAwwKam9lQGJhbmFuYTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA vDYFgMgxi5W488d9J7UpCInl0NXmZQpJDEHE4hvkaRlH7pnC71H0DLt0/3zATRP1 JzY2+eqBmbGl4/sgZKYv8UrLnNyQNUTsNx1iZAfPUflf5FwgVsai8BM0pUciq1NB xD429VFcrGZNucvFLh72RuRFIKH8WUpiK/iZNFkWhZ0CAwEAAaN3MHUwDgYDVR0P AQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMB Af8EAjAAMBkGA1UdDgQSBBCVgnFBCWgL/iwCqnGrhTPQMBsGA1UdIwQUMBKAEKey Um2o4k2WiEVA0ldQvNYwDQYJKoZIhvcNAQELBQADgYEAYK986R4E3L1v+Q6esBtW JrUwA9UmJRSQr0N5w3o9XzarU37/bkjOP0Fw0k/A6Vv1n3vlciYfBFaBIam1qRHr 5dMsYf4CZS6w50r7hyzqyrwDoyNxkLnd2PdcHT/sym1QmflsjEs7pejtnohO6N2H wQW6M0H7Zt8claGRla4fKkg= -----END CERTIFICATE-----
func (o AuthConfigOutput) ClientCertificateEncryptedPrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringPtrOutput { return v.ClientCertificateEncryptedPrivateKey }).(pulumi.StringPtrOutput)
}

// 'passphrase' should be left unset if private key is not encrypted. Note that 'passphrase' is not the password for web server, but an extra layer of security to protected private key.
func (o AuthConfigOutput) ClientCertificatePassphrase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringPtrOutput { return v.ClientCertificatePassphrase }).(pulumi.StringPtrOutput)
}

// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines. For example, -----BEGIN CERTIFICATE----- MIICTTCCAbagAwIBAgIJAPT0tSKNxan/MA0GCSqGSIb3DQEBCwUAMCoxFzAVBgNV BAoTDkdvb2dsZSBURVNUSU5HMQ8wDQYDVQQDEwZ0ZXN0Q0EwHhcNMTUwMTAxMDAw MDAwWhcNMjUwMTAxMDAwMDAwWjAuMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElORzET MBEGA1UEAwwKam9lQGJhbmFuYTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA vDYFgMgxi5W488d9J7UpCInl0NXmZQpJDEHE4hvkaRlH7pnC71H0DLt0/3zATRP1 JzY2+eqBmbGl4/sgZKYv8UrLnNyQNUTsNx1iZAfPUflf5FwgVsai8BM0pUciq1NB xD429VFcrGZNucvFLh72RuRFIKH8WUpiK/iZNFkWhZ0CAwEAAaN3MHUwDgYDVR0P AQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMB Af8EAjAAMBkGA1UdDgQSBBCVgnFBCWgL/iwCqnGrhTPQMBsGA1UdIwQUMBKAEKey Um2o4k2WiEVA0ldQvNYwDQYJKoZIhvcNAQELBQADgYEAYK986R4E3L1v+Q6esBtW JrUwA9UmJRSQr0N5w3o9XzarU37/bkjOP0Fw0k/A6Vv1n3vlciYfBFaBIam1qRHr 5dMsYf4CZS6w50r7hyzqyrwDoyNxkLnd2PdcHT/sym1QmflsjEs7pejtnohO6N2H wQW6M0H7Zt8claGRla4fKkg= -----END CERTIFICATE-----
func (o AuthConfigOutput) ClientCertificateSslCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringPtrOutput { return v.ClientCertificateSslCertificate }).(pulumi.StringPtrOutput)
}

// The timestamp when the auth config is created.
func (o AuthConfigOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The creator's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
func (o AuthConfigOutput) CreatorEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.CreatorEmail }).(pulumi.StringOutput)
}

// Credential type of the encrypted credential.
func (o AuthConfigOutput) CredentialType() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.CredentialType }).(pulumi.StringOutput)
}

// Raw auth credentials.
func (o AuthConfigOutput) DecryptedCredential() GoogleCloudIntegrationsV1alphaCredentialResponseOutput {
	return o.ApplyT(func(v *AuthConfig) GoogleCloudIntegrationsV1alphaCredentialResponseOutput {
		return v.DecryptedCredential
	}).(GoogleCloudIntegrationsV1alphaCredentialResponseOutput)
}

// A description of the auth config.
func (o AuthConfigOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The name of the auth config.
func (o AuthConfigOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Auth credential encrypted by Cloud KMS. Can be decrypted as Credential with proper KMS key.
func (o AuthConfigOutput) EncryptedCredential() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.EncryptedCredential }).(pulumi.StringOutput)
}

// User can define the time to receive notification after which the auth config becomes invalid. Support up to 30 days. Support granularity in hours.
func (o AuthConfigOutput) ExpiryNotificationDuration() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringArrayOutput { return v.ExpiryNotificationDuration }).(pulumi.StringArrayOutput)
}

// The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
func (o AuthConfigOutput) LastModifierEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.LastModifierEmail }).(pulumi.StringOutput)
}

func (o AuthConfigOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name of the SFDC instance projects/{project}/locations/{location}/authConfigs/{authConfig}.
func (o AuthConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// User provided expiry time to override. For the example of Salesforce, username/password credentials can be valid for 6 months depending on the instance settings.
func (o AuthConfigOutput) OverrideValidTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.OverrideValidTime }).(pulumi.StringOutput)
}

func (o AuthConfigOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.ProductId }).(pulumi.StringOutput)
}

func (o AuthConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The reason / details of the current status.
func (o AuthConfigOutput) Reason() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.Reason }).(pulumi.StringOutput)
}

// The status of the auth config.
func (o AuthConfigOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The timestamp when the auth config is modified.
func (o AuthConfigOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// The time until the auth config is valid. Empty or max value is considered the auth config won't expire.
func (o AuthConfigOutput) ValidTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.ValidTime }).(pulumi.StringOutput)
}

// The visibility of the auth config.
func (o AuthConfigOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthConfig) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthConfigInput)(nil)).Elem(), &AuthConfig{})
	pulumi.RegisterOutputType(AuthConfigOutput{})
}
