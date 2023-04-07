// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V2
{
    /// <summary>
    /// Creates a new Service in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:run/v2:Service")]
    public partial class Service : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects. Cloud Run API v2 does not support annotations with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system annotations in v1 now have a corresponding field in v2 Service. This field follows Kubernetes annotations' namespacing, limits, and rules. More info: https://kubernetes.io/docs/user-guide/annotations
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableDictionary<string, string>> Annotations { get; private set; } = null!;

        /// <summary>
        /// Settings for the Binary Authorization feature.
        /// </summary>
        [Output("binaryAuthorization")]
        public Output<Outputs.GoogleCloudRunV2BinaryAuthorizationResponse> BinaryAuthorization { get; private set; } = null!;

        /// <summary>
        /// Arbitrary identifier for the API client.
        /// </summary>
        [Output("client")]
        public Output<string> Client { get; private set; } = null!;

        /// <summary>
        /// Arbitrary version identifier for the API client.
        /// </summary>
        [Output("clientVersion")]
        public Output<string> ClientVersion { get; private set; } = null!;

        /// <summary>
        /// The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Service does not reach its Serving state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
        /// </summary>
        [Output("conditions")]
        public Output<ImmutableArray<Outputs.GoogleCloudRunV2ConditionResponse>> Conditions { get; private set; } = null!;

        /// <summary>
        /// The creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Email address of the authenticated creator.
        /// </summary>
        [Output("creator")]
        public Output<string> Creator { get; private set; } = null!;

        /// <summary>
        /// The deletion time.
        /// </summary>
        [Output("deleteTime")]
        public Output<string> DeleteTime { get; private set; } = null!;

        /// <summary>
        /// User-provided description of the Service. This field currently has a 512-character limit.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// For a deleted resource, the time after which it will be permamently deleted.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// A number that monotonically increases every time the user modifies the desired state. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a `string` instead of an `integer`.
        /// </summary>
        [Output("generation")]
        public Output<string> Generation { get; private set; } = null!;

        /// <summary>
        /// Provides the ingress settings for this Service. On output, returns the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED if no revision is active.
        /// </summary>
        [Output("ingress")]
        public Output<string> Ingress { get; private set; } = null!;

        /// <summary>
        /// Map of string keys and values that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run API v2 does not support labels with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system labels in v1 now have a corresponding field in v2 Service.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Email address of the last authenticated modifier.
        /// </summary>
        [Output("lastModifier")]
        public Output<string> LastModifier { get; private set; } = null!;

        /// <summary>
        /// Name of the last created revision. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
        /// </summary>
        [Output("latestCreatedRevision")]
        public Output<string> LatestCreatedRevision { get; private set; } = null!;

        /// <summary>
        /// Name of the latest revision that is serving traffic. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
        /// </summary>
        [Output("latestReadyRevision")]
        public Output<string> LatestReadyRevision { get; private set; } = null!;

        /// <summary>
        /// The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/terms/launch-stages). Cloud Run supports `ALPHA`, `BETA`, and `GA`. If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features. For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output.
        /// </summary>
        [Output("launchStage")]
        public Output<string> LaunchStage { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The fully qualified name of this Service. In CreateServiceRequest, this field is ignored, and instead composed from CreateServiceRequest.parent and CreateServiceRequest.service_id. Format: projects/{project}/locations/{location}/services/{service_id}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The generation of this Service currently serving traffic. See comments in `reconciling` for additional information on reconciliation process in Cloud Run. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a `string` instead of an `integer`.
        /// </summary>
        [Output("observedGeneration")]
        public Output<string> ObservedGeneration { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Returns true if the Service is currently being acted upon by the system to bring it into the desired state. When a new Service is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Service to the desired serving state. This process is called reconciliation. While reconciliation is in process, `observed_generation`, `latest_ready_revison`, `traffic_statuses`, and `uri` will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the serving state matches the Service, or there was an error, and reconciliation failed. This state can be found in `terminal_condition.state`. If reconciliation succeeded, the following fields will match: `traffic` and `traffic_statuses`, `observed_generation` and `generation`, `latest_ready_revision` and `latest_created_revision`. If reconciliation failed, `traffic_statuses`, `observed_generation`, and `latest_ready_revision` will have the state of the last serving revision, or empty for newly created Services. Additional information on the failure can be found in `terminal_condition` and `conditions`.
        /// </summary>
        [Output("reconciling")]
        public Output<bool> Reconciling { get; private set; } = null!;

        /// <summary>
        /// Reserved for future use.
        /// </summary>
        [Output("satisfiesPzs")]
        public Output<bool> SatisfiesPzs { get; private set; } = null!;

        /// <summary>
        /// Required. The unique identifier for the Service. It must begin with letter, and cannot end with hyphen; must contain fewer than 50 characters. The name of the service becomes {parent}/services/{service_id}.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;

        /// <summary>
        /// The template used to create revisions for this Service.
        /// </summary>
        [Output("template")]
        public Output<Outputs.GoogleCloudRunV2RevisionTemplateResponse> Template { get; private set; } = null!;

        /// <summary>
        /// The Condition of this Service, containing its readiness status, and detailed error information in case it did not reach a serving state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
        /// </summary>
        [Output("terminalCondition")]
        public Output<Outputs.GoogleCloudRunV2ConditionResponse> TerminalCondition { get; private set; } = null!;

        /// <summary>
        /// Specifies how to distribute traffic over a collection of Revisions belonging to the Service. If traffic is empty or not provided, defaults to 100% traffic to the latest `Ready` Revision.
        /// </summary>
        [Output("traffic")]
        public Output<ImmutableArray<Outputs.GoogleCloudRunV2TrafficTargetResponse>> Traffic { get; private set; } = null!;

        /// <summary>
        /// Detailed status information for corresponding traffic targets. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
        /// </summary>
        [Output("trafficStatuses")]
        public Output<ImmutableArray<Outputs.GoogleCloudRunV2TrafficTargetStatusResponse>> TrafficStatuses { get; private set; } = null!;

        /// <summary>
        /// Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The last-modified time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// The main URI in which this Service is serving traffic.
        /// </summary>
        [Output("uri")]
        public Output<string> Uri { get; private set; } = null!;

        /// <summary>
        /// Indicates that the request should be validated and default values populated, without persisting the request or creating any resources.
        /// </summary>
        [Output("validateOnly")]
        public Output<bool?> ValidateOnly { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("google-native:run/v2:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:run/v2:Service", name, null, MakeResourceOptions(options, id))
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
                    "serviceId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Service(name, id, options);
        }
    }

    public sealed class ServiceArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects. Cloud Run API v2 does not support annotations with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system annotations in v1 now have a corresponding field in v2 Service. This field follows Kubernetes annotations' namespacing, limits, and rules. More info: https://kubernetes.io/docs/user-guide/annotations
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// Settings for the Binary Authorization feature.
        /// </summary>
        [Input("binaryAuthorization")]
        public Input<Inputs.GoogleCloudRunV2BinaryAuthorizationArgs>? BinaryAuthorization { get; set; }

        /// <summary>
        /// Arbitrary identifier for the API client.
        /// </summary>
        [Input("client")]
        public Input<string>? Client { get; set; }

        /// <summary>
        /// Arbitrary version identifier for the API client.
        /// </summary>
        [Input("clientVersion")]
        public Input<string>? ClientVersion { get; set; }

        /// <summary>
        /// User-provided description of the Service. This field currently has a 512-character limit.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Provides the ingress settings for this Service. On output, returns the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED if no revision is active.
        /// </summary>
        [Input("ingress")]
        public Input<Pulumi.GoogleNative.Run.V2.ServiceIngress>? Ingress { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Map of string keys and values that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run API v2 does not support labels with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system labels in v1 now have a corresponding field in v2 Service.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/terms/launch-stages). Cloud Run supports `ALPHA`, `BETA`, and `GA`. If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features. For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output.
        /// </summary>
        [Input("launchStage")]
        public Input<Pulumi.GoogleNative.Run.V2.ServiceLaunchStage>? LaunchStage { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The fully qualified name of this Service. In CreateServiceRequest, this field is ignored, and instead composed from CreateServiceRequest.parent and CreateServiceRequest.service_id. Format: projects/{project}/locations/{location}/services/{service_id}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Required. The unique identifier for the Service. It must begin with letter, and cannot end with hyphen; must contain fewer than 50 characters. The name of the service becomes {parent}/services/{service_id}.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        /// <summary>
        /// The template used to create revisions for this Service.
        /// </summary>
        [Input("template", required: true)]
        public Input<Inputs.GoogleCloudRunV2RevisionTemplateArgs> Template { get; set; } = null!;

        [Input("traffic")]
        private InputList<Inputs.GoogleCloudRunV2TrafficTargetArgs>? _traffic;

        /// <summary>
        /// Specifies how to distribute traffic over a collection of Revisions belonging to the Service. If traffic is empty or not provided, defaults to 100% traffic to the latest `Ready` Revision.
        /// </summary>
        public InputList<Inputs.GoogleCloudRunV2TrafficTargetArgs> Traffic
        {
            get => _traffic ?? (_traffic = new InputList<Inputs.GoogleCloudRunV2TrafficTargetArgs>());
            set => _traffic = value;
        }

        /// <summary>
        /// Indicates that the request should be validated and default values populated, without persisting the request or creating any resources.
        /// </summary>
        [Input("validateOnly")]
        public Input<bool>? ValidateOnly { get; set; }

        public ServiceArgs()
        {
        }
        public static new ServiceArgs Empty => new ServiceArgs();
    }
}
