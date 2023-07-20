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

// Creates a DataScan resource.
// Auto-naming is currently not supported for this resource.
type DataScan struct {
	pulumi.CustomResourceState

	// The time when the scan was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The data source for DataScan.
	Data GoogleCloudDataplexV1DataSourceResponseOutput `pulumi:"data"`
	// The result of the data profile scan.
	DataProfileResult GoogleCloudDataplexV1DataProfileResultResponseOutput `pulumi:"dataProfileResult"`
	// DataProfileScan related setting.
	DataProfileSpec GoogleCloudDataplexV1DataProfileSpecResponseOutput `pulumi:"dataProfileSpec"`
	// The result of the data quality scan.
	DataQualityResult GoogleCloudDataplexV1DataQualityResultResponseOutput `pulumi:"dataQualityResult"`
	// DataQualityScan related setting.
	DataQualitySpec GoogleCloudDataplexV1DataQualitySpecResponseOutput `pulumi:"dataQualitySpec"`
	// Required. DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter. Must be between 1-63 characters. Must be unique within the customer project / location.
	DataScanId pulumi.StringOutput `pulumi:"dataScanId"`
	// Optional. Description of the scan. Must be between 1-1024 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. User friendly display name. Must be between 1-256 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Optional. DataScan execution settings.If not specified, the fields in it will use their default values.
	ExecutionSpec GoogleCloudDataplexV1DataScanExecutionSpecResponseOutput `pulumi:"executionSpec"`
	// Status of the data scan execution.
	ExecutionStatus GoogleCloudDataplexV1DataScanExecutionStatusResponseOutput `pulumi:"executionStatus"`
	// Optional. User-defined labels for the scan.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// The relative resource name of the scan, of the form: projects/{project}/locations/{location_id}/dataScans/{datascan_id}, where project refers to a project_id or project_number and location_id refers to a GCP region.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Current state of the DataScan.
	State pulumi.StringOutput `pulumi:"state"`
	// The type of DataScan.
	Type pulumi.StringOutput `pulumi:"type"`
	// System generated globally unique ID for the scan. This ID will be different if the scan is deleted and re-created with the same name.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The time when the scan was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewDataScan registers a new resource with the given unique name, arguments, and options.
func NewDataScan(ctx *pulumi.Context,
	name string, args *DataScanArgs, opts ...pulumi.ResourceOption) (*DataScan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Data == nil {
		return nil, errors.New("invalid value for required argument 'Data'")
	}
	if args.DataScanId == nil {
		return nil, errors.New("invalid value for required argument 'DataScanId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"dataScanId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataScan
	err := ctx.RegisterResource("google-native:dataplex/v1:DataScan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataScan gets an existing DataScan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataScan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataScanState, opts ...pulumi.ResourceOption) (*DataScan, error) {
	var resource DataScan
	err := ctx.ReadResource("google-native:dataplex/v1:DataScan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataScan resources.
type dataScanState struct {
}

type DataScanState struct {
}

func (DataScanState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataScanState)(nil)).Elem()
}

type dataScanArgs struct {
	// The data source for DataScan.
	Data GoogleCloudDataplexV1DataSource `pulumi:"data"`
	// DataProfileScan related setting.
	DataProfileSpec *GoogleCloudDataplexV1DataProfileSpec `pulumi:"dataProfileSpec"`
	// DataQualityScan related setting.
	DataQualitySpec *GoogleCloudDataplexV1DataQualitySpec `pulumi:"dataQualitySpec"`
	// Required. DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter. Must be between 1-63 characters. Must be unique within the customer project / location.
	DataScanId string `pulumi:"dataScanId"`
	// Optional. Description of the scan. Must be between 1-1024 characters.
	Description *string `pulumi:"description"`
	// Optional. User friendly display name. Must be between 1-256 characters.
	DisplayName *string `pulumi:"displayName"`
	// Optional. DataScan execution settings.If not specified, the fields in it will use their default values.
	ExecutionSpec *GoogleCloudDataplexV1DataScanExecutionSpec `pulumi:"executionSpec"`
	// Optional. User-defined labels for the scan.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	Project  *string           `pulumi:"project"`
}

// The set of arguments for constructing a DataScan resource.
type DataScanArgs struct {
	// The data source for DataScan.
	Data GoogleCloudDataplexV1DataSourceInput
	// DataProfileScan related setting.
	DataProfileSpec GoogleCloudDataplexV1DataProfileSpecPtrInput
	// DataQualityScan related setting.
	DataQualitySpec GoogleCloudDataplexV1DataQualitySpecPtrInput
	// Required. DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter. Must be between 1-63 characters. Must be unique within the customer project / location.
	DataScanId pulumi.StringInput
	// Optional. Description of the scan. Must be between 1-1024 characters.
	Description pulumi.StringPtrInput
	// Optional. User friendly display name. Must be between 1-256 characters.
	DisplayName pulumi.StringPtrInput
	// Optional. DataScan execution settings.If not specified, the fields in it will use their default values.
	ExecutionSpec GoogleCloudDataplexV1DataScanExecutionSpecPtrInput
	// Optional. User-defined labels for the scan.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
}

func (DataScanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataScanArgs)(nil)).Elem()
}

type DataScanInput interface {
	pulumi.Input

	ToDataScanOutput() DataScanOutput
	ToDataScanOutputWithContext(ctx context.Context) DataScanOutput
}

func (*DataScan) ElementType() reflect.Type {
	return reflect.TypeOf((**DataScan)(nil)).Elem()
}

func (i *DataScan) ToDataScanOutput() DataScanOutput {
	return i.ToDataScanOutputWithContext(context.Background())
}

func (i *DataScan) ToDataScanOutputWithContext(ctx context.Context) DataScanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataScanOutput)
}

