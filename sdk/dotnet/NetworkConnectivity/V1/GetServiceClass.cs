// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkConnectivity.V1
{
    public static class GetServiceClass
    {
        /// <summary>
        /// Gets details of a single ServiceClass.
        /// </summary>
        public static Task<GetServiceClassResult> InvokeAsync(GetServiceClassArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceClassResult>("google-native:networkconnectivity/v1:getServiceClass", args ?? new GetServiceClassArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single ServiceClass.
        /// </summary>
        public static Output<GetServiceClassResult> Invoke(GetServiceClassInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceClassResult>("google-native:networkconnectivity/v1:getServiceClass", args ?? new GetServiceClassInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceClassArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("serviceClassId", required: true)]
        public string ServiceClassId { get; set; } = null!;

        public GetServiceClassArgs()
        {
        }
        public static new GetServiceClassArgs Empty => new GetServiceClassArgs();
    }

    public sealed class GetServiceClassInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("serviceClassId", required: true)]
        public Input<string> ServiceClassId { get; set; } = null!;

        public GetServiceClassInvokeArgs()
        {
        }
        public static new GetServiceClassInvokeArgs Empty => new GetServiceClassInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceClassResult
    {
        /// <summary>
        /// Time when the ServiceClass was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// A description of this resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// User-defined labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Immutable. The name of a ServiceClass resource. Format: projects/{project}/locations/{location}/serviceClasses/{service_class} See: https://google.aip.dev/122#fields-representing-resource-names
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The generated service class name. Use this name to refer to the Service class in Service Connection Maps and Service Connection Policies.
        /// </summary>
        public readonly string ServiceClass;
        /// <summary>
        /// URIs of all Service Connection Maps using this service class.
        /// </summary>
        public readonly ImmutableArray<string> ServiceConnectionMaps;
        /// <summary>
        /// Time when the ServiceClass was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetServiceClassResult(
            string createTime,

            string description,

            ImmutableDictionary<string, string> labels,

            string name,

            string serviceClass,

            ImmutableArray<string> serviceConnectionMaps,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            Labels = labels;
            Name = name;
            ServiceClass = serviceClass;
            ServiceConnectionMaps = serviceConnectionMaps;
            UpdateTime = updateTime;
        }
    }
}
