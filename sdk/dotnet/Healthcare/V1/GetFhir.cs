// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1
{
    public static class GetFhir
    {
        /// <summary>
        /// Gets the contents of a FHIR resource. Implements the FHIR standard read interaction ([DSTU2](http://hl7.org/implement/standards/fhir/DSTU2/http.html#read), [STU3](http://hl7.org/implement/standards/fhir/STU3/http.html#read), [R4](http://hl7.org/implement/standards/fhir/R4/http.html#read)). Also supports the FHIR standard conditional read interaction ([DSTU2](http://hl7.org/implement/standards/fhir/DSTU2/http.html#cread), [STU3](http://hl7.org/implement/standards/fhir/STU3/http.html#cread), [R4](http://hl7.org/implement/standards/fhir/R4/http.html#cread)) specified by supplying an `If-Modified-Since` header with a date/time value or an `If-None-Match` header with an ETag value. On success, the response body contains a JSON-encoded representation of the resource. Errors generated by the FHIR store contain a JSON-encoded `OperationOutcome` resource describing the reason for the error. If the request cannot be mapped to a valid API method on a FHIR store, a generic GCP error might be returned instead. For samples that show how to call `read`, see [Getting a FHIR resource](/healthcare/docs/how-tos/fhir-resources#getting_a_fhir_resource).
        /// </summary>
        public static Task<GetFhirResult> InvokeAsync(GetFhirArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFhirResult>("google-native:healthcare/v1:getFhir", args ?? new GetFhirArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the contents of a FHIR resource. Implements the FHIR standard read interaction ([DSTU2](http://hl7.org/implement/standards/fhir/DSTU2/http.html#read), [STU3](http://hl7.org/implement/standards/fhir/STU3/http.html#read), [R4](http://hl7.org/implement/standards/fhir/R4/http.html#read)). Also supports the FHIR standard conditional read interaction ([DSTU2](http://hl7.org/implement/standards/fhir/DSTU2/http.html#cread), [STU3](http://hl7.org/implement/standards/fhir/STU3/http.html#cread), [R4](http://hl7.org/implement/standards/fhir/R4/http.html#cread)) specified by supplying an `If-Modified-Since` header with a date/time value or an `If-None-Match` header with an ETag value. On success, the response body contains a JSON-encoded representation of the resource. Errors generated by the FHIR store contain a JSON-encoded `OperationOutcome` resource describing the reason for the error. If the request cannot be mapped to a valid API method on a FHIR store, a generic GCP error might be returned instead. For samples that show how to call `read`, see [Getting a FHIR resource](/healthcare/docs/how-tos/fhir-resources#getting_a_fhir_resource).
        /// </summary>
        public static Output<GetFhirResult> Invoke(GetFhirInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFhirResult>("google-native:healthcare/v1:getFhir", args ?? new GetFhirInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFhirArgs : global::Pulumi.InvokeArgs
    {
        [Input("datasetId", required: true)]
        public string DatasetId { get; set; } = null!;

        [Input("fhirId", required: true)]
        public string FhirId { get; set; } = null!;

        [Input("fhirId1", required: true)]
        public string FhirId1 { get; set; } = null!;

        [Input("fhirStoreId", required: true)]
        public string FhirStoreId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetFhirArgs()
        {
        }
        public static new GetFhirArgs Empty => new GetFhirArgs();
    }

    public sealed class GetFhirInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("fhirId", required: true)]
        public Input<string> FhirId { get; set; } = null!;

        [Input("fhirId1", required: true)]
        public Input<string> FhirId1 { get; set; } = null!;

        [Input("fhirStoreId", required: true)]
        public Input<string> FhirStoreId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetFhirInvokeArgs()
        {
        }
        public static new GetFhirInvokeArgs Empty => new GetFhirInvokeArgs();
    }


    [OutputType]
    public sealed class GetFhirResult
    {
        /// <summary>
        /// The HTTP Content-Type header value specifying the content type of the body.
        /// </summary>
        public readonly string ContentType;
        /// <summary>
        /// The HTTP request/response body as raw binary.
        /// </summary>
        public readonly string Data;
        /// <summary>
        /// Application specific response metadata. Must be set in the first response for streaming APIs.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, string>> Extensions;

        [OutputConstructor]
        private GetFhirResult(
            string contentType,

            string data,

            ImmutableArray<ImmutableDictionary<string, string>> extensions)
        {
            ContentType = contentType;
            Data = data;
            Extensions = extensions;
        }
    }
}
