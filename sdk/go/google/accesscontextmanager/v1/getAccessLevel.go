// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get an Access Level by resource name.
func LookupAccessLevel(ctx *pulumi.Context, args *LookupAccessLevelArgs, opts ...pulumi.InvokeOption) (*LookupAccessLevelResult, error) {
	var rv LookupAccessLevelResult
	err := ctx.Invoke("google-native:accesscontextmanager/v1:getAccessLevel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessLevelArgs struct {
	AccessLevelFormat *string `pulumi:"accessLevelFormat"`
	AccessLevelId     string  `pulumi:"accessLevelId"`
	AccessPolicyId    string  `pulumi:"accessPolicyId"`
}

type LookupAccessLevelResult struct {
	// A `BasicLevel` composed of `Conditions`.
	Basic BasicLevelResponse `pulumi:"basic"`
	// A `CustomLevel` written in the Common Expression Language.
	Custom CustomLevelResponse `pulumi:"custom"`
	// Description of the `AccessLevel` and its use. Does not affect behavior.
	Description string `pulumi:"description"`
	// Resource name for the Access Level. The `short_name` component must begin with a letter and only include alphanumeric and '_'. Format: `accessPolicies/{policy_id}/accessLevels/{short_name}`. The maximum length of the `short_name` component is 50 characters.
	Name string `pulumi:"name"`
	// Human readable title. Must be unique within the Policy.
	Title string `pulumi:"title"`
}