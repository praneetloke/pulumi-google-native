// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a Dataset.
func LookupDataset(ctx *pulumi.Context, args *LookupDatasetArgs, opts ...pulumi.InvokeOption) (*LookupDatasetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDatasetResult
	err := ctx.Invoke("google-native:aiplatform/v1beta1:getDataset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatasetArgs struct {
	DatasetId string  `pulumi:"datasetId"`
	Location  string  `pulumi:"location"`
	Project   *string `pulumi:"project"`
	ReadMask  *string `pulumi:"readMask"`
}

type LookupDatasetResult struct {
	// Timestamp when this Dataset was created.
	CreateTime string `pulumi:"createTime"`
	// The number of DataItems in this Dataset. Only apply for non-structured Dataset.
	DataItemCount string `pulumi:"dataItemCount"`
	// The description of the Dataset.
	Description string `pulumi:"description"`
	// The user-defined name of the Dataset. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName string `pulumi:"displayName"`
	// Customer-managed encryption key spec for a Dataset. If set, this Dataset and all sub-resources of this Dataset will be secured by this key.
	EncryptionSpec GoogleCloudAiplatformV1beta1EncryptionSpecResponse `pulumi:"encryptionSpec"`
	// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
	Etag string `pulumi:"etag"`
	// The labels with user-defined metadata to organize your Datasets. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one Dataset (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "aiplatform.googleapis.com/" and are immutable. Following system labels exist for each Dataset: * "aiplatform.googleapis.com/dataset_metadata_schema": output only, its value is the metadata_schema's title.
	Labels map[string]string `pulumi:"labels"`
	// Additional information about the Dataset.
	Metadata interface{} `pulumi:"metadata"`
	// The resource name of the Artifact that was created in MetadataStore when creating the Dataset. The Artifact resource name pattern is `projects/{project}/locations/{location}/metadataStores/{metadata_store}/artifacts/{artifact}`.
	MetadataArtifact string `pulumi:"metadataArtifact"`
	// Points to a YAML file stored on Google Cloud Storage describing additional information about the Dataset. The schema is defined as an OpenAPI 3.0.2 Schema Object. The schema files that can be used here are found in gs://google-cloud-aiplatform/schema/dataset/metadata/.
	MetadataSchemaUri string `pulumi:"metadataSchemaUri"`
	// The resource name of the Dataset.
	Name string `pulumi:"name"`
	// All SavedQueries belong to the Dataset will be returned in List/Get Dataset response. The annotation_specs field will not be populated except for UI cases which will only use annotation_spec_count. In CreateDataset request, a SavedQuery is created together if this field is set, up to one SavedQuery can be set in CreateDatasetRequest. The SavedQuery should not contain any AnnotationSpec.
	SavedQueries []GoogleCloudAiplatformV1beta1SavedQueryResponse `pulumi:"savedQueries"`
	// Timestamp when this Dataset was last updated.
	UpdateTime string `pulumi:"updateTime"`
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
	Location  pulumi.StringInput    `pulumi:"location"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
	ReadMask  pulumi.StringPtrInput `pulumi:"readMask"`
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

// Timestamp when this Dataset was created.
func (o LookupDatasetResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The number of DataItems in this Dataset. Only apply for non-structured Dataset.
func (o LookupDatasetResultOutput) DataItemCount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.DataItemCount }).(pulumi.StringOutput)
}

// The description of the Dataset.
func (o LookupDatasetResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.Description }).(pulumi.StringOutput)
}

// The user-defined name of the Dataset. The name can be up to 128 characters long and can consist of any UTF-8 characters.
func (o LookupDatasetResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Customer-managed encryption key spec for a Dataset. If set, this Dataset and all sub-resources of this Dataset will be secured by this key.
func (o LookupDatasetResultOutput) EncryptionSpec() GoogleCloudAiplatformV1beta1EncryptionSpecResponseOutput {
	return o.ApplyT(func(v LookupDatasetResult) GoogleCloudAiplatformV1beta1EncryptionSpecResponse {
		return v.EncryptionSpec
	}).(GoogleCloudAiplatformV1beta1EncryptionSpecResponseOutput)
}

// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
func (o LookupDatasetResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The labels with user-defined metadata to organize your Datasets. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one Dataset (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "aiplatform.googleapis.com/" and are immutable. Following system labels exist for each Dataset: * "aiplatform.googleapis.com/dataset_metadata_schema": output only, its value is the metadata_schema's title.
func (o LookupDatasetResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatasetResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Additional information about the Dataset.
func (o LookupDatasetResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDatasetResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// The resource name of the Artifact that was created in MetadataStore when creating the Dataset. The Artifact resource name pattern is `projects/{project}/locations/{location}/metadataStores/{metadata_store}/artifacts/{artifact}`.
func (o LookupDatasetResultOutput) MetadataArtifact() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.MetadataArtifact }).(pulumi.StringOutput)
}

// Points to a YAML file stored on Google Cloud Storage describing additional information about the Dataset. The schema is defined as an OpenAPI 3.0.2 Schema Object. The schema files that can be used here are found in gs://google-cloud-aiplatform/schema/dataset/metadata/.
func (o LookupDatasetResultOutput) MetadataSchemaUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.MetadataSchemaUri }).(pulumi.StringOutput)
}

// The resource name of the Dataset.
func (o LookupDatasetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.Name }).(pulumi.StringOutput)
}

// All SavedQueries belong to the Dataset will be returned in List/Get Dataset response. The annotation_specs field will not be populated except for UI cases which will only use annotation_spec_count. In CreateDataset request, a SavedQuery is created together if this field is set, up to one SavedQuery can be set in CreateDatasetRequest. The SavedQuery should not contain any AnnotationSpec.
func (o LookupDatasetResultOutput) SavedQueries() GoogleCloudAiplatformV1beta1SavedQueryResponseArrayOutput {
	return o.ApplyT(func(v LookupDatasetResult) []GoogleCloudAiplatformV1beta1SavedQueryResponse { return v.SavedQueries }).(GoogleCloudAiplatformV1beta1SavedQueryResponseArrayOutput)
}

// Timestamp when this Dataset was last updated.
func (o LookupDatasetResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatasetResultOutput{})
}