// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a function with the given name from the requested project.
func LookupFunction(ctx *pulumi.Context, args *LookupFunctionArgs, opts ...pulumi.InvokeOption) (*LookupFunctionResult, error) {
	var rv LookupFunctionResult
	err := ctx.Invoke("google-native:cloudfunctions/v1:getFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFunctionArgs struct {
	FunctionId string  `pulumi:"functionId"`
	Location   string  `pulumi:"location"`
	Project    *string `pulumi:"project"`
}

type LookupFunctionResult struct {
	// The amount of memory in MB available for a function. Defaults to 256MB.
	AvailableMemoryMb int `pulumi:"availableMemoryMb"`
	// Build environment variables that shall be available during build time.
	BuildEnvironmentVariables map[string]string `pulumi:"buildEnvironmentVariables"`
	// The Cloud Build ID of the latest successful deployment of the function.
	BuildId string `pulumi:"buildId"`
	// Name of the Cloud Build Custom Worker Pool that should be used to build the function. The format of this field is `projects/{project}/locations/{region}/workerPools/{workerPool}` where {project} and {region} are the project id and region respectively where the worker pool is defined and {workerPool} is the short name of the worker pool. If the project id is not the same as the function, then the Cloud Functions Service Agent (service-@gcf-admin-robot.iam.gserviceaccount.com) must be granted the role Cloud Build Custom Workers Builder (roles/cloudbuild.customworkers.builder) in the project.
	BuildWorkerPool string `pulumi:"buildWorkerPool"`
	// User-provided description of a function.
	Description string `pulumi:"description"`
	// The name of the function (as defined in source code) that will be executed. Defaults to the resource name suffix, if not specified. For backward compatibility, if function with given name is not found, then the system will try to use function named "function". For Node.js this is name of a function exported by the module specified in `source_location`.
	EntryPoint string `pulumi:"entryPoint"`
	// Environment variables that shall be available during function execution.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// A source that fires events in response to a condition in another service.
	EventTrigger EventTriggerResponse `pulumi:"eventTrigger"`
	// An HTTPS endpoint type of source that can be triggered via URL.
	HttpsTrigger HttpsTriggerResponse `pulumi:"httpsTrigger"`
	// The ingress settings for the function, controlling what traffic can reach it.
	IngressSettings string `pulumi:"ingressSettings"`
	// Labels associated with this Cloud Function.
	Labels map[string]string `pulumi:"labels"`
	// The limit on the maximum number of function instances that may coexist at a given time. In some cases, such as rapid traffic surges, Cloud Functions may, for a short period of time, create more instances than the specified max instances limit. If your function cannot tolerate this temporary behavior, you may want to factor in a safety margin and set a lower max instances value than your function can tolerate. See the [Max Instances](https://cloud.google.com/functions/docs/max-instances) Guide for more details.
	MaxInstances int `pulumi:"maxInstances"`
	// A user-defined name of the function. Function names must be unique globally and match pattern `projects/*/locations/*/functions/*`
	Name string `pulumi:"name"`
	// The VPC Network that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network resource. If the short network name is used, the network must belong to the same project. Otherwise, it must belong to a project within the same organization. The format of this field is either `projects/{project}/global/networks/{network}` or `{network}`, where {project} is a project id where the network is defined, and {network} is the short name of the network. This field is mutually exclusive with `vpc_connector` and will be replaced by it. See [the VPC documentation](https://cloud.google.com/compute/docs/vpc) for more information on connecting Cloud projects.
	Network string `pulumi:"network"`
	// The runtime in which to run the function. Required when deploying a new function, optional when updating an existing function. For a complete list of possible choices, see the [`gcloud` command reference](/sdk/gcloud/reference/functions/deploy#--runtime).
	Runtime string `pulumi:"runtime"`
	// Secret environment variables configuration.
	SecretEnvironmentVariables []SecretEnvVarResponse `pulumi:"secretEnvironmentVariables"`
	// Secret volumes configuration.
	SecretVolumes []SecretVolumeResponse `pulumi:"secretVolumes"`
	// The email of the function's service account. If empty, defaults to `{project_id}@appspot.gserviceaccount.com`.
	ServiceAccountEmail string `pulumi:"serviceAccountEmail"`
	// The Google Cloud Storage URL, starting with gs://, pointing to the zip archive which contains the function.
	SourceArchiveUrl string `pulumi:"sourceArchiveUrl"`
	// **Beta Feature** The source repository where a function is hosted.
	SourceRepository SourceRepositoryResponse `pulumi:"sourceRepository"`
	// Input only. An identifier for Firebase function sources. Disclaimer: This field is only supported for Firebase function deployments.
	SourceToken string `pulumi:"sourceToken"`
	// The Google Cloud Storage signed URL used for source uploading, generated by google.cloud.functions.v1.GenerateUploadUrl
	SourceUploadUrl string `pulumi:"sourceUploadUrl"`
	// Status of the function deployment.
	Status string `pulumi:"status"`
	// The function execution timeout. Execution is considered failed and can be terminated if the function is not completed at the end of the timeout period. Defaults to 60 seconds.
	Timeout string `pulumi:"timeout"`
	// The last update timestamp of a Cloud Function.
	UpdateTime string `pulumi:"updateTime"`
	// The version identifier of the Cloud Function. Each deployment attempt results in a new version of a function being created.
	VersionId string `pulumi:"versionId"`
	// The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*/locations/*/connectors/*` This field is mutually exclusive with `network` field and will eventually replace it. See [the VPC documentation](https://cloud.google.com/compute/docs/vpc) for more information on connecting Cloud projects.
	VpcConnector string `pulumi:"vpcConnector"`
	// The egress settings for the connector, controlling what traffic is diverted through it.
	VpcConnectorEgressSettings string `pulumi:"vpcConnectorEgressSettings"`
}