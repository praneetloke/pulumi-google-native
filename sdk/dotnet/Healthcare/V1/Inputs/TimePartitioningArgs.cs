// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1.Inputs
{

    /// <summary>
    /// Configuration for FHIR BigQuery time-partitioned tables.
    /// </summary>
    public sealed class TimePartitioningArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of milliseconds for which to keep the storage for a partition.
        /// </summary>
        [Input("expirationMs")]
        public Input<string>? ExpirationMs { get; set; }

        /// <summary>
        /// Type of partitioning.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Healthcare.V1.TimePartitioningType>? Type { get; set; }

        public TimePartitioningArgs()
        {
        }
        public static new TimePartitioningArgs Empty => new TimePartitioningArgs();
    }
}
