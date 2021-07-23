// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.APIGateway.V1
{
    public static class GetApi
    {
        /// <summary>
        /// Gets details of a single Api.
        /// </summary>
        public static Task<GetApiResult> InvokeAsync(GetApiArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApiResult>("google-native:apigateway/v1:getApi", args ?? new GetApiArgs(), options.WithVersion());
    }


    public sealed class GetApiArgs : Pulumi.InvokeArgs
    {
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetApiArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApiResult
    {
        /// <summary>
        /// Created time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. Display name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Optional. Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed). If not specified, a new Service will automatically be created in the same project as this API.
        /// </summary>
        public readonly string ManagedService;
        /// <summary>
        /// Resource name of the API. Format: projects/{project}/locations/global/apis/{api}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// State of the API.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Updated time.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetApiResult(
            string createTime,

            string displayName,

            ImmutableDictionary<string, string> labels,

            string managedService,

            string name,

            string state,

            string updateTime)
        {
            CreateTime = createTime;
            DisplayName = displayName;
            Labels = labels;
            ManagedService = managedService;
            Name = name;
            State = state;
            UpdateTime = updateTime;
        }
    }
}