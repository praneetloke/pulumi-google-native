// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ToolResults.V1Beta3.Inputs
{

    /// <summary>
    /// A game loop test of an iOS application.
    /// </summary>
    public sealed class IosTestLoopArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bundle ID of the app.
        /// </summary>
        [Input("bundleId")]
        public Input<string>? BundleId { get; set; }

        public IosTestLoopArgs()
        {
        }
        public static new IosTestLoopArgs Empty => new IosTestLoopArgs();
    }
}
