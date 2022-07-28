// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta.Inputs
{

    /// <summary>
    /// Readiness checking configuration for VM instances. Unhealthy instances are removed from traffic rotation.
    /// </summary>
    public sealed class ReadinessCheckArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A maximum time limit on application initialization, measured from moment the application successfully replies to a healthcheck until it is ready to serve traffic.
        /// </summary>
        [Input("appStartTimeout")]
        public Input<string>? AppStartTimeout { get; set; }

        /// <summary>
        /// Interval between health checks.
        /// </summary>
        [Input("checkInterval")]
        public Input<string>? CheckInterval { get; set; }

        /// <summary>
        /// Number of consecutive failed checks required before removing traffic.
        /// </summary>
        [Input("failureThreshold")]
        public Input<int>? FailureThreshold { get; set; }

        /// <summary>
        /// Host header to send when performing a HTTP Readiness check. Example: "myapp.appspot.com"
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The request path.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Number of consecutive successful checks required before receiving traffic.
        /// </summary>
        [Input("successThreshold")]
        public Input<int>? SuccessThreshold { get; set; }

        /// <summary>
        /// Time before the check is considered failed.
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        public ReadinessCheckArgs()
        {
        }
        public static new ReadinessCheckArgs Empty => new ReadinessCheckArgs();
    }
}
