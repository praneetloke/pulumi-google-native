// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1Beta1
{
    /// <summary>
    /// Create a DeploymentResourcePool.
    /// </summary>
    [GoogleNativeResourceType("google-native:aiplatform/v1beta1:DeploymentResourcePool")]
    public partial class DeploymentResourcePool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Timestamp when this DeploymentResourcePool was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The underlying DedicatedResources that the DeploymentResourcePool uses.
        /// </summary>
        [Output("dedicatedResources")]
        public Output<Outputs.GoogleCloudAiplatformV1beta1DedicatedResourcesResponse> DedicatedResources { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Immutable. The resource name of the DeploymentResourcePool. Format: `projects/{project}/locations/{location}/deploymentResourcePools/{deployment_resource_pool}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a DeploymentResourcePool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeploymentResourcePool(string name, DeploymentResourcePoolArgs args, CustomResourceOptions? options = null)
            : base("google-native:aiplatform/v1beta1:DeploymentResourcePool", name, args ?? new DeploymentResourcePoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeploymentResourcePool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:aiplatform/v1beta1:DeploymentResourcePool", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DeploymentResourcePool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeploymentResourcePool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DeploymentResourcePool(name, id, options);
        }
    }

    public sealed class DeploymentResourcePoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The underlying DedicatedResources that the DeploymentResourcePool uses.
        /// </summary>
        [Input("dedicatedResources", required: true)]
        public Input<Inputs.GoogleCloudAiplatformV1beta1DedicatedResourcesArgs> DedicatedResources { get; set; } = null!;

        /// <summary>
        /// The ID to use for the DeploymentResourcePool, which will become the final component of the DeploymentResourcePool's resource name. The maximum length is 63 characters, and valid characters are `/^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$/`.
        /// </summary>
        [Input("deploymentResourcePoolId", required: true)]
        public Input<string> DeploymentResourcePoolId { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Immutable. The resource name of the DeploymentResourcePool. Format: `projects/{project}/locations/{location}/deploymentResourcePools/{deployment_resource_pool}`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public DeploymentResourcePoolArgs()
        {
        }
        public static new DeploymentResourcePoolArgs Empty => new DeploymentResourcePoolArgs();
    }
}