type DataScanOutput struct{ *pulumi.OutputState }

func (DataScanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataScan)(nil)).Elem()
}

func (o DataScanOutput) ToDataScanOutput() DataScanOutput {
	return o
}

func (o DataScanOutput) ToDataScanOutputWithContext(ctx context.Context) DataScanOutput {
	return o
}

// The time when the scan was created.
func (o DataScanOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataScan) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The data source for DataScan.
func (o DataScanOutput) Data() GoogleCloudDataplexV1DataSourceResponseOutput {
	return o.ApplyT(func(v *DataScan) GoogleCloudDataplexV1DataSourceResponseOutput { return v.Data }).(GoogleCloudDataplexV1DataSourceResponseOutput)
}

// The result of the data profile scan.
func (o DataScanOutput) DataProfileResult() GoogleCloudDataplexV1DataProfileResultResponseOutput {
	return o.ApplyT(func(v *DataScan) GoogleCloudDataplexV1DataProfileResultResponseOutput { return v.DataProfileResult }).(GoogleCloudDataplexV1DataProfileResultResponseOutput)
}

// DataProfileScan related setting.
func (o DataScanOutput) DataProfileSpec() GoogleCloudDataplexV1DataProfileSpecResponseOutput {
	return o.ApplyT(func(v *DataScan) GoogleCloudDataplexV1DataProfileSpecResponseOutput { return v.DataProfileSpec }).(GoogleCloudDataplexV1DataProfileSpecResponseOutput)
}

// The result of the data quality scan.
func (o DataScanOutput) DataQualityResult() GoogleCloudDataplexV1DataQualityResultResponseOutput {
	return o.ApplyT(func(v *DataScan) GoogleCloudDataplexV1DataQualityResultResponseOutput { return v.DataQualityResult }).(GoogleCloudDataplexV1DataQualityResultResponseOutput)
}

// DataQualityScan related setting.
func (o DataScanOutput) DataQualitySpec() GoogleCloudDataplexV1DataQualitySpecResponseOutput {
	return o.ApplyT(func(v *DataScan) GoogleCloudDataplexV1DataQualitySpecResponseOutput { return v.DataQualitySpec }).(GoogleCloudDataplexV1DataQualitySpecResponseOutput)
}

// Required. DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter. Must be between 1-63 characters. Must be unique within the customer project / location.
func (o DataScanOutput) DataScanId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataScan) pulumi.StringOutput { return v.DataScanId }).(pulumi.StringOutput)
}

// Optional. Description of the scan. Must be between 1-1024 characters.
func (o DataScanOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *DataScan) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Optional. User friendly display name. Must be between 1-256 characters.
func (o DataScanOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DataScan) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Optional. DataScan execution settings.If not specified, the fields in it will use their default values.
func (o DataScanOutput) ExecutionSpec() GoogleCloudDataplexV1DataScanExecutionSpecResponseOutput {
	return o.ApplyT(func(v *DataScan) GoogleCloudDataplexV1DataScanExecutionSpecResponseOutput { return v.ExecutionSpec }).(GoogleCloudDataplexV1DataScanExecutionSpecResponseOutput)
}

// Status of the data scan execution.
func (o DataScanOutput) ExecutionStatus() GoogleCloudDataplexV1DataScanExecutionStatusResponseOutput {
	return o.ApplyT(func(v *DataScan) GoogleCloudDataplexV1DataScanExecutionStatusResponseOutput { return v.ExecutionStatus }).(GoogleCloudDataplexV1DataScanExecutionStatusResponseOutput)
}

// Optional. User-defined labels for the scan.
func (o DataScanOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataScan) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o DataScanOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataScan) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The relative resource name of the scan, of the form: projects/{project}/locations/{location_id}/dataScans/{datascan_id}, where project refers to a project_id or project_number and location_id refers to a GCP region.
func (o DataScanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataScan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DataScanOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataScan) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Current state of the DataScan.
func (o DataScanOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *DataScan) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The type of DataScan.
func (o DataScanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataScan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// System generated globally unique ID for the scan. This ID will be different if the scan is deleted and re-created with the same name.
func (o DataScanOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *DataScan) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The time when the scan was last updated.
func (o DataScanOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataScan) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataScanInput)(nil)).Elem(), &DataScan{})
	pulumi.RegisterOutputType(DataScanOutput{})
}
