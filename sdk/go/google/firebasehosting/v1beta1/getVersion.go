// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the specified version that has been created for the specified site. This can include versions that were created for the default `live` channel or for any active preview channels for the specified site.
func LookupVersion(ctx *pulumi.Context, args *LookupVersionArgs, opts ...pulumi.InvokeOption) (*LookupVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVersionResult
	err := ctx.Invoke("google-native:firebasehosting/v1beta1:getVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVersionArgs struct {
	Project   *string `pulumi:"project"`
	SiteId    string  `pulumi:"siteId"`
	VersionId string  `pulumi:"versionId"`
}

type LookupVersionResult struct {
	// The configuration for the behavior of the site. This configuration exists in the [`firebase.json`](https://firebase.google.com/docs/cli/#the_firebasejson_file) file.
	Config ServingConfigResponse `pulumi:"config"`
	// The time at which the version was created.
	CreateTime string `pulumi:"createTime"`
	// Identifies the user who created the version.
	CreateUser ActingUserResponse `pulumi:"createUser"`
	// The time at which the version was `DELETED`.
	DeleteTime string `pulumi:"deleteTime"`
	// Identifies the user who `DELETED` the version.
	DeleteUser ActingUserResponse `pulumi:"deleteUser"`
	// The total number of files associated with the version. This value is calculated after a version is `FINALIZED`.
	FileCount string `pulumi:"fileCount"`
	// The time at which the version was `FINALIZED`.
	FinalizeTime string `pulumi:"finalizeTime"`
	// Identifies the user who `FINALIZED` the version.
	FinalizeUser ActingUserResponse `pulumi:"finalizeUser"`
	// The labels used for extra metadata and/or filtering.
	Labels map[string]string `pulumi:"labels"`
	// The fully-qualified resource name for the version, in the format: sites/ SITE_ID/versions/VERSION_ID This name is provided in the response body when you call [`CreateVersion`](sites.versions/create).
	Name string `pulumi:"name"`
	// The deploy status of the version. For a successful deploy, call [`CreateVersion`](sites.versions/create) to make a new version (`CREATED` status), [upload all desired files](sites.versions/populateFiles) to the version, then [update](sites.versions/patch) the version to the `FINALIZED` status. Note that if you leave the version in the `CREATED` state for more than 12 hours, the system will automatically mark the version as `ABANDONED`. You can also change the status of a version to `DELETED` by calling [`DeleteVersion`](sites.versions/delete).
	Status string `pulumi:"status"`
	// The total stored bytesize of the version. This value is calculated after a version is `FINALIZED`.
	VersionBytes string `pulumi:"versionBytes"`
}

func LookupVersionOutput(ctx *pulumi.Context, args LookupVersionOutputArgs, opts ...pulumi.InvokeOption) LookupVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVersionResult, error) {
			args := v.(LookupVersionArgs)
			r, err := LookupVersion(ctx, &args, opts...)
			var s LookupVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVersionResultOutput)
}

type LookupVersionOutputArgs struct {
	Project   pulumi.StringPtrInput `pulumi:"project"`
	SiteId    pulumi.StringInput    `pulumi:"siteId"`
	VersionId pulumi.StringInput    `pulumi:"versionId"`
}

func (LookupVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVersionArgs)(nil)).Elem()
}

type LookupVersionResultOutput struct{ *pulumi.OutputState }

func (LookupVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVersionResult)(nil)).Elem()
}

func (o LookupVersionResultOutput) ToLookupVersionResultOutput() LookupVersionResultOutput {
	return o
}

func (o LookupVersionResultOutput) ToLookupVersionResultOutputWithContext(ctx context.Context) LookupVersionResultOutput {
	return o
}

// The configuration for the behavior of the site. This configuration exists in the [`firebase.json`](https://firebase.google.com/docs/cli/#the_firebasejson_file) file.
func (o LookupVersionResultOutput) Config() ServingConfigResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) ServingConfigResponse { return v.Config }).(ServingConfigResponseOutput)
}

// The time at which the version was created.
func (o LookupVersionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Identifies the user who created the version.
func (o LookupVersionResultOutput) CreateUser() ActingUserResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) ActingUserResponse { return v.CreateUser }).(ActingUserResponseOutput)
}

// The time at which the version was `DELETED`.
func (o LookupVersionResultOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.DeleteTime }).(pulumi.StringOutput)
}

// Identifies the user who `DELETED` the version.
func (o LookupVersionResultOutput) DeleteUser() ActingUserResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) ActingUserResponse { return v.DeleteUser }).(ActingUserResponseOutput)
}

// The total number of files associated with the version. This value is calculated after a version is `FINALIZED`.
func (o LookupVersionResultOutput) FileCount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.FileCount }).(pulumi.StringOutput)
}

// The time at which the version was `FINALIZED`.
func (o LookupVersionResultOutput) FinalizeTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.FinalizeTime }).(pulumi.StringOutput)
}

// Identifies the user who `FINALIZED` the version.
func (o LookupVersionResultOutput) FinalizeUser() ActingUserResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) ActingUserResponse { return v.FinalizeUser }).(ActingUserResponseOutput)
}

// The labels used for extra metadata and/or filtering.
func (o LookupVersionResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVersionResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The fully-qualified resource name for the version, in the format: sites/ SITE_ID/versions/VERSION_ID This name is provided in the response body when you call [`CreateVersion`](sites.versions/create).
func (o LookupVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The deploy status of the version. For a successful deploy, call [`CreateVersion`](sites.versions/create) to make a new version (`CREATED` status), [upload all desired files](sites.versions/populateFiles) to the version, then [update](sites.versions/patch) the version to the `FINALIZED` status. Note that if you leave the version in the `CREATED` state for more than 12 hours, the system will automatically mark the version as `ABANDONED`. You can also change the status of a version to `DELETED` by calling [`DeleteVersion`](sites.versions/delete).
func (o LookupVersionResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.Status }).(pulumi.StringOutput)
}

// The total stored bytesize of the version. This value is calculated after a version is `FINALIZED`.
func (o LookupVersionResultOutput) VersionBytes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.VersionBytes }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVersionResultOutput{})
}
