// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// A set of instance tags.
    /// </summary>
    public sealed class TagsArgs : global::Pulumi.ResourceArgs
    {
        [Input("items")]
        private InputList<string>? _items;

        /// <summary>
        /// An array of tags. Each tag must be 1-63 characters long, and comply with RFC1035.
        /// </summary>
        public InputList<string> Items
        {
            get => _items ?? (_items = new InputList<string>());
            set => _items = value;
        }

        public TagsArgs()
        {
        }
        public static new TagsArgs Empty => new TagsArgs();
    }
}
