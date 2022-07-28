// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1.Inputs
{

    /// <summary>
    /// Target scaling by request utilization. Only applicable in the App Engine flexible environment.
    /// </summary>
    public sealed class RequestUtilizationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Target number of concurrent requests.
        /// </summary>
        [Input("targetConcurrentRequests")]
        public Input<int>? TargetConcurrentRequests { get; set; }

        /// <summary>
        /// Target requests per second.
        /// </summary>
        [Input("targetRequestCountPerSecond")]
        public Input<int>? TargetRequestCountPerSecond { get; set; }

        public RequestUtilizationArgs()
        {
        }
        public static new RequestUtilizationArgs Empty => new RequestUtilizationArgs();
    }
}
