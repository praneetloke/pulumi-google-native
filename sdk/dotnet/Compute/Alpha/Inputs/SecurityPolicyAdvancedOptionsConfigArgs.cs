// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    public sealed class SecurityPolicyAdvancedOptionsConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom configuration to apply the JSON parsing. Only applicable when json_parsing is set to STANDARD.
        /// </summary>
        [Input("jsonCustomConfig")]
        public Input<Inputs.SecurityPolicyAdvancedOptionsConfigJsonCustomConfigArgs>? JsonCustomConfig { get; set; }

        [Input("jsonParsing")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SecurityPolicyAdvancedOptionsConfigJsonParsing>? JsonParsing { get; set; }

        [Input("logLevel")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SecurityPolicyAdvancedOptionsConfigLogLevel>? LogLevel { get; set; }

        [Input("userIpRequestHeaders")]
        private InputList<string>? _userIpRequestHeaders;

        /// <summary>
        /// An optional list of case-insensitive request header names to use for resolving the callers client IP address.
        /// </summary>
        public InputList<string> UserIpRequestHeaders
        {
            get => _userIpRequestHeaders ?? (_userIpRequestHeaders = new InputList<string>());
            set => _userIpRequestHeaders = value;
        }

        public SecurityPolicyAdvancedOptionsConfigArgs()
        {
        }
        public static new SecurityPolicyAdvancedOptionsConfigArgs Empty => new SecurityPolicyAdvancedOptionsConfigArgs();
    }
}
