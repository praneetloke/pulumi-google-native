// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a Service.
func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceResult
	err := ctx.Invoke("google-native:run/v2:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	Location  string  `pulumi:"location"`
	Project   *string `pulumi:"project"`
	ServiceId string  `pulumi:"serviceId"`
}

type LookupServiceResult struct {
	// Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects. Cloud Run API v2 does not support annotations with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected in new resources. All system annotations in v1 now have a corresponding field in v2 Service. This field follows Kubernetes annotations' namespacing, limits, and rules.
	Annotations map[string]string `pulumi:"annotations"`
	// Settings for the Binary Authorization feature.
	BinaryAuthorization GoogleCloudRunV2BinaryAuthorizationResponse `pulumi:"binaryAuthorization"`
	// Arbitrary identifier for the API client.
	Client string `pulumi:"client"`
	// Arbitrary version identifier for the API client.
	ClientVersion string `pulumi:"clientVersion"`
	// The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Service does not reach its Serving state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
	Conditions []GoogleCloudRunV2ConditionResponse `pulumi:"conditions"`
	// The creation time.
	CreateTime string `pulumi:"createTime"`
	// Email address of the authenticated creator.
	Creator string `pulumi:"creator"`
	// The deletion time.
	DeleteTime string `pulumi:"deleteTime"`
	// User-provided description of the Service. This field currently has a 512-character limit.
	Description string `pulumi:"description"`
	// A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.
	Etag string `pulumi:"etag"`
	// For a deleted resource, the time after which it will be permamently deleted.
	ExpireTime string `pulumi:"expireTime"`
	// A number that monotonically increases every time the user modifies the desired state. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a `string` instead of an `integer`.
	Generation string `pulumi:"generation"`
	// Provides the ingress settings for this Service. On output, returns the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED if no revision is active.
	Ingress string `pulumi:"ingress"`
	// Unstructured key value map that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels. Cloud Run API v2 does not support labels with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system labels in v1 now have a corresponding field in v2 Service.
	Labels map[string]string `pulumi:"labels"`
	// Email address of the last authenticated modifier.
	LastModifier string `pulumi:"lastModifier"`
	// Name of the last created revision. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
	LatestCreatedRevision string `pulumi:"latestCreatedRevision"`
	// Name of the latest revision that is serving traffic. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
	LatestReadyRevision string `pulumi:"latestReadyRevision"`
	// The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/terms/launch-stages). Cloud Run supports `ALPHA`, `BETA`, and `GA`. If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features. For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output.
	LaunchStage string `pulumi:"launchStage"`
	// The fully qualified name of this Service. In CreateServiceRequest, this field is ignored, and instead composed from CreateServiceRequest.parent and CreateServiceRequest.service_id. Format: projects/{project}/locations/{location}/services/{service_id}
	Name string `pulumi:"name"`
	// The generation of this Service currently serving traffic. See comments in `reconciling` for additional information on reconciliation process in Cloud Run. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a `string` instead of an `integer`.
	ObservedGeneration string `pulumi:"observedGeneration"`
	// Returns true if the Service is currently being acted upon by the system to bring it into the desired state. When a new Service is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Service to the desired serving state. This process is called reconciliation. While reconciliation is in process, `observed_generation`, `latest_ready_revison`, `traffic_statuses`, and `uri` will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the serving state matches the Service, or there was an error, and reconciliation failed. This state can be found in `terminal_condition.state`. If reconciliation succeeded, the following fields will match: `traffic` and `traffic_statuses`, `observed_generation` and `generation`, `latest_ready_revision` and `latest_created_revision`. If reconciliation failed, `traffic_statuses`, `observed_generation`, and `latest_ready_revision` will have the state of the last serving revision, or empty for newly created Services. Additional information on the failure can be found in `terminal_condition` and `conditions`.
	Reconciling bool `pulumi:"reconciling"`
	// Reserved for future use.
	SatisfiesPzs bool `pulumi:"satisfiesPzs"`
	// The template used to create revisions for this Service.
	Template GoogleCloudRunV2RevisionTemplateResponse `pulumi:"template"`
	// The Condition of this Service, containing its readiness status, and detailed error information in case it did not reach a serving state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
	TerminalCondition GoogleCloudRunV2ConditionResponse `pulumi:"terminalCondition"`
	// Specifies how to distribute traffic over a collection of Revisions belonging to the Service. If traffic is empty or not provided, defaults to 100% traffic to the latest `Ready` Revision.
	Traffic []GoogleCloudRunV2TrafficTargetResponse `pulumi:"traffic"`
	// Detailed status information for corresponding traffic targets. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
	TrafficStatuses []GoogleCloudRunV2TrafficTargetStatusResponse `pulumi:"trafficStatuses"`
	// Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
	Uid string `pulumi:"uid"`
	// The last-modified time.
	UpdateTime string `pulumi:"updateTime"`
	// The main URI in which this Service is serving traffic.
	Uri string `pulumi:"uri"`
}

