// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Inputs
{

    public sealed class EnterpriseCrmEventbusProtoSerializedObjectParameterArgs : global::Pulumi.ResourceArgs
    {
        [Input("objectValue")]
        public Input<string>? ObjectValue { get; set; }

        public EnterpriseCrmEventbusProtoSerializedObjectParameterArgs()
        {
        }
        public static new EnterpriseCrmEventbusProtoSerializedObjectParameterArgs Empty => new EnterpriseCrmEventbusProtoSerializedObjectParameterArgs();
    }
}
