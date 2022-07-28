// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1
{
    /// <summary>
    /// Creates a new EndpointPolicy in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:networkservices/v1:EndpointPolicy")]
    public partial class EndpointPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. This field specifies the URL of AuthorizationPolicy resource that applies authorization policies to the inbound traffic at the matched endpoints. Refer to Authorization. If this field is not specified, authorization is disabled(no authz checks) for this endpoint.
        /// </summary>
        [Output("authorizationPolicy")]
        public Output<string> AuthorizationPolicy { get; private set; } = null!;

        /// <summary>
        /// Optional. A URL referring to a ClientTlsPolicy resource. ClientTlsPolicy can be set to specify the authentication for traffic from the proxy to the actual endpoints. More specifically, it is applied to the outgoing traffic from the proxy to the endpoint. This is typically used for sidecar model where the proxy identifies itself as endpoint to the control plane, with the connection between sidecar and endpoint requiring authentication. If this field is not set, authentication is disabled(open). Applicable only when EndpointPolicyType is SIDECAR_PROXY.
        /// </summary>
        [Output("clientTlsPolicy")]
        public Output<string> ClientTlsPolicy { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the resource was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// A matcher that selects endpoints to which the policies should be applied.
        /// </summary>
        [Output("endpointMatcher")]
        public Output<Outputs.EndpointMatcherResponse> EndpointMatcher { get; private set; } = null!;

        /// <summary>
        /// Required. Short name of the EndpointPolicy resource to be created. E.g. "CustomECS".
        /// </summary>
        [Output("endpointPolicyId")]
        public Output<string> EndpointPolicyId { get; private set; } = null!;

        /// <summary>
        /// Optional. Set of label tags associated with the EndpointPolicy resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Name of the EndpointPolicy resource. It matches pattern `projects/{project}/locations/global/endpointPolicies/{endpoint_policy}`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. A URL referring to ServerTlsPolicy resource. ServerTlsPolicy is used to determine the authentication policy to be applied to terminate the inbound traffic at the identified backends. If this field is not set, authentication is disabled(open) for this endpoint.
        /// </summary>
        [Output("serverTlsPolicy")]
        public Output<string> ServerTlsPolicy { get; private set; } = null!;

        /// <summary>
        /// Optional. Port selector for the (matched) endpoints. If no port selector is provided, the matched config is applied to all ports.
        /// </summary>
        [Output("trafficPortSelector")]
        public Output<Outputs.TrafficPortSelectorResponse> TrafficPortSelector { get; private set; } = null!;

        /// <summary>
        /// The type of endpoint policy. This is primarily used to validate the configuration.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the resource was updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a EndpointPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EndpointPolicy(string name, EndpointPolicyArgs args, CustomResourceOptions? options = null)
            : base("google-native:networkservices/v1:EndpointPolicy", name, args ?? new EndpointPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EndpointPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:networkservices/v1:EndpointPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "endpointPolicyId",
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
        /// Get an existing EndpointPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EndpointPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EndpointPolicy(name, id, options);
        }
    }

    public sealed class EndpointPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. This field specifies the URL of AuthorizationPolicy resource that applies authorization policies to the inbound traffic at the matched endpoints. Refer to Authorization. If this field is not specified, authorization is disabled(no authz checks) for this endpoint.
        /// </summary>
        [Input("authorizationPolicy")]
        public Input<string>? AuthorizationPolicy { get; set; }

        /// <summary>
        /// Optional. A URL referring to a ClientTlsPolicy resource. ClientTlsPolicy can be set to specify the authentication for traffic from the proxy to the actual endpoints. More specifically, it is applied to the outgoing traffic from the proxy to the endpoint. This is typically used for sidecar model where the proxy identifies itself as endpoint to the control plane, with the connection between sidecar and endpoint requiring authentication. If this field is not set, authentication is disabled(open). Applicable only when EndpointPolicyType is SIDECAR_PROXY.
        /// </summary>
        [Input("clientTlsPolicy")]
        public Input<string>? ClientTlsPolicy { get; set; }

        /// <summary>
        /// Optional. A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A matcher that selects endpoints to which the policies should be applied.
        /// </summary>
        [Input("endpointMatcher", required: true)]
        public Input<Inputs.EndpointMatcherArgs> EndpointMatcher { get; set; } = null!;

        /// <summary>
        /// Required. Short name of the EndpointPolicy resource to be created. E.g. "CustomECS".
        /// </summary>
        [Input("endpointPolicyId", required: true)]
        public Input<string> EndpointPolicyId { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Set of label tags associated with the EndpointPolicy resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the EndpointPolicy resource. It matches pattern `projects/{project}/locations/global/endpointPolicies/{endpoint_policy}`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. A URL referring to ServerTlsPolicy resource. ServerTlsPolicy is used to determine the authentication policy to be applied to terminate the inbound traffic at the identified backends. If this field is not set, authentication is disabled(open) for this endpoint.
        /// </summary>
        [Input("serverTlsPolicy")]
        public Input<string>? ServerTlsPolicy { get; set; }

        /// <summary>
        /// Optional. Port selector for the (matched) endpoints. If no port selector is provided, the matched config is applied to all ports.
        /// </summary>
        [Input("trafficPortSelector")]
        public Input<Inputs.TrafficPortSelectorArgs>? TrafficPortSelector { get; set; }

        /// <summary>
        /// The type of endpoint policy. This is primarily used to validate the configuration.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.NetworkServices.V1.EndpointPolicyType> Type { get; set; } = null!;

        public EndpointPolicyArgs()
        {
        }
        public static new EndpointPolicyArgs Empty => new EndpointPolicyArgs();
    }
}
