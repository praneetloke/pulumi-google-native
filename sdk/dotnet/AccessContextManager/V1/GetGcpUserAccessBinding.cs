// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AccessContextManager.V1
{
    public static class GetGcpUserAccessBinding
    {
        /// <summary>
        /// Gets the GcpUserAccessBinding with the given name.
        /// </summary>
        public static Task<GetGcpUserAccessBindingResult> InvokeAsync(GetGcpUserAccessBindingArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGcpUserAccessBindingResult>("google-native:accesscontextmanager/v1:getGcpUserAccessBinding", args ?? new GetGcpUserAccessBindingArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the GcpUserAccessBinding with the given name.
        /// </summary>
        public static Output<GetGcpUserAccessBindingResult> Invoke(GetGcpUserAccessBindingInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGcpUserAccessBindingResult>("google-native:accesscontextmanager/v1:getGcpUserAccessBinding", args ?? new GetGcpUserAccessBindingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGcpUserAccessBindingArgs : global::Pulumi.InvokeArgs
    {
        [Input("gcpUserAccessBindingId", required: true)]
        public string GcpUserAccessBindingId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        public GetGcpUserAccessBindingArgs()
        {
        }
        public static new GetGcpUserAccessBindingArgs Empty => new GetGcpUserAccessBindingArgs();
    }

    public sealed class GetGcpUserAccessBindingInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("gcpUserAccessBindingId", required: true)]
        public Input<string> GcpUserAccessBindingId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        public GetGcpUserAccessBindingInvokeArgs()
        {
        }
        public static new GetGcpUserAccessBindingInvokeArgs Empty => new GetGcpUserAccessBindingInvokeArgs();
    }


    [OutputType]
    public sealed class GetGcpUserAccessBindingResult
    {
        /// <summary>
        /// Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
        /// </summary>
        public readonly ImmutableArray<string> AccessLevels;
        /// <summary>
        /// Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the [G Suite Directory API's Groups resource] (https://developers.google.com/admin-sdk/directory/v1/reference/groups#resource). If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
        /// </summary>
        public readonly string GroupKey;
        /// <summary>
        /// Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by [RFC 3986 Section 2.3](https://tools.ietf.org/html/rfc3986#section-2.3)). Should not be specified by the client during creation. Example: "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetGcpUserAccessBindingResult(
            ImmutableArray<string> accessLevels,

            string groupKey,

            string name)
        {
            AccessLevels = accessLevels;
            GroupKey = groupKey;
            Name = name;
        }
    }
}
