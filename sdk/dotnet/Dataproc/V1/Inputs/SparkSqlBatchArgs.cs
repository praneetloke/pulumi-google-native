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
    /// A configuration for running Apache Spark SQL (https://spark.apache.org/sql/) queries as a batch workload.
    /// </summary>
    public sealed class SparkSqlBatchArgs : global::Pulumi.ResourceArgs
    {
        [Input("jarFileUris")]
        private InputList<string>? _jarFileUris;

        /// <summary>
        /// Optional. HCFS URIs of jar files to be added to the Spark CLASSPATH.
        /// </summary>
        public InputList<string> JarFileUris
        {
            get => _jarFileUris ?? (_jarFileUris = new InputList<string>());
            set => _jarFileUris = value;
        }

        /// <summary>
        /// The HCFS URI of the script that contains Spark SQL queries to execute.
        /// </summary>
        [Input("queryFileUri", required: true)]
        public Input<string> QueryFileUri { get; set; } = null!;

        [Input("queryVariables")]
        private InputMap<string>? _queryVariables;

        /// <summary>
        /// Optional. Mapping of query variable names to values (equivalent to the Spark SQL command: SET name="value";).
        /// </summary>
        public InputMap<string> QueryVariables
        {
            get => _queryVariables ?? (_queryVariables = new InputMap<string>());
            set => _queryVariables = value;
        }

        public SparkSqlBatchArgs()
        {
        }
        public static new SparkSqlBatchArgs Empty => new SparkSqlBatchArgs();
    }
}
