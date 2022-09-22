// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Pubsub.V1Beta1a
{
    public static class GetSubscription
    {
        /// <summary>
        /// Gets the configuration details of a subscription.
        /// </summary>
        public static Task<GetSubscriptionResult> InvokeAsync(GetSubscriptionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSubscriptionResult>("google-native:pubsub/v1beta1a:getSubscription", args ?? new GetSubscriptionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the configuration details of a subscription.
        /// </summary>
        public static Output<GetSubscriptionResult> Invoke(GetSubscriptionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSubscriptionResult>("google-native:pubsub/v1beta1a:getSubscription", args ?? new GetSubscriptionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSubscriptionArgs : global::Pulumi.InvokeArgs
    {
        [Input("subscriptionId", required: true)]
        public string SubscriptionId { get; set; } = null!;

        public GetSubscriptionArgs()
        {
        }
        public static new GetSubscriptionArgs Empty => new GetSubscriptionArgs();
    }

    public sealed class GetSubscriptionInvokeArgs : global::Pulumi.InvokeArgs
    {
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
        /// For either push or pull delivery, the value is the maximum time after a subscriber receives a message before the subscriber should acknowledge or Nack the message. If the Ack deadline for a message passes without an Ack or a Nack, the Pub/Sub system will eventually redeliver the message. If a subscriber acknowledges after the deadline, the Pub/Sub system may accept the Ack, but it is possible that the message has been already delivered again. Multiple Acks to the message are allowed and will succeed. For push delivery, this value is used to set the request timeout for the call to the push endpoint. For pull delivery, this value is used as the initial value for the Ack deadline. It may be overridden for each message using its corresponding ack_id with ModifyAckDeadline. While a message is outstanding (i.e. it has been delivered to a pull subscriber and the subscriber has not yet Acked or Nacked), the Pub/Sub system will not deliver that message to another pull subscriber (on a best-effort basis).
        /// </summary>
        public readonly int AckDeadlineSeconds;
        /// <summary>
        /// Name of the subscription.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// If push delivery is used with this subscription, this field is used to configure it.
        /// </summary>
        public readonly Outputs.PushConfigResponse PushConfig;
        /// <summary>
        /// The name of the topic from which this subscription is receiving messages.
        /// </summary>
        public readonly string Topic;

        [OutputConstructor]
        private GetSubscriptionResult(
            int ackDeadlineSeconds,

            string name,

            Outputs.PushConfigResponse pushConfig,

            string topic)
        {
            AckDeadlineSeconds = ackDeadlineSeconds;
            Name = name;
            PushConfig = pushConfig;
            Topic = topic;
        }
    }
}
