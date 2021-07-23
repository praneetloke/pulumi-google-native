// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an environment group attachment.
func LookupEnvgroupAttachment(ctx *pulumi.Context, args *LookupEnvgroupAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvgroupAttachmentResult, error) {
	var rv LookupEnvgroupAttachmentResult
	err := ctx.Invoke("google-native:apigee/v1:getEnvgroupAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvgroupAttachmentArgs struct {
	AttachmentId   string `pulumi:"attachmentId"`
	EnvgroupId     string `pulumi:"envgroupId"`
	OrganizationId string `pulumi:"organizationId"`
}

type LookupEnvgroupAttachmentResult struct {
	// The time at which the environment group attachment was created as milliseconds since epoch.
	CreatedAt string `pulumi:"createdAt"`
	// ID of the attached environment.
	Environment string `pulumi:"environment"`
	// ID of the environment group attachment.
	Name string `pulumi:"name"`
}