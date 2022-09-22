// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SourceRepo.V1
{
    public static class GetRepo
    {
        /// <summary>
        /// Returns information about a repo.
        /// </summary>
        public static Task<GetRepoResult> InvokeAsync(GetRepoArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepoResult>("google-native:sourcerepo/v1:getRepo", args ?? new GetRepoArgs(), options.WithDefaults());

        /// <summary>
        /// Returns information about a repo.
        /// </summary>
        public static Output<GetRepoResult> Invoke(GetRepoInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepoResult>("google-native:sourcerepo/v1:getRepo", args ?? new GetRepoInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepoArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("repoId", required: true)]
        public string RepoId { get; set; } = null!;

        public GetRepoArgs()
        {
        }
        public static new GetRepoArgs Empty => new GetRepoArgs();
    }

    public sealed class GetRepoInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("repoId", required: true)]
        public Input<string> RepoId { get; set; } = null!;

        public GetRepoInvokeArgs()
        {
        }
        public static new GetRepoInvokeArgs Empty => new GetRepoInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepoResult
    {
        /// <summary>
        /// How this repository mirrors a repository managed by another service. Read-only field.
        /// </summary>
        public readonly Outputs.MirrorConfigResponse MirrorConfig;
        /// <summary>
        /// Resource name of the repository, of the form `projects//repos/`. The repo name may contain slashes. eg, `projects/myproject/repos/name/with/slash`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// How this repository publishes a change in the repository through Cloud Pub/Sub. Keyed by the topic names.
        /// </summary>
        public readonly ImmutableDictionary<string, string> PubsubConfigs;
        /// <summary>
        /// The disk usage of the repo, in bytes. Read-only field. Size is only returned by GetRepo.
        /// </summary>
        public readonly string Size;
        /// <summary>
        /// URL to clone the repository from Google Cloud Source Repositories. Read-only field.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetRepoResult(
            Outputs.MirrorConfigResponse mirrorConfig,

            string name,

            ImmutableDictionary<string, string> pubsubConfigs,

            string size,

            string url)
        {
            MirrorConfig = mirrorConfig;
            Name = name;
            PubsubConfigs = pubsubConfigs;
            Size = size;
            Url = url;
        }
    }
}
