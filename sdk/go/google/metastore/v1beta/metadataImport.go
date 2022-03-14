// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new MetadataImport in a given project and location.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type MetadataImport struct {
	pulumi.CustomResourceState

	// The time when the metadata import was started.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Immutable. A database dump from a pre-existing metastore's database.
	DatabaseDump DatabaseDumpResponseOutput `pulumi:"databaseDump"`
	// The description of the metadata import.
	Description pulumi.StringOutput `pulumi:"description"`
	// The time when the metadata import finished.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// Immutable. The relative resource name of the metadata import, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}/metadataImports/{metadata_import_id}.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current state of the metadata import.
	State pulumi.StringOutput `pulumi:"state"`
	// The time when the metadata import was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewMetadataImport registers a new resource with the given unique name, arguments, and options.
func NewMetadataImport(ctx *pulumi.Context,
	name string, args *MetadataImportArgs, opts ...pulumi.ResourceOption) (*MetadataImport, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MetadataImportId == nil {
		return nil, errors.New("invalid value for required argument 'MetadataImportId'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	var resource MetadataImport
	err := ctx.RegisterResource("google-native:metastore/v1beta:MetadataImport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetadataImport gets an existing MetadataImport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetadataImport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetadataImportState, opts ...pulumi.ResourceOption) (*MetadataImport, error) {
	var resource MetadataImport
	err := ctx.ReadResource("google-native:metastore/v1beta:MetadataImport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetadataImport resources.
type metadataImportState struct {
}

type MetadataImportState struct {
}

func (MetadataImportState) ElementType() reflect.Type {
	return reflect.TypeOf((*metadataImportState)(nil)).Elem()
}

type metadataImportArgs struct {
	// Immutable. A database dump from a pre-existing metastore's database.
	DatabaseDump *DatabaseDump `pulumi:"databaseDump"`
	// The description of the metadata import.
	Description *string `pulumi:"description"`
	Location    *string `pulumi:"location"`
	// Required. The ID of the metadata import, which is used as the final component of the metadata import's name.This value must be between 1 and 64 characters long, begin with a letter, end with a letter or number, and consist of alpha-numeric ASCII characters or hyphens.
	MetadataImportId string `pulumi:"metadataImportId"`
	// Immutable. The relative resource name of the metadata import, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}/metadataImports/{metadata_import_id}.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Optional. A request ID. Specify a unique request ID to allow the server to ignore the request if it has completed. The server will ignore subsequent requests that provide a duplicate request ID for at least 60 minutes after the first request.For example, if an initial request times out, followed by another request with the same request ID, the server ignores the second request to prevent the creation of duplicate commitments.The request ID must be a valid UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier#Format) A zero UUID (00000000-0000-0000-0000-000000000000) is not supported.
	RequestId *string `pulumi:"requestId"`
	ServiceId string  `pulumi:"serviceId"`
}

// The set of arguments for constructing a MetadataImport resource.
type MetadataImportArgs struct {
	// Immutable. A database dump from a pre-existing metastore's database.
	DatabaseDump DatabaseDumpPtrInput
	// The description of the metadata import.
	Description pulumi.StringPtrInput
	Location    pulumi.StringPtrInput
	// Required. The ID of the metadata import, which is used as the final component of the metadata import's name.This value must be between 1 and 64 characters long, begin with a letter, end with a letter or number, and consist of alpha-numeric ASCII characters or hyphens.
	MetadataImportId pulumi.StringInput
	// Immutable. The relative resource name of the metadata import, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}/metadataImports/{metadata_import_id}.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Optional. A request ID. Specify a unique request ID to allow the server to ignore the request if it has completed. The server will ignore subsequent requests that provide a duplicate request ID for at least 60 minutes after the first request.For example, if an initial request times out, followed by another request with the same request ID, the server ignores the second request to prevent the creation of duplicate commitments.The request ID must be a valid UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier#Format) A zero UUID (00000000-0000-0000-0000-000000000000) is not supported.
	RequestId pulumi.StringPtrInput
	ServiceId pulumi.StringInput
}

func (MetadataImportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metadataImportArgs)(nil)).Elem()
}

type MetadataImportInput interface {
	pulumi.Input

	ToMetadataImportOutput() MetadataImportOutput
	ToMetadataImportOutputWithContext(ctx context.Context) MetadataImportOutput
}

func (*MetadataImport) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataImport)(nil)).Elem()
}

func (i *MetadataImport) ToMetadataImportOutput() MetadataImportOutput {
	return i.ToMetadataImportOutputWithContext(context.Background())
}

func (i *MetadataImport) ToMetadataImportOutputWithContext(ctx context.Context) MetadataImportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataImportOutput)
}

type MetadataImportOutput struct{ *pulumi.OutputState }

func (MetadataImportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataImport)(nil)).Elem()
}

func (o MetadataImportOutput) ToMetadataImportOutput() MetadataImportOutput {
	return o
}

func (o MetadataImportOutput) ToMetadataImportOutputWithContext(ctx context.Context) MetadataImportOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetadataImportInput)(nil)).Elem(), &MetadataImport{})
	pulumi.RegisterOutputType(MetadataImportOutput{})
}
