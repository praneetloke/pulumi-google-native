// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Version resource. By default, only a BASIC_VIEW will be returned. Specify the FULL_VIEW parameter to get the full resource.
func LookupVersion(ctx *pulumi.Context, args *LookupVersionArgs, opts ...pulumi.InvokeOption) (*LookupVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVersionResult
	err := ctx.Invoke("google-native:appengine/v1beta:getVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVersionArgs struct {
	AppId     string  `pulumi:"appId"`
	ServiceId string  `pulumi:"serviceId"`
	VersionId string  `pulumi:"versionId"`
	View      *string `pulumi:"view"`
}

type LookupVersionResult struct {
	// Serving configuration for Google Cloud Endpoints (https://cloud.google.com/endpoints).Only returned in GET requests if view=FULL is set.
	ApiConfig ApiConfigHandlerResponse `pulumi:"apiConfig"`
	// Allows App Engine second generation runtimes to access the legacy bundled services.
	AppEngineApis bool `pulumi:"appEngineApis"`
	// Automatic scaling is based on request rate, response latencies, and other application metrics. Instances are dynamically created and destroyed as needed in order to handle traffic.
	AutomaticScaling AutomaticScalingResponse `pulumi:"automaticScaling"`
	// A service with basic scaling will create an instance when the application receives a request. The instance will be turned down when the app becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
	BasicScaling BasicScalingResponse `pulumi:"basicScaling"`
	// Metadata settings that are supplied to this version to enable beta runtime features.
	BetaSettings map[string]string `pulumi:"betaSettings"`
	// Environment variables available to the build environment.Only returned in GET requests if view=FULL is set.
	BuildEnvVariables map[string]string `pulumi:"buildEnvVariables"`
	// Time that this version was created.
	CreateTime string `pulumi:"createTime"`
	// Email address of the user who created this version.
	CreatedBy string `pulumi:"createdBy"`
	// Duration that static files should be cached by web proxies and browsers. Only applicable if the corresponding StaticFilesHandler (https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1beta/apps.services.versions#StaticFilesHandler) does not specify its own expiration time.Only returned in GET requests if view=FULL is set.
	DefaultExpiration string `pulumi:"defaultExpiration"`
	// Code and application artifacts that make up this version.Only returned in GET requests if view=FULL is set.
	Deployment DeploymentResponse `pulumi:"deployment"`
	// Total size in bytes of all the files that are included in this version and currently hosted on the App Engine disk.
	DiskUsageBytes string `pulumi:"diskUsageBytes"`
	// Cloud Endpoints configuration.If endpoints_api_service is set, the Cloud Endpoints Extensible Service Proxy will be provided to serve the API implemented by the app.
	EndpointsApiService EndpointsApiServiceResponse `pulumi:"endpointsApiService"`
	// The entrypoint for the application.
	Entrypoint EntrypointResponse `pulumi:"entrypoint"`
	// App Engine execution environment for this version.Defaults to standard.
	Env string `pulumi:"env"`
	// Environment variables available to the application.Only returned in GET requests if view=FULL is set.
	EnvVariables map[string]string `pulumi:"envVariables"`
	// Custom static error pages. Limited to 10KB per page.Only returned in GET requests if view=FULL is set.
	ErrorHandlers []ErrorHandlerResponse `pulumi:"errorHandlers"`
	// Settings for App Engine flexible runtimes.
	FlexibleRuntimeSettings FlexibleRuntimeSettingsResponse `pulumi:"flexibleRuntimeSettings"`
	// An ordered list of URL-matching patterns that should be applied to incoming requests. The first matching URL handles the request and other request handlers are not attempted.Only returned in GET requests if view=FULL is set.
	Handlers []UrlMapResponse `pulumi:"handlers"`
	// Configures health checking for instances. Unhealthy instances are stopped and replaced with new instances. Only applicable in the App Engine flexible environment.Only returned in GET requests if view=FULL is set.
	HealthCheck HealthCheckResponse `pulumi:"healthCheck"`
	// Before an application can receive email or XMPP messages, the application must be configured to enable the service.
	InboundServices []string `pulumi:"inboundServices"`
	// Instance class that is used to run this version. Valid values are: AutomaticScaling: F1, F2, F4, F4_1G ManualScaling or BasicScaling: B1, B2, B4, B8, B4_1GDefaults to F1 for AutomaticScaling and B1 for ManualScaling or BasicScaling.
	InstanceClass string `pulumi:"instanceClass"`
	// Configuration for third-party Python runtime libraries that are required by the application.Only returned in GET requests if view=FULL is set.
	Libraries []LibraryResponse `pulumi:"libraries"`
	// Configures liveness health checking for instances. Unhealthy instances are stopped and replaced with new instancesOnly returned in GET requests if view=FULL is set.
	LivenessCheck LivenessCheckResponse `pulumi:"livenessCheck"`
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time. Manually scaled versions are sometimes referred to as "backends".
	ManualScaling ManualScalingResponse `pulumi:"manualScaling"`
	// Full path to the Version resource in the API. Example: apps/myapp/services/default/versions/v1.
	Name string `pulumi:"name"`
	// Extra network settings. Only applicable in the App Engine flexible environment.
	Network NetworkResponse `pulumi:"network"`
	// Files that match this pattern will not be built into this version. Only applicable for Go runtimes.Only returned in GET requests if view=FULL is set.
	NobuildFilesRegex string `pulumi:"nobuildFilesRegex"`
	// Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.Only returned in GET requests if view=FULL is set.
	ReadinessCheck ReadinessCheckResponse `pulumi:"readinessCheck"`
	// Machine resources for this version. Only applicable in the App Engine flexible environment.
	Resources ResourcesResponse `pulumi:"resources"`
	// Desired runtime. Example: python27.
	Runtime string `pulumi:"runtime"`
	// The version of the API in the given runtime environment. Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
	RuntimeApiVersion string `pulumi:"runtimeApiVersion"`
	// The channel of the runtime to use. Only available for some runtimes. Defaults to the default channel.
	RuntimeChannel string `pulumi:"runtimeChannel"`
	// The path or name of the app's main executable.
	RuntimeMainExecutablePath string `pulumi:"runtimeMainExecutablePath"`
	// The identity that the deployed version will run as. Admin API will use the App Engine Appspot service account as default if this field is neither provided in app.yaml file nor through CLI flag.
	ServiceAccount string `pulumi:"serviceAccount"`
	// Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.SERVING_STATUS_UNSPECIFIED is an invalid value. Defaults to SERVING.
	ServingStatus string `pulumi:"servingStatus"`
	// Whether multiple requests can be dispatched to this version at once.
	Threadsafe bool `pulumi:"threadsafe"`
	// Serving URL for this version. Example: "https://myversion-dot-myservice-dot-myapp.appspot.com"
	VersionUrl string `pulumi:"versionUrl"`
	// Whether to deploy this version in a container on a virtual machine.
	Vm bool `pulumi:"vm"`
	// Enables VPC connectivity for standard apps.
	VpcAccessConnector VpcAccessConnectorResponse `pulumi:"vpcAccessConnector"`
	// The Google Compute Engine zones that are supported by this version in the App Engine flexible environment. Deprecated.
	//
	// Deprecated: The Google Compute Engine zones that are supported by this version in the App Engine flexible environment. Deprecated.
	Zones []string `pulumi:"zones"`
}

func LookupVersionOutput(ctx *pulumi.Context, args LookupVersionOutputArgs, opts ...pulumi.InvokeOption) LookupVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVersionResult, error) {
			args := v.(LookupVersionArgs)
			r, err := LookupVersion(ctx, &args, opts...)
			var s LookupVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVersionResultOutput)
}

