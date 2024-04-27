// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the database specified by the resource name.
func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabaseResult
	err := ctx.Invoke("google-native:biglake/v1:getDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseArgs struct {
	CatalogId  string  `pulumi:"catalogId"`
	DatabaseId string  `pulumi:"databaseId"`
	Location   string  `pulumi:"location"`
	Project    *string `pulumi:"project"`
}

type LookupDatabaseResult struct {
	// The creation time of the database.
	CreateTime string `pulumi:"createTime"`
	// The deletion time of the database. Only set after the database is deleted.
	DeleteTime string `pulumi:"deleteTime"`
	// The time when this database is considered expired. Only set after the database is deleted.
	ExpireTime string `pulumi:"expireTime"`
	// Options of a Hive database.
	HiveOptions HiveDatabaseOptionsResponse `pulumi:"hiveOptions"`
	// The resource name. Format: projects/{project_id_or_number}/locations/{location_id}/catalogs/{catalog_id}/databases/{database_id}
	Name string `pulumi:"name"`
	// The database type.
	Type string `pulumi:"type"`
	// The last modification time of the database.
	UpdateTime string `pulumi:"updateTime"`
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
	CatalogId  pulumi.StringInput    `pulumi:"catalogId"`
	DatabaseId pulumi.StringInput    `pulumi:"databaseId"`
	Location   pulumi.StringInput    `pulumi:"location"`
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

// The creation time of the database.
func (o LookupDatabaseResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The deletion time of the database. Only set after the database is deleted.
func (o LookupDatabaseResultOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.DeleteTime }).(pulumi.StringOutput)
}

// The time when this database is considered expired. Only set after the database is deleted.
func (o LookupDatabaseResultOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.ExpireTime }).(pulumi.StringOutput)
}

// Options of a Hive database.
func (o LookupDatabaseResultOutput) HiveOptions() HiveDatabaseOptionsResponseOutput {
	return o.ApplyT(func(v LookupDatabaseResult) HiveDatabaseOptionsResponse { return v.HiveOptions }).(HiveDatabaseOptionsResponseOutput)
}

// The resource name. Format: projects/{project_id_or_number}/locations/{location_id}/catalogs/{catalog_id}/databases/{database_id}
func (o LookupDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// The database type.
func (o LookupDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

// The last modification time of the database.
func (o LookupDatabaseResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseResultOutput{})
}