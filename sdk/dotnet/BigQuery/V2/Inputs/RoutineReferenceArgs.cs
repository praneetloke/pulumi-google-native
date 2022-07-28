// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Inputs
{

    public sealed class RoutineReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Required] The ID of the dataset containing this routine.
        /// </summary>
        [Input("datasetId")]
        public Input<string>? DatasetId { get; set; }

        /// <summary>
        /// [Required] The ID of the project containing this routine.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// [Required] The ID of the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.
        /// </summary>
        [Input("routineId")]
        public Input<string>? RoutineId { get; set; }

        public RoutineReferenceArgs()
        {
        }
        public static new RoutineReferenceArgs Empty => new RoutineReferenceArgs();
    }
}