type LookupVersionOutputArgs struct {
	AppId     pulumi.StringInput    `pulumi:"appId"`
	ServiceId pulumi.StringInput    `pulumi:"serviceId"`
	VersionId pulumi.StringInput    `pulumi:"versionId"`
	View      pulumi.StringPtrInput `pulumi:"view"`
}

func (LookupVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVersionArgs)(nil)).Elem()
}

type LookupVersionResultOutput struct{ *pulumi.OutputState }

func (LookupVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVersionResult)(nil)).Elem()
}

func (o LookupVersionResultOutput) ToLookupVersionResultOutput() LookupVersionResultOutput {
	return o
}

func (o LookupVersionResultOutput) ToLookupVersionResultOutputWithContext(ctx context.Context) LookupVersionResultOutput {
	return o
}

// Serving configuration for Google Cloud Endpoints (https://cloud.google.com/endpoints).Only returned in GET requests if view=FULL is set.
func (o LookupVersionResultOutput) ApiConfig() ApiConfigHandlerResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) ApiConfigHandlerResponse { return v.ApiConfig }).(ApiConfigHandlerResponseOutput)
}

// Allows App Engine second generation runtimes to access the legacy bundled services.
func (o LookupVersionResultOutput) AppEngineApis() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVersionResult) bool { return v.AppEngineApis }).(pulumi.BoolOutput)
}

