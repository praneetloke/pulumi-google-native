// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    /// <summary>
    /// Creates a commitment in the specified project using the data included in the request.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/v1:RegionCommitment")]
    public partial class RegionCommitment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether to enable automatic renewal for the commitment. The default value is false if not specified. The field can be updated until the day of the commitment expiration at 12:00am PST. If the field is set to true, the commitment will be automatically renewed for either one or three years according to the terms of the existing commitment.
        /// </summary>
        [Output("autoRenew")]
        public Output<bool> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
        /// </summary>
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Commitment end time in RFC3339 text format.
        /// </summary>
        [Output("endTimestamp")]
        public Output<string> EndTimestamp { get; private set; } = null!;

        /// <summary>
        /// Type of the resource. Always compute#commitment for commitments.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The license specification required as part of a license commitment.
        /// </summary>
        [Output("licenseResource")]
        public Output<Outputs.LicenseResourceCommitmentResponse> LicenseResource { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
        /// </summary>
        [Output("plan")]
        public Output<string> Plan { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// List of reservations in this commitment.
        /// </summary>
        [Output("reservations")]
        public Output<ImmutableArray<Outputs.ReservationResponse>> Reservations { get; private set; } = null!;

        /// <summary>
        /// A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
        /// </summary>
        [Output("resources")]
        public Output<ImmutableArray<Outputs.ResourceCommitmentResponse>> Resources { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Commitment start time in RFC3339 text format.
        /// </summary>
        [Output("startTimestamp")]
        public Output<string> StartTimestamp { get; private set; } = null!;

        /// <summary>
        /// Status of the commitment with regards to eventual expiration (each commitment has an end date defined). One of the following values: NOT_YET_ACTIVE, ACTIVE, EXPIRED.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// An optional, human-readable explanation of the status.
        /// </summary>
        [Output("statusMessage")]
        public Output<string> StatusMessage { get; private set; } = null!;

        /// <summary>
        /// The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a RegionCommitment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionCommitment(string name, RegionCommitmentArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:RegionCommitment", name, args ?? new RegionCommitmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegionCommitment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:RegionCommitment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "project",
                    "region",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RegionCommitment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionCommitment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RegionCommitment(name, id, options);
        }
    }

    public sealed class RegionCommitmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to enable automatic renewal for the commitment. The default value is false if not specified. The field can be updated until the day of the commitment expiration at 12:00am PST. If the field is set to true, the commitment will be automatically renewed for either one or three years according to the terms of the existing commitment.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
        /// </summary>
        [Input("category")]
        public Input<Pulumi.GoogleNative.Compute.V1.RegionCommitmentCategory>? Category { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The license specification required as part of a license commitment.
        /// </summary>
        [Input("licenseResource")]
        public Input<Inputs.LicenseResourceCommitmentArgs>? LicenseResource { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
        /// </summary>
        [Input("plan")]
        public Input<Pulumi.GoogleNative.Compute.V1.RegionCommitmentPlan>? Plan { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        [Input("reservations")]
        private InputList<Inputs.ReservationArgs>? _reservations;

        /// <summary>
        /// List of reservations in this commitment.
        /// </summary>
        public InputList<Inputs.ReservationArgs> Reservations
        {
            get => _reservations ?? (_reservations = new InputList<Inputs.ReservationArgs>());
            set => _reservations = value;
        }

        [Input("resources")]
        private InputList<Inputs.ResourceCommitmentArgs>? _resources;

        /// <summary>
        /// A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
        /// </summary>
        public InputList<Inputs.ResourceCommitmentArgs> Resources
        {
            get => _resources ?? (_resources = new InputList<Inputs.ResourceCommitmentArgs>());
            set => _resources = value;
        }

        /// <summary>
        /// The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Compute.V1.RegionCommitmentType>? Type { get; set; }

        public RegionCommitmentArgs()
        {
        }
        public static new RegionCommitmentArgs Empty => new RegionCommitmentArgs();
    }
}
