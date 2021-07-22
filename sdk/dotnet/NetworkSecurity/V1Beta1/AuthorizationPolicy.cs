// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkSecurity.V1Beta1
{
    /// <summary>
    /// Creates a new AuthorizationPolicy in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:networksecurity/v1beta1:AuthorizationPolicy")]
    public partial class AuthorizationPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The action to take when a rule match is found. Possible values are "ALLOW" or "DENY".
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the resource was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Free-text description of the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Optional. Set of label tags associated with the AuthorizationPolicy resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Name of the AuthorizationPolicy resource. It matches pattern `projects/{project}/locations/{location}/authorizationPolicies/`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. List of rules to match. Note that at least one of the rules must match in order for the action specified in the 'action' field to be taken. A rule is a match if there is a matching source and destination. If left blank, the action specified in the `action` field will be applied on every request.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.RuleResponse>> Rules { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the resource was updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a AuthorizationPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthorizationPolicy(string name, AuthorizationPolicyArgs args, CustomResourceOptions? options = null)
            : base("google-native:networksecurity/v1beta1:AuthorizationPolicy", name, args ?? new AuthorizationPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthorizationPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:networksecurity/v1beta1:AuthorizationPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AuthorizationPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthorizationPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AuthorizationPolicy(name, id, options);
        }
    }

    public sealed class AuthorizationPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to take when a rule match is found. Possible values are "ALLOW" or "DENY".
        /// </summary>
        [Input("action", required: true)]
        public Input<Pulumi.GoogleNative.NetworkSecurity.V1Beta1.AuthorizationPolicyAction> Action { get; set; } = null!;

        [Input("authorizationPolicyId", required: true)]
        public Input<string> AuthorizationPolicyId { get; set; } = null!;

        /// <summary>
        /// Optional. Free-text description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Set of label tags associated with the AuthorizationPolicy resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the AuthorizationPolicy resource. It matches pattern `projects/{project}/locations/{location}/authorizationPolicies/`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("rules")]
        private InputList<Inputs.RuleArgs>? _rules;

        /// <summary>
        /// Optional. List of rules to match. Note that at least one of the rules must match in order for the action specified in the 'action' field to be taken. A rule is a match if there is a matching source and destination. If left blank, the action specified in the `action` field will be applied on every request.
        /// </summary>
        public InputList<Inputs.RuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.RuleArgs>());
            set => _rules = value;
        }

        public AuthorizationPolicyArgs()
        {
        }
    }
}
