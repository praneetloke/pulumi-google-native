// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Service Perimeter by resource name.
func LookupServicePerimeter(ctx *pulumi.Context, args *LookupServicePerimeterArgs, opts ...pulumi.InvokeOption) (*LookupServicePerimeterResult, error) {
	var rv LookupServicePerimeterResult
	err := ctx.Invoke("google-native:accesscontextmanager/v1beta:getServicePerimeter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServicePerimeterArgs struct {
	AccessPolicyId     string `pulumi:"accessPolicyId"`
	ServicePerimeterId string `pulumi:"servicePerimeterId"`
}

type LookupServicePerimeterResult struct {
	// Description of the `ServicePerimeter` and its use. Does not affect behavior.
	Description string `pulumi:"description"`
	// Resource name for the ServicePerimeter. The `short_name` component must begin with a letter and only include alphanumeric and '_'. Format: `accessPolicies/{policy_id}/servicePerimeters/{short_name}`
	Name string `pulumi:"name"`
	// Perimeter type indicator. A single project is allowed to be a member of single regular perimeter, but multiple service perimeter bridges. A project cannot be a included in a perimeter bridge without being included in regular perimeter. For perimeter bridges, restricted/unrestricted service lists as well as access lists must be empty.
	PerimeterType string `pulumi:"perimeterType"`
	// Current ServicePerimeter configuration. Specifies sets of resources, restricted/unrestricted services and access levels that determine perimeter content and boundaries.
	Status ServicePerimeterConfigResponse `pulumi:"status"`
	// Human readable title. Must be unique within the Policy.
	Title string `pulumi:"title"`
}