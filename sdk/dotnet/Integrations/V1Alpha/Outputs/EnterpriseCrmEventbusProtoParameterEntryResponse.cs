// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Outputs
{

    /// <summary>
    /// Key-value pair of EventBus parameters.
    /// </summary>
    [OutputType]
    public sealed class EnterpriseCrmEventbusProtoParameterEntryResponse
    {
        /// <summary>
        /// Key is used to retrieve the corresponding parameter value. This should be unique for a given fired event. These parameters must be predefined in the integration definition.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Values for the defined keys. Each value can either be string, int, double or any proto message.
        /// </summary>
        public readonly Outputs.EnterpriseCrmEventbusProtoParameterValueTypeResponse Value;

        [OutputConstructor]
        private EnterpriseCrmEventbusProtoParameterEntryResponse(
            string key,

            Outputs.EnterpriseCrmEventbusProtoParameterValueTypeResponse value)
        {
            Key = key;
            Value = value;
        }
    }
}
