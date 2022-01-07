// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V2.Inputs
{

    /// <summary>
    /// Settings for Binary Authorization feature.
    /// </summary>
    public sealed class GoogleCloudRunOpV2BinaryAuthorizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If present, indicates to use Breakglass using this justification. If use_default is False, then it must be empty. For more information on breakglass, see https://cloud.google.com/binary-authorization/docs/using-breakglass
        /// </summary>
        [Input("breakglassJustification")]
        public Input<string>? BreakglassJustification { get; set; }

        /// <summary>
        /// If True, indicates to use the default project's binary authorization policy. If False, binary authorization will be disabled.
        /// </summary>
        [Input("useDefault")]
        public Input<bool>? UseDefault { get; set; }

        public GoogleCloudRunOpV2BinaryAuthorizationArgs()
        {
        }
    }
}
