// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Cloudkms.V1
{
    /// <summary>
    /// Creates a new EkmConnection in a given Project and Location.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudkms/v1:EkmConnection")]
    public partial class EkmConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// The time at which the EkmConnection was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// This checksum is computed by the server based on the value of other fields, and may be sent on update requests to ensure the client has an up-to-date value before proceeding.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The resource name for the EkmConnection in the format `projects/*/locations/*/ekmConnections/*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of ServiceResolvers where the EKM can be reached. There should be one ServiceResolver per EKM replica. Currently, only a single ServiceResolver is supported.
        /// </summary>
        [Output("serviceResolvers")]
        public Output<ImmutableArray<Outputs.ServiceResolverResponse>> ServiceResolvers { get; private set; } = null!;


        /// <summary>
        /// Create a EkmConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EkmConnection(string name, EkmConnectionArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:cloudkms/v1:EkmConnection", name, args ?? new EkmConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EkmConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudkms/v1:EkmConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EkmConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EkmConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EkmConnection(name, id, options);
        }
    }

    public sealed class EkmConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. It must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`.
        /// </summary>
        [Input("ekmConnectionId")]
        public Input<string>? EkmConnectionId { get; set; }

        /// <summary>
        /// This checksum is computed by the server based on the value of other fields, and may be sent on update requests to ensure the client has an up-to-date value before proceeding.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("serviceResolvers")]
        private InputList<Inputs.ServiceResolverArgs>? _serviceResolvers;

        /// <summary>
        /// A list of ServiceResolvers where the EKM can be reached. There should be one ServiceResolver per EKM replica. Currently, only a single ServiceResolver is supported.
        /// </summary>
        public InputList<Inputs.ServiceResolverArgs> ServiceResolvers
        {
            get => _serviceResolvers ?? (_serviceResolvers = new InputList<Inputs.ServiceResolverArgs>());
            set => _serviceResolvers = value;
        }

        public EkmConnectionArgs()
        {
        }
    }
}
