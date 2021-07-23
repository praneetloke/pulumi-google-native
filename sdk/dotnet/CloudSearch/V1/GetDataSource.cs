// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudSearch.V1
{
    public static class GetDataSource
    {
        /// <summary>
        /// Gets a datasource. **Note:** This API requires an admin account to execute.
        /// </summary>
        public static Task<GetDataSourceResult> InvokeAsync(GetDataSourceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDataSourceResult>("google-native:cloudsearch/v1:getDataSource", args ?? new GetDataSourceArgs(), options.WithVersion());
    }


    public sealed class GetDataSourceArgs : Pulumi.InvokeArgs
    {
        [Input("datasourceId", required: true)]
        public string DatasourceId { get; set; } = null!;

        [Input("debugOptionsEnableDebugging")]
        public string? DebugOptionsEnableDebugging { get; set; }

        public GetDataSourceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDataSourceResult
    {
        /// <summary>
        /// If true, sets the datasource to read-only mode. In read-only mode, the Indexing API rejects any requests to index or delete items in this source. Enabling read-only mode does not stop the processing of previously accepted data.
        /// </summary>
        public readonly bool DisableModifications;
        /// <summary>
        /// Disable serving any search or assist results.
        /// </summary>
        public readonly bool DisableServing;
        /// <summary>
        /// Display name of the datasource The maximum length is 300 characters.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// List of service accounts that have indexing access.
        /// </summary>
        public readonly ImmutableArray<string> IndexingServiceAccounts;
        /// <summary>
        /// This field restricts visibility to items at the datasource level. Items within the datasource are restricted to the union of users and groups included in this field. Note that, this does not ensure access to a specific item, as users need to have ACL permissions on the contained items. This ensures a high level access on the entire datasource, and that the individual items are not shared outside this visibility.
        /// </summary>
        public readonly ImmutableArray<Outputs.GSuitePrincipalResponse> ItemsVisibility;
        /// <summary>
        /// Name of the datasource resource. Format: datasources/{source_id}. The name is ignored when creating a datasource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// IDs of the Long Running Operations (LROs) currently running for this schema.
        /// </summary>
        public readonly ImmutableArray<string> OperationIds;
        /// <summary>
        /// A short name or alias for the source. This value will be used to match the 'source' operator. For example, if the short name is *&lt;value&gt;* then queries like *source:&lt;value&gt;* will only return results for this source. The value must be unique across all datasources. The value must only contain alphanumeric characters (a-zA-Z0-9). The value cannot start with 'google' and cannot be one of the following: mail, gmail, docs, drive, groups, sites, calendar, hangouts, gplus, keep, people, teams. Its maximum length is 32 characters.
        /// </summary>
        public readonly string ShortName;

        [OutputConstructor]
        private GetDataSourceResult(
            bool disableModifications,

            bool disableServing,

            string displayName,

            ImmutableArray<string> indexingServiceAccounts,

            ImmutableArray<Outputs.GSuitePrincipalResponse> itemsVisibility,

            string name,

            ImmutableArray<string> operationIds,

            string shortName)
        {
            DisableModifications = disableModifications;
            DisableServing = disableServing;
            DisplayName = displayName;
            IndexingServiceAccounts = indexingServiceAccounts;
            ItemsVisibility = itemsVisibility;
            Name = name;
            OperationIds = operationIds;
            ShortName = shortName;
        }
    }
}