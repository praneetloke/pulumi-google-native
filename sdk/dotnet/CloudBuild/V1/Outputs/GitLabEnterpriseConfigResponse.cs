// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Outputs
{

    /// <summary>
    /// GitLabEnterpriseConfig represents the configuration for a GitLabEnterprise integration.
    /// </summary>
    [OutputType]
    public sealed class GitLabEnterpriseConfigResponse
    {
        /// <summary>
        /// Immutable. The URI of the GitlabEnterprise host.
        /// </summary>
        public readonly string HostUri;
        /// <summary>
        /// The Service Directory configuration to be used when reaching out to the GitLab Enterprise instance.
        /// </summary>
        public readonly Outputs.ServiceDirectoryConfigResponse ServiceDirectoryConfig;
        /// <summary>
        /// The SSL certificate to use in requests to GitLab Enterprise instances.
        /// </summary>
        public readonly string SslCa;

        [OutputConstructor]
        private GitLabEnterpriseConfigResponse(
            string hostUri,

            Outputs.ServiceDirectoryConfigResponse serviceDirectoryConfig,

            string sslCa)
        {
            HostUri = hostUri;
            ServiceDirectoryConfig = serviceDirectoryConfig;
            SslCa = sslCa;
        }
    }
}
