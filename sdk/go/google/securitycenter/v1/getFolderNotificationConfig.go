// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a notification config.
func LookupFolderNotificationConfig(ctx *pulumi.Context, args *LookupFolderNotificationConfigArgs, opts ...pulumi.InvokeOption) (*LookupFolderNotificationConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFolderNotificationConfigResult
	err := ctx.Invoke("google-native:securitycenter/v1:getFolderNotificationConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFolderNotificationConfigArgs struct {
	FolderId             string `pulumi:"folderId"`
	NotificationConfigId string `pulumi:"notificationConfigId"`
}

type LookupFolderNotificationConfigResult struct {
	// The description of the notification config (max of 1024 characters).
	Description string `pulumi:"description"`
	// The relative resource name of this notification config. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/notificationConfigs/notify_public_bucket", "folders/{folder_id}/notificationConfigs/notify_public_bucket", or "projects/{project_id}/notificationConfigs/notify_public_bucket".
	Name string `pulumi:"name"`
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic string `pulumi:"pubsubTopic"`
	// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
	ServiceAccount string `pulumi:"serviceAccount"`
	// The config for triggering streaming-based notifications.
	StreamingConfig StreamingConfigResponse `pulumi:"streamingConfig"`
}

func LookupFolderNotificationConfigOutput(ctx *pulumi.Context, args LookupFolderNotificationConfigOutputArgs, opts ...pulumi.InvokeOption) LookupFolderNotificationConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFolderNotificationConfigResult, error) {
			args := v.(LookupFolderNotificationConfigArgs)
			r, err := LookupFolderNotificationConfig(ctx, &args, opts...)
			var s LookupFolderNotificationConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFolderNotificationConfigResultOutput)
}

type LookupFolderNotificationConfigOutputArgs struct {
	FolderId             pulumi.StringInput `pulumi:"folderId"`
	NotificationConfigId pulumi.StringInput `pulumi:"notificationConfigId"`
}

func (LookupFolderNotificationConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderNotificationConfigArgs)(nil)).Elem()
}

type LookupFolderNotificationConfigResultOutput struct{ *pulumi.OutputState }

func (LookupFolderNotificationConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderNotificationConfigResult)(nil)).Elem()
}

func (o LookupFolderNotificationConfigResultOutput) ToLookupFolderNotificationConfigResultOutput() LookupFolderNotificationConfigResultOutput {
	return o
}

func (o LookupFolderNotificationConfigResultOutput) ToLookupFolderNotificationConfigResultOutputWithContext(ctx context.Context) LookupFolderNotificationConfigResultOutput {
	return o
}

// The description of the notification config (max of 1024 characters).
func (o LookupFolderNotificationConfigResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderNotificationConfigResult) string { return v.Description }).(pulumi.StringOutput)
}

// The relative resource name of this notification config. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/notificationConfigs/notify_public_bucket", "folders/{folder_id}/notificationConfigs/notify_public_bucket", or "projects/{project_id}/notificationConfigs/notify_public_bucket".
func (o LookupFolderNotificationConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderNotificationConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
func (o LookupFolderNotificationConfigResultOutput) PubsubTopic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderNotificationConfigResult) string { return v.PubsubTopic }).(pulumi.StringOutput)
}

// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
func (o LookupFolderNotificationConfigResultOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderNotificationConfigResult) string { return v.ServiceAccount }).(pulumi.StringOutput)
}

// The config for triggering streaming-based notifications.
func (o LookupFolderNotificationConfigResultOutput) StreamingConfig() StreamingConfigResponseOutput {
	return o.ApplyT(func(v LookupFolderNotificationConfigResult) StreamingConfigResponse { return v.StreamingConfig }).(StreamingConfigResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFolderNotificationConfigResultOutput{})
}
