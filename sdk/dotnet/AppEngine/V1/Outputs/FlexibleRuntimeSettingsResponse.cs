// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1.Outputs
{

    /// <summary>
    /// Runtime settings for the App Engine flexible environment.
    /// </summary>
    [OutputType]
    public sealed class FlexibleRuntimeSettingsResponse
    {
        /// <summary>
        /// The operating system of the application runtime.
        /// </summary>
        public readonly string OperatingSystem;
        /// <summary>
        /// The runtime version of an App Engine flexible application.
        /// </summary>
        public readonly string RuntimeVersion;

        [OutputConstructor]
        private FlexibleRuntimeSettingsResponse(
            string operatingSystem,

            string runtimeVersion)
        {
            OperatingSystem = operatingSystem;
            RuntimeVersion = runtimeVersion;
        }
    }
}
