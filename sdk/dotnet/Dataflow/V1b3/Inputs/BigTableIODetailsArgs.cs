// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Dataflow.V1b3.Inputs
{

    /// <summary>
    /// Metadata for a Cloud BigTable connector used by the job.
    /// </summary>
    public sealed class BigTableIODetailsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// InstanceId accessed in the connection.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// ProjectId accessed in the connection.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// TableId accessed in the connection.
        /// </summary>
        [Input("tableId")]
        public Input<string>? TableId { get; set; }

        public BigTableIODetailsArgs()
        {
        }
    }
}
