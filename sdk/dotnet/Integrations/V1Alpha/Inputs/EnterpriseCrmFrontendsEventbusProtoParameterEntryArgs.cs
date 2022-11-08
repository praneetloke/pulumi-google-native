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
    /// Key-value pair of EventBus parameters.
    /// </summary>
    public sealed class EnterpriseCrmFrontendsEventbusProtoParameterEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Explicitly getting the type of the parameter.
        /// </summary>
        [Input("dataType")]
        public Input<Pulumi.GoogleNative.Integrations.V1Alpha.EnterpriseCrmFrontendsEventbusProtoParameterEntryDataType>? DataType { get; set; }

        /// <summary>
        /// Key is used to retrieve the corresponding parameter value. This should be unique for a given fired event. These parameters must be predefined in the workflow definition.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Values for the defined keys. Each value can either be string, int, double or any proto message.
        /// </summary>
        [Input("value")]
        public Input<Inputs.EnterpriseCrmFrontendsEventbusProtoParameterValueTypeArgs>? Value { get; set; }

        public EnterpriseCrmFrontendsEventbusProtoParameterEntryArgs()
        {
        }
        public static new EnterpriseCrmFrontendsEventbusProtoParameterEntryArgs Empty => new EnterpriseCrmFrontendsEventbusProtoParameterEntryArgs();
    }
}
