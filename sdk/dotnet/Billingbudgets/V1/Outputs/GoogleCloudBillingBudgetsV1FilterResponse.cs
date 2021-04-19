// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Billingbudgets.V1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudBillingBudgetsV1FilterResponse
    {
        /// <summary>
        /// Optional. Specifies to track usage for recurring calendar period. For example, assume that CalendarPeriod.QUARTER is set. The budget will track usage from April 1 to June 30, when the current calendar month is April, May, June. After that, it will track usage from July 1 to September 30 when the current calendar month is July, August, September, so on.
        /// </summary>
        public readonly string CalendarPeriod;
        /// <summary>
        /// Optional. If Filter.credit_types_treatment is INCLUDE_SPECIFIED_CREDITS, this is a list of credit types to be subtracted from gross cost to determine the spend for threshold calculations. See [a list of acceptable credit type values](https://cloud.google.com/billing/docs/how-to/export-data-bigquery-tables#credits-type). If Filter.credit_types_treatment is **not** INCLUDE_SPECIFIED_CREDITS, this field must be empty.
        /// </summary>
        public readonly ImmutableArray<string> CreditTypes;
        /// <summary>
        /// Optional. If not set, default behavior is `INCLUDE_ALL_CREDITS`.
        /// </summary>
        public readonly string CreditTypesTreatment;
        /// <summary>
        /// Optional. Specifies to track usage from any start date (required) to any end date (optional). This time period is static, it does not recur.
        /// </summary>
        public readonly Outputs.GoogleCloudBillingBudgetsV1CustomPeriodResponse CustomPeriod;
        /// <summary>
        /// Optional. A single label and value pair specifying that usage from only this set of labeled resources should be included in the budget. Currently, multiple entries or multiple values per entry are not allowed. If omitted, the report will include all labeled and unlabeled usage.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Optional. A set of projects of the form `projects/{project}`, specifying that usage from only this set of projects should be included in the budget. If omitted, the report will include all usage for the billing account, regardless of which project the usage occurred on. Only zero or one project can be specified currently.
        /// </summary>
        public readonly ImmutableArray<string> Projects;
        /// <summary>
        /// Optional. A set of services of the form `services/{service_id}`, specifying that usage from only this set of services should be included in the budget. If omitted, the report will include usage for all the services. The service names are available through the Catalog API: https://cloud.google.com/billing/v1/how-tos/catalog-api.
        /// </summary>
        public readonly ImmutableArray<string> Services;
        /// <summary>
        /// Optional. A set of subaccounts of the form `billingAccounts/{account_id}`, specifying that usage from only this set of subaccounts should be included in the budget. If a subaccount is set to the name of the parent account, usage from the parent account will be included. If the field is omitted, the report will include usage from the parent account and all subaccounts, if they exist.
        /// </summary>
        public readonly ImmutableArray<string> Subaccounts;

        [OutputConstructor]
        private GoogleCloudBillingBudgetsV1FilterResponse(
            string calendarPeriod,

            ImmutableArray<string> creditTypes,

            string creditTypesTreatment,

            Outputs.GoogleCloudBillingBudgetsV1CustomPeriodResponse customPeriod,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<string> projects,

            ImmutableArray<string> services,

            ImmutableArray<string> subaccounts)
        {
            CalendarPeriod = calendarPeriod;
            CreditTypes = creditTypes;
            CreditTypesTreatment = creditTypesTreatment;
            CustomPeriod = customPeriod;
            Labels = labels;
            Projects = projects;
            Services = services;
            Subaccounts = subaccounts;
        }
    }
}
