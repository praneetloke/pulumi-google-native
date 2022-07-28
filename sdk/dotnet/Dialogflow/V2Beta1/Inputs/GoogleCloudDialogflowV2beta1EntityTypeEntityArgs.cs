// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1.Inputs
{

    /// <summary>
    /// An **entity entry** for an associated entity type.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1EntityTypeEntityArgs : global::Pulumi.ResourceArgs
    {
        [Input("synonyms", required: true)]
        private InputList<string>? _synonyms;

        /// <summary>
        /// A collection of value synonyms. For example, if the entity type is *vegetable*, and `value` is *scallions*, a synonym could be *green onions*. For `KIND_LIST` entity types: * This collection must contain exactly one synonym equal to `value`.
        /// </summary>
        public InputList<string> Synonyms
        {
            get => _synonyms ?? (_synonyms = new InputList<string>());
            set => _synonyms = value;
        }

        /// <summary>
        /// The primary value associated with this entity entry. For example, if the entity type is *vegetable*, the value could be *scallions*. For `KIND_MAP` entity types: * A reference value to be used in place of synonyms. For `KIND_LIST` entity types: * A string that can contain references to other entity types (with or without aliases).
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public GoogleCloudDialogflowV2beta1EntityTypeEntityArgs()
        {
        }
        public static new GoogleCloudDialogflowV2beta1EntityTypeEntityArgs Empty => new GoogleCloudDialogflowV2beta1EntityTypeEntityArgs();
    }
}
