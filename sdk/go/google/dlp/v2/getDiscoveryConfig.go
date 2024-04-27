// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a discovery configuration.
func LookupDiscoveryConfig(ctx *pulumi.Context, args *LookupDiscoveryConfigArgs, opts ...pulumi.InvokeOption) (*LookupDiscoveryConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDiscoveryConfigResult
	err := ctx.Invoke("google-native:dlp/v2:getDiscoveryConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiscoveryConfigArgs struct {
	DiscoveryConfigId string  `pulumi:"discoveryConfigId"`
	Location          string  `pulumi:"location"`
	Project           *string `pulumi:"project"`
}

type LookupDiscoveryConfigResult struct {
	// Actions to execute at the completion of scanning.
	Actions []GooglePrivacyDlpV2DataProfileActionResponse `pulumi:"actions"`
	// The creation timestamp of a DiscoveryConfig.
	CreateTime string `pulumi:"createTime"`
	// Display name (max 100 chars)
	DisplayName string `pulumi:"displayName"`
	// A stream of errors encountered when the config was activated. Repeated errors may result in the config automatically being paused. Output only field. Will return the last 100 errors. Whenever the config is modified this list will be cleared.
	Errors []GooglePrivacyDlpV2ErrorResponse `pulumi:"errors"`
	// Detection logic for profile generation. Not all template features are used by Discovery. FindingLimits, include_quote and exclude_info_types have no impact on Discovery. Multiple templates may be provided if there is data in multiple regions. At most one template must be specified per-region (including "global"). Each region is scanned using the applicable template. If no region-specific template is specified, but a "global" template is specified, it will be copied to that region and used instead. If no global or region-specific template is provided for a region with data, that region's data will not be scanned. For more information, see https://cloud.google.com/dlp/docs/data-profiles#data-residency.
	InspectTemplates []string `pulumi:"inspectTemplates"`
	// The timestamp of the last time this config was executed.
	LastRunTime string `pulumi:"lastRunTime"`
	// Unique resource name for the DiscoveryConfig, assigned by the service when the DiscoveryConfig is created, for example `projects/dlp-test-project/locations/global/discoveryConfigs/53234423`.
	Name string `pulumi:"name"`
	// Only set when the parent is an org.
	OrgConfig GooglePrivacyDlpV2OrgConfigResponse `pulumi:"orgConfig"`
	// A status for this configuration.
	Status string `pulumi:"status"`
	// Target to match against for determining what to scan and how frequently.
	Targets []GooglePrivacyDlpV2DiscoveryTargetResponse `pulumi:"targets"`
	// The last update timestamp of a DiscoveryConfig.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupDiscoveryConfigOutput(ctx *pulumi.Context, args LookupDiscoveryConfigOutputArgs, opts ...pulumi.InvokeOption) LookupDiscoveryConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiscoveryConfigResult, error) {
			args := v.(LookupDiscoveryConfigArgs)
			r, err := LookupDiscoveryConfig(ctx, &args, opts...)
			var s LookupDiscoveryConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiscoveryConfigResultOutput)
}

type LookupDiscoveryConfigOutputArgs struct {
	DiscoveryConfigId pulumi.StringInput    `pulumi:"discoveryConfigId"`
	Location          pulumi.StringInput    `pulumi:"location"`
	Project           pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDiscoveryConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiscoveryConfigArgs)(nil)).Elem()
}

type LookupDiscoveryConfigResultOutput struct{ *pulumi.OutputState }

func (LookupDiscoveryConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiscoveryConfigResult)(nil)).Elem()
}

func (o LookupDiscoveryConfigResultOutput) ToLookupDiscoveryConfigResultOutput() LookupDiscoveryConfigResultOutput {
	return o
}

func (o LookupDiscoveryConfigResultOutput) ToLookupDiscoveryConfigResultOutputWithContext(ctx context.Context) LookupDiscoveryConfigResultOutput {
	return o
}

// Actions to execute at the completion of scanning.
func (o LookupDiscoveryConfigResultOutput) Actions() GooglePrivacyDlpV2DataProfileActionResponseArrayOutput {
	return o.ApplyT(func(v LookupDiscoveryConfigResult) []GooglePrivacyDlpV2DataProfileActionResponse { return v.Actions }).(GooglePrivacyDlpV2DataProfileActionResponseArrayOutput)
}

// The creation timestamp of a DiscoveryConfig.
func (o LookupDiscoveryConfigResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiscoveryConfigResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Display name (max 100 chars)
func (o LookupDiscoveryConfigResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiscoveryConfigResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// A stream of errors encountered when the config was activated. Repeated errors may result in the config automatically being paused. Output only field. Will return the last 100 errors. Whenever the config is modified this list will be cleared.
func (o LookupDiscoveryConfigResultOutput) Errors() GooglePrivacyDlpV2ErrorResponseArrayOutput {
	return o.ApplyT(func(v LookupDiscoveryConfigResult) []GooglePrivacyDlpV2ErrorResponse { return v.Errors }).(GooglePrivacyDlpV2ErrorResponseArrayOutput)
}

// Detection logic for profile generation. Not all template features are used by Discovery. FindingLimits, include_quote and exclude_info_types have no impact on Discovery. Multiple templates may be provided if there is data in multiple regions. At most one template must be specified per-region (including "global"). Each region is scanned using the applicable template. If no region-specific template is specified, but a "global" template is specified, it will be copied to that region and used instead. If no global or region-specific template is provided for a region with data, that region's data will not be scanned. For more information, see https://cloud.google.com/dlp/docs/data-profiles#data-residency.
func (o LookupDiscoveryConfigResultOutput) InspectTemplates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDiscoveryConfigResult) []string { return v.InspectTemplates }).(pulumi.StringArrayOutput)
}

// The timestamp of the last time this config was executed.
func (o LookupDiscoveryConfigResultOutput) LastRunTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiscoveryConfigResult) string { return v.LastRunTime }).(pulumi.StringOutput)
}

// Unique resource name for the DiscoveryConfig, assigned by the service when the DiscoveryConfig is created, for example `projects/dlp-test-project/locations/global/discoveryConfigs/53234423`.
func (o LookupDiscoveryConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiscoveryConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// Only set when the parent is an org.
func (o LookupDiscoveryConfigResultOutput) OrgConfig() GooglePrivacyDlpV2OrgConfigResponseOutput {
	return o.ApplyT(func(v LookupDiscoveryConfigResult) GooglePrivacyDlpV2OrgConfigResponse { return v.OrgConfig }).(GooglePrivacyDlpV2OrgConfigResponseOutput)
}

// A status for this configuration.
func (o LookupDiscoveryConfigResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiscoveryConfigResult) string { return v.Status }).(pulumi.StringOutput)
}

// Target to match against for determining what to scan and how frequently.
func (o LookupDiscoveryConfigResultOutput) Targets() GooglePrivacyDlpV2DiscoveryTargetResponseArrayOutput {
	return o.ApplyT(func(v LookupDiscoveryConfigResult) []GooglePrivacyDlpV2DiscoveryTargetResponse { return v.Targets }).(GooglePrivacyDlpV2DiscoveryTargetResponseArrayOutput)
}

// The last update timestamp of a DiscoveryConfig.
func (o LookupDiscoveryConfigResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiscoveryConfigResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiscoveryConfigResultOutput{})
}