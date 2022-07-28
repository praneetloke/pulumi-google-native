// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1
{
    public static class GetConsentStore
    {
        /// <summary>
        /// Gets the specified consent store.
        /// </summary>
        public static Task<GetConsentStoreResult> InvokeAsync(GetConsentStoreArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConsentStoreResult>("google-native:healthcare/v1:getConsentStore", args ?? new GetConsentStoreArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified consent store.
        /// </summary>
        public static Output<GetConsentStoreResult> Invoke(GetConsentStoreInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConsentStoreResult>("google-native:healthcare/v1:getConsentStore", args ?? new GetConsentStoreInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConsentStoreArgs : global::Pulumi.InvokeArgs
    {
        [Input("consentStoreId", required: true)]
        public string ConsentStoreId { get; set; } = null!;

        [Input("datasetId", required: true)]
        public string DatasetId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetConsentStoreArgs()
        {
        }
        public static new GetConsentStoreArgs Empty => new GetConsentStoreArgs();
    }

    public sealed class GetConsentStoreInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("consentStoreId", required: true)]
        public Input<string> ConsentStoreId { get; set; } = null!;

        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetConsentStoreInvokeArgs()
        {
        }
        public static new GetConsentStoreInvokeArgs Empty => new GetConsentStoreInvokeArgs();
    }


    [OutputType]
    public sealed class GetConsentStoreResult
    {
        /// <summary>
        /// Optional. Default time to live for Consents created in this store. Must be at least 24 hours. Updating this field will not affect the expiration time of existing consents.
        /// </summary>
        public readonly string DefaultConsentTtl;
        /// <summary>
        /// Optional. If `true`, UpdateConsent creates the Consent if it does not already exist. If unspecified, defaults to `false`.
        /// </summary>
        public readonly bool EnableConsentCreateOnUpdate;
        /// <summary>
        /// Optional. User-supplied key-value pairs used to organize consent stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62}. Label values must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}. No more than 64 labels can be associated with a given store. For more information: https://cloud.google.com/healthcare/docs/how-tos/labeling-resources
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Resource name of the consent store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}`. Cannot be changed after creation.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetConsentStoreResult(
            string defaultConsentTtl,

            bool enableConsentCreateOnUpdate,

            ImmutableDictionary<string, string> labels,

            string name)
        {
            DefaultConsentTtl = defaultConsentTtl;
            EnableConsentCreateOnUpdate = enableConsentCreateOnUpdate;
            Labels = labels;
            Name = name;
        }
    }
}
