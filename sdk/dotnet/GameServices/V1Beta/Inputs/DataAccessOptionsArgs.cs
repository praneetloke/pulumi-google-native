// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GameServices.V1Beta.Inputs
{

    /// <summary>
    /// Write a Data Access (Gin) log
    /// </summary>
    public sealed class DataAccessOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("logMode")]
        public Input<Pulumi.GoogleNative.GameServices.V1Beta.DataAccessOptionsLogMode>? LogMode { get; set; }

        public DataAccessOptionsArgs()
        {
        }
        public static new DataAccessOptionsArgs Empty => new DataAccessOptionsArgs();
    }
}
