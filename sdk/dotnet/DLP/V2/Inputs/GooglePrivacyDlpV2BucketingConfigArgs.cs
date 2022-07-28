// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Inputs
{

    /// <summary>
    /// Generalization function that buckets values based on ranges. The ranges and replacement values are dynamically provided by the user for custom behavior, such as 1-30 -&gt; LOW 31-65 -&gt; MEDIUM 66-100 -&gt; HIGH This can be used on data of type: number, long, string, timestamp. If the bound `Value` type differs from the type of data being transformed, we will first attempt converting the type of the data to be transformed to match the type of the bound before comparing. See https://cloud.google.com/dlp/docs/concepts-bucketing to learn more.
    /// </summary>
    public sealed class GooglePrivacyDlpV2BucketingConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("buckets")]
        private InputList<Inputs.GooglePrivacyDlpV2BucketArgs>? _buckets;

        /// <summary>
        /// Set of buckets. Ranges must be non-overlapping.
        /// </summary>
        public InputList<Inputs.GooglePrivacyDlpV2BucketArgs> Buckets
        {
            get => _buckets ?? (_buckets = new InputList<Inputs.GooglePrivacyDlpV2BucketArgs>());
            set => _buckets = value;
        }

        public GooglePrivacyDlpV2BucketingConfigArgs()
        {
        }
        public static new GooglePrivacyDlpV2BucketingConfigArgs Empty => new GooglePrivacyDlpV2BucketingConfigArgs();
    }
}
