// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Service in a given project and location.
type Service struct {
	pulumi.CustomResourceState

	// Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects. Cloud Run API v2 does not support annotations with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system annotations in v1 now have a corresponding field in v2 Service. This field follows Kubernetes annotations' namespacing, limits, and rules. More info: https://kubernetes.io/docs/user-guide/annotations
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Settings for the Binary Authorization feature.
	BinaryAuthorization GoogleCloudRunV2BinaryAuthorizationResponseOutput `pulumi:"binaryAuthorization"`
	// Arbitrary identifier for the API client.
	Client pulumi.StringOutput `pulumi:"client"`
	// Arbitrary version identifier for the API client.
	ClientVersion pulumi.StringOutput `pulumi:"clientVersion"`
	// The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Service does not reach its Serving state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
	Conditions GoogleCloudRunV2ConditionResponseArrayOutput `pulumi:"conditions"`
	// The creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Email address of the authenticated creator.
	Creator pulumi.StringOutput `pulumi:"creator"`
	// The deletion time.
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// User-provided description of the Service. This field currently has a 512-character limit.
	Description pulumi.StringOutput `pulumi:"description"`
	// A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// For a deleted resource, the time after which it will be permamently deleted.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// A number that monotonically increases every time the user modifies the desired state. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a `string` instead of an `integer`.
	Generation pulumi.StringOutput `pulumi:"generation"`
	// Provides the ingress settings for this Service. On output, returns the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED if no revision is active.
	Ingress pulumi.StringOutput `pulumi:"ingress"`
	// Map of string keys and values that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run API v2 does not support labels with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system labels in v1 now have a corresponding field in v2 Service.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Email address of the last authenticated modifier.
	LastModifier pulumi.StringOutput `pulumi:"lastModifier"`
	// Name of the last created revision. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
	LatestCreatedRevision pulumi.StringOutput `pulumi:"latestCreatedRevision"`
	// Name of the latest revision that is serving traffic. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
	LatestReadyRevision pulumi.StringOutput `pulumi:"latestReadyRevision"`
	// The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/terms/launch-stages). Cloud Run supports `ALPHA`, `BETA`, and `GA`. If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features. For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output.
	LaunchStage pulumi.StringOutput `pulumi:"launchStage"`
	Location    pulumi.StringOutput `pulumi:"location"`
	// The fully qualified name of this Service. In CreateServiceRequest, this field is ignored, and instead composed from CreateServiceRequest.parent and CreateServiceRequest.service_id. Format: projects/{project}/locations/{location}/services/{service_id}
	Name pulumi.StringOutput `pulumi:"name"`
	// The generation of this Service currently serving traffic. See comments in `reconciling` for additional information on reconciliation process in Cloud Run. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a `string` instead of an `integer`.
	ObservedGeneration pulumi.StringOutput `pulumi:"observedGeneration"`
	Project            pulumi.StringOutput `pulumi:"project"`
	// Returns true if the Service is currently being acted upon by the system to bring it into the desired state. When a new Service is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Service to the desired serving state. This process is called reconciliation. While reconciliation is in process, `observed_generation`, `latest_ready_revison`, `traffic_statuses`, and `uri` will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the serving state matches the Service, or there was an error, and reconciliation failed. This state can be found in `terminal_condition.state`. If reconciliation succeeded, the following fields will match: `traffic` and `traffic_statuses`, `observed_generation` and `generation`, `latest_ready_revision` and `latest_created_revision`. If reconciliation failed, `traffic_statuses`, `observed_generation`, and `latest_ready_revision` will have the state of the last serving revision, or empty for newly created Services. Additional information on the failure can be found in `terminal_condition` and `conditions`.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// Reserved for future use.
	SatisfiesPzs pulumi.BoolOutput `pulumi:"satisfiesPzs"`
	// Required. The unique identifier for the Service. It must begin with letter, and cannot end with hyphen; must contain fewer than 50 characters. The name of the service becomes {parent}/services/{service_id}.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// The template used to create revisions for this Service.
	Template GoogleCloudRunV2RevisionTemplateResponseOutput `pulumi:"template"`
	// The Condition of this Service, containing its readiness status, and detailed error information in case it did not reach a serving state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
	TerminalCondition GoogleCloudRunV2ConditionResponseOutput `pulumi:"terminalCondition"`
	// Specifies how to distribute traffic over a collection of Revisions belonging to the Service. If traffic is empty or not provided, defaults to 100% traffic to the latest `Ready` Revision.
	Traffic GoogleCloudRunV2TrafficTargetResponseArrayOutput `pulumi:"traffic"`
	// Detailed status information for corresponding traffic targets. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
	TrafficStatuses GoogleCloudRunV2TrafficTargetStatusResponseArrayOutput `pulumi:"trafficStatuses"`
	// Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The last-modified time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The main URI in which this Service is serving traffic.
	Uri pulumi.StringOutput `pulumi:"uri"`
	// Indicates that the request should be validated and default values populated, without persisting the request or creating any resources.
	ValidateOnly pulumi.BoolPtrOutput `pulumi:"validateOnly"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	if args.Template == nil {
		return nil, errors.New("invalid value for required argument 'Template'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"serviceId",
	})
	opts = append(opts, replaceOnChanges)
	var resource Service
	err := ctx.RegisterResource("google-native:run/v2:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("google-native:run/v2:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
}

type ServiceState struct {
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects. Cloud Run API v2 does not support annotations with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system annotations in v1 now have a corresponding field in v2 Service. This field follows Kubernetes annotations' namespacing, limits, and rules. More info: https://kubernetes.io/docs/user-guide/annotations
	Annotations map[string]string `pulumi:"annotations"`
	// Settings for the Binary Authorization feature.
	BinaryAuthorization *GoogleCloudRunV2BinaryAuthorization `pulumi:"binaryAuthorization"`
	// Arbitrary identifier for the API client.
	Client *string `pulumi:"client"`
	// Arbitrary version identifier for the API client.
	ClientVersion *string `pulumi:"clientVersion"`
	// User-provided description of the Service. This field currently has a 512-character limit.
	Description *string `pulumi:"description"`
	// Provides the ingress settings for this Service. On output, returns the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED if no revision is active.
	Ingress *ServiceIngress `pulumi:"ingress"`
	// Map of string keys and values that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run API v2 does not support labels with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system labels in v1 now have a corresponding field in v2 Service.
	Labels map[string]string `pulumi:"labels"`
	// The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/terms/launch-stages). Cloud Run supports `ALPHA`, `BETA`, and `GA`. If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features. For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output.
	LaunchStage *ServiceLaunchStage `pulumi:"launchStage"`
	Location    *string             `pulumi:"location"`
	// The fully qualified name of this Service. In CreateServiceRequest, this field is ignored, and instead composed from CreateServiceRequest.parent and CreateServiceRequest.service_id. Format: projects/{project}/locations/{location}/services/{service_id}
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Required. The unique identifier for the Service. It must begin with letter, and cannot end with hyphen; must contain fewer than 50 characters. The name of the service becomes {parent}/services/{service_id}.
	ServiceId string `pulumi:"serviceId"`
	// The template used to create revisions for this Service.
	Template GoogleCloudRunV2RevisionTemplate `pulumi:"template"`
	// Specifies how to distribute traffic over a collection of Revisions belonging to the Service. If traffic is empty or not provided, defaults to 100% traffic to the latest `Ready` Revision.
	Traffic []GoogleCloudRunV2TrafficTarget `pulumi:"traffic"`
	// Indicates that the request should be validated and default values populated, without persisting the request or creating any resources.
	ValidateOnly *bool `pulumi:"validateOnly"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects. Cloud Run API v2 does not support annotations with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system annotations in v1 now have a corresponding field in v2 Service. This field follows Kubernetes annotations' namespacing, limits, and rules. More info: https://kubernetes.io/docs/user-guide/annotations
	Annotations pulumi.StringMapInput
	// Settings for the Binary Authorization feature.
	BinaryAuthorization GoogleCloudRunV2BinaryAuthorizationPtrInput
	// Arbitrary identifier for the API client.
	Client pulumi.StringPtrInput
	// Arbitrary version identifier for the API client.
	ClientVersion pulumi.StringPtrInput
	// User-provided description of the Service. This field currently has a 512-character limit.
	Description pulumi.StringPtrInput
	// Provides the ingress settings for this Service. On output, returns the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED if no revision is active.
	Ingress ServiceIngressPtrInput
	// Map of string keys and values that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run API v2 does not support labels with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system labels in v1 now have a corresponding field in v2 Service.
	Labels pulumi.StringMapInput
	// The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/terms/launch-stages). Cloud Run supports `ALPHA`, `BETA`, and `GA`. If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features. For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output.
	LaunchStage ServiceLaunchStagePtrInput
	Location    pulumi.StringPtrInput
	// The fully qualified name of this Service. In CreateServiceRequest, this field is ignored, and instead composed from CreateServiceRequest.parent and CreateServiceRequest.service_id. Format: projects/{project}/locations/{location}/services/{service_id}
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Required. The unique identifier for the Service. It must begin with letter, and cannot end with hyphen; must contain fewer than 50 characters. The name of the service becomes {parent}/services/{service_id}.
	ServiceId pulumi.StringInput
	// The template used to create revisions for this Service.
	Template GoogleCloudRunV2RevisionTemplateInput
	// Specifies how to distribute traffic over a collection of Revisions belonging to the Service. If traffic is empty or not provided, defaults to 100% traffic to the latest `Ready` Revision.
	Traffic GoogleCloudRunV2TrafficTargetArrayInput
	// Indicates that the request should be validated and default values populated, without persisting the request or creating any resources.
	ValidateOnly pulumi.BoolPtrInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

// Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects. Cloud Run API v2 does not support annotations with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system annotations in v1 now have a corresponding field in v2 Service. This field follows Kubernetes annotations' namespacing, limits, and rules. More info: https://kubernetes.io/docs/user-guide/annotations
func (o ServiceOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Service) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Settings for the Binary Authorization feature.
func (o ServiceOutput) BinaryAuthorization() GoogleCloudRunV2BinaryAuthorizationResponseOutput {
	return o.ApplyT(func(v *Service) GoogleCloudRunV2BinaryAuthorizationResponseOutput { return v.BinaryAuthorization }).(GoogleCloudRunV2BinaryAuthorizationResponseOutput)
}

// Arbitrary identifier for the API client.
func (o ServiceOutput) Client() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Client }).(pulumi.StringOutput)
}

