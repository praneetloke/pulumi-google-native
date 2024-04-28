// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Billingbudgets.V1Beta1.Inputs
{

    /// <summary>
    /// A filter for a budget, limiting the scope of the cost to calculate.
    /// </summary>
    public sealed class GoogleCloudBillingBudgetsV1beta1FilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Specifies to track usage for recurring calendar period. For example, assume that CalendarPeriod.QUARTER is set. The budget will track usage from April 1 to June 30, when the current calendar month is April, May, June. After that, it will track usage from July 1 to September 30 when the current calendar month is July, August, September, so on.
        /// </summary>
        [Input("calendarPeriod")]
        public Input<Pulumi.GoogleNative.Billingbudgets.V1Beta1.GoogleCloudBillingBudgetsV1beta1FilterCalendarPeriod>? CalendarPeriod { get; set; }

        [Input("creditTypes")]
        private InputList<string>? _creditTypes;

        /// <summary>
        /// Optional. If Filter.credit_types_treatment is INCLUDE_SPECIFIED_CREDITS, this is a list of credit types to be subtracted from gross cost to determine the spend for threshold calculations. See [a list of acceptable credit type values](https://cloud.google.com/billing/docs/how-to/export-data-bigquery-tables#credits-type). If Filter.credit_types_treatment is **not** INCLUDE_SPECIFIED_CREDITS, this field must be empty.
        /// </summary>
        public InputList<string> CreditTypes
        {
            get => _creditTypes ?? (_creditTypes = new InputList<string>());
            set => _creditTypes = value;
        }

        /// <summary>
        /// Optional. If not set, default behavior is `INCLUDE_ALL_CREDITS`.
        /// </summary>
        [Input("creditTypesTreatment")]
        public Input<Pulumi.GoogleNative.Billingbudgets.V1Beta1.GoogleCloudBillingBudgetsV1beta1FilterCreditTypesTreatment>? CreditTypesTreatment { get; set; }

        /// <summary>
        /// Optional. Specifies to track usage from any start date (required) to any end date (optional). This time period is static, it does not recur.
        /// </summary>
        [Input("customPeriod")]
        public Input<Inputs.GoogleCloudBillingBudgetsV1beta1CustomPeriodArgs>? CustomPeriod { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// Optional. A single label and value pair specifying that usage from only this set of labeled resources should be included in the budget. If omitted, the report will include all labeled and unlabeled usage. An object containing a single `"key": value` pair. Example: `{ "name": "wrench" }`. _Currently, multiple entries or multiple values per entry are not allowed._
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        [Input("projects")]
        private InputList<string>? _projects;

        /// <summary>
        /// Optional. A set of projects of the form `projects/{project}`, specifying that usage from only this set of projects should be included in the budget. If omitted, the report will include all usage for the billing account, regardless of which project the usage occurred on.
        /// </summary>
        public InputList<string> Projects
        {
            get => _projects ?? (_projects = new InputList<string>());
            set => _projects = value;
        }

        [Input("resourceAncestors")]
        private InputList<string>? _resourceAncestors;

        /// <summary>
        /// Optional. A set of folder and organization names of the form `folders/{folderId}` or `organizations/{organizationId}`, specifying that usage from only this set of folders and organizations should be included in the budget. If omitted, the budget includes all usage that the billing account pays for. If the folder or organization contains projects that are paid for by a different Cloud Billing account, the budget *doesn't* apply to those projects.
        /// </summary>
        public InputList<string> ResourceAncestors
        {
            get => _resourceAncestors ?? (_resourceAncestors = new InputList<string>());
            set => _resourceAncestors = value;
        }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// Optional. A set of services of the form `services/{service_id}`, specifying that usage from only this set of services should be included in the budget. If omitted, the report will include usage for all the services. The service names are available through the Catalog API: https://cloud.google.com/billing/v1/how-tos/catalog-api.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        [Input("subaccounts")]
        private InputList<string>? _subaccounts;

        /// <summary>
        /// Optional. A set of subaccounts of the form `billingAccounts/{account_id}`, specifying that usage from only this set of subaccounts should be included in the budget. If a subaccount is set to the name of the parent account, usage from the parent account will be included. If omitted, the report will include usage from the parent account and all subaccounts, if they exist.
        /// </summary>
        public InputList<string> Subaccounts
        {
            get => _subaccounts ?? (_subaccounts = new InputList<string>());
            set => _subaccounts = value;
        }

        public GoogleCloudBillingBudgetsV1beta1FilterArgs()
        {
        }
        public static new GoogleCloudBillingBudgetsV1beta1FilterArgs Empty => new GoogleCloudBillingBudgetsV1beta1FilterArgs();
    }
}
