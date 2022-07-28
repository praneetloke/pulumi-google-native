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
    /// Message representing a set of files in a Cloud Storage bucket. Regular expressions are used to allow fine-grained control over which files in the bucket to include. Included files are those that match at least one item in `include_regex` and do not match any items in `exclude_regex`. Note that a file that matches items from both lists will _not_ be included. For a match to occur, the entire file path (i.e., everything in the url after the bucket name) must match the regular expression. For example, given the input `{bucket_name: "mybucket", include_regex: ["directory1/.*"], exclude_regex: ["directory1/excluded.*"]}`: * `gs://mybucket/directory1/myfile` will be included * `gs://mybucket/directory1/directory2/myfile` will be included (`.*` matches across `/`) * `gs://mybucket/directory0/directory1/myfile` will _not_ be included (the full path doesn't match any items in `include_regex`) * `gs://mybucket/directory1/excludedfile` will _not_ be included (the path matches an item in `exclude_regex`) If `include_regex` is left empty, it will match all files by default (this is equivalent to setting `include_regex: [".*"]`). Some other common use cases: * `{bucket_name: "mybucket", exclude_regex: [".*\.pdf"]}` will include all files in `mybucket` except for .pdf files * `{bucket_name: "mybucket", include_regex: ["directory/[^/]+"]}` will include all files directly under `gs://mybucket/directory/`, without matching across `/`
    /// </summary>
    public sealed class GooglePrivacyDlpV2CloudStorageRegexFileSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of a Cloud Storage bucket. Required.
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        [Input("excludeRegex")]
        private InputList<string>? _excludeRegex;

        /// <summary>
        /// A list of regular expressions matching file paths to exclude. All files in the bucket that match at least one of these regular expressions will be excluded from the scan. Regular expressions use RE2 [syntax](https://github.com/google/re2/wiki/Syntax); a guide can be found under the google/re2 repository on GitHub.
        /// </summary>
        public InputList<string> ExcludeRegex
        {
            get => _excludeRegex ?? (_excludeRegex = new InputList<string>());
            set => _excludeRegex = value;
        }

        [Input("includeRegex")]
        private InputList<string>? _includeRegex;

        /// <summary>
        /// A list of regular expressions matching file paths to include. All files in the bucket that match at least one of these regular expressions will be included in the set of files, except for those that also match an item in `exclude_regex`. Leaving this field empty will match all files by default (this is equivalent to including `.*` in the list). Regular expressions use RE2 [syntax](https://github.com/google/re2/wiki/Syntax); a guide can be found under the google/re2 repository on GitHub.
        /// </summary>
        public InputList<string> IncludeRegex
        {
            get => _includeRegex ?? (_includeRegex = new InputList<string>());
            set => _includeRegex = value;
        }

        public GooglePrivacyDlpV2CloudStorageRegexFileSetArgs()
        {
        }
        public static new GooglePrivacyDlpV2CloudStorageRegexFileSetArgs Empty => new GooglePrivacyDlpV2CloudStorageRegexFileSetArgs();
    }
}
