// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ArtifactRegistry.V1Beta2
{
    /// <summary>
    /// Creates a tag.
    /// </summary>
    [GoogleNativeResourceType("google-native:artifactregistry/v1beta2:RepositoryPackageTag")]
    public partial class RepositoryPackageTag : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the tag, for example: "projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/tags/tag1".
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the version the tag refers to, for example: "projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/versions/sha256:5243811"
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryPackageTag resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryPackageTag(string name, RepositoryPackageTagArgs args, CustomResourceOptions? options = null)
            : base("google-native:artifactregistry/v1beta2:RepositoryPackageTag", name, args ?? new RepositoryPackageTagArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryPackageTag(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:artifactregistry/v1beta2:RepositoryPackageTag", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryPackageTag resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryPackageTag Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RepositoryPackageTag(name, id, options);
        }
    }

    public sealed class RepositoryPackageTagArgs : Pulumi.ResourceArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the tag, for example: "projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/tags/tag1".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("packageId", required: true)]
        public Input<string> PackageId { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("repositoryId", required: true)]
        public Input<string> RepositoryId { get; set; } = null!;

        [Input("tagId")]
        public Input<string>? TagId { get; set; }

        /// <summary>
        /// The name of the version the tag refers to, for example: "projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/versions/sha256:5243811"
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public RepositoryPackageTagArgs()
        {
        }
    }
}
