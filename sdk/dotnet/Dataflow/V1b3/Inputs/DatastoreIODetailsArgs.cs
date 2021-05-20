// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataflow.V1b3.Inputs
{

    /// <summary>
    /// Metadata for a Datastore connector used by the job.
    /// </summary>
    public sealed class DatastoreIODetailsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Namespace used in the connection.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// ProjectId accessed in the connection.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public DatastoreIODetailsArgs()
        {
        }
    }
}
