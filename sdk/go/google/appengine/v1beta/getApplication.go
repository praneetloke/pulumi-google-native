// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an application.
func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("google-native:appengine/v1beta:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	ApplicationId string  `pulumi:"applicationId"`
	Location      string  `pulumi:"location"`
	Project       *string `pulumi:"project"`
}

type LookupApplicationResult struct {
	// Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
	AuthDomain string `pulumi:"authDomain"`
	// Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.
	CodeBucket string `pulumi:"codeBucket"`
	// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
	DatabaseType string `pulumi:"databaseType"`
	// Google Cloud Storage bucket that can be used by this application to store content.
	DefaultBucket string `pulumi:"defaultBucket"`
	// Cookie expiration policy for this application.
	DefaultCookieExpiration string `pulumi:"defaultCookieExpiration"`
	// Hostname used to reach this application, as resolved by App Engine.
	DefaultHostname string `pulumi:"defaultHostname"`
	// HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
	DispatchRules []UrlDispatchRuleResponse `pulumi:"dispatchRules"`
	// The feature specific settings to be used in the application.
	FeatureSettings FeatureSettingsResponse `pulumi:"featureSettings"`
	// The Google Container Registry domain used for storing managed build docker images for this application.
	GcrDomain string                     `pulumi:"gcrDomain"`
	Iap       IdentityAwareProxyResponse `pulumi:"iap"`
	// Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
	Location string `pulumi:"location"`
	// Full path to the Application resource in the API. Example: apps/myapp.
	Name string `pulumi:"name"`
	// The service account associated with the application. This is the app-level default identity. If no identity provided during create version, Admin API will fallback to this one.
	ServiceAccount string `pulumi:"serviceAccount"`
	// Serving status of this application.
	ServingStatus string `pulumi:"servingStatus"`
}

func LookupApplicationOutput(ctx *pulumi.Context, args LookupApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationResult, error) {
			args := v.(LookupApplicationArgs)
			r, err := LookupApplication(ctx, &args, opts...)
			var s LookupApplicationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationResultOutput)
}

type LookupApplicationOutputArgs struct {
	ApplicationId pulumi.StringInput    `pulumi:"applicationId"`
	Location      pulumi.StringInput    `pulumi:"location"`
	Project       pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}

type LookupApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResult)(nil)).Elem()
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutput() LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutputWithContext(ctx context.Context) LookupApplicationResultOutput {
	return o
}

// Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
func (o LookupApplicationResultOutput) AuthDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.AuthDomain }).(pulumi.StringOutput)
}

// Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.
func (o LookupApplicationResultOutput) CodeBucket() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.CodeBucket }).(pulumi.StringOutput)
}

// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
func (o LookupApplicationResultOutput) DatabaseType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.DatabaseType }).(pulumi.StringOutput)
}

// Google Cloud Storage bucket that can be used by this application to store content.
func (o LookupApplicationResultOutput) DefaultBucket() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.DefaultBucket }).(pulumi.StringOutput)
}

// Cookie expiration policy for this application.
func (o LookupApplicationResultOutput) DefaultCookieExpiration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.DefaultCookieExpiration }).(pulumi.StringOutput)
}

// Hostname used to reach this application, as resolved by App Engine.
func (o LookupApplicationResultOutput) DefaultHostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.DefaultHostname }).(pulumi.StringOutput)
}

// HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
func (o LookupApplicationResultOutput) DispatchRules() UrlDispatchRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []UrlDispatchRuleResponse { return v.DispatchRules }).(UrlDispatchRuleResponseArrayOutput)
}

// The feature specific settings to be used in the application.
func (o LookupApplicationResultOutput) FeatureSettings() FeatureSettingsResponseOutput {
	return o.ApplyT(func(v LookupApplicationResult) FeatureSettingsResponse { return v.FeatureSettings }).(FeatureSettingsResponseOutput)
}

// The Google Container Registry domain used for storing managed build docker images for this application.
func (o LookupApplicationResultOutput) GcrDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.GcrDomain }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Iap() IdentityAwareProxyResponseOutput {
	return o.ApplyT(func(v LookupApplicationResult) IdentityAwareProxyResponse { return v.Iap }).(IdentityAwareProxyResponseOutput)
}

// Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
func (o LookupApplicationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Location }).(pulumi.StringOutput)
}

// Full path to the Application resource in the API. Example: apps/myapp.
func (o LookupApplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The service account associated with the application. This is the app-level default identity. If no identity provided during create version, Admin API will fallback to this one.
func (o LookupApplicationResultOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.ServiceAccount }).(pulumi.StringOutput)
}

// Serving status of this application.
func (o LookupApplicationResultOutput) ServingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.ServingStatus }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}