// Automatic scaling is based on request rate, response latencies, and other application metrics. Instances are dynamically created and destroyed as needed in order to handle traffic.
func (o LookupVersionResultOutput) AutomaticScaling() AutomaticScalingResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) AutomaticScalingResponse { return v.AutomaticScaling }).(AutomaticScalingResponseOutput)
}

// A service with basic scaling will create an instance when the application receives a request. The instance will be turned down when the app becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
func (o LookupVersionResultOutput) BasicScaling() BasicScalingResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) BasicScalingResponse { return v.BasicScaling }).(BasicScalingResponseOutput)
}

// Metadata settings that are supplied to this version to enable beta runtime features.
func (o LookupVersionResultOutput) BetaSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVersionResult) map[string]string { return v.BetaSettings }).(pulumi.StringMapOutput)
}

// Environment variables available to the build environment.Only returned in GET requests if view=FULL is set.
func (o LookupVersionResultOutput) BuildEnvVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVersionResult) map[string]string { return v.BuildEnvVariables }).(pulumi.StringMapOutput)
}

// Time that this version was created.
func (o LookupVersionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Email address of the user who created this version.
func (o LookupVersionResultOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.CreatedBy }).(pulumi.StringOutput)
}

// Duration that static files should be cached by web proxies and browsers. Only applicable if the corresponding StaticFilesHandler (https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1beta/apps.services.versions#StaticFilesHandler) does not specify its own expiration time.Only returned in GET requests if view=FULL is set.
func (o LookupVersionResultOutput) DefaultExpiration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.DefaultExpiration }).(pulumi.StringOutput)
}

// Code and application artifacts that make up this version.Only returned in GET requests if view=FULL is set.
func (o LookupVersionResultOutput) Deployment() DeploymentResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) DeploymentResponse { return v.Deployment }).(DeploymentResponseOutput)
}

// Total size in bytes of all the files that are included in this version and currently hosted on the App Engine disk.
func (o LookupVersionResultOutput) DiskUsageBytes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.DiskUsageBytes }).(pulumi.StringOutput)
}

// Cloud Endpoints configuration.If endpoints_api_service is set, the Cloud Endpoints Extensible Service Proxy will be provided to serve the API implemented by the app.
func (o LookupVersionResultOutput) EndpointsApiService() EndpointsApiServiceResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) EndpointsApiServiceResponse { return v.EndpointsApiService }).(EndpointsApiServiceResponseOutput)
}

// The entrypoint for the application.
func (o LookupVersionResultOutput) Entrypoint() EntrypointResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) EntrypointResponse { return v.Entrypoint }).(EntrypointResponseOutput)
}

// App Engine execution environment for this version.Defaults to standard.
func (o LookupVersionResultOutput) Env() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.Env }).(pulumi.StringOutput)
}

// Environment variables available to the application.Only returned in GET requests if view=FULL is set.
func (o LookupVersionResultOutput) EnvVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVersionResult) map[string]string { return v.EnvVariables }).(pulumi.StringMapOutput)
}

// Custom static error pages. Limited to 10KB per page.Only returned in GET requests if view=FULL is set.
func (o LookupVersionResultOutput) ErrorHandlers() ErrorHandlerResponseArrayOutput {
	return o.ApplyT(func(v LookupVersionResult) []ErrorHandlerResponse { return v.ErrorHandlers }).(ErrorHandlerResponseArrayOutput)
}

// Settings for App Engine flexible runtimes.
func (o LookupVersionResultOutput) FlexibleRuntimeSettings() FlexibleRuntimeSettingsResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) FlexibleRuntimeSettingsResponse { return v.FlexibleRuntimeSettings }).(FlexibleRuntimeSettingsResponseOutput)
}

// An ordered list of URL-matching patterns that should be applied to incoming requests. The first matching URL handles the request and other request handlers are not attempted.Only returned in GET requests if view=FULL is set.
func (o LookupVersionResultOutput) Handlers() UrlMapResponseArrayOutput {
	return o.ApplyT(func(v LookupVersionResult) []UrlMapResponse { return v.Handlers }).(UrlMapResponseArrayOutput)
}

// Configures health checking for instances. Unhealthy instances are stopped and replaced with new instances. Only applicable in the App Engine flexible environment.Only returned in GET requests if view=FULL is set.
func (o LookupVersionResultOutput) HealthCheck() HealthCheckResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) HealthCheckResponse { return v.HealthCheck }).(HealthCheckResponseOutput)
}

// Before an application can receive email or XMPP messages, the application must be configured to enable the service.
func (o LookupVersionResultOutput) InboundServices() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVersionResult) []string { return v.InboundServices }).(pulumi.StringArrayOutput)
}

