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
    /// A metadata key/value entry.
    /// </summary>
    public sealed class MetadataArgs : global::Pulumi.ResourceArgs
    {
        [Input("items")]
        private InputList<Inputs.MetadataItemsItemArgs>? _items;

        /// <summary>
        /// Array of key/value pairs. The total size of all keys and values must be less than 512 KB.
        /// </summary>
        public InputList<Inputs.MetadataItemsItemArgs> Items
        {
            get => _items ?? (_items = new InputList<Inputs.MetadataItemsItemArgs>());
            set => _items = value;
        }

        public MetadataArgs()
        {
        }
        public static new MetadataArgs Empty => new MetadataArgs();
    }
}
