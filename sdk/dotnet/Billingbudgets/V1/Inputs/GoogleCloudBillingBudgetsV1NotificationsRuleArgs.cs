// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Billingbudgets.V1.Inputs
{

    /// <summary>
    /// NotificationsRule defines notifications that are sent based on budget spend and thresholds.
    /// </summary>
    public sealed class GoogleCloudBillingBudgetsV1NotificationsRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. When set to true, disables default notifications sent when a threshold is exceeded. Default notifications are sent to those with Billing Account Administrator and Billing Account User IAM roles for the target account.
        /// </summary>
        [Input("disableDefaultIamRecipients")]
        public Input<bool>? DisableDefaultIamRecipients { get; set; }

        [Input("monitoringNotificationChannels")]
        private InputList<string>? _monitoringNotificationChannels;

        /// <summary>
        /// Optional. Email targets to send notifications to when a threshold is exceeded. This is in addition to the `DefaultIamRecipients` who receive alert emails based on their billing account IAM role. The value is the full REST resource name of a Cloud Monitoring email notification channel with the form `projects/{project_id}/notificationChannels/{channel_id}`. A maximum of 5 email notifications are allowed. To customize budget alert email recipients with monitoring notification channels, you _must create the monitoring notification channels before you link them to a budget_. For guidance on setting up notification channels to use with budgets, see [Customize budget alert email recipients](https://cloud.google.com/billing/docs/how-to/budgets-notification-recipients). For Cloud Billing budget alerts, you _must use email notification channels_. The other types of notification channels are _not_ supported, such as Slack, SMS, or PagerDuty. If you want to [send budget notifications to Slack](https://cloud.google.com/billing/docs/how-to/notify#send_notifications_to_slack), use a pubsubTopic and configure [programmatic notifications](https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications).
        /// </summary>
        public InputList<string> MonitoringNotificationChannels
        {
            get => _monitoringNotificationChannels ?? (_monitoringNotificationChannels = new InputList<string>());
            set => _monitoringNotificationChannels = value;
        }

        /// <summary>
        /// Optional. The name of the Pub/Sub topic where budget-related messages are published, in the form `projects/{project_id}/topics/{topic_id}`. Updates are sent to the topic at regular intervals; the timing of the updates is not dependent on the [threshold rules](#thresholdrule) you've set. Note that if you want your [Pub/Sub JSON object](https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#notification_format) to contain data for `alertThresholdExceeded`, you need at least one [alert threshold rule](#thresholdrule). When you set threshold rules, you must also enable at least one of the email notification options, either using the default IAM recipients or Cloud Monitoring email notification channels. To use Pub/Sub topics with budgets, you must do the following: 1. Create the Pub/Sub topic before connecting it to your budget. For guidance, see [Manage programmatic budget alert notifications](https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications). 2. Grant the API caller the `pubsub.topics.setIamPolicy` permission on the Pub/Sub topic. If not set, the API call fails with PERMISSION_DENIED. For additional details on Pub/Sub roles and permissions, see [Permissions required for this task](https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#permissions_required_for_this_task).
        /// </summary>
        [Input("pubsubTopic")]
        public Input<string>? PubsubTopic { get; set; }

        /// <summary>
        /// Optional. Required when NotificationsRule.pubsub_topic is set. The schema version of the notification sent to NotificationsRule.pubsub_topic. Only "1.0" is accepted. It represents the JSON schema as defined in https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#notification_format.
        /// </summary>
        [Input("schemaVersion")]
        public Input<string>? SchemaVersion { get; set; }

        public GoogleCloudBillingBudgetsV1NotificationsRuleArgs()
        {
        }
        public static new GoogleCloudBillingBudgetsV1NotificationsRuleArgs Empty => new GoogleCloudBillingBudgetsV1NotificationsRuleArgs();
    }
}
