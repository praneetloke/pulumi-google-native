// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// Represents a match condition that incoming traffic is evaluated against. Exactly one field must be specified.
    /// </summary>
    public sealed class SecurityPolicyRuleMatcherArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration options available when specifying versioned_expr. This field must be specified if versioned_expr is specified and cannot be specified if versioned_expr is not specified.
        /// </summary>
        [Input("config")]
        public Input<Inputs.SecurityPolicyRuleMatcherConfigArgs>? Config { get; set; }

        /// <summary>
        /// User defined CEVAL expression. A CEVAL expression is used to specify match criteria such as origin.ip, source.region_code and contents in the request header.
        /// </summary>
        [Input("expr")]
        public Input<Inputs.ExprArgs>? Expr { get; set; }

        /// <summary>
        /// Preconfigured versioned expression. If this field is specified, config must also be specified. Available preconfigured expressions along with their requirements are: SRC_IPS_V1 - must specify the corresponding src_ip_range field in config.
        /// </summary>
        [Input("versionedExpr")]
        public Input<Pulumi.GoogleNative.Compute.V1.SecurityPolicyRuleMatcherVersionedExpr>? VersionedExpr { get; set; }

        public SecurityPolicyRuleMatcherArgs()
        {
        }
        public static new SecurityPolicyRuleMatcherArgs Empty => new SecurityPolicyRuleMatcherArgs();
    }
}
