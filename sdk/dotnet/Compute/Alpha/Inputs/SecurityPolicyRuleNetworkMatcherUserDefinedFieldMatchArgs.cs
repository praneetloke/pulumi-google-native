// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    public sealed class SecurityPolicyRuleNetworkMatcherUserDefinedFieldMatchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the user-defined field, as given in the definition.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// Matching values of the field. Each element can be a 32-bit unsigned decimal or hexadecimal (starting with "0x") number (e.g. "64") or range (e.g. "0x400-0x7ff").
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public SecurityPolicyRuleNetworkMatcherUserDefinedFieldMatchArgs()
        {
        }
        public static new SecurityPolicyRuleNetworkMatcherUserDefinedFieldMatchArgs Empty => new SecurityPolicyRuleNetworkMatcherUserDefinedFieldMatchArgs();
    }
}
