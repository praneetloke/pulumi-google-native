// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha.Outputs
{

    /// <summary>
    /// MembershipBindingLifecycleState describes the state of a Binding resource.
    /// </summary>
    [OutputType]
    public sealed class MembershipBindingLifecycleStateResponse
    {
        /// <summary>
        /// The current state of the MembershipBinding resource.
        /// </summary>
        public readonly string Code;

        [OutputConstructor]
        private MembershipBindingLifecycleStateResponse(string code)
        {
            Code = code;
        }
    }
}
