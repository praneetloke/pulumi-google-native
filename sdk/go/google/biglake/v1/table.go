// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new table.
// Auto-naming is currently not supported for this resource.
type Table struct {
	pulumi.CustomResourceState

	CatalogId pulumi.StringOutput `pulumi:"catalogId"`
	// The creation time of the table.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	DatabaseId pulumi.StringOutput `pulumi:"databaseId"`
	// The deletion time of the table. Only set after the table is deleted.
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// The checksum of a table object computed by the server based on the value of other fields. It may be sent on update requests to ensure the client has an up-to-date value before proceeding. It is only checked for update table operations.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The time when this table is considered expired. Only set after the table is deleted.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// Options of a Hive table.
	HiveOptions HiveTableOptionsResponseOutput `pulumi:"hiveOptions"`
	Location    pulumi.StringOutput            `pulumi:"location"`
	// The resource name. Format: projects/{project_id_or_number}/locations/{location_id}/catalogs/{catalog_id}/databases/{database_id}/tables/{table_id}
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Required. The ID to use for the table, which will become the final component of the table's resource name.
	TableId pulumi.StringOutput `pulumi:"tableId"`
	// The table type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The last modification time of the table.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CatalogId == nil {
		return nil, errors.New("invalid value for required argument 'CatalogId'")
	}
	if args.DatabaseId == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseId'")
	}
	if args.TableId == nil {
		return nil, errors.New("invalid value for required argument 'TableId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"catalogId",
		"databaseId",
		"location",
		"project",
		"tableId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Table
	err := ctx.RegisterResource("google-native:biglake/v1:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("google-native:biglake/v1:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
}

type TableState struct {
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	CatalogId  string `pulumi:"catalogId"`
	DatabaseId string `pulumi:"databaseId"`
	// The checksum of a table object computed by the server based on the value of other fields. It may be sent on update requests to ensure the client has an up-to-date value before proceeding. It is only checked for update table operations.
	Etag *string `pulumi:"etag"`
	// Options of a Hive table.
	HiveOptions *HiveTableOptions `pulumi:"hiveOptions"`
	Location    *string           `pulumi:"location"`
	Project     *string           `pulumi:"project"`
	// Required. The ID to use for the table, which will become the final component of the table's resource name.
	TableId string `pulumi:"tableId"`
	// The table type.
	Type *TableType `pulumi:"type"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	CatalogId  pulumi.StringInput
	DatabaseId pulumi.StringInput
	// The checksum of a table object computed by the server based on the value of other fields. It may be sent on update requests to ensure the client has an up-to-date value before proceeding. It is only checked for update table operations.
	Etag pulumi.StringPtrInput
	// Options of a Hive table.
	HiveOptions HiveTableOptionsPtrInput
	Location    pulumi.StringPtrInput
	Project     pulumi.StringPtrInput
	// Required. The ID to use for the table, which will become the final component of the table's resource name.
	TableId pulumi.StringInput
	// The table type.
	Type TableTypePtrInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

type TableInput interface {
	pulumi.Input

	ToTableOutput() TableOutput
	ToTableOutputWithContext(ctx context.Context) TableOutput
}

func (*Table) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (i *Table) ToTableOutput() TableOutput {
	return i.ToTableOutputWithContext(context.Background())
}

func (i *Table) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableOutput)
}

type TableOutput struct{ *pulumi.OutputState }

func (TableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (o TableOutput) ToTableOutput() TableOutput {
	return o
}

func (o TableOutput) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return o
}

func (o TableOutput) CatalogId() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.CatalogId }).(pulumi.StringOutput)
}

// The creation time of the table.
func (o TableOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o TableOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.DatabaseId }).(pulumi.StringOutput)
}

// The deletion time of the table. Only set after the table is deleted.
func (o TableOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.DeleteTime }).(pulumi.StringOutput)
}

// The checksum of a table object computed by the server based on the value of other fields. It may be sent on update requests to ensure the client has an up-to-date value before proceeding. It is only checked for update table operations.
func (o TableOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The time when this table is considered expired. Only set after the table is deleted.
func (o TableOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.ExpireTime }).(pulumi.StringOutput)
}

// Options of a Hive table.
func (o TableOutput) HiveOptions() HiveTableOptionsResponseOutput {
	return o.ApplyT(func(v *Table) HiveTableOptionsResponseOutput { return v.HiveOptions }).(HiveTableOptionsResponseOutput)
}

func (o TableOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name. Format: projects/{project_id_or_number}/locations/{location_id}/catalogs/{catalog_id}/databases/{database_id}/tables/{table_id}
func (o TableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TableOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Required. The ID to use for the table, which will become the final component of the table's resource name.
func (o TableOutput) TableId() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.TableId }).(pulumi.StringOutput)
}

// The table type.
func (o TableOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The last modification time of the table.
func (o TableOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TableInput)(nil)).Elem(), &Table{})
	pulumi.RegisterOutputType(TableOutput{})
}