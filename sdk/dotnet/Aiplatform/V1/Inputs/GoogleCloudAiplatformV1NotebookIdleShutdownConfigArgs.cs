// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1.Inputs
{

    /// <summary>
    /// The idle shutdown configuration of NotebookRuntimeTemplate, which contains the idle_timeout as required field.
    /// </summary>
    public sealed class GoogleCloudAiplatformV1NotebookIdleShutdownConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether Idle Shutdown is disabled in this NotebookRuntimeTemplate.
        /// </summary>
        [Input("idleShutdownDisabled")]
        public Input<bool>? IdleShutdownDisabled { get; set; }

        /// <summary>
        /// Duration is accurate to the second. In Notebook, Idle Timeout is accurate to minute so the range of idle_timeout (second) is: 10 * 60 ~ 1440 * 60.
        /// </summary>
        [Input("idleTimeout", required: true)]
        public Input<string> IdleTimeout { get; set; } = null!;

        public GoogleCloudAiplatformV1NotebookIdleShutdownConfigArgs()
        {
        }
        public static new GoogleCloudAiplatformV1NotebookIdleShutdownConfigArgs Empty => new GoogleCloudAiplatformV1NotebookIdleShutdownConfigArgs();
    }
}