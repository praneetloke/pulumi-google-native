// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1
{
    /// <summary>
    /// Creates a `WorkerPool`.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudbuild/v1:WorkerPool")]
    public partial class WorkerPool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableDictionary<string, string>> Annotations { get; private set; } = null!;

        /// <summary>
        /// Time at which the request to create the `WorkerPool` was received.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Time at which the request to delete the `WorkerPool` was received.
        /// </summary>
        [Output("deleteTime")]
        public Output<string> DeleteTime { get; private set; } = null!;

        /// <summary>
        /// A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Checksum computed by the server. May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name of the `WorkerPool`, with format `projects/{project}/locations/{location}/workerPools/{worker_pool}`. The value of `{worker_pool}` is provided by `worker_pool_id` in `CreateWorkerPool` request and the value of `{location}` is determined by the endpoint accessed.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Legacy Private Pool configuration.
        /// </summary>
        [Output("privatePoolV1Config")]
        public Output<Outputs.PrivatePoolV1ConfigResponse> PrivatePoolV1Config { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// `WorkerPool` state.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for the `WorkerPool`.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// Time at which the request to update the `WorkerPool` was received.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// If set, validate the request and preview the response, but do not actually post it.
        /// </summary>
        [Output("validateOnly")]
        public Output<string?> ValidateOnly { get; private set; } = null!;

        /// <summary>
        /// Required. Immutable. The ID to use for the `WorkerPool`, which will become the final component of the resource name. This value should be 1-63 characters, and valid characters are /a-z-/.
        /// </summary>
        [Output("workerPoolId")]
        public Output<string> WorkerPoolId { get; private set; } = null!;


        /// <summary>
        /// Create a WorkerPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkerPool(string name, WorkerPoolArgs args, CustomResourceOptions? options = null)
            : base("google-native:cloudbuild/v1:WorkerPool", name, args ?? new WorkerPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkerPool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudbuild/v1:WorkerPool", name, null, MakeResourceOptions(options, id))
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
                    "workerPoolId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WorkerPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkerPool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WorkerPool(name, id, options);
        }
    }

    public sealed class WorkerPoolArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Legacy Private Pool configuration.
        /// </summary>
        [Input("privatePoolV1Config")]
        public Input<Inputs.PrivatePoolV1ConfigArgs>? PrivatePoolV1Config { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// If set, validate the request and preview the response, but do not actually post it.
        /// </summary>
        [Input("validateOnly")]
        public Input<string>? ValidateOnly { get; set; }

        /// <summary>
        /// Required. Immutable. The ID to use for the `WorkerPool`, which will become the final component of the resource name. This value should be 1-63 characters, and valid characters are /a-z-/.
        /// </summary>
        [Input("workerPoolId", required: true)]
        public Input<string> WorkerPoolId { get; set; } = null!;

        public WorkerPoolArgs()
        {
        }
        public static new WorkerPoolArgs Empty => new WorkerPoolArgs();
    }
}
