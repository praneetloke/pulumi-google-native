// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V1.Inputs
{

    /// <summary>
    /// A widget that displays a stream of log.
    /// </summary>
    public sealed class LogsPanelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A filter that chooses which log entries to return. See Advanced Logs Queries (https://cloud.google.com/logging/docs/view/advanced-queries). Only log entries that match the filter are returned. An empty filter matches all log entries.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        [Input("resourceNames")]
        private InputList<string>? _resourceNames;

        /// <summary>
        /// The names of logging resources to collect logs for. Currently only projects are supported. If empty, the widget will default to the host project.
        /// </summary>
        public InputList<string> ResourceNames
        {
            get => _resourceNames ?? (_resourceNames = new InputList<string>());
            set => _resourceNames = value;
        }

        public LogsPanelArgs()
        {
        }
        public static new LogsPanelArgs Empty => new LogsPanelArgs();
    }
}
