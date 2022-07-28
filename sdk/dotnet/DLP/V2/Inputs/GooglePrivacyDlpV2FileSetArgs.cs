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
    /// Set of files to scan.
    /// </summary>
    public sealed class GooglePrivacyDlpV2FileSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The regex-filtered set of files to scan. Exactly one of `url` or `regex_file_set` must be set.
        /// </summary>
        [Input("regexFileSet")]
        public Input<Inputs.GooglePrivacyDlpV2CloudStorageRegexFileSetArgs>? RegexFileSet { get; set; }

        /// <summary>
        /// The Cloud Storage url of the file(s) to scan, in the format `gs:///`. Trailing wildcard in the path is allowed. If the url ends in a trailing slash, the bucket or directory represented by the url will be scanned non-recursively (content in sub-directories will not be scanned). This means that `gs://mybucket/` is equivalent to `gs://mybucket/*`, and `gs://mybucket/directory/` is equivalent to `gs://mybucket/directory/*`. Exactly one of `url` or `regex_file_set` must be set.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public GooglePrivacyDlpV2FileSetArgs()
        {
        }
        public static new GooglePrivacyDlpV2FileSetArgs Empty => new GooglePrivacyDlpV2FileSetArgs();
    }
}
