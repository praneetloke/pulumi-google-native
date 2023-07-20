// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List all of the ordered rules present in a single specified policy.
func LookupRegionSecurityPolicy(ctx *pulumi.Context, args *LookupRegionSecurityPolicyArgs, opts ...pulumi.InvokeOption) (*LookupRegionSecurityPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRegionSecurityPolicyResult
	err := ctx.Invoke("google-native:compute/beta:getRegionSecurityPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegionSecurityPolicyArgs struct {
	Project        *string `pulumi:"project"`
	Region         string  `pulumi:"region"`
	SecurityPolicy string  `pulumi:"securityPolicy"`
}

type LookupRegionSecurityPolicyResult struct {
	AdaptiveProtectionConfig SecurityPolicyAdaptiveProtectionConfigResponse `pulumi:"adaptiveProtectionConfig"`
	AdvancedOptionsConfig    SecurityPolicyAdvancedOptionsConfigResponse    `pulumi:"advancedOptionsConfig"`
	// A list of associations that belong to this policy.
	Associations []SecurityPolicyAssociationResponse `pulumi:"associations"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp    string                                     `pulumi:"creationTimestamp"`
	DdosProtectionConfig SecurityPolicyDdosProtectionConfigResponse `pulumi:"ddosProtectionConfig"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// User-provided name of the Organization security plicy. The name should be unique in the organization in which the security policy is created. This should only be used when SecurityPolicyType is FIREWALL. The name must be 1-63 characters long, and comply with https://www.ietf.org/rfc/rfc1035.txt. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName string `pulumi:"displayName"`
	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the security policy.
	Fingerprint string `pulumi:"fingerprint"`
	// [Output only] Type of the resource. Always compute#securityPolicyfor security policies
	Kind string `pulumi:"kind"`
	// A fingerprint for the labels being applied to this security policy, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels. To see the latest fingerprint, make get() request to the security policy.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// The parent of the security policy.
	Parent                 string                                       `pulumi:"parent"`
	RecaptchaOptionsConfig SecurityPolicyRecaptchaOptionsConfigResponse `pulumi:"recaptchaOptionsConfig"`
	// URL of the region where the regional security policy resides. This field is not applicable to global security policies.
	Region string `pulumi:"region"`
	// Total count of all security policy rule tuples. A security policy can not exceed a set number of tuples.
	RuleTupleCount int `pulumi:"ruleTupleCount"`
	// A list of rules that belong to this policy. There must always be a default rule which is a rule with priority 2147483647 and match all condition (for the match condition this means match "*" for srcIpRanges and for the networkMatch condition every field must be either match "*" or not set). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
	Rules []SecurityPolicyRuleResponse `pulumi:"rules"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// The type indicates the intended use of the security policy. - CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. - CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache. - CLOUD_ARMOR_INTERNAL_SERVICE: Cloud Armor internal service policies can be configured to filter HTTP requests targeting services managed by Traffic Director in a service mesh. They filter requests before the request is served from the application. - CLOUD_ARMOR_NETWORK: Cloud Armor network policies can be configured to filter packets targeting network load balancing resources such as backend services, target pools, target instances, and instances with external IPs. They filter requests before the request is served from the application. This field can be set only at resource creation time.
	Type string `pulumi:"type"`
	// Definitions of user-defined fields for CLOUD_ARMOR_NETWORK policies. A user-defined field consists of up to 4 bytes extracted from a fixed offset in the packet, relative to the IPv4, IPv6, TCP, or UDP header, with an optional mask to select certain bits. Rules may then specify matching values for these fields. Example: userDefinedFields: - name: "ipv4_fragment_offset" base: IPV4 offset: 6 size: 2 mask: "0x1fff"
	UserDefinedFields []SecurityPolicyUserDefinedFieldResponse `pulumi:"userDefinedFields"`
}

func LookupRegionSecurityPolicyOutput(ctx *pulumi.Context, args LookupRegionSecurityPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupRegionSecurityPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegionSecurityPolicyResult, error) {
			args := v.(LookupRegionSecurityPolicyArgs)
			r, err := LookupRegionSecurityPolicy(ctx, &args, opts...)
			var s LookupRegionSecurityPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegionSecurityPolicyResultOutput)
}

type LookupRegionSecurityPolicyOutputArgs struct {
	Project        pulumi.StringPtrInput `pulumi:"project"`
	Region         pulumi.StringInput    `pulumi:"region"`
	SecurityPolicy pulumi.StringInput    `pulumi:"securityPolicy"`
}

func (LookupRegionSecurityPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionSecurityPolicyArgs)(nil)).Elem()
}

type LookupRegionSecurityPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupRegionSecurityPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionSecurityPolicyResult)(nil)).Elem()
}

func (o LookupRegionSecurityPolicyResultOutput) ToLookupRegionSecurityPolicyResultOutput() LookupRegionSecurityPolicyResultOutput {
	return o
}

func (o LookupRegionSecurityPolicyResultOutput) ToLookupRegionSecurityPolicyResultOutputWithContext(ctx context.Context) LookupRegionSecurityPolicyResultOutput {
	return o
}

func (o LookupRegionSecurityPolicyResultOutput) AdaptiveProtectionConfig() SecurityPolicyAdaptiveProtectionConfigResponseOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) SecurityPolicyAdaptiveProtectionConfigResponse {
		return v.AdaptiveProtectionConfig
	}).(SecurityPolicyAdaptiveProtectionConfigResponseOutput)
}

func (o LookupRegionSecurityPolicyResultOutput) AdvancedOptionsConfig() SecurityPolicyAdvancedOptionsConfigResponseOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) SecurityPolicyAdvancedOptionsConfigResponse {
		return v.AdvancedOptionsConfig
	}).(SecurityPolicyAdvancedOptionsConfigResponseOutput)
}

// A list of associations that belong to this policy.
func (o LookupRegionSecurityPolicyResultOutput) Associations() SecurityPolicyAssociationResponseArrayOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) []SecurityPolicyAssociationResponse { return v.Associations }).(SecurityPolicyAssociationResponseArrayOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupRegionSecurityPolicyResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

func (o LookupRegionSecurityPolicyResultOutput) DdosProtectionConfig() SecurityPolicyDdosProtectionConfigResponseOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) SecurityPolicyDdosProtectionConfigResponse {
		return v.DdosProtectionConfig
	}).(SecurityPolicyDdosProtectionConfigResponseOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupRegionSecurityPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// User-provided name of the Organization security plicy. The name should be unique in the organization in which the security policy is created. This should only be used when SecurityPolicyType is FIREWALL. The name must be 1-63 characters long, and comply with https://www.ietf.org/rfc/rfc1035.txt. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupRegionSecurityPolicyResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the security policy.
func (o LookupRegionSecurityPolicyResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// [Output only] Type of the resource. Always compute#securityPolicyfor security policies
func (o LookupRegionSecurityPolicyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) string { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this security policy, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels. To see the latest fingerprint, make get() request to the security policy.
func (o LookupRegionSecurityPolicyResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
func (o LookupRegionSecurityPolicyResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupRegionSecurityPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The parent of the security policy.
func (o LookupRegionSecurityPolicyResultOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) string { return v.Parent }).(pulumi.StringOutput)
}

func (o LookupRegionSecurityPolicyResultOutput) RecaptchaOptionsConfig() SecurityPolicyRecaptchaOptionsConfigResponseOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) SecurityPolicyRecaptchaOptionsConfigResponse {
		return v.RecaptchaOptionsConfig
	}).(SecurityPolicyRecaptchaOptionsConfigResponseOutput)
}

// URL of the region where the regional security policy resides. This field is not applicable to global security policies.
func (o LookupRegionSecurityPolicyResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) string { return v.Region }).(pulumi.StringOutput)
}

// Total count of all security policy rule tuples. A security policy can not exceed a set number of tuples.
func (o LookupRegionSecurityPolicyResultOutput) RuleTupleCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) int { return v.RuleTupleCount }).(pulumi.IntOutput)
}

// A list of rules that belong to this policy. There must always be a default rule which is a rule with priority 2147483647 and match all condition (for the match condition this means match "*" for srcIpRanges and for the networkMatch condition every field must be either match "*" or not set). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
func (o LookupRegionSecurityPolicyResultOutput) Rules() SecurityPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) []SecurityPolicyRuleResponse { return v.Rules }).(SecurityPolicyRuleResponseArrayOutput)
}

// Server-defined URL for the resource.
func (o LookupRegionSecurityPolicyResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o LookupRegionSecurityPolicyResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// The type indicates the intended use of the security policy. - CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. - CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache. - CLOUD_ARMOR_INTERNAL_SERVICE: Cloud Armor internal service policies can be configured to filter HTTP requests targeting services managed by Traffic Director in a service mesh. They filter requests before the request is served from the application. - CLOUD_ARMOR_NETWORK: Cloud Armor network policies can be configured to filter packets targeting network load balancing resources such as backend services, target pools, target instances, and instances with external IPs. They filter requests before the request is served from the application. This field can be set only at resource creation time.
func (o LookupRegionSecurityPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

// Definitions of user-defined fields for CLOUD_ARMOR_NETWORK policies. A user-defined field consists of up to 4 bytes extracted from a fixed offset in the packet, relative to the IPv4, IPv6, TCP, or UDP header, with an optional mask to select certain bits. Rules may then specify matching values for these fields. Example: userDefinedFields: - name: "ipv4_fragment_offset" base: IPV4 offset: 6 size: 2 mask: "0x1fff"
func (o LookupRegionSecurityPolicyResultOutput) UserDefinedFields() SecurityPolicyUserDefinedFieldResponseArrayOutput {
	return o.ApplyT(func(v LookupRegionSecurityPolicyResult) []SecurityPolicyUserDefinedFieldResponse {
		return v.UserDefinedFields
	}).(SecurityPolicyUserDefinedFieldResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegionSecurityPolicyResultOutput{})
}
