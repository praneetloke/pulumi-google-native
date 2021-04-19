// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Billingbudgets.V1Beta1
{
    /// <summary>
    /// Creates a new budget. See Quotas and limits for more information on the limits of the number of budgets you can create.
    /// </summary>
    [GoogleNativeResourceType("google-native:billingbudgets/v1beta1:BillingAccountBudget")]
    public partial class BillingAccountBudget : Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. Rules to apply to notifications sent based on budget spend and thresholds.
        /// </summary>
        [Output("allUpdatesRule")]
        public Output<Outputs.GoogleCloudBillingBudgetsV1beta1AllUpdatesRuleResponse> AllUpdatesRule { get; private set; } = null!;

        /// <summary>
        /// Required. Budgeted amount.
        /// </summary>
        [Output("amount")]
        public Output<Outputs.GoogleCloudBillingBudgetsV1beta1BudgetAmountResponse> Amount { get; private set; } = null!;

        /// <summary>
        /// Optional. Filters that define which resources are used to compute the actual spend against the budget amount, such as projects, services, and the budget's time period, as well as other filters.
        /// </summary>
        [Output("budgetFilter")]
        public Output<Outputs.GoogleCloudBillingBudgetsV1beta1FilterResponse> BudgetFilter { get; private set; } = null!;

        /// <summary>
        /// User data for display name in UI. Validation: &lt;= 60 chars.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Optional. Etag to validate that the object is unchanged for a read-modify-write operation. An empty etag will cause an update to overwrite other changes.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource name of the budget. The resource name implies the scope of a budget. Values are of the form `billingAccounts/{billingAccountId}/budgets/{budgetId}`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of the budget.
        /// </summary>
        [Output("thresholdRules")]
        public Output<ImmutableArray<Outputs.GoogleCloudBillingBudgetsV1beta1ThresholdRuleResponse>> ThresholdRules { get; private set; } = null!;


        /// <summary>
        /// Create a BillingAccountBudget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BillingAccountBudget(string name, BillingAccountBudgetArgs args, CustomResourceOptions? options = null)
            : base("google-native:billingbudgets/v1beta1:BillingAccountBudget", name, args ?? new BillingAccountBudgetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BillingAccountBudget(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:billingbudgets/v1beta1:BillingAccountBudget", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing BillingAccountBudget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BillingAccountBudget Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BillingAccountBudget(name, id, options);
        }
    }

    public sealed class BillingAccountBudgetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Rules to apply to notifications sent based on budget spend and thresholds.
        /// </summary>
        [Input("allUpdatesRule")]
        public Input<Inputs.GoogleCloudBillingBudgetsV1beta1AllUpdatesRuleArgs>? AllUpdatesRule { get; set; }

        /// <summary>
        /// Required. Budgeted amount.
        /// </summary>
        [Input("amount")]
        public Input<Inputs.GoogleCloudBillingBudgetsV1beta1BudgetAmountArgs>? Amount { get; set; }

        [Input("billingAccountsId", required: true)]
        public Input<string> BillingAccountsId { get; set; } = null!;

        /// <summary>
        /// Optional. Filters that define which resources are used to compute the actual spend against the budget amount, such as projects, services, and the budget's time period, as well as other filters.
        /// </summary>
        [Input("budgetFilter")]
        public Input<Inputs.GoogleCloudBillingBudgetsV1beta1FilterArgs>? BudgetFilter { get; set; }

        [Input("budgetsId", required: true)]
        public Input<string> BudgetsId { get; set; } = null!;

        /// <summary>
        /// User data for display name in UI. Validation: &lt;= 60 chars.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Optional. Etag to validate that the object is unchanged for a read-modify-write operation. An empty etag will cause an update to overwrite other changes.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("thresholdRules")]
        private InputList<Inputs.GoogleCloudBillingBudgetsV1beta1ThresholdRuleArgs>? _thresholdRules;

        /// <summary>
        /// Optional. Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of the budget.
        /// </summary>
        public InputList<Inputs.GoogleCloudBillingBudgetsV1beta1ThresholdRuleArgs> ThresholdRules
        {
            get => _thresholdRules ?? (_thresholdRules = new InputList<Inputs.GoogleCloudBillingBudgetsV1beta1ThresholdRuleArgs>());
            set => _thresholdRules = value;
        }

        public BillingAccountBudgetArgs()
        {
        }
    }
}
