// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    /// <summary>
    /// CreateSecurityAction creates a SecurityAction.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:SecurityAction")]
    public partial class SecurityAction : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Allow a request through if it matches this SecurityAction.
        /// </summary>
        [Output("allow")]
        public Output<Outputs.GoogleCloudApigeeV1SecurityActionAllowResponse> Allow { get; private set; } = null!;

        /// <summary>
        /// A valid SecurityAction must contain at least one condition.
        /// </summary>
        [Output("conditionConfig")]
        public Output<Outputs.GoogleCloudApigeeV1SecurityActionConditionConfigResponse> ConditionConfig { get; private set; } = null!;

        /// <summary>
        /// The create time for this SecurityAction.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Deny a request through if it matches this SecurityAction.
        /// </summary>
        [Output("deny")]
        public Output<Outputs.GoogleCloudApigeeV1SecurityActionDenyResponse> Deny { get; private set; } = null!;

        /// <summary>
        /// Optional. An optional user provided description of the SecurityAction.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("environmentId")]
        public Output<string> EnvironmentId { get; private set; } = null!;

        /// <summary>
        /// The expiration for this SecurityAction.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// Flag a request through if it matches this SecurityAction.
        /// </summary>
        [Output("flag")]
        public Output<Outputs.GoogleCloudApigeeV1SecurityActionFlagResponse> Flag { get; private set; } = null!;

        /// <summary>
        /// Immutable. This field is ignored during creation as per AIP-133. Please set the `security_action_id` field in the CreateSecurityActionRequest when creating a new SecurityAction. Format: organizations/{org}/environments/{env}/securityActions/{security_action}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// Required. The ID to use for the SecurityAction, which will become the final component of the action's resource name. This value should be 0-61 characters, and valid format is (^[a-z]([a-z0-9-]{​0,61}[a-z0-9])?$).
        /// </summary>
        [Output("securityActionId")]
        public Output<string> SecurityActionId { get; private set; } = null!;

        /// <summary>
        /// Only an ENABLED SecurityAction is enforced. An ENABLED SecurityAction past its expiration time will not be enforced.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Input only. The TTL for this SecurityAction.
        /// </summary>
        [Output("ttl")]
        public Output<string> Ttl { get; private set; } = null!;

        /// <summary>
        /// The update time for this SecurityAction. This reflects when this SecurityAction changed states.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityAction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityAction(string name, SecurityActionArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:SecurityAction", name, args ?? new SecurityActionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityAction(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:SecurityAction", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "environmentId",
                    "organizationId",
                    "securityActionId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecurityAction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityAction Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SecurityAction(name, id, options);
        }
    }

    public sealed class SecurityActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow a request through if it matches this SecurityAction.
        /// </summary>
        [Input("allow")]
        public Input<Inputs.GoogleCloudApigeeV1SecurityActionAllowArgs>? Allow { get; set; }

        /// <summary>
        /// A valid SecurityAction must contain at least one condition.
        /// </summary>
        [Input("conditionConfig", required: true)]
        public Input<Inputs.GoogleCloudApigeeV1SecurityActionConditionConfigArgs> ConditionConfig { get; set; } = null!;

        /// <summary>
        /// Deny a request through if it matches this SecurityAction.
        /// </summary>
        [Input("deny")]
        public Input<Inputs.GoogleCloudApigeeV1SecurityActionDenyArgs>? Deny { get; set; }

        /// <summary>
        /// Optional. An optional user provided description of the SecurityAction.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        /// <summary>
        /// The expiration for this SecurityAction.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        /// <summary>
        /// Flag a request through if it matches this SecurityAction.
        /// </summary>
        [Input("flag")]
        public Input<Inputs.GoogleCloudApigeeV1SecurityActionFlagArgs>? Flag { get; set; }

        /// <summary>
        /// Immutable. This field is ignored during creation as per AIP-133. Please set the `security_action_id` field in the CreateSecurityActionRequest when creating a new SecurityAction. Format: organizations/{org}/environments/{env}/securityActions/{security_action}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        /// <summary>
        /// Required. The ID to use for the SecurityAction, which will become the final component of the action's resource name. This value should be 0-61 characters, and valid format is (^[a-z]([a-z0-9-]{​0,61}[a-z0-9])?$).
        /// </summary>
        [Input("securityActionId", required: true)]
        public Input<string> SecurityActionId { get; set; } = null!;

        /// <summary>
        /// Only an ENABLED SecurityAction is enforced. An ENABLED SecurityAction past its expiration time will not be enforced.
        /// </summary>
        [Input("state", required: true)]
        public Input<Pulumi.GoogleNative.Apigee.V1.SecurityActionState> State { get; set; } = null!;

        /// <summary>
        /// Input only. The TTL for this SecurityAction.
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        public SecurityActionArgs()
        {
        }
        public static new SecurityActionArgs Empty => new SecurityActionArgs();
    }
}