// Instance class that is used to run this version. Valid values are: AutomaticScaling: F1, F2, F4, F4_1G ManualScaling or BasicScaling: B1, B2, B4, B8, B4_1GDefaults to F1 for AutomaticScaling and B1 for ManualScaling or BasicScaling.
func (o LookupVersionResultOutput) InstanceClass() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.InstanceClass }).(pulumi.StringOutput)
}

// Configuration for third-party Python runtime libraries that are required by the application.Only returned in GET requests if view=FULL is set.
func (o LookupVersionResultOutput) Libraries() LibraryResponseArrayOutput {
	return o.ApplyT(func(v LookupVersionResult) []LibraryResponse { return v.Libraries }).(LibraryResponseArrayOutput)
}

// Configures liveness health checking for instances. Unhealthy instances are stopped and replaced with new instancesOnly returned in GET requests if view=FULL is set.
func (o LookupVersionResultOutput) LivenessCheck() LivenessCheckResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) LivenessCheckResponse { return v.LivenessCheck }).(LivenessCheckResponseOutput)
}

// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time. Manually scaled versions are sometimes referred to as "backends".
func (o LookupVersionResultOutput) ManualScaling() ManualScalingResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) ManualScalingResponse { return v.ManualScaling }).(ManualScalingResponseOutput)
}

// Full path to the Version resource in the API. Example: apps/myapp/services/default/versions/v1.
func (o LookupVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Extra network settings. Only applicable in the App Engine flexible environment.
func (o LookupVersionResultOutput) Network() NetworkResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) NetworkResponse { return v.Network }).(NetworkResponseOutput)
}

// Files that match this pattern will not be built into this version. Only applicable for Go runtimes.Only returned in GET requests if view=FULL is set.
func (o LookupVersionResultOutput) NobuildFilesRegex() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.NobuildFilesRegex }).(pulumi.StringOutput)
}

// Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.Only returned in GET requests if view=FULL is set.
func (o LookupVersionResultOutput) ReadinessCheck() ReadinessCheckResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) ReadinessCheckResponse { return v.ReadinessCheck }).(ReadinessCheckResponseOutput)
}

// Machine resources for this version. Only applicable in the App Engine flexible environment.
func (o LookupVersionResultOutput) Resources() ResourcesResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) ResourcesResponse { return v.Resources }).(ResourcesResponseOutput)
}

// Desired runtime. Example: python27.
func (o LookupVersionResultOutput) Runtime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.Runtime }).(pulumi.StringOutput)
}

// The version of the API in the given runtime environment. Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
func (o LookupVersionResultOutput) RuntimeApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.RuntimeApiVersion }).(pulumi.StringOutput)
}

// The channel of the runtime to use. Only available for some runtimes. Defaults to the default channel.
func (o LookupVersionResultOutput) RuntimeChannel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.RuntimeChannel }).(pulumi.StringOutput)
}

// The path or name of the app's main executable.
func (o LookupVersionResultOutput) RuntimeMainExecutablePath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.RuntimeMainExecutablePath }).(pulumi.StringOutput)
}

// The identity that the deployed version will run as. Admin API will use the App Engine Appspot service account as default if this field is neither provided in app.yaml file nor through CLI flag.
func (o LookupVersionResultOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.ServiceAccount }).(pulumi.StringOutput)
}

// Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.SERVING_STATUS_UNSPECIFIED is an invalid value. Defaults to SERVING.
func (o LookupVersionResultOutput) ServingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.ServingStatus }).(pulumi.StringOutput)
}

// Whether multiple requests can be dispatched to this version at once.
func (o LookupVersionResultOutput) Threadsafe() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVersionResult) bool { return v.Threadsafe }).(pulumi.BoolOutput)
}

// Serving URL for this version. Example: "https://myversion-dot-myservice-dot-myapp.appspot.com"
func (o LookupVersionResultOutput) VersionUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.VersionUrl }).(pulumi.StringOutput)
}

// Whether to deploy this version in a container on a virtual machine.
func (o LookupVersionResultOutput) Vm() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVersionResult) bool { return v.Vm }).(pulumi.BoolOutput)
}

// Enables VPC connectivity for standard apps.
func (o LookupVersionResultOutput) VpcAccessConnector() VpcAccessConnectorResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) VpcAccessConnectorResponse { return v.VpcAccessConnector }).(VpcAccessConnectorResponseOutput)
}

// The Google Compute Engine zones that are supported by this version in the App Engine flexible environment. Deprecated.
//
// Deprecated: The Google Compute Engine zones that are supported by this version in the App Engine flexible environment. Deprecated.
func (o LookupVersionResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVersionResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVersionResultOutput{})
}
