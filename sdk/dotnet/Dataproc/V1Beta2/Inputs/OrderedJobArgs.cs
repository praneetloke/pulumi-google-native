// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1Beta2.Inputs
{

    /// <summary>
    /// A job executed by the workflow.
    /// </summary>
    public sealed class OrderedJobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Job is a Hadoop job.
        /// </summary>
        [Input("hadoopJob")]
        public Input<Inputs.HadoopJobArgs>? HadoopJob { get; set; }

        /// <summary>
        /// Optional. Job is a Hive job.
        /// </summary>
        [Input("hiveJob")]
        public Input<Inputs.HiveJobArgs>? HiveJob { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. The labels to associate with this job.Label keys must be between 1 and 63 characters long, and must conform to the following regular expression: \p{Ll}\p{Lo}{0,62}Label values must be between 1 and 63 characters long, and must conform to the following regular expression: \p{Ll}\p{Lo}\p{N}_-{0,63}No more than 32 labels can be associated with a given job.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Optional. Job is a Pig job.
        /// </summary>
        [Input("pigJob")]
        public Input<Inputs.PigJobArgs>? PigJob { get; set; }

        [Input("prerequisiteStepIds")]
        private InputList<string>? _prerequisiteStepIds;

        /// <summary>
        /// Optional. The optional list of prerequisite job step_ids. If not specified, the job will start at the beginning of workflow.
        /// </summary>
        public InputList<string> PrerequisiteStepIds
        {
            get => _prerequisiteStepIds ?? (_prerequisiteStepIds = new InputList<string>());
            set => _prerequisiteStepIds = value;
        }

        /// <summary>
        /// Optional. Job is a Presto job.
        /// </summary>
        [Input("prestoJob")]
        public Input<Inputs.PrestoJobArgs>? PrestoJob { get; set; }

        /// <summary>
        /// Optional. Job is a PySpark job.
        /// </summary>
        [Input("pysparkJob")]
        public Input<Inputs.PySparkJobArgs>? PysparkJob { get; set; }

        /// <summary>
        /// Optional. Job scheduling configuration.
        /// </summary>
        [Input("scheduling")]
        public Input<Inputs.JobSchedulingArgs>? Scheduling { get; set; }

        /// <summary>
        /// Optional. Job is a Spark job.
        /// </summary>
        [Input("sparkJob")]
        public Input<Inputs.SparkJobArgs>? SparkJob { get; set; }

        /// <summary>
        /// Optional. Job is a SparkR job.
        /// </summary>
        [Input("sparkRJob")]
        public Input<Inputs.SparkRJobArgs>? SparkRJob { get; set; }

        /// <summary>
        /// Optional. Job is a SparkSql job.
        /// </summary>
        [Input("sparkSqlJob")]
        public Input<Inputs.SparkSqlJobArgs>? SparkSqlJob { get; set; }

        /// <summary>
        /// The step id. The id must be unique among all jobs within the template.The step id is used as prefix for job id, as job goog-dataproc-workflow-step-id label, and in prerequisiteStepIds field from other steps.The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
        /// </summary>
        [Input("stepId", required: true)]
        public Input<string> StepId { get; set; } = null!;

        public OrderedJobArgs()
        {
        }
        public static new OrderedJobArgs Empty => new OrderedJobArgs();
    }
}
