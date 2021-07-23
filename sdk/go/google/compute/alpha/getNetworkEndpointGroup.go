// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified network endpoint group. Gets a list of available network endpoint groups by making a list() request.
func LookupNetworkEndpointGroup(ctx *pulumi.Context, args *LookupNetworkEndpointGroupArgs, opts ...pulumi.InvokeOption) (*LookupNetworkEndpointGroupResult, error) {
	var rv LookupNetworkEndpointGroupResult
	err := ctx.Invoke("google-native:compute/alpha:getNetworkEndpointGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkEndpointGroupArgs struct {
	NetworkEndpointGroup string  `pulumi:"networkEndpointGroup"`
	Project              *string `pulumi:"project"`
	Zone                 string  `pulumi:"zone"`
}

type LookupNetworkEndpointGroupResult struct {
	// Metadata defined as annotations on the network endpoint group.
	Annotations map[string]string `pulumi:"annotations"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	AppEngine NetworkEndpointGroupAppEngineResponse `pulumi:"appEngine"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudFunction NetworkEndpointGroupCloudFunctionResponse `pulumi:"cloudFunction"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudRun NetworkEndpointGroupCloudRunResponse `pulumi:"cloudRun"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// The default port used if the port number is not specified in the network endpoint.
	DefaultPort int `pulumi:"defaultPort"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
	Kind string `pulumi:"kind"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
	Network string `pulumi:"network"`
	// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
	NetworkEndpointType string `pulumi:"networkEndpointType"`
	// The target service url used to set up private service connection to a Google API. An example value is: "asia-northeast3-cloudkms.googleapis.com"
	PscTargetService string `pulumi:"pscTargetService"`
	// The URL of the region where the network endpoint group is located.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine, cloudFunction or serverlessDeployment may be set.
	ServerlessDeployment NetworkEndpointGroupServerlessDeploymentResponse `pulumi:"serverlessDeployment"`
	// [Output only] Number of network endpoints in the network endpoint group.
	Size int `pulumi:"size"`
	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	Subnetwork string `pulumi:"subnetwork"`
	// Specify the type of this network endpoint group. Only LOAD_BALANCING is valid for now.
	Type string `pulumi:"type"`
	// The URL of the zone where the network endpoint group is located.
	Zone string `pulumi:"zone"`
}