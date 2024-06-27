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
    /// Contains information about how a particular google.dataflow.v1beta3.Step will be executed.
    /// </summary>
    public sealed class JobExecutionStageInfoArgs : global::Pulumi.ResourceArgs
    {
        [Input("stepName")]
        private InputList<string>? _stepName;

        /// <summary>
        /// The steps associated with the execution stage. Note that stages may have several steps, and that a given step might be run by more than one stage.
        /// </summary>
        public InputList<string> StepName
        {
            get => _stepName ?? (_stepName = new InputList<string>());
            set => _stepName = value;
        }

        public JobExecutionStageInfoArgs()
        {
        }
        public static new JobExecutionStageInfoArgs Empty => new JobExecutionStageInfoArgs();
    }
}