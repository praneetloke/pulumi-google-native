// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a single Uptime check configuration.
func LookupUptimeCheckConfig(ctx *pulumi.Context, args *LookupUptimeCheckConfigArgs, opts ...pulumi.InvokeOption) (*LookupUptimeCheckConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUptimeCheckConfigResult
	err := ctx.Invoke("google-native:monitoring/v3:getUptimeCheckConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUptimeCheckConfigArgs struct {
	Project             *string `pulumi:"project"`
	UptimeCheckConfigId string  `pulumi:"uptimeCheckConfigId"`
}

type LookupUptimeCheckConfigResult struct {
	// The type of checkers to use to execute the Uptime check.
	CheckerType string `pulumi:"checkerType"`
	// The content that is expected to appear in the data returned by the target server against which the check is run. Currently, only the first entry in the content_matchers list is supported, and additional entries will be ignored. This field is optional and should only be specified if a content match is required as part of the/ Uptime check.
	ContentMatchers []ContentMatcherResponse `pulumi:"contentMatchers"`
	// A human-friendly name for the Uptime check configuration. The display name should be unique within a Cloud Monitoring Workspace in order to make it easier to identify; however, uniqueness is not enforced. Required.
	DisplayName string `pulumi:"displayName"`
	// Contains information needed to make an HTTP or HTTPS check.
	HttpCheck HttpCheckResponse `pulumi:"httpCheck"`
	// The internal checkers that this check will egress from. If is_internal is true and this list is empty, the check will egress from all the InternalCheckers configured for the project that owns this UptimeCheckConfig.
	InternalCheckers []InternalCheckerResponse `pulumi:"internalCheckers"`
	// If this is true, then checks are made only from the 'internal_checkers'. If it is false, then checks are made only from the 'selected_regions'. It is an error to provide 'selected_regions' when is_internal is true, or to provide 'internal_checkers' when is_internal is false.
	IsInternal bool `pulumi:"isInternal"`
	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are valid for this field: uptime_url, gce_instance, gae_app, aws_ec2_instance, aws_elb_load_balancer k8s_service servicedirectory_service cloud_run_revision
	MonitoredResource MonitoredResourceResponse `pulumi:"monitoredResource"`
	// A unique resource name for this Uptime check configuration. The format is: projects/[PROJECT_ID_OR_NUMBER]/uptimeCheckConfigs/[UPTIME_CHECK_ID] [PROJECT_ID_OR_NUMBER] is the Workspace host project associated with the Uptime check.This field should be omitted when creating the Uptime check configuration; on create, the resource name is assigned by the server and included in the response.
	Name string `pulumi:"name"`
	// How often, in seconds, the Uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 60s.
	Period string `pulumi:"period"`
	// The group resource associated with the configuration.
	ResourceGroup ResourceGroupResponse `pulumi:"resourceGroup"`
	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions must be provided to include a minimum of 3 locations. Not specifying this field will result in Uptime checks running from all available regions.
	SelectedRegions []string `pulumi:"selectedRegions"`
	// Contains information needed to make a TCP check.
	TcpCheck TcpCheckResponse `pulumi:"tcpCheck"`
	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Required.
	Timeout string `pulumi:"timeout"`
	// User-supplied key/value data to be used for organizing and identifying the UptimeCheckConfig objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels map[string]string `pulumi:"userLabels"`
}

func LookupUptimeCheckConfigOutput(ctx *pulumi.Context, args LookupUptimeCheckConfigOutputArgs, opts ...pulumi.InvokeOption) LookupUptimeCheckConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUptimeCheckConfigResult, error) {
			args := v.(LookupUptimeCheckConfigArgs)
			r, err := LookupUptimeCheckConfig(ctx, &args, opts...)
			var s LookupUptimeCheckConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUptimeCheckConfigResultOutput)
}

type LookupUptimeCheckConfigOutputArgs struct {
	Project             pulumi.StringPtrInput `pulumi:"project"`
	UptimeCheckConfigId pulumi.StringInput    `pulumi:"uptimeCheckConfigId"`
}

func (LookupUptimeCheckConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUptimeCheckConfigArgs)(nil)).Elem()
}

type LookupUptimeCheckConfigResultOutput struct{ *pulumi.OutputState }

func (LookupUptimeCheckConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUptimeCheckConfigResult)(nil)).Elem()
}

func (o LookupUptimeCheckConfigResultOutput) ToLookupUptimeCheckConfigResultOutput() LookupUptimeCheckConfigResultOutput {
	return o
}

func (o LookupUptimeCheckConfigResultOutput) ToLookupUptimeCheckConfigResultOutputWithContext(ctx context.Context) LookupUptimeCheckConfigResultOutput {
	return o
}

