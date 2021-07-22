// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new ApiConfig in a given project and location.
// Auto-naming is currently not supported for this resource.
type Config struct {
	pulumi.CustomResourceState

	// Created time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Immutable. The Google Cloud IAM Service Account that Gateways serving this config should use to authenticate to other services. This may either be the Service Account's email (`{ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com`) or its full resource name (`projects/{PROJECT}/accounts/{UNIQUE_ID}`). This is most often used when the service is a GCP resource such as a Cloud Run Service or an IAP-secured service.
	GatewayServiceAccount pulumi.StringOutput `pulumi:"gatewayServiceAccount"`
	// Optional. gRPC service definition files. If specified, openapi_documents must not be included.
	GrpcServices ApigatewayApiConfigGrpcServiceDefinitionResponseArrayOutput `pulumi:"grpcServices"`
	// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
	ManagedServiceConfigs ApigatewayApiConfigFileResponseArrayOutput `pulumi:"managedServiceConfigs"`
	// Resource name of the API Config. Format: projects/{project}/locations/global/apis/{api}/configs/{api_config}
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. OpenAPI specification documents. If specified, grpc_services and managed_service_configs must not be included.
	OpenapiDocuments ApigatewayApiConfigOpenApiDocumentResponseArrayOutput `pulumi:"openapiDocuments"`
	// The ID of the associated Service Config ( https://cloud.google.com/service-infrastructure/docs/glossary#config).
	ServiceConfigId pulumi.StringOutput `pulumi:"serviceConfigId"`
	// State of the API Config.
	State pulumi.StringOutput `pulumi:"state"`
	// Updated time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewConfig registers a new resource with the given unique name, arguments, and options.
func NewConfig(ctx *pulumi.Context,
	name string, args *ConfigArgs, opts ...pulumi.ResourceOption) (*Config, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ApiConfigId'")
	}
	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	var resource Config
	err := ctx.RegisterResource("google-native:apigateway/v1:Config", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfig gets an existing Config resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigState, opts ...pulumi.ResourceOption) (*Config, error) {
	var resource Config
	err := ctx.ReadResource("google-native:apigateway/v1:Config", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Config resources.
type configState struct {
}

type ConfigState struct {
}

func (ConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*configState)(nil)).Elem()
}

type configArgs struct {
	ApiConfigId string `pulumi:"apiConfigId"`
	ApiId       string `pulumi:"apiId"`
	// Optional. Display name.
	DisplayName *string `pulumi:"displayName"`
	// Immutable. The Google Cloud IAM Service Account that Gateways serving this config should use to authenticate to other services. This may either be the Service Account's email (`{ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com`) or its full resource name (`projects/{PROJECT}/accounts/{UNIQUE_ID}`). This is most often used when the service is a GCP resource such as a Cloud Run Service or an IAP-secured service.
	GatewayServiceAccount *string `pulumi:"gatewayServiceAccount"`
	// Optional. gRPC service definition files. If specified, openapi_documents must not be included.
	GrpcServices []ApigatewayApiConfigGrpcServiceDefinition `pulumi:"grpcServices"`
	// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
	ManagedServiceConfigs []ApigatewayApiConfigFile `pulumi:"managedServiceConfigs"`
	// Optional. OpenAPI specification documents. If specified, grpc_services and managed_service_configs must not be included.
	OpenapiDocuments []ApigatewayApiConfigOpenApiDocument `pulumi:"openapiDocuments"`
	Project          *string                              `pulumi:"project"`
}

// The set of arguments for constructing a Config resource.
type ConfigArgs struct {
	ApiConfigId pulumi.StringInput
	ApiId       pulumi.StringInput
	// Optional. Display name.
	DisplayName pulumi.StringPtrInput
	// Immutable. The Google Cloud IAM Service Account that Gateways serving this config should use to authenticate to other services. This may either be the Service Account's email (`{ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com`) or its full resource name (`projects/{PROJECT}/accounts/{UNIQUE_ID}`). This is most often used when the service is a GCP resource such as a Cloud Run Service or an IAP-secured service.
	GatewayServiceAccount pulumi.StringPtrInput
	// Optional. gRPC service definition files. If specified, openapi_documents must not be included.
	GrpcServices ApigatewayApiConfigGrpcServiceDefinitionArrayInput
	// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
	ManagedServiceConfigs ApigatewayApiConfigFileArrayInput
	// Optional. OpenAPI specification documents. If specified, grpc_services and managed_service_configs must not be included.
	OpenapiDocuments ApigatewayApiConfigOpenApiDocumentArrayInput
	Project          pulumi.StringPtrInput
}

func (ConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configArgs)(nil)).Elem()
}

type ConfigInput interface {
	pulumi.Input

	ToConfigOutput() ConfigOutput
	ToConfigOutputWithContext(ctx context.Context) ConfigOutput
}

func (*Config) ElementType() reflect.Type {
	return reflect.TypeOf((*Config)(nil))
}

func (i *Config) ToConfigOutput() ConfigOutput {
	return i.ToConfigOutputWithContext(context.Background())
}

func (i *Config) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigOutput)
}

type ConfigOutput struct {
	*pulumi.OutputState
}

func (ConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Config)(nil))
}

func (o ConfigOutput) ToConfigOutput() ConfigOutput {
	return o
}

func (o ConfigOutput) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConfigOutput{})
}
