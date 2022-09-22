// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudSupport.V2Beta
{
    public static class GetCase
    {
        /// <summary>
        /// Retrieve the specified case.
        /// </summary>
        public static Task<GetCaseResult> InvokeAsync(GetCaseArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCaseResult>("google-native:cloudsupport/v2beta:getCase", args ?? new GetCaseArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieve the specified case.
        /// </summary>
        public static Output<GetCaseResult> Invoke(GetCaseInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCaseResult>("google-native:cloudsupport/v2beta:getCase", args ?? new GetCaseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCaseArgs : global::Pulumi.InvokeArgs
    {
        [Input("caseId", required: true)]
        public string CaseId { get; set; } = null!;

        [Input("v2betaId1", required: true)]
        public string V2betaId1 { get; set; } = null!;

        [Input("v2betumId", required: true)]
        public string V2betumId { get; set; } = null!;

        public GetCaseArgs()
        {
        }
        public static new GetCaseArgs Empty => new GetCaseArgs();
    }

    public sealed class GetCaseInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("caseId", required: true)]
        public Input<string> CaseId { get; set; } = null!;

        [Input("v2betaId1", required: true)]
        public Input<string> V2betaId1 { get; set; } = null!;

        [Input("v2betumId", required: true)]
        public Input<string> V2betumId { get; set; } = null!;

        public GetCaseInvokeArgs()
        {
        }
        public static new GetCaseInvokeArgs Empty => new GetCaseInvokeArgs();
    }


    [OutputType]
    public sealed class GetCaseResult
    {
        /// <summary>
        /// The issue classification applicable to this case.
        /// </summary>
        public readonly Outputs.CaseClassificationResponse Classification;
        /// <summary>
        /// The time this case was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The user who created the case. Note: The name and email will be obfuscated if the case was created by Google Support.
        /// </summary>
        public readonly Outputs.ActorResponse Creator;
        /// <summary>
        /// A broad description of the issue.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The short summary of the issue reported in this case.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Whether the case is currently escalated.
        /// </summary>
        public readonly bool Escalated;
        /// <summary>
        /// The resource name for the case.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The priority of this case. If this is set, do not set severity.
        /// </summary>
        public readonly string Priority;
        /// <summary>
        /// The severity of this case. Deprecated. Use priority instead.
        /// </summary>
        public readonly string Severity;
        /// <summary>
        /// The current status of the support case.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The email addresses to receive updates on this case.
        /// </summary>
        public readonly ImmutableArray<string> SubscriberEmailAddresses;
        /// <summary>
        /// Whether this case was created for internal API testing and should not be acted on by the support team.
        /// </summary>
        public readonly bool TestCase;
        /// <summary>
        /// The timezone of the user who created the support case. It should be in a format IANA recognizes: https://www.iana.org/time-zones. There is no additional validation done by the API.
        /// </summary>
        public readonly string TimeZone;
        /// <summary>
        /// The time this case was last updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetCaseResult(
            Outputs.CaseClassificationResponse classification,

            string createTime,

            Outputs.ActorResponse creator,

            string description,

            string displayName,

            bool escalated,

            string name,

            string priority,

            string severity,

            string state,

            ImmutableArray<string> subscriberEmailAddresses,

            bool testCase,

            string timeZone,

            string updateTime)
        {
            Classification = classification;
            CreateTime = createTime;
            Creator = creator;
            Description = description;
            DisplayName = displayName;
            Escalated = escalated;
            Name = name;
            Priority = priority;
            Severity = severity;
            State = state;
            SubscriberEmailAddresses = subscriberEmailAddresses;
            TestCase = testCase;
            TimeZone = timeZone;
            UpdateTime = updateTime;
        }
    }
}
