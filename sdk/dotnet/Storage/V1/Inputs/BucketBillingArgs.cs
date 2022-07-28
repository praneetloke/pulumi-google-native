// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Storage.V1.Inputs
{

    /// <summary>
    /// The bucket's billing configuration.
    /// </summary>
    public sealed class BucketBillingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When set to true, Requester Pays is enabled for this bucket.
        /// </summary>
        [Input("requesterPays")]
        public Input<bool>? RequesterPays { get; set; }

        public BucketBillingArgs()
        {
        }
        public static new BucketBillingArgs Empty => new BucketBillingArgs();
    }
}
