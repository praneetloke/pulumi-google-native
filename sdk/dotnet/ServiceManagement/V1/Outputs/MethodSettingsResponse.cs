// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ServiceManagement.V1.Outputs
{

    /// <summary>
    /// Describes the generator configuration for a method.
    /// </summary>
    [OutputType]
    public sealed class MethodSettingsResponse
    {
        /// <summary>
        /// List of top-level fields of the request message, that should be automatically populated by the client libraries based on their (google.api.field_info).format. Currently supported format: UUID4. Example of a YAML configuration: publishing: method_settings: - selector: google.example.v1.ExampleService.CreateExample auto_populated_fields: - request_id
        /// </summary>
        public readonly ImmutableArray<string> AutoPopulatedFields;
        /// <summary>
        /// Describes settings to use for long-running operations when generating API methods for RPCs. Complements RPCs that use the annotations in google/longrunning/operations.proto. Example of a YAML configuration:: publishing: method_settings: - selector: google.cloud.speech.v2.Speech.BatchRecognize long_running: initial_poll_delay: seconds: 60 # 1 minute poll_delay_multiplier: 1.5 max_poll_delay: seconds: 360 # 6 minutes total_poll_timeout: seconds: 54000 # 90 minutes
        /// </summary>
        public readonly Outputs.LongRunningResponse LongRunning;
        /// <summary>
        /// The fully qualified name of the method, for which the options below apply. This is used to find the method to apply the options.
        /// </summary>
        public readonly string Selector;

        [OutputConstructor]
        private MethodSettingsResponse(
            ImmutableArray<string> autoPopulatedFields,

            Outputs.LongRunningResponse longRunning,

            string selector)
        {
            AutoPopulatedFields = autoPopulatedFields;
            LongRunning = longRunning;
            Selector = selector;
        }
    }
}