// The type of checkers to use to execute the Uptime check.
func (o LookupUptimeCheckConfigResultOutput) CheckerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUptimeCheckConfigResult) string { return v.CheckerType }).(pulumi.StringOutput)
}

// The content that is expected to appear in the data returned by the target server against which the check is run. Currently, only the first entry in the content_matchers list is supported, and additional entries will be ignored. This field is optional and should only be specified if a content match is required as part of the/ Uptime check.
func (o LookupUptimeCheckConfigResultOutput) ContentMatchers() ContentMatcherResponseArrayOutput {
	return o.ApplyT(func(v LookupUptimeCheckConfigResult) []ContentMatcherResponse { return v.ContentMatchers }).(ContentMatcherResponseArrayOutput)
}

// A human-friendly name for the Uptime check configuration. The display name should be unique within a Cloud Monitoring Workspace in order to make it easier to identify; however, uniqueness is not enforced. Required.
func (o LookupUptimeCheckConfigResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUptimeCheckConfigResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Contains information needed to make an HTTP or HTTPS check.
func (o LookupUptimeCheckConfigResultOutput) HttpCheck() HttpCheckResponseOutput {
	return o.ApplyT(func(v LookupUptimeCheckConfigResult) HttpCheckResponse { return v.HttpCheck }).(HttpCheckResponseOutput)
}

// The internal checkers that this check will egress from. If is_internal is true and this list is empty, the check will egress from all the InternalCheckers configured for the project that owns this UptimeCheckConfig.
func (o LookupUptimeCheckConfigResultOutput) InternalCheckers() InternalCheckerResponseArrayOutput {
	return o.ApplyT(func(v LookupUptimeCheckConfigResult) []InternalCheckerResponse { return v.InternalCheckers }).(InternalCheckerResponseArrayOutput)
}

// If this is true, then checks are made only from the 'internal_checkers'. If it is false, then checks are made only from the 'selected_regions'. It is an error to provide 'selected_regions' when is_internal is true, or to provide 'internal_checkers' when is_internal is false.
func (o LookupUptimeCheckConfigResultOutput) IsInternal() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupUptimeCheckConfigResult) bool { return v.IsInternal }).(pulumi.BoolOutput)
}

// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are valid for this field: uptime_url, gce_instance, gae_app, aws_ec2_instance, aws_elb_load_balancer k8s_service servicedirectory_service cloud_run_revision
func (o LookupUptimeCheckConfigResultOutput) MonitoredResource() MonitoredResourceResponseOutput {
	return o.ApplyT(func(v LookupUptimeCheckConfigResult) MonitoredResourceResponse { return v.MonitoredResource }).(MonitoredResourceResponseOutput)
}

// A unique resource name for this Uptime check configuration. The format is: projects/[PROJECT_ID_OR_NUMBER]/uptimeCheckConfigs/[UPTIME_CHECK_ID] [PROJECT_ID_OR_NUMBER] is the Workspace host project associated with the Uptime check.This field should be omitted when creating the Uptime check configuration; on create, the resource name is assigned by the server and included in the response.
func (o LookupUptimeCheckConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUptimeCheckConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// How often, in seconds, the Uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 60s.
func (o LookupUptimeCheckConfigResultOutput) Period() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUptimeCheckConfigResult) string { return v.Period }).(pulumi.StringOutput)
}

// The group resource associated with the configuration.
func (o LookupUptimeCheckConfigResultOutput) ResourceGroup() ResourceGroupResponseOutput {
	return o.ApplyT(func(v LookupUptimeCheckConfigResult) ResourceGroupResponse { return v.ResourceGroup }).(ResourceGroupResponseOutput)
}

// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions must be provided to include a minimum of 3 locations. Not specifying this field will result in Uptime checks running from all available regions.
func (o LookupUptimeCheckConfigResultOutput) SelectedRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUptimeCheckConfigResult) []string { return v.SelectedRegions }).(pulumi.StringArrayOutput)
}

// Contains information needed to make a TCP check.
func (o LookupUptimeCheckConfigResultOutput) TcpCheck() TcpCheckResponseOutput {
	return o.ApplyT(func(v LookupUptimeCheckConfigResult) TcpCheckResponse { return v.TcpCheck }).(TcpCheckResponseOutput)
}

// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Required.
func (o LookupUptimeCheckConfigResultOutput) Timeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUptimeCheckConfigResult) string { return v.Timeout }).(pulumi.StringOutput)
}

// User-supplied key/value data to be used for organizing and identifying the UptimeCheckConfig objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
func (o LookupUptimeCheckConfigResultOutput) UserLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupUptimeCheckConfigResult) map[string]string { return v.UserLabels }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUptimeCheckConfigResultOutput{})
}
