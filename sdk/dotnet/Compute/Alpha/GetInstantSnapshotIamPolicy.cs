// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetInstantSnapshotIamPolicy
    {
        /// <summary>
        /// Gets the access control policy for a resource. May be empty if no such policy or resource exists.
        /// </summary>
        public static Task<GetInstantSnapshotIamPolicyResult> InvokeAsync(GetInstantSnapshotIamPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstantSnapshotIamPolicyResult>("google-native:compute/alpha:getInstantSnapshotIamPolicy", args ?? new GetInstantSnapshotIamPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the access control policy for a resource. May be empty if no such policy or resource exists.
        /// </summary>
        public static Output<GetInstantSnapshotIamPolicyResult> Invoke(GetInstantSnapshotIamPolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstantSnapshotIamPolicyResult>("google-native:compute/alpha:getInstantSnapshotIamPolicy", args ?? new GetInstantSnapshotIamPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstantSnapshotIamPolicyArgs : global::Pulumi.InvokeArgs
    {
        [Input("optionsRequestedPolicyVersion")]
        public string? OptionsRequestedPolicyVersion { get; set; }

        [Input("project")]
        public string? Project { get; set; }

        [Input("resource", required: true)]
        public string Resource { get; set; } = null!;

        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetInstantSnapshotIamPolicyArgs()
        {
        }
        public static new GetInstantSnapshotIamPolicyArgs Empty => new GetInstantSnapshotIamPolicyArgs();
    }

    public sealed class GetInstantSnapshotIamPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("optionsRequestedPolicyVersion")]
        public Input<string>? OptionsRequestedPolicyVersion { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetInstantSnapshotIamPolicyInvokeArgs()
        {
        }
        public static new GetInstantSnapshotIamPolicyInvokeArgs Empty => new GetInstantSnapshotIamPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstantSnapshotIamPolicyResult
    {
        /// <summary>
        /// Specifies cloud audit logging configuration for this policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.AuditConfigResponse> AuditConfigs;
        /// <summary>
        /// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
        /// </summary>
        public readonly ImmutableArray<Outputs.BindingResponse> Bindings;
        /// <summary>
        /// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleResponse> Rules;
        /// <summary>
        /// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        /// </summary>
        public readonly int Version;

        [OutputConstructor]
        private GetInstantSnapshotIamPolicyResult(
            ImmutableArray<Outputs.AuditConfigResponse> auditConfigs,

            ImmutableArray<Outputs.BindingResponse> bindings,

            string etag,

            ImmutableArray<Outputs.RuleResponse> rules,

            int version)
        {
            AuditConfigs = auditConfigs;
            Bindings = bindings;
            Etag = etag;
            Rules = rules;
            Version = version;
        }
    }
}