// Arbitrary version identifier for the API client.
func (o ServiceOutput) ClientVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ClientVersion }).(pulumi.StringOutput)
}

// The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Service does not reach its Serving state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
func (o ServiceOutput) Conditions() GoogleCloudRunV2ConditionResponseArrayOutput {
	return o.ApplyT(func(v *Service) GoogleCloudRunV2ConditionResponseArrayOutput { return v.Conditions }).(GoogleCloudRunV2ConditionResponseArrayOutput)
}

// The creation time.
func (o ServiceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Email address of the authenticated creator.
func (o ServiceOutput) Creator() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Creator }).(pulumi.StringOutput)
}

// The deletion time.
func (o ServiceOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.DeleteTime }).(pulumi.StringOutput)
}

// User-provided description of the Service. This field currently has a 512-character limit.
func (o ServiceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.
func (o ServiceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// For a deleted resource, the time after which it will be permamently deleted.
func (o ServiceOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ExpireTime }).(pulumi.StringOutput)
}

// A number that monotonically increases every time the user modifies the desired state. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a `string` instead of an `integer`.
func (o ServiceOutput) Generation() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Generation }).(pulumi.StringOutput)
}

// Provides the ingress settings for this Service. On output, returns the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED if no revision is active.
func (o ServiceOutput) Ingress() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Ingress }).(pulumi.StringOutput)
}

