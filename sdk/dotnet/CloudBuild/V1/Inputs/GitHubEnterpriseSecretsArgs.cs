// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Inputs
{

    /// <summary>
    /// GitHubEnterpriseSecrets represents the names of all necessary secrets in Secret Manager for a GitHub Enterprise server. Format is: projects//secrets/.
    /// </summary>
    public sealed class GitHubEnterpriseSecretsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource name for the OAuth client ID secret in Secret Manager.
        /// </summary>
        [Input("oauthClientIdName")]
        public Input<string>? OauthClientIdName { get; set; }

        /// <summary>
        /// The resource name for the OAuth client ID secret version in Secret Manager.
        /// </summary>
        [Input("oauthClientIdVersionName")]
        public Input<string>? OauthClientIdVersionName { get; set; }

        /// <summary>
        /// The resource name for the OAuth secret in Secret Manager.
        /// </summary>
        [Input("oauthSecretName")]
        public Input<string>? OauthSecretName { get; set; }

        /// <summary>
        /// The resource name for the OAuth secret secret version in Secret Manager.
        /// </summary>
        [Input("oauthSecretVersionName")]
        public Input<string>? OauthSecretVersionName { get; set; }

        /// <summary>
        /// The resource name for the private key secret.
        /// </summary>
        [Input("privateKeyName")]
        public Input<string>? PrivateKeyName { get; set; }

        /// <summary>
        /// The resource name for the private key secret version.
        /// </summary>
        [Input("privateKeyVersionName")]
        public Input<string>? PrivateKeyVersionName { get; set; }

        /// <summary>
        /// The resource name for the webhook secret in Secret Manager.
        /// </summary>
        [Input("webhookSecretName")]
        public Input<string>? WebhookSecretName { get; set; }

        /// <summary>
        /// The resource name for the webhook secret secret version in Secret Manager.
        /// </summary>
        [Input("webhookSecretVersionName")]
        public Input<string>? WebhookSecretVersionName { get; set; }

        public GitHubEnterpriseSecretsArgs()
        {
        }
        public static new GitHubEnterpriseSecretsArgs Empty => new GitHubEnterpriseSecretsArgs();
    }
}
