// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contentwarehouse.V1.Outputs
{

    /// <summary>
    /// Represents the string value of the enum field.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudContentwarehouseV1EnumValueResponse
    {
        /// <summary>
        /// String value of the enum field. This must match defined set of enums in document schema using EnumTypeOptions.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GoogleCloudContentwarehouseV1EnumValueResponse(string value)
        {
            Value = value;
        }
    }
}
