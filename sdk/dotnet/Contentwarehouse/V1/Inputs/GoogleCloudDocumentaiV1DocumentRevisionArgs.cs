// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contentwarehouse.V1.Inputs
{

    /// <summary>
    /// Contains past or forward revisions of this document.
    /// </summary>
    public sealed class GoogleCloudDocumentaiV1DocumentRevisionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If the change was made by a person specify the name or id of that person.
        /// </summary>
        [Input("agent")]
        public Input<string>? Agent { get; set; }

        /// <summary>
        /// The time that the revision was created, internally generated by doc proto storage at the time of create.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Human Review information of this revision.
        /// </summary>
        [Input("humanReview")]
        public Input<Inputs.GoogleCloudDocumentaiV1DocumentRevisionHumanReviewArgs>? HumanReview { get; set; }

        /// <summary>
        /// Id of the revision, internally generated by doc proto storage. Unique within the context of the document.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("parent")]
        private InputList<int>? _parent;

        /// <summary>
        /// The revisions that this revision is based on. This can include one or more parent (when documents are merged.) This field represents the index into the `revisions` field.
        /// </summary>
        public InputList<int> Parent
        {
            get => _parent ?? (_parent = new InputList<int>());
            set => _parent = value;
        }

        [Input("parentIds")]
        private InputList<string>? _parentIds;

        /// <summary>
        /// The revisions that this revision is based on. Must include all the ids that have anything to do with this revision - eg. there are `provenance.parent.revision` fields that index into this field.
        /// </summary>
        public InputList<string> ParentIds
        {
            get => _parentIds ?? (_parentIds = new InputList<string>());
            set => _parentIds = value;
        }

        /// <summary>
        /// If the annotation was made by processor identify the processor by its resource name.
        /// </summary>
        [Input("processor")]
        public Input<string>? Processor { get; set; }

        public GoogleCloudDocumentaiV1DocumentRevisionArgs()
        {
        }
        public static new GoogleCloudDocumentaiV1DocumentRevisionArgs Empty => new GoogleCloudDocumentaiV1DocumentRevisionArgs();
    }
}