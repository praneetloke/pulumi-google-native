// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Outputs
{

    [OutputType]
    public sealed class EnterpriseCrmFrontendsEventbusProtoProtoParameterArrayResponse
    {
        public readonly ImmutableArray<ImmutableDictionary<string, object>> ProtoValues;

        [OutputConstructor]
        private EnterpriseCrmFrontendsEventbusProtoProtoParameterArrayResponse(ImmutableArray<ImmutableDictionary<string, object>> protoValues)
        {
            ProtoValues = protoValues;
        }
    }
}
