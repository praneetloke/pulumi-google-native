// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Inputs
{

    public sealed class GoogleCloudDataplexV1EnvironmentSessionSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. If True, this causes sessions to be pre-created and available for faster startup to enable interactive exploration use-cases. This defaults to False to avoid additional billed charges. These can only be set to True for the environment with name set to "default", and with default configuration.
        /// </summary>
        [Input("enableFastStartup")]
        public Input<bool>? EnableFastStartup { get; set; }

        /// <summary>
        /// Optional. The idle time configuration of the session. The session will be auto-terminated at the end of this period.
        /// </summary>
        [Input("maxIdleDuration")]
        public Input<string>? MaxIdleDuration { get; set; }

        public GoogleCloudDataplexV1EnvironmentSessionSpecArgs()
        {
        }
        public static new GoogleCloudDataplexV1EnvironmentSessionSpecArgs Empty => new GoogleCloudDataplexV1EnvironmentSessionSpecArgs();
    }
}
