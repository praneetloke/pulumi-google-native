// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1
{
    public static class GetStream
    {
        /// <summary>
        /// Use this method to get details about a stream.
        /// </summary>
        public static Task<GetStreamResult> InvokeAsync(GetStreamArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStreamResult>("google-native:datastream/v1:getStream", args ?? new GetStreamArgs(), options.WithDefaults());

        /// <summary>
        /// Use this method to get details about a stream.
        /// </summary>
        public static Output<GetStreamResult> Invoke(GetStreamInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStreamResult>("google-native:datastream/v1:getStream", args ?? new GetStreamInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStreamArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("streamId", required: true)]
        public string StreamId { get; set; } = null!;

        public GetStreamArgs()
        {
        }
        public static new GetStreamArgs Empty => new GetStreamArgs();
    }

    public sealed class GetStreamInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("streamId", required: true)]
        public Input<string> StreamId { get; set; } = null!;

        public GetStreamInvokeArgs()
        {
        }
        public static new GetStreamInvokeArgs Empty => new GetStreamInvokeArgs();
    }


    [OutputType]
    public sealed class GetStreamResult
    {
        /// <summary>
        /// Automatically backfill objects included in the stream source configuration. Specific objects can be excluded.
        /// </summary>
        public readonly Outputs.BackfillAllStrategyResponse BackfillAll;
        /// <summary>
        /// Do not automatically backfill any objects.
        /// </summary>
        public readonly Outputs.BackfillNoneStrategyResponse BackfillNone;
        /// <summary>
        /// The creation time of the stream.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Immutable. A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data will be encrypted using an internal Stream-specific encryption key provisioned through KMS.
        /// </summary>
        public readonly string CustomerManagedEncryptionKey;
        /// <summary>
        /// Destination connection profile configuration.
        /// </summary>
        public readonly Outputs.DestinationConfigResponse DestinationConfig;
        /// <summary>
        /// Display name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Errors on the Stream.
        /// </summary>
        public readonly ImmutableArray<Outputs.ErrorResponse> Errors;
        /// <summary>
        /// Labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The stream's name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Source connection profile configuration.
        /// </summary>
        public readonly Outputs.SourceConfigResponse SourceConfig;
        /// <summary>
        /// The state of the stream.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The last update time of the stream.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetStreamResult(
            Outputs.BackfillAllStrategyResponse backfillAll,

            Outputs.BackfillNoneStrategyResponse backfillNone,

            string createTime,

            string customerManagedEncryptionKey,

            Outputs.DestinationConfigResponse destinationConfig,

            string displayName,

            ImmutableArray<Outputs.ErrorResponse> errors,

            ImmutableDictionary<string, string> labels,

            string name,

            Outputs.SourceConfigResponse sourceConfig,

            string state,

            string updateTime)
        {
            BackfillAll = backfillAll;
            BackfillNone = backfillNone;
            CreateTime = createTime;
            CustomerManagedEncryptionKey = customerManagedEncryptionKey;
            DestinationConfig = destinationConfig;
            DisplayName = displayName;
            Errors = errors;
            Labels = labels;
            Name = name;
            SourceConfig = sourceConfig;
            State = state;
            UpdateTime = updateTime;
        }
    }
}
