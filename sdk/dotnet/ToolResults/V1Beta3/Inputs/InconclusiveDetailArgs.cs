// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ToolResults.V1Beta3.Inputs
{

    /// <summary>
    /// Details for an outcome with an INCONCLUSIVE outcome summary.
    /// </summary>
    public sealed class InconclusiveDetailArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If the end user aborted the test execution before a pass or fail could be determined. For example, the user pressed ctrl-c which sent a kill signal to the test runner while the test was running.
        /// </summary>
        [Input("abortedByUser")]
        public Input<bool>? AbortedByUser { get; set; }

        /// <summary>
        /// If results are being provided to the user in certain cases of infrastructure failures
        /// </summary>
        [Input("hasErrorLogs")]
        public Input<bool>? HasErrorLogs { get; set; }

        /// <summary>
        /// If the test runner could not determine success or failure because the test depends on a component other than the system under test which failed. For example, a mobile test requires provisioning a device where the test executes, and that provisioning can fail.
        /// </summary>
        [Input("infrastructureFailure")]
        public Input<bool>? InfrastructureFailure { get; set; }

        public InconclusiveDetailArgs()
        {
        }
        public static new InconclusiveDetailArgs Empty => new InconclusiveDetailArgs();
    }
}
