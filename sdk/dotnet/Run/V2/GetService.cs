// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V2
{
    public static class GetService
    {
        /// <summary>
        /// Gets information about a Service.
        /// </summary>
        public static Task<GetServiceResult> InvokeAsync(GetServiceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceResult>("google-native:run/v2:getService", args ?? new GetServiceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Service.
        /// </summary>
        public static Output<GetServiceResult> Invoke(GetServiceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceResult>("google-native:run/v2:getService", args ?? new GetServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetServiceArgs()
        {
        }
        public static new GetServiceArgs Empty => new GetServiceArgs();
    }

    public sealed class GetServiceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetServiceInvokeArgs()
        {
        }
        public static new GetServiceInvokeArgs Empty => new GetServiceInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceResult
    {
        /// <summary>
        /// Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects. Cloud Run will populate some annotations using 'run.googleapis.com' or 'serving.knative.dev' namespaces. This field follows Kubernetes annotations' namespacing, limits, and rules. More info: https://kubernetes.io/docs/user-guide/annotations
        /// </summary>
        public readonly ImmutableDictionary<string, string> Annotations;
        /// <summary>
        /// Settings for the Binary Authorization feature.
        /// </summary>
        public readonly Outputs.GoogleCloudRunV2BinaryAuthorizationResponse BinaryAuthorization;
        /// <summary>
        /// Arbitrary identifier for the API client.
        /// </summary>
        public readonly string Client;
        /// <summary>
        /// Arbitrary version identifier for the API client.
        /// </summary>
        public readonly string ClientVersion;
        /// <summary>
        /// The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Service does not reach its Serving state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudRunV2ConditionResponse> Conditions;
        /// <summary>
        /// The creation time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Email address of the authenticated creator.
        /// </summary>
        public readonly string Creator;
        /// <summary>
        /// The deletion time.
        /// </summary>
        public readonly string DeleteTime;
        /// <summary>
        /// User-provided description of the Service. This field currently has a 512-character limit.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// For a deleted resource, the time after which it will be permamently deleted.
        /// </summary>
        public readonly string ExpireTime;
        /// <summary>
        /// A number that monotonically increases every time the user modifies the desired state. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a `string` instead of an `integer`.
        /// </summary>
        public readonly string Generation;
        /// <summary>
        /// Provides the ingress settings for this Service. On output, returns the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED if no revision is active.
        /// </summary>
        public readonly string Ingress;
        /// <summary>
        /// Map of string keys and values that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run will populate some labels with 'run.googleapis.com' or 'serving.knative.dev' namespaces. Those labels are read-only, and user changes will not be preserved.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Email address of the last authenticated modifier.
        /// </summary>
        public readonly string LastModifier;
        /// <summary>
        /// Name of the last created revision. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
        /// </summary>
        public readonly string LatestCreatedRevision;
        /// <summary>
        /// Name of the latest revision that is serving traffic. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
        /// </summary>
        public readonly string LatestReadyRevision;
        /// <summary>
        /// The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/terms/launch-stages). Cloud Run supports `ALPHA`, `BETA`, and `GA`. If no value is specified, GA is assumed.
        /// </summary>
        public readonly string LaunchStage;
        /// <summary>
        /// The fully qualified name of this Service. In CreateServiceRequest, this field is ignored, and instead composed from CreateServiceRequest.parent and CreateServiceRequest.service_id. Format: projects/{project}/locations/{location}/services/{service_id}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The generation of this Service currently serving traffic. See comments in `reconciling` for additional information on reconciliation process in Cloud Run. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a `string` instead of an `integer`.
        /// </summary>
        public readonly string ObservedGeneration;
        /// <summary>
        /// Returns true if the Service is currently being acted upon by the system to bring it into the desired state. When a new Service is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Service to the desired serving state. This process is called reconciliation. While reconciliation is in process, `observed_generation`, `latest_ready_revison`, `traffic_statuses`, and `uri` will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the serving state matches the Service, or there was an error, and reconciliation failed. This state can be found in `terminal_condition.state`. If reconciliation succeeded, the following fields will match: `traffic` and `traffic_statuses`, `observed_generation` and `generation`, `latest_ready_revision` and `latest_created_revision`. If reconciliation failed, `traffic_statuses`, `observed_generation`, and `latest_ready_revision` will have the state of the last serving revision, or empty for newly created Services. Additional information on the failure can be found in `terminal_condition` and `conditions`.
        /// </summary>
        public readonly bool Reconciling;
        /// <summary>
        /// The template used to create revisions for this Service.
        /// </summary>
        public readonly Outputs.GoogleCloudRunV2RevisionTemplateResponse Template;
        /// <summary>
        /// The Condition of this Service, containing its readiness status, and detailed error information in case it did not reach a serving state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
        /// </summary>
        public readonly Outputs.GoogleCloudRunV2ConditionResponse TerminalCondition;
        /// <summary>
        /// Specifies how to distribute traffic over a collection of Revisions belonging to the Service. If traffic is empty or not provided, defaults to 100% traffic to the latest `Ready` Revision.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudRunV2TrafficTargetResponse> Traffic;
        /// <summary>
        /// Detailed status information for corresponding traffic targets. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudRunV2TrafficTargetStatusResponse> TrafficStatuses;
        /// <summary>
        /// Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// The last-modified time.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The main URI in which this Service is serving traffic.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private GetServiceResult(
            ImmutableDictionary<string, string> annotations,

            Outputs.GoogleCloudRunV2BinaryAuthorizationResponse binaryAuthorization,

            string client,

            string clientVersion,

            ImmutableArray<Outputs.GoogleCloudRunV2ConditionResponse> conditions,

            string createTime,

            string creator,

            string deleteTime,

            string description,

            string etag,

            string expireTime,

            string generation,

            string ingress,

            ImmutableDictionary<string, string> labels,

            string lastModifier,

            string latestCreatedRevision,

            string latestReadyRevision,

            string launchStage,

            string name,

            string observedGeneration,

            bool reconciling,

            Outputs.GoogleCloudRunV2RevisionTemplateResponse template,

            Outputs.GoogleCloudRunV2ConditionResponse terminalCondition,

            ImmutableArray<Outputs.GoogleCloudRunV2TrafficTargetResponse> traffic,

            ImmutableArray<Outputs.GoogleCloudRunV2TrafficTargetStatusResponse> trafficStatuses,

            string uid,

            string updateTime,

            string uri)
        {
            Annotations = annotations;
            BinaryAuthorization = binaryAuthorization;
            Client = client;
            ClientVersion = clientVersion;
            Conditions = conditions;
            CreateTime = createTime;
            Creator = creator;
            DeleteTime = deleteTime;
            Description = description;
            Etag = etag;
            ExpireTime = expireTime;
            Generation = generation;
            Ingress = ingress;
            Labels = labels;
            LastModifier = lastModifier;
            LatestCreatedRevision = latestCreatedRevision;
            LatestReadyRevision = latestReadyRevision;
            LaunchStage = launchStage;
            Name = name;
            ObservedGeneration = observedGeneration;
            Reconciling = reconciling;
            Template = template;
            TerminalCondition = terminalCondition;
            Traffic = traffic;
            TrafficStatuses = trafficStatuses;
            Uid = uid;
            UpdateTime = updateTime;
            Uri = uri;
        }
    }
}
