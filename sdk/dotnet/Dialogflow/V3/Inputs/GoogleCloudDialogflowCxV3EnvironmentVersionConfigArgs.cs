// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Inputs
{

    /// <summary>
    /// Configuration for the version.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3EnvironmentVersionConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Format: projects//locations//agents//flows//versions/.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public GoogleCloudDialogflowCxV3EnvironmentVersionConfigArgs()
        {
        }
        public static new GoogleCloudDialogflowCxV3EnvironmentVersionConfigArgs Empty => new GoogleCloudDialogflowCxV3EnvironmentVersionConfigArgs();
    }
}
