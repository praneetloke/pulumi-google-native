// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Beta.Inputs
{

    /// <summary>
    /// CommonFleetDefaultMemberConfigSpec contains default configuration information for memberships of a fleet
    /// </summary>
    public sealed class CommonFleetDefaultMemberConfigSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identity Service-specific spec.
        /// </summary>
        [Input("identityservice")]
        public Input<Inputs.IdentityServiceMembershipSpecArgs>? Identityservice { get; set; }

        public CommonFleetDefaultMemberConfigSpecArgs()
        {
        }
        public static new CommonFleetDefaultMemberConfigSpecArgs Empty => new CommonFleetDefaultMemberConfigSpecArgs();
    }
}