// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Firestore.V1Beta1
{
    public static class GetIndex
    {
        /// <summary>
        /// Gets an index.
        /// </summary>
        public static Task<GetIndexResult> InvokeAsync(GetIndexArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIndexResult>("google-native:firestore/v1beta1:getIndex", args ?? new GetIndexArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an index.
        /// </summary>
        public static Output<GetIndexResult> Invoke(GetIndexInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIndexResult>("google-native:firestore/v1beta1:getIndex", args ?? new GetIndexInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIndexArgs : global::Pulumi.InvokeArgs
    {
        [Input("databaseId", required: true)]
        public string DatabaseId { get; set; } = null!;

        [Input("indexId", required: true)]
        public string IndexId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetIndexArgs()
        {
        }
        public static new GetIndexArgs Empty => new GetIndexArgs();
    }

    public sealed class GetIndexInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("databaseId", required: true)]
        public Input<string> DatabaseId { get; set; } = null!;

        [Input("indexId", required: true)]
        public Input<string> IndexId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetIndexInvokeArgs()
        {
        }
        public static new GetIndexInvokeArgs Empty => new GetIndexInvokeArgs();
    }


    [OutputType]
    public sealed class GetIndexResult
    {
        /// <summary>
        /// The collection ID to which this index applies. Required.
        /// </summary>
        public readonly string CollectionId;
        /// <summary>
        /// The fields to index.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleFirestoreAdminV1beta1IndexFieldResponse> Fields;
        /// <summary>
        /// The resource name of the index. Output only.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The state of the index. Output only.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetIndexResult(
            string collectionId,

            ImmutableArray<Outputs.GoogleFirestoreAdminV1beta1IndexFieldResponse> fields,

            string name,

            string state)
        {
            CollectionId = collectionId;
            Fields = fields;
            Name = name;
            State = state;
        }
    }
}
