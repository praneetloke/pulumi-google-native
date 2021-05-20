// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an alias from a key/certificate pair. The structure of the request is controlled by the `format` query parameter: - `keycertfile` - Separate PEM-encoded key and certificate files are uploaded. Set `Content-Type: multipart/form-data` and include the `keyFile`, `certFile`, and `password` (if keys are encrypted) fields in the request body. If uploading to a truststore, omit `keyFile`. - `pkcs12` - A PKCS12 file is uploaded. Set `Content-Type: multipart/form-data`, provide the file in the `file` field, and include the `password` field if the file is encrypted in the request body. - `selfsignedcert` - A new private key and certificate are generated. Set `Content-Type: application/json` and include CertificateGenerationSpec in the request body.
type OrganizationEnvironmentKeystoreAlias struct {
	pulumi.CustomResourceState

	// Resource ID for this alias. Values must match the regular expression `[^/]{1,255}`.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// Chain of certificates under this alias.
	CertsInfo GoogleCloudApigeeV1CertificateResponseOutput `pulumi:"certsInfo"`
	// Type of alias.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOrganizationEnvironmentKeystoreAlias registers a new resource with the given unique name, arguments, and options.
func NewOrganizationEnvironmentKeystoreAlias(ctx *pulumi.Context,
	name string, args *OrganizationEnvironmentKeystoreAliasArgs, opts ...pulumi.ResourceOption) (*OrganizationEnvironmentKeystoreAlias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AliasId == nil {
		return nil, errors.New("invalid value for required argument 'AliasId'")
	}
	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.Format == nil {
		return nil, errors.New("invalid value for required argument 'Format'")
	}
	if args.KeystoreId == nil {
		return nil, errors.New("invalid value for required argument 'KeystoreId'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource OrganizationEnvironmentKeystoreAlias
	err := ctx.RegisterResource("google-native:apigee/v1:OrganizationEnvironmentKeystoreAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationEnvironmentKeystoreAlias gets an existing OrganizationEnvironmentKeystoreAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationEnvironmentKeystoreAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationEnvironmentKeystoreAliasState, opts ...pulumi.ResourceOption) (*OrganizationEnvironmentKeystoreAlias, error) {
	var resource OrganizationEnvironmentKeystoreAlias
	err := ctx.ReadResource("google-native:apigee/v1:OrganizationEnvironmentKeystoreAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationEnvironmentKeystoreAlias resources.
type organizationEnvironmentKeystoreAliasState struct {
	// Resource ID for this alias. Values must match the regular expression `[^/]{1,255}`.
	Alias *string `pulumi:"alias"`
	// Chain of certificates under this alias.
	CertsInfo *GoogleCloudApigeeV1CertificateResponse `pulumi:"certsInfo"`
	// Type of alias.
	Type *string `pulumi:"type"`
}

type OrganizationEnvironmentKeystoreAliasState struct {
	// Resource ID for this alias. Values must match the regular expression `[^/]{1,255}`.
	Alias pulumi.StringPtrInput
	// Chain of certificates under this alias.
	CertsInfo GoogleCloudApigeeV1CertificateResponsePtrInput
	// Type of alias.
	Type pulumi.StringPtrInput
}

func (OrganizationEnvironmentKeystoreAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationEnvironmentKeystoreAliasState)(nil)).Elem()
}

type organizationEnvironmentKeystoreAliasArgs struct {
	Alias   *string `pulumi:"alias"`
	AliasId string  `pulumi:"aliasId"`
	// The HTTP Content-Type header value specifying the content type of the body.
	ContentType *string `pulumi:"contentType"`
	// The HTTP request/response body as raw binary.
	Data          *string `pulumi:"data"`
	EnvironmentId string  `pulumi:"environmentId"`
	// Application specific response metadata. Must be set in the first response for streaming APIs.
	Extensions              []map[string]string `pulumi:"extensions"`
	Format                  string              `pulumi:"format"`
	IgnoreExpiryValidation  *string             `pulumi:"ignoreExpiryValidation"`
	IgnoreNewlineValidation *string             `pulumi:"ignoreNewlineValidation"`
	KeystoreId              string              `pulumi:"keystoreId"`
	OrganizationId          string              `pulumi:"organizationId"`
	Password                *string             `pulumi:"password"`
}

// The set of arguments for constructing a OrganizationEnvironmentKeystoreAlias resource.
type OrganizationEnvironmentKeystoreAliasArgs struct {
	Alias   pulumi.StringPtrInput
	AliasId pulumi.StringInput
	// The HTTP Content-Type header value specifying the content type of the body.
	ContentType pulumi.StringPtrInput
	// The HTTP request/response body as raw binary.
	Data          pulumi.StringPtrInput
	EnvironmentId pulumi.StringInput
	// Application specific response metadata. Must be set in the first response for streaming APIs.
	Extensions              pulumi.StringMapArrayInput
	Format                  pulumi.StringInput
	IgnoreExpiryValidation  pulumi.StringPtrInput
	IgnoreNewlineValidation pulumi.StringPtrInput
	KeystoreId              pulumi.StringInput
	OrganizationId          pulumi.StringInput
	Password                pulumi.StringPtrInput
}

func (OrganizationEnvironmentKeystoreAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationEnvironmentKeystoreAliasArgs)(nil)).Elem()
}

type OrganizationEnvironmentKeystoreAliasInput interface {
	pulumi.Input

	ToOrganizationEnvironmentKeystoreAliasOutput() OrganizationEnvironmentKeystoreAliasOutput
	ToOrganizationEnvironmentKeystoreAliasOutputWithContext(ctx context.Context) OrganizationEnvironmentKeystoreAliasOutput
}

func (*OrganizationEnvironmentKeystoreAlias) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationEnvironmentKeystoreAlias)(nil))
}

func (i *OrganizationEnvironmentKeystoreAlias) ToOrganizationEnvironmentKeystoreAliasOutput() OrganizationEnvironmentKeystoreAliasOutput {
	return i.ToOrganizationEnvironmentKeystoreAliasOutputWithContext(context.Background())
}

func (i *OrganizationEnvironmentKeystoreAlias) ToOrganizationEnvironmentKeystoreAliasOutputWithContext(ctx context.Context) OrganizationEnvironmentKeystoreAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationEnvironmentKeystoreAliasOutput)
}

type OrganizationEnvironmentKeystoreAliasOutput struct {
	*pulumi.OutputState
}

func (OrganizationEnvironmentKeystoreAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationEnvironmentKeystoreAlias)(nil))
}

func (o OrganizationEnvironmentKeystoreAliasOutput) ToOrganizationEnvironmentKeystoreAliasOutput() OrganizationEnvironmentKeystoreAliasOutput {
	return o
}

func (o OrganizationEnvironmentKeystoreAliasOutput) ToOrganizationEnvironmentKeystoreAliasOutputWithContext(ctx context.Context) OrganizationEnvironmentKeystoreAliasOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationEnvironmentKeystoreAliasOutput{})
}
