// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudIdentity.V1Beta1
{
    public static class GetGroup
    {
        /// <summary>
        /// Retrieves a `Group`.
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("google-native:cloudidentity/v1beta1:getGroup", args ?? new GetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a `Group`.
        /// </summary>
        public static Output<GetGroupResult> Invoke(GetGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupResult>("google-native:cloudidentity/v1beta1:getGroup", args ?? new GetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("groupId", required: true)]
        public string GroupId { get; set; } = null!;

        public GetGroupArgs()
        {
        }
        public static new GetGroupArgs Empty => new GetGroupArgs();
    }

    public sealed class GetGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        public GetGroupInvokeArgs()
        {
        }
        public static new GetGroupInvokeArgs Empty => new GetGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        /// <summary>
        /// Additional group keys associated with the Group.
        /// </summary>
        public readonly ImmutableArray<Outputs.EntityKeyResponse> AdditionalGroupKeys;
        /// <summary>
        /// The time when the `Group` was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// An extended description to help users determine the purpose of a `Group`. Must not be longer than 4,096 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The display name of the `Group`.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Optional. Dynamic group metadata like queries and status.
        /// </summary>
        public readonly Outputs.DynamicGroupMetadataResponse DynamicGroupMetadata;
        /// <summary>
        /// The `EntityKey` of the `Group`.
        /// </summary>
        public readonly Outputs.EntityKeyResponse GroupKey;
        /// <summary>
        /// One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value. Google Groups are the default type of group and have a label with a key of `cloudidentity.googleapis.com/groups.discussion_forum` and an empty value. Existing Google Groups can have an additional label with a key of `cloudidentity.googleapis.com/groups.security` and an empty value added to them. **This is an immutable change and the security label cannot be removed once added.** Dynamic groups have a label with a key of `cloudidentity.googleapis.com/groups.dynamic`. Identity-mapped groups for Cloud Search have a label with a key of `system/groups/external` and an empty value.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Group`. Shall be of the form `groups/{group_id}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Immutable. The resource name of the entity under which this `Group` resides in the Cloud Identity resource hierarchy. Must be of the form `identitysources/{identity_source}` for external [identity-mapped groups](https://support.google.com/a/answer/9039510) or `customers/{customer_id}` for Google Groups. The `customer_id` must begin with "C" (for example, 'C046psxkn'). [Find your customer ID.] (https://support.google.com/cloudidentity/answer/10070793)
        /// </summary>
        public readonly string Parent;
        /// <summary>
        /// Optional. The POSIX groups associated with the `Group`.
        /// </summary>
        public readonly ImmutableArray<Outputs.PosixGroupResponse> PosixGroups;
        /// <summary>
        /// The time when the `Group` was last updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetGroupResult(
            ImmutableArray<Outputs.EntityKeyResponse> additionalGroupKeys,

            string createTime,

            string description,

            string displayName,

            Outputs.DynamicGroupMetadataResponse dynamicGroupMetadata,

            Outputs.EntityKeyResponse groupKey,

            ImmutableDictionary<string, string> labels,

            string name,

            string parent,

            ImmutableArray<Outputs.PosixGroupResponse> posixGroups,

            string updateTime)
        {
            AdditionalGroupKeys = additionalGroupKeys;
            CreateTime = createTime;
            Description = description;
            DisplayName = displayName;
            DynamicGroupMetadata = dynamicGroupMetadata;
            GroupKey = groupKey;
            Labels = labels;
            Name = name;
            Parent = parent;
            PosixGroups = posixGroups;
            UpdateTime = updateTime;
        }
    }
}
