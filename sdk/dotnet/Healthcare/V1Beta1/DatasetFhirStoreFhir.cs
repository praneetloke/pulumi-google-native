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
    /// Creates a FHIR resource. Implements the FHIR standard create interaction ([DSTU2](https://hl7.org/implement/standards/fhir/DSTU2/http.html#create), [STU3](https://hl7.org/implement/standards/fhir/STU3/http.html#create), [R4](https://hl7.org/implement/standards/fhir/R4/http.html#create)), which creates a new resource with a server-assigned resource ID. Also supports the FHIR standard conditional create interaction ([DSTU2](https://hl7.org/implement/standards/fhir/DSTU2/http.html#ccreate), [STU3](https://hl7.org/implement/standards/fhir/STU3/http.html#ccreate), [R4](https://hl7.org/implement/standards/fhir/R4/http.html#ccreate)), specified by supplying an `If-None-Exist` header containing a FHIR search query. If no resources match this search query, the server processes the create operation as normal. The request body must contain a JSON-encoded FHIR resource, and the request headers must contain `Content-Type: application/fhir+json`. On success, the response body contains a JSON-encoded representation of the resource as it was created on the server, including the server-assigned resource ID and version ID. Errors generated by the FHIR store contain a JSON-encoded `OperationOutcome` resource describing the reason for the error. If the request cannot be mapped to a valid API method on a FHIR store, a generic GCP error might be returned instead. For samples that show how to call `create`, see [Creating a FHIR resource](/healthcare/docs/how-tos/fhir-resources#creating_a_fhir_resource).
    /// </summary>
    [GoogleNativeResourceType("google-native:healthcare/v1beta1:DatasetFhirStoreFhir")]
    public partial class DatasetFhirStoreFhir : Pulumi.CustomResource
    {
        /// <summary>
        /// The HTTP Content-Type header value specifying the content type of the body.
        /// </summary>
        [Output("contentType")]
        public Output<string> ContentType { get; private set; } = null!;

        /// <summary>
        /// The HTTP request/response body as raw binary.
        /// </summary>
        [Output("data")]
        public Output<string> Data { get; private set; } = null!;

        /// <summary>
        /// Application specific response metadata. Must be set in the first response for streaming APIs.
        /// </summary>
        [Output("extensions")]
        public Output<ImmutableArray<ImmutableDictionary<string, string>>> Extensions { get; private set; } = null!;


        /// <summary>
        /// Create a DatasetFhirStoreFhir resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatasetFhirStoreFhir(string name, DatasetFhirStoreFhirArgs args, CustomResourceOptions? options = null)
            : base("google-native:healthcare/v1beta1:DatasetFhirStoreFhir", name, args ?? new DatasetFhirStoreFhirArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatasetFhirStoreFhir(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:healthcare/v1beta1:DatasetFhirStoreFhir", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DatasetFhirStoreFhir resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatasetFhirStoreFhir Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DatasetFhirStoreFhir(name, id, options);
        }
    }

    public sealed class DatasetFhirStoreFhirArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The HTTP Content-Type header value specifying the content type of the body.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// The HTTP request/response body as raw binary.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("extensions")]
        private InputList<ImmutableDictionary<string, string>>? _extensions;

        /// <summary>
        /// Application specific response metadata. Must be set in the first response for streaming APIs.
        /// </summary>
        public InputList<ImmutableDictionary<string, string>> Extensions
        {
            get => _extensions ?? (_extensions = new InputList<ImmutableDictionary<string, string>>());
            set => _extensions = value;
        }

        [Input("fhirId", required: true)]
        public Input<string> FhirId { get; set; } = null!;

        [Input("fhirId1", required: true)]
        public Input<string> FhirId1 { get; set; } = null!;

        [Input("fhirStoreId", required: true)]
        public Input<string> FhirStoreId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public DatasetFhirStoreFhirArgs()
        {
        }
    }
}
