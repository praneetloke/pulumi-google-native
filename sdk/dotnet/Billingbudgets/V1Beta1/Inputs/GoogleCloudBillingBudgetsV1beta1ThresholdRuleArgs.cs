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
    /// ThresholdRule contains the definition of a threshold. Threshold rules define the triggering events used to generate a budget notification email. When a threshold is crossed (spend exceeds the specified percentages of the budget), budget alert emails are sent to the email recipients you specify in the [NotificationsRule](#notificationsrule). Threshold rules also affect the fields included in the [JSON data object](https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#notification_format) sent to a Pub/Sub topic. Threshold rules are _required_ if using email notifications. Threshold rules are _optional_ if only setting a [`pubsubTopic` NotificationsRule](#NotificationsRule), unless you want your JSON data object to include data about the thresholds you set. For more information, see [set budget threshold rules and actions](https://cloud.google.com/billing/docs/how-to/budgets#budget-actions).
    /// </summary>
    public sealed class GoogleCloudBillingBudgetsV1beta1ThresholdRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The type of basis used to determine if spend has passed the threshold. Behavior defaults to CURRENT_SPEND if not set.
        /// </summary>
        [Input("spendBasis")]
        public Input<Pulumi.GoogleNative.Billingbudgets.V1Beta1.GoogleCloudBillingBudgetsV1beta1ThresholdRuleSpendBasis>? SpendBasis { get; set; }

        /// <summary>
        /// Send an alert when this threshold is exceeded. This is a 1.0-based percentage, so 0.5 = 50%. Validation: non-negative number.
        /// </summary>
        [Input("thresholdPercent", required: true)]
        public Input<double> ThresholdPercent { get; set; } = null!;

        public GoogleCloudBillingBudgetsV1beta1ThresholdRuleArgs()
        {
        }
        public static new GoogleCloudBillingBudgetsV1beta1ThresholdRuleArgs Empty => new GoogleCloudBillingBudgetsV1beta1ThresholdRuleArgs();
    }
}
