// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Inputs
{

    /// <summary>
    /// Indicates that the builder claims certain fields in this message to be complete.
    /// </summary>
    public sealed class CompletenessArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, the builder claims that recipe.arguments is complete, meaning that all external inputs are properly captured in the recipe.
        /// </summary>
        [Input("arguments")]
        public Input<bool>? Arguments { get; set; }

        /// <summary>
        /// If true, the builder claims that recipe.environment is claimed to be complete.
        /// </summary>
        [Input("environment")]
        public Input<bool>? Environment { get; set; }

        /// <summary>
        /// If true, the builder claims that materials are complete, usually through some controls to prevent network access. Sometimes called "hermetic".
        /// </summary>
        [Input("materials")]
        public Input<bool>? Materials { get; set; }

        public CompletenessArgs()
        {
        }
        public static new CompletenessArgs Empty => new CompletenessArgs();
    }
}
