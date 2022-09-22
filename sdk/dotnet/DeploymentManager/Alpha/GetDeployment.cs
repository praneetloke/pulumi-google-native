// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.Alpha
{
    public static class GetDeployment
    {
        /// <summary>
        /// Gets information about a specific deployment.
        /// </summary>
        public static Task<GetDeploymentResult> InvokeAsync(GetDeploymentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeploymentResult>("google-native:deploymentmanager/alpha:getDeployment", args ?? new GetDeploymentArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a specific deployment.
        /// </summary>
        public static Output<GetDeploymentResult> Invoke(GetDeploymentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeploymentResult>("google-native:deploymentmanager/alpha:getDeployment", args ?? new GetDeploymentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeploymentArgs : global::Pulumi.InvokeArgs
    {
        [Input("deployment", required: true)]
        public string Deployment { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetDeploymentArgs()
        {
        }
        public static new GetDeploymentArgs Empty => new GetDeploymentArgs();
    }

    public sealed class GetDeploymentInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("deployment", required: true)]
        public Input<string> Deployment { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetDeploymentInvokeArgs()
        {
        }
        public static new GetDeploymentInvokeArgs Empty => new GetDeploymentInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeploymentResult
    {
        /// <summary>
        /// User provided default credential for the deployment.
        /// </summary>
        public readonly Outputs.CredentialResponse Credential;
        /// <summary>
        /// An optional user-provided description of the deployment.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Provides a fingerprint to use in requests to modify a deployment, such as `update()`, `stop()`, and `cancelPreview()` requests. A fingerprint is a randomly generated value that must be provided with `update()`, `stop()`, and `cancelPreview()` requests to perform optimistic locking. This ensures optimistic concurrency so that only one request happens at a time. The fingerprint is initially generated by Deployment Manager and changes after every request to modify data. To get the latest fingerprint value, perform a `get()` request to a deployment.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string InsertTime;
        /// <summary>
        /// Map of One Platform labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
        /// </summary>
        public readonly ImmutableArray<Outputs.DeploymentLabelEntryResponse> Labels;
        /// <summary>
        /// URL of the manifest representing the last manifest that was successfully deployed. If no manifest has been successfully deployed, this field will be absent.
        /// </summary>
        public readonly string Manifest;
        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Operation that most recently ran, or is currently running, on this deployment.
        /// </summary>
        public readonly Outputs.OperationResponse Operation;
        /// <summary>
        /// List of outputs from the last manifest that deployed successfully.
        /// </summary>
        public readonly ImmutableArray<Outputs.DeploymentOutputEntryResponse> Outputs;
        /// <summary>
        /// Server defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// [Input Only] The parameters that define your deployment, including the deployment configuration and relevant templates.
        /// </summary>
        public readonly Outputs.TargetConfigurationResponse Target;
        /// <summary>
        /// If Deployment Manager is currently updating or previewing an update to this deployment, the updated configuration appears here.
        /// </summary>
        public readonly Outputs.DeploymentUpdateResponse Update;
        /// <summary>
        /// Update timestamp in RFC3339 text format.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetDeploymentResult(
            Outputs.CredentialResponse credential,

            string description,

            string fingerprint,

            string insertTime,

            ImmutableArray<Outputs.DeploymentLabelEntryResponse> labels,

            string manifest,

            string name,

            Outputs.OperationResponse operation,

            ImmutableArray<Outputs.DeploymentOutputEntryResponse> outputs,

            string selfLink,

            Outputs.TargetConfigurationResponse target,

            Outputs.DeploymentUpdateResponse update,

            string updateTime)
        {
            Credential = credential;
            Description = description;
            Fingerprint = fingerprint;
            InsertTime = insertTime;
            Labels = labels;
            Manifest = manifest;
            Name = name;
            Operation = operation;
            Outputs = outputs;
            SelfLink = selfLink;
            Target = target;
            Update = update;
            UpdateTime = updateTime;
        }
    }
}
