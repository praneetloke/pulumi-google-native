// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1
{
    /// <summary>
    /// Create a metadata partition.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:dataplex/v1:Partition")]
    public partial class Partition : Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. The etag for this partition.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Immutable. The location of the entity data within the partition, for example, gs://bucket/path/to/entity/key1=value1/key2=value2. Or projects//datasets//tables/
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Partition values used in the HTTP URL must be double encoded. For example, url_encode(url_encode(value)) can be used to encode "US:CA/CA#Sunnyvale so that the request URL ends with "/partitions/US%253ACA/CA%2523Sunnyvale". The name field in the response retains the encoded format.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Immutable. The set of values representing the partition, which correspond to the partition schema defined in the parent entity.
        /// </summary>
        [Output("values")]
        public Output<ImmutableArray<string>> Values { get; private set; } = null!;


        /// <summary>
        /// Create a Partition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Partition(string name, PartitionArgs args, CustomResourceOptions? options = null)
            : base("google-native:dataplex/v1:Partition", name, args ?? new PartitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Partition(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dataplex/v1:Partition", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Partition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Partition Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Partition(name, id, options);
        }
    }

    public sealed class PartitionArgs : Pulumi.ResourceArgs
    {
        [Input("entityId", required: true)]
        public Input<string> EntityId { get; set; } = null!;

        /// <summary>
        /// Optional. The etag for this partition.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("lakeId", required: true)]
        public Input<string> LakeId { get; set; } = null!;

        /// <summary>
        /// Immutable. The location of the entity data within the partition, for example, gs://bucket/path/to/entity/key1=value1/key2=value2. Or projects//datasets//tables/
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. Only validate the request, but do not perform mutations. The default is false.
        /// </summary>
        [Input("validateOnly")]
        public Input<string>? ValidateOnly { get; set; }

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// Immutable. The set of values representing the partition, which correspond to the partition schema defined in the parent entity.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public PartitionArgs()
        {
        }
    }
}
