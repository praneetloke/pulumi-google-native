// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SecretManager.V1
{
    public static class GetSecret
    {
        /// <summary>
        /// Gets metadata for a given Secret.
        /// </summary>
        public static Task<GetSecretResult> InvokeAsync(GetSecretArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecretResult>("google-native:secretmanager/v1:getSecret", args ?? new GetSecretArgs(), options.WithDefaults());

        /// <summary>
        /// Gets metadata for a given Secret.
        /// </summary>
        public static Output<GetSecretResult> Invoke(GetSecretInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSecretResult>("google-native:secretmanager/v1:getSecret", args ?? new GetSecretInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecretArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("secretId", required: true)]
        public string SecretId { get; set; } = null!;

        public GetSecretArgs()
        {
        }
    }

    public sealed class GetSecretInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("secretId", required: true)]
        public Input<string> SecretId { get; set; } = null!;

        public GetSecretInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSecretResult
    {
        /// <summary>
        /// The time at which the Secret was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. Etag of the currently stored Secret.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Optional. Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
        /// </summary>
        public readonly string ExpireTime;
        /// <summary>
        /// The labels assigned to this Secret. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: `\p{Ll}\p{Lo}{0,62}` Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}` No more than 64 labels can be assigned to a given resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The resource name of the Secret in the format `projects/*/secrets/*`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Immutable. The replication policy of the secret data attached to the Secret. The replication policy cannot be changed after the Secret has been created.
        /// </summary>
        public readonly Outputs.ReplicationResponse Replication;
        /// <summary>
        /// Optional. Rotation policy attached to the Secret. May be excluded if there is no rotation policy.
        /// </summary>
        public readonly Outputs.RotationResponse Rotation;
        /// <summary>
        /// Optional. A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
        /// </summary>
        public readonly ImmutableArray<Outputs.TopicResponse> Topics;
        /// <summary>
        /// Input only. The TTL for the Secret.
        /// </summary>
        public readonly string Ttl;
        /// <summary>
        /// Optional. Mapping from version alias to version name. A version alias is a string with a maximum length of 63 characters and can contain uppercase and lowercase letters, numerals, and the hyphen (`-`) and underscore ('_') characters. An alias string must start with a letter and cannot be the string 'latest' or 'NEW'. No more than 50 aliases can be assigned to a given secret. Version-Alias pairs will be viewable via GetSecret and modifiable via UpdateSecret. At launch Access by Allias will only be supported on GetSecretVersion and AccessSecretVersion.
        /// </summary>
        public readonly ImmutableDictionary<string, string> VersionAliases;

        [OutputConstructor]
        private GetSecretResult(
            string createTime,

            string etag,

            string expireTime,

            ImmutableDictionary<string, string> labels,

            string name,

            Outputs.ReplicationResponse replication,

            Outputs.RotationResponse rotation,

            ImmutableArray<Outputs.TopicResponse> topics,

            string ttl,

            ImmutableDictionary<string, string> versionAliases)
        {
            CreateTime = createTime;
            Etag = etag;
            ExpireTime = expireTime;
            Labels = labels;
            Name = name;
            Replication = replication;
            Rotation = rotation;
            Topics = topics;
            Ttl = ttl;
            VersionAliases = versionAliases;
        }
    }
}
