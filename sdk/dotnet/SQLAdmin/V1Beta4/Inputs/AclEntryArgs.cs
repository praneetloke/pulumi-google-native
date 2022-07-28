// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1Beta4.Inputs
{

    /// <summary>
    /// An entry for an Access Control list.
    /// </summary>
    public sealed class AclEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time when this access control entry expires in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        [Input("expirationTime")]
        public Input<string>? ExpirationTime { get; set; }

        /// <summary>
        /// This is always `sql#aclEntry`.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Optional. A label to identify this entry.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The allowlisted value for the access control list.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public AclEntryArgs()
        {
        }
        public static new AclEntryArgs Empty => new AclEntryArgs();
    }
}
