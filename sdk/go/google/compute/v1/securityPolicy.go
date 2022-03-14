// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new policy in the specified project using the data included in the request.
type SecurityPolicy struct {
	pulumi.CustomResourceState

	AdaptiveProtectionConfig SecurityPolicyAdaptiveProtectionConfigResponseOutput `pulumi:"adaptiveProtectionConfig"`
	AdvancedOptionsConfig    SecurityPolicyAdvancedOptionsConfigResponseOutput    `pulumi:"advancedOptionsConfig"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the security policy.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// [Output only] Type of the resource. Always compute#securityPolicyfor security policies
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name                   pulumi.StringOutput                                `pulumi:"name"`
	RecaptchaOptionsConfig SecurityPolicyRecaptchaOptionsConfigResponseOutput `pulumi:"recaptchaOptionsConfig"`
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
	Rules SecurityPolicyRuleResponseArrayOutput `pulumi:"rules"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The type indicates the intended use of the security policy. - CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. - CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache. - CLOUD_ARMOR_INTERNAL_SERVICE: Cloud Armor internal service policies can be configured to filter HTTP requests targeting services managed by Traffic Director in a service mesh. They filter requests before the request is served from the application. This field can be set only at resource creation time.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSecurityPolicy registers a new resource with the given unique name, arguments, and options.
func NewSecurityPolicy(ctx *pulumi.Context,
	name string, args *SecurityPolicyArgs, opts ...pulumi.ResourceOption) (*SecurityPolicy, error) {
	if args == nil {
		args = &SecurityPolicyArgs{}
	}

	var resource SecurityPolicy
	err := ctx.RegisterResource("google-native:compute/v1:SecurityPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityPolicy gets an existing SecurityPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityPolicyState, opts ...pulumi.ResourceOption) (*SecurityPolicy, error) {
	var resource SecurityPolicy
	err := ctx.ReadResource("google-native:compute/v1:SecurityPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityPolicy resources.
type securityPolicyState struct {
}

type SecurityPolicyState struct {
}

func (SecurityPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPolicyState)(nil)).Elem()
}

type securityPolicyArgs struct {
	AdaptiveProtectionConfig *SecurityPolicyAdaptiveProtectionConfig `pulumi:"adaptiveProtectionConfig"`
	AdvancedOptionsConfig    *SecurityPolicyAdvancedOptionsConfig    `pulumi:"advancedOptionsConfig"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name                   *string                               `pulumi:"name"`
	Project                *string                               `pulumi:"project"`
	RecaptchaOptionsConfig *SecurityPolicyRecaptchaOptionsConfig `pulumi:"recaptchaOptionsConfig"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
	Rules []SecurityPolicyRule `pulumi:"rules"`
	// The type indicates the intended use of the security policy. - CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. - CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache. - CLOUD_ARMOR_INTERNAL_SERVICE: Cloud Armor internal service policies can be configured to filter HTTP requests targeting services managed by Traffic Director in a service mesh. They filter requests before the request is served from the application. This field can be set only at resource creation time.
	Type *SecurityPolicyType `pulumi:"type"`
}

// The set of arguments for constructing a SecurityPolicy resource.
type SecurityPolicyArgs struct {
	AdaptiveProtectionConfig SecurityPolicyAdaptiveProtectionConfigPtrInput
	AdvancedOptionsConfig    SecurityPolicyAdvancedOptionsConfigPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name                   pulumi.StringPtrInput
	Project                pulumi.StringPtrInput
	RecaptchaOptionsConfig SecurityPolicyRecaptchaOptionsConfigPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
	Rules SecurityPolicyRuleArrayInput
	// The type indicates the intended use of the security policy. - CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. - CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache. - CLOUD_ARMOR_INTERNAL_SERVICE: Cloud Armor internal service policies can be configured to filter HTTP requests targeting services managed by Traffic Director in a service mesh. They filter requests before the request is served from the application. This field can be set only at resource creation time.
	Type SecurityPolicyTypePtrInput
}

func (SecurityPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPolicyArgs)(nil)).Elem()
}

type SecurityPolicyInput interface {
	pulumi.Input

	ToSecurityPolicyOutput() SecurityPolicyOutput
	ToSecurityPolicyOutputWithContext(ctx context.Context) SecurityPolicyOutput
}

func (*SecurityPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPolicy)(nil)).Elem()
}

func (i *SecurityPolicy) ToSecurityPolicyOutput() SecurityPolicyOutput {
	return i.ToSecurityPolicyOutputWithContext(context.Background())
}

func (i *SecurityPolicy) ToSecurityPolicyOutputWithContext(ctx context.Context) SecurityPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPolicyOutput)
}

type SecurityPolicyOutput struct{ *pulumi.OutputState }

func (SecurityPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPolicy)(nil)).Elem()
}

func (o SecurityPolicyOutput) ToSecurityPolicyOutput() SecurityPolicyOutput {
	return o
}

func (o SecurityPolicyOutput) ToSecurityPolicyOutputWithContext(ctx context.Context) SecurityPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityPolicyInput)(nil)).Elem(), &SecurityPolicy{})
	pulumi.RegisterOutputType(SecurityPolicyOutput{})
}