// Map of string keys and values that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run API v2 does not support labels with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system labels in v1 now have a corresponding field in v2 Service.
func (o ServiceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Service) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Email address of the last authenticated modifier.
func (o ServiceOutput) LastModifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.LastModifier }).(pulumi.StringOutput)
}

// Name of the last created revision. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
func (o ServiceOutput) LatestCreatedRevision() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.LatestCreatedRevision }).(pulumi.StringOutput)
}

// Name of the latest revision that is serving traffic. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
func (o ServiceOutput) LatestReadyRevision() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.LatestReadyRevision }).(pulumi.StringOutput)
}

// The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/terms/launch-stages). Cloud Run supports `ALPHA`, `BETA`, and `GA`. If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features. For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output.
func (o ServiceOutput) LaunchStage() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.LaunchStage }).(pulumi.StringOutput)
}

func (o ServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The fully qualified name of this Service. In CreateServiceRequest, this field is ignored, and instead composed from CreateServiceRequest.parent and CreateServiceRequest.service_id. Format: projects/{project}/locations/{location}/services/{service_id}
func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The generation of this Service currently serving traffic. See comments in `reconciling` for additional information on reconciliation process in Cloud Run. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a `string` instead of an `integer`.
func (o ServiceOutput) ObservedGeneration() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ObservedGeneration }).(pulumi.StringOutput)
}

