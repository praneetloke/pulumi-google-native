// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Inputs
{

    /// <summary>
    /// Job scheduling options.
    /// </summary>
    public sealed class JobSchedulingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Maximum number of times per hour a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed.A job may be reported as thrashing if driver exits with non-zero code 4 times within 10 minute window.Maximum value is 10.Note: Currently, this restartable job option is not supported in Dataproc workflow template (https://cloud.google.com/dataproc/docs/concepts/workflows/using-workflows#adding_jobs_to_a_template) jobs.
        /// </summary>
        [Input("maxFailuresPerHour")]
        public Input<int>? MaxFailuresPerHour { get; set; }

        /// <summary>
        /// Optional. Maximum number of times in total a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed. Maximum value is 240.Note: Currently, this restartable job option is not supported in Dataproc workflow template (https://cloud.google.com/dataproc/docs/concepts/workflows/using-workflows#adding_jobs_to_a_template) jobs.
        /// </summary>
        [Input("maxFailuresTotal")]
        public Input<int>? MaxFailuresTotal { get; set; }

        public JobSchedulingArgs()
        {
        }
        public static new JobSchedulingArgs Empty => new JobSchedulingArgs();
    }
}
