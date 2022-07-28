// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    public sealed class RouterMd5AuthenticationKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Input only] Value of the key. For patch and update calls, it can be skipped to copy the value from the previous configuration. This is allowed if the key with the same name existed before the operation. Maximum length is 80 characters. Can only contain printable ASCII characters.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Name used to identify the key. Must be unique within a router. Must be referenced by at least one bgpPeer. Must comply with RFC1035.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public RouterMd5AuthenticationKeyArgs()
        {
        }
        public static new RouterMd5AuthenticationKeyArgs Empty => new RouterMd5AuthenticationKeyArgs();
    }
}
