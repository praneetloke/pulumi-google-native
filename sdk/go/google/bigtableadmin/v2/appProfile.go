// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an app profile within an instance.
type AppProfile struct {
	pulumi.CustomResourceState

	// Required. The ID to be used when referring to the new app profile within its instance, e.g., just `myprofile` rather than `projects/myproject/instances/myinstance/appProfiles/myprofile`.
	AppProfileId pulumi.StringOutput `pulumi:"appProfileId"`
	// Long form description of the use case for this AppProfile.
	Description pulumi.StringOutput `pulumi:"description"`
	// Strongly validated etag for optimistic concurrency control. Preserve the value returned from `GetAppProfile` when calling `UpdateAppProfile` to fail the request if there has been a modification in the mean time. The `update_mask` of the request need not include `etag` for this protection to apply. See [Wikipedia](https://en.wikipedia.org/wiki/HTTP_ETag) and [RFC 7232](https://tools.ietf.org/html/rfc7232#section-2.3) for more details.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// If true, ignore safety checks when creating the app profile.
	IgnoreWarnings pulumi.BoolPtrOutput `pulumi:"ignoreWarnings"`
	InstanceId     pulumi.StringOutput  `pulumi:"instanceId"`
	// Use a multi-cluster routing policy.
	MultiClusterRoutingUseAny MultiClusterRoutingUseAnyResponseOutput `pulumi:"multiClusterRoutingUseAny"`
	// The unique name of the app profile. Values are of the form `projects/{project}/instances/{instance}/appProfiles/_a-zA-Z0-9*`.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Use a single-cluster routing policy.
	SingleClusterRouting SingleClusterRoutingResponseOutput `pulumi:"singleClusterRouting"`
}

// NewAppProfile registers a new resource with the given unique name, arguments, and options.
func NewAppProfile(ctx *pulumi.Context,
	name string, args *AppProfileArgs, opts ...pulumi.ResourceOption) (*AppProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppProfileId == nil {
		return nil, errors.New("invalid value for required argument 'AppProfileId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"appProfileId",
		"instanceId",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AppProfile
	err := ctx.RegisterResource("google-native:bigtableadmin/v2:AppProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppProfile gets an existing AppProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppProfileState, opts ...pulumi.ResourceOption) (*AppProfile, error) {
	var resource AppProfile
	err := ctx.ReadResource("google-native:bigtableadmin/v2:AppProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppProfile resources.
type appProfileState struct {
}

type AppProfileState struct {
}

func (AppProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*appProfileState)(nil)).Elem()
}

type appProfileArgs struct {
	// Required. The ID to be used when referring to the new app profile within its instance, e.g., just `myprofile` rather than `projects/myproject/instances/myinstance/appProfiles/myprofile`.
	AppProfileId string `pulumi:"appProfileId"`
	// Long form description of the use case for this AppProfile.
	Description *string `pulumi:"description"`
	// Strongly validated etag for optimistic concurrency control. Preserve the value returned from `GetAppProfile` when calling `UpdateAppProfile` to fail the request if there has been a modification in the mean time. The `update_mask` of the request need not include `etag` for this protection to apply. See [Wikipedia](https://en.wikipedia.org/wiki/HTTP_ETag) and [RFC 7232](https://tools.ietf.org/html/rfc7232#section-2.3) for more details.
	Etag *string `pulumi:"etag"`
	// If true, ignore safety checks when creating the app profile.
	IgnoreWarnings *bool  `pulumi:"ignoreWarnings"`
	InstanceId     string `pulumi:"instanceId"`
	// Use a multi-cluster routing policy.
	MultiClusterRoutingUseAny *MultiClusterRoutingUseAny `pulumi:"multiClusterRoutingUseAny"`
	// The unique name of the app profile. Values are of the form `projects/{project}/instances/{instance}/appProfiles/_a-zA-Z0-9*`.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Use a single-cluster routing policy.
	SingleClusterRouting *SingleClusterRouting `pulumi:"singleClusterRouting"`
}

// The set of arguments for constructing a AppProfile resource.
type AppProfileArgs struct {
	// Required. The ID to be used when referring to the new app profile within its instance, e.g., just `myprofile` rather than `projects/myproject/instances/myinstance/appProfiles/myprofile`.
	AppProfileId pulumi.StringInput
	// Long form description of the use case for this AppProfile.
	Description pulumi.StringPtrInput
	// Strongly validated etag for optimistic concurrency control. Preserve the value returned from `GetAppProfile` when calling `UpdateAppProfile` to fail the request if there has been a modification in the mean time. The `update_mask` of the request need not include `etag` for this protection to apply. See [Wikipedia](https://en.wikipedia.org/wiki/HTTP_ETag) and [RFC 7232](https://tools.ietf.org/html/rfc7232#section-2.3) for more details.
	Etag pulumi.StringPtrInput
	// If true, ignore safety checks when creating the app profile.
	IgnoreWarnings pulumi.BoolPtrInput
	InstanceId     pulumi.StringInput
	// Use a multi-cluster routing policy.
	MultiClusterRoutingUseAny MultiClusterRoutingUseAnyPtrInput
	// The unique name of the app profile. Values are of the form `projects/{project}/instances/{instance}/appProfiles/_a-zA-Z0-9*`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Use a single-cluster routing policy.
	SingleClusterRouting SingleClusterRoutingPtrInput
}

func (AppProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appProfileArgs)(nil)).Elem()
}

type AppProfileInput interface {
	pulumi.Input

	ToAppProfileOutput() AppProfileOutput
	ToAppProfileOutputWithContext(ctx context.Context) AppProfileOutput
}

func (*AppProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**AppProfile)(nil)).Elem()
}

