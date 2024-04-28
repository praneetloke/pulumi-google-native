// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha.Inputs
{

    /// <summary>
    /// **Cloud Build**: Configurations for each Cloud Build enabled cluster.
    /// </summary>
    public sealed class CloudBuildMembershipSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether it is allowed to run the privileged builds on the cluster or not.
        /// </summary>
        [Input("securityPolicy")]
        public Input<Pulumi.GoogleNative.GKEHub.V1Alpha.CloudBuildMembershipSpecSecurityPolicy>? SecurityPolicy { get; set; }

        /// <summary>
        /// Version of the cloud build software on the cluster.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public CloudBuildMembershipSpecArgs()
        {
        }
        public static new CloudBuildMembershipSpecArgs Empty => new CloudBuildMembershipSpecArgs();
    }
}
