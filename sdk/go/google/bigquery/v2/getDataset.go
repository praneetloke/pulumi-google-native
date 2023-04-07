// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the dataset specified by datasetID.
func LookupDataset(ctx *pulumi.Context, args *LookupDatasetArgs, opts ...pulumi.InvokeOption) (*LookupDatasetResult, error) {
	var rv LookupDatasetResult
	err := ctx.Invoke("google-native:bigquery/v2:getDataset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatasetArgs struct {
	DatasetId string  `pulumi:"datasetId"`
	Project   *string `pulumi:"project"`
}

type LookupDatasetResult struct {
	// [Optional] An array of objects that define dataset access for one or more entities. You can set this property when inserting or updating a dataset in order to control who is allowed to access the data. If unspecified at dataset creation time, BigQuery adds default dataset access for the following entities: access.specialGroup: projectReaders; access.role: READER; access.specialGroup: projectWriters; access.role: WRITER; access.specialGroup: projectOwners; access.role: OWNER; access.userByEmail: [dataset creator email]; access.role: OWNER;
	Access []DatasetAccessItemResponse `pulumi:"access"`
	// The time when this dataset was created, in milliseconds since the epoch.
	CreationTime string `pulumi:"creationTime"`
	// [Required] A reference that identifies the dataset.
	DatasetReference DatasetReferenceResponse `pulumi:"datasetReference"`
	// The default collation of the dataset.
	DefaultCollation               string                          `pulumi:"defaultCollation"`
	DefaultEncryptionConfiguration EncryptionConfigurationResponse `pulumi:"defaultEncryptionConfiguration"`
	// [Optional] The default partition expiration for all partitioned tables in the dataset, in milliseconds. Once this property is set, all newly-created partitioned tables in the dataset will have an expirationMs property in the timePartitioning settings set to this value, and changing the value will only affect new tables, not existing ones. The storage in a partition will have an expiration time of its partition time plus this value. Setting this property overrides the use of defaultTableExpirationMs for partitioned tables: only one of defaultTableExpirationMs and defaultPartitionExpirationMs will be used for any new partitioned table. If you provide an explicit timePartitioning.expirationMs when creating or updating a partitioned table, that value takes precedence over the default partition expiration time indicated by this property.
	DefaultPartitionExpirationMs string `pulumi:"defaultPartitionExpirationMs"`
	// The default rounding mode of the dataset.
	DefaultRoundingMode string `pulumi:"defaultRoundingMode"`
	// [Optional] The default lifetime of all tables in the dataset, in milliseconds. The minimum value is 3600000 milliseconds (one hour). Once this property is set, all newly-created tables in the dataset will have an expirationTime property set to the creation time plus the value in this property, and changing the value will only affect new tables, not existing ones. When the expirationTime for a given table is reached, that table will be deleted automatically. If a table's expirationTime is modified or removed before the table expires, or if you provide an explicit expirationTime when creating a table, that value takes precedence over the default expiration time indicated by this property.
	DefaultTableExpirationMs string `pulumi:"defaultTableExpirationMs"`
	// [Optional] A user-friendly description of the dataset.
	Description string `pulumi:"description"`
	// A hash of the resource.
	Etag string `pulumi:"etag"`
	// [Optional] A descriptive name for the dataset.
	FriendlyName string `pulumi:"friendlyName"`
	// [Optional] Indicates if table names are case insensitive in the dataset.
	IsCaseInsensitive bool `pulumi:"isCaseInsensitive"`
	// The resource type.
	Kind string `pulumi:"kind"`
	// The labels associated with this dataset. You can use these to organize and group your datasets. You can set this property when inserting or updating a dataset. See Creating and Updating Dataset Labels for more information.
	Labels map[string]string `pulumi:"labels"`
	// The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.
	LastModifiedTime string `pulumi:"lastModifiedTime"`
	// The geographic location where the dataset should reside. The default value is US. See details at https://cloud.google.com/bigquery/docs/locations.
	Location string `pulumi:"location"`
	// [Optional] Number of hours for the max time travel for all tables in the dataset.
	MaxTimeTravelHours string `pulumi:"maxTimeTravelHours"`
	// Reserved for future use.
	SatisfiesPzs bool `pulumi:"satisfiesPzs"`
	// A URL that can be used to access the resource again. You can use this URL in Get or Update requests to the resource.
	SelfLink string `pulumi:"selfLink"`
	// [Optional] Storage billing model to be used for all tables in the dataset. Can be set to PHYSICAL. Default is LOGICAL.
	StorageBillingModel string `pulumi:"storageBillingModel"`
	// [Optional]The tags associated with this dataset. Tag keys are globally unique.
	Tags []DatasetTagsItemResponse `pulumi:"tags"`
}

func LookupDatasetOutput(ctx *pulumi.Context, args LookupDatasetOutputArgs, opts ...pulumi.InvokeOption) LookupDatasetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatasetResult, error) {
			args := v.(LookupDatasetArgs)
			r, err := LookupDataset(ctx, &args, opts...)
			var s LookupDatasetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatasetResultOutput)
}

