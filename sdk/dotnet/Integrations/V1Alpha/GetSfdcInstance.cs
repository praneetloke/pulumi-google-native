// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha
{
    public static class GetSfdcInstance
    {
        /// <summary>
        /// Gets an sfdc instance. If the instance doesn't exist, Code.NOT_FOUND exception will be thrown.
        /// </summary>
        public static Task<GetSfdcInstanceResult> InvokeAsync(GetSfdcInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSfdcInstanceResult>("google-native:integrations/v1alpha:getSfdcInstance", args ?? new GetSfdcInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an sfdc instance. If the instance doesn't exist, Code.NOT_FOUND exception will be thrown.
        /// </summary>
        public static Output<GetSfdcInstanceResult> Invoke(GetSfdcInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSfdcInstanceResult>("google-native:integrations/v1alpha:getSfdcInstance", args ?? new GetSfdcInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSfdcInstanceArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("productId", required: true)]
        public string ProductId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("sfdcInstanceId", required: true)]
        public string SfdcInstanceId { get; set; } = null!;

        public GetSfdcInstanceArgs()
        {
        }
        public static new GetSfdcInstanceArgs Empty => new GetSfdcInstanceArgs();
    }

    public sealed class GetSfdcInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("productId", required: true)]
        public Input<string> ProductId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sfdcInstanceId", required: true)]
        public Input<string> SfdcInstanceId { get; set; } = null!;

        public GetSfdcInstanceInvokeArgs()
        {
        }
        public static new GetSfdcInstanceInvokeArgs Empty => new GetSfdcInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetSfdcInstanceResult
    {
        /// <summary>
        /// A list of AuthConfigs that can be tried to open the channel to SFDC
        /// </summary>
        public readonly ImmutableArray<string> AuthConfigId;
        /// <summary>
        /// Time when the instance is created
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Time when the instance was deleted. Empty if not deleted.
        /// </summary>
        public readonly string DeleteTime;
        /// <summary>
        /// A description of the sfdc instance.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// User selected unique name/alias to easily reference an instance.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Resource name of the SFDC instance projects/{project}/locations/{location}/sfdcInstances/{sfdcInstance}.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// URL used for API calls after authentication (the login authority is configured within the referenced AuthConfig).
        /// </summary>
        public readonly string ServiceAuthority;
        /// <summary>
        /// The SFDC Org Id. This is defined in salesforce.
        /// </summary>
        public readonly string SfdcOrgId;
        /// <summary>
        /// Time when the instance was last updated
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetSfdcInstanceResult(
            ImmutableArray<string> authConfigId,

            string createTime,

            string deleteTime,

            string description,

            string displayName,

            string name,

            string serviceAuthority,

            string sfdcOrgId,

            string updateTime)
        {
            AuthConfigId = authConfigId;
            CreateTime = createTime;
            DeleteTime = deleteTime;
            Description = description;
            DisplayName = displayName;
            Name = name;
            ServiceAuthority = serviceAuthority;
            SfdcOrgId = sfdcOrgId;
            UpdateTime = updateTime;
        }
    }
}
