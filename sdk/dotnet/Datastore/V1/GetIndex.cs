// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastore.V1
{
    public static class GetIndex
    {
        /// <summary>
        /// Gets an index.
        /// </summary>
        public static Task<GetIndexResult> InvokeAsync(GetIndexArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIndexResult>("google-native:datastore/v1:getIndex", args ?? new GetIndexArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an index.
        /// </summary>
        public static Output<GetIndexResult> Invoke(GetIndexInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetIndexResult>("google-native:datastore/v1:getIndex", args ?? new GetIndexInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIndexArgs : global::Pulumi.InvokeArgs
    {
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
        /// The index's ancestor mode. Must not be ANCESTOR_MODE_UNSPECIFIED.
        /// </summary>
        public readonly string Ancestor;
        /// <summary>
        /// The resource ID of the index.
        /// </summary>
        public readonly string IndexId;
        /// <summary>
        /// The entity kind to which this index applies.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Project ID.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// An ordered sequence of property names and their index attributes.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleDatastoreAdminV1IndexedPropertyResponse> Properties;
        /// <summary>
        /// The state of the index.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetIndexResult(
            string ancestor,

            string indexId,

            string kind,

            string project,

            ImmutableArray<Outputs.GoogleDatastoreAdminV1IndexedPropertyResponse> properties,

            string state)
        {
            Ancestor = ancestor;
            IndexId = indexId;
            Kind = kind;
            Project = project;
            Properties = properties;
            State = state;
        }
    }
}
