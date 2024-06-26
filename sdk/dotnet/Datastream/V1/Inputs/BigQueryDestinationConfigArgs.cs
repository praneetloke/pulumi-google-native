// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1.Inputs
{

    /// <summary>
    /// BigQuery destination configuration
    /// </summary>
    public sealed class BigQueryDestinationConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The guaranteed data freshness (in seconds) when querying tables created by the stream. Editing this field will only affect new tables created in the future, but existing tables will not be impacted. Lower values mean that queries will return fresher data, but may result in higher cost.
        /// </summary>
        [Input("dataFreshness")]
        public Input<string>? DataFreshness { get; set; }

        /// <summary>
        /// Single destination dataset.
        /// </summary>
        [Input("singleTargetDataset")]
        public Input<Inputs.SingleTargetDatasetArgs>? SingleTargetDataset { get; set; }

        /// <summary>
        /// Source hierarchy datasets.
        /// </summary>
        [Input("sourceHierarchyDatasets")]
        public Input<Inputs.SourceHierarchyDatasetsArgs>? SourceHierarchyDatasets { get; set; }

        public BigQueryDestinationConfigArgs()
        {
        }
        public static new BigQueryDestinationConfigArgs Empty => new BigQueryDestinationConfigArgs();
    }
}
