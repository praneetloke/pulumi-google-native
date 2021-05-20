// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    /// <summary>
    /// Creates a Reference in the specified environment.
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:OrganizationEnvironmentReference")]
    public partial class OrganizationEnvironmentReference : Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. A human-readable description of this reference.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Required. The resource id of this reference. Values must match the regular expression [\w\s\-.]+.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Required. The id of the resource to which this reference refers. Must be the id of a resource that exists in the parent environment and is of the given resource_type.
        /// </summary>
        [Output("refers")]
        public Output<string> Refers { get; private set; } = null!;

        /// <summary>
        /// The type of resource referred to by this reference. Valid values are 'KeyStore' or 'TrustStore'.
        /// </summary>
        [Output("resourceType")]
        public Output<string> ResourceType { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationEnvironmentReference resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationEnvironmentReference(string name, OrganizationEnvironmentReferenceArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:OrganizationEnvironmentReference", name, args ?? new OrganizationEnvironmentReferenceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationEnvironmentReference(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:OrganizationEnvironmentReference", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationEnvironmentReference resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationEnvironmentReference Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OrganizationEnvironmentReference(name, id, options);
        }
    }

    public sealed class OrganizationEnvironmentReferenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. A human-readable description of this reference.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        /// <summary>
        /// Required. The resource id of this reference. Values must match the regular expression [\w\s\-.]+.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        [Input("referenceId", required: true)]
        public Input<string> ReferenceId { get; set; } = null!;

        /// <summary>
        /// Required. The id of the resource to which this reference refers. Must be the id of a resource that exists in the parent environment and is of the given resource_type.
        /// </summary>
        [Input("refers")]
        public Input<string>? Refers { get; set; }

        /// <summary>
        /// The type of resource referred to by this reference. Valid values are 'KeyStore' or 'TrustStore'.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        public OrganizationEnvironmentReferenceArgs()
        {
        }
    }
}
