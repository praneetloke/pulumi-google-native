// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V1
{
    public static class GetDashboard
    {
        /// <summary>
        /// Fetches a specific dashboard.This method requires the monitoring.dashboards.get permission on the specified dashboard. For more information, see Cloud Identity and Access Management (https://cloud.google.com/iam).
        /// </summary>
        public static Task<GetDashboardResult> InvokeAsync(GetDashboardArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDashboardResult>("google-native:monitoring/v1:getDashboard", args ?? new GetDashboardArgs(), options.WithDefaults());

        /// <summary>
        /// Fetches a specific dashboard.This method requires the monitoring.dashboards.get permission on the specified dashboard. For more information, see Cloud Identity and Access Management (https://cloud.google.com/iam).
        /// </summary>
        public static Output<GetDashboardResult> Invoke(GetDashboardInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDashboardResult>("google-native:monitoring/v1:getDashboard", args ?? new GetDashboardInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDashboardArgs : Pulumi.InvokeArgs
    {
        [Input("dashboardId", required: true)]
        public string DashboardId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetDashboardArgs()
        {
        }
    }

    public sealed class GetDashboardInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("dashboardId", required: true)]
        public Input<string> DashboardId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetDashboardInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDashboardResult
    {
        /// <summary>
        /// The content is divided into equally spaced columns and the widgets are arranged vertically.
        /// </summary>
        public readonly Outputs.ColumnLayoutResponse ColumnLayout;
        /// <summary>
        /// Filters to reduce the amount of data charted based on the filter criteria.
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardFilterResponse> DashboardFilters;
        /// <summary>
        /// The mutable, human-readable name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. An etag is returned in the response to GetDashboard, and users are expected to put that etag in the request to UpdateDashboard to ensure that their change will be applied to the same version of the Dashboard configuration. The field should not be passed during dashboard creation.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Content is arranged with a basic layout that re-flows a simple list of informational elements like widgets or tiles.
        /// </summary>
        public readonly Outputs.GridLayoutResponse GridLayout;
        /// <summary>
        /// Labels applied to the dashboard
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The content is arranged as a grid of tiles, with each content widget occupying one or more grid blocks.
        /// </summary>
        public readonly Outputs.MosaicLayoutResponse MosaicLayout;
        /// <summary>
        /// Immutable. The resource name of the dashboard.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The content is divided into equally spaced rows and the widgets are arranged horizontally.
        /// </summary>
        public readonly Outputs.RowLayoutResponse RowLayout;

        [OutputConstructor]
        private GetDashboardResult(
            Outputs.ColumnLayoutResponse columnLayout,

            ImmutableArray<Outputs.DashboardFilterResponse> dashboardFilters,

            string displayName,

            string etag,

            Outputs.GridLayoutResponse gridLayout,

            ImmutableDictionary<string, string> labels,

            Outputs.MosaicLayoutResponse mosaicLayout,

            string name,

            Outputs.RowLayoutResponse rowLayout)
        {
            ColumnLayout = columnLayout;
            DashboardFilters = dashboardFilters;
            DisplayName = displayName;
            Etag = etag;
            GridLayout = gridLayout;
            Labels = labels;
            MosaicLayout = mosaicLayout;
            Name = name;
            RowLayout = rowLayout;
        }
    }
}
