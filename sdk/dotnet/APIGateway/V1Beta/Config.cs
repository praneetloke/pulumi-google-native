// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.APIGateway.V1Beta
{
    /// <summary>
    /// Creates a new ApiConfig in a given project and location.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:apigateway/v1beta:Config")]
    public partial class Config : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Required. Identifier to assign to the API Config. Must be unique within scope of the parent resource.
        /// </summary>
        [Output("apiConfigId")]
        public Output<string> ApiConfigId { get; private set; } = null!;

        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// Created time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Immutable. Gateway specific configuration.
        /// </summary>
        [Output("gatewayConfig")]
        public Output<Outputs.ApigatewayGatewayConfigResponse> GatewayConfig { get; private set; } = null!;

        /// <summary>
        /// Immutable. The Google Cloud IAM Service Account that Gateways serving this config should use to authenticate to other services. This may either be the Service Account's email (`{ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com`) or its full resource name (`projects/{PROJECT}/accounts/{UNIQUE_ID}`). This is most often used when the service is a GCP resource such as a Cloud Run Service or an IAP-secured service.
        /// </summary>
        [Output("gatewayServiceAccount")]
        public Output<string> GatewayServiceAccount { get; private set; } = null!;

        /// <summary>
        /// Optional. gRPC service definition files. If specified, openapi_documents must not be included.
        /// </summary>
        [Output("grpcServices")]
        public Output<ImmutableArray<Outputs.ApigatewayApiConfigGrpcServiceDefinitionResponse>> GrpcServices { get; private set; } = null!;

        /// <summary>
        /// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
        /// </summary>
        [Output("managedServiceConfigs")]
        public Output<ImmutableArray<Outputs.ApigatewayApiConfigFileResponse>> ManagedServiceConfigs { get; private set; } = null!;

        /// <summary>
        /// Resource name of the API Config. Format: projects/{project}/locations/global/apis/{api}/configs/{api_config}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. OpenAPI specification documents. If specified, grpc_services and managed_service_configs must not be included.
        /// </summary>
        [Output("openapiDocuments")]
        public Output<ImmutableArray<Outputs.ApigatewayApiConfigOpenApiDocumentResponse>> OpenapiDocuments { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated Service Config ( https://cloud.google.com/service-infrastructure/docs/glossary#config).
        /// </summary>
        [Output("serviceConfigId")]
        public Output<string> ServiceConfigId { get; private set; } = null!;

        /// <summary>
        /// State of the API Config.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Updated time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Config resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Config(string name, ConfigArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigateway/v1beta:Config", name, args ?? new ConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Config(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigateway/v1beta:Config", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "apiConfigId",
                    "apiId",
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Config resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Config Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Config(name, id, options);
        }
    }

    public sealed class ConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. Identifier to assign to the API Config. Must be unique within scope of the parent resource.
        /// </summary>
        [Input("apiConfigId", required: true)]
        public Input<string> ApiConfigId { get; set; } = null!;

        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// Optional. Display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Immutable. Gateway specific configuration.
        /// </summary>
        [Input("gatewayConfig")]
        public Input<Inputs.ApigatewayGatewayConfigArgs>? GatewayConfig { get; set; }

        /// <summary>
        /// Immutable. The Google Cloud IAM Service Account that Gateways serving this config should use to authenticate to other services. This may either be the Service Account's email (`{ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com`) or its full resource name (`projects/{PROJECT}/accounts/{UNIQUE_ID}`). This is most often used when the service is a GCP resource such as a Cloud Run Service or an IAP-secured service.
        /// </summary>
        [Input("gatewayServiceAccount")]
        public Input<string>? GatewayServiceAccount { get; set; }

        [Input("grpcServices")]
        private InputList<Inputs.ApigatewayApiConfigGrpcServiceDefinitionArgs>? _grpcServices;

        /// <summary>
        /// Optional. gRPC service definition files. If specified, openapi_documents must not be included.
        /// </summary>
        public InputList<Inputs.ApigatewayApiConfigGrpcServiceDefinitionArgs> GrpcServices
        {
            get => _grpcServices ?? (_grpcServices = new InputList<Inputs.ApigatewayApiConfigGrpcServiceDefinitionArgs>());
            set => _grpcServices = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("managedServiceConfigs")]
        private InputList<Inputs.ApigatewayApiConfigFileArgs>? _managedServiceConfigs;

        /// <summary>
        /// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
        /// </summary>
        public InputList<Inputs.ApigatewayApiConfigFileArgs> ManagedServiceConfigs
        {
            get => _managedServiceConfigs ?? (_managedServiceConfigs = new InputList<Inputs.ApigatewayApiConfigFileArgs>());
            set => _managedServiceConfigs = value;
        }

        [Input("openapiDocuments")]
        private InputList<Inputs.ApigatewayApiConfigOpenApiDocumentArgs>? _openapiDocuments;

        /// <summary>
        /// Optional. OpenAPI specification documents. If specified, grpc_services and managed_service_configs must not be included.
        /// </summary>
        public InputList<Inputs.ApigatewayApiConfigOpenApiDocumentArgs> OpenapiDocuments
        {
            get => _openapiDocuments ?? (_openapiDocuments = new InputList<Inputs.ApigatewayApiConfigOpenApiDocumentArgs>());
            set => _openapiDocuments = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public ConfigArgs()
        {
        }
        public static new ConfigArgs Empty => new ConfigArgs();
    }
}
