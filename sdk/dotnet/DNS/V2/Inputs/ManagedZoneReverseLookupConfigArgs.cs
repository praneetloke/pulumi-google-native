// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V2.Inputs
{

    public sealed class ManagedZoneReverseLookupConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        public ManagedZoneReverseLookupConfigArgs()
        {
        }
        public static new ManagedZoneReverseLookupConfigArgs Empty => new ManagedZoneReverseLookupConfigArgs();
    }
}
