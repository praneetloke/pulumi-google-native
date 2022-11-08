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
    /// Describes string and array limits when writing to logs. When a limit is exceeded the *shortener_type* describes how to shorten the field. next_id: 6
    /// </summary>
    [OutputType]
    public sealed class EnterpriseCrmLoggingGwsFieldLimitsResponse
    {
        public readonly string LogAction;
        /// <summary>
        /// To which type(s) of logs the limits apply.
        /// </summary>
        public readonly ImmutableArray<string> LogType;
        /// <summary>
        /// maximum array size. If the array exceds this size, the field (list) is truncated.
        /// </summary>
        public readonly int MaxArraySize;
        /// <summary>
        /// maximum string length. If the field exceeds this amount the field is shortened.
        /// </summary>
        public readonly int MaxStringLength;
        public readonly string ShortenerType;

        [OutputConstructor]
        private EnterpriseCrmLoggingGwsFieldLimitsResponse(
            string logAction,

            ImmutableArray<string> logType,

            int maxArraySize,

            int maxStringLength,

            string shortenerType)
        {
            LogAction = logAction;
            LogType = logType;
            MaxArraySize = maxArraySize;
            MaxStringLength = maxStringLength;
            ShortenerType = shortenerType;
        }
    }
}
