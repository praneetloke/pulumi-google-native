// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single Hub.
func LookupHub(ctx *pulumi.Context, args *LookupHubArgs, opts ...pulumi.InvokeOption) (*LookupHubResult, error) {
	var rv LookupHubResult
	err := ctx.Invoke("google-native:networkconnectivity/v1alpha1:getHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHubArgs struct {
	HubId   string  `pulumi:"hubId"`
	Project *string `pulumi:"project"`
}

type LookupHubResult struct {
	// Time when the Hub was created.
	CreateTime string `pulumi:"createTime"`
	// Short description of the hub resource.
	Description string `pulumi:"description"`
	// User-defined labels.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. The name of a Hub resource.
	Name string `pulumi:"name"`
	// The current lifecycle state of this Hub.
	State string `pulumi:"state"`
	// Google-generated UUID for this resource. This is unique across all Hub resources. If a Hub resource is deleted and another with the same name is created, it gets a different unique_id.
	UniqueId string `pulumi:"uniqueId"`
	// Time when the Hub was updated.
	UpdateTime string `pulumi:"updateTime"`
}