// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Inputs
{

    public sealed class EnterpriseCrmEventbusProtoParamSpecEntryValidationRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("doubleRange")]
        public Input<Inputs.EnterpriseCrmEventbusProtoParamSpecEntryValidationRuleDoubleRangeArgs>? DoubleRange { get; set; }

        [Input("intRange")]
        public Input<Inputs.EnterpriseCrmEventbusProtoParamSpecEntryValidationRuleIntRangeArgs>? IntRange { get; set; }

        [Input("stringRegex")]
        public Input<Inputs.EnterpriseCrmEventbusProtoParamSpecEntryValidationRuleStringRegexArgs>? StringRegex { get; set; }

        public EnterpriseCrmEventbusProtoParamSpecEntryValidationRuleArgs()
        {
        }
        public static new EnterpriseCrmEventbusProtoParamSpecEntryValidationRuleArgs Empty => new EnterpriseCrmEventbusProtoParamSpecEntryValidationRuleArgs();
    }
}