type LookupDatasetOutputArgs struct {
	DatasetId pulumi.StringInput    `pulumi:"datasetId"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDatasetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatasetArgs)(nil)).Elem()
}

type LookupDatasetResultOutput struct{ *pulumi.OutputState }

func (LookupDatasetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatasetResult)(nil)).Elem()
}

func (o LookupDatasetResultOutput) ToLookupDatasetResultOutput() LookupDatasetResultOutput {
	return o
}

func (o LookupDatasetResultOutput) ToLookupDatasetResultOutputWithContext(ctx context.Context) LookupDatasetResultOutput {
	return o
}

// [Optional] An array of objects that define dataset access for one or more entities. You can set this property when inserting or updating a dataset in order to control who is allowed to access the data. If unspecified at dataset creation time, BigQuery adds default dataset access for the following entities: access.specialGroup: projectReaders; access.role: READER; access.specialGroup: projectWriters; access.role: WRITER; access.specialGroup: projectOwners; access.role: OWNER; access.userByEmail: [dataset creator email]; access.role: OWNER;
func (o LookupDatasetResultOutput) Access() DatasetAccessItemResponseArrayOutput {
	return o.ApplyT(func(v LookupDatasetResult) []DatasetAccessItemResponse { return v.Access }).(DatasetAccessItemResponseArrayOutput)
}

// The time when this dataset was created, in milliseconds since the epoch.
func (o LookupDatasetResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// [Required] A reference that identifies the dataset.
func (o LookupDatasetResultOutput) DatasetReference() DatasetReferenceResponseOutput {
	return o.ApplyT(func(v LookupDatasetResult) DatasetReferenceResponse { return v.DatasetReference }).(DatasetReferenceResponseOutput)
}

// The default collation of the dataset.
func (o LookupDatasetResultOutput) DefaultCollation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.DefaultCollation }).(pulumi.StringOutput)
}

func (o LookupDatasetResultOutput) DefaultEncryptionConfiguration() EncryptionConfigurationResponseOutput {
	return o.ApplyT(func(v LookupDatasetResult) EncryptionConfigurationResponse { return v.DefaultEncryptionConfiguration }).(EncryptionConfigurationResponseOutput)
}

// [Optional] The default partition expiration for all partitioned tables in the dataset, in milliseconds. Once this property is set, all newly-created partitioned tables in the dataset will have an expirationMs property in the timePartitioning settings set to this value, and changing the value will only affect new tables, not existing ones. The storage in a partition will have an expiration time of its partition time plus this value. Setting this property overrides the use of defaultTableExpirationMs for partitioned tables: only one of defaultTableExpirationMs and defaultPartitionExpirationMs will be used for any new partitioned table. If you provide an explicit timePartitioning.expirationMs when creating or updating a partitioned table, that value takes precedence over the default partition expiration time indicated by this property.
func (o LookupDatasetResultOutput) DefaultPartitionExpirationMs() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.DefaultPartitionExpirationMs }).(pulumi.StringOutput)
}

// The default rounding mode of the dataset.
func (o LookupDatasetResultOutput) DefaultRoundingMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.DefaultRoundingMode }).(pulumi.StringOutput)
}

// [Optional] The default lifetime of all tables in the dataset, in milliseconds. The minimum value is 3600000 milliseconds (one hour). Once this property is set, all newly-created tables in the dataset will have an expirationTime property set to the creation time plus the value in this property, and changing the value will only affect new tables, not existing ones. When the expirationTime for a given table is reached, that table will be deleted automatically. If a table's expirationTime is modified or removed before the table expires, or if you provide an explicit expirationTime when creating a table, that value takes precedence over the default expiration time indicated by this property.
func (o LookupDatasetResultOutput) DefaultTableExpirationMs() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.DefaultTableExpirationMs }).(pulumi.StringOutput)
}

// [Optional] A user-friendly description of the dataset.
func (o LookupDatasetResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.Description }).(pulumi.StringOutput)
}

// A hash of the resource.
func (o LookupDatasetResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.Etag }).(pulumi.StringOutput)
}

// [Optional] A descriptive name for the dataset.
func (o LookupDatasetResultOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.FriendlyName }).(pulumi.StringOutput)
}

// [Optional] Indicates if table names are case insensitive in the dataset.
func (o LookupDatasetResultOutput) IsCaseInsensitive() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDatasetResult) bool { return v.IsCaseInsensitive }).(pulumi.BoolOutput)
}

// The resource type.
func (o LookupDatasetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The labels associated with this dataset. You can use these to organize and group your datasets. You can set this property when inserting or updating a dataset. See Creating and Updating Dataset Labels for more information.
func (o LookupDatasetResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatasetResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.
func (o LookupDatasetResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// The geographic location where the dataset should reside. The default value is US. See details at https://cloud.google.com/bigquery/docs/locations.
func (o LookupDatasetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.Location }).(pulumi.StringOutput)
}

// [Optional] Number of hours for the max time travel for all tables in the dataset.
func (o LookupDatasetResultOutput) MaxTimeTravelHours() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.MaxTimeTravelHours }).(pulumi.StringOutput)
}

// Reserved for future use.
func (o LookupDatasetResultOutput) SatisfiesPzs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDatasetResult) bool { return v.SatisfiesPzs }).(pulumi.BoolOutput)
}

// A URL that can be used to access the resource again. You can use this URL in Get or Update requests to the resource.
func (o LookupDatasetResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// [Optional] Storage billing model to be used for all tables in the dataset. Can be set to PHYSICAL. Default is LOGICAL.
func (o LookupDatasetResultOutput) StorageBillingModel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.StorageBillingModel }).(pulumi.StringOutput)
}

// [Optional]The tags associated with this dataset. Tag keys are globally unique.
func (o LookupDatasetResultOutput) Tags() DatasetTagsItemResponseArrayOutput {
	return o.ApplyT(func(v LookupDatasetResult) []DatasetTagsItemResponse { return v.Tags }).(DatasetTagsItemResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatasetResultOutput{})
}
