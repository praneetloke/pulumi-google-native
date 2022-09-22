// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudResourceManager.V3
{
    public static class GetTagValue
    {
        /// <summary>
        /// Retrieves a TagValue. This method will return `PERMISSION_DENIED` if the value does not exist or the user does not have permission to view it.
        /// </summary>
        public static Task<GetTagValueResult> InvokeAsync(GetTagValueArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTagValueResult>("google-native:cloudresourcemanager/v3:getTagValue", args ?? new GetTagValueArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a TagValue. This method will return `PERMISSION_DENIED` if the value does not exist or the user does not have permission to view it.
        /// </summary>
        public static Output<GetTagValueResult> Invoke(GetTagValueInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTagValueResult>("google-native:cloudresourcemanager/v3:getTagValue", args ?? new GetTagValueInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTagValueArgs : global::Pulumi.InvokeArgs
    {
        [Input("tagValueId", required: true)]
        public string TagValueId { get; set; } = null!;

        public GetTagValueArgs()
        {
        }
        public static new GetTagValueArgs Empty => new GetTagValueArgs();
    }

    public sealed class GetTagValueInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("tagValueId", required: true)]
        public Input<string> TagValueId { get; set; } = null!;

        public GetTagValueInvokeArgs()
        {
        }
        public static new GetTagValueInvokeArgs Empty => new GetTagValueInvokeArgs();
    }


    [OutputType]
    public sealed class GetTagValueResult
    {
        /// <summary>
        /// Creation time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. User-assigned description of the TagValue. Must not exceed 256 characters. Read-write.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. Entity tag which users can pass to prevent race conditions. This field is always set in server responses. See UpdateTagValueRequest for details.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Immutable. Resource name for TagValue in the format `tagValues/456`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Namespaced name of the TagValue. Now only supported in the format `{organization_id}/{tag_key_short_name}/{short_name}`. Other formats will be supported when we add non-org parented tags.
        /// </summary>
        public readonly string NamespacedName;
        /// <summary>
        /// Immutable. The resource name of the new TagValue's parent TagKey. Must be of the form `tagKeys/{tag_key_id}`.
        /// </summary>
        public readonly string Parent;
        /// <summary>
        /// Immutable. User-assigned short name for TagValue. The short name should be unique for TagValues within the same parent TagKey. The short name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
        /// </summary>
        public readonly string ShortName;
        /// <summary>
        /// Update time.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetTagValueResult(
            string createTime,

            string description,

            string etag,

            string name,

            string namespacedName,

            string parent,

            string shortName,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            Etag = etag;
            Name = name;
            NamespacedName = namespacedName;
            Parent = parent;
            ShortName = shortName;
            UpdateTime = updateTime;
        }
    }
}
