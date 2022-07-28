// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// Additional disk params.
    /// </summary>
    public sealed class DiskParamsArgs : global::Pulumi.ResourceArgs
    {
        [Input("resourceManagerTags")]
        private InputMap<string>? _resourceManagerTags;

        /// <summary>
        /// Resource manager tags to be bound to the disk. Tag keys and values have the same definition as resource manager tags. Keys must be in the format `tagKeys/{tag_key_id}`, and values are in the format `tagValues/456`. The field is ignored (both PUT &amp; PATCH) when empty.
        /// </summary>
        public InputMap<string> ResourceManagerTags
        {
            get => _resourceManagerTags ?? (_resourceManagerTags = new InputMap<string>());
            set => _resourceManagerTags = value;
        }

        public DiskParamsArgs()
        {
        }
        public static new DiskParamsArgs Empty => new DiskParamsArgs();
    }
}
