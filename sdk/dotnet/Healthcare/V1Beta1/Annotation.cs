// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1
{
    /// <summary>
    /// Creates a new Annotation record. It is valid to create Annotation objects for the same source more than once since a unique ID is assigned to each record by this service.
    /// </summary>
    [GoogleNativeResourceType("google-native:healthcare/v1beta1:Annotation")]
    public partial class Annotation : Pulumi.CustomResource
    {
        /// <summary>
        /// Details of the source.
        /// </summary>
        [Output("annotationSource")]
        public Output<Outputs.AnnotationSourceResponse> AnnotationSource { get; private set; } = null!;

        /// <summary>
        /// Additional information for this annotation record, such as annotator and verifier information or study campaign.
        /// </summary>
        [Output("customData")]
        public Output<ImmutableDictionary<string, string>> CustomData { get; private set; } = null!;

        /// <summary>
        /// Annotations for images. For example, bounding polygons.
        /// </summary>
        [Output("imageAnnotation")]
        public Output<Outputs.ImageAnnotationResponse> ImageAnnotation { get; private set; } = null!;

        /// <summary>
        /// Resource name of the Annotation, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/annotationStores/{annotation_store_id}/annotations/{annotation_id}`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Annotations for resource. For example, classification tags.
        /// </summary>
        [Output("resourceAnnotation")]
        public Output<Outputs.ResourceAnnotationResponse> ResourceAnnotation { get; private set; } = null!;

        /// <summary>
        /// Annotations for sensitive texts. For example, a range that describes the location of sensitive text.
        /// </summary>
        [Output("textAnnotation")]
        public Output<Outputs.SensitiveTextAnnotationResponse> TextAnnotation { get; private set; } = null!;


        /// <summary>
        /// Create a Annotation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Annotation(string name, AnnotationArgs args, CustomResourceOptions? options = null)
            : base("google-native:healthcare/v1beta1:Annotation", name, args ?? new AnnotationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Annotation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:healthcare/v1beta1:Annotation", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Annotation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Annotation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Annotation(name, id, options);
        }
    }

    public sealed class AnnotationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Details of the source.
        /// </summary>
        [Input("annotationSource")]
        public Input<Inputs.AnnotationSourceArgs>? AnnotationSource { get; set; }

        [Input("annotationStoreId", required: true)]
        public Input<string> AnnotationStoreId { get; set; } = null!;

        [Input("customData")]
        private InputMap<string>? _customData;

        /// <summary>
        /// Additional information for this annotation record, such as annotator and verifier information or study campaign.
        /// </summary>
        public InputMap<string> CustomData
        {
            get => _customData ?? (_customData = new InputMap<string>());
            set => _customData = value;
        }

        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        /// <summary>
        /// Annotations for images. For example, bounding polygons.
        /// </summary>
        [Input("imageAnnotation")]
        public Input<Inputs.ImageAnnotationArgs>? ImageAnnotation { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Resource name of the Annotation, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/annotationStores/{annotation_store_id}/annotations/{annotation_id}`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Annotations for resource. For example, classification tags.
        /// </summary>
        [Input("resourceAnnotation")]
        public Input<Inputs.ResourceAnnotationArgs>? ResourceAnnotation { get; set; }

        /// <summary>
        /// Annotations for sensitive texts. For example, a range that describes the location of sensitive text.
        /// </summary>
        [Input("textAnnotation")]
        public Input<Inputs.SensitiveTextAnnotationArgs>? TextAnnotation { get; set; }

        public AnnotationArgs()
        {
        }
    }
}
