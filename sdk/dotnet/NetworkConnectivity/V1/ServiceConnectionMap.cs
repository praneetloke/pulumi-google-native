// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkConnectivity.V1
{
    /// <summary>
    /// Creates a new ServiceConnectionMap in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:networkconnectivity/v1:ServiceConnectionMap")]
    public partial class ServiceConnectionMap : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The PSC configurations on consumer side.
        /// </summary>
        [Output("consumerPscConfigs")]
        public Output<ImmutableArray<Outputs.ConsumerPscConfigResponse>> ConsumerPscConfigs { get; private set; } = null!;

        /// <summary>
        /// PSC connection details on consumer side.
        /// </summary>
        [Output("consumerPscConnections")]
        public Output<ImmutableArray<Outputs.ConsumerPscConnectionResponse>> ConsumerPscConnections { get; private set; } = null!;

        /// <summary>
        /// Time when the ServiceConnectionMap was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// A description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The infrastructure used for connections between consumers/producers.
        /// </summary>
        [Output("infrastructure")]
        public Output<string> Infrastructure { get; private set; } = null!;

        /// <summary>
        /// User-defined labels.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Immutable. The name of a ServiceConnectionMap. Format: projects/{project}/locations/{location}/serviceConnectionMaps/{service_connection_map} See: https://google.aip.dev/122#fields-representing-resource-names
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The PSC configurations on producer side.
        /// </summary>
        [Output("producerPscConfigs")]
        public Output<ImmutableArray<Outputs.ProducerPscConfigResponse>> ProducerPscConfigs { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// The service class identifier this ServiceConnectionMap is for. The user of ServiceConnectionMap create API needs to have networkconnecitivty.serviceclasses.use iam permission for the service class.
        /// </summary>
        [Output("serviceClass")]
        public Output<string> ServiceClass { get; private set; } = null!;

        /// <summary>
        /// The service class uri this ServiceConnectionMap is for.
        /// </summary>
        [Output("serviceClassUri")]
        public Output<string> ServiceClassUri { get; private set; } = null!;

        /// <summary>
        /// Optional. Resource ID (i.e. 'foo' in '[...]/projects/p/locations/l/serviceConnectionMaps/foo') See https://google.aip.dev/122#resource-id-segments Unique per location. If one is not provided, one will be generated.
        /// </summary>
        [Output("serviceConnectionMapId")]
        public Output<string?> ServiceConnectionMapId { get; private set; } = null!;

        /// <summary>
        /// The token provided by the consumer. This token authenticates that the consumer can create a connecton within the specified project and network.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// Time when the ServiceConnectionMap was updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceConnectionMap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceConnectionMap(string name, ServiceConnectionMapArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:networkconnectivity/v1:ServiceConnectionMap", name, args ?? new ServiceConnectionMapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceConnectionMap(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:networkconnectivity/v1:ServiceConnectionMap", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
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
        /// Get an existing ServiceConnectionMap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceConnectionMap Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServiceConnectionMap(name, id, options);
        }
    }

    public sealed class ServiceConnectionMapArgs : global::Pulumi.ResourceArgs
    {
        [Input("consumerPscConfigs")]
        private InputList<Inputs.ConsumerPscConfigArgs>? _consumerPscConfigs;

        /// <summary>
        /// The PSC configurations on consumer side.
        /// </summary>
        public InputList<Inputs.ConsumerPscConfigArgs> ConsumerPscConfigs
        {
            get => _consumerPscConfigs ?? (_consumerPscConfigs = new InputList<Inputs.ConsumerPscConfigArgs>());
            set => _consumerPscConfigs = value;
        }

        /// <summary>
        /// A description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-defined labels.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Immutable. The name of a ServiceConnectionMap. Format: projects/{project}/locations/{location}/serviceConnectionMaps/{service_connection_map} See: https://google.aip.dev/122#fields-representing-resource-names
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("producerPscConfigs")]
        private InputList<Inputs.ProducerPscConfigArgs>? _producerPscConfigs;

        /// <summary>
        /// The PSC configurations on producer side.
        /// </summary>
        public InputList<Inputs.ProducerPscConfigArgs> ProducerPscConfigs
        {
            get => _producerPscConfigs ?? (_producerPscConfigs = new InputList<Inputs.ProducerPscConfigArgs>());
            set => _producerPscConfigs = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// The service class identifier this ServiceConnectionMap is for. The user of ServiceConnectionMap create API needs to have networkconnecitivty.serviceclasses.use iam permission for the service class.
        /// </summary>
        [Input("serviceClass")]
        public Input<string>? ServiceClass { get; set; }

        /// <summary>
        /// Optional. Resource ID (i.e. 'foo' in '[...]/projects/p/locations/l/serviceConnectionMaps/foo') See https://google.aip.dev/122#resource-id-segments Unique per location. If one is not provided, one will be generated.
        /// </summary>
        [Input("serviceConnectionMapId")]
        public Input<string>? ServiceConnectionMapId { get; set; }

        /// <summary>
        /// The token provided by the consumer. This token authenticates that the consumer can create a connecton within the specified project and network.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public ServiceConnectionMapArgs()
        {
        }
        public static new ServiceConnectionMapArgs Empty => new ServiceConnectionMapArgs();
    }
}
