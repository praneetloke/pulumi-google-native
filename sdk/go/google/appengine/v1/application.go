// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an App Engine application for a Google Cloud Platform project. Required fields: id - The ID of the target Cloud Platform project. location - The region (https://cloud.google.com/appengine/docs/locations) where you want the App Engine application located.For more information about App Engine applications, see Managing Projects, Applications, and Billing (https://cloud.google.com/appengine/docs/standard/python/console/).
// Auto-naming is currently not supported for this resource.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type Application struct {
	pulumi.CustomResourceState

	// Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
	AuthDomain pulumi.StringOutput `pulumi:"authDomain"`
	// Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.
	CodeBucket pulumi.StringOutput `pulumi:"codeBucket"`
	// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
	DatabaseType pulumi.StringOutput `pulumi:"databaseType"`
	// Google Cloud Storage bucket that can be used by this application to store content.
	DefaultBucket pulumi.StringOutput `pulumi:"defaultBucket"`
	// Cookie expiration policy for this application.
	DefaultCookieExpiration pulumi.StringOutput `pulumi:"defaultCookieExpiration"`
	// Hostname used to reach this application, as resolved by App Engine.
	DefaultHostname pulumi.StringOutput `pulumi:"defaultHostname"`
	// HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
	DispatchRules UrlDispatchRuleResponseArrayOutput `pulumi:"dispatchRules"`
	// The feature specific settings to be used in the application.
	FeatureSettings FeatureSettingsResponseOutput `pulumi:"featureSettings"`
	// The Google Container Registry domain used for storing managed build docker images for this application.
	GcrDomain pulumi.StringOutput              `pulumi:"gcrDomain"`
	Iap       IdentityAwareProxyResponseOutput `pulumi:"iap"`
	Location  pulumi.StringOutput              `pulumi:"location"`
	// Full path to the Application resource in the API. Example: apps/myapp.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The service account associated with the application. This is the app-level default identity. If no identity provided during create version, Admin API will fallback to this one.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// Serving status of this application.
	ServingStatus pulumi.StringOutput `pulumi:"servingStatus"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		args = &ApplicationArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Application
	err := ctx.RegisterResource("google-native:appengine/v1:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("google-native:appengine/v1:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
}

type ApplicationState struct {
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
	AuthDomain *string `pulumi:"authDomain"`
	// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
	DatabaseType *ApplicationDatabaseType `pulumi:"databaseType"`
	// Cookie expiration policy for this application.
	DefaultCookieExpiration *string `pulumi:"defaultCookieExpiration"`
	// HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
	DispatchRules []UrlDispatchRule `pulumi:"dispatchRules"`
	// The feature specific settings to be used in the application.
	FeatureSettings *FeatureSettings    `pulumi:"featureSettings"`
	Iap             *IdentityAwareProxy `pulumi:"iap"`
	// Identifier of the Application resource. This identifier is equivalent to the project ID of the Google Cloud Platform project where you want to deploy your application. Example: myapp.
	Id *string `pulumi:"id"`
	// Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
	Location *string `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// The service account associated with the application. This is the app-level default identity. If no identity provided during create version, Admin API will fallback to this one.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Serving status of this application.
	ServingStatus *ApplicationServingStatus `pulumi:"servingStatus"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
	AuthDomain pulumi.StringPtrInput
	// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
	DatabaseType ApplicationDatabaseTypePtrInput
	// Cookie expiration policy for this application.
	DefaultCookieExpiration pulumi.StringPtrInput
	// HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
	DispatchRules UrlDispatchRuleArrayInput
	// The feature specific settings to be used in the application.
	FeatureSettings FeatureSettingsPtrInput
	Iap             IdentityAwareProxyPtrInput
	// Identifier of the Application resource. This identifier is equivalent to the project ID of the Google Cloud Platform project where you want to deploy your application. Example: myapp.
	Id pulumi.StringPtrInput
	// Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// The service account associated with the application. This is the app-level default identity. If no identity provided during create version, Admin API will fallback to this one.
	ServiceAccount pulumi.StringPtrInput
	// Serving status of this application.
	ServingStatus ApplicationServingStatusPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

// Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
func (o ApplicationOutput) AuthDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.AuthDomain }).(pulumi.StringOutput)
}

// Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.
func (o ApplicationOutput) CodeBucket() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.CodeBucket }).(pulumi.StringOutput)
}

// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
func (o ApplicationOutput) DatabaseType() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.DatabaseType }).(pulumi.StringOutput)
}

// Google Cloud Storage bucket that can be used by this application to store content.
func (o ApplicationOutput) DefaultBucket() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.DefaultBucket }).(pulumi.StringOutput)
}

// Cookie expiration policy for this application.
func (o ApplicationOutput) DefaultCookieExpiration() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.DefaultCookieExpiration }).(pulumi.StringOutput)
}

// Hostname used to reach this application, as resolved by App Engine.
func (o ApplicationOutput) DefaultHostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.DefaultHostname }).(pulumi.StringOutput)
}

// HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
func (o ApplicationOutput) DispatchRules() UrlDispatchRuleResponseArrayOutput {
	return o.ApplyT(func(v *Application) UrlDispatchRuleResponseArrayOutput { return v.DispatchRules }).(UrlDispatchRuleResponseArrayOutput)
}

// The feature specific settings to be used in the application.
func (o ApplicationOutput) FeatureSettings() FeatureSettingsResponseOutput {
	return o.ApplyT(func(v *Application) FeatureSettingsResponseOutput { return v.FeatureSettings }).(FeatureSettingsResponseOutput)
}

// The Google Container Registry domain used for storing managed build docker images for this application.
func (o ApplicationOutput) GcrDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.GcrDomain }).(pulumi.StringOutput)
}

func (o ApplicationOutput) Iap() IdentityAwareProxyResponseOutput {
	return o.ApplyT(func(v *Application) IdentityAwareProxyResponseOutput { return v.Iap }).(IdentityAwareProxyResponseOutput)
}

func (o ApplicationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Full path to the Application resource in the API. Example: apps/myapp.
func (o ApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The service account associated with the application. This is the app-level default identity. If no identity provided during create version, Admin API will fallback to this one.
func (o ApplicationOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ServiceAccount }).(pulumi.StringOutput)
}

// Serving status of this application.
func (o ApplicationOutput) ServingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ServingStatus }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInput)(nil)).Elem(), &Application{})
	pulumi.RegisterOutputType(ApplicationOutput{})
}