func (o ServiceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Returns true if the Service is currently being acted upon by the system to bring it into the desired state. When a new Service is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Service to the desired serving state. This process is called reconciliation. While reconciliation is in process, `observed_generation`, `latest_ready_revison`, `traffic_statuses`, and `uri` will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the serving state matches the Service, or there was an error, and reconciliation failed. This state can be found in `terminal_condition.state`. If reconciliation succeeded, the following fields will match: `traffic` and `traffic_statuses`, `observed_generation` and `generation`, `latest_ready_revision` and `latest_created_revision`. If reconciliation failed, `traffic_statuses`, `observed_generation`, and `latest_ready_revision` will have the state of the last serving revision, or empty for newly created Services. Additional information on the failure can be found in `terminal_condition` and `conditions`.
func (o ServiceOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *Service) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// Reserved for future use.
func (o ServiceOutput) SatisfiesPzs() pulumi.BoolOutput {
	return o.ApplyT(func(v *Service) pulumi.BoolOutput { return v.SatisfiesPzs }).(pulumi.BoolOutput)
}

// Required. The unique identifier for the Service. It must begin with letter, and cannot end with hyphen; must contain fewer than 50 characters. The name of the service becomes {parent}/services/{service_id}.
func (o ServiceOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

// The template used to create revisions for this Service.
func (o ServiceOutput) Template() GoogleCloudRunV2RevisionTemplateResponseOutput {
	return o.ApplyT(func(v *Service) GoogleCloudRunV2RevisionTemplateResponseOutput { return v.Template }).(GoogleCloudRunV2RevisionTemplateResponseOutput)
}

// The Condition of this Service, containing its readiness status, and detailed error information in case it did not reach a serving state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
func (o ServiceOutput) TerminalCondition() GoogleCloudRunV2ConditionResponseOutput {
	return o.ApplyT(func(v *Service) GoogleCloudRunV2ConditionResponseOutput { return v.TerminalCondition }).(GoogleCloudRunV2ConditionResponseOutput)
}

// Specifies how to distribute traffic over a collection of Revisions belonging to the Service. If traffic is empty or not provided, defaults to 100% traffic to the latest `Ready` Revision.
func (o ServiceOutput) Traffic() GoogleCloudRunV2TrafficTargetResponseArrayOutput {
	return o.ApplyT(func(v *Service) GoogleCloudRunV2TrafficTargetResponseArrayOutput { return v.Traffic }).(GoogleCloudRunV2TrafficTargetResponseArrayOutput)
}

// Detailed status information for corresponding traffic targets. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
func (o ServiceOutput) TrafficStatuses() GoogleCloudRunV2TrafficTargetStatusResponseArrayOutput {
	return o.ApplyT(func(v *Service) GoogleCloudRunV2TrafficTargetStatusResponseArrayOutput { return v.TrafficStatuses }).(GoogleCloudRunV2TrafficTargetStatusResponseArrayOutput)
}

// Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
func (o ServiceOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The last-modified time.
func (o ServiceOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// The main URI in which this Service is serving traffic.
func (o ServiceOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

// Indicates that the request should be validated and default values populated, without persisting the request or creating any resources.
func (o ServiceOutput) ValidateOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.BoolPtrOutput { return v.ValidateOnly }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterOutputType(ServiceOutput{})
}
