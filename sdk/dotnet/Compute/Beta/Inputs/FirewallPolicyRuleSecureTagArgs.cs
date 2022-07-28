// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    public sealed class FirewallPolicyRuleSecureTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the secure tag, created with TagManager's TagValue API.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallPolicyRuleSecureTagArgs()
        {
        }
        public static new FirewallPolicyRuleSecureTagArgs Empty => new FirewallPolicyRuleSecureTagArgs();
    }
}
