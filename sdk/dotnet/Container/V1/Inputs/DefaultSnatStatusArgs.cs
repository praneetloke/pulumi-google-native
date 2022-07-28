// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Inputs
{

    /// <summary>
    /// DefaultSnatStatus contains the desired state of whether default sNAT should be disabled on the cluster.
    /// </summary>
    public sealed class DefaultSnatStatusArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disables cluster default sNAT rules.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        public DefaultSnatStatusArgs()
        {
        }
        public static new DefaultSnatStatusArgs Empty => new DefaultSnatStatusArgs();
    }
}
