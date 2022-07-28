// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AnalyticsHub.V1Beta1
{
    /// <summary>
    /// Creates a new listing.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:analyticshub/v1beta1:Listing")]
    public partial class Listing : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Shared dataset i.e. BigQuery dataset source.
        /// </summary>
        [Output("bigqueryDataset")]
        public Output<Outputs.BigQueryDatasetSourceResponse> BigqueryDataset { get; private set; } = null!;

        /// <summary>
        /// Optional. Categories of the listing. Up to two categories are allowed.
        /// </summary>
        [Output("categories")]
        public Output<ImmutableArray<string>> Categories { get; private set; } = null!;

        [Output("dataExchangeId")]
        public Output<string> DataExchangeId { get; private set; } = null!;

        /// <summary>
        /// Optional. Details of the data provider who owns the source data.
        /// </summary>
        [Output("dataProvider")]
        public Output<Outputs.DataProviderResponse> DataProvider { get; private set; } = null!;

        /// <summary>
        /// Optional. Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF). Default value is an empty string. Max length: 2000 bytes.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&amp;) and can't start or end with spaces. Default value is an empty string. Max length: 63 bytes.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Optional. Documentation describing the listing.
        /// </summary>
        [Output("documentation")]
        public Output<string> Documentation { get; private set; } = null!;

        /// <summary>
        /// Optional. Base64 encoded image representing the listing. Max Size: 3.0MiB Expected image dimensions are 512x512 pixels, however the API only performs validation on size of the encoded data. Note: For byte fields, the contents of the field are base64-encoded (which increases the size of the data by 33-36%) when using JSON on the wire.
        /// </summary>
        [Output("icon")]
        public Output<string> Icon { get; private set; } = null!;

        /// <summary>
        /// Required. The ID of the listing to create. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces. Max length: 100 bytes.
        /// </summary>
        [Output("listingId")]
        public Output<string> ListingId { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name of the listing. e.g. `projects/myproject/locations/US/dataExchanges/123/listings/456`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. Email or URL of the primary point of contact of the listing. Max Length: 1000 bytes.
        /// </summary>
        [Output("primaryContact")]
        public Output<string> PrimaryContact { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. Details of the publisher who owns the listing and who can share the source data.
        /// </summary>
        [Output("publisher")]
        public Output<Outputs.PublisherResponse> Publisher { get; private set; } = null!;

        /// <summary>
        /// Optional. Email or URL of the request access of the listing. Subscribers can use this reference to request access. Max Length: 1000 bytes.
        /// </summary>
        [Output("requestAccess")]
        public Output<string> RequestAccess { get; private set; } = null!;

        /// <summary>
        /// Current state of the listing.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a Listing resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Listing(string name, ListingArgs args, CustomResourceOptions? options = null)
            : base("google-native:analyticshub/v1beta1:Listing", name, args ?? new ListingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Listing(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:analyticshub/v1beta1:Listing", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "dataExchangeId",
                    "listingId",
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
        /// Get an existing Listing resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Listing Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Listing(name, id, options);
        }
    }

    public sealed class ListingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Shared dataset i.e. BigQuery dataset source.
        /// </summary>
        [Input("bigqueryDataset", required: true)]
        public Input<Inputs.BigQueryDatasetSourceArgs> BigqueryDataset { get; set; } = null!;

        [Input("categories")]
        private InputList<Pulumi.GoogleNative.AnalyticsHub.V1Beta1.ListingCategoriesItem>? _categories;

        /// <summary>
        /// Optional. Categories of the listing. Up to two categories are allowed.
        /// </summary>
        public InputList<Pulumi.GoogleNative.AnalyticsHub.V1Beta1.ListingCategoriesItem> Categories
        {
            get => _categories ?? (_categories = new InputList<Pulumi.GoogleNative.AnalyticsHub.V1Beta1.ListingCategoriesItem>());
            set => _categories = value;
        }

        [Input("dataExchangeId", required: true)]
        public Input<string> DataExchangeId { get; set; } = null!;

        /// <summary>
        /// Optional. Details of the data provider who owns the source data.
        /// </summary>
        [Input("dataProvider")]
        public Input<Inputs.DataProviderArgs>? DataProvider { get; set; }

        /// <summary>
        /// Optional. Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF). Default value is an empty string. Max length: 2000 bytes.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&amp;) and can't start or end with spaces. Default value is an empty string. Max length: 63 bytes.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Optional. Documentation describing the listing.
        /// </summary>
        [Input("documentation")]
        public Input<string>? Documentation { get; set; }

        /// <summary>
        /// Optional. Base64 encoded image representing the listing. Max Size: 3.0MiB Expected image dimensions are 512x512 pixels, however the API only performs validation on size of the encoded data. Note: For byte fields, the contents of the field are base64-encoded (which increases the size of the data by 33-36%) when using JSON on the wire.
        /// </summary>
        [Input("icon")]
        public Input<string>? Icon { get; set; }

        /// <summary>
        /// Required. The ID of the listing to create. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces. Max length: 100 bytes.
        /// </summary>
        [Input("listingId", required: true)]
        public Input<string> ListingId { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Optional. Email or URL of the primary point of contact of the listing. Max Length: 1000 bytes.
        /// </summary>
        [Input("primaryContact")]
        public Input<string>? PrimaryContact { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. Details of the publisher who owns the listing and who can share the source data.
        /// </summary>
        [Input("publisher")]
        public Input<Inputs.PublisherArgs>? Publisher { get; set; }

        /// <summary>
        /// Optional. Email or URL of the request access of the listing. Subscribers can use this reference to request access. Max Length: 1000 bytes.
        /// </summary>
        [Input("requestAccess")]
        public Input<string>? RequestAccess { get; set; }

        public ListingArgs()
        {
        }
        public static new ListingArgs Empty => new ListingArgs();
    }
}
