// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Pubsub.V1Beta2
{
    public static class GetTopic
    {
        /// <summary>
        /// Gets the configuration of a topic.
        /// </summary>
        public static Task<GetTopicResult> InvokeAsync(GetTopicArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTopicResult>("google-native:pubsub/v1beta2:getTopic", args ?? new GetTopicArgs(), options.WithVersion());
    }


    public sealed class GetTopicArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("topicId", required: true)]
        public string TopicId { get; set; } = null!;

        public GetTopicArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTopicResult
    {
        /// <summary>
        /// The name of the topic. It must have the format `"projects/{project}/topics/{topic}"`. `{topic}` must start with a letter, and contain only letters (`[A-Za-z]`), numbers (`[0-9]`), dashes (`-`), underscores (`_`), periods (`.`), tildes (`~`), plus (`+`) or percent signs (`%`). It must be between 3 and 255 characters in length, and it must not start with `"goog"`.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetTopicResult(string name)
        {
            Name = name;
        }
    }
}