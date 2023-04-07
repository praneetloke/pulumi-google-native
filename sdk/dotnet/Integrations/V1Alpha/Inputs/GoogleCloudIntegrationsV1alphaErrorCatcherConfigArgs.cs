// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Inputs
{

    /// <summary>
    /// Configuration detail of a error catch task
    /// </summary>
    public sealed class GoogleCloudIntegrationsV1alphaErrorCatcherConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. User-provided description intended to give more business context about the error catcher config.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// An error catcher id is string representation for the error catcher config. Within a workflow, error_catcher_id uniquely identifies an error catcher config among all error catcher configs for the workflow
        /// </summary>
        [Input("errorCatcherId", required: true)]
        public Input<string> ErrorCatcherId { get; set; } = null!;

        /// <summary>
        /// A number to uniquely identify each error catcher config within the workflow on UI.
        /// </summary>
        [Input("errorCatcherNumber", required: true)]
        public Input<string> ErrorCatcherNumber { get; set; } = null!;

        /// <summary>
        /// Optional. The user created label for a particular error catcher. Optional.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Optional. Informs the front-end application where to draw this error catcher config on the UI.
        /// </summary>
        [Input("position")]
        public Input<Inputs.GoogleCloudIntegrationsV1alphaCoordinateArgs>? Position { get; set; }

        [Input("startErrorTasks", required: true)]
        private InputList<Inputs.GoogleCloudIntegrationsV1alphaNextTaskArgs>? _startErrorTasks;

        /// <summary>
        /// The set of start tasks that are to be executed for the error catch flow
        /// </summary>
        public InputList<Inputs.GoogleCloudIntegrationsV1alphaNextTaskArgs> StartErrorTasks
        {
            get => _startErrorTasks ?? (_startErrorTasks = new InputList<Inputs.GoogleCloudIntegrationsV1alphaNextTaskArgs>());
            set => _startErrorTasks = value;
        }

        public GoogleCloudIntegrationsV1alphaErrorCatcherConfigArgs()
        {
        }
        public static new GoogleCloudIntegrationsV1alphaErrorCatcherConfigArgs Empty => new GoogleCloudIntegrationsV1alphaErrorCatcherConfigArgs();
    }
}
