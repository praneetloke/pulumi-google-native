// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudApigeeV1RateRangeResponse
    {
        /// <summary>
        /// Ending value of the range. Set to 0 or `null` for the last range of values.
        /// </summary>
        public readonly string End;
        /// <summary>
        /// Fee to charge when total number of API calls falls within this range.
        /// </summary>
        public readonly Outputs.GoogleTypeMoneyResponse Fee;
        /// <summary>
        /// Starting value of the range. Set to 0 or `null` for the initial range of values.
        /// </summary>
        public readonly string Start;

        [OutputConstructor]
        private GoogleCloudApigeeV1RateRangeResponse(
            string end,

            Outputs.GoogleTypeMoneyResponse fee,

            string start)
        {
            End = end;
            Fee = fee;
            Start = start;
        }
    }
}
