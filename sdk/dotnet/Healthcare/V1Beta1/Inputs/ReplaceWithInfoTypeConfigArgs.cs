// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1.Inputs
{

    /// <summary>
    /// When using the INSPECT_AND_TRANSFORM action, each match is replaced with the name of the info_type. For example, "My name is Jane" becomes "My name is [PERSON_NAME]." The TRANSFORM action is equivalent to redacting.
    /// </summary>
    public sealed class ReplaceWithInfoTypeConfigArgs : global::Pulumi.ResourceArgs
    {
        public ReplaceWithInfoTypeConfigArgs()
        {
        }
        public static new ReplaceWithInfoTypeConfigArgs Empty => new ReplaceWithInfoTypeConfigArgs();
    }
}
