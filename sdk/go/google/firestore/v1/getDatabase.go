// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a database.
func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	var rv LookupDatabaseResult
	err := ctx.Invoke("google-native:firestore/v1:getDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseArgs struct {
	DatabaseId string  `pulumi:"databaseId"`
	Project    *string `pulumi:"project"`
}

type LookupDatabaseResult struct {
	// The App Engine integration mode to use for this database.
	AppEngineIntegrationMode string `pulumi:"appEngineIntegrationMode"`
	// The concurrency control mode to use for this database.
	ConcurrencyMode string `pulumi:"concurrencyMode"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag string `pulumi:"etag"`
	// The key_prefix for this database. This key_prefix is used, in combination with the project id ("~") to construct the application id that is returned from the Cloud Datastore APIs in Google App Engine first generation runtimes. This value may be empty in which case the appid to use for URL-encoded keys is the project_id (eg: foo instead of v~foo).
	KeyPrefix string `pulumi:"keyPrefix"`
	// The location of the database. Available databases are listed at https://cloud.google.com/firestore/docs/locations.
	Location string `pulumi:"location"`
	// The resource name of the Database. Format: `projects/{project}/databases/{database}`
	Name string `pulumi:"name"`
	// The type of the database. See https://cloud.google.com/datastore/docs/firestore-or-datastore for information about how to choose.
	Type string `pulumi:"type"`
}

func LookupDatabaseOutput(ctx *pulumi.Context, args LookupDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseResult, error) {
			args := v.(LookupDatabaseArgs)
			r, err := LookupDatabase(ctx, &args, opts...)
			var s LookupDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseResultOutput)
}

type LookupDatabaseOutputArgs struct {
	DatabaseId pulumi.StringInput    `pulumi:"databaseId"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseArgs)(nil)).Elem()
}

type LookupDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseResult)(nil)).Elem()
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutput() LookupDatabaseResultOutput {
	return o
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutputWithContext(ctx context.Context) LookupDatabaseResultOutput {
	return o
}

// The App Engine integration mode to use for this database.
func (o LookupDatabaseResultOutput) AppEngineIntegrationMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.AppEngineIntegrationMode }).(pulumi.StringOutput)
}

// The concurrency control mode to use for this database.
func (o LookupDatabaseResultOutput) ConcurrencyMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.ConcurrencyMode }).(pulumi.StringOutput)
}

// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
func (o LookupDatabaseResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The key_prefix for this database. This key_prefix is used, in combination with the project id ("~") to construct the application id that is returned from the Cloud Datastore APIs in Google App Engine first generation runtimes. This value may be empty in which case the appid to use for URL-encoded keys is the project_id (eg: foo instead of v~foo).
func (o LookupDatabaseResultOutput) KeyPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.KeyPrefix }).(pulumi.StringOutput)
}

// The location of the database. Available databases are listed at https://cloud.google.com/firestore/docs/locations.
func (o LookupDatabaseResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the Database. Format: `projects/{project}/databases/{database}`
func (o LookupDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the database. See https://cloud.google.com/datastore/docs/firestore-or-datastore for information about how to choose.
func (o LookupDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseResultOutput{})
}