func (i *AppProfile) ToAppProfileOutput() AppProfileOutput {
	return i.ToAppProfileOutputWithContext(context.Background())
}

func (i *AppProfile) ToAppProfileOutputWithContext(ctx context.Context) AppProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppProfileOutput)
}

type AppProfileOutput struct{ *pulumi.OutputState }

func (AppProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppProfile)(nil)).Elem()
}

func (o AppProfileOutput) ToAppProfileOutput() AppProfileOutput {
	return o
}

func (o AppProfileOutput) ToAppProfileOutputWithContext(ctx context.Context) AppProfileOutput {
	return o
}

// Required. The ID to be used when referring to the new app profile within its instance, e.g., just `myprofile` rather than `projects/myproject/instances/myinstance/appProfiles/myprofile`.
func (o AppProfileOutput) AppProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppProfile) pulumi.StringOutput { return v.AppProfileId }).(pulumi.StringOutput)
}

// Long form description of the use case for this AppProfile.
func (o AppProfileOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *AppProfile) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Strongly validated etag for optimistic concurrency control. Preserve the value returned from `GetAppProfile` when calling `UpdateAppProfile` to fail the request if there has been a modification in the mean time. The `update_mask` of the request need not include `etag` for this protection to apply. See [Wikipedia](https://en.wikipedia.org/wiki/HTTP_ETag) and [RFC 7232](https://tools.ietf.org/html/rfc7232#section-2.3) for more details.
func (o AppProfileOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AppProfile) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// If true, ignore safety checks when creating the app profile.
func (o AppProfileOutput) IgnoreWarnings() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppProfile) pulumi.BoolPtrOutput { return v.IgnoreWarnings }).(pulumi.BoolPtrOutput)
}

func (o AppProfileOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppProfile) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Use a multi-cluster routing policy.
func (o AppProfileOutput) MultiClusterRoutingUseAny() MultiClusterRoutingUseAnyResponseOutput {
	return o.ApplyT(func(v *AppProfile) MultiClusterRoutingUseAnyResponseOutput { return v.MultiClusterRoutingUseAny }).(MultiClusterRoutingUseAnyResponseOutput)
}

// The unique name of the app profile. Values are of the form `projects/{project}/instances/{instance}/appProfiles/_a-zA-Z0-9*`.
func (o AppProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AppProfileOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AppProfile) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Use a single-cluster routing policy.
func (o AppProfileOutput) SingleClusterRouting() SingleClusterRoutingResponseOutput {
	return o.ApplyT(func(v *AppProfile) SingleClusterRoutingResponseOutput { return v.SingleClusterRouting }).(SingleClusterRoutingResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppProfileInput)(nil)).Elem(), &AppProfile{})
	pulumi.RegisterOutputType(AppProfileOutput{})
}
