// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQueryReservation.V1
{
    /// <summary>
    /// Creates a new capacity commitment resource.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:bigqueryreservation/v1:CapacityCommitment")]
    public partial class CapacityCommitment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The optional capacity commitment ID. Capacity commitment name will be generated automatically if this field is empty. This field must only contain lower case alphanumeric characters or dashes. The first and last character cannot be a dash. Max length is 64 characters. NOTE: this ID won't be kept if the capacity commitment is split or merged.
        /// </summary>
        [Output("capacityCommitmentId")]
        public Output<string?> CapacityCommitmentId { get; private set; } = null!;

        /// <summary>
        /// The end of the current commitment period. It is applicable only for ACTIVE capacity commitments.
        /// </summary>
        [Output("commitmentEndTime")]
        public Output<string> CommitmentEndTime { get; private set; } = null!;

        /// <summary>
        /// The start of the current commitment period. It is applicable only for ACTIVE capacity commitments.
        /// </summary>
        [Output("commitmentStartTime")]
        public Output<string> CommitmentStartTime { get; private set; } = null!;

        /// <summary>
        /// Edition of the capacity commitment.
        /// </summary>
        [Output("edition")]
        public Output<string> Edition { get; private set; } = null!;

        /// <summary>
        /// If true, fail the request if another project in the organization has a capacity commitment.
        /// </summary>
        [Output("enforceSingleAdminProjectPerOrg")]
        public Output<bool?> EnforceSingleAdminProjectPerOrg { get; private set; } = null!;

        /// <summary>
        /// For FAILED commitment plan, provides the reason of failure.
        /// </summary>
        [Output("failureStatus")]
        public Output<Outputs.StatusResponse> FailureStatus { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Applicable only for commitments located within one of the BigQuery multi-regions (US or EU). If set to true, this commitment is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this commitment is placed in the organization's default region. NOTE: this is a preview feature. Project must be allow-listed in order to set this field.
        /// </summary>
        [Output("multiRegionAuxiliary")]
        public Output<bool> MultiRegionAuxiliary { get; private set; } = null!;

        /// <summary>
        /// The resource name of the capacity commitment, e.g., `projects/myproject/locations/US/capacityCommitments/123` The commitment_id must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Capacity commitment commitment plan.
        /// </summary>
        [Output("plan")]
        public Output<string> Plan { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The plan this capacity commitment is converted to after commitment_end_time passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable for ANNUAL and TRIAL commitments.
        /// </summary>
        [Output("renewalPlan")]
        public Output<string> RenewalPlan { get; private set; } = null!;

        /// <summary>
        /// Number of slots in this commitment.
        /// </summary>
        [Output("slotCount")]
        public Output<string> SlotCount { get; private set; } = null!;

        /// <summary>
        /// State of the commitment.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a CapacityCommitment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CapacityCommitment(string name, CapacityCommitmentArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:bigqueryreservation/v1:CapacityCommitment", name, args ?? new CapacityCommitmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CapacityCommitment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:bigqueryreservation/v1:CapacityCommitment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CapacityCommitment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CapacityCommitment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CapacityCommitment(name, id, options);
        }
    }

    public sealed class CapacityCommitmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The optional capacity commitment ID. Capacity commitment name will be generated automatically if this field is empty. This field must only contain lower case alphanumeric characters or dashes. The first and last character cannot be a dash. Max length is 64 characters. NOTE: this ID won't be kept if the capacity commitment is split or merged.
        /// </summary>
        [Input("capacityCommitmentId")]
        public Input<string>? CapacityCommitmentId { get; set; }

        /// <summary>
        /// Edition of the capacity commitment.
        /// </summary>
        [Input("edition")]
        public Input<Pulumi.GoogleNative.BigQueryReservation.V1.CapacityCommitmentEdition>? Edition { get; set; }

        /// <summary>
        /// If true, fail the request if another project in the organization has a capacity commitment.
        /// </summary>
        [Input("enforceSingleAdminProjectPerOrg")]
        public Input<bool>? EnforceSingleAdminProjectPerOrg { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Applicable only for commitments located within one of the BigQuery multi-regions (US or EU). If set to true, this commitment is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this commitment is placed in the organization's default region. NOTE: this is a preview feature. Project must be allow-listed in order to set this field.
        /// </summary>
        [Input("multiRegionAuxiliary")]
        public Input<bool>? MultiRegionAuxiliary { get; set; }

        /// <summary>
        /// Capacity commitment commitment plan.
        /// </summary>
        [Input("plan")]
        public Input<Pulumi.GoogleNative.BigQueryReservation.V1.CapacityCommitmentPlan>? Plan { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The plan this capacity commitment is converted to after commitment_end_time passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable for ANNUAL and TRIAL commitments.
        /// </summary>
        [Input("renewalPlan")]
        public Input<Pulumi.GoogleNative.BigQueryReservation.V1.CapacityCommitmentRenewalPlan>? RenewalPlan { get; set; }

        /// <summary>
        /// Number of slots in this commitment.
        /// </summary>
        [Input("slotCount")]
        public Input<string>? SlotCount { get; set; }

        public CapacityCommitmentArgs()
        {
        }
        public static new CapacityCommitmentArgs Empty => new CapacityCommitmentArgs();
    }
}
