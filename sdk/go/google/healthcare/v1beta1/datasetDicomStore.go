// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new DICOM store within the parent dataset.
type DatasetDicomStore struct {
	pulumi.CustomResourceState

	// User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Resource name of the DICOM store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/dicomStores/{dicom_store_id}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Notification destination for new DICOM instances. Supplied by the client.
	NotificationConfig NotificationConfigResponseOutput `pulumi:"notificationConfig"`
	// A list of streaming configs used to configure the destination of streaming exports for every DICOM instance insertion in this DICOM store. After a new config is added to `stream_configs`, DICOM instance insertions are streamed to the new destination. When a config is removed from `stream_configs`, the server stops streaming to that destination. Each config must contain a unique destination.
	StreamConfigs GoogleCloudHealthcareV1beta1DicomStreamConfigResponseArrayOutput `pulumi:"streamConfigs"`
}

// NewDatasetDicomStore registers a new resource with the given unique name, arguments, and options.
func NewDatasetDicomStore(ctx *pulumi.Context,
	name string, args *DatasetDicomStoreArgs, opts ...pulumi.ResourceOption) (*DatasetDicomStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource DatasetDicomStore
	err := ctx.RegisterResource("google-native:healthcare/v1beta1:DatasetDicomStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetDicomStore gets an existing DatasetDicomStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetDicomStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetDicomStoreState, opts ...pulumi.ResourceOption) (*DatasetDicomStore, error) {
	var resource DatasetDicomStore
	err := ctx.ReadResource("google-native:healthcare/v1beta1:DatasetDicomStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetDicomStore resources.
type datasetDicomStoreState struct {
	// User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels map[string]string `pulumi:"labels"`
	// Resource name of the DICOM store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/dicomStores/{dicom_store_id}`.
	Name *string `pulumi:"name"`
	// Notification destination for new DICOM instances. Supplied by the client.
	NotificationConfig *NotificationConfigResponse `pulumi:"notificationConfig"`
	// A list of streaming configs used to configure the destination of streaming exports for every DICOM instance insertion in this DICOM store. After a new config is added to `stream_configs`, DICOM instance insertions are streamed to the new destination. When a config is removed from `stream_configs`, the server stops streaming to that destination. Each config must contain a unique destination.
	StreamConfigs []GoogleCloudHealthcareV1beta1DicomStreamConfigResponse `pulumi:"streamConfigs"`
}

type DatasetDicomStoreState struct {
	// User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels pulumi.StringMapInput
	// Resource name of the DICOM store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/dicomStores/{dicom_store_id}`.
	Name pulumi.StringPtrInput
	// Notification destination for new DICOM instances. Supplied by the client.
	NotificationConfig NotificationConfigResponsePtrInput
	// A list of streaming configs used to configure the destination of streaming exports for every DICOM instance insertion in this DICOM store. After a new config is added to `stream_configs`, DICOM instance insertions are streamed to the new destination. When a config is removed from `stream_configs`, the server stops streaming to that destination. Each config must contain a unique destination.
	StreamConfigs GoogleCloudHealthcareV1beta1DicomStreamConfigResponseArrayInput
}

func (DatasetDicomStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetDicomStoreState)(nil)).Elem()
}

type datasetDicomStoreArgs struct {
	DatasetId    string  `pulumi:"datasetId"`
	DicomStoreId *string `pulumi:"dicomStoreId"`
	// User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels   map[string]string `pulumi:"labels"`
	Location string            `pulumi:"location"`
	// Resource name of the DICOM store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/dicomStores/{dicom_store_id}`.
	Name *string `pulumi:"name"`
	// Notification destination for new DICOM instances. Supplied by the client.
	NotificationConfig *NotificationConfig `pulumi:"notificationConfig"`
	Project            string              `pulumi:"project"`
	// A list of streaming configs used to configure the destination of streaming exports for every DICOM instance insertion in this DICOM store. After a new config is added to `stream_configs`, DICOM instance insertions are streamed to the new destination. When a config is removed from `stream_configs`, the server stops streaming to that destination. Each config must contain a unique destination.
	StreamConfigs []GoogleCloudHealthcareV1beta1DicomStreamConfig `pulumi:"streamConfigs"`
}

// The set of arguments for constructing a DatasetDicomStore resource.
type DatasetDicomStoreArgs struct {
	DatasetId    pulumi.StringInput
	DicomStoreId pulumi.StringPtrInput
	// User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels   pulumi.StringMapInput
	Location pulumi.StringInput
	// Resource name of the DICOM store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/dicomStores/{dicom_store_id}`.
	Name pulumi.StringPtrInput
	// Notification destination for new DICOM instances. Supplied by the client.
	NotificationConfig NotificationConfigPtrInput
	Project            pulumi.StringInput
	// A list of streaming configs used to configure the destination of streaming exports for every DICOM instance insertion in this DICOM store. After a new config is added to `stream_configs`, DICOM instance insertions are streamed to the new destination. When a config is removed from `stream_configs`, the server stops streaming to that destination. Each config must contain a unique destination.
	StreamConfigs GoogleCloudHealthcareV1beta1DicomStreamConfigArrayInput
}

func (DatasetDicomStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetDicomStoreArgs)(nil)).Elem()
}

type DatasetDicomStoreInput interface {
	pulumi.Input

	ToDatasetDicomStoreOutput() DatasetDicomStoreOutput
	ToDatasetDicomStoreOutputWithContext(ctx context.Context) DatasetDicomStoreOutput
}

func (*DatasetDicomStore) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetDicomStore)(nil))
}

func (i *DatasetDicomStore) ToDatasetDicomStoreOutput() DatasetDicomStoreOutput {
	return i.ToDatasetDicomStoreOutputWithContext(context.Background())
}

func (i *DatasetDicomStore) ToDatasetDicomStoreOutputWithContext(ctx context.Context) DatasetDicomStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetDicomStoreOutput)
}

type DatasetDicomStoreOutput struct {
	*pulumi.OutputState
}

func (DatasetDicomStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetDicomStore)(nil))
}

func (o DatasetDicomStoreOutput) ToDatasetDicomStoreOutput() DatasetDicomStoreOutput {
	return o
}

func (o DatasetDicomStoreOutput) ToDatasetDicomStoreOutputWithContext(ctx context.Context) DatasetDicomStoreOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatasetDicomStoreOutput{})
}
