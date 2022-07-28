// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GameServices.V1Beta.Inputs
{

    /// <summary>
    /// Specifies what kind of log the caller must write
    /// </summary>
    public sealed class LogConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud audit options.
        /// </summary>
        [Input("cloudAudit")]
        public Input<Inputs.CloudAuditOptionsArgs>? CloudAudit { get; set; }

        /// <summary>
        /// Counter options.
        /// </summary>
        [Input("counter")]
        public Input<Inputs.CounterOptionsArgs>? Counter { get; set; }

        /// <summary>
        /// Data access options.
        /// </summary>
        [Input("dataAccess")]
        public Input<Inputs.DataAccessOptionsArgs>? DataAccess { get; set; }

        public LogConfigArgs()
        {
        }
        public static new LogConfigArgs Empty => new LogConfigArgs();
    }
}
