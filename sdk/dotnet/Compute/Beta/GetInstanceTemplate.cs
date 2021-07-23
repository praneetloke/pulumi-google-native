// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetInstanceTemplate
    {
        /// <summary>
        /// Returns the specified instance template. Gets a list of available instance templates by making a list() request.
        /// </summary>
        public static Task<GetInstanceTemplateResult> InvokeAsync(GetInstanceTemplateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceTemplateResult>("google-native:compute/beta:getInstanceTemplate", args ?? new GetInstanceTemplateArgs(), options.WithVersion());
    }


    public sealed class GetInstanceTemplateArgs : Pulumi.InvokeArgs
    {
        [Input("instanceTemplate", required: true)]
        public string InstanceTemplate { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetInstanceTemplateArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceTemplateResult
    {
        /// <summary>
        /// The creation timestamp for this instance template in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The resource type, which is always compute#instanceTemplate for instance templates.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The instance properties for this instance template.
        /// </summary>
        public readonly Outputs.InstancePropertiesResponse Properties;
        /// <summary>
        /// The URL for this instance template. The server defines this URL.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The source instance used to create the template. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance 
        /// </summary>
        public readonly string SourceInstance;
        /// <summary>
        /// The source instance params to use to create this instance template.
        /// </summary>
        public readonly Outputs.SourceInstanceParamsResponse SourceInstanceParams;

        [OutputConstructor]
        private GetInstanceTemplateResult(
            string creationTimestamp,

            string description,

            string kind,

            string name,

            Outputs.InstancePropertiesResponse properties,

            string selfLink,

            string sourceInstance,

            Outputs.SourceInstanceParamsResponse sourceInstanceParams)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            Kind = kind;
            Name = name;
            Properties = properties;
            SelfLink = selfLink;
            SourceInstance = sourceInstance;
            SourceInstanceParams = sourceInstanceParams;
        }
    }
}