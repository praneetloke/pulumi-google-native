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
    /// Creates an autoscaler in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/v1:Autoscaler")]
    public partial class Autoscaler : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The configuration parameters for the autoscaling algorithm. You can define one or more signals for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization. If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
        /// </summary>
        [Output("autoscalingPolicy")]
        public Output<Outputs.AutoscalingPolicyResponse> AutoscalingPolicy { get; private set; } = null!;

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
        /// Type of the resource. Always compute#autoscaler for autoscalers.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Target recommended MIG size (number of instances) computed by autoscaler. Autoscaler calculates the recommended MIG size even when the autoscaling policy mode is different from ON. This field is empty when autoscaler is not connected to an existing managed instance group or autoscaler did not generate its prediction.
        /// </summary>
        [Output("recommendedSize")]
        public Output<int> RecommendedSize { get; private set; } = null!;

        /// <summary>
        /// URL of the region where the instance group resides (for autoscalers living in regional scope).
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// Status information of existing scaling schedules.
        /// </summary>
        [Output("scalingScheduleStatus")]
        public Output<Outputs.ScalingScheduleStatusResponse> ScalingScheduleStatus { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The status of the autoscaler configuration. Current set of possible values: - PENDING: Autoscaler backend hasn't read new/updated configuration. - DELETING: Configuration is being deleted. - ACTIVE: Configuration is acknowledged to be effective. Some warnings might be present in the statusDetails field. - ERROR: Configuration has errors. Actionable for users. Details are present in the statusDetails field. New values might be added in the future.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Human-readable details about the current state of the autoscaler. Read the documentation for Commonly returned status messages for examples of status messages you might encounter.
        /// </summary>
        [Output("statusDetails")]
        public Output<ImmutableArray<Outputs.AutoscalerStatusDetailsResponse>> StatusDetails { get; private set; } = null!;

        /// <summary>
        /// URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
        /// </summary>
        [Output("target")]
        public Output<string> Target { get; private set; } = null!;

        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Autoscaler resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Autoscaler(string name, AutoscalerArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:Autoscaler", name, args ?? new AutoscalerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Autoscaler(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:Autoscaler", name, null, MakeResourceOptions(options, id))
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
                    "zone",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Autoscaler resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Autoscaler Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Autoscaler(name, id, options);
        }
    }

    public sealed class AutoscalerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration parameters for the autoscaling algorithm. You can define one or more signals for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization. If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
        /// </summary>
        [Input("autoscalingPolicy")]
        public Input<Inputs.AutoscalingPolicyArgs>? AutoscalingPolicy { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public AutoscalerArgs()
        {
        }
        public static new AutoscalerArgs Empty => new AutoscalerArgs();
    }
}
