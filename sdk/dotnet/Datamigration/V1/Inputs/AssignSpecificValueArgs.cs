// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1.Inputs
{

    /// <summary>
    /// Set to a specific value (value is converted to fit the target data type)
    /// </summary>
    public sealed class AssignSpecificValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specific value to be assigned
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public AssignSpecificValueArgs()
        {
        }
        public static new AssignSpecificValueArgs Empty => new AssignSpecificValueArgs();
    }
}