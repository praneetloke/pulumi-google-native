// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Pubsub.V1
{
    public static class GetSubscription
    {
        /// <summary>
        /// Gets the configuration details of a subscription.
        /// </summary>
        public static Task<GetSubscriptionResult> InvokeAsync(GetSubscriptionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubscriptionResult>("google-native:pubsub/v1:getSubscription", args ?? new GetSubscriptionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the configuration details of a subscription.
        /// </summary>
        public static Output<GetSubscriptionResult> Invoke(GetSubscriptionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSubscriptionResult>("google-native:pubsub/v1:getSubscription", args ?? new GetSubscriptionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSubscriptionArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("subscriptionId", required: true)]
        public string SubscriptionId { get; set; } = null!;

        public GetSubscriptionArgs()
        {
        }
        public static new GetSubscriptionArgs Empty => new GetSubscriptionArgs();
    }

    public sealed class GetSubscriptionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("subscriptionId", required: true)]
        public Input<string> SubscriptionId { get; set; } = null!;

        public GetSubscriptionInvokeArgs()
        {
        }
        public static new GetSubscriptionInvokeArgs Empty => new GetSubscriptionInvokeArgs();
    }


    [OutputType]
    public sealed class GetSubscriptionResult
    {
        /// <summary>
        /// The approximate amount of time (on a best-effort basis) Pub/Sub waits for the subscriber to acknowledge receipt before resending the message. In the interval after the message is delivered and before it is acknowledged, it is considered to be *outstanding*. During that time period, the message will not be redelivered (on a best-effort basis). For pull subscriptions, this value is used as the initial value for the ack deadline. To override this value for a given message, call `ModifyAckDeadline` with the corresponding `ack_id` if using non-streaming pull or send the `ack_id` in a `StreamingModifyAckDeadlineRequest` if using streaming pull. The minimum custom deadline you can specify is 10 seconds. The maximum custom deadline you can specify is 600 seconds (10 minutes). If this parameter is 0, a default value of 10 seconds is used. For push delivery, this value is also used to set the request timeout for the call to the push endpoint. If the subscriber never acknowledges the message, the Pub/Sub system will eventually redeliver the message.
        /// </summary>
        public readonly int AckDeadlineSeconds;
        /// <summary>
        /// If delivery to BigQuery is used with this subscription, this field is used to configure it. Either `pushConfig` or `bigQueryConfig` can be set, but not both. If both are empty, then the subscriber will pull and ack messages using API methods.
        /// </summary>
        public readonly Outputs.BigQueryConfigResponse BigqueryConfig;
        /// <summary>
        /// A policy that specifies the conditions for dead lettering messages in this subscription. If dead_letter_policy is not set, dead lettering is disabled. The Cloud Pub/Sub service account associated with this subscriptions's parent project (i.e., service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have permission to Acknowledge() messages on this subscription.
        /// </summary>
        public readonly Outputs.DeadLetterPolicyResponse DeadLetterPolicy;
        /// <summary>
        /// Indicates whether the subscription is detached from its topic. Detached subscriptions don't receive messages from their topic and don't retain any backlog. `Pull` and `StreamingPull` requests will return FAILED_PRECONDITION. If the subscription is a push subscription, pushes to the endpoint will not be made.
        /// </summary>
        public readonly bool Detached;
        /// <summary>
        /// If true, Pub/Sub provides the following guarantees for the delivery of a message with a given value of `message_id` on this subscription: * The message sent to a subscriber is guaranteed not to be resent before the message's acknowledgement deadline expires. * An acknowledged message will not be resent to a subscriber. Note that subscribers may still receive multiple copies of a message when `enable_exactly_once_delivery` is true if the message was published multiple times by a publisher client. These copies are considered distinct by Pub/Sub and have distinct `message_id` values.
        /// </summary>
        public readonly bool EnableExactlyOnceDelivery;
        /// <summary>
        /// If true, messages published with the same `ordering_key` in `PubsubMessage` will be delivered to the subscribers in the order in which they are received by the Pub/Sub system. Otherwise, they may be delivered in any order.
        /// </summary>
        public readonly bool EnableMessageOrdering;
        /// <summary>
        /// A policy that specifies the conditions for this subscription's expiration. A subscription is considered active as long as any connected subscriber is successfully consuming messages from the subscription or is issuing operations on the subscription. If `expiration_policy` is not set, a *default policy* with `ttl` of 31 days will be used. The minimum allowed value for `expiration_policy.ttl` is 1 day. If `expiration_policy` is set, but `expiration_policy.ttl` is not set, the subscription never expires.
        /// </summary>
        public readonly Outputs.ExpirationPolicyResponse ExpirationPolicy;
        /// <summary>
        /// An expression written in the Pub/Sub [filter language](https://cloud.google.com/pubsub/docs/filtering). If non-empty, then only `PubsubMessage`s whose `attributes` field matches the filter are delivered on this subscription. If empty, then no messages are filtered out.
        /// </summary>
        public readonly string Filter;
        /// <summary>
        /// See Creating and managing labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// How long to retain unacknowledged messages in the subscription's backlog, from the moment a message is published. If `retain_acked_messages` is true, then this also configures the retention of acknowledged messages, and thus configures how far back in time a `Seek` can be done. Defaults to 7 days. Cannot be more than 7 days or less than 10 minutes.
        /// </summary>
        public readonly string MessageRetentionDuration;
        /// <summary>
        /// The name of the subscription. It must have the format `"projects/{project}/subscriptions/{subscription}"`. `{subscription}` must start with a letter, and contain only letters (`[A-Za-z]`), numbers (`[0-9]`), dashes (`-`), underscores (`_`), periods (`.`), tildes (`~`), plus (`+`) or percent signs (`%`). It must be between 3 and 255 characters in length, and it must not start with `"goog"`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// If push delivery is used with this subscription, this field is used to configure it. Either `pushConfig` or `bigQueryConfig` can be set, but not both. If both are empty, then the subscriber will pull and ack messages using API methods.
        /// </summary>
        public readonly Outputs.PushConfigResponse PushConfig;
        /// <summary>
        /// Indicates whether to retain acknowledged messages. If true, then messages are not expunged from the subscription's backlog, even if they are acknowledged, until they fall out of the `message_retention_duration` window. This must be true if you would like to [`Seek` to a timestamp] (https://cloud.google.com/pubsub/docs/replay-overview#seek_to_a_time) in the past to replay previously-acknowledged messages.
        /// </summary>
        public readonly bool RetainAckedMessages;
        /// <summary>
        /// A policy that specifies how Pub/Sub retries message delivery for this subscription. If not set, the default retry policy is applied. This generally implies that messages will be retried as soon as possible for healthy subscribers. RetryPolicy will be triggered on NACKs or acknowledgement deadline exceeded events for a given message.
        /// </summary>
        public readonly Outputs.RetryPolicyResponse RetryPolicy;
        /// <summary>
        /// An output-only field indicating whether or not the subscription can receive messages.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The name of the topic from which this subscription is receiving messages. Format is `projects/{project}/topics/{topic}`. The value of this field will be `_deleted-topic_` if the topic has been deleted.
        /// </summary>
        public readonly string Topic;
        /// <summary>
        /// Indicates the minimum duration for which a message is retained after it is published to the subscription's topic. If this field is set, messages published to the subscription's topic in the last `topic_message_retention_duration` are always available to subscribers. See the `message_retention_duration` field in `Topic`. This field is set only in responses from the server; it is ignored if it is set in any requests.
        /// </summary>
        public readonly string TopicMessageRetentionDuration;

        [OutputConstructor]
        private GetSubscriptionResult(
            int ackDeadlineSeconds,

            Outputs.BigQueryConfigResponse bigqueryConfig,

            Outputs.DeadLetterPolicyResponse deadLetterPolicy,

            bool detached,

            bool enableExactlyOnceDelivery,

            bool enableMessageOrdering,

            Outputs.ExpirationPolicyResponse expirationPolicy,

            string filter,

            ImmutableDictionary<string, string> labels,

            string messageRetentionDuration,

            string name,

            Outputs.PushConfigResponse pushConfig,

            bool retainAckedMessages,

            Outputs.RetryPolicyResponse retryPolicy,

            string state,

            string topic,

            string topicMessageRetentionDuration)
        {
            AckDeadlineSeconds = ackDeadlineSeconds;
            BigqueryConfig = bigqueryConfig;
            DeadLetterPolicy = deadLetterPolicy;
            Detached = detached;
            EnableExactlyOnceDelivery = enableExactlyOnceDelivery;
            EnableMessageOrdering = enableMessageOrdering;
            ExpirationPolicy = expirationPolicy;
            Filter = filter;
            Labels = labels;
            MessageRetentionDuration = messageRetentionDuration;
            Name = name;
            PushConfig = pushConfig;
            RetainAckedMessages = retainAckedMessages;
            RetryPolicy = retryPolicy;
            State = state;
            Topic = topic;
            TopicMessageRetentionDuration = topicMessageRetentionDuration;
        }
    }
}
