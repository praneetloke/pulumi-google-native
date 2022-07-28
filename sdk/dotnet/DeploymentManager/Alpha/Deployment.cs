// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.Alpha
{
    /// <summary>
    /// Creates a deployment and all of the resources described by the deployment manifest.
    /// </summary>
    [GoogleNativeResourceType("google-native:deploymentmanager/alpha:Deployment")]
    public partial class Deployment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sets the policy to use for creating new resources.
        /// </summary>
        [Output("createPolicy")]
        public Output<string?> CreatePolicy { get; private set; } = null!;

        /// <summary>
        /// User provided default credential for the deployment.
        /// </summary>
        [Output("credential")]
        public Output<Outputs.CredentialResponse> Credential { get; private set; } = null!;

        /// <summary>
        /// An optional user-provided description of the deployment.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Provides a fingerprint to use in requests to modify a deployment, such as `update()`, `stop()`, and `cancelPreview()` requests. A fingerprint is a randomly generated value that must be provided with `update()`, `stop()`, and `cancelPreview()` requests to perform optimistic locking. This ensures optimistic concurrency so that only one request happens at a time. The fingerprint is initially generated by Deployment Manager and changes after every request to modify data. To get the latest fingerprint value, perform a `get()` request to a deployment.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("insertTime")]
        public Output<string> InsertTime { get; private set; } = null!;

        /// <summary>
        /// Map of One Platform labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<Outputs.DeploymentLabelEntryResponse>> Labels { get; private set; } = null!;

        /// <summary>
        /// URL of the manifest representing the last manifest that was successfully deployed. If no manifest has been successfully deployed, this field will be absent.
        /// </summary>
        [Output("manifest")]
        public Output<string> Manifest { get; private set; } = null!;

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Operation that most recently ran, or is currently running, on this deployment.
        /// </summary>
        [Output("operation")]
        public Output<Outputs.OperationResponse> Operation { get; private set; } = null!;

        /// <summary>
        /// List of outputs from the last manifest that deployed successfully.
        /// </summary>
        [Output("outputs")]
        public Output<ImmutableArray<Outputs.DeploymentOutputEntryResponse>> Outputs { get; private set; } = null!;

        /// <summary>
        /// If set to true, creates a deployment and creates "shell" resources but does not actually instantiate these resources. This allows you to preview what your deployment looks like. After previewing a deployment, you can deploy your resources by making a request with the `update()` method or you can use the `cancelPreview()` method to cancel the preview altogether. Note that the deployment will still exist after you cancel the preview and you must separately delete this deployment if you want to remove it.
        /// </summary>
        [Output("preview")]
        public Output<string?> Preview { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Server defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// [Input Only] The parameters that define your deployment, including the deployment configuration and relevant templates.
        /// </summary>
        [Output("target")]
        public Output<Outputs.TargetConfigurationResponse> Target { get; private set; } = null!;

        /// <summary>
        /// If Deployment Manager is currently updating or previewing an update to this deployment, the updated configuration appears here.
        /// </summary>
        [Output("update")]
        public Output<Outputs.DeploymentUpdateResponse> Update { get; private set; } = null!;

        /// <summary>
        /// Update timestamp in RFC3339 text format.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Deployment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Deployment(string name, DeploymentArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:deploymentmanager/alpha:Deployment", name, args ?? new DeploymentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Deployment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:deploymentmanager/alpha:Deployment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Deployment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Deployment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Deployment(name, id, options);
        }
    }

    public sealed class DeploymentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sets the policy to use for creating new resources.
        /// </summary>
        [Input("createPolicy")]
        public Input<string>? CreatePolicy { get; set; }

        /// <summary>
        /// User provided default credential for the deployment.
        /// </summary>
        [Input("credential")]
        public Input<Inputs.CredentialArgs>? Credential { get; set; }

        /// <summary>
        /// An optional user-provided description of the deployment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("labels")]
        private InputList<Inputs.DeploymentLabelEntryArgs>? _labels;

        /// <summary>
        /// Map of One Platform labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
        /// </summary>
        public InputList<Inputs.DeploymentLabelEntryArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.DeploymentLabelEntryArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If set to true, creates a deployment and creates "shell" resources but does not actually instantiate these resources. This allows you to preview what your deployment looks like. After previewing a deployment, you can deploy your resources by making a request with the `update()` method or you can use the `cancelPreview()` method to cancel the preview altogether. Note that the deployment will still exist after you cancel the preview and you must separately delete this deployment if you want to remove it.
        /// </summary>
        [Input("preview")]
        public Input<string>? Preview { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// [Input Only] The parameters that define your deployment, including the deployment configuration and relevant templates.
        /// </summary>
        [Input("target")]
        public Input<Inputs.TargetConfigurationArgs>? Target { get; set; }

        public DeploymentArgs()
        {
        }
        public static new DeploymentArgs Empty => new DeploymentArgs();
    }
}
