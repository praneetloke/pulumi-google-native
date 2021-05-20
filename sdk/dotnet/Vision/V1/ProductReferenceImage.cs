// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Vision.V1
{
    /// <summary>
    /// Creates and returns a new ReferenceImage resource. The `bounding_poly` field is optional. If `bounding_poly` is not specified, the system will try to detect regions of interest in the image that are compatible with the product_category on the parent product. If it is specified, detection is ALWAYS skipped. The system converts polygons into non-rotated rectangles. Note that the pipeline will resize the image if the image resolution is too large to process (above 50MP). Possible errors: * Returns INVALID_ARGUMENT if the image_uri is missing or longer than 4096 characters. * Returns INVALID_ARGUMENT if the product does not exist. * Returns INVALID_ARGUMENT if bounding_poly is not provided, and nothing compatible with the parent product's product_category is detected. * Returns INVALID_ARGUMENT if bounding_poly contains more than 10 polygons.
    /// </summary>
    [GoogleNativeResourceType("google-native:vision/v1:ProductReferenceImage")]
    public partial class ProductReferenceImage : Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. Bounding polygons around the areas of interest in the reference image. If this field is empty, the system will try to detect regions of interest. At most 10 bounding polygons will be used. The provided shape is converted into a non-rotated rectangle. Once converted, the small edge of the rectangle must be greater than or equal to 300 pixels. The aspect ratio must be 1:4 or less (i.e. 1:3 is ok; 1:5 is not).
        /// </summary>
        [Output("boundingPolys")]
        public Output<ImmutableArray<Outputs.BoundingPolyResponse>> BoundingPolys { get; private set; } = null!;

        /// <summary>
        /// The resource name of the reference image. Format is: `projects/PROJECT_ID/locations/LOC_ID/products/PRODUCT_ID/referenceImages/IMAGE_ID`. This field is ignored when creating a reference image.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Required. The Google Cloud Storage URI of the reference image. The URI must start with `gs://`.
        /// </summary>
        [Output("uri")]
        public Output<string> Uri { get; private set; } = null!;


        /// <summary>
        /// Create a ProductReferenceImage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProductReferenceImage(string name, ProductReferenceImageArgs args, CustomResourceOptions? options = null)
            : base("google-native:vision/v1:ProductReferenceImage", name, args ?? new ProductReferenceImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProductReferenceImage(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:vision/v1:ProductReferenceImage", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ProductReferenceImage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProductReferenceImage Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ProductReferenceImage(name, id, options);
        }
    }

    public sealed class ProductReferenceImageArgs : Pulumi.ResourceArgs
    {
        [Input("boundingPolys")]
        private InputList<Inputs.BoundingPolyArgs>? _boundingPolys;

        /// <summary>
        /// Optional. Bounding polygons around the areas of interest in the reference image. If this field is empty, the system will try to detect regions of interest. At most 10 bounding polygons will be used. The provided shape is converted into a non-rotated rectangle. Once converted, the small edge of the rectangle must be greater than or equal to 300 pixels. The aspect ratio must be 1:4 or less (i.e. 1:3 is ok; 1:5 is not).
        /// </summary>
        public InputList<Inputs.BoundingPolyArgs> BoundingPolys
        {
            get => _boundingPolys ?? (_boundingPolys = new InputList<Inputs.BoundingPolyArgs>());
            set => _boundingPolys = value;
        }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The resource name of the reference image. Format is: `projects/PROJECT_ID/locations/LOC_ID/products/PRODUCT_ID/referenceImages/IMAGE_ID`. This field is ignored when creating a reference image.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("productId", required: true)]
        public Input<string> ProductId { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("referenceImageId")]
        public Input<string>? ReferenceImageId { get; set; }

        /// <summary>
        /// Required. The Google Cloud Storage URI of the reference image. The URI must start with `gs://`.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public ProductReferenceImageArgs()
        {
        }
    }
}
