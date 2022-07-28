// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CertificateManager.V1
{
    /// <summary>
    /// Creates a new CertificateMap in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:certificatemanager/v1:CertificateMap")]
    public partial class CertificateMap : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Required. A user-provided name of the certificate map.
        /// </summary>
        [Output("certificateMapId")]
        public Output<string> CertificateMapId { get; private set; } = null!;

        /// <summary>
        /// The creation timestamp of a Certificate Map.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// One or more paragraphs of text description of a certificate map.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// A list of GCLB targets which use this Certificate Map.
        /// </summary>
        [Output("gclbTargets")]
        public Output<ImmutableArray<Outputs.GclbTargetResponse>> GclbTargets { get; private set; } = null!;

        /// <summary>
        /// Set of labels associated with a Certificate Map.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// A user-defined name of the Certificate Map. Certificate Map names must be unique globally and match pattern `projects/*/locations/*/certificateMaps/*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The update timestamp of a Certificate Map.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a CertificateMap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertificateMap(string name, CertificateMapArgs args, CustomResourceOptions? options = null)
            : base("google-native:certificatemanager/v1:CertificateMap", name, args ?? new CertificateMapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertificateMap(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:certificatemanager/v1:CertificateMap", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "certificateMapId",
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
        /// Get an existing CertificateMap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertificateMap Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CertificateMap(name, id, options);
        }
    }

    public sealed class CertificateMapArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. A user-provided name of the certificate map.
        /// </summary>
        [Input("certificateMapId", required: true)]
        public Input<string> CertificateMapId { get; set; } = null!;

        /// <summary>
        /// One or more paragraphs of text description of a certificate map.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of labels associated with a Certificate Map.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A user-defined name of the Certificate Map. Certificate Map names must be unique globally and match pattern `projects/*/locations/*/certificateMaps/*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public CertificateMapArgs()
        {
        }
        public static new CertificateMapArgs Empty => new CertificateMapArgs();
    }
}
