// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1
{
    public static class GetEntryGroup
    {
        /// <summary>
        /// Gets an entry group.
        /// </summary>
        public static Task<GetEntryGroupResult> InvokeAsync(GetEntryGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEntryGroupResult>("google-native:datacatalog/v1:getEntryGroup", args ?? new GetEntryGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an entry group.
        /// </summary>
        public static Output<GetEntryGroupResult> Invoke(GetEntryGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEntryGroupResult>("google-native:datacatalog/v1:getEntryGroup", args ?? new GetEntryGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEntryGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("entryGroupId", required: true)]
        public string EntryGroupId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("readMask")]
        public string? ReadMask { get; set; }

        public GetEntryGroupArgs()
        {
        }
        public static new GetEntryGroupArgs Empty => new GetEntryGroupArgs();
    }

    public sealed class GetEntryGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("entryGroupId", required: true)]
        public Input<string> EntryGroupId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("readMask")]
        public Input<string>? ReadMask { get; set; }

        public GetEntryGroupInvokeArgs()
        {
        }
        public static new GetEntryGroupInvokeArgs Empty => new GetEntryGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetEntryGroupResult
    {
        /// <summary>
        /// Timestamps of the entry group. Default value is empty.
        /// </summary>
        public readonly Outputs.GoogleCloudDatacatalogV1SystemTimestampsResponse DataCatalogTimestamps;
        /// <summary>
        /// Entry group description. Can consist of several sentences or paragraphs that describe the entry group contents. Default value is an empty string.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// A short name to identify the entry group, for example, "analytics data - jan 2011". Default value is an empty string.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The resource name of the entry group in URL format. Note: The entry group itself and its child resources might not be stored in the location specified in its name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetEntryGroupResult(
            Outputs.GoogleCloudDatacatalogV1SystemTimestampsResponse dataCatalogTimestamps,

            string description,

            string displayName,

            string name)
        {
            DataCatalogTimestamps = dataCatalogTimestamps;
            Description = description;
            DisplayName = displayName;
            Name = name;
        }
    }
}
