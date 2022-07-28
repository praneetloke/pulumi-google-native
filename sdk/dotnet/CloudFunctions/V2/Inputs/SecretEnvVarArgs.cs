// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudFunctions.V2.Inputs
{

    /// <summary>
    /// Configuration for a secret environment variable. It has the information necessary to fetch the secret value from secret manager and expose it as an environment variable.
    /// </summary>
    public sealed class SecretEnvVarArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the environment variable.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Project identifier (preferably project number but can also be the project ID) of the project that contains the secret. If not set, it is assumed that the secret is in the same project as the function.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Name of the secret in secret manager (not the full resource name).
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        /// <summary>
        /// Version of the secret (version number or the string 'latest'). It is recommended to use a numeric version for secret environment variables as any updates to the secret value is not reflected until new instances start.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public SecretEnvVarArgs()
        {
        }
        public static new SecretEnvVarArgs Empty => new SecretEnvVarArgs();
    }
}
