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
    /// This is deprecated and has no effect. Do not use.
    /// </summary>
    public sealed class LogConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        [Input("cloudAudit")]
        public Input<Inputs.LogConfigCloudAuditOptionsArgs>? CloudAudit { get; set; }

        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        [Input("counter")]
        public Input<Inputs.LogConfigCounterOptionsArgs>? Counter { get; set; }

        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        [Input("dataAccess")]
        public Input<Inputs.LogConfigDataAccessOptionsArgs>? DataAccess { get; set; }

        public LogConfigArgs()
        {
        }
        public static new LogConfigArgs Empty => new LogConfigArgs();
    }
}