func LookupServiceOutput(ctx *pulumi.Context, args LookupServiceOutputArgs, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceResult, error) {
			args := v.(LookupServiceArgs)
			r, err := LookupService(ctx, &args, opts...)
			var s LookupServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceResultOutput)
}

type LookupServiceOutputArgs struct {
	Location  pulumi.StringInput    `pulumi:"location"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
	ServiceId pulumi.StringInput    `pulumi:"serviceId"`
}

func (LookupServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}

type LookupServiceResultOutput struct{ *pulumi.OutputState }

func (LookupServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceResult)(nil)).Elem()
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutput() LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutputWithContext(ctx context.Context) LookupServiceResultOutput {
	return o
}

// Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects. Cloud Run API v2 does not support annotations with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected in new resources. All system annotations in v1 now have a corresponding field in v2 Service. This field follows Kubernetes annotations' namespacing, limits, and rules.
func (o LookupServiceResultOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceResult) map[string]string { return v.Annotations }).(pulumi.StringMapOutput)
}

// Settings for the Binary Authorization feature.
func (o LookupServiceResultOutput) BinaryAuthorization() GoogleCloudRunV2BinaryAuthorizationResponseOutput {
	return o.ApplyT(func(v LookupServiceResult) GoogleCloudRunV2BinaryAuthorizationResponse { return v.BinaryAuthorization }).(GoogleCloudRunV2BinaryAuthorizationResponseOutput)
}

// Arbitrary identifier for the API client.
func (o LookupServiceResultOutput) Client() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Client }).(pulumi.StringOutput)
}

// Arbitrary version identifier for the API client.
func (o LookupServiceResultOutput) ClientVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.ClientVersion }).(pulumi.StringOutput)
}

// The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Service does not reach its Serving state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
func (o LookupServiceResultOutput) Conditions() GoogleCloudRunV2ConditionResponseArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []GoogleCloudRunV2ConditionResponse { return v.Conditions }).(GoogleCloudRunV2ConditionResponseArrayOutput)
}

// The creation time.
func (o LookupServiceResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Email address of the authenticated creator.
func (o LookupServiceResultOutput) Creator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Creator }).(pulumi.StringOutput)
}

// The deletion time.
func (o LookupServiceResultOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.DeleteTime }).(pulumi.StringOutput)
}

// User-provided description of the Service. This field currently has a 512-character limit.
func (o LookupServiceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Description }).(pulumi.StringOutput)
}

// A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.
func (o LookupServiceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Etag }).(pulumi.StringOutput)
}

// For a deleted resource, the time after which it will be permamently deleted.
func (o LookupServiceResultOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.ExpireTime }).(pulumi.StringOutput)
}

// A number that monotonically increases every time the user modifies the desired state. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a `string` instead of an `integer`.
func (o LookupServiceResultOutput) Generation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Generation }).(pulumi.StringOutput)
}

// Provides the ingress settings for this Service. On output, returns the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED if no revision is active.
func (o LookupServiceResultOutput) Ingress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Ingress }).(pulumi.StringOutput)
}

// Unstructured key value map that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels. Cloud Run API v2 does not support labels with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system labels in v1 now have a corresponding field in v2 Service.
func (o LookupServiceResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Email address of the last authenticated modifier.
func (o LookupServiceResultOutput) LastModifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.LastModifier }).(pulumi.StringOutput)
}

// Name of the last created revision. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
func (o LookupServiceResultOutput) LatestCreatedRevision() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.LatestCreatedRevision }).(pulumi.StringOutput)
}

// Name of the latest revision that is serving traffic. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
func (o LookupServiceResultOutput) LatestReadyRevision() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.LatestReadyRevision }).(pulumi.StringOutput)
}

// The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/terms/launch-stages). Cloud Run supports `ALPHA`, `BETA`, and `GA`. If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features. For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output.
func (o LookupServiceResultOutput) LaunchStage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.LaunchStage }).(pulumi.StringOutput)
}

// The fully qualified name of this Service. In CreateServiceRequest, this field is ignored, and instead composed from CreateServiceRequest.parent and CreateServiceRequest.service_id. Format: projects/{project}/locations/{location}/services/{service_id}
func (o LookupServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The generation of this Service currently serving traffic. See comments in `reconciling` for additional information on reconciliation process in Cloud Run. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a `string` instead of an `integer`.
func (o LookupServiceResultOutput) ObservedGeneration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.ObservedGeneration }).(pulumi.StringOutput)
}

// Returns true if the Service is currently being acted upon by the system to bring it into the desired state. When a new Service is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Service to the desired serving state. This process is called reconciliation. While reconciliation is in process, `observed_generation`, `latest_ready_revison`, `traffic_statuses`, and `uri` will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the serving state matches the Service, or there was an error, and reconciliation failed. This state can be found in `terminal_condition.state`. If reconciliation succeeded, the following fields will match: `traffic` and `traffic_statuses`, `observed_generation` and `generation`, `latest_ready_revision` and `latest_created_revision`. If reconciliation failed, `traffic_statuses`, `observed_generation`, and `latest_ready_revision` will have the state of the last serving revision, or empty for newly created Services. Additional information on the failure can be found in `terminal_condition` and `conditions`.
func (o LookupServiceResultOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServiceResult) bool { return v.Reconciling }).(pulumi.BoolOutput)
}

// Reserved for future use.
func (o LookupServiceResultOutput) SatisfiesPzs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServiceResult) bool { return v.SatisfiesPzs }).(pulumi.BoolOutput)
}

// The template used to create revisions for this Service.
func (o LookupServiceResultOutput) Template() GoogleCloudRunV2RevisionTemplateResponseOutput {
	return o.ApplyT(func(v LookupServiceResult) GoogleCloudRunV2RevisionTemplateResponse { return v.Template }).(GoogleCloudRunV2RevisionTemplateResponseOutput)
}

// The Condition of this Service, containing its readiness status, and detailed error information in case it did not reach a serving state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
func (o LookupServiceResultOutput) TerminalCondition() GoogleCloudRunV2ConditionResponseOutput {
	return o.ApplyT(func(v LookupServiceResult) GoogleCloudRunV2ConditionResponse { return v.TerminalCondition }).(GoogleCloudRunV2ConditionResponseOutput)
}

// Specifies how to distribute traffic over a collection of Revisions belonging to the Service. If traffic is empty or not provided, defaults to 100% traffic to the latest `Ready` Revision.
func (o LookupServiceResultOutput) Traffic() GoogleCloudRunV2TrafficTargetResponseArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []GoogleCloudRunV2TrafficTargetResponse { return v.Traffic }).(GoogleCloudRunV2TrafficTargetResponseArrayOutput)
}

// Detailed status information for corresponding traffic targets. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
func (o LookupServiceResultOutput) TrafficStatuses() GoogleCloudRunV2TrafficTargetStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []GoogleCloudRunV2TrafficTargetStatusResponse { return v.TrafficStatuses }).(GoogleCloudRunV2TrafficTargetStatusResponseArrayOutput)
}

// Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
func (o LookupServiceResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Uid }).(pulumi.StringOutput)
}

// The last-modified time.
func (o LookupServiceResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// The main URI in which this Service is serving traffic.
func (o LookupServiceResultOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Uri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}
