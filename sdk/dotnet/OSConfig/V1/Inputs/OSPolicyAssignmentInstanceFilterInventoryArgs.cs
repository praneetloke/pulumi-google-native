// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1.Inputs
{

    /// <summary>
    /// VM inventory details.
    /// </summary>
    public sealed class OSPolicyAssignmentInstanceFilterInventoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The OS short name
        /// </summary>
        [Input("osShortName", required: true)]
        public Input<string> OsShortName { get; set; } = null!;

        /// <summary>
        /// The OS version Prefix matches are supported if asterisk(*) is provided as the last character. For example, to match all versions with a major version of `7`, specify the following value for this field `7.*` An empty string matches all OS versions.
        /// </summary>
        [Input("osVersion")]
        public Input<string>? OsVersion { get; set; }

        public OSPolicyAssignmentInstanceFilterInventoryArgs()
        {
        }
        public static new OSPolicyAssignmentInstanceFilterInventoryArgs Empty => new OSPolicyAssignmentInstanceFilterInventoryArgs();
    }